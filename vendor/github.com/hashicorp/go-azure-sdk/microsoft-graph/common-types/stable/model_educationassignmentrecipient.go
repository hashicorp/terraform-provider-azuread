package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentRecipient interface {
	EducationAssignmentRecipient() BaseEducationAssignmentRecipientImpl
}

var _ EducationAssignmentRecipient = BaseEducationAssignmentRecipientImpl{}

type BaseEducationAssignmentRecipientImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationAssignmentRecipientImpl) EducationAssignmentRecipient() BaseEducationAssignmentRecipientImpl {
	return s
}

var _ EducationAssignmentRecipient = RawEducationAssignmentRecipientImpl{}

// RawEducationAssignmentRecipientImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEducationAssignmentRecipientImpl struct {
	educationAssignmentRecipient BaseEducationAssignmentRecipientImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawEducationAssignmentRecipientImpl) EducationAssignmentRecipient() BaseEducationAssignmentRecipientImpl {
	return s.educationAssignmentRecipient
}

func (s RawEducationAssignmentRecipientImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEducationAssignmentRecipientImplementation(input []byte) (EducationAssignmentRecipient, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentRecipient into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentClassRecipient") {
		var out EducationAssignmentClassRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentClassRecipient: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentGroupRecipient") {
		var out EducationAssignmentGroupRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentGroupRecipient: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.educationAssignmentIndividualRecipient") {
		var out EducationAssignmentIndividualRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationAssignmentIndividualRecipient: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationAssignmentRecipientImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationAssignmentRecipientImpl: %+v", err)
	}

	return RawEducationAssignmentRecipientImpl{
		educationAssignmentRecipient: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
