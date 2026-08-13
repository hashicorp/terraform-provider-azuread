package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataPasswordSettings interface {
	IndustryDataPasswordSettings() BaseIndustryDataPasswordSettingsImpl
}

var _ IndustryDataPasswordSettings = BaseIndustryDataPasswordSettingsImpl{}

type BaseIndustryDataPasswordSettingsImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseIndustryDataPasswordSettingsImpl) IndustryDataPasswordSettings() BaseIndustryDataPasswordSettingsImpl {
	return s
}

var _ IndustryDataPasswordSettings = RawIndustryDataPasswordSettingsImpl{}

// RawIndustryDataPasswordSettingsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawIndustryDataPasswordSettingsImpl struct {
	industryDataPasswordSettings BaseIndustryDataPasswordSettingsImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawIndustryDataPasswordSettingsImpl) IndustryDataPasswordSettings() BaseIndustryDataPasswordSettingsImpl {
	return s.industryDataPasswordSettings
}

func (s RawIndustryDataPasswordSettingsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalIndustryDataPasswordSettingsImplementation(input []byte) (IndustryDataPasswordSettings, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IndustryDataPasswordSettings into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.industryData.simplePasswordSettings") {
		var out IndustryDataSimplePasswordSettings
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into IndustryDataSimplePasswordSettings: %+v", err)
		}
		return out, nil
	}

	var parent BaseIndustryDataPasswordSettingsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseIndustryDataPasswordSettingsImpl: %+v", err)
	}

	return RawIndustryDataPasswordSettingsImpl{
		industryDataPasswordSettings: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
