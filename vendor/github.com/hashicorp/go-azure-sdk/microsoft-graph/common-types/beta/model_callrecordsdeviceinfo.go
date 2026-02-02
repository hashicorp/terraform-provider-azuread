package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsDeviceInfo struct {
	// Name of the capture device driver used by the media endpoint.
	CaptureDeviceDriver nullable.Type[string] `json:"captureDeviceDriver,omitempty"`

	// Name of the capture device used by the media endpoint.
	CaptureDeviceName nullable.Type[string] `json:"captureDeviceName,omitempty"`

	// Number of times during the call that the media endpoint detected howling or screeching audio.
	HowlingEventCount nullable.Type[int64] `json:"howlingEventCount,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Average energy level of received audio for audio classified as mono noise or left channel of stereo noise by the
	// media endpoint.
	ReceivedNoiseLevel nullable.Type[int64] `json:"receivedNoiseLevel,omitempty"`

	// Average energy level of received audio for audio classified as mono speech, or left channel of stereo speech by the
	// media endpoint.
	ReceivedSignalLevel nullable.Type[int64] `json:"receivedSignalLevel,omitempty"`

	// Name of the render device driver used by the media endpoint.
	RenderDeviceDriver nullable.Type[string] `json:"renderDeviceDriver,omitempty"`

	// Name of the render device used by the media endpoint.
	RenderDeviceName nullable.Type[string] `json:"renderDeviceName,omitempty"`

	// Average energy level of sent audio for audio classified as mono noise or left channel of stereo noise by the media
	// endpoint.
	SentNoiseLevel nullable.Type[int64] `json:"sentNoiseLevel,omitempty"`

	// Average energy level of sent audio for audio classified as mono speech, or left channel of stereo speech by the media
	// endpoint.
	SentSignalLevel nullable.Type[int64] `json:"sentSignalLevel,omitempty"`
}
