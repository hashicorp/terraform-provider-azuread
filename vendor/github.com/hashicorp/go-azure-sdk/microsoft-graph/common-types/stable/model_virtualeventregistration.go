package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VirtualEventRegistration{}

type VirtualEventRegistration struct {
	// Date and time when the registrant cancels their registration for the virtual event. Only appears when applicable. The
	// Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CancelationDateTime nullable.Type[string] `json:"cancelationDateTime,omitempty"`

	// Email address of the registrant.
	Email nullable.Type[string] `json:"email,omitempty"`

	// First name of the registrant.
	FirstName nullable.Type[string] `json:"firstName,omitempty"`

	// Last name of the registrant.
	LastName nullable.Type[string] `json:"lastName,omitempty"`

	PreferredLanguage nullable.Type[string] `json:"preferredLanguage,omitempty"`
	PreferredTimezone nullable.Type[string] `json:"preferredTimezone,omitempty"`

	// Date and time when the registrant registers for the virtual event. The Timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z.
	RegistrationDateTime nullable.Type[string] `json:"registrationDateTime,omitempty"`

	// The registrant's answer to the registration questions.
	RegistrationQuestionAnswers *[]VirtualEventRegistrationQuestionAnswer `json:"registrationQuestionAnswers,omitempty"`

	Sessions *[]VirtualEventSession `json:"sessions,omitempty"`

	// Registration status of the registrant. Read-only.
	Status *VirtualEventAttendeeRegistrationStatus `json:"status,omitempty"`

	// The registrant's ID in Microsoft Entra ID. Only appears when the registrant is registered in Microsoft Entra ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s VirtualEventRegistration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventRegistration{}

func (s VirtualEventRegistration) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventRegistration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventRegistration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventRegistration: %+v", err)
	}

	delete(decoded, "status")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventRegistration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventRegistration: %+v", err)
	}

	return encoded, nil
}
