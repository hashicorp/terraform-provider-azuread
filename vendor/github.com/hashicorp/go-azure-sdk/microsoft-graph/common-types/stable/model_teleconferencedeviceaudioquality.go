package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeleconferenceDeviceMediaQuality = TeleconferenceDeviceAudioQuality{}

type TeleconferenceDeviceAudioQuality struct {

	// Fields inherited from TeleconferenceDeviceMediaQuality

	// The average inbound stream network jitter.
	AverageInboundJitter nullable.Type[string] `json:"averageInboundJitter,omitempty"`

	// The average inbound stream network round trip delay.
	AverageInboundRoundTripDelay nullable.Type[string] `json:"averageInboundRoundTripDelay,omitempty"`

	// The average outbound stream network jitter.
	AverageOutboundJitter nullable.Type[string] `json:"averageOutboundJitter,omitempty"`

	// The average outbound stream network round trip delay.
	AverageOutboundRoundTripDelay nullable.Type[string] `json:"averageOutboundRoundTripDelay,omitempty"`

	// The channel index of media. Indexing begins with 1. If a media session contains 3 video modalities, channel indexes
	// will be 1, 2, and 3.
	ChannelIndex *int64 `json:"channelIndex,omitempty"`

	// The total number of the inbound packets.
	InboundPackets nullable.Type[int64] `json:"inboundPackets,omitempty"`

	// the local IP address for the media session.
	LocalIPAddress nullable.Type[string] `json:"localIPAddress,omitempty"`

	// The local media port.
	LocalPort nullable.Type[int64] `json:"localPort,omitempty"`

	// The maximum inbound stream network jitter.
	MaximumInboundJitter nullable.Type[string] `json:"maximumInboundJitter,omitempty"`

	// The maximum inbound stream network round trip delay.
	MaximumInboundRoundTripDelay nullable.Type[string] `json:"maximumInboundRoundTripDelay,omitempty"`

	// The maximum outbound stream network jitter.
	MaximumOutboundJitter nullable.Type[string] `json:"maximumOutboundJitter,omitempty"`

	// The maximum outbound stream network round trip delay.
	MaximumOutboundRoundTripDelay nullable.Type[string] `json:"maximumOutboundRoundTripDelay,omitempty"`

	// The total modality duration. If the media enabled and disabled multiple times, MediaDuration will the summation of
	// all of the durations.
	MediaDuration nullable.Type[string] `json:"mediaDuration,omitempty"`

	// The network link speed in bytes
	NetworkLinkSpeedInBytes nullable.Type[int64] `json:"networkLinkSpeedInBytes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The total number of the outbound packets.
	OutboundPackets nullable.Type[int64] `json:"outboundPackets,omitempty"`

	// The remote IP address for the media session.
	RemoteIPAddress nullable.Type[string] `json:"remoteIPAddress,omitempty"`

	// The remote media port.
	RemotePort nullable.Type[int64] `json:"remotePort,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TeleconferenceDeviceAudioQuality) TeleconferenceDeviceMediaQuality() BaseTeleconferenceDeviceMediaQualityImpl {
	return BaseTeleconferenceDeviceMediaQualityImpl{
		AverageInboundJitter:          s.AverageInboundJitter,
		AverageInboundRoundTripDelay:  s.AverageInboundRoundTripDelay,
		AverageOutboundJitter:         s.AverageOutboundJitter,
		AverageOutboundRoundTripDelay: s.AverageOutboundRoundTripDelay,
		ChannelIndex:                  s.ChannelIndex,
		InboundPackets:                s.InboundPackets,
		LocalIPAddress:                s.LocalIPAddress,
		LocalPort:                     s.LocalPort,
		MaximumInboundJitter:          s.MaximumInboundJitter,
		MaximumInboundRoundTripDelay:  s.MaximumInboundRoundTripDelay,
		MaximumOutboundJitter:         s.MaximumOutboundJitter,
		MaximumOutboundRoundTripDelay: s.MaximumOutboundRoundTripDelay,
		MediaDuration:                 s.MediaDuration,
		NetworkLinkSpeedInBytes:       s.NetworkLinkSpeedInBytes,
		ODataId:                       s.ODataId,
		ODataType:                     s.ODataType,
		OutboundPackets:               s.OutboundPackets,
		RemoteIPAddress:               s.RemoteIPAddress,
		RemotePort:                    s.RemotePort,
	}
}

var _ json.Marshaler = TeleconferenceDeviceAudioQuality{}

func (s TeleconferenceDeviceAudioQuality) MarshalJSON() ([]byte, error) {
	type wrapper TeleconferenceDeviceAudioQuality
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TeleconferenceDeviceAudioQuality: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TeleconferenceDeviceAudioQuality: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.teleconferenceDeviceAudioQuality"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TeleconferenceDeviceAudioQuality: %+v", err)
	}

	return encoded, nil
}
