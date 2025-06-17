// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

// This package contains the code for the Token APIs in the Deepgram Auth API
package auth

import (
	authv1 "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/auth/v1"
	interfaces "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/interfaces/v1"
)

const (
	PackageVersion = authv1.PackageVersion
)

// Alias
type Client = authv1.Client

/*
NewWithDefaults creates a new analyze/read client with all default options

Notes:
  - The Deepgram API KEY is read from the environment variable DEEPGRAM_API_KEY
*/
func NewWithDefaults() *Client {
	return authv1.NewWithDefaults()
}

/*
New creates a new analyze/read client with the specified options

Input parameters:
- [Optional] apiKey: string containing the Deepgram API key. If empty, the API key is read from the environment variable DEEPGRAM_API_KEY
- [Optional] options: ClientOptions which allows overriding things like hostname, version of the API, etc.
*/
func New(apiKey string, options *interfaces.ClientOptions) *Client {
	return authv1.New(apiKey, options)
}
