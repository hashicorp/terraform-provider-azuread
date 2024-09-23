package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRuleDestination = NetworkaccessIPAddress{}

type NetworkaccessIPAddress struct {
	// Defines the IP address used in a destination for a rule.
	Value *string `json:"value,omitempty"`

	// Fields inherited from NetworkaccessRuleDestination

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessIPAddress) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return BaseNetworkaccessRuleDestinationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessIPAddress{}

func (s NetworkaccessIPAddress) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessIPAddress
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessIPAddress: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessIPAddress: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.ipAddress"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessIPAddress: %+v", err)
	}

	return encoded, nil
}
