package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipants struct {
	// Information about the meeting attendees.
	Attendees *[]MeetingParticipantInfo `json:"attendees,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information about the meeting organizer.
	Organizer MeetingParticipantInfo `json:"organizer"`
}

var _ json.Unmarshaler = &MeetingParticipants{}

func (s *MeetingParticipants) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId   *string `json:"@odata.id,omitempty"`
		ODataType *string `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MeetingParticipants into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["attendees"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Attendees into list []json.RawMessage: %+v", err)
		}

		output := make([]MeetingParticipantInfo, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMeetingParticipantInfoImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Attendees' for 'MeetingParticipants': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Attendees = &output
	}

	if v, ok := temp["organizer"]; ok {
		impl, err := UnmarshalMeetingParticipantInfoImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Organizer' for 'MeetingParticipants': %+v", err)
		}
		s.Organizer = impl
	}

	return nil
}
