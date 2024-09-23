package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationAssignmentRecipient = EducationAssignmentIndividualRecipient{}

type EducationAssignmentIndividualRecipient struct {
	// A collection of ids of the recipients.
	Recipients *[]string `json:"recipients,omitempty"`

	// Fields inherited from EducationAssignmentRecipient

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationAssignmentIndividualRecipient) EducationAssignmentRecipient() BaseEducationAssignmentRecipientImpl {
	return BaseEducationAssignmentRecipientImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentIndividualRecipient{}

func (s EducationAssignmentIndividualRecipient) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentIndividualRecipient
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentIndividualRecipient: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentIndividualRecipient: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentIndividualRecipient"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentIndividualRecipient: %+v", err)
	}

	return encoded, nil
}
