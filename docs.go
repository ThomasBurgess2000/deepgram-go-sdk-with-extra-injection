// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

/*
	This SDK provides Go library for performing Prerecorded and Live/Streaming operations
	on the Deepgram.com Platform.

	GitHub repo: https://github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection
	Go SDK Examples: https://github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/tree/main/examples

	Deepgram Platform API reference: https://developers.deepgram.com/reference
	Documentation: https://developers.deepgram.com/docs

	The main entry point for this SDK is:
	1. pkg/client/live - contains the SDK objects, functions, etc for performing Live/Stream operations
	2. pkg/client/prerecorded - contains the SDK objects, functions, etc for performing operations on Prerecorded media
*/

package sdk

import (
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/analyze"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/listen"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/manage"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/speak"

	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/analyze/v1"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/listen/v1/rest"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/listen/v1/websocket"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/manage/v1"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/speak/v1/rest"
	_ "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/speak/v1/websocket"
)
