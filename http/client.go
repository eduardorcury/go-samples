package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type ClientBuilder struct {
	httpClient    *http.Client
	url, method   string
	gatewayID     string
	queryParams   map[string]string
	customHeaders map[string]string
	body          any
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		httpClient:    http.DefaultClient,
		method:        http.MethodGet,
		queryParams:   make(map[string]string),
		customHeaders: make(map[string]string),
	}
}

func (c *ClientBuilder) WithUrl(url string) *ClientBuilder {
	c.url = url
	return c
}

func (c *ClientBuilder) WithMethod(method string) *ClientBuilder {
	c.method = method
	return c
}

func (c *ClientBuilder) WithGatewayID(gatewayID string) *ClientBuilder {
	c.gatewayID = gatewayID
	return c
}

func (c *ClientBuilder) WithQueryParams(queryParams map[string]string) *ClientBuilder {
	for k, v := range queryParams {
		c.queryParams[k] = v
	}
	return c
}

func (c *ClientBuilder) WithCustomHeaders(customHeaders map[string]string) *ClientBuilder {
	for k, v := range customHeaders {
		c.customHeaders[k] = v
	}
	return c
}

func (c *ClientBuilder) WithBody(body any) *ClientBuilder {
	c.body = body
	return c
}

func (c *ClientBuilder) Do() (*http.Response, error) {
	parsedUrl, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}

	q := parsedUrl.Query()
	for k, v := range c.queryParams {
		q.Set(k, v)
	}
	parsedUrl.RawQuery = q.Encode()

	var bodyReader *bytes.Reader
	if c.body != nil {
		b, err := json.Marshal(c.body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(b)
	} else {
		bodyReader = bytes.NewReader(nil)
	}

	req, err := http.NewRequest(c.method, parsedUrl.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	for k, v := range c.customHeaders {
		req.Header.Set(k, v)
	}
	if c.gatewayID != "" {
		req.Header.Set("X-Gateway-ID", c.gatewayID)
	}

	return c.httpClient.Do(req)
}
