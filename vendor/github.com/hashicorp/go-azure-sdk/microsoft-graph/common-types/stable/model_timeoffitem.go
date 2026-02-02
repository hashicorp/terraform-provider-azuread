package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScheduleEntity = TimeOffItem{}

type TimeOffItem struct {
	// ID of the timeOffReason for this timeOffItem. Required.
	TimeOffReasonId nullable.Type[string] `json:"timeOffReasonId,omitempty"`

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

func (s TimeOffItem) ScheduleEntity() BaseScheduleEntityImpl {
	return BaseScheduleEntityImpl{
		EndDateTime:   s.EndDateTime,
		ODataId:       s.ODataId,
		ODataType:     s.ODataType,
		StartDateTime: s.StartDateTime,
		Theme:         s.Theme,
	}
}

var _ json.Marshaler = TimeOffItem{}

func (s TimeOffItem) MarshalJSON() ([]byte, error) {
	type wrapper TimeOffItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TimeOffItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TimeOffItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.timeOffItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TimeOffItem: %+v", err)
	}

	return encoded, nil
}
