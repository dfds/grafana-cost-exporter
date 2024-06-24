package prom_client

type Vector struct {
	Metric VectorMetricLabel `json:"metric"`
	Value  VectorValue       `json:"value"`
}

type VectorMetricLabel struct {
	Namespace string `json:"namespace"`
}

type VectorValue struct {
	Time  int64
	Value string
}

type vectorMidParse struct {
	Metric struct {
		Instance string `json:"instance"`
		Job      string `json:"job"`
		KafkaID  string `json:"kafka_id"`
		Topic    string `json:"topic"`
	} `json:"metric"`
	Value []interface{} `json:"value"`
}

type QueryResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string        `json:"resultType"`
		Result     []interface{} `json:"result"`
	} `json:"data"`
}

type QueryRangeResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Namespace string `json:"namespace,omitempty"`
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

type SeriesResponse struct {
	Status string `json:"status"`
	Data   []struct {
		Name      string `json:"__name__"`
		Cluster   string `json:"cluster"`
		Endpoint  string `json:"endpoint"`
		Instance  string `json:"instance"`
		Job       string `json:"job"`
		Name1     string `json:"name"`
		Namespace string `json:"namespace"`
		Pod       string `json:"pod"`
		Service   string `json:"service"`
	} `json:"data"`
}

type ListSeriesResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

type ListLabelsResponse struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}
