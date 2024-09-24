package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationCustomizationsBase interface {
	EducationSynchronizationCustomizationsBase() BaseEducationSynchronizationCustomizationsBaseImpl
}

var _ EducationSynchronizationCustomizationsBase = BaseEducationSynchronizationCustomizationsBaseImpl{}

type BaseEducationSynchronizationCustomizationsBaseImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationSynchronizationCustomizationsBaseImpl) EducationSynchronizationCustomizationsBase() BaseEducationSynchronizationCustomizationsBaseImpl {
	return s
}

var _ EducationSynchronizationCustomizationsBase = RawEducationSynchronizationCustomizationsBaseImpl{}

// RawEducationSynchronizationCustomizationsBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationSynchronizationCustomizationsBaseImpl struct {
	educationSynchronizationCustomizationsBase BaseEducationSynchronizationCustomizationsBaseImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawEducationSynchronizationCustomizationsBaseImpl) EducationSynchronizationCustomizationsBase() BaseEducationSynchronizationCustomizationsBaseImpl {
	return s.educationSynchronizationCustomizationsBase
}

func UnmarshalEducationSynchronizationCustomizationsBaseImplementation(input []byte) (EducationSynchronizationCustomizationsBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationCustomizationsBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSynchronizationCustomizations") {
		var out EducationSynchronizationCustomizations
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSynchronizationCustomizations: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationSynchronizationCustomizationsBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationSynchronizationCustomizationsBaseImpl: %+v", err)
	}

	return RawEducationSynchronizationCustomizationsBaseImpl{
		educationSynchronizationCustomizationsBase: parent,
		Type:   value,
		Values: temp,
	}, nil

}
