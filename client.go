package alephalpha

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://docs.aleph-alpha.com/api/aleph-alpha-api/

const BaseURL = "https://api.aleph-alpha.com/"

type transport struct {
	apiKey              string
	underlyingTransport http.RoundTripper
}

type Client struct {
	*http.Client
	BaseURL string

	availableModels []AvailableModel
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+t.apiKey)
	return t.underlyingTransport.RoundTrip(req)
}

func NewClient(apiKey string) *Client {
	return &Client{
		Client: &http.Client{
			Transport: &transport{
				apiKey:              apiKey,
				underlyingTransport: http.DefaultTransport,
			},
		},
		BaseURL: BaseURL,
	}
}

func NewClientFromEnv() *Client {
	return NewClient(os.Getenv("AA_TOKEN"))
}

type newable[T any] interface {
	New() T
}

func req[D, T any](c *Client, method string, path string, data D) (T, error) {
	var jsonData io.Reader
	if method == http.MethodGet {
		jsonData = nil
	} else {
		jsonDataS, err := json.Marshal(data)
		if err != nil {
			return *new(T), err
		}
		// log.Println(string(jsonDataS))
		jsonData = bytes.NewReader(jsonDataS)
	}

	req, err := http.NewRequest(method, c.BaseURL+path, jsonData)
	if err != nil {
		return *new(T), err
	}
	res, err := c.Do(req)
	if err != nil {
		return *new(T), err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return *new(T), fmt.Errorf("status: %s", res.Status)
	}

	var result T
	if n, ok := interface{}(result).(newable[T]); ok {
		result = n.New()
	}

	err = json.NewDecoder(res.Body).Decode(&result)

	return result, err
}

func APIVersion() (string, error) {
	resp, err := http.Get(BaseURL + "version")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	v, err := io.ReadAll(resp.Body)
	return string(v), err
}
