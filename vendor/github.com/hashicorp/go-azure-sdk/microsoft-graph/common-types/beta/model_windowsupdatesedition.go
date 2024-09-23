package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = WindowsUpdatesEdition{}

type WindowsUpdatesEdition struct {
	// The device family targeted by the edition.
	DeviceFamily *string `json:"deviceFamily,omitempty"`

	// The date and time when the edition reached the end of service. The timestamp type represents date and time
	// information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is
	// 2014-01-01T00:00:00Z. Read-only.
	EndOfServiceDateTime *string `json:"endOfServiceDateTime,omitempty"`

	// The date and time when the edition became available to the general customers for the first time. The timestamp type
	// represents date and time information using ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1,
	// 2014 is 2014-01-01T00:00:00Z. Read-only.
	GeneralAvailabilityDateTime *string `json:"generalAvailabilityDateTime,omitempty"`

	// Indicates whether the edition is in service or out of service.
	IsInService *bool `json:"isInService,omitempty"`

	// The name of the edition. Read-only.
	Name *string `json:"name,omitempty"`

	// The public name of the edition. Read-only.
	ReleasedName *string `json:"releasedName,omitempty"`

	ServicingPeriods *[]WindowsUpdatesServicingPeriod `json:"servicingPeriods,omitempty"`

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

func (s WindowsUpdatesEdition) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsUpdatesEdition{}

func (s WindowsUpdatesEdition) MarshalJSON() ([]byte, error) {
	type wrapper WindowsUpdatesEdition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsUpdatesEdition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesEdition: %+v", err)
	}

	delete(decoded, "endOfServiceDateTime")
	delete(decoded, "generalAvailabilityDateTime")
	delete(decoded, "name")
	delete(decoded, "releasedName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.edition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsUpdatesEdition: %+v", err)
	}

	return encoded, nil
}
