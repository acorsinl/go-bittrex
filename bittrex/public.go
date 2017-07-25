// Copyright 2017 the go-bittrex authors. All rights reserved.
//
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bittrex provides a client to use bittres HTTP API
package bittrex

// PublicService handles communication with the public services
// exposed by bittrex API
type PublicService struct {
	client *Client
}

// Market stores the API result about markets
type Market struct{}

// GetMarkets gets open and available markets at bittrex
func (p *PublicService) GetMarkets() ([]Market, error) {
	url := "public/getMarkets"
	req, err := p.client.NewRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	markets := new([]Market)
	_, err = p.client.Do(req, markets)

	return *markets, err
}
