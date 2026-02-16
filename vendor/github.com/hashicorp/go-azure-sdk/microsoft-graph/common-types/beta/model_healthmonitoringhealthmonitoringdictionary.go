package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringHealthMonitoringDictionary interface {
	HealthMonitoringDictionary
	HealthMonitoringHealthMonitoringDictionary() BaseHealthMonitoringHealthMonitoringDictionaryImpl
}

var _ HealthMonitoringHealthMonitoringDictionary = BaseHealthMonitoringHealthMonitoringDictionaryImpl{}

type BaseHealthMonitoringHealthMonitoringDictionaryImpl struct {

	// Fields inherited from HealthMonitoringDictionary

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseHealthMonitoringHealthMonitoringDictionaryImpl) HealthMonitoringHealthMonitoringDictionary() BaseHealthMonitoringHealthMonitoringDictionaryImpl {
	return s
}

func (s BaseHealthMonitoringHealthMonitoringDictionaryImpl) HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl {
	return BaseHealthMonitoringDictionaryImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ HealthMonitoringHealthMonitoringDictionary = RawHealthMonitoringHealthMonitoringDictionaryImpl{}

// RawHealthMonitoringHealthMonitoringDictionaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawHealthMonitoringHealthMonitoringDictionaryImpl struct {
	healthMonitoringHealthMonitoringDictionary BaseHealthMonitoringHealthMonitoringDictionaryImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawHealthMonitoringHealthMonitoringDictionaryImpl) HealthMonitoringHealthMonitoringDictionary() BaseHealthMonitoringHealthMonitoringDictionaryImpl {
	return s.healthMonitoringHealthMonitoringDictionary
}

func (s RawHealthMonitoringHealthMonitoringDictionaryImpl) HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl {
	return s.healthMonitoringHealthMonitoringDictionary.HealthMonitoringDictionary()
}

var _ json.Marshaler = BaseHealthMonitoringHealthMonitoringDictionaryImpl{}

func (s BaseHealthMonitoringHealthMonitoringDictionaryImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseHealthMonitoringHealthMonitoringDictionaryImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseHealthMonitoringHealthMonitoringDictionaryImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseHealthMonitoringHealthMonitoringDictionaryImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.healthMonitoring.healthMonitoringDictionary"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseHealthMonitoringHealthMonitoringDictionaryImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalHealthMonitoringHealthMonitoringDictionaryImplementation(input []byte) (HealthMonitoringHealthMonitoringDictionary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringHealthMonitoringDictionary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.documentation") {
		var out HealthMonitoringDocumentation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringDocumentation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.signals") {
		var out HealthMonitoringSignals
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringSignals: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.supportingData") {
		var out HealthMonitoringSupportingData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringSupportingData: %+v", err)
		}
		return out, nil
	}

	var parent BaseHealthMonitoringHealthMonitoringDictionaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseHealthMonitoringHealthMonitoringDictionaryImpl: %+v", err)
	}

	return RawHealthMonitoringHealthMonitoringDictionaryImpl{
		healthMonitoringHealthMonitoringDictionary: parent,
		Type:   value,
		Values: temp,
	}, nil

}
