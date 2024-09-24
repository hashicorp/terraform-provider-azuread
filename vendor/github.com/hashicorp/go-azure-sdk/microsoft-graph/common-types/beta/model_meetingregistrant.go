package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingRegistrantBase = MeetingRegistrant{}

type MeetingRegistrant struct {
	// The registrant's answer to custom questions.
	CustomQuestionAnswers *[]CustomQuestionAnswer `json:"customQuestionAnswers,omitempty"`

	// The email address of the registrant.
	Email nullable.Type[string] `json:"email,omitempty"`

	// The first name of the registrant.
	FirstName nullable.Type[string] `json:"firstName,omitempty"`

	// The family name of the registrant.
	LastName nullable.Type[string] `json:"lastName,omitempty"`

	// Time in UTC when the registrant registers for the meeting. Read-only.
	RegistrationDateTime nullable.Type[string] `json:"registrationDateTime,omitempty"`

	// The registration status of the registrant. Read-only.
	Status *MeetingRegistrantStatus `json:"status,omitempty"`

	// Fields inherited from MeetingRegistrantBase

	// A unique web URL for the registrant to join the meeting. Read-only.
	JoinWebUrl nullable.Type[string] `json:"joinWebUrl,omitempty"`

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

func (s MeetingRegistrant) MeetingRegistrantBase() BaseMeetingRegistrantBaseImpl {
	return BaseMeetingRegistrantBaseImpl{
		JoinWebUrl: s.JoinWebUrl,
		Id:         s.Id,
		ODataId:    s.ODataId,
		ODataType:  s.ODataType,
	}
}

func (s MeetingRegistrant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MeetingRegistrant{}

func (s MeetingRegistrant) MarshalJSON() ([]byte, error) {
	type wrapper MeetingRegistrant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MeetingRegistrant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MeetingRegistrant: %+v", err)
	}

	delete(decoded, "registrationDateTime")
	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.meetingRegistrant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MeetingRegistrant: %+v", err)
	}

	return encoded, nil
}
