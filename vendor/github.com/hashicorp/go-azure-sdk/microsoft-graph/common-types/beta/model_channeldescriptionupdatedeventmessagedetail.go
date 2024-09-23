package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = ChannelDescriptionUpdatedEventMessageDetail{}

type ChannelDescriptionUpdatedEventMessageDetail struct {
	// The updated description of the channel.
	ChannelDescription nullable.Type[string] `json:"channelDescription,omitempty"`

	// Unique identifier of the channel.
	ChannelId nullable.Type[string] `json:"channelId,omitempty"`

	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChannelDescriptionUpdatedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChannelDescriptionUpdatedEventMessageDetail{}

func (s ChannelDescriptionUpdatedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ChannelDescriptionUpdatedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChannelDescriptionUpdatedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChannelDescriptionUpdatedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.channelDescriptionUpdatedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChannelDescriptionUpdatedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChannelDescriptionUpdatedEventMessageDetail{}

func (s *ChannelDescriptionUpdatedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChannelDescription nullable.Type[string] `json:"channelDescription,omitempty"`
		ChannelId          nullable.Type[string] `json:"channelId,omitempty"`
		ODataId            *string               `json:"@odata.id,omitempty"`
		ODataType          *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChannelDescription = decoded.ChannelDescription
	s.ChannelId = decoded.ChannelId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChannelDescriptionUpdatedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'ChannelDescriptionUpdatedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
