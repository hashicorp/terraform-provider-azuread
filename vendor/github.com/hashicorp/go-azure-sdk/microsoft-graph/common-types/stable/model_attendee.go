package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AttendeeBase = Attendee{}

type Attendee struct {
	// An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't
	// proposed another time, then this property isn't included in a response of a GET event.
	ProposedNewTime *TimeSlot `json:"proposedNewTime,omitempty"`

	// The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
	Status *ResponseStatus `json:"status,omitempty"`

	// Fields inherited from AttendeeBase

	// The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person,
	// findMeetingTimes always considers the person is of the Required type.
	Type *AttendeeType `json:"type,omitempty"`

	// Fields inherited from Recipient

	// The recipient's email address.
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`

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
