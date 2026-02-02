package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CalendarGroup{}

type CalendarGroup struct {
	// The calendars in the calendar group. Navigation property. Read-only. Nullable.
	Calendars *[]Calendar `json:"calendars,omitempty"`

	// Identifies the version of the calendar group. Every time the calendar group is changed, ChangeKey changes as well.
	// This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey nullable.Type[string] `json:"changeKey,omitempty"`

	// The class identifier. Read-only.
	ClassId nullable.Type[string] `json:"classId,omitempty"`

	// The group name.
	Name nullable.Type[string] `json:"name,omitempty"`

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

func (s CalendarGroup) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CalendarGroup{}

func (s CalendarGroup) MarshalJSON() ([]byte, error) {
	type wrapper CalendarGroup
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CalendarGroup: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CalendarGroup: %+v", err)
	}

	delete(decoded, "calendars")
	delete(decoded, "changeKey")
	delete(decoded, "classId")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.calendarGroup"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CalendarGroup: %+v", err)
	}

	return encoded, nil
}
