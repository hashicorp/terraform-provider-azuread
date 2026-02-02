package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicy interface {
	Entity
	NetworkaccessPolicy() BaseNetworkaccessPolicyImpl
}

var _ NetworkaccessPolicy = BaseNetworkaccessPolicyImpl{}

type BaseNetworkaccessPolicyImpl struct {
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

func (s BaseNetworkaccessPolicyImpl) NetworkaccessPolicy() BaseNetworkaccessPolicyImpl {
	return s
}

func (s BaseNetworkaccessPolicyImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NetworkaccessPolicy = RawNetworkaccessPolicyImpl{}

// RawNetworkaccessPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessPolicyImpl struct {
	networkaccessPolicy BaseNetworkaccessPolicyImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawNetworkaccessPolicyImpl) NetworkaccessPolicy() BaseNetworkaccessPolicyImpl {
	return s.networkaccessPolicy
}

func (s RawNetworkaccessPolicyImpl) Entity() BaseEntityImpl {
	return s.networkaccessPolicy.Entity()
}

var _ json.Marshaler = BaseNetworkaccessPolicyImpl{}

func (s BaseNetworkaccessPolicyImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNetworkaccessPolicyImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNetworkaccessPolicyImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNetworkaccessPolicyImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.policy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNetworkaccessPolicyImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseNetworkaccessPolicyImpl{}

func (s *BaseNetworkaccessPolicyImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description nullable.Type[string] `json:"description,omitempty"`
		Name        *string               `json:"name,omitempty"`
		Version     *string               `json:"version,omitempty"`
		Id          *string               `json:"id,omitempty"`
		ODataId     *string               `json:"@odata.id,omitempty"`
		ODataType   *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.Name = decoded.Name
	s.Version = decoded.Version
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseNetworkaccessPolicyImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'PolicyRules' for 'BaseNetworkaccessPolicyImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PolicyRules = &output
	}

	return nil
}

func UnmarshalNetworkaccessPolicyImplementation(input []byte) (NetworkaccessPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringPolicy") {
		var out NetworkaccessFilteringPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingPolicy") {
		var out NetworkaccessForwardingPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessPolicyImpl: %+v", err)
	}

	return RawNetworkaccessPolicyImpl{
		networkaccessPolicy: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
