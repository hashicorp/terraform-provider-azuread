package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEvent interface {
	Entity
	VirtualEvent() BaseVirtualEventImpl
}

var _ VirtualEvent = BaseVirtualEventImpl{}

type BaseVirtualEventImpl struct {
	// The identity information for the creator of the virtual event. Inherited from virtualEvent.
	CreatedBy *CommunicationsIdentitySet `json:"createdBy,omitempty"`

	// A description of the virtual event.
	Description *ItemBody `json:"description,omitempty"`

	// The display name of the virtual event.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The end time of the virtual event. The timeZone property can be set to any of the time zones currently supported by
	// Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
	EndDateTime *DateTimeTimeZone `json:"endDateTime,omitempty"`

	// The external information of a virtual event. Returned only for event organizers or coorganizers; otherwise, null.
	ExternalEventInformation *[]VirtualEventExternalInformation `json:"externalEventInformation,omitempty"`

	// The virtual event presenters.
	Presenters *[]VirtualEventPresenter `json:"presenters,omitempty"`

	// The sessions for the virtual event.
	Sessions *[]VirtualEventSession `json:"sessions,omitempty"`

	// The virtual event settings.
	Settings *VirtualEventSettings `json:"settings,omitempty"`

	// Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by
	// Windows. For details on how to get all available time zones using PowerShell, see Get-TimeZone.
	StartDateTime *DateTimeTimeZone `json:"startDateTime,omitempty"`

	// The status of the virtual event. The possible values are: draft, published, canceled, and unknownFutureValue.
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

func (s BaseVirtualEventImpl) VirtualEvent() BaseVirtualEventImpl {
	return s
}

func (s BaseVirtualEventImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ VirtualEvent = RawVirtualEventImpl{}

// RawVirtualEventImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawVirtualEventImpl struct {
	virtualEvent BaseVirtualEventImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawVirtualEventImpl) VirtualEvent() BaseVirtualEventImpl {
	return s.virtualEvent
}

func (s RawVirtualEventImpl) Entity() BaseEntityImpl {
	return s.virtualEvent.Entity()
}

var _ json.Marshaler = BaseVirtualEventImpl{}

func (s BaseVirtualEventImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseVirtualEventImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseVirtualEventImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseVirtualEventImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEvent"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseVirtualEventImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalVirtualEventImplementation(input []byte) (VirtualEvent, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEvent into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventTownhall") {
		var out VirtualEventTownhall
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventTownhall: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.virtualEventWebinar") {
		var out VirtualEventWebinar
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VirtualEventWebinar: %+v", err)
		}
		return out, nil
	}

	var parent BaseVirtualEventImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseVirtualEventImpl: %+v", err)
	}

	return RawVirtualEventImpl{
		virtualEvent: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
