package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PolicyLocation = PolicyLocationUrl{}

type PolicyLocationUrl struct {

	// Fields inherited from PolicyLocation

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The actual value representing the location (for example, 'contoso.com', 'https://partner.contoso.com/upload',
	// '83ef198a-0396-4893-9d4f-d36efbffcaaa').
	Value *string `json:"value,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PolicyLocationUrl) PolicyLocation() BasePolicyLocationImpl {
	return BasePolicyLocationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Value:     s.Value,
	}
}

var _ json.Marshaler = PolicyLocationUrl{}

func (s PolicyLocationUrl) MarshalJSON() ([]byte, error) {
	type wrapper PolicyLocationUrl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicyLocationUrl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyLocationUrl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policyLocationUrl"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicyLocationUrl: %+v", err)
	}

	return encoded, nil
}
