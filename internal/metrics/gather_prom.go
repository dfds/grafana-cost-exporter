package metrics

import (
	"context"
	"fmt"
	"go.dfds.cloud/grafana-cost-exporter/internal/prom_client"
	"golang.org/x/sync/semaphore"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

func (g *Gatherer) gatherProm() map[MetricKey]map[LabelKey][]*MetricDataI64 {
	dataStore30Days := make(map[MetricKey]map[LabelKey][]*MetricDataI64)
	now := time.Now()

	fmt.Println(now.Unix())
	fmt.Println(now.Add(time.Duration(-(24 * 30)) * time.Hour).Unix())

	//start := now.Add(time.Duration(-(24 * 30)) * time.Hour).Unix()
	//end := now
	//start := int64(1716625158) // 24th of May
	//end := int64(1719217158) // 24th of June
	start := int64(1714525883) // 1st of May
	end := int64(1717031483)   // 30th of May

	// Prom data
	volumeResp, err := g.promClient.QueryRange(`sum by (namespace) (scrape_samples_scraped)`, "1d", start, end)
	if err != nil {
		log.Fatal(err)
	}

	seriesResp, err := g.promClient.ListSeries()
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := dataStore30Days[MetricKeyPromSamplesScraped]; !ok {
		dataStore30Days[MetricKeyPromSamplesScraped] = make(map[LabelKey][]*MetricDataI64)
	}

	wg := &sync.WaitGroup{}
	sem := semaphore.NewWeighted(80)
	ctx := context.Background()
	mu := &sync.Mutex{}

	var totalForFirst int64 = 0
	for _, val := range volumeResp.Data.Result {
		wg.Add(1)
		sem.Acquire(ctx, 1)
		innerVal := val
		go func() {
			defer sem.Release(1)
			defer wg.Done()
			vec, err := prom_client.ResultToVector(innerVal.Values)
			if err != nil {
				log.Fatal(err)
			}

			if innerVal.Metric.Namespace == "" { // Skip if it doesn't contain the namespace label
				return
			}

			fmt.Println(innerVal.Metric.Namespace)
			for i, metric := range vec {
				fmt.Printf("%s - %s\n", time.Unix(metric.Value.Time, 0).Format(time.RFC3339), metric.Value.Value)
				value, err := strconv.ParseInt(metric.Value.Value, 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				if i == 0 {
					totalForFirst = totalForFirst + value
				}

				// lock
				mu.Lock()
				dataStore30Days[MetricKeyPromSamplesScraped][LabelKey(innerVal.Metric.Namespace)] = append(dataStore30Days[MetricKeyPromSamplesScraped][LabelKey(innerVal.Metric.Namespace)], &MetricDataI64{
					Time:  metric.Value.Time,
					Value: value,
					Cost:  0,
				})
				mu.Unlock()
			}
		}()

	}

	wg.Wait()

	sem = semaphore.NewWeighted(80)
	mu = &sync.Mutex{}
	wg = &sync.WaitGroup{}
	var totalTimeSeries int64 = 0
	timeSeriesValues := make(map[string]int64)

	for _, series := range seriesResp.Data {
		wg.Add(1)
		sem.Acquire(ctx, 1)
		lookupSeries := series

		go func() {
			defer sem.Release(1)
			defer wg.Done()

			resp, err := g.promClient.Series(lookupSeries, start, end)
			if err != nil {
				log.Fatal(err)
			}

			mu.Lock()
			totalTimeSeries = totalTimeSeries + int64(len(resp.Data))
			timeSeriesValues[lookupSeries] = int64(len(resp.Data))
			mu.Unlock()
		}()
	}

	wg.Wait()

	// sanity check numbers
	var ss []kv
	for k, v := range timeSeriesValues {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, v := range ss {
		fmt.Printf("%d - %s\n", v.Value, v.Key)
	}

	fmt.Printf("Total data samples: %d\n", totalForFirst)
	fmt.Printf("Total series count: %d\n", len(seriesResp.Data))
	fmt.Printf("Total time series count: %d\n", totalTimeSeries)

	// calc method #1
	{
		seriesCount := float64(totalTimeSeries)
		dpm := seriesCount
		usage := dpm / 1

		cost := usage * (4.56 / 1000)
		fmt.Printf("Estimated cost calc 01: %.2f\n", cost)

	}

	// calc method #2
	{
		seriesCount := float64(totalTimeSeries)
		seriesCountDivided := seriesCount / 1000

		cost := seriesCount * (4.56 / seriesCountDivided)
		fmt.Printf("Estimated cost calc 02: %.2f\n", cost)
	}

	// calc method #3
	{
		seriesCount := float64(totalForFirst)
		dpm := 1.0

		cost := seriesCount * dpm * (4.56 / 1000)
		fmt.Printf("Estimated cost calc 03: %.2f\n", cost)
	}

	return dataStore30Days
}

type kv struct {
	Key   string
	Value int64
}
