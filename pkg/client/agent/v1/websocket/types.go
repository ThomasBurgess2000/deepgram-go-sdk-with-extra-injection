// Copyright 2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package websocketv1

import (
	"context"

	msginterface "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/api/agent/v1/websocket/interfaces"
	common "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/common/v1"
	commoninterfaces "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/common/v1/interfaces"
	interfaces "github.com/ThomasBurgess2000/deepgram-go-sdk-with-extra-injection/v3/pkg/client/interfaces"
)

// client messages
type SettingsOptions interfaces.SettingsOptions
type UpdatePrompt msginterface.UpdatePrompt
type UpdateSpeak msginterface.UpdateSpeak
type InjectAgentMessage msginterface.InjectAgentMessage
type FunctionCallResponse msginterface.FunctionCallResponse
type KeepAlive msginterface.KeepAlive

// WSChannel is a struct representing the websocket client connection using channels
type WSChannel struct {
	*common.WSClient
	ctx       context.Context
	ctxCancel context.CancelFunc

	cOptions *interfaces.ClientOptions
	tOptions *interfaces.SettingsOptions

	chans  []*msginterface.AgentMessageChan
	router *commoninterfaces.Router
}
