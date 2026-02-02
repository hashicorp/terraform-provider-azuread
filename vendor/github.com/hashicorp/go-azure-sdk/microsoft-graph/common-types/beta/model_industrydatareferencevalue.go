package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataReferenceValue interface {
	IndustryDataReferenceValue() BaseIndustryDataReferenceValueImpl
}

var _ IndustryDataReferenceValue = BaseIndustryDataReferenceValueImpl{}

type BaseIndustryDataReferenceValueImpl struct {
	// The code of the desired referenceDefinition entry.
	Code nullable.Type[string] `json:"code,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Value *IndustryDataReferenceDefinition `json:"value,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIndustryDataReferenceValueImpl) IndustryDataReferenceValue() BaseIndustryDataReferenceValueImpl {
	return s
}

var _ IndustryDataReferenceValue = RawIndustryDataReferenceValueImpl{}

// RawIndustryDataReferenceValueImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataReferenceValueImpl struct {
	industryDataReferenceValue BaseIndustryDataReferenceValueImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawIndustryDataReferenceValueImpl) IndustryDataReferenceValue() BaseIndustryDataReferenceValueImpl {
	return s.industryDataReferenceValue
}

func UnmarshalIndustryDataReferenceValueImplementation(input []byte) (IndustryDataReferenceValue, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataReferenceValue into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.fileFormatReferenceValue") {
		var out IndustryDataFileFormatReferenceValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataFileFormatReferenceValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.identifierTypeReferenceValue") {
		var out IndustryDataIdentifierTypeReferenceValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataIdentifierTypeReferenceValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.roleReferenceValue") {
		var out IndustryDataRoleReferenceValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataRoleReferenceValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.sectionRoleReferenceValue") {
		var out IndustryDataSectionRoleReferenceValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataSectionRoleReferenceValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.userMatchTargetReferenceValue") {
		var out IndustryDataUserMatchTargetReferenceValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataUserMatchTargetReferenceValue: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.yearReferenceValue") {
		var out IndustryDataYearReferenceValue
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataYearReferenceValue: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataReferenceValueImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataReferenceValueImpl: %+v", err)
	}

	return RawIndustryDataReferenceValueImpl{
		industryDataReferenceValue: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
