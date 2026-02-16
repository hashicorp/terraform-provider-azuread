package beta

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

	// If present, represents a tag @mentioned in a team message.
	Tag *TeamworkTagIdentity `json:"tag,omitempty"`

	// Fields inherited from IdentitySet

	// The Identity of the Application. This property is read-only.
	Application Identity `json:"application"`

	// The Identity of the Device. This property is read-only.
	Device Identity `json:"device"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Identity of the User. This property is read-only.
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
		Tag          *TeamworkTagIdentity          `json:"tag,omitempty"`
		ODataId      *string                       `json:"@odata.id,omitempty"`
		ODataType    *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Conversation = decoded.Conversation
	s.Tag = decoded.Tag
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
