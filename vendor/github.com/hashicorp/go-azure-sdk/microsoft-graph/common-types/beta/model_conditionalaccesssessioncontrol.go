package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessSessionControl interface {
	ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl
}

var _ ConditionalAccessSessionControl = BaseConditionalAccessSessionControlImpl{}

type BaseConditionalAccessSessionControlImpl struct {
	// Specifies whether the session control is enabled.
	IsEnabled nullable.Type[bool] `json:"isEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseConditionalAccessSessionControlImpl) ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl {
	return s
}

var _ ConditionalAccessSessionControl = RawConditionalAccessSessionControlImpl{}

// RawConditionalAccessSessionControlImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConditionalAccessSessionControlImpl struct {
	conditionalAccessSessionControl BaseConditionalAccessSessionControlImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawConditionalAccessSessionControlImpl) ConditionalAccessSessionControl() BaseConditionalAccessSessionControlImpl {
	return s.conditionalAccessSessionControl
}

func UnmarshalConditionalAccessSessionControlImplementation(input []byte) (ConditionalAccessSessionControl, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessSessionControl into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.applicationEnforcedRestrictionsSessionControl") {
		var out ApplicationEnforcedRestrictionsSessionControl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ApplicationEnforcedRestrictionsSessionControl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.cloudAppSecuritySessionControl") {
		var out CloudAppSecuritySessionControl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CloudAppSecuritySessionControl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.globalSecureAccessFilteringProfileSessionControl") {
		var out GlobalSecureAccessFilteringProfileSessionControl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GlobalSecureAccessFilteringProfileSessionControl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.persistentBrowserSessionControl") {
		var out PersistentBrowserSessionControl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PersistentBrowserSessionControl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.secureSignInSessionControl") {
		var out SecureSignInSessionControl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecureSignInSessionControl: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.signInFrequencySessionControl") {
		var out SignInFrequencySessionControl
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SignInFrequencySessionControl: %+v", err)
		}
		return out, nil
	}

	var parent BaseConditionalAccessSessionControlImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseConditionalAccessSessionControlImpl: %+v", err)
	}

	return RawConditionalAccessSessionControlImpl{
		conditionalAccessSessionControl: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
