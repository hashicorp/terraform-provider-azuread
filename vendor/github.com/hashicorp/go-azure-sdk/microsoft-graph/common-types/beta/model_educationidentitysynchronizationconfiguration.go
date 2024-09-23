package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationIdentitySynchronizationConfiguration interface {
	EducationIdentitySynchronizationConfiguration() BaseEducationIdentitySynchronizationConfigurationImpl
}

var _ EducationIdentitySynchronizationConfiguration = BaseEducationIdentitySynchronizationConfigurationImpl{}

type BaseEducationIdentitySynchronizationConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationIdentitySynchronizationConfigurationImpl) EducationIdentitySynchronizationConfiguration() BaseEducationIdentitySynchronizationConfigurationImpl {
	return s
}

var _ EducationIdentitySynchronizationConfiguration = RawEducationIdentitySynchronizationConfigurationImpl{}

// RawEducationIdentitySynchronizationConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationIdentitySynchronizationConfigurationImpl struct {
	educationIdentitySynchronizationConfiguration BaseEducationIdentitySynchronizationConfigurationImpl
	Type                                          string
	Values                                        map[string]interface{}
}

func (s RawEducationIdentitySynchronizationConfigurationImpl) EducationIdentitySynchronizationConfiguration() BaseEducationIdentitySynchronizationConfigurationImpl {
	return s.educationIdentitySynchronizationConfiguration
}

func UnmarshalEducationIdentitySynchronizationConfigurationImplementation(input []byte) (EducationIdentitySynchronizationConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationIdentitySynchronizationConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationIdentityCreationConfiguration") {
		var out EducationIdentityCreationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationIdentityCreationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationIdentityMatchingConfiguration") {
		var out EducationIdentityMatchingConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationIdentityMatchingConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationIdentitySynchronizationConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationIdentitySynchronizationConfigurationImpl: %+v", err)
	}

	return RawEducationIdentitySynchronizationConfigurationImpl{
		educationIdentitySynchronizationConfiguration: parent,
		Type:   value,
		Values: temp,
	}, nil

}
