package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookDocumentTask{}

type WorkbookDocumentTask struct {
	// A collection of user identities the task is assigned to.
	Assignees *[]WorkbookEmailIdentity `json:"assignees,omitempty"`

	// A collection of task change histories.
	Changes *[]WorkbookDocumentTaskChange `json:"changes,omitempty"`

	// The comment that the task is associated with.
	Comment *WorkbookComment `json:"comment,omitempty"`

	// The identity of the user who completed the task. Nullable.
	CompletedBy *WorkbookEmailIdentity `json:"completedBy,omitempty"`

	// Date and time when the task was completed. Nullable. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// A user identity that creates the task. Nullable.
	CreatedBy *WorkbookEmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the task was created. Nullable. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// An integer value from 0 to 100 that represents the percentage of the completion of the task. 100 means that the task
	// is completed. Nullable.
	PercentComplete nullable.Type[int64] `json:"percentComplete,omitempty"`

	// An integer value from 0 to 10 that represents the priority of the task. A lower value indicates a higher priority.
	// Nullable.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// Start and due date of the task. Nullable.
	StartAndDueDateTime *WorkbookDocumentTaskSchedule `json:"startAndDueDateTime,omitempty"`

	// The title of the task.
	Title nullable.Type[string] `json:"title,omitempty"`

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

func (s WorkbookDocumentTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookDocumentTask{}

func (s WorkbookDocumentTask) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookDocumentTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookDocumentTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookDocumentTask: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookDocumentTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookDocumentTask: %+v", err)
	}

	return encoded, nil
}
