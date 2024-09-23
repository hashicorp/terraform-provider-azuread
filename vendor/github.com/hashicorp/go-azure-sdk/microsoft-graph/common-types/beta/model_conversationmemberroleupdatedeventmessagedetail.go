package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventMessageDetail = ConversationMemberRoleUpdatedEventMessageDetail{}

type ConversationMemberRoleUpdatedEventMessageDetail struct {
	// Roles for the coversation member user.
	ConversationMemberRoles *[]string `json:"conversationMemberRoles,omitempty"`

	// Identity of the conversation member user.
	ConversationMemberUser *TeamworkUserIdentity `json:"conversationMemberUser,omitempty"`

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

func (s ConversationMemberRoleUpdatedEventMessageDetail) EventMessageDetail() BaseEventMessageDetailImpl {
	return BaseEventMessageDetailImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConversationMemberRoleUpdatedEventMessageDetail{}

func (s ConversationMemberRoleUpdatedEventMessageDetail) MarshalJSON() ([]byte, error) {
	type wrapper ConversationMemberRoleUpdatedEventMessageDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConversationMemberRoleUpdatedEventMessageDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConversationMemberRoleUpdatedEventMessageDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conversationMemberRoleUpdatedEventMessageDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConversationMemberRoleUpdatedEventMessageDetail: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ConversationMemberRoleUpdatedEventMessageDetail{}

func (s *ConversationMemberRoleUpdatedEventMessageDetail) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ConversationMemberRoles *[]string             `json:"conversationMemberRoles,omitempty"`
		ConversationMemberUser  *TeamworkUserIdentity `json:"conversationMemberUser,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ConversationMemberRoles = decoded.ConversationMemberRoles
	s.ConversationMemberUser = decoded.ConversationMemberUser
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ConversationMemberRoleUpdatedEventMessageDetail into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["initiator"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Initiator' for 'ConversationMemberRoleUpdatedEventMessageDetail': %+v", err)
		}
		s.Initiator = impl
	}

	return nil
}
