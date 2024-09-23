package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationSynchronizationError{}

type EducationSynchronizationError struct {
	// Represents the sync entity (school, section, student, teacher).
	EntryType nullable.Type[string] `json:"entryType,omitempty"`

	// Represents the error code for this error.
	ErrorCode nullable.Type[string] `json:"errorCode,omitempty"`

	// Contains a description of the error.
	ErrorMessage nullable.Type[string] `json:"errorMessage,omitempty"`

	// The unique identifier for the entry.
	JoiningValue nullable.Type[string] `json:"joiningValue,omitempty"`

	// The time of occurrence of this error.
	RecordedDateTime nullable.Type[string] `json:"recordedDateTime,omitempty"`

	// The identifier of this error entry.
	ReportableIdentifier nullable.Type[string] `json:"reportableIdentifier,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s EducationSynchronizationError) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSynchronizationError{}

func (s EducationSynchronizationError) MarshalJSON() ([]byte, error) {
	type wrapper EducationSynchronizationError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSynchronizationError: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationError: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSynchronizationError"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSynchronizationError: %+v", err)
	}

	return encoded, nil
}
