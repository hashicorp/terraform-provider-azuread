package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRuleDestination interface {
	NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl
}

var _ NetworkaccessRuleDestination = BaseNetworkaccessRuleDestinationImpl{}

type BaseNetworkaccessRuleDestinationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseNetworkaccessRuleDestinationImpl) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return s
}

var _ NetworkaccessRuleDestination = RawNetworkaccessRuleDestinationImpl{}

// RawNetworkaccessRuleDestinationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawNetworkaccessRuleDestinationImpl struct {
	networkaccessRuleDestination BaseNetworkaccessRuleDestinationImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawNetworkaccessRuleDestinationImpl) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return s.networkaccessRuleDestination
}

func (s RawNetworkaccessRuleDestinationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalNetworkaccessRuleDestinationImplementation(input []byte) (NetworkaccessRuleDestination, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessRuleDestination into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.fqdn") {
		var out NetworkaccessFqdn
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessFqdn: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.ipAddress") {
		var out NetworkaccessIPAddress
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessIPAddress: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.ipRange") {
		var out NetworkaccessIPRange
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessIPRange: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.ipSubnet") {
		var out NetworkaccessIPSubnet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessIPSubnet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.url") {
		var out NetworkaccessUrl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessUrl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.networkaccess.webCategory") {
		var out NetworkaccessWebCategory
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NetworkaccessWebCategory: %+v", err)
		}
		return out, nil
	}

	var parent BaseNetworkaccessRuleDestinationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNetworkaccessRuleDestinationImpl: %+v", err)
	}

	return RawNetworkaccessRuleDestinationImpl{
		networkaccessRuleDestination: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
