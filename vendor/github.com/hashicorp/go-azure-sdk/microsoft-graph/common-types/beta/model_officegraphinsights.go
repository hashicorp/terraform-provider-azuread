package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeGraphInsights interface {
	Entity
	OfficeGraphInsights() BaseOfficeGraphInsightsImpl
}

var _ OfficeGraphInsights = BaseOfficeGraphInsightsImpl{}

type BaseOfficeGraphInsightsImpl struct {
	// Access this property from the derived type itemInsights.
	Shared *[]SharedInsight `json:"shared,omitempty"`

	// Access this property from the derived type itemInsights.
	Trending *[]Trending `json:"trending,omitempty"`

	// Access this property from the derived type itemInsights.
	Used *[]UsedInsight `json:"used,omitempty"`

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

func (s BaseOfficeGraphInsightsImpl) OfficeGraphInsights() BaseOfficeGraphInsightsImpl {
	return s
}

func (s BaseOfficeGraphInsightsImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ OfficeGraphInsights = RawOfficeGraphInsightsImpl{}

// RawOfficeGraphInsightsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawOfficeGraphInsightsImpl struct {
	officeGraphInsights BaseOfficeGraphInsightsImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawOfficeGraphInsightsImpl) OfficeGraphInsights() BaseOfficeGraphInsightsImpl {
	return s.officeGraphInsights
}

func (s RawOfficeGraphInsightsImpl) Entity() BaseEntityImpl {
	return s.officeGraphInsights.Entity()
}

var _ json.Marshaler = BaseOfficeGraphInsightsImpl{}

func (s BaseOfficeGraphInsightsImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseOfficeGraphInsightsImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseOfficeGraphInsightsImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseOfficeGraphInsightsImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.officeGraphInsights"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseOfficeGraphInsightsImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalOfficeGraphInsightsImplementation(input []byte) (OfficeGraphInsights, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling OfficeGraphInsights into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.itemInsights") {
		var out ItemInsights
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ItemInsights: %+v", err)
		}
		return out, nil
	}

	var parent BaseOfficeGraphInsightsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseOfficeGraphInsightsImpl: %+v", err)
	}

	return RawOfficeGraphInsightsImpl{
		officeGraphInsights: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
