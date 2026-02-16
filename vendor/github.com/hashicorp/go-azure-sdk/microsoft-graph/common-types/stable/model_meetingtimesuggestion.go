package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestion struct {
	// An array that shows the availability status of each attendee for this meeting suggestion.
	AttendeeAvailability *[]AttendeeAvailability `json:"attendeeAvailability,omitempty"`

	// An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
	Locations *[]Location `json:"locations,omitempty"`

	// A time period suggested for the meeting.
	MeetingTimeSlot *TimeSlot `json:"meetingTimeSlot,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if
	// there are suggestions with the same confidence.
	Order nullable.Type[int64] `json:"order,omitempty"`

	// Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy,
	// oof, workingElsewhere, unknown.
	OrganizerAvailability *FreeBusyStatus `json:"organizerAvailability,omitempty"`

	// Reason for suggesting the meeting time.
	SuggestionReason nullable.Type[string] `json:"suggestionReason,omitempty"`
}

var _ json.Unmarshaler = &MeetingTimeSuggestion{}

func (s *MeetingTimeSuggestion) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AttendeeAvailability  *[]AttendeeAvailability `json:"attendeeAvailability,omitempty"`
		MeetingTimeSlot       *TimeSlot               `json:"meetingTimeSlot,omitempty"`
		ODataId               *string                 `json:"@odata.id,omitempty"`
		ODataType             *string                 `json:"@odata.type,omitempty"`
		Order                 nullable.Type[int64]    `json:"order,omitempty"`
		OrganizerAvailability *FreeBusyStatus         `json:"organizerAvailability,omitempty"`
		SuggestionReason      nullable.Type[string]   `json:"suggestionReason,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AttendeeAvailability = decoded.AttendeeAvailability
	s.MeetingTimeSlot = decoded.MeetingTimeSlot
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Order = decoded.Order
	s.OrganizerAvailability = decoded.OrganizerAvailability
	s.SuggestionReason = decoded.SuggestionReason

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MeetingTimeSuggestion into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["locations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Locations into list []json.RawMessage: %+v", err)
		}

		output := make([]Location, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalLocationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Locations' for 'MeetingTimeSuggestion': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Locations = &output
	}

	return nil
}
