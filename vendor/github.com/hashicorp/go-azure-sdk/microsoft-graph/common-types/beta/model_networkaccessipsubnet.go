package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRuleDestination = NetworkaccessIPSubnet{}

type NetworkaccessIPSubnet struct {
	// Defines the IP address of the subset used in a destination for a rule.
	Value *string `json:"value,omitempty"`

	// Fields inherited from NetworkaccessRuleDestination

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessIPSubnet) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return BaseNetworkaccessRuleDestinationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessIPSubnet{}

func (s NetworkaccessIPSubnet) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessIPSubnet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessIPSubnet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessIPSubnet: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.ipSubnet"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessIPSubnet: %+v", err)
	}

	return encoded, nil
}
