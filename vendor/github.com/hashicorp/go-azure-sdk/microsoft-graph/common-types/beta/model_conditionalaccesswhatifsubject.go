package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessWhatIfSubject interface {
	ConditionalAccessWhatIfSubject() BaseConditionalAccessWhatIfSubjectImpl
}

var _ ConditionalAccessWhatIfSubject = BaseConditionalAccessWhatIfSubjectImpl{}

type BaseConditionalAccessWhatIfSubjectImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseConditionalAccessWhatIfSubjectImpl) ConditionalAccessWhatIfSubject() BaseConditionalAccessWhatIfSubjectImpl {
	return s
}

var _ ConditionalAccessWhatIfSubject = RawConditionalAccessWhatIfSubjectImpl{}

// RawConditionalAccessWhatIfSubjectImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawConditionalAccessWhatIfSubjectImpl struct {
	conditionalAccessWhatIfSubject BaseConditionalAccessWhatIfSubjectImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawConditionalAccessWhatIfSubjectImpl) ConditionalAccessWhatIfSubject() BaseConditionalAccessWhatIfSubjectImpl {
	return s.conditionalAccessWhatIfSubject
}

func UnmarshalConditionalAccessWhatIfSubjectImplementation(input []byte) (ConditionalAccessWhatIfSubject, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ConditionalAccessWhatIfSubject into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.servicePrincipalSubject") {
		var out ServicePrincipalSubject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServicePrincipalSubject: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.userSubject") {
		var out UserSubject
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UserSubject: %+v", err)
		}
		return out, nil
	}

	var parent BaseConditionalAccessWhatIfSubjectImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseConditionalAccessWhatIfSubjectImpl: %+v", err)
	}

	return RawConditionalAccessWhatIfSubjectImpl{
		conditionalAccessWhatIfSubject: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
