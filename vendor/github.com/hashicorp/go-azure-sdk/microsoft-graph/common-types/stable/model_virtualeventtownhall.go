package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VirtualEvent = VirtualEventTownhall{}

type VirtualEventTownhall struct {
	Audience         *MeetingAudience              `json:"audience,omitempty"`
	CoOrganizers     *[]CommunicationsUserIdentity `json:"coOrganizers,omitempty"`
	InvitedAttendees *[]Identity                   `json:"invitedAttendees,omitempty"`
	IsInviteOnly     nullable.Type[bool]           `json:"isInviteOnly,omitempty"`

	// Fields inherited from VirtualEvent

	// Identity information for the creator of the virtual event. Inherited from virtualEvent.
	CreatedBy *CommunicationsIdentitySet `json:"createdBy,omitempty"`

	// Description of the virtual event.
	Description *ItemBody `json:"description,omitempty"`

	// Display name of the virtual event.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// End time of the virtual event. The timeZone property can be set to any of the time zones currently supported by
	// Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
	EndDateTime *DateTimeTimeZone `json:"endDateTime,omitempty"`

	Presenters *[]VirtualEventPresenter `json:"presenters,omitempty"`

	// Sessions for the virtual event.
	Sessions *[]VirtualEventSession `json:"sessions,omitempty"`

	// Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by
	// Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`

	// Status of the virtual event. The possible values are: draft, published, canceled, unknownFutureValue.
	Status *VirtualEventStatus `json:"status,omitempty"`

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

func (s VirtualEventTownhall) VirtualEvent() BaseVirtualEventImpl {
	return BaseVirtualEventImpl{
		CreatedBy:     s.CreatedBy,
		Description:   s.Description,
		DisplayName:   s.DisplayName,
		EndDateTime:   s.EndDateTime,
		Presenters:    s.Presenters,
		Sessions:      s.Sessions,
		StartDateTime: s.StartDateTime,
		Status:        s.Status,
		Id:            s.Id,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
	}
}

func (s VirtualEventTownhall) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventTownhall{}

func (s VirtualEventTownhall) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventTownhall
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventTownhall: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventTownhall: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventTownhall"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventTownhall: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VirtualEventTownhall{}

func (s *VirtualEventTownhall) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Audience      *MeetingAudience              `json:"audience,omitempty"`
		CoOrganizers  *[]CommunicationsUserIdentity `json:"coOrganizers,omitempty"`
		IsInviteOnly  nullable.Type[bool]           `json:"isInviteOnly,omitempty"`
		CreatedBy     *CommunicationsIdentitySet    `json:"createdBy,omitempty"`
		Description   *ItemBody                     `json:"description,omitempty"`
		DisplayName   nullable.Type[string]         `json:"displayName,omitempty"`
		EndDateTime   *DateTimeTimeZone             `json:"endDateTime,omitempty"`
		Presenters    *[]VirtualEventPresenter      `json:"presenters,omitempty"`
		Sessions      *[]VirtualEventSession        `json:"sessions,omitempty"`
		StartDateTime *DateTimeTimeZone             `json:"startDateTime,omitempty"`
		Status        *VirtualEventStatus           `json:"status,omitempty"`
		Id            *string                       `json:"id,omitempty"`
		ODataId       *string                       `json:"@odata.id,omitempty"`
		ODataType     *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Audience = decoded.Audience
	s.CoOrganizers = decoded.CoOrganizers
	s.IsInviteOnly = decoded.IsInviteOnly
	s.CreatedBy = decoded.CreatedBy
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EndDateTime = decoded.EndDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Presenters = decoded.Presenters
	s.Sessions = decoded.Sessions
	s.StartDateTime = decoded.StartDateTime
	s.Status = decoded.Status

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualEventTownhall into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["invitedAttendees"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling InvitedAttendees into list []json.RawMessage: %+v", err)
		}

		output := make([]Identity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIdentityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'InvitedAttendees' for 'VirtualEventTownhall': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.InvitedAttendees = &output
	}

	return nil
}
