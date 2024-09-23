package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PlannerTask{}

type PlannerTask struct {
	// Number of checklist items with value set to false, representing incomplete items.
	ActiveChecklistItemCount nullable.Type[int64] `json:"activeChecklistItemCount,omitempty"`

	// The categories to which the task has been applied. See applied Categories for possible values.
	AppliedCategories *PlannerAppliedCategories `json:"appliedCategories,omitempty"`

	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
	AssignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`

	// Hint used to order items of this type in a list view. The format is defined as outlined here.
	AssigneePriority nullable.Type[string] `json:"assigneePriority,omitempty"`

	// The set of assignees the task is assigned to.
	Assignments *PlannerAssignments `json:"assignments,omitempty"`

	// Bucket ID to which the task belongs. The bucket needs to be in the plan that the task is in. It's 28 characters long
	// and case-sensitive. Format validation is done on the service.
	BucketId nullable.Type[string] `json:"bucketId,omitempty"`

	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
	BucketTaskBoardFormat *PlannerBucketTaskBoardTaskFormat `json:"bucketTaskBoardFormat,omitempty"`

	// Number of checklist items that are present on the task.
	ChecklistItemCount nullable.Type[int64] `json:"checklistItemCount,omitempty"`

	// Identity of the user that completed the task.
	CompletedBy IdentitySet `json:"completedBy"`

	// Read-only. Date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type represents
	// date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014
	// is 2014-01-01T00:00:00Z
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// Thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
	ConversationThreadId nullable.Type[string] `json:"conversationThreadId,omitempty"`

	// Identity of the user that created the task.
	CreatedBy IdentitySet `json:"createdBy"`

	// Read-only. Date and time at which the task is created. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Read-only. Nullable. More details about the task.
	Details *PlannerTaskDetails `json:"details,omitempty"`

	// Date and time at which the task is due. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DueDateTime nullable.Type[string] `json:"dueDateTime,omitempty"`

	// Read-only. Value is true if the details object of the task has a nonempty description and false otherwise.
	HasDescription nullable.Type[bool] `json:"hasDescription,omitempty"`

	// Hint used to order items of this type in a list view. The format is defined as outlined here.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`

	// Percentage of task completion. When set to 100, the task is considered completed.
	PercentComplete nullable.Type[int64] `json:"percentComplete,omitempty"`

	// Plan ID to which the task belongs.
	PlanId nullable.Type[string] `json:"planId,omitempty"`

	// This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist,
	// description, reference.
	PreviewType *PlannerPreviewType `json:"previewType,omitempty"`

	// Priority of the task. The valid range of values is between 0 and 10, with the increasing value being lower priority
	// (0 has the highest priority and 10 has the lowest priority). Currently, Planner interprets values 0 and 1 as
	// 'urgent', 2, 3 and 4 as 'important', 5, 6, and 7 as 'medium', and 8, 9, and 10 as 'low'. Additionally, Planner sets
	// the value 1 for 'urgent', 3 for 'important', 5 for 'medium', and 9 for 'low'.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
	ProgressTaskBoardFormat *PlannerProgressTaskBoardTaskFormat `json:"progressTaskBoardFormat,omitempty"`

	// Number of external references that exist on the task.
	ReferenceCount nullable.Type[int64] `json:"referenceCount,omitempty"`

	// Date and time at which the task starts. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// Title of the task.
	Title *string `json:"title,omitempty"`

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

func (s PlannerTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PlannerTask{}

func (s PlannerTask) MarshalJSON() ([]byte, error) {
	type wrapper PlannerTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PlannerTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PlannerTask: %+v", err)
	}

	delete(decoded, "assignedToTaskBoardFormat")
	delete(decoded, "bucketTaskBoardFormat")
	delete(decoded, "completedDateTime")
	delete(decoded, "createdDateTime")
	delete(decoded, "details")
	delete(decoded, "hasDescription")
	delete(decoded, "progressTaskBoardFormat")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.plannerTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PlannerTask: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PlannerTask{}

func (s *PlannerTask) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActiveChecklistItemCount  nullable.Type[int64]                  `json:"activeChecklistItemCount,omitempty"`
		AppliedCategories         *PlannerAppliedCategories             `json:"appliedCategories,omitempty"`
		AssignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`
		AssigneePriority          nullable.Type[string]                 `json:"assigneePriority,omitempty"`
		Assignments               *PlannerAssignments                   `json:"assignments,omitempty"`
		BucketId                  nullable.Type[string]                 `json:"bucketId,omitempty"`
		BucketTaskBoardFormat     *PlannerBucketTaskBoardTaskFormat     `json:"bucketTaskBoardFormat,omitempty"`
		ChecklistItemCount        nullable.Type[int64]                  `json:"checklistItemCount,omitempty"`
		CompletedDateTime         nullable.Type[string]                 `json:"completedDateTime,omitempty"`
		ConversationThreadId      nullable.Type[string]                 `json:"conversationThreadId,omitempty"`
		CreatedDateTime           nullable.Type[string]                 `json:"createdDateTime,omitempty"`
		Details                   *PlannerTaskDetails                   `json:"details,omitempty"`
		DueDateTime               nullable.Type[string]                 `json:"dueDateTime,omitempty"`
		HasDescription            nullable.Type[bool]                   `json:"hasDescription,omitempty"`
		OrderHint                 nullable.Type[string]                 `json:"orderHint,omitempty"`
		PercentComplete           nullable.Type[int64]                  `json:"percentComplete,omitempty"`
		PlanId                    nullable.Type[string]                 `json:"planId,omitempty"`
		PreviewType               *PlannerPreviewType                   `json:"previewType,omitempty"`
		Priority                  nullable.Type[int64]                  `json:"priority,omitempty"`
		ProgressTaskBoardFormat   *PlannerProgressTaskBoardTaskFormat   `json:"progressTaskBoardFormat,omitempty"`
		ReferenceCount            nullable.Type[int64]                  `json:"referenceCount,omitempty"`
		StartDateTime             nullable.Type[string]                 `json:"startDateTime,omitempty"`
		Title                     *string                               `json:"title,omitempty"`
		Id                        *string                               `json:"id,omitempty"`
		ODataId                   *string                               `json:"@odata.id,omitempty"`
		ODataType                 *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActiveChecklistItemCount = decoded.ActiveChecklistItemCount
	s.AppliedCategories = decoded.AppliedCategories
	s.AssignedToTaskBoardFormat = decoded.AssignedToTaskBoardFormat
	s.AssigneePriority = decoded.AssigneePriority
	s.Assignments = decoded.Assignments
	s.BucketId = decoded.BucketId
	s.BucketTaskBoardFormat = decoded.BucketTaskBoardFormat
	s.ChecklistItemCount = decoded.ChecklistItemCount
	s.CompletedDateTime = decoded.CompletedDateTime
	s.ConversationThreadId = decoded.ConversationThreadId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Details = decoded.Details
	s.DueDateTime = decoded.DueDateTime
	s.HasDescription = decoded.HasDescription
	s.OrderHint = decoded.OrderHint
	s.PercentComplete = decoded.PercentComplete
	s.PlanId = decoded.PlanId
	s.PreviewType = decoded.PreviewType
	s.Priority = decoded.Priority
	s.ProgressTaskBoardFormat = decoded.ProgressTaskBoardFormat
	s.ReferenceCount = decoded.ReferenceCount
	s.StartDateTime = decoded.StartDateTime
	s.Title = decoded.Title
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerTask into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["completedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CompletedBy' for 'PlannerTask': %+v", err)
		}
		s.CompletedBy = impl
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'PlannerTask': %+v", err)
		}
		s.CreatedBy = impl
	}

	return nil
}
