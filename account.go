// Copyright 2017 the go-bittrex authors. All rights reserved.
//
// Use of this code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bittrex provides a client to use bittres HTTP API
package bittrex

// PublicService handles communication with the public services
// exposed by bittrex API
type AccountService struct {
	client *Client
}
