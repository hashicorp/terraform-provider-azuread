package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyLocation interface {
	PolicyLocation() BasePolicyLocationImpl
}

var _ PolicyLocation = BasePolicyLocationImpl{}

type BasePolicyLocationImpl struct {
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

func (s BasePolicyLocationImpl) PolicyLocation() BasePolicyLocationImpl {
	return s
}

var _ PolicyLocation = RawPolicyLocationImpl{}

// RawPolicyLocationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPolicyLocationImpl struct {
	policyLocation BasePolicyLocationImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawPolicyLocationImpl) PolicyLocation() BasePolicyLocationImpl {
	return s.policyLocation
}

func UnmarshalPolicyLocationImplementation(input []byte) (PolicyLocation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyLocation into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.policyLocationApplication") {
		var out PolicyLocationApplication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyLocationApplication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyLocationDomain") {
		var out PolicyLocationDomain
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyLocationDomain: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.policyLocationUrl") {
		var out PolicyLocationUrl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PolicyLocationUrl: %+v", err)
		}
		return out, nil
	}

	var parent BasePolicyLocationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePolicyLocationImpl: %+v", err)
	}

	return RawPolicyLocationImpl{
		policyLocation: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
