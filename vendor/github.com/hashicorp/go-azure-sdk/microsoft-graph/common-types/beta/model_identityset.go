package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentitySet interface {
	IdentitySet() BaseIdentitySetImpl
}

var _ IdentitySet = BaseIdentitySetImpl{}

type BaseIdentitySetImpl struct {
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

func (s BaseIdentitySetImpl) IdentitySet() BaseIdentitySetImpl {
	return s
}

var _ IdentitySet = RawIdentitySetImpl{}

// RawIdentitySetImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIdentitySetImpl struct {
	identitySet BaseIdentitySetImpl
	Type        string
	Values      map[string]interface{}
}

func (s RawIdentitySetImpl) IdentitySet() BaseIdentitySetImpl {
	return s.identitySet
}

var _ json.Unmarshaler = &BaseIdentitySetImpl{}

func (s *BaseIdentitySetImpl) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling BaseIdentitySetImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["application"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Application' for 'BaseIdentitySetImpl': %+v", err)
		}
		s.Application = impl
	}

	if v, ok := temp["device"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Device' for 'BaseIdentitySetImpl': %+v", err)
		}
		s.Device = impl
	}

	if v, ok := temp["user"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'User' for 'BaseIdentitySetImpl': %+v", err)
		}
		s.User = impl
	}

	return nil
}

func UnmarshalIdentitySetImplementation(input []byte) (IdentitySet, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IdentitySet into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.aiInteractionMentionedIdentitySet") {
		var out AiInteractionMentionedIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AiInteractionMentionedIdentitySet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.approvalIdentitySet") {
		var out ApprovalIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApprovalIdentitySet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageFromIdentitySet") {
		var out ChatMessageFromIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageFromIdentitySet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageMentionedIdentitySet") {
		var out ChatMessageMentionedIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageMentionedIdentitySet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.chatMessageReactionIdentitySet") {
		var out ChatMessageReactionIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChatMessageReactionIdentitySet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.communicationsIdentitySet") {
		var out CommunicationsIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommunicationsIdentitySet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.sharePointIdentitySet") {
		var out SharePointIdentitySet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharePointIdentitySet: %+v", err)
		}
		return out, nil
	}

	var parent BaseIdentitySetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIdentitySetImpl: %+v", err)
	}

	return RawIdentitySetImpl{
		identitySet: parent,
		Type:        value,
		Values:      temp,
	}, nil

}
