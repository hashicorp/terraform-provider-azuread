package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSubmissionRecipient interface {
	EducationSubmissionRecipient() BaseEducationSubmissionRecipientImpl
}

var _ EducationSubmissionRecipient = BaseEducationSubmissionRecipientImpl{}

type BaseEducationSubmissionRecipientImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseEducationSubmissionRecipientImpl) EducationSubmissionRecipient() BaseEducationSubmissionRecipientImpl {
	return s
}

var _ EducationSubmissionRecipient = RawEducationSubmissionRecipientImpl{}

// RawEducationSubmissionRecipientImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEducationSubmissionRecipientImpl struct {
	educationSubmissionRecipient BaseEducationSubmissionRecipientImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawEducationSubmissionRecipientImpl) EducationSubmissionRecipient() BaseEducationSubmissionRecipientImpl {
	return s.educationSubmissionRecipient
}

func (s RawEducationSubmissionRecipientImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEducationSubmissionRecipientImplementation(input []byte) (EducationSubmissionRecipient, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSubmissionRecipient into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.educationSubmissionIndividualRecipient") {
		var out EducationSubmissionIndividualRecipient
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EducationSubmissionIndividualRecipient: %+v", err)
		}
		return out, nil
	}

	var parent BaseEducationSubmissionRecipientImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEducationSubmissionRecipientImpl: %+v", err)
	}

	return RawEducationSubmissionRecipientImpl{
		educationSubmissionRecipient: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
