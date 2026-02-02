package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IPRange = IPv6Range{}

type IPv6Range struct {
	// Lower address.
	LowerAddress *string `json:"lowerAddress,omitempty"`

	// Upper address.
	UpperAddress *string `json:"upperAddress,omitempty"`

	// Fields inherited from IPRange

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IPv6Range) IPRange() BaseIPRangeImpl {
	return BaseIPRangeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IPv6Range{}

func (s IPv6Range) MarshalJSON() ([]byte, error) {
	type wrapper IPv6Range
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IPv6Range: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IPv6Range: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iPv6Range"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IPv6Range: %+v", err)
	}

	return encoded, nil
}
