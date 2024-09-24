package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeleconferenceDeviceMediaQuality interface {
	TeleconferenceDeviceMediaQuality() BaseTeleconferenceDeviceMediaQualityImpl
}

var _ TeleconferenceDeviceMediaQuality = BaseTeleconferenceDeviceMediaQualityImpl{}

type BaseTeleconferenceDeviceMediaQualityImpl struct {
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

func (s BaseTeleconferenceDeviceMediaQualityImpl) TeleconferenceDeviceMediaQuality() BaseTeleconferenceDeviceMediaQualityImpl {
	return s
}

var _ TeleconferenceDeviceMediaQuality = RawTeleconferenceDeviceMediaQualityImpl{}

// RawTeleconferenceDeviceMediaQualityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTeleconferenceDeviceMediaQualityImpl struct {
	teleconferenceDeviceMediaQuality BaseTeleconferenceDeviceMediaQualityImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawTeleconferenceDeviceMediaQualityImpl) TeleconferenceDeviceMediaQuality() BaseTeleconferenceDeviceMediaQualityImpl {
	return s.teleconferenceDeviceMediaQuality
}

func UnmarshalTeleconferenceDeviceMediaQualityImplementation(input []byte) (TeleconferenceDeviceMediaQuality, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TeleconferenceDeviceMediaQuality into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.teleconferenceDeviceAudioQuality") {
		var out TeleconferenceDeviceAudioQuality
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeleconferenceDeviceAudioQuality: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.teleconferenceDeviceVideoQuality") {
		var out TeleconferenceDeviceVideoQuality
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TeleconferenceDeviceVideoQuality: %+v", err)
		}
		return out, nil
	}

	var parent BaseTeleconferenceDeviceMediaQualityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTeleconferenceDeviceMediaQualityImpl: %+v", err)
	}

	return RawTeleconferenceDeviceMediaQualityImpl{
		teleconferenceDeviceMediaQuality: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
