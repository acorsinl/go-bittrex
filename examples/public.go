// Copyright 2017 the go-bittrex authors. All rights reserved.
//
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/acorsinl/go-bittrex"
)

func main() {
	client := bittrex.NewClient(nil)
	client.APIKey = ""
	client.APISecret = ""

	markets, err := client.Public.GetMarkets()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, m := range markets {
		fmt.Println(m)
	}

	currencies, err := client.Public.GetCurrencies()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, m := range currencies {
		fmt.Println(m)
	}

	marketSummaries, err := client.Public.GetMarketSummaries()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, m := range marketSummaries {
		fmt.Println(m)
	}

	ticker, err := client.Public.GetTicker("ETH-LTC")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(ticker)

	marketSummary, err := client.Public.GetMarketSummary("ETH-LTC")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, m := range marketSummary {
		fmt.Println(m)
	}

	marketHistory, err := client.Public.GetMarketHistory("ETH-LTC")
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, m := range marketHistory {
		fmt.Println(m)
	}
}
