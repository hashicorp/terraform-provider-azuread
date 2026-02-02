package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySet = ChatMessageMentionedIdentitySet{}

type ChatMessageMentionedIdentitySet struct {
	// If present, represents a conversation (for example, team, channel, or chat) @mentioned in a message.
	Conversation *TeamworkConversationIdentity `json:"conversation,omitempty"`

	// Fields inherited from IdentitySet

	// Optional. The application associated with this action.
	Application Identity `json:"application"`

	// Optional. The device associated with this action.
	Device Identity `json:"device"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Optional. The user associated with this action.
	User Identity `json:"user"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ChatMessageMentionedIdentitySet) IdentitySet() BaseIdentitySetImpl {
	return BaseIdentitySetImpl{
		Application: s.Application,
		Device:      s.Device,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		User:        s.User,
	}
}

var _ json.Marshaler = ChatMessageMentionedIdentitySet{}

func (s ChatMessageMentionedIdentitySet) MarshalJSON() ([]byte, error) {
	type wrapper ChatMessageMentionedIdentitySet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMessageMentionedIdentitySet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMessageMentionedIdentitySet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatMessageMentionedIdentitySet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMessageMentionedIdentitySet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChatMessageMentionedIdentitySet{}

func (s *ChatMessageMentionedIdentitySet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Conversation *TeamworkConversationIdentity `json:"conversation,omitempty"`
		ODataId      *string                       `json:"@odata.id,omitempty"`
		ODataType    *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Conversation = decoded.Conversation
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChatMessageMentionedIdentitySet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'ChatMessageMentionedIdentitySet': %+v", err)
		}
		s.Application = impl
	}

	if v, ok := temp["device"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Device' for 'ChatMessageMentionedIdentitySet': %+v", err)
		}
		s.Device = impl
	}

	if v, ok := temp["user"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'User' for 'ChatMessageMentionedIdentitySet': %+v", err)
		}
		s.User = impl
	}

	return nil
}
