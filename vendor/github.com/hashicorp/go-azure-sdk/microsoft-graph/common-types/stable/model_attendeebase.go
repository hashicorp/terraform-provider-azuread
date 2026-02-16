package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeBase interface {
	Recipient
	AttendeeBase() BaseAttendeeBaseImpl
}

var _ AttendeeBase = BaseAttendeeBaseImpl{}

type BaseAttendeeBaseImpl struct {
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

func (s BaseAttendeeBaseImpl) AttendeeBase() BaseAttendeeBaseImpl {
	return s
}

func (s BaseAttendeeBaseImpl) Recipient() BaseRecipientImpl {
	return BaseRecipientImpl{
		EmailAddress: s.EmailAddress,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

var _ AttendeeBase = RawAttendeeBaseImpl{}

// RawAttendeeBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAttendeeBaseImpl struct {
	attendeeBase BaseAttendeeBaseImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawAttendeeBaseImpl) AttendeeBase() BaseAttendeeBaseImpl {
	return s.attendeeBase
}

func (s RawAttendeeBaseImpl) Recipient() BaseRecipientImpl {
	return s.attendeeBase.Recipient()
}

var _ json.Marshaler = BaseAttendeeBaseImpl{}

func (s BaseAttendeeBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAttendeeBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAttendeeBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAttendeeBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.attendeeBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAttendeeBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAttendeeBaseImplementation(input []byte) (AttendeeBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AttendeeBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.attendee") {
		var out Attendee
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Attendee: %+v", err)
		}
		return out, nil
	}

	var parent BaseAttendeeBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAttendeeBaseImpl: %+v", err)
	}

	return RawAttendeeBaseImpl{
		attendeeBase: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
