package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationSynchronizationProfileStatus{}

type EducationSynchronizationProfileStatus struct {
	// Number of errors during synchronization.
	ErrorCount *int64 `json:"errorCount,omitempty"`

	// Date and time when most recent changes were observed in the profile.
	LastActivityDateTime nullable.Type[string] `json:"lastActivityDateTime,omitempty"`

	// Date and time of the most recent successful synchronization.
	LastSynchronizationDateTime nullable.Type[string] `json:"lastSynchronizationDateTime,omitempty"`

	// The status of a sync. The possible values are: paused, inProgress, success, error, validationError, quarantined,
	// unknownFutureValue, extracting, validating. Note that you must use the Prefer: include-unknown-enum-members request
	// header to get the following values in this evolvable enum: extracting, validating.
	Status *EducationSynchronizationStatus `json:"status,omitempty"`

	// Status message for the synchronization stage of the current profile.
	StatusMessage *string `json:"statusMessage,omitempty"`

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

func (s EducationSynchronizationProfileStatus) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationSynchronizationProfileStatus{}

func (s EducationSynchronizationProfileStatus) MarshalJSON() ([]byte, error) {
	type wrapper EducationSynchronizationProfileStatus
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationSynchronizationProfileStatus: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationSynchronizationProfileStatus: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationSynchronizationProfileStatus"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationSynchronizationProfileStatus: %+v", err)
	}

	return encoded, nil
}
