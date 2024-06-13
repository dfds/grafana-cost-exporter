package loki_client

import (
	"encoding/base64"
	"encoding/json"
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

func (c *Client) LabelValues(label string) (*LabelValuesResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/label/%s/values", c.endpoint, label), nil)
	if err != nil {
		return nil, err
	}
	c.PrepareHttpRequest(req)

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *LabelValuesResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) Volume(query string, start int64, end int64) (*VolumeResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/index/volume", c.endpoint), nil)
	if err != nil {
		return nil, err
	}
	c.PrepareHttpRequest(req)

	queryValues := req.URL.Query()
	queryValues.Set("query", query)
	queryValues.Set("start", fmt.Sprintf("%d", start))
	queryValues.Set("end", fmt.Sprintf("%d", end))
	req.URL.RawQuery = queryValues.Encode()

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.StatusCode)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var payload *VolumeResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}

func (c *Client) VolumeRange(query string, start int64, end int64) (*VolumeRangeResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/index/volume_range", c.endpoint), nil)
	if err != nil {
		return nil, err
	}
	c.PrepareHttpRequest(req)

	queryValues := req.URL.Query()
	queryValues.Set("query", query)
	queryValues.Set("start", fmt.Sprintf("%d", start))
	queryValues.Set("end", fmt.Sprintf("%d", end))
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

	var payload *VolumeRangeResponse

	err = json.Unmarshal(data, &payload)

	return payload, err
}
