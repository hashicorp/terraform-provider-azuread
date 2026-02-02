package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = ChannelRenamedEventMessageDetail{}

type ChannelRenamedEventMessageDetail struct {
	// The updated name of the channel.
	ChannelDisplayName nullable.Type[string] `json:"channelDisplayName,omitempty"`

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

func (s ChannelRenamedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChannelRenamedEventMessageDetail{}

func (s ChannelRenamedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ChannelRenamedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChannelRenamedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChannelRenamedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.channelRenamedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChannelRenamedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChannelRenamedEventMessageDetail{}

func (s *ChannelRenamedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ChannelDisplayName nullable.Type[string] `json:"channelDisplayName,omitempty"`
		ChannelId          nullable.Type[string] `json:"channelId,omitempty"`
		ODataId            *string               `json:"@odata.id,omitempty"`
		ODataType          *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ChannelDisplayName = decoded.ChannelDisplayName
	s.ChannelId = decoded.ChannelId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChannelRenamedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'ChannelRenamedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
