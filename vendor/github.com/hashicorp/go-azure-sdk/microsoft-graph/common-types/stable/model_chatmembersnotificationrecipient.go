package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamworkNotificationRecipient = ChatMembersNotificationRecipient{}

type ChatMembersNotificationRecipient struct {
	// The unique identifier for the chat whose members should receive the notifications.
	ChatId *string `json:"chatId,omitempty"`

	// Fields inherited from TeamworkNotificationRecipient

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChatMembersNotificationRecipient) TeamworkNotificationRecipient() BaseTeamworkNotificationRecipientImpl {
	return BaseTeamworkNotificationRecipientImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ChatMembersNotificationRecipient{}

func (s ChatMembersNotificationRecipient) MarshalJSON() ([]byte, error) {
	type wrapper ChatMembersNotificationRecipient
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMembersNotificationRecipient: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMembersNotificationRecipient: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatMembersNotificationRecipient"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMembersNotificationRecipient: %+v", err)
	}

	return encoded, nil
}
