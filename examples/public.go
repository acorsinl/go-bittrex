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
}
