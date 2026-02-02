package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ValidatingDomains = EnumeratedDomains{}

type EnumeratedDomains struct {
	// List of federated or managed root domains that Microsoft Entra ID validates.
	DomainNames *[]string `json:"domainNames,omitempty"`

	// Fields inherited from ValidatingDomains

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RootDomains *RootDomains `json:"rootDomains,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EnumeratedDomains) ValidatingDomains() BaseValidatingDomainsImpl {
	return BaseValidatingDomainsImpl{
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		RootDomains: s.RootDomains,
	}
}

var _ json.Marshaler = EnumeratedDomains{}

func (s EnumeratedDomains) MarshalJSON() ([]byte, error) {
	type wrapper EnumeratedDomains
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EnumeratedDomains: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EnumeratedDomains: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.enumeratedDomains"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EnumeratedDomains: %+v", err)
	}

	return encoded, nil
}
