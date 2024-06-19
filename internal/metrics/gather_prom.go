package metrics

import (
	"fmt"
	"go.dfds.cloud/grafana-cost-exporter/internal/prom_client"
	"log"
	"strconv"
	"time"
)

func (g *Gatherer) gatherProm() map[MetricKey]map[LabelKey][]*MetricDataI64 {
	dataStore30Days := make(map[MetricKey]map[LabelKey][]*MetricDataI64)
	now := time.Now()

	// Prom data
	volumeResp, err := g.promClient.QueryRange(`sum by (namespace) (scrape_samples_scraped)`, "1d", now.Add(time.Duration(-(24*30))*time.Hour).Unix(), now.Unix())
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := dataStore30Days[MetricKeyPromSamplesScraped]; !ok {
		dataStore30Days[MetricKeyPromSamplesScraped] = make(map[LabelKey][]*MetricDataI64)
	}

	for _, val := range volumeResp.Data.Result {
		vec, err := prom_client.ResultToVector(val.Values)
		if err != nil {
			log.Fatal(err)
		}

		if val.Metric.Namespace == "" { // Skip if it doesn't contain the namespace label
			continue
		}

		fmt.Println(val.Metric.Namespace)
		for _, metric := range vec {
			fmt.Printf("%s - %s\n", time.Unix(metric.Value.Time, 0).Format(time.RFC3339), metric.Value.Value)
			value, err := strconv.ParseInt(metric.Value.Value, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			dataStore30Days[MetricKeyPromSamplesScraped][LabelKey(val.Metric.Namespace)] = append(dataStore30Days[MetricKeyPromSamplesScraped][LabelKey(val.Metric.Namespace)], &MetricDataI64{
				Time:  metric.Value.Time,
				Value: value,
				Cost:  0,
			})
		}

		//bytesValue := val.Value[1].(string)
		//bytes, err := strconv.ParseFloat(bytesValue, 64)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//bytesToGib := bytes / 1024 / 1024 / 1024
		//cost := bytesToGib * 0.23
		//
		//dataStore30Days[MetricKeyLokiIngest][LabelKey(val.Metric.Namespace)] = &MetricData{
		//	Time:  now.Unix(),
		//	Value: bytesToGib,
		//	Cost:  cost,
		//}
	}

	return dataStore30Days
}
