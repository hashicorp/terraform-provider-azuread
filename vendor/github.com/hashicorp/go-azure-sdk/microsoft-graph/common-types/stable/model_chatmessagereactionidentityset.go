package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IdentitySet = ChatMessageReactionIdentitySet{}

type ChatMessageReactionIdentitySet struct {

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

func (s ChatMessageReactionIdentitySet) IdentitySet() BaseIdentitySetImpl {
	return BaseIdentitySetImpl{
		Application: s.Application,
		Device:      s.Device,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		User:        s.User,
	}
}

var _ json.Marshaler = ChatMessageReactionIdentitySet{}

func (s ChatMessageReactionIdentitySet) MarshalJSON() ([]byte, error) {
	type wrapper ChatMessageReactionIdentitySet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChatMessageReactionIdentitySet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChatMessageReactionIdentitySet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.chatMessageReactionIdentitySet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChatMessageReactionIdentitySet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChatMessageReactionIdentitySet{}

func (s *ChatMessageReactionIdentitySet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChatMessageReactionIdentitySet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'ChatMessageReactionIdentitySet': %+v", err)
		}
		s.Application = impl
	}

	if v, ok := temp["device"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Device' for 'ChatMessageReactionIdentitySet': %+v", err)
		}
		s.Device = impl
	}

	if v, ok := temp["user"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'User' for 'ChatMessageReactionIdentitySet': %+v", err)
		}
		s.User = impl
	}

	return nil
}
