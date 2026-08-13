package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentGradeType interface {
	EducationAssignmentGradeType() BaseEducationAssignmentGradeTypeImpl
}

var _ EducationAssignmentGradeType = BaseEducationAssignmentGradeTypeImpl{}

type BaseEducationAssignmentGradeTypeImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationAssignmentGradeTypeImpl) EducationAssignmentGradeType() BaseEducationAssignmentGradeTypeImpl {
	return s
}

var _ EducationAssignmentGradeType = RawEducationAssignmentGradeTypeImpl{}

// RawEducationAssignmentGradeTypeImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEducationAssignmentGradeTypeImpl struct {
	educationAssignmentGradeType BaseEducationAssignmentGradeTypeImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawEducationAssignmentGradeTypeImpl) EducationAssignmentGradeType() BaseEducationAssignmentGradeTypeImpl {
	return s.educationAssignmentGradeType
}

func (s RawEducationAssignmentGradeTypeImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEducationAssignmentGradeTypeImplementation(input []byte) (EducationAssignmentGradeType, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentGradeType into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentPointsGradeType") {
		var out EducationAssignmentPointsGradeType
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentPointsGradeType: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationAssignmentGradeTypeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationAssignmentGradeTypeImpl: %+v", err)
	}

	return RawEducationAssignmentGradeTypeImpl{
		educationAssignmentGradeType: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
