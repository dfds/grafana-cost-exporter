package metrics

import (
	"fmt"
	"go.dfds.cloud/grafana-cost-exporter/internal/loki_client"
	"go.dfds.cloud/grafana-cost-exporter/internal/prom_client"
	"log"
	"os"
	"strconv"
	"time"
)

type Gatherer struct {
	promClient *prom_client.Client
	lokiClient *loki_client.Client
}

func NewGatherer(lokiClient *loki_client.Client, promClient *prom_client.Client) *Gatherer {
	return &Gatherer{lokiClient: lokiClient, promClient: promClient}
}

type MetricKey string
type LabelKey string
type CapabilityId string

const MetricKeyLokiIngest MetricKey = "loki.ingest"
const MetricKeyLokiRetention MetricKey = "loki.retention"

type MetricData struct {
	Time  int64
	Value float64
	Cost  float64
}

type AllMetricsResponse struct {
	Days30 map[MetricKey]map[LabelKey]*MetricData
}

func (g *Gatherer) GetAllMetrics() *AllMetricsResponse {
	dataStore30Days := make(map[MetricKey]map[LabelKey]*MetricData)
	now := time.Now()

	// Loki ingest
	volumeResp, err := g.lokiClient.Volume(`{namespace=~".+"}`, now.Add(time.Duration(-(24*30))*time.Hour).Unix(), now.Unix())
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := dataStore30Days[MetricKeyLokiIngest]; !ok {
		dataStore30Days[MetricKeyLokiIngest] = make(map[LabelKey]*MetricData)
	}

	for _, val := range volumeResp.Data.Result {
		bytesValue := val.Value[1].(string)
		bytes, err := strconv.ParseFloat(bytesValue, 64)
		if err != nil {
			log.Fatal(err)
		}

		bytesToGib := bytes / 1024 / 1024 / 1024
		cost := bytesToGib * 0.23

		dataStore30Days[MetricKeyLokiIngest][LabelKey(val.Metric.Namespace)] = &MetricData{
			Time:  now.Unix(),
			Value: bytesToGib,
			Cost:  cost,
		}
	}

	// Loki retention

	if _, ok := dataStore30Days[MetricKeyLokiRetention]; !ok {
		dataStore30Days[MetricKeyLokiRetention] = make(map[LabelKey]*MetricData)
	}

	for namespace, data := range dataStore30Days[MetricKeyLokiIngest] {
		cost := data.Value * 0.06
		dataStore30Days[MetricKeyLokiRetention][namespace] = &MetricData{
			Time:  now.Unix(),
			Value: data.Value,
			Cost:  cost,
		}
	}

	for metricKey, metricData := range dataStore30Days {
		fmt.Printf("\n:: %s ::\n\n", metricKey)

		for namespace, data := range metricData {
			fmt.Printf("%s - %f GiB - %f USD\n", namespace, data.Value, data.Cost)
		}
	}

	os.Exit(1)

	return &AllMetricsResponse{
		Days30: dataStore30Days,
	}
}
