package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationListener interface {
	Entity
	AuthenticationListener() BaseAuthenticationListenerImpl
}

var _ AuthenticationListener = BaseAuthenticationListenerImpl{}

type BaseAuthenticationListenerImpl struct {
	// The priority of the listener. Determines the order of evaluation when an event has multiple listeners. The priority
	// is evaluated from low to high.
	Priority *int64 `json:"priority,omitempty"`

	// Filter based on the source of the authentication that is used to determine whether the listener is evaluated, and is
	// currently limited to evaluations based on application the user is authenticating to.
	SourceFilter *AuthenticationSourceFilter `json:"sourceFilter,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAuthenticationListenerImpl) AuthenticationListener() BaseAuthenticationListenerImpl {
	return s
}

func (s BaseAuthenticationListenerImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationListener = RawAuthenticationListenerImpl{}

// RawAuthenticationListenerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationListenerImpl struct {
	authenticationListener BaseAuthenticationListenerImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawAuthenticationListenerImpl) AuthenticationListener() BaseAuthenticationListenerImpl {
	return s.authenticationListener
}

func (s RawAuthenticationListenerImpl) Entity() BaseEntityImpl {
	return s.authenticationListener.Entity()
}

var _ json.Marshaler = BaseAuthenticationListenerImpl{}

func (s BaseAuthenticationListenerImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationListenerImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationListenerImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationListenerImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationListener"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationListenerImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationListenerImplementation(input []byte) (AuthenticationListener, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationListener into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.invokeUserFlowListener") {
		var out InvokeUserFlowListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InvokeUserFlowListener: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationListenerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationListenerImpl: %+v", err)
	}

	return RawAuthenticationListenerImpl{
		authenticationListener: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
