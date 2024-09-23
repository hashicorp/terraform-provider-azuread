package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRuleDestination = NetworkaccessIPRange{}

type NetworkaccessIPRange struct {
	// Specifies the starting IP address of the IP range.
	BeginAddress *string `json:"beginAddress,omitempty"`

	// Specifies the ending IP address of the IP range.
	EndAddress *string `json:"endAddress,omitempty"`

	// Fields inherited from NetworkaccessRuleDestination

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessIPRange) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return BaseNetworkaccessRuleDestinationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessIPRange{}

func (s NetworkaccessIPRange) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessIPRange
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessIPRange: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessIPRange: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.ipRange"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessIPRange: %+v", err)
	}

	return encoded, nil
}
