package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EducationAssignmentDefaults{}

type EducationAssignmentDefaults struct {
	// Optional field to control adding assignments to students' and teachers' calendars when the assignment is published.
	// The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Use
	// the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum:
	// studentsOnly. The default value is none.
	AddToCalendarAction *EducationAddToCalendarOptions `json:"addToCalendarAction,omitempty"`

	// Class-level default behavior for handling students who are added after the assignment is published. Possible values
	// are: none, assignIfOpen.
	AddedStudentAction *EducationAddedStudentAction `json:"addedStudentAction,omitempty"`

	// Class-level default value for due time field. Default value is 23:59:00.
	DueTime nullable.Type[string] `json:"dueTime,omitempty"`

	// Default Teams channel to which notifications are sent. Default value is null.
	NotificationChannelUrl nullable.Type[string] `json:"notificationChannelUrl,omitempty"`

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

func (s EducationAssignmentDefaults) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EducationAssignmentDefaults{}

func (s EducationAssignmentDefaults) MarshalJSON() ([]byte, error) {
	type wrapper EducationAssignmentDefaults
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EducationAssignmentDefaults: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EducationAssignmentDefaults: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.educationAssignmentDefaults"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EducationAssignmentDefaults: %+v", err)
	}

	return encoded, nil
}
