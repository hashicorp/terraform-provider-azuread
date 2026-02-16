package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InactiveUsersByApplicationMetricBase interface {
	Entity
	InactiveUsersByApplicationMetricBase() BaseInactiveUsersByApplicationMetricBaseImpl
}

var _ InactiveUsersByApplicationMetricBase = BaseInactiveUsersByApplicationMetricBaseImpl{}

type BaseInactiveUsersByApplicationMetricBaseImpl struct {
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

func (s BaseInactiveUsersByApplicationMetricBaseImpl) InactiveUsersByApplicationMetricBase() BaseInactiveUsersByApplicationMetricBaseImpl {
	return s
}

func (s BaseInactiveUsersByApplicationMetricBaseImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ InactiveUsersByApplicationMetricBase = RawInactiveUsersByApplicationMetricBaseImpl{}

// RawInactiveUsersByApplicationMetricBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInactiveUsersByApplicationMetricBaseImpl struct {
	inactiveUsersByApplicationMetricBase BaseInactiveUsersByApplicationMetricBaseImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawInactiveUsersByApplicationMetricBaseImpl) InactiveUsersByApplicationMetricBase() BaseInactiveUsersByApplicationMetricBaseImpl {
	return s.inactiveUsersByApplicationMetricBase
}

func (s RawInactiveUsersByApplicationMetricBaseImpl) Entity() BaseEntityImpl {
	return s.inactiveUsersByApplicationMetricBase.Entity()
}

var _ json.Marshaler = BaseInactiveUsersByApplicationMetricBaseImpl{}

func (s BaseInactiveUsersByApplicationMetricBaseImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseInactiveUsersByApplicationMetricBaseImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseInactiveUsersByApplicationMetricBaseImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseInactiveUsersByApplicationMetricBaseImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.inactiveUsersByApplicationMetricBase"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseInactiveUsersByApplicationMetricBaseImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalInactiveUsersByApplicationMetricBaseImplementation(input []byte) (InactiveUsersByApplicationMetricBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InactiveUsersByApplicationMetricBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.dailyInactiveUsersByApplicationMetric") {
		var out DailyInactiveUsersByApplicationMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DailyInactiveUsersByApplicationMetric: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.monthlyInactiveUsersByApplicationMetric") {
		var out MonthlyInactiveUsersByApplicationMetric
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MonthlyInactiveUsersByApplicationMetric: %+v", err)
		}
		return out, nil
	}

	var parent BaseInactiveUsersByApplicationMetricBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseInactiveUsersByApplicationMetricBaseImpl: %+v", err)
	}

	return RawInactiveUsersByApplicationMetricBaseImpl{
		inactiveUsersByApplicationMetricBase: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
