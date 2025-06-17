package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PlannerTask = BusinessScenarioTask{}

type BusinessScenarioTask struct {
	// Scenario-specific properties of the task. externalObjectId and externalBucketId properties must be specified when
	// creating a task.
	BusinessScenarioProperties *BusinessScenarioProperties `json:"businessScenarioProperties,omitempty"`

	// Target of the task that specifies where the task should be placed. Must be specified when creating a task.
	Target BusinessScenarioTaskTargetBase `json:"target"`

	// Fields inherited from PlannerTask

	// The number of checklist items with value set to false, representing incomplete items.
	ActiveChecklistItemCount nullable.Type[int64] `json:"activeChecklistItemCount,omitempty"`

	// The categories to which the task is applied. See plannerAppliedCategories resource type for possible values.
	AppliedCategories *PlannerAppliedCategories `json:"appliedCategories,omitempty"`

	// Read-only. Nullable. Contains information about who archived or unarchived the task and why.
	ArchivalInfo *PlannerArchivalInfo `json:"archivalInfo,omitempty"`

	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
	AssignedToTaskBoardFormat *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`

	// A hint that is used to order items of this type in a list view. For more information, see Using order hints in
	// planner.
	AssigneePriority nullable.Type[string] `json:"assigneePriority,omitempty"`

	// The set of assignees the task is assigned to.
	Assignments *PlannerAssignments `json:"assignments,omitempty"`

	// Bucket ID to which the task belongs. The bucket needs to be in the same plan as the task. The value of the bucketId
	// property is 28 characters long and case-sensitive. Format validation is done on the service.
	BucketId nullable.Type[string] `json:"bucketId,omitempty"`

	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
	BucketTaskBoardFormat *PlannerBucketTaskBoardTaskFormat `json:"bucketTaskBoardFormat,omitempty"`

	// The number of checklist items that are present on the task.
	ChecklistItemCount nullable.Type[int64] `json:"checklistItemCount,omitempty"`

	// The identity of the user that completed the task.
	CompletedBy IdentitySet `json:"completedBy"`

	// Read-only. The date and time at which the 'percentComplete' of the task is set to '100'. The Timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z
	CompletedDateTime nullable.Type[string] `json:"completedDateTime,omitempty"`

	// The thread ID of the conversation on the task. This is the ID of the conversation thread object created in the group.
	ConversationThreadId nullable.Type[string] `json:"conversationThreadId,omitempty"`

	// The identity of the user who created the task.
	CreatedBy IdentitySet `json:"createdBy"`

	// Read-only. The date and time at which the task is created. The Timestamp type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Information about the origin of the task.
	CreationSource PlannerTaskCreation `json:"creationSource"`

	// Read-only. Nullable. More details about the task.
	Details *PlannerTaskDetails `json:"details,omitempty"`

	// The date and time at which the task is due. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DueDateTime nullable.Type[string] `json:"dueDateTime,omitempty"`

	// Read-only. This value is true if the details object of the task has a nonempty description. Otherwise,false.
	HasDescription nullable.Type[bool] `json:"hasDescription,omitempty"`

	// Read-only. If set to true, the task is archived. An archived task is read-only.
	IsArchived nullable.Type[bool] `json:"isArchived,omitempty"`

	// Indicates whether to show this task in the MyDay view. If true, it shows the task.
	IsOnMyDay nullable.Type[bool] `json:"isOnMyDay,omitempty"`

	// Read-only. The date on which task is added to or removed from MyDay.
	IsOnMyDayLastModifiedDate nullable.Type[string] `json:"isOnMyDayLastModifiedDate,omitempty"`

	LastModifiedBy       IdentitySet           `json:"lastModifiedBy"`
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The hint used to order items of this type in a list view. For more information, see Using order hints in plannern.
	OrderHint nullable.Type[string] `json:"orderHint,omitempty"`

	// The percentage of task completion. When set to 100, the task is completed.
	PercentComplete nullable.Type[int64] `json:"percentComplete,omitempty"`

	// Plan ID to which the task belongs.
	PlanId nullable.Type[string] `json:"planId,omitempty"`

	// The type of preview that shows up on the task. Possible values are: automatic, noPreview, checklist, description,
	// reference.
	PreviewType *PlannerPreviewType `json:"previewType,omitempty"`

	// The priority of the task. Valid values are between 0 and 10, inclusive. Larger values indicate lower priority. For
	// example, 0 has the highest priority and 10 has the lowest priority. Currently, planner interprets values 0 and 1 as
	// 'urgent', 2 and 3 and 4 as 'important', 5, 6, and 7 as 'medium', and 8, 9, and 10 as 'low'. Currently, planner sets
	// the value 1 for 'urgent', 3 for 'important', 5 for 'medium', and 9 for 'low'.
	Priority nullable.Type[int64] `json:"priority,omitempty"`

	// Read-only. Nullable. Used to render the task correctly in the task board view when grouped by progress.
	ProgressTaskBoardFormat *PlannerProgressTaskBoardTaskFormat `json:"progressTaskBoardFormat,omitempty"`

	// Defines active or inactive recurrence for the task. null when the recurrence has never been defined for the task.
	Recurrence *PlannerTaskRecurrence `json:"recurrence,omitempty"`

	// Number of external references that exist on the task.
	ReferenceCount nullable.Type[int64] `json:"referenceCount,omitempty"`

	// Indicates all the requirements specified on the plannerTask. Possible values are: none, checklistCompletion,
	// unknownFutureValue, formCompletion, approvalCompletion. Read-only. Use the Prefer: include-unknown-enum-members
	// request header to get the following values in this evolvable enum: formCompletion, approvalCompletion. The
	// plannerTaskCompletionRequirementDetails in plannerTaskDetails has details of the requirements specified, if any.
	SpecifiedCompletionRequirements *PlannerTaskCompletionRequirements `json:"specifiedCompletionRequirements,omitempty"`

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

func (s BusinessScenarioTask) PlannerTask() BasePlannerTaskImpl {
	return BasePlannerTaskImpl{
		ActiveChecklistItemCount:        s.ActiveChecklistItemCount,
		AppliedCategories:               s.AppliedCategories,
		ArchivalInfo:                    s.ArchivalInfo,
		AssignedToTaskBoardFormat:       s.AssignedToTaskBoardFormat,
		AssigneePriority:                s.AssigneePriority,
		Assignments:                     s.Assignments,
		BucketId:                        s.BucketId,
		BucketTaskBoardFormat:           s.BucketTaskBoardFormat,
		ChecklistItemCount:              s.ChecklistItemCount,
		CompletedBy:                     s.CompletedBy,
		CompletedDateTime:               s.CompletedDateTime,
		ConversationThreadId:            s.ConversationThreadId,
		CreatedBy:                       s.CreatedBy,
		CreatedDateTime:                 s.CreatedDateTime,
		CreationSource:                  s.CreationSource,
		Details:                         s.Details,
		DueDateTime:                     s.DueDateTime,
		HasDescription:                  s.HasDescription,
		IsArchived:                      s.IsArchived,
		IsOnMyDay:                       s.IsOnMyDay,
		IsOnMyDayLastModifiedDate:       s.IsOnMyDayLastModifiedDate,
		LastModifiedBy:                  s.LastModifiedBy,
		LastModifiedDateTime:            s.LastModifiedDateTime,
		OrderHint:                       s.OrderHint,
		PercentComplete:                 s.PercentComplete,
		PlanId:                          s.PlanId,
		PreviewType:                     s.PreviewType,
		Priority:                        s.Priority,
		ProgressTaskBoardFormat:         s.ProgressTaskBoardFormat,
		Recurrence:                      s.Recurrence,
		ReferenceCount:                  s.ReferenceCount,
		SpecifiedCompletionRequirements: s.SpecifiedCompletionRequirements,
		StartDateTime:                   s.StartDateTime,
		Title:                           s.Title,
		Id:                              s.Id,
		ODataId:                         s.ODataId,
		ODataType:                       s.ODataType,
	}
}

func (s BusinessScenarioTask) PlannerDelta() BasePlannerDeltaImpl {
	return BasePlannerDeltaImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s BusinessScenarioTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = BusinessScenarioTask{}

func (s BusinessScenarioTask) MarshalJSON() ([]byte, error) {
	type wrapper BusinessScenarioTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BusinessScenarioTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BusinessScenarioTask: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.businessScenarioTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BusinessScenarioTask: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BusinessScenarioTask{}

func (s *BusinessScenarioTask) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		BusinessScenarioProperties      *BusinessScenarioProperties           `json:"businessScenarioProperties,omitempty"`
		ActiveChecklistItemCount        nullable.Type[int64]                  `json:"activeChecklistItemCount,omitempty"`
		AppliedCategories               *PlannerAppliedCategories             `json:"appliedCategories,omitempty"`
		ArchivalInfo                    *PlannerArchivalInfo                  `json:"archivalInfo,omitempty"`
		AssignedToTaskBoardFormat       *PlannerAssignedToTaskBoardTaskFormat `json:"assignedToTaskBoardFormat,omitempty"`
		AssigneePriority                nullable.Type[string]                 `json:"assigneePriority,omitempty"`
		Assignments                     *PlannerAssignments                   `json:"assignments,omitempty"`
		BucketId                        nullable.Type[string]                 `json:"bucketId,omitempty"`
		BucketTaskBoardFormat           *PlannerBucketTaskBoardTaskFormat     `json:"bucketTaskBoardFormat,omitempty"`
		ChecklistItemCount              nullable.Type[int64]                  `json:"checklistItemCount,omitempty"`
		CompletedDateTime               nullable.Type[string]                 `json:"completedDateTime,omitempty"`
		ConversationThreadId            nullable.Type[string]                 `json:"conversationThreadId,omitempty"`
		CreatedDateTime                 nullable.Type[string]                 `json:"createdDateTime,omitempty"`
		Details                         *PlannerTaskDetails                   `json:"details,omitempty"`
		DueDateTime                     nullable.Type[string]                 `json:"dueDateTime,omitempty"`
		HasDescription                  nullable.Type[bool]                   `json:"hasDescription,omitempty"`
		IsArchived                      nullable.Type[bool]                   `json:"isArchived,omitempty"`
		IsOnMyDay                       nullable.Type[bool]                   `json:"isOnMyDay,omitempty"`
		IsOnMyDayLastModifiedDate       nullable.Type[string]                 `json:"isOnMyDayLastModifiedDate,omitempty"`
		LastModifiedDateTime            nullable.Type[string]                 `json:"lastModifiedDateTime,omitempty"`
		OrderHint                       nullable.Type[string]                 `json:"orderHint,omitempty"`
		PercentComplete                 nullable.Type[int64]                  `json:"percentComplete,omitempty"`
		PlanId                          nullable.Type[string]                 `json:"planId,omitempty"`
		PreviewType                     *PlannerPreviewType                   `json:"previewType,omitempty"`
		Priority                        nullable.Type[int64]                  `json:"priority,omitempty"`
		ProgressTaskBoardFormat         *PlannerProgressTaskBoardTaskFormat   `json:"progressTaskBoardFormat,omitempty"`
		Recurrence                      *PlannerTaskRecurrence                `json:"recurrence,omitempty"`
		ReferenceCount                  nullable.Type[int64]                  `json:"referenceCount,omitempty"`
		SpecifiedCompletionRequirements *PlannerTaskCompletionRequirements    `json:"specifiedCompletionRequirements,omitempty"`
		StartDateTime                   nullable.Type[string]                 `json:"startDateTime,omitempty"`
		Title                           *string                               `json:"title,omitempty"`
		Id                              *string                               `json:"id,omitempty"`
		ODataId                         *string                               `json:"@odata.id,omitempty"`
		ODataType                       *string                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.BusinessScenarioProperties = decoded.BusinessScenarioProperties
	s.ActiveChecklistItemCount = decoded.ActiveChecklistItemCount
	s.AppliedCategories = decoded.AppliedCategories
	s.ArchivalInfo = decoded.ArchivalInfo
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
	s.Id = decoded.Id
	s.IsArchived = decoded.IsArchived
	s.IsOnMyDay = decoded.IsOnMyDay
	s.IsOnMyDayLastModifiedDate = decoded.IsOnMyDayLastModifiedDate
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.OrderHint = decoded.OrderHint
	s.PercentComplete = decoded.PercentComplete
	s.PlanId = decoded.PlanId
	s.PreviewType = decoded.PreviewType
	s.Priority = decoded.Priority
	s.ProgressTaskBoardFormat = decoded.ProgressTaskBoardFormat
	s.Recurrence = decoded.Recurrence
	s.ReferenceCount = decoded.ReferenceCount
	s.SpecifiedCompletionRequirements = decoded.SpecifiedCompletionRequirements
	s.StartDateTime = decoded.StartDateTime
	s.Title = decoded.Title

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BusinessScenarioTask into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["completedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CompletedBy' for 'BusinessScenarioTask': %+v", err)
		}
		s.CompletedBy = impl
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'BusinessScenarioTask': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["creationSource"]; ok {
		impl, err := UnmarshalPlannerTaskCreationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreationSource' for 'BusinessScenarioTask': %+v", err)
		}
		s.CreationSource = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'BusinessScenarioTask': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalBusinessScenarioTaskTargetBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'BusinessScenarioTask': %+v", err)
		}
		s.Target = impl
	}

	return nil
}
