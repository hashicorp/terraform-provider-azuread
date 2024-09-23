package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SynchronizationJob{}

type SynchronizationJob struct {
	// The bulk upload operation for the job.
	BulkUpload *BulkUpload `json:"bulkUpload,omitempty"`

	// Schedule used to run the job. Read-only.
	Schedule *SynchronizationSchedule `json:"schedule,omitempty"`

	// The synchronization schema configured for the job.
	Schema *SynchronizationSchema `json:"schema,omitempty"`

	// Status of the job, which includes when the job was last run, current job state, and errors.
	Status *SynchronizationStatus `json:"status,omitempty"`

	// Settings associated with the job. Some settings are inherited from the template.
	SynchronizationJobSettings *[]KeyValuePair `json:"synchronizationJobSettings,omitempty"`

	// Identifier of the synchronization template this job is based on.
	TemplateId nullable.Type[string] `json:"templateId,omitempty"`

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

func (s SynchronizationJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SynchronizationJob{}

func (s SynchronizationJob) MarshalJSON() ([]byte, error) {
	type wrapper SynchronizationJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SynchronizationJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SynchronizationJob: %+v", err)
	}

	delete(decoded, "schedule")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.synchronizationJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SynchronizationJob: %+v", err)
	}

	return encoded, nil
}
