package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MeetingAttendanceReport{}

type MeetingAttendanceReport struct {
	// List of attendance records of an attendance report. Read-only.
	AttendanceRecords *[]AttendanceRecord `json:"attendanceRecords,omitempty"`

	// The external information of a virtual event. Returned only for event organizers or coorganizers. Read-only.
	ExternalEventInformation *[]VirtualEventExternalInformation `json:"externalEventInformation,omitempty"`

	// UTC time when the meeting ended. Read-only.
	MeetingEndDateTime nullable.Type[string] `json:"meetingEndDateTime,omitempty"`

	// UTC time when the meeting started. Read-only.
	MeetingStartDateTime nullable.Type[string] `json:"meetingStartDateTime,omitempty"`

	// Total number of participants. Read-only.
	TotalParticipantCount nullable.Type[int64] `json:"totalParticipantCount,omitempty"`

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

func (s MeetingAttendanceReport) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MeetingAttendanceReport{}

func (s MeetingAttendanceReport) MarshalJSON() ([]byte, error) {
	type wrapper MeetingAttendanceReport
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MeetingAttendanceReport: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingAttendanceReport: %+v", err)
	}

	delete(decoded, "attendanceRecords")
	delete(decoded, "externalEventInformation")
	delete(decoded, "meetingEndDateTime")
	delete(decoded, "meetingStartDateTime")
	delete(decoded, "totalParticipantCount")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.meetingAttendanceReport"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MeetingAttendanceReport: %+v", err)
	}

	return encoded, nil
}
