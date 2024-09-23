package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ShiftItem = OpenShiftItem{}

type OpenShiftItem struct {
	// Count of the number of slots for the given open shift.
	OpenSlotCount *int64 `json:"openSlotCount,omitempty"`

	// Fields inherited from ShiftItem

	// An incremental part of a shift which can cover details of when and where an employee is during their shift. For
	// example, an assignment or a scheduled break or lunch. Required.
	Activities []ShiftActivity `json:"activities"`

	// The shift label of the shiftItem.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The shift notes for the shiftItem.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// Fields inherited from ScheduleEntity

	EndDateTime nullable.Type[string] `json:"endDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	StartDateTime nullable.Type[string] `json:"startDateTime,omitempty"`
	Theme         *ScheduleEntityTheme  `json:"theme,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OpenShiftItem) ShiftItem() BaseShiftItemImpl {
	return BaseShiftItemImpl{
		Activities:    s.Activities,
		DisplayName:   s.DisplayName,
		Notes:         s.Notes,
		EndDateTime:   s.EndDateTime,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		StartDateTime: s.StartDateTime,
		Theme:         s.Theme,
	}
}

func (s OpenShiftItem) ScheduleEntity() BaseScheduleEntityImpl {
	return BaseScheduleEntityImpl{
		EndDateTime:   s.EndDateTime,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		StartDateTime: s.StartDateTime,
		Theme:         s.Theme,
	}
}

var _ json.Marshaler = OpenShiftItem{}

func (s OpenShiftItem) MarshalJSON() ([]byte, error) {
	type wrapper OpenShiftItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OpenShiftItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OpenShiftItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.openShiftItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OpenShiftItem: %+v", err)
	}

	return encoded, nil
}
