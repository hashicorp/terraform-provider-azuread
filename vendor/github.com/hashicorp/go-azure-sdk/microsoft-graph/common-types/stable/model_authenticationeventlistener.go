package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventListener interface {
	Entity
	AuthenticationEventListener() BaseAuthenticationEventListenerImpl
}

var _ AuthenticationEventListener = BaseAuthenticationEventListenerImpl{}

type BaseAuthenticationEventListenerImpl struct {
	// Indicates the authenticationEventListener is associated with an authenticationEventsFlow. Read-only.
	AuthenticationEventsFlowId nullable.Type[string] `json:"authenticationEventsFlowId,omitempty"`

	// The conditions on which this authenticationEventListener should trigger.
	Conditions *AuthenticationConditions `json:"conditions,omitempty"`

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

func (s BaseAuthenticationEventListenerImpl) AuthenticationEventListener() BaseAuthenticationEventListenerImpl {
	return s
}

func (s BaseAuthenticationEventListenerImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationEventListener = RawAuthenticationEventListenerImpl{}

// RawAuthenticationEventListenerImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationEventListenerImpl struct {
	authenticationEventListener BaseAuthenticationEventListenerImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawAuthenticationEventListenerImpl) AuthenticationEventListener() BaseAuthenticationEventListenerImpl {
	return s.authenticationEventListener
}

func (s RawAuthenticationEventListenerImpl) Entity() BaseEntityImpl {
	return s.authenticationEventListener.Entity()
}

var _ json.Marshaler = BaseAuthenticationEventListenerImpl{}

func (s BaseAuthenticationEventListenerImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationEventListenerImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationEventListenerImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationEventListenerImpl: %+v", err)
	}

	delete(decoded, "authenticationEventsFlowId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationEventListener"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationEventListenerImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationEventListenerImplementation(input []byte) (AuthenticationEventListener, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationEventListener into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionListener") {
		var out OnAttributeCollectionListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionStartListener") {
		var out OnAttributeCollectionStartListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionStartListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAttributeCollectionSubmitListener") {
		var out OnAttributeCollectionSubmitListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAttributeCollectionSubmitListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onAuthenticationMethodLoadStartListener") {
		var out OnAuthenticationMethodLoadStartListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnAuthenticationMethodLoadStartListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onInteractiveAuthFlowStartListener") {
		var out OnInteractiveAuthFlowStartListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnInteractiveAuthFlowStartListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onTokenIssuanceStartListener") {
		var out OnTokenIssuanceStartListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnTokenIssuanceStartListener: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.onUserCreateStartListener") {
		var out OnUserCreateStartListener
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OnUserCreateStartListener: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationEventListenerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationEventListenerImpl: %+v", err)
	}

	return RawAuthenticationEventListenerImpl{
		authenticationEventListener: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
