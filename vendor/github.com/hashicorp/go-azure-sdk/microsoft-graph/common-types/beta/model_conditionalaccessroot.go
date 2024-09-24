package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ConditionalAccessRoot{}

type ConditionalAccessRoot struct {
	// Read-only. Nullable. Returns a collection of the specified authentication context class references.
	AuthenticationContextClassReferences *[]AuthenticationContextClassReference `json:"authenticationContextClassReferences,omitempty"`

	// Defines the authentication strength policies, valid authentication method combinations, and authentication method
	// mode details that can be required by a conditional access policy.
	AuthenticationStrength *AuthenticationStrengthRoot `json:"authenticationStrength,omitempty"`

	// DEPRECATED. See the authenticationStrength relationship instead.
	AuthenticationStrengths *AuthenticationStrengthRoot `json:"authenticationStrengths,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified named locations.
	NamedLocations *[]NamedLocation `json:"namedLocations,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified Conditional Access policies.
	Policies *[]ConditionalAccessPolicy `json:"policies,omitempty"`

	// Read-only. Nullable. Returns a collection of the specified Conditional Access templates.
	Templates *[]ConditionalAccessTemplate `json:"templates,omitempty"`

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

func (s ConditionalAccessRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ConditionalAccessRoot{}

func (s ConditionalAccessRoot) MarshalJSON() ([]byte, error) {
	type wrapper ConditionalAccessRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConditionalAccessRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessRoot: %+v", err)
	}

	delete(decoded, "authenticationContextClassReferences")
	delete(decoded, "namedLocations")
	delete(decoded, "policies")
	delete(decoded, "templates")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.conditionalAccessRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConditionalAccessRoot: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ConditionalAccessRoot{}

func (s *ConditionalAccessRoot) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AuthenticationContextClassReferences *[]AuthenticationContextClassReference `json:"authenticationContextClassReferences,omitempty"`
		AuthenticationStrength               *AuthenticationStrengthRoot            `json:"authenticationStrength,omitempty"`
		AuthenticationStrengths              *AuthenticationStrengthRoot            `json:"authenticationStrengths,omitempty"`
		Templates                            *[]ConditionalAccessTemplate           `json:"templates,omitempty"`
		Id                                   *string                                `json:"id,omitempty"`
		ODataId                              *string                                `json:"@odata.id,omitempty"`
		ODataType                            *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AuthenticationContextClassReferences = decoded.AuthenticationContextClassReferences
	s.AuthenticationStrength = decoded.AuthenticationStrength
	s.AuthenticationStrengths = decoded.AuthenticationStrengths
	s.Templates = decoded.Templates
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ConditionalAccessRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["namedLocations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling NamedLocations into list []json.RawMessage: %+v", err)
		}

		output := make([]NamedLocation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalNamedLocationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'NamedLocations' for 'ConditionalAccessRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.NamedLocations = &output
	}

	if v, ok := temp["policies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Policies into list []json.RawMessage: %+v", err)
		}

		output := make([]ConditionalAccessPolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalConditionalAccessPolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Policies' for 'ConditionalAccessRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Policies = &output
	}

	return nil
}
