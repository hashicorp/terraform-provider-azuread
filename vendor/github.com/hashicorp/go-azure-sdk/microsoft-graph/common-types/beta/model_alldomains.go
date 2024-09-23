package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ValidatingDomains = AllDomains{}

type AllDomains struct {

	// Fields inherited from ValidatingDomains

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RootDomains *RootDomains `json:"rootDomains,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AllDomains) ValidatingDomains() BaseValidatingDomainsImpl {
	return BaseValidatingDomainsImpl{
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		RootDomains: s.RootDomains,
	}
}

var _ json.Marshaler = AllDomains{}

func (s AllDomains) MarshalJSON() ([]byte, error) {
	type wrapper AllDomains
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AllDomains: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AllDomains: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.allDomains"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AllDomains: %+v", err)
	}

	return encoded, nil
}
