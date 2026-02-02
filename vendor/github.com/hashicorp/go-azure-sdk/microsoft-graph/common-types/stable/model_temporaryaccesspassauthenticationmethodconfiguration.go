package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AuthenticationMethodConfiguration = TemporaryAccessPassAuthenticationMethodConfiguration{}

type TemporaryAccessPassAuthenticationMethodConfiguration struct {
	// Default length in characters of a Temporary Access Pass object. Must be between 8 and 48 characters.
	DefaultLength nullable.Type[int64] `json:"defaultLength,omitempty"`

	// Default lifetime in minutes for a Temporary Access Pass. Value can be any integer between the
	// minimumLifetimeInMinutes and maximumLifetimeInMinutes.
	DefaultLifetimeInMinutes nullable.Type[int64] `json:"defaultLifetimeInMinutes,omitempty"`

	// A collection of groups that are enabled to use the authentication method.
	IncludeTargets *[]AuthenticationMethodTarget `json:"includeTargets,omitempty"`

	// If true, all the passes in the tenant will be restricted to one-time use. If false, passes in the tenant can be
	// created to be either one-time use or reusable.
	IsUsableOnce nullable.Type[bool] `json:"isUsableOnce,omitempty"`

	// Maximum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200
	// minutes (equivalent to 30 days).
	MaximumLifetimeInMinutes nullable.Type[int64] `json:"maximumLifetimeInMinutes,omitempty"`

	// Minimum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200
	// minutes (equivalent to 30 days).
	MinimumLifetimeInMinutes nullable.Type[int64] `json:"minimumLifetimeInMinutes,omitempty"`

	// Fields inherited from AuthenticationMethodConfiguration

	// Groups of users that are excluded from a policy.
	ExcludeTargets *[]ExcludeTarget `json:"excludeTargets,omitempty"`

	// The state of the policy. Possible values are: enabled, disabled.
	State *AuthenticationMethodState `json:"state,omitempty"`

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

func (s TemporaryAccessPassAuthenticationMethodConfiguration) AuthenticationMethodConfiguration() BaseAuthenticationMethodConfigurationImpl {
	return BaseAuthenticationMethodConfigurationImpl{
		ExcludeTargets: s.ExcludeTargets,
		State:          s.State,
		Id:             s.Id,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
	}
}

func (s TemporaryAccessPassAuthenticationMethodConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TemporaryAccessPassAuthenticationMethodConfiguration{}

func (s TemporaryAccessPassAuthenticationMethodConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper TemporaryAccessPassAuthenticationMethodConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TemporaryAccessPassAuthenticationMethodConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TemporaryAccessPassAuthenticationMethodConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.temporaryAccessPassAuthenticationMethodConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TemporaryAccessPassAuthenticationMethodConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TemporaryAccessPassAuthenticationMethodConfiguration{}

func (s *TemporaryAccessPassAuthenticationMethodConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DefaultLength            nullable.Type[int64]       `json:"defaultLength,omitempty"`
		DefaultLifetimeInMinutes nullable.Type[int64]       `json:"defaultLifetimeInMinutes,omitempty"`
		IsUsableOnce             nullable.Type[bool]        `json:"isUsableOnce,omitempty"`
		MaximumLifetimeInMinutes nullable.Type[int64]       `json:"maximumLifetimeInMinutes,omitempty"`
		MinimumLifetimeInMinutes nullable.Type[int64]       `json:"minimumLifetimeInMinutes,omitempty"`
		ExcludeTargets           *[]ExcludeTarget           `json:"excludeTargets,omitempty"`
		State                    *AuthenticationMethodState `json:"state,omitempty"`
		Id                       *string                    `json:"id,omitempty"`
		ODataId                  *string                    `json:"@odata.id,omitempty"`
		ODataType                *string                    `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DefaultLength = decoded.DefaultLength
	s.DefaultLifetimeInMinutes = decoded.DefaultLifetimeInMinutes
	s.IsUsableOnce = decoded.IsUsableOnce
	s.MaximumLifetimeInMinutes = decoded.MaximumLifetimeInMinutes
	s.MinimumLifetimeInMinutes = decoded.MinimumLifetimeInMinutes
	s.ExcludeTargets = decoded.ExcludeTargets
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TemporaryAccessPassAuthenticationMethodConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["includeTargets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling IncludeTargets into list []json.RawMessage: %+v", err)
		}

		output := make([]AuthenticationMethodTarget, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAuthenticationMethodTargetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'IncludeTargets' for 'TemporaryAccessPassAuthenticationMethodConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.IncludeTargets = &output
	}

	return nil
}
