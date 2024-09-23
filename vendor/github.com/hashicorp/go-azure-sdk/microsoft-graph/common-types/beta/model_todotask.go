package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TodoTask{}

type TodoTask struct {
	AttachmentSessions *[]AttachmentSession `json:"attachmentSessions,omitempty"`

	// A collection of file attachments for the task.
	Attachments *[]AttachmentBase `json:"attachments,omitempty"`

	// The task body that typically contains information about the task.
	Body *ItemBody `json:"body,omitempty"`

	// The date and time when the task body was last modified. By default, it is in UTC. You can provide a custom time zone
	// in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC
	// on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
	BodyLastModifiedDateTime *string `json:"bodyLastModifiedDateTime,omitempty"`

	// The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory
	// that the user has defined.
	Categories *[]string `json:"categories,omitempty"`

	// A collection of smaller subtasks linked to the more complex parent task.
	ChecklistItems *[]ChecklistItem `json:"checklistItems,omitempty"`

	// The date and time in the specified time zone that the task was finished.
	CompletedDateTime *DateTimeTimeZone `json:"completedDateTime,omitempty"`

	// The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the
	// request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like
	// this: '2020-01-01T00:00:00Z'.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The date and time in the specified time zone that the task is to be finished.
	DueDateTime *DateTimeTimeZone `json:"dueDateTime,omitempty"`

	// The collection of open extensions defined for the task. Nullable.
	Extensions *[]Extension `json:"extensions,omitempty"`

	// Indicates whether the task has attachments.
	HasAttachments nullable.Type[bool] `json:"hasAttachments,omitempty"`

	Importance *Importance `json:"importance,omitempty"`

	// Set to true if an alert is set to remind the user of the task.
	IsReminderOn *bool `json:"isReminderOn,omitempty"`

	// The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in
	// the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on
	// Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// A collection of resources linked to the task.
	LinkedResources *[]LinkedResource `json:"linkedResources,omitempty"`

	// The recurrence pattern for the task.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// The date and time in the specified time zone for a reminder alert of the task to occur.
	ReminderDateTime *DateTimeTimeZone `json:"reminderDateTime,omitempty"`

	// The date and time in the specified time zone at which the task is scheduled to start.
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`

	Status *TaskStatus `json:"status,omitempty"`

	// A brief description of the task.
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

func (s TodoTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TodoTask{}

func (s TodoTask) MarshalJSON() ([]byte, error) {
	type wrapper TodoTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TodoTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TodoTask: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.todoTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TodoTask: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &TodoTask{}

func (s *TodoTask) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AttachmentSessions       *[]AttachmentSession  `json:"attachmentSessions,omitempty"`
		Body                     *ItemBody             `json:"body,omitempty"`
		BodyLastModifiedDateTime *string               `json:"bodyLastModifiedDateTime,omitempty"`
		Categories               *[]string             `json:"categories,omitempty"`
		ChecklistItems           *[]ChecklistItem      `json:"checklistItems,omitempty"`
		CompletedDateTime        *DateTimeTimeZone     `json:"completedDateTime,omitempty"`
		CreatedDateTime          *string               `json:"createdDateTime,omitempty"`
		DueDateTime              *DateTimeTimeZone     `json:"dueDateTime,omitempty"`
		HasAttachments           nullable.Type[bool]   `json:"hasAttachments,omitempty"`
		Importance               *Importance           `json:"importance,omitempty"`
		IsReminderOn             *bool                 `json:"isReminderOn,omitempty"`
		LastModifiedDateTime     *string               `json:"lastModifiedDateTime,omitempty"`
		LinkedResources          *[]LinkedResource     `json:"linkedResources,omitempty"`
		Recurrence               *PatternedRecurrence  `json:"recurrence,omitempty"`
		ReminderDateTime         *DateTimeTimeZone     `json:"reminderDateTime,omitempty"`
		StartDateTime            *DateTimeTimeZone     `json:"startDateTime,omitempty"`
		Status                   *TaskStatus           `json:"status,omitempty"`
		Title                    nullable.Type[string] `json:"title,omitempty"`
		Id                       *string               `json:"id,omitempty"`
		ODataId                  *string               `json:"@odata.id,omitempty"`
		ODataType                *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AttachmentSessions = decoded.AttachmentSessions
	s.Body = decoded.Body
	s.BodyLastModifiedDateTime = decoded.BodyLastModifiedDateTime
	s.Categories = decoded.Categories
	s.ChecklistItems = decoded.ChecklistItems
	s.CompletedDateTime = decoded.CompletedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DueDateTime = decoded.DueDateTime
	s.HasAttachments = decoded.HasAttachments
	s.Importance = decoded.Importance
	s.IsReminderOn = decoded.IsReminderOn
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LinkedResources = decoded.LinkedResources
	s.Recurrence = decoded.Recurrence
	s.ReminderDateTime = decoded.ReminderDateTime
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status
	s.Title = decoded.Title
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling TodoTask into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attachments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attachments into list []json.RawMessage: %+v", err)
		}

		output := make([]AttachmentBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAttachmentBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attachments' for 'TodoTask': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attachments = &output
	}

	if v, ok := temp["extensions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Extensions into list []json.RawMessage: %+v", err)
		}

		output := make([]Extension, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExtensionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Extensions' for 'TodoTask': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Extensions = &output
	}

	return nil
}
