package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AttendeeBase = Attendee{}

type Attendee struct {
	// An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't
	// proposed another time, then this property isn't included in a response of a GET event.
	ProposedNewTime *TimeSlot `json:"proposedNewTime,omitempty"`

	// The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
	Status *ResponseStatus `json:"status,omitempty"`

	// Fields inherited from AttendeeBase

	// The type of attendee. Possible values are: required, optional, resource. Currently if the attendee is a person,
	// findMeetingTimes always considers the person is of the Required type.
	Type *AttendeeType `json:"type,omitempty"`

	// Fields inherited from Recipient

	// The recipient's email address.
	EmailAddress EmailAddress `json:"emailAddress"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Attendee) AttendeeBase() BaseAttendeeBaseImpl {
	return BaseAttendeeBaseImpl{
		Type:         s.Type,
		EmailAddress: s.EmailAddress,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s Attendee) Recipient() BaseRecipientImpl {
	return BaseRecipientImpl{
		EmailAddress: s.EmailAddress,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

var _ json.Marshaler = Attendee{}

func (s Attendee) MarshalJSON() ([]byte, error) {
	type wrapper Attendee
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Attendee: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Attendee: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attendee"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Attendee: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Attendee{}

func (s *Attendee) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ProposedNewTime *TimeSlot       `json:"proposedNewTime,omitempty"`
		Status          *ResponseStatus `json:"status,omitempty"`
		Type            *AttendeeType   `json:"type,omitempty"`
		ODataId         *string         `json:"@odata.id,omitempty"`
		ODataType       *string         `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ProposedNewTime = decoded.ProposedNewTime
	s.Status = decoded.Status
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Attendee into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["emailAddress"]; ok {
		impl, err := UnmarshalEmailAddressImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EmailAddress' for 'Attendee': %+v", err)
		}
		s.EmailAddress = impl
	}

	return nil
}
