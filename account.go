// Copyright 2017 the go-bittrex authors. All rights reserved.
//
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bittrex provides a client to use bittres HTTP API
package bittrex

// AccountService handles communication with the public services
// exposed by bittrex API
type AccountService struct {
	client *Client
}

// Balance stores API returned balances
type Balance struct {
	Currency      string  `json:"Currency"`
	Balance       float64 `json:"Balance"`
	Available     float64 `json:"Available"`
	Pending       float64 `json:"Pending"`
	CryptoAddress string  `json:"CryptoAddress"`
	Requested     bool    `json:"Requested"`
	UUID          string  `json:"Uuid"`
}

// GetBalances gets all balances from your account
func (a *AccountService) GetBalances() ([]Balance, error) {
	url := "account/getbalances?apikey=" + a.client.APIKey
	req, err := a.client.NewRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	balance := new([]Balance)
	_, err = a.client.Do(req, balance)

	return *balance, err
}
