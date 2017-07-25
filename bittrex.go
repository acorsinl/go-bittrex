// Copyright 2017 the go-bittrex authors. All rights reserved.
//
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bittrex provides a client to use bittres HTTP API
package bittrex

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// LibraryVersion contains this library current version
	LibraryVersion = "0.1"

	// BaseURL where bittrex API is listening
	BaseURL = "https://bittrex.com/api/v1.1/"

	// UserAgent to be sent in each request by this package
	UserAgent = "github.com/acorsinl/go-bittrex v" + LibraryVersion
)

// Client manages communication with bittrex API
type Client struct {
	// HTTP client to communicate with the API
	client *http.Client

	// BaseURL for API requests
	BaseURL *url.URL

	// UserAgent to be sent in each request by this package
	UserAgent string

	// API Key
	APIKey string

	// API Secret
	APISecret string

	// Services used to access different API sections
	Public *PublicService
	// @todo Market  *MarketService
	Account *AccountService

	Response *Response
}

// Response handles bittrex API response
type Response struct {
	Response *http.Response
	Success  bool
	Message  string
	Result   interface{}
}

// NewRequest creates an API request. A relative URL can be provided in
// URLStr, in which case it is added to the BaseURL.
func (c *Client) NewRequest(method, URLStr, body string) (*http.Request, error) {
	rel, err := url.Parse(URLStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	// @todo: handle non public urls with apikey and apisecret
	req, err := http.NewRequest(method, u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("Accept", "application/json")

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := &Response{Response: res}
	if v != nil {
		r.Result = v
		err = json.NewDecoder(res.Body).Decode(r)
		c.Response = r
	}

	return res, err
}

// NewClient returns a bittrex API client. If a nil http client is provided
// http.DefaultClient will be used
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(BaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
	}
	c.Public = &PublicService{client: c}

	return c
}
