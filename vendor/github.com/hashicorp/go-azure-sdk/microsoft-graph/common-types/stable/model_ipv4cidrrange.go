package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IPRange = IPv4CIDRRange{}

type IPv4CIDRRange struct {
	// IPv4 address in CIDR notation. Not nullable.
	CIDRAddress *string `json:"cidrAddress,omitempty"`

	// Fields inherited from IPRange

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IPv4CIDRRange) IPRange() BaseIPRangeImpl {
	return BaseIPRangeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IPv4CIDRRange{}

func (s IPv4CIDRRange) MarshalJSON() ([]byte, error) {
	type wrapper IPv4CIDRRange
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPv4CIDRRange: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPv4CIDRRange: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iPv4CidrRange"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPv4CIDRRange: %+v", err)
	}

	return encoded, nil
}
