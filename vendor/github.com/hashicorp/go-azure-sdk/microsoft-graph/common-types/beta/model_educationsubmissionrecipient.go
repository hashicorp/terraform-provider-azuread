package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
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

// RawEducationSubmissionRecipientImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEducationSubmissionRecipientImpl struct {
	educationSubmissionRecipient BaseEducationSubmissionRecipientImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawEducationSubmissionRecipientImpl) EducationSubmissionRecipient() BaseEducationSubmissionRecipientImpl {
	return s.educationSubmissionRecipient
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
