package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InactiveUsersMetricBase interface {
	Entity
	InactiveUsersMetricBase() BaseInactiveUsersMetricBaseImpl
}

var _ InactiveUsersMetricBase = BaseInactiveUsersMetricBaseImpl{}

type BaseInactiveUsersMetricBaseImpl struct {
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

func (s BaseInactiveUsersMetricBaseImpl) InactiveUsersMetricBase() BaseInactiveUsersMetricBaseImpl {
	return s
}

func (s BaseInactiveUsersMetricBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ InactiveUsersMetricBase = RawInactiveUsersMetricBaseImpl{}

// RawInactiveUsersMetricBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInactiveUsersMetricBaseImpl struct {
	inactiveUsersMetricBase BaseInactiveUsersMetricBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawInactiveUsersMetricBaseImpl) InactiveUsersMetricBase() BaseInactiveUsersMetricBaseImpl {
	return s.inactiveUsersMetricBase
}

func (s RawInactiveUsersMetricBaseImpl) Entity() BaseEntityImpl {
	return s.inactiveUsersMetricBase.Entity()
}

var _ json.Marshaler = BaseInactiveUsersMetricBaseImpl{}

func (s BaseInactiveUsersMetricBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseInactiveUsersMetricBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseInactiveUsersMetricBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseInactiveUsersMetricBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.inactiveUsersMetricBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseInactiveUsersMetricBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalInactiveUsersMetricBaseImplementation(input []byte) (InactiveUsersMetricBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InactiveUsersMetricBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.dailyInactiveUsersMetric") {
		var out DailyInactiveUsersMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DailyInactiveUsersMetric: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.monthlyInactiveUsersMetric") {
		var out MonthlyInactiveUsersMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MonthlyInactiveUsersMetric: %+v", err)
		}
		return out, nil
	}

	var parent BaseInactiveUsersMetricBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseInactiveUsersMetricBaseImpl: %+v", err)
	}

	return RawInactiveUsersMetricBaseImpl{
		inactiveUsersMetricBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
