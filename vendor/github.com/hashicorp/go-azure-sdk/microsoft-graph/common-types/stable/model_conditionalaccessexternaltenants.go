package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessExternalTenants interface {
	ConditionalAccessExternalTenants() BaseConditionalAccessExternalTenantsImpl
}

var _ ConditionalAccessExternalTenants = BaseConditionalAccessExternalTenantsImpl{}

type BaseConditionalAccessExternalTenantsImpl struct {
	Members *[]string `json:"members,omitempty"`

	// The membership kind. Possible values are: all, enumerated, unknownFutureValue. The enumerated member references an
	// conditionalAccessEnumeratedExternalTenants object.
	MembershipKind *ConditionalAccessExternalTenantsMembershipKind `json:"membershipKind,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseConditionalAccessExternalTenantsImpl) ConditionalAccessExternalTenants() BaseConditionalAccessExternalTenantsImpl {
	return s
}

var _ ConditionalAccessExternalTenants = RawConditionalAccessExternalTenantsImpl{}

// RawConditionalAccessExternalTenantsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConditionalAccessExternalTenantsImpl struct {
	conditionalAccessExternalTenants BaseConditionalAccessExternalTenantsImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawConditionalAccessExternalTenantsImpl) ConditionalAccessExternalTenants() BaseConditionalAccessExternalTenantsImpl {
	return s.conditionalAccessExternalTenants
}

func UnmarshalConditionalAccessExternalTenantsImplementation(input []byte) (ConditionalAccessExternalTenants, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessExternalTenants into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessAllExternalTenants") {
		var out ConditionalAccessAllExternalTenants
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessAllExternalTenants: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.conditionalAccessEnumeratedExternalTenants") {
		var out ConditionalAccessEnumeratedExternalTenants
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConditionalAccessEnumeratedExternalTenants: %+v", err)
		}
		return out, nil
	}

	var parent BaseConditionalAccessExternalTenantsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseConditionalAccessExternalTenantsImpl: %+v", err)
	}

	return RawConditionalAccessExternalTenantsImpl{
		conditionalAccessExternalTenants: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
