package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AttendanceRecord{}

type AttendanceRecord struct {
	// List of time periods between joining and leaving a meeting.
	AttendanceIntervals *[]AttendanceInterval `json:"attendanceIntervals,omitempty"`

	// Email address of the user associated with this attendance record.
	EmailAddress nullable.Type[string] `json:"emailAddress,omitempty"`

	// Identity of the user associated with this attendance record.
	Identity Identity `json:"identity"`

	// Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
	Role nullable.Type[string] `json:"role,omitempty"`

	// Total duration of the attendances in seconds.
	TotalAttendanceInSeconds nullable.Type[int64] `json:"totalAttendanceInSeconds,omitempty"`

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

func (s AttendanceRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AttendanceRecord{}

func (s AttendanceRecord) MarshalJSON() ([]byte, error) {
	type wrapper AttendanceRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AttendanceRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AttendanceRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attendanceRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AttendanceRecord: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AttendanceRecord{}

func (s *AttendanceRecord) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AttendanceIntervals      *[]AttendanceInterval `json:"attendanceIntervals,omitempty"`
		EmailAddress             nullable.Type[string] `json:"emailAddress,omitempty"`
		Role                     nullable.Type[string] `json:"role,omitempty"`
		TotalAttendanceInSeconds nullable.Type[int64]  `json:"totalAttendanceInSeconds,omitempty"`
		Id                       *string               `json:"id,omitempty"`
		ODataId                  *string               `json:"@odata.id,omitempty"`
		ODataType                *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AttendanceIntervals = decoded.AttendanceIntervals
	s.EmailAddress = decoded.EmailAddress
	s.Role = decoded.Role
	s.TotalAttendanceInSeconds = decoded.TotalAttendanceInSeconds
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AttendanceRecord into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'AttendanceRecord': %+v", err)
		}
		s.Identity = impl
	}

	return nil
}
