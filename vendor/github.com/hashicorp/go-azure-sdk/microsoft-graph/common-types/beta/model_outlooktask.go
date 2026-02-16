package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ OutlookItem = OutlookTask{}

type OutlookTask struct {
	// The name of the person who has been assigned the task in Outlook. Read-only.
	AssignedTo nullable.Type[string] `json:"assignedTo,omitempty"`

	// The collection of fileAttachment, itemAttachment, and referenceAttachment attachments for the task. Read-only.
	// Nullable.
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// The task body that typically contains information about the task. Only the HTML type is supported.
	Body *ItemBody `json:"body,omitempty"`

	// The date in the specified time zone that the task was finished.
	CompletedDateTime *DateTimeTimeZone `json:"completedDateTime,omitempty"`

	// The date in the specified time zone that the task is to be finished.
	DueDateTime *DateTimeTimeZone `json:"dueDateTime,omitempty"`

	// Set to true if the task has attachments.
	HasAttachments nullable.Type[bool] `json:"hasAttachments,omitempty"`

	// The importance of the event. Possible values are: low, normal, high.
	Importance *Importance `json:"importance,omitempty"`

	// Set to true if an alert is set to remind the user of the task.
	IsReminderOn nullable.Type[bool] `json:"isReminderOn,omitempty"`

	// The collection of multi-value extended properties defined for the task. Read-only. Nullable.
	MultiValueExtendedProperties *[]MultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`

	// The name of the person who created the task.
	Owner nullable.Type[string] `json:"owner,omitempty"`

	// The unique identifier for the task's parent folder.
	ParentFolderId nullable.Type[string] `json:"parentFolderId,omitempty"`

	// The recurrence pattern for the task.
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`

	// The date and time for a reminder alert of the task to occur.
	ReminderDateTime *DateTimeTimeZone `json:"reminderDateTime,omitempty"`

	// Indicates the level of privacy for the task. Possible values are: normal, personal, private, confidential.
	Sensitivity *Sensitivity `json:"sensitivity,omitempty"`

	// The collection of single-value extended properties defined for the task. Read-only. Nullable.
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`

	// The date in the specified time zone when the task is to begin.
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`

	// Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed, waitingOnOthers,
	// deferred.
	Status *TaskStatus `json:"status,omitempty"`

	// A brief description or title of the task.
	Subject nullable.Type[string] `json:"subject,omitempty"`

	// Fields inherited from OutlookItem

	// The categories associated with the item.
	Categories *[]string `json:"categories,omitempty"`

	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange
	// to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

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

func (s OutlookTask) OutlookItem() BaseOutlookItemImpl {
	return BaseOutlookItemImpl{
		Categories:           s.Categories,
		ChangeKey:            s.ChangeKey,
		CreatedDateTime:      s.CreatedDateTime,
		LastModifiedDateTime: s.LastModifiedDateTime,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s OutlookTask) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OutlookTask{}

func (s OutlookTask) MarshalJSON() ([]byte, error) {
	type wrapper OutlookTask
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OutlookTask: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OutlookTask: %+v", err)
	}

	delete(decoded, "assignedTo")
	delete(decoded, "attachments")
	delete(decoded, "multiValueExtendedProperties")
	delete(decoded, "singleValueExtendedProperties")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.outlookTask"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OutlookTask: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &OutlookTask{}

func (s *OutlookTask) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssignedTo                    nullable.Type[string]                `json:"assignedTo,omitempty"`
		Body                          *ItemBody                            `json:"body,omitempty"`
		CompletedDateTime             *DateTimeTimeZone                    `json:"completedDateTime,omitempty"`
		DueDateTime                   *DateTimeTimeZone                    `json:"dueDateTime,omitempty"`
		HasAttachments                nullable.Type[bool]                  `json:"hasAttachments,omitempty"`
		Importance                    *Importance                          `json:"importance,omitempty"`
		IsReminderOn                  nullable.Type[bool]                  `json:"isReminderOn,omitempty"`
		MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
		Owner                         nullable.Type[string]                `json:"owner,omitempty"`
		ParentFolderId                nullable.Type[string]                `json:"parentFolderId,omitempty"`
		Recurrence                    *PatternedRecurrence                 `json:"recurrence,omitempty"`
		ReminderDateTime              *DateTimeTimeZone                    `json:"reminderDateTime,omitempty"`
		Sensitivity                   *Sensitivity                         `json:"sensitivity,omitempty"`
		SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
		StartDateTime                 *DateTimeTimeZone                    `json:"startDateTime,omitempty"`
		Status                        *TaskStatus                          `json:"status,omitempty"`
		Subject                       nullable.Type[string]                `json:"subject,omitempty"`
		Categories                    *[]string                            `json:"categories,omitempty"`
		ChangeKey                     nullable.Type[string]                `json:"changeKey,omitempty"`
		CreatedDateTime               nullable.Type[string]                `json:"createdDateTime,omitempty"`
		LastModifiedDateTime          nullable.Type[string]                `json:"lastModifiedDateTime,omitempty"`
		Id                            *string                              `json:"id,omitempty"`
		ODataId                       *string                              `json:"@odata.id,omitempty"`
		ODataType                     *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssignedTo = decoded.AssignedTo
	s.Body = decoded.Body
	s.CompletedDateTime = decoded.CompletedDateTime
	s.DueDateTime = decoded.DueDateTime
	s.HasAttachments = decoded.HasAttachments
	s.Importance = decoded.Importance
	s.IsReminderOn = decoded.IsReminderOn
	s.MultiValueExtendedProperties = decoded.MultiValueExtendedProperties
	s.Owner = decoded.Owner
	s.ParentFolderId = decoded.ParentFolderId
	s.Recurrence = decoded.Recurrence
	s.ReminderDateTime = decoded.ReminderDateTime
	s.Sensitivity = decoded.Sensitivity
	s.SingleValueExtendedProperties = decoded.SingleValueExtendedProperties
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status
	s.Subject = decoded.Subject
	s.Categories = decoded.Categories
	s.ChangeKey = decoded.ChangeKey
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling OutlookTask into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attachments"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attachments into list []json.RawMessage: %+v", err)
		}

		output := make([]Attachment, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAttachmentImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attachments' for 'OutlookTask': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attachments = &output
	}

	return nil
}
