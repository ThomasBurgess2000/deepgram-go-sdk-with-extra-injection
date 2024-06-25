// Copyright 2023-2024 Deepgram SDK contributors. All Rights Reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// SPDX-License-Identifier: MIT

package interfacesv1

// These are the message types that can be received from the live API
const (
	// message types
	TypeOpenResponse          string = "Open"
	TypeMessageResponse       string = "Results"
	TypeMetadataResponse      string = "Metadata"
	TypeUtteranceEndResponse  string = "UtteranceEnd"
	TypeSpeechStartedResponse string = "SpeechStarted"
	TypeFinalizeResponse      string = "Finalize"
	TypeCloseStreamResponse   string = "CloseStream"
	TypeCloseResponse         string = "Close"

	// Error type
	TypeErrorResponse string = "Error"
)
