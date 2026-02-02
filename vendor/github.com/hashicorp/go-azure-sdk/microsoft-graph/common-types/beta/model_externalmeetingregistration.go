package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingRegistrationBase = ExternalMeetingRegistration{}

type ExternalMeetingRegistration struct {

	// Fields inherited from MeetingRegistrationBase

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

func (s ExternalMeetingRegistration) MeetingRegistrationBase() BaseMeetingRegistrationBaseImpl {
	return BaseMeetingRegistrationBaseImpl{
		AllowedRegistrant: s.AllowedRegistrant,
		Registrants:       s.Registrants,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s ExternalMeetingRegistration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalMeetingRegistration{}

func (s ExternalMeetingRegistration) MarshalJSON() ([]byte, error) {
	type wrapper ExternalMeetingRegistration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalMeetingRegistration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalMeetingRegistration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalMeetingRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalMeetingRegistration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExternalMeetingRegistration{}

func (s *ExternalMeetingRegistration) UnmarshalJSON(bytes []byte) error {
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
		return fmt.Errorf("unmarshaling ExternalMeetingRegistration into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Registrants' for 'ExternalMeetingRegistration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Registrants = &output
	}

	return nil
}
