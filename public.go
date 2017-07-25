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
type Market struct {
	MarketCurrency     string  `json:"MarketCurrency"`
	BaseCurrency       string  `json:"BaseCurrency"`
	MarketCurrencyLong string  `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string  `json:"BaseCurrencyLong"`
	MinTradeSize       float64 `json:"MinTradeSize"`
	MarketName         string  `json:"MarketName"`
	IsActive           bool    `json:"IsActive"`
	Created            string  `json:"Created"`
	Notice             string  `json:"Notice"`
	IsSponsored        bool    `json:"IsSponsored"`
	LogoURL            string  `json:"LogoUrl"`
}

// Currency stores the API result about currencies
type Currency struct {
	Currency        string  `json:"Currency"`
	CurrencyLong    string  `json:"CurrencyLong"`
	MinConfirmation int     `json:"MinConfirmation"`
	TxFee           float64 `json:"TxFee"`
	IsActive        bool    `json:"IsActive"`
	CoinType        string  `json:"CoinType"`
	BaseAddress     string  `json:"BaseAddress"`
	Notice          string  `json:"Notice"`
}

// MarketSummary stores the API result about market summaries
type MarketSummary struct {
	MarketName     string  `json:"MarketName"`
	High           float64 `json:"High"`
	Low            float64 `json:"Low"`
	Volume         float64 `json:"Volume"`
	Last           float64 `json:"Last"`
	BaseVolume     float64 `json:"BaseVolume"`
	TimeStamp      string  `json:"TimeStamp"`
	Bid            float64 `json:"Bid"`
	Ask            float64 `json:"Ask"`
	OpenBuyOrders  int     `json:"OpenBuyOrders"`
	OpenSellOrders int     `json:"OpenSellOrders"`
	PrevDay        float64 `json:"PrevDay"`
	Created        string  `json:"Created"`
}

// Ticker stores current tick value for a market
type Ticker struct {
	Bid  float64 `json:"Bid"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}

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

// GetCurrencies gets all supported currencies
func (p *PublicService) GetCurrencies() ([]Currency, error) {
	url := "public/getcurrencies"
	req, err := p.client.NewRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	currencies := new([]Currency)
	_, err = p.client.Do(req, currencies)

	return *currencies, err
}

// GetMarketSummaries gets all supported currencies
func (p *PublicService) GetMarketSummaries() ([]MarketSummary, error) {
	url := "public/getmarketsummaries"
	req, err := p.client.NewRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	marketSummaries := new([]MarketSummary)
	_, err = p.client.Do(req, marketSummaries)

	return *marketSummaries, err
}

// GetMarketSummary gets last 24h summary of all active exchanges
func (p *PublicService) GetMarketSummary(market string) ([]MarketSummary, error) {
	url := "public/getmarketsummary?market=" + market
	req, err := p.client.NewRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	marketSummary := new([]MarketSummary)
	_, err = p.client.Do(req, marketSummary)

	return *marketSummary, err
}

// GetTicker returns current tick values for a market
func (p *PublicService) GetTicker(market string) (Ticker, error) {
	url := "public/getticker?market=" + market
	req, err := p.client.NewRequest("GET", url, "")
	if err != nil {
		return Ticker{}, err
	}

	ticker := new(Ticker)
	_, err = p.client.Do(req, ticker)

	return *ticker, err
}
