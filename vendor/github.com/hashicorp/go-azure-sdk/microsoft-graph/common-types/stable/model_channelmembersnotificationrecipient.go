package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamworkNotificationRecipient = ChannelMembersNotificationRecipient{}

type ChannelMembersNotificationRecipient struct {
	// The unique identifier for the channel whose members should receive the notification.
	ChannelId *string `json:"channelId,omitempty"`

	// The unique identifier for the team under which the channel resides.
	TeamId *string `json:"teamId,omitempty"`

	// Fields inherited from TeamworkNotificationRecipient

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChannelMembersNotificationRecipient) TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl {
	return BaseTeamworkNotificationRecipientImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChannelMembersNotificationRecipient{}

func (s ChannelMembersNotificationRecipient) MarshalJSON() ([]byte, error) {
	type wrapper ChannelMembersNotificationRecipient
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChannelMembersNotificationRecipient: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChannelMembersNotificationRecipient: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.channelMembersNotificationRecipient"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChannelMembersNotificationRecipient: %+v", err)
	}

	return encoded, nil
}
