package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = ChannelSharingUpdatedEventMessageDetail{}

type ChannelSharingUpdatedEventMessageDetail struct {
	// Initiator of the event.
	Initiator IdentitySet `json:"initiator"`

	// The ID of the team to which the shared channel belongs.
	OwnerTeamId nullable.Type[string] `json:"ownerTeamId,omitempty"`

	// The ID of the tenant to which the shared channel belongs.
	OwnerTenantId nullable.Type[string] `json:"ownerTenantId,omitempty"`

	// The ID of the shared channel.
	SharedChannelId nullable.Type[string] `json:"sharedChannelId,omitempty"`

	// Fields inherited from EventMessageDetail

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChannelSharingUpdatedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChannelSharingUpdatedEventMessageDetail{}

func (s ChannelSharingUpdatedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ChannelSharingUpdatedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChannelSharingUpdatedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChannelSharingUpdatedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.channelSharingUpdatedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChannelSharingUpdatedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChannelSharingUpdatedEventMessageDetail{}

func (s *ChannelSharingUpdatedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		OwnerTeamId     nullable.Type[string] `json:"ownerTeamId,omitempty"`
		OwnerTenantId   nullable.Type[string] `json:"ownerTenantId,omitempty"`
		SharedChannelId nullable.Type[string] `json:"sharedChannelId,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.OwnerTeamId = decoded.OwnerTeamId
	s.OwnerTenantId = decoded.OwnerTenantId
	s.SharedChannelId = decoded.SharedChannelId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChannelSharingUpdatedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'ChannelSharingUpdatedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
