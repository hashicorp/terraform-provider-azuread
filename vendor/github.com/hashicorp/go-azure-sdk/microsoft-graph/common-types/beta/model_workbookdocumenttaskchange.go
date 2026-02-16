package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WorkbookDocumentTaskChange{}

type WorkbookDocumentTaskChange struct {
	// The user identity the task is assigned to. Only present when the type property is assign. Nullable.
	Assignee *WorkbookEmailIdentity `json:"assignee,omitempty"`

	ChangedBy *WorkbookEmailIdentity `json:"changedBy,omitempty"`

	// The identifier of the associated comment.
	CommentId nullable.Type[string] `json:"commentId,omitempty"`

	// Date and time when the task was changed. Nullable. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The due date and time for the task. Only present when the type property is setSchedule. Nullable. The Timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z.
	DueDateTime nullable.Type[string] `json:"dueDateTime,omitempty"`

	// An integer value from 0 to 100 that represents the percentage of the completion of the task and associated comment.
	// 100 means that the task and associated comment are completed. If you change the completion from 100 to a lower value,
	// the associated task and comment are reactivated. Only present when the type property is setPercentComplete. Nullable.
	PercentComplete nullable.Type[int64] `json:"percentComplete,omitempty"`

	// An integer value from 0 to 10 that represents the priority of the task. A lower value indicates a higher priority. 5
	// indicates the default priority if not set. Only present when the type property is setPriority. Nullable.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// The start date and time for the task. Only present when the type property is setSchedule. Nullable. The Timestamp
	// type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC
	// on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The title of the task. Only present when the type property is setTitle. Nullable.
	Title nullable.Type[string] `json:"title,omitempty"`

	// The type of the change history. Possible values are: create, assign, unassign, unassignAll, setPriority, setTitle,
	// setPercentComplete, setSchedule, remove, restore, undo.
	Type *string `json:"type,omitempty"`

	// The ID of the workbookDocumentTaskChange that was undone for the undo change action. Only exists on an undo change
	// history. Nullable.
	UndoChangeId nullable.Type[string] `json:"undoChangeId,omitempty"`

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

func (s WorkbookDocumentTaskChange) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WorkbookDocumentTaskChange{}

func (s WorkbookDocumentTaskChange) MarshalJSON() ([]byte, error) {
	type wrapper WorkbookDocumentTaskChange
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WorkbookDocumentTaskChange: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WorkbookDocumentTaskChange: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.workbookDocumentTaskChange"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WorkbookDocumentTaskChange: %+v", err)
	}

	return encoded, nil
}
