package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessContext interface {
	ConditionalAccessContext() BaseConditionalAccessContextImpl
}

var _ ConditionalAccessContext = BaseConditionalAccessContextImpl{}

type BaseConditionalAccessContextImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseConditionalAccessContextImpl) ConditionalAccessContext() BaseConditionalAccessContextImpl {
	return s
}

var _ ConditionalAccessContext = RawConditionalAccessContextImpl{}

// RawConditionalAccessContextImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConditionalAccessContextImpl struct {
	conditionalAccessContext BaseConditionalAccessContextImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawConditionalAccessContextImpl) ConditionalAccessContext() BaseConditionalAccessContextImpl {
	return s.conditionalAccessContext
}

func UnmarshalConditionalAccessContextImplementation(input []byte) (ConditionalAccessContext, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessContext into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.whatIfApplicationContext") {
		var out WhatIfApplicationContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WhatIfApplicationContext: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.whatIfAuthenticationContext") {
		var out WhatIfAuthenticationContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WhatIfAuthenticationContext: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.whatIfUserActionContext") {
		var out WhatIfUserActionContext
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WhatIfUserActionContext: %+v", err)
		}
		return out, nil
	}

	var parent BaseConditionalAccessContextImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseConditionalAccessContextImpl: %+v", err)
	}

	return RawConditionalAccessContextImpl{
		conditionalAccessContext: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
