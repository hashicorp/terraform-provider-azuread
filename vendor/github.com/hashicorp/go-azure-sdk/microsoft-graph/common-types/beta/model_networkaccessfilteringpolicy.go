package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessPolicy = NetworkaccessFilteringPolicy{}

type NetworkaccessFilteringPolicy struct {
	Action *NetworkaccessFilteringPolicyAction `json:"action,omitempty"`

	// The date and time when the filtering Policy was originally created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time when a particular profile was last modified or updated.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Fields inherited from NetworkaccessPolicy

	// Description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Policy name.
	Name *string `json:"name,omitempty"`

	// Represents the definition of the policy ruleset that makes up the core definition of a policy.
	PolicyRules *[]NetworkaccessPolicyRule `json:"policyRules,omitempty"`

	// Version.
	Version *string `json:"version,omitempty"`

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

func (s NetworkaccessFilteringPolicy) NetworkaccessPolicy() BaseNetworkaccessPolicyImpl {
	return BaseNetworkaccessPolicyImpl{
		Description: s.Description,
		Name:        s.Name,
		PolicyRules: s.PolicyRules,
		Version:     s.Version,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

func (s NetworkaccessFilteringPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessFilteringPolicy{}

func (s NetworkaccessFilteringPolicy) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessFilteringPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessFilteringPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessFilteringPolicy: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.filteringPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessFilteringPolicy: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessFilteringPolicy{}

func (s *NetworkaccessFilteringPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action               *NetworkaccessFilteringPolicyAction `json:"action,omitempty"`
		CreatedDateTime      *string                             `json:"createdDateTime,omitempty"`
		LastModifiedDateTime *string                             `json:"lastModifiedDateTime,omitempty"`
		Description          nullable.Type[string]               `json:"description,omitempty"`
		Name                 *string                             `json:"name,omitempty"`
		Version              *string                             `json:"version,omitempty"`
		Id                   *string                             `json:"id,omitempty"`
		ODataId              *string                             `json:"@odata.id,omitempty"`
		ODataType            *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.CreatedDateTime = decoded.CreatedDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Description = decoded.Description
	s.Id = decoded.Id
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessFilteringPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["policyRules"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PolicyRules into list []json.RawMessage: %+v", err)
		}

		output := make([]NetworkaccessPolicyRule, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalNetworkaccessPolicyRuleImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PolicyRules' for 'NetworkaccessFilteringPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PolicyRules = &output
	}

	return nil
}
