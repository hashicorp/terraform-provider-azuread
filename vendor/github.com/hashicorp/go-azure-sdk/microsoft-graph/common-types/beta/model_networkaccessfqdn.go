package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ NetworkaccessRuleDestination = NetworkaccessFqdn{}

type NetworkaccessFqdn struct {
	// Defines the FQDN used in a destination for a rule.
	Value *string `json:"value,omitempty"`

	// Fields inherited from NetworkaccessRuleDestination

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s NetworkaccessFqdn) NetworkaccessRuleDestination() BaseNetworkaccessRuleDestinationImpl {
	return BaseNetworkaccessRuleDestinationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessFqdn{}

func (s NetworkaccessFqdn) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessFqdn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessFqdn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessFqdn: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.fqdn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessFqdn: %+v", err)
	}

	return encoded, nil
}
