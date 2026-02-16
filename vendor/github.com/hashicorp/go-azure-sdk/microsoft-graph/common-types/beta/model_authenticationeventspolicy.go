package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationEventsPolicy{}

type AuthenticationEventsPolicy struct {
	// A list of applicable actions to be taken on sign-up.
	OnSignupStart *[]AuthenticationListener `json:"onSignupStart,omitempty"`

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

func (s AuthenticationEventsPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationEventsPolicy{}

func (s AuthenticationEventsPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationEventsPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationEventsPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationEventsPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationEventsPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationEventsPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AuthenticationEventsPolicy{}

func (s *AuthenticationEventsPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id        *string `json:"id,omitempty"`
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AuthenticationEventsPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["onSignupStart"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OnSignupStart into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationListener, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationListenerImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OnSignupStart' for 'AuthenticationEventsPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OnSignupStart = &output
	}

	return nil
}
