package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EducationAssignmentRecipient = EducationAssignmentClassRecipient{}

type EducationAssignmentClassRecipient struct {

	// Fields inherited from EducationAssignmentRecipient

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationAssignmentClassRecipient) EducationAssignmentRecipient() BaseEducationAssignmentRecipientImpl {
	return BaseEducationAssignmentRecipientImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentClassRecipient{}

func (s EducationAssignmentClassRecipient) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentClassRecipient
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentClassRecipient: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentClassRecipient: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentClassRecipient"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentClassRecipient: %+v", err)
	}

	return encoded, nil
}
