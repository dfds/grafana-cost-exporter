package prom_client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	endpoint string
	http     *http.Client
}

func NewClient(endpoint string) *Client {
	return &Client{
		endpoint: endpoint,
		http:     http.DefaultClient,
	}
}

func (c *Client) Query(query string, time float64) (*QueryResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/query", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	queryValues := req.URL.Query()
	queryValues.Set("query", query)
	queryValues.Set("time", fmt.Sprintf("%f", time))
	req.URL.RawQuery = queryValues.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *QueryResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) QueryRange(query string, step string, start int64, end int64) (*QueryResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/query_range", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	queryValues := req.URL.Query()
	queryValues.Set("query", query)
	queryValues.Set("start", fmt.Sprintf("%d", start))
	queryValues.Set("end", fmt.Sprintf("%d", end))
	queryValues.Set("step", step)
	req.URL.RawQuery = queryValues.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *QueryResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) ListSeries() (*ListSeriesResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/label/__name__/values", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	queryValues := req.URL.Query()
	req.URL.RawQuery = queryValues.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *ListSeriesResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) ListLabels() (*ListLabelsResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/labels", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	queryValues := req.URL.Query()
	req.URL.RawQuery = queryValues.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *ListLabelsResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func ResultToVector(data []interface{}) ([]Vector, error) {
	var casted []vectorMidParse
	for _, d := range data {
		var deserialised vectorMidParse
		serialised, err := json.Marshal(d)
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(serialised, &deserialised)
		if err != nil {
			log.Fatal(err)
		}

		casted = append(casted, deserialised)
	}
	var payload []Vector

	for _, vec := range casted {
		newVec := Vector{
			Metric: VectorMetricLabel{
				Instance: vec.Metric.Instance,
				Job:      vec.Metric.Job,
				KafkaID:  vec.Metric.KafkaID,
				Topic:    vec.Metric.Topic,
			},
			Value: VectorValue{},
		}
		newVec.Value.Time = vec.Value[0].(float64)
		newVec.Value.Value = vec.Value[1].(string)
		payload = append(payload, newVec)
	}

	return payload, nil
}
