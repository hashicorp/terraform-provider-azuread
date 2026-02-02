package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ InactiveUsersMetricBase = DailyInactiveUsersMetric{}

type DailyInactiveUsersMetric struct {
	Inactive1DayCount nullable.Type[int64] `json:"inactive1DayCount,omitempty"`

	// Fields inherited from InactiveUsersMetricBase

	AppId              nullable.Type[string] `json:"appId,omitempty"`
	FactDate           *string               `json:"factDate,omitempty"`
	Inactive30DayCount nullable.Type[int64]  `json:"inactive30DayCount,omitempty"`
	Inactive60DayCount nullable.Type[int64]  `json:"inactive60DayCount,omitempty"`
	Inactive90DayCount nullable.Type[int64]  `json:"inactive90DayCount,omitempty"`

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

func (s DailyInactiveUsersMetric) InactiveUsersMetricBase() BaseInactiveUsersMetricBaseImpl {
	return BaseInactiveUsersMetricBaseImpl{
		AppId:              s.AppId,
		FactDate:           s.FactDate,
		Inactive30DayCount: s.Inactive30DayCount,
		Inactive60DayCount: s.Inactive60DayCount,
		Inactive90DayCount: s.Inactive90DayCount,
		Id:                 s.Id,
		ODataId:            s.ODataId,
		ODataType:          s.ODataType,
	}
}

func (s DailyInactiveUsersMetric) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DailyInactiveUsersMetric{}

func (s DailyInactiveUsersMetric) MarshalJSON() ([]byte, error) {
	type wrapper DailyInactiveUsersMetric
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DailyInactiveUsersMetric: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DailyInactiveUsersMetric: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dailyInactiveUsersMetric"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DailyInactiveUsersMetric: %+v", err)
	}

	return encoded, nil
}
