package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationBase interface {
	Entity
	MeetingRegistrationBase() BaseMeetingRegistrationBaseImpl
}

var _ MeetingRegistrationBase = BaseMeetingRegistrationBaseImpl{}

type BaseMeetingRegistrationBaseImpl struct {
	// Specifies who can register for the meeting.
	AllowedRegistrant *MeetingAudience `json:"allowedRegistrant,omitempty"`

	// Registrants of the online meeting.
	Registrants *[]MeetingRegistrantBase `json:"registrants,omitempty"`

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

func (s BaseMeetingRegistrationBaseImpl) MeetingRegistrationBase() BaseMeetingRegistrationBaseImpl {
	return s
}

func (s BaseMeetingRegistrationBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ MeetingRegistrationBase = RawMeetingRegistrationBaseImpl{}

// RawMeetingRegistrationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMeetingRegistrationBaseImpl struct {
	meetingRegistrationBase BaseMeetingRegistrationBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawMeetingRegistrationBaseImpl) MeetingRegistrationBase() BaseMeetingRegistrationBaseImpl {
	return s.meetingRegistrationBase
}

func (s RawMeetingRegistrationBaseImpl) Entity() BaseEntityImpl {
	return s.meetingRegistrationBase.Entity()
}

var _ json.Marshaler = BaseMeetingRegistrationBaseImpl{}

func (s BaseMeetingRegistrationBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseMeetingRegistrationBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseMeetingRegistrationBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseMeetingRegistrationBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.meetingRegistrationBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseMeetingRegistrationBaseImpl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BaseMeetingRegistrationBaseImpl{}

func (s *BaseMeetingRegistrationBaseImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowedRegistrant *MeetingAudience `json:"allowedRegistrant,omitempty"`
		Id                *string          `json:"id,omitempty"`
		ODataId           *string          `json:"@odata.id,omitempty"`
		ODataType         *string          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowedRegistrant = decoded.AllowedRegistrant
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseMeetingRegistrationBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["registrants"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Registrants into list []json.RawMessage: %+v", err)
		}

		output := make([]MeetingRegistrantBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMeetingRegistrantBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Registrants' for 'BaseMeetingRegistrationBaseImpl': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Registrants = &output
	}

	return nil
}

func UnmarshalMeetingRegistrationBaseImplementation(input []byte) (MeetingRegistrationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingRegistrationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.externalMeetingRegistration") {
		var out ExternalMeetingRegistration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalMeetingRegistration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.meetingRegistration") {
		var out MeetingRegistration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MeetingRegistration: %+v", err)
		}
		return out, nil
	}

	var parent BaseMeetingRegistrationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMeetingRegistrationBaseImpl: %+v", err)
	}

	return RawMeetingRegistrationBaseImpl{
		meetingRegistrationBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
