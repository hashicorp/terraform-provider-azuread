package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessForwardingRule = NetworkaccessInternetAccessForwardingRule{}

type NetworkaccessInternetAccessForwardingRule struct {
	Ports    *[]string                        `json:"ports,omitempty"`
	Protocol *NetworkaccessNetworkingProtocol `json:"protocol,omitempty"`

	// Fields inherited from NetworkaccessForwardingRule

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

func (s NetworkaccessInternetAccessForwardingRule) NetworkaccessForwardingRule() BaseNetworkaccessForwardingRuleImpl {
	return BaseNetworkaccessForwardingRuleImpl{
		Action:       s.Action,
		Destinations: s.Destinations,
		RuleType:     s.RuleType,
		Name:         s.Name,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s NetworkaccessInternetAccessForwardingRule) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return BaseNetworkaccessPolicyRuleImpl{
		Name:      s.Name,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s NetworkaccessInternetAccessForwardingRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessInternetAccessForwardingRule{}

func (s NetworkaccessInternetAccessForwardingRule) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessInternetAccessForwardingRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessInternetAccessForwardingRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessInternetAccessForwardingRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.internetAccessForwardingRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessInternetAccessForwardingRule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessInternetAccessForwardingRule{}

func (s *NetworkaccessInternetAccessForwardingRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Ports     *[]string                            `json:"ports,omitempty"`
		Protocol  *NetworkaccessNetworkingProtocol     `json:"protocol,omitempty"`
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

	s.Ports = decoded.Ports
	s.Protocol = decoded.Protocol
	s.Action = decoded.Action
	s.Id = decoded.Id
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RuleType = decoded.RuleType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessInternetAccessForwardingRule into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NetworkaccessInternetAccessForwardingRule': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
