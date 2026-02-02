package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MeetingRegistrantBase = ExternalMeetingRegistrant{}

type ExternalMeetingRegistrant struct {
	// The tenant ID of this registrant if in Microsoft Entra ID.
	TenantId nullable.Type[string] `json:"tenantId,omitempty"`

	// The user ID of this registrant if in Microsoft Entra ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

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

func (s ExternalMeetingRegistrant) MeetingRegistrantBase() BaseMeetingRegistrantBaseImpl {
	return BaseMeetingRegistrantBaseImpl{
		JoinWebUrl: s.JoinWebUrl,
		Id:         s.Id,
		ODataId:    s.ODataId,
		ODataType:  s.ODataType,
	}
}

func (s ExternalMeetingRegistrant) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalMeetingRegistrant{}

func (s ExternalMeetingRegistrant) MarshalJSON() ([]byte, error) {
	type wrapper ExternalMeetingRegistrant
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalMeetingRegistrant: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalMeetingRegistrant: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalMeetingRegistrant"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalMeetingRegistrant: %+v", err)
	}

	return encoded, nil
}
