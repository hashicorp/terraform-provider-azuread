package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ VirtualEvent = VirtualEventWebinar{}

type VirtualEventWebinar struct {
	// To whom the webinar is visible.
	Audience *MeetingAudience `json:"audience,omitempty"`

	// Identity information of coorganizers of the webinar.
	CoOrganizers *[]CommunicationsUserIdentity `json:"coOrganizers,omitempty"`

	RegistrationConfiguration *VirtualEventWebinarRegistrationConfiguration `json:"registrationConfiguration,omitempty"`

	// Registration records of the webinar.
	Registrations *[]VirtualEventRegistration `json:"registrations,omitempty"`

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

func (s VirtualEventWebinar) VirtualEvent() BaseVirtualEventImpl {
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

func (s VirtualEventWebinar) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventWebinar{}

func (s VirtualEventWebinar) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventWebinar
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventWebinar: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventWebinar: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventWebinar"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventWebinar: %+v", err)
	}

	return encoded, nil
}
