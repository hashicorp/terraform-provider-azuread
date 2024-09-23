package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationDataProvider interface {
	EducationSynchronizationDataProvider() BaseEducationSynchronizationDataProviderImpl
}

var _ EducationSynchronizationDataProvider = BaseEducationSynchronizationDataProviderImpl{}

type BaseEducationSynchronizationDataProviderImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationSynchronizationDataProviderImpl) EducationSynchronizationDataProvider() BaseEducationSynchronizationDataProviderImpl {
	return s
}

var _ EducationSynchronizationDataProvider = RawEducationSynchronizationDataProviderImpl{}

// RawEducationSynchronizationDataProviderImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationSynchronizationDataProviderImpl struct {
	educationSynchronizationDataProvider BaseEducationSynchronizationDataProviderImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawEducationSynchronizationDataProviderImpl) EducationSynchronizationDataProvider() BaseEducationSynchronizationDataProviderImpl {
	return s.educationSynchronizationDataProvider
}

func UnmarshalEducationSynchronizationDataProviderImplementation(input []byte) (EducationSynchronizationDataProvider, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationDataProvider into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationCsvDataProvider") {
		var out EducationCsvDataProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationCsvDataProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationOneRosterApiDataProvider") {
		var out EducationOneRosterApiDataProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationOneRosterApiDataProvider: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationPowerSchoolDataProvider") {
		var out EducationPowerSchoolDataProvider
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationPowerSchoolDataProvider: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationSynchronizationDataProviderImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationSynchronizationDataProviderImpl: %+v", err)
	}

	return RawEducationSynchronizationDataProviderImpl{
		educationSynchronizationDataProvider: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
