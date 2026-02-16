package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessFilteringRule = NetworkaccessFqdnFilteringRule{}

type NetworkaccessFqdnFilteringRule struct {

	// Fields inherited from NetworkaccessFilteringRule

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

func (s NetworkaccessFqdnFilteringRule) NetworkaccessFilteringRule() BaseNetworkaccessFilteringRuleImpl {
	return BaseNetworkaccessFilteringRuleImpl{
		Destinations: s.Destinations,
		RuleType:     s.RuleType,
		Name:         s.Name,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s NetworkaccessFqdnFilteringRule) NetworkaccessPolicyRule() BaseNetworkaccessPolicyRuleImpl {
	return BaseNetworkaccessPolicyRuleImpl{
		Name:      s.Name,
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s NetworkaccessFqdnFilteringRule) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessFqdnFilteringRule{}

func (s NetworkaccessFqdnFilteringRule) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessFqdnFilteringRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessFqdnFilteringRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessFqdnFilteringRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.fqdnFilteringRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessFqdnFilteringRule: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessFqdnFilteringRule{}

func (s *NetworkaccessFqdnFilteringRule) UnmarshalJSON(bytes []byte) error {
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

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RuleType = decoded.RuleType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessFqdnFilteringRule into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Destinations' for 'NetworkaccessFqdnFilteringRule': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Destinations = &output
	}

	return nil
}
