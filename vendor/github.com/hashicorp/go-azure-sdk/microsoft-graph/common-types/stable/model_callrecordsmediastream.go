package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStream struct {
	// Codec name used to encode audio for transmission on the network. Possible values are: unknown, invalid, cn, pcma,
	// pcmu, amrWide, g722, g7221, g7221c, g729, multiChannelAudio, muchv2, opus, satin, satinFullband, rtAudio8, rtAudio16,
	// silk, silkNarrow, silkWide, siren, xmsRta, unknownFutureValue.
	AudioCodec *CallRecordsAudioCodec `json:"audioCodec,omitempty"`

	// Average jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is
	// denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second
	// designator.
	AverageAudioNetworkJitter nullable.Type[string] `json:"averageAudioNetworkJitter,omitempty"`

	// Average estimated bandwidth available between two endpoints in bits per second.
	AverageBandwidthEstimate nullable.Type[int64] `json:"averageBandwidthEstimate,omitempty"`

	// Average duration of the received freezing time in the video stream.
	AverageFreezeDuration nullable.Type[string] `json:"averageFreezeDuration,omitempty"`

	// Average jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is
	// denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second
	// designator.
	AverageJitter nullable.Type[string] `json:"averageJitter,omitempty"`

	// Average network propagation round-trip time computed as specified in RFC 3550, denoted in ISO 8601 format. For
	// example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is
	// the second designator.
	AverageRoundTripTime nullable.Type[string] `json:"averageRoundTripTime,omitempty"`

	// UTC time when the stream ended. The DateTimeOffset type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. This field is only
	// available for streams that use the SIP protocol.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// Indicates whether the forward error correction (FEC) was used at some point during the session. The default value is
	// null.
	IsAudioForwardErrorCorrectionUsed nullable.Type[bool] `json:"isAudioForwardErrorCorrectionUsed,omitempty"`

	// Maximum of audio network jitter computed over each of the 20 second windows during the session, denoted in ISO 8601
	// format. For example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator,
	// and 'S' is the second designator.
	MaxAudioNetworkJitter nullable.Type[string] `json:"maxAudioNetworkJitter,omitempty"`

	// Maximum jitter for the stream computed as specified in RFC 3550, denoted in ISO 8601 format. For example, 1 second is
	// denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is the second
	// designator.
	MaxJitter nullable.Type[string] `json:"maxJitter,omitempty"`

	// Maximum network propagation round-trip time computed as specified in RFC 3550, denoted in ISO 8601 format. For
	// example, 1 second is denoted as 'PT1S', where 'P' is the duration designator, 'T' is the time designator, and 'S' is
	// the second designator.
	MaxRoundTripTime nullable.Type[string] `json:"maxRoundTripTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Packet count for the stream.
	PacketUtilization nullable.Type[int64] `json:"packetUtilization,omitempty"`

	// Average duration of the received freezing time in the video stream represented in root mean square.
	RmsFreezeDuration nullable.Type[string] `json:"rmsFreezeDuration,omitempty"`

	// UTC time when the stream started. The DateTimeOffset type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. This field is only
	// available for streams that use the SIP protocol.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	StreamDirection *CallRecordsMediaStreamDirection `json:"streamDirection,omitempty"`

	// Unique identifier for the stream.
	StreamId nullable.Type[string] `json:"streamId,omitempty"`

	// Codec name used to encode video for transmission on the network. Possible values are: unknown, invalid, av1, h263,
	// h264, h264s, h264uc, h265, rtvc1, rtVideo, xrtvc1, unknownFutureValue.
	VideoCodec *CallRecordsVideoCodec `json:"videoCodec,omitempty"`

	// True if the media stream bypassed the Mediation Server and went straight between client and PSTN Gateway/PBX, false
	// otherwise.
	WasMediaBypassed nullable.Type[bool] `json:"wasMediaBypassed,omitempty"`
}
