package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringDictionary interface {
	HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl
}

var _ HealthMonitoringDictionary = BaseHealthMonitoringDictionaryImpl{}

type BaseHealthMonitoringDictionaryImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseHealthMonitoringDictionaryImpl) HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl {
	return s
}

var _ HealthMonitoringDictionary = RawHealthMonitoringDictionaryImpl{}

// RawHealthMonitoringDictionaryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawHealthMonitoringDictionaryImpl struct {
	healthMonitoringDictionary BaseHealthMonitoringDictionaryImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawHealthMonitoringDictionaryImpl) HealthMonitoringDictionary() BaseHealthMonitoringDictionaryImpl {
	return s.healthMonitoringDictionary
}

func UnmarshalHealthMonitoringDictionaryImplementation(input []byte) (HealthMonitoringDictionary, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling HealthMonitoringDictionary into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.healthMonitoring.healthMonitoringDictionary") {
		var out HealthMonitoringHealthMonitoringDictionary
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HealthMonitoringHealthMonitoringDictionary: %+v", err)
		}
		return out, nil
	}

	var parent BaseHealthMonitoringDictionaryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseHealthMonitoringDictionaryImpl: %+v", err)
	}

	return RawHealthMonitoringDictionaryImpl{
		healthMonitoringDictionary: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
