package metrics

import (
	"go.dfds.cloud/grafana-cost-exporter/internal/loki_client"
	"go.dfds.cloud/grafana-cost-exporter/internal/prom_client"
	"os"
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
const MetricKeyPromSamplesScraped MetricKey = "prom.samples_scraped"

type MetricData struct {
	Time  int64
	Value float64
	Cost  float64
}

type MetricDataI64 struct {
	Time  int64
	Value int64
	Cost  float64
}

type AllMetricsResponse struct {
	Days30Loki map[MetricKey]map[LabelKey]*MetricData
	Days30Prom map[MetricKey]map[LabelKey][]*MetricDataI64
}

func (g *Gatherer) GetAllMetrics() *AllMetricsResponse {
	// Loki
	loki30DaysResult := g.gatherLoki()
	prom30DaysResult := g.gatherProm()

	os.Exit(1)

	return &AllMetricsResponse{
		Days30Loki: loki30DaysResult,
		Days30Prom: prom30DaysResult,
	}
}
