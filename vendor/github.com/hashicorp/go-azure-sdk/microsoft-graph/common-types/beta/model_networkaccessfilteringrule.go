package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringRule interface {
	Entity
	NetworkaccessPolicyRule
	NetworkaccessFilteringRule() BaseNetworkaccessFilteringRuleImpl
}

var _ NetworkaccessFilteringRule = BaseNetworkaccessFilteringRuleImpl{}

type BaseNetworkaccessFilteringRuleImpl struct {
	// Possible destinations and types of destinations accessed by the user in accordance with the network filtering policy,
	// such as IP addresses and FQDNs/URLs.
	Destinations *[]NetworkaccessRuleDestination `json:"destinations,omitempty"`

	RuleType *NetworkaccessNetworkDestinationType `json:"ruleType,omitempty"`

	// Fields inherited from NetworkaccessPolicyRule

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

func (s BaseNetworkaccessFilteringRuleImpl) NetworkaccessFilteringRule() BaseNetworkaccessFilteringRuleImpl {
	return s
}

func (s BaseNetworkaccessFilteringRuleImpl) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return BaseNetworkaccessPolicyRuleImpl{
		Name:      s.Name,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BaseNetworkaccessFilteringRuleImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NetworkaccessFilteringRule = RawNetworkaccessFilteringRuleImpl{}

// RawNetworkaccessFilteringRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessFilteringRuleImpl struct {
	networkaccessFilteringRule BaseNetworkaccessFilteringRuleImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawNetworkaccessFilteringRuleImpl) NetworkaccessFilteringRule() BaseNetworkaccessFilteringRuleImpl {
	return s.networkaccessFilteringRule
}

func (s RawNetworkaccessFilteringRuleImpl) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return s.networkaccessFilteringRule.NetworkaccessPolicyRule()
}

func (s RawNetworkaccessFilteringRuleImpl) Entity() BaseEntityImpl {
	return s.networkaccessFilteringRule.Entity()
}

var _ json.Marshaler = BaseNetworkaccessFilteringRuleImpl{}

func (s BaseNetworkaccessFilteringRuleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNetworkaccessFilteringRuleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNetworkaccessFilteringRuleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNetworkaccessFilteringRuleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.filteringRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNetworkaccessFilteringRuleImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseNetworkaccessFilteringRuleImpl{}

func (s *BaseNetworkaccessFilteringRuleImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		RuleType  *NetworkaccessNetworkDestinationType `json:"ruleType,omitempty"`
		Name      *string                              `json:"name,omitempty"`
		Id        *string                              `json:"id,omitempty"`
		ODataId   *string                              `json:"@odata.id,omitempty"`
		ODataType *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.RuleType = decoded.RuleType
	s.Id = decoded.Id
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseNetworkaccessFilteringRuleImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["destinations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Destinations into list []json.RawMessage: %+v", err)
		}

		output := make([]NetworkaccessRuleDestination, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalNetworkaccessRuleDestinationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'BaseNetworkaccessFilteringRuleImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}

func UnmarshalNetworkaccessFilteringRuleImplementation(input []byte) (NetworkaccessFilteringRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessFilteringRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.fqdnFilteringRule") {
		var out NetworkaccessFqdnFilteringRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFqdnFilteringRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.webCategoryFilteringRule") {
		var out NetworkaccessWebCategoryFilteringRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessWebCategoryFilteringRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessFilteringRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessFilteringRuleImpl: %+v", err)
	}

	return RawNetworkaccessFilteringRuleImpl{
		networkaccessFilteringRule: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
