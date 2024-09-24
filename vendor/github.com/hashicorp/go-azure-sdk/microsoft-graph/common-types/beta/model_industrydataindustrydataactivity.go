package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivity interface {
	Entity
	IndustryDataIndustryDataActivity() BaseIndustryDataIndustryDataActivityImpl
}

var _ IndustryDataIndustryDataActivity = BaseIndustryDataIndustryDataActivityImpl{}

type BaseIndustryDataIndustryDataActivityImpl struct {
	// The name of the activity. Maximum supported length is 100 characters.
	DisplayName *string `json:"displayName,omitempty"`

	ReadinessStatus *IndustryDataReadinessStatus `json:"readinessStatus,omitempty"`

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

func (s BaseIndustryDataIndustryDataActivityImpl) IndustryDataIndustryDataActivity() BaseIndustryDataIndustryDataActivityImpl {
	return s
}

func (s BaseIndustryDataIndustryDataActivityImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ IndustryDataIndustryDataActivity = RawIndustryDataIndustryDataActivityImpl{}

// RawIndustryDataIndustryDataActivityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawIndustryDataIndustryDataActivityImpl struct {
	industryDataIndustryDataActivity BaseIndustryDataIndustryDataActivityImpl
	Type                             string
	Values                           map[string]interface{}
}

func (s RawIndustryDataIndustryDataActivityImpl) IndustryDataIndustryDataActivity() BaseIndustryDataIndustryDataActivityImpl {
	return s.industryDataIndustryDataActivity
}

func (s RawIndustryDataIndustryDataActivityImpl) Entity() BaseEntityImpl {
	return s.industryDataIndustryDataActivity.Entity()
}

var _ json.Marshaler = BaseIndustryDataIndustryDataActivityImpl{}

func (s BaseIndustryDataIndustryDataActivityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseIndustryDataIndustryDataActivityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseIndustryDataIndustryDataActivityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseIndustryDataIndustryDataActivityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.industryData.industryDataActivity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseIndustryDataIndustryDataActivityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalIndustryDataIndustryDataActivityImplementation(input []byte) (IndustryDataIndustryDataActivity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataIndustryDataActivity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.inboundFlow") {
		var out IndustryDataInboundFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataInboundFlow: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataIndustryDataActivityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataIndustryDataActivityImpl: %+v", err)
	}

	return RawIndustryDataIndustryDataActivityImpl{
		industryDataIndustryDataActivity: parent,
		Type:                             value,
		Values:                           temp,
	}, nil

}
