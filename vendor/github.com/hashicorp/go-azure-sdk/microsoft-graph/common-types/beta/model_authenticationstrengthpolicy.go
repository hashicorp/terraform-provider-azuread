package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AuthenticationStrengthPolicy{}

type AuthenticationStrengthPolicy struct {
	// A collection of authentication method modes that are required be used to satify this authentication strength.
	AllowedCombinations *[]AuthenticationMethodModes `json:"allowedCombinations,omitempty"`

	// Settings that may be used to require specific types or instances of an authentication method to be used when
	// authenticating with a specified combination of authentication methods.
	CombinationConfigurations *[]AuthenticationCombinationConfiguration `json:"combinationConfigurations,omitempty"`

	// The datetime when this policy was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The human-readable description of this policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The human-readable display name of this policy. Supports $filter (eq, ne, not , and in).
	DisplayName *string `json:"displayName,omitempty"`

	// The datetime when this policy was last modified.
	ModifiedDateTime *string `json:"modifiedDateTime,omitempty"`

	PolicyType            *AuthenticationStrengthPolicyType   `json:"policyType,omitempty"`
	RequirementsSatisfied *AuthenticationStrengthRequirements `json:"requirementsSatisfied,omitempty"`

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

func (s AuthenticationStrengthPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AuthenticationStrengthPolicy{}

func (s AuthenticationStrengthPolicy) MarshalJSON() ([]byte, error) {
	type wrapper AuthenticationStrengthPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuthenticationStrengthPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuthenticationStrengthPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.authenticationStrengthPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuthenticationStrengthPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AuthenticationStrengthPolicy{}

func (s *AuthenticationStrengthPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowedCombinations   *[]AuthenticationMethodModes        `json:"allowedCombinations,omitempty"`
		CreatedDateTime       *string                             `json:"createdDateTime,omitempty"`
		Description           nullable.Type[string]               `json:"description,omitempty"`
		DisplayName           *string                             `json:"displayName,omitempty"`
		ModifiedDateTime      *string                             `json:"modifiedDateTime,omitempty"`
		PolicyType            *AuthenticationStrengthPolicyType   `json:"policyType,omitempty"`
		RequirementsSatisfied *AuthenticationStrengthRequirements `json:"requirementsSatisfied,omitempty"`
		Id                    *string                             `json:"id,omitempty"`
		ODataId               *string                             `json:"@odata.id,omitempty"`
		ODataType             *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowedCombinations = decoded.AllowedCombinations
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ModifiedDateTime = decoded.ModifiedDateTime
	s.PolicyType = decoded.PolicyType
	s.RequirementsSatisfied = decoded.RequirementsSatisfied
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AuthenticationStrengthPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["combinationConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CombinationConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationCombinationConfiguration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationCombinationConfigurationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CombinationConfigurations' for 'AuthenticationStrengthPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CombinationConfigurations = &output
	}

	return nil
}
