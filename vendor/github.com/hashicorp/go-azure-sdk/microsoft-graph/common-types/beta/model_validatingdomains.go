package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidatingDomains interface {
	ValidatingDomains() BaseValidatingDomainsImpl
}

var _ ValidatingDomains = BaseValidatingDomainsImpl{}

type BaseValidatingDomainsImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	RootDomains *RootDomains `json:"rootDomains,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseValidatingDomainsImpl) ValidatingDomains() BaseValidatingDomainsImpl {
	return s
}

var _ ValidatingDomains = RawValidatingDomainsImpl{}

// RawValidatingDomainsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawValidatingDomainsImpl struct {
	validatingDomains BaseValidatingDomainsImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawValidatingDomainsImpl) ValidatingDomains() BaseValidatingDomainsImpl {
	return s.validatingDomains
}

func UnmarshalValidatingDomainsImplementation(input []byte) (ValidatingDomains, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidatingDomains into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.allDomains") {
		var out AllDomains
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllDomains: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.enumeratedDomains") {
		var out EnumeratedDomains
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EnumeratedDomains: %+v", err)
		}
		return out, nil
	}

	var parent BaseValidatingDomainsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseValidatingDomainsImpl: %+v", err)
	}

	return RawValidatingDomainsImpl{
		validatingDomains: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
