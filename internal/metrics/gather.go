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

		fmt.Printf("%s - %f GiB - %f USD\n", val.Metric.Namespace, bytesToGib, cost)

		dataStore30Days[MetricKeyLokiIngest][LabelKey(val.Metric.Namespace)] = &MetricData{
			Time:  now.Unix(),
			Value: bytesToGib,
			Cost:  cost,
		}

	}

	os.Exit(1)

	//for _, metricKey := range ConfluentMetrics {
	//	// check if metricKey exists in dataStore30Days and dataStorePerDay, if not, init
	//	if _, ok := dataStore30Days[metricKey]; !ok {
	//		dataStore30Days[metricKey] = make(map[ClusterId]map[string]float64)
	//	}
	//	if _, ok := dataStorePerDay[metricKey]; !ok {
	//		dataStorePerDay[metricKey] = make(map[ClusterId]map[string][]MetricData)
	//	}
	//
	//	baseQuery := fmt.Sprintf("sum_over_time(%s[1d]", metricKey)
	//	for i := 0; i <= 30; i++ {
	//		var query string = baseQuery
	//		timestamp := now
	//		if i != 0 {
	//			query = fmt.Sprintf("%s offset %dd)", baseQuery, i)
	//			timestamp = now.Add(time.Duration(i*24) * -time.Hour)
	//			if metricKey == ConfluentKafkaServerRetainedBytes {
	//				query = fmt.Sprintf("%s offset %dd", ConfluentKafkaServerRetainedBytes, i)
	//			}
	//		} else {
	//			query = fmt.Sprintf("%s)", query)
	//			if metricKey == ConfluentKafkaServerRetainedBytes {
	//				query = fmt.Sprintf("%s offset 1h", ConfluentKafkaServerRetainedBytes) // not perfect, WIP
	//			}
	//		}
	//
	//		fmt.Println(query)
	//
	//		queryResp, err := g.client.Query(query, float64(now.Unix()))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		data, err := prom_client.ResultToVector(queryResp.Data.Result)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		for _, vector := range data {
	//			if _, ok := dataStorePerDay[metricKey][ClusterId(vector.Metric.KafkaID)]; !ok {
	//				dataStorePerDay[metricKey][ClusterId(vector.Metric.KafkaID)] = map[string][]MetricData{}
	//			}
	//			if _, ok := dataStorePerDay[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic]; !ok {
	//				dataStorePerDay[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic] = []MetricData{}
	//			}
	//			if _, ok := dataStore30Days[metricKey][ClusterId(vector.Metric.KafkaID)]; !ok {
	//				dataStore30Days[metricKey][ClusterId(vector.Metric.KafkaID)] = map[string]float64{}
	//			}
	//
	//			f64, _ := strconv.ParseFloat(vector.Value.Value, 64)
	//
	//			dataStorePerDay[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic] = append(dataStorePerDay[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic], MetricData{
	//				Time:  float64(timestamp.Unix()),
	//				Value: f64,
	//			})
	//
	//			if metricKey == ConfluentKafkaServerRetainedBytes {
	//				if i == 0 {
	//					dataStore30Days[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic] = f64
	//				}
	//			} else {
	//				dataStore30Days[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic] = dataStore30Days[metricKey][ClusterId(vector.Metric.KafkaID)][vector.Metric.Topic] + f64
	//			}
	//
	//		}
	//	}
	//}

	return &AllMetricsResponse{
		Days30: dataStore30Days,
	}
}
