package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventsFlow interface {
	Entity
	AuthenticationEventsFlow() BaseAuthenticationEventsFlowImpl
}

var _ AuthenticationEventsFlow = BaseAuthenticationEventsFlowImpl{}

type BaseAuthenticationEventsFlowImpl struct {
	// The conditions representing the context of the authentication request that's used to decide whether the events policy
	// is invoked. Supports $filter (eq). See support for filtering on user flows for syntax information.
	Conditions *AuthenticationConditions `json:"conditions,omitempty"`

	// The description of the events policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Required. The display name for the events policy.
	DisplayName string `json:"displayName"`

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

func (s BaseAuthenticationEventsFlowImpl) AuthenticationEventsFlow() BaseAuthenticationEventsFlowImpl {
	return s
}

func (s BaseAuthenticationEventsFlowImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AuthenticationEventsFlow = RawAuthenticationEventsFlowImpl{}

// RawAuthenticationEventsFlowImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAuthenticationEventsFlowImpl struct {
	authenticationEventsFlow BaseAuthenticationEventsFlowImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawAuthenticationEventsFlowImpl) AuthenticationEventsFlow() BaseAuthenticationEventsFlowImpl {
	return s.authenticationEventsFlow
}

func (s RawAuthenticationEventsFlowImpl) Entity() BaseEntityImpl {
	return s.authenticationEventsFlow.Entity()
}

var _ json.Marshaler = BaseAuthenticationEventsFlowImpl{}

func (s BaseAuthenticationEventsFlowImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAuthenticationEventsFlowImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAuthenticationEventsFlowImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAuthenticationEventsFlowImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationEventsFlow"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAuthenticationEventsFlowImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAuthenticationEventsFlowImplementation(input []byte) (AuthenticationEventsFlow, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationEventsFlow into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.externalUsersSelfServiceSignUpEventsFlow") {
		var out ExternalUsersSelfServiceSignUpEventsFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalUsersSelfServiceSignUpEventsFlow: %+v", err)
		}
		return out, nil
	}

	var parent BaseAuthenticationEventsFlowImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAuthenticationEventsFlowImpl: %+v", err)
	}

	return RawAuthenticationEventsFlowImpl{
		authenticationEventsFlow: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
