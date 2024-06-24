package prom_client

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	endpoint string
	auth     *Auth
	http     *http.Client
}

type Auth struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewClient(endpoint string, auth *Auth) *Client {
	return &Client{
		endpoint: endpoint,
		http:     http.DefaultClient,
		auth:     auth,
	}
}

func (c *Client) PrepareHttpRequest(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", c.auth.Username, c.auth.Token)))))
}

func (c *Client) Query(query string, time float64) (*QueryResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/query", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	c.PrepareHttpRequest(req)

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

func (c *Client) QueryRange(query string, step string, start int64, end int64) (*QueryRangeResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/query_range", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	c.PrepareHttpRequest(req)

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

	var payload *QueryRangeResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) Series(match string, start int64, end int64) (*SeriesResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/series", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	c.PrepareHttpRequest(req)

	queryValues := req.URL.Query()
	queryValues.Set("match[]", match)
	queryValues.Set("start", fmt.Sprintf("%d", start))
	queryValues.Set("end", fmt.Sprintf("%d", end))
	req.URL.RawQuery = queryValues.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		fmt.Println(resp.Header)
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *SeriesResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) ListSeries() (*ListSeriesResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/label/__name__/values", c.endpoint), nil)
	if err != nil {
		return nil, err
	}

	c.PrepareHttpRequest(req)

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

	c.PrepareHttpRequest(req)

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

func ResultToVector(data [][]interface{}) ([]Vector, error) {
	var payload []Vector

	for _, x := range data {
		timestamp := x[0].(float64)
		value := x[1].(string)

		newVec := Vector{Value: VectorValue{}}
		newVec.Value.Value = value
		newVec.Value.Time = int64(timestamp)
		payload = append(payload, newVec)
	}

	return payload, nil
}
