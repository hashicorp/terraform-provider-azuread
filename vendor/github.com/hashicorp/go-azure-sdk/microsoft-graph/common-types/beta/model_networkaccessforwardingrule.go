package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingRule interface {
	Entity
	NetworkaccessPolicyRule
	NetworkaccessForwardingRule() BaseNetworkaccessForwardingRuleImpl
}

var _ NetworkaccessForwardingRule = BaseNetworkaccessForwardingRuleImpl{}

type BaseNetworkaccessForwardingRuleImpl struct {
	Action *NetworkaccessForwardingRuleAction `json:"action,omitempty"`

	// Destinations maintain a list of potential destinations and destination types that the user may access within the
	// context of a network filtering policy. This includes IP addresses and fully qualified domain names (FQDNs)/URLs.
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

func (s BaseNetworkaccessForwardingRuleImpl) NetworkaccessForwardingRule() BaseNetworkaccessForwardingRuleImpl {
	return s
}

func (s BaseNetworkaccessForwardingRuleImpl) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return BaseNetworkaccessPolicyRuleImpl{
		Name:      s.Name,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BaseNetworkaccessForwardingRuleImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ NetworkaccessForwardingRule = RawNetworkaccessForwardingRuleImpl{}

// RawNetworkaccessForwardingRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNetworkaccessForwardingRuleImpl struct {
	networkaccessForwardingRule BaseNetworkaccessForwardingRuleImpl
	Type                        string
	Values                      map[string]interface{}
}

func (s RawNetworkaccessForwardingRuleImpl) NetworkaccessForwardingRule() BaseNetworkaccessForwardingRuleImpl {
	return s.networkaccessForwardingRule
}

func (s RawNetworkaccessForwardingRuleImpl) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return s.networkaccessForwardingRule.NetworkaccessPolicyRule()
}

func (s RawNetworkaccessForwardingRuleImpl) Entity() BaseEntityImpl {
	return s.networkaccessForwardingRule.Entity()
}

var _ json.Marshaler = BaseNetworkaccessForwardingRuleImpl{}

func (s BaseNetworkaccessForwardingRuleImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseNetworkaccessForwardingRuleImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseNetworkaccessForwardingRuleImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseNetworkaccessForwardingRuleImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.forwardingRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseNetworkaccessForwardingRuleImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseNetworkaccessForwardingRuleImpl{}

func (s *BaseNetworkaccessForwardingRuleImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Action    *NetworkaccessForwardingRuleAction   `json:"action,omitempty"`
		RuleType  *NetworkaccessNetworkDestinationType `json:"ruleType,omitempty"`
		Name      *string                              `json:"name,omitempty"`
		Id        *string                              `json:"id,omitempty"`
		ODataId   *string                              `json:"@odata.id,omitempty"`
		ODataType *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Action = decoded.Action
	s.RuleType = decoded.RuleType
	s.Id = decoded.Id
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseNetworkaccessForwardingRuleImpl into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'BaseNetworkaccessForwardingRuleImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}

func UnmarshalNetworkaccessForwardingRuleImplementation(input []byte) (NetworkaccessForwardingRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessForwardingRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.internetAccessForwardingRule") {
		var out NetworkaccessInternetAccessForwardingRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessInternetAccessForwardingRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.m365ForwardingRule") {
		var out NetworkaccessM365ForwardingRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessM365ForwardingRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.privateAccessForwardingRule") {
		var out NetworkaccessPrivateAccessForwardingRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessPrivateAccessForwardingRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessForwardingRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessForwardingRuleImpl: %+v", err)
	}

	return RawNetworkaccessForwardingRuleImpl{
		networkaccessForwardingRule: parent,
		Type:                        value,
		Values:                      temp,
	}, nil

}
