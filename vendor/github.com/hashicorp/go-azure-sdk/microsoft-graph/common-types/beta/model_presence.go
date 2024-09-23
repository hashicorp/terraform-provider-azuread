package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Presence{}

type Presence struct {
	// The supplemental information to a user's availability. Possible values are Available, Away, BeRightBack, Busy,
	// DoNotDisturb, InACall, InAConferenceCall, Inactive, InAMeeting, Offline, OffWork, OutOfOffice, PresenceUnknown,
	// Presenting, UrgentInterruptionsOnly.
	Activity nullable.Type[string] `json:"activity,omitempty"`

	// The base presence information for a user. Possible values are Available, AvailableIdle, Away, BeRightBack, Busy,
	// BusyIdle, DoNotDisturb, Offline, PresenceUnknown.
	Availability nullable.Type[string] `json:"availability,omitempty"`

	// The out of office settings for a user.
	OutOfOfficeSettings *OutOfOfficeSettings `json:"outOfOfficeSettings,omitempty"`

	// The presence status message of a user.
	StatusMessage *PresenceStatusMessage `json:"statusMessage,omitempty"`

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

func (s Presence) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Presence{}

func (s Presence) MarshalJSON() ([]byte, error) {
	type wrapper Presence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Presence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Presence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.presence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Presence: %+v", err)
	}

	return encoded, nil
}
