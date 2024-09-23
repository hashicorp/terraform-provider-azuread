package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingRegistrationBase = MeetingRegistration{}

type MeetingRegistration struct {
	// Custom registration questions.
	CustomQuestions *[]MeetingRegistrationQuestion `json:"customQuestions,omitempty"`

	// The description of the meeting.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The meeting end time in UTC.
	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The number of times the registration page has been visited. Read-only.
	RegistrationPageViewCount nullable.Type[int64] `json:"registrationPageViewCount,omitempty"`

	// The URL of the registration page. Read-only.
	RegistrationPageWebUrl nullable.Type[string] `json:"registrationPageWebUrl,omitempty"`

	// The meeting speaker's information.
	Speakers *[]MeetingSpeaker `json:"speakers,omitempty"`

	// The meeting start time in UTC.
	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`

	// The subject of the meeting.
	Subject nullable.Type[string] `json:"subject,omitempty"`

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

func (s MeetingRegistration) MeetingRegistrationBase() BaseMeetingRegistrationBaseImpl {
	return BaseMeetingRegistrationBaseImpl{
		AllowedRegistrant: s.AllowedRegistrant,
		Registrants:       s.Registrants,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s MeetingRegistration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MeetingRegistration{}

func (s MeetingRegistration) MarshalJSON() ([]byte, error) {
	type wrapper MeetingRegistration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MeetingRegistration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingRegistration: %+v", err)
	}

	delete(decoded, "registrationPageViewCount")
	delete(decoded, "registrationPageWebUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.meetingRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MeetingRegistration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MeetingRegistration{}

func (s *MeetingRegistration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CustomQuestions           *[]MeetingRegistrationQuestion `json:"customQuestions,omitempty"`
		Description               nullable.Type[string]          `json:"description,omitempty"`
		EndDateTime               nullable.Type[string]          `json:"endDateTime,omitempty"`
		RegistrationPageViewCount nullable.Type[int64]           `json:"registrationPageViewCount,omitempty"`
		RegistrationPageWebUrl    nullable.Type[string]          `json:"registrationPageWebUrl,omitempty"`
		Speakers                  *[]MeetingSpeaker              `json:"speakers,omitempty"`
		StartDateTime             nullable.Type[string]          `json:"startDateTime,omitempty"`
		Subject                   nullable.Type[string]          `json:"subject,omitempty"`
		AllowedRegistrant         *MeetingAudience               `json:"allowedRegistrant,omitempty"`
		Id                        *string                        `json:"id,omitempty"`
		ODataId                   *string                        `json:"@odata.id,omitempty"`
		ODataType                 *string                        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CustomQuestions = decoded.CustomQuestions
	s.Description = decoded.Description
	s.EndDateTime = decoded.EndDateTime
	s.RegistrationPageViewCount = decoded.RegistrationPageViewCount
	s.RegistrationPageWebUrl = decoded.RegistrationPageWebUrl
	s.Speakers = decoded.Speakers
	s.StartDateTime = decoded.StartDateTime
	s.Subject = decoded.Subject
	s.AllowedRegistrant = decoded.AllowedRegistrant
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MeetingRegistration into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Registrants' for 'MeetingRegistration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Registrants = &output
	}

	return nil
}
