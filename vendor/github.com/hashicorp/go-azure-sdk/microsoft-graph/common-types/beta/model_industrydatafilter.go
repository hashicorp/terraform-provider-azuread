package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFilter interface {
	IndustryDataFilter() BaseIndustryDataFilterImpl
}

var _ IndustryDataFilter = BaseIndustryDataFilterImpl{}

type BaseIndustryDataFilterImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIndustryDataFilterImpl) IndustryDataFilter() BaseIndustryDataFilterImpl {
	return s
}

var _ IndustryDataFilter = RawIndustryDataFilterImpl{}

// RawIndustryDataFilterImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIndustryDataFilterImpl struct {
	industryDataFilter BaseIndustryDataFilterImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawIndustryDataFilterImpl) IndustryDataFilter() BaseIndustryDataFilterImpl {
	return s.industryDataFilter
}

func (s RawIndustryDataFilterImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIndustryDataFilterImplementation(input []byte) (IndustryDataFilter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataFilter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.basicFilter") {
		var out IndustryDataBasicFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataBasicFilter: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataFilterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataFilterImpl: %+v", err)
	}

	return RawIndustryDataFilterImpl{
		industryDataFilter: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
