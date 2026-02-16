package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Video struct {
	// Number of audio bits per sample.
	AudioBitsPerSample nullable.Type[int64] `json:"audioBitsPerSample,omitempty"`

	// Number of audio channels.
	AudioChannels nullable.Type[int64] `json:"audioChannels,omitempty"`

	// Name of the audio format (AAC, MP3, etc.).
	AudioFormat nullable.Type[string] `json:"audioFormat,omitempty"`

	// Number of audio samples per second.
	AudioSamplesPerSecond nullable.Type[int64] `json:"audioSamplesPerSecond,omitempty"`

	// Bit rate of the video in bits per second.
	Bitrate nullable.Type[int64] `json:"bitrate,omitempty"`

	// Duration of the file in milliseconds.
	Duration nullable.Type[int64] `json:"duration,omitempty"`

	// 'Four character code' name of the video format.
	FourCC nullable.Type[string] `json:"fourCC,omitempty"`

	// Height of the video, in pixels.
	Height nullable.Type[int64] `json:"height,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Width of the video, in pixels.
	Width nullable.Type[int64] `json:"width,omitempty"`
}
