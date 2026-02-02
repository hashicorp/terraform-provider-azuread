package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPolicyRule interface {
	Entity
	NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl
}

var _ NetworkaccessPolicyRule = BaseNetworkaccessPolicyRuleImpl{}

type BaseNetworkaccessPolicyRuleImpl struct {
	// Name.
	Name *string `json:"name,omitempty"`

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

func (s BaseNetworkaccessPolicyRuleImpl) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return s
}

func (s BaseNetworkaccessPolicyRuleImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NetworkaccessPolicyRule = RawNetworkaccessPolicyRuleImpl{}

// RawNetworkaccessPolicyRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessPolicyRuleImpl struct {
	networkaccessPolicyRule BaseNetworkaccessPolicyRuleImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawNetworkaccessPolicyRuleImpl) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return s.networkaccessPolicyRule
}

func (s RawNetworkaccessPolicyRuleImpl) Entity() BaseEntityImpl {
	return s.networkaccessPolicyRule.Entity()
}

var _ json.Marshaler = BaseNetworkaccessPolicyRuleImpl{}

func (s BaseNetworkaccessPolicyRuleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNetworkaccessPolicyRuleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNetworkaccessPolicyRuleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNetworkaccessPolicyRuleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.policyRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNetworkaccessPolicyRuleImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalNetworkaccessPolicyRuleImplementation(input []byte) (NetworkaccessPolicyRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessPolicyRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.filteringRule") {
		var out NetworkaccessFilteringRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFilteringRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.forwardingRule") {
		var out NetworkaccessForwardingRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessForwardingRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessPolicyRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessPolicyRuleImpl: %+v", err)
	}

	return RawNetworkaccessPolicyRuleImpl{
		networkaccessPolicyRule: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
