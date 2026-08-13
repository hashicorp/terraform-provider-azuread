package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionEndpointConfiguration interface {
	CustomExtensionEndpointConfiguration() BaseCustomExtensionEndpointConfigurationImpl
}

var _ CustomExtensionEndpointConfiguration = BaseCustomExtensionEndpointConfigurationImpl{}

type BaseCustomExtensionEndpointConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomExtensionEndpointConfigurationImpl) CustomExtensionEndpointConfiguration() BaseCustomExtensionEndpointConfigurationImpl {
	return s
}

var _ CustomExtensionEndpointConfiguration = RawCustomExtensionEndpointConfigurationImpl{}

// RawCustomExtensionEndpointConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawCustomExtensionEndpointConfigurationImpl struct {
	customExtensionEndpointConfiguration BaseCustomExtensionEndpointConfigurationImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawCustomExtensionEndpointConfigurationImpl) CustomExtensionEndpointConfiguration() BaseCustomExtensionEndpointConfigurationImpl {
	return s.customExtensionEndpointConfiguration
}

func (s RawCustomExtensionEndpointConfigurationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalCustomExtensionEndpointConfigurationImplementation(input []byte) (CustomExtensionEndpointConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionEndpointConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.httpRequestEndpoint") {
		var out HttpRequestEndpoint
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into HttpRequestEndpoint: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.logicAppTriggerEndpointConfiguration") {
		var out LogicAppTriggerEndpointConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LogicAppTriggerEndpointConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomExtensionEndpointConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomExtensionEndpointConfigurationImpl: %+v", err)
	}

	return RawCustomExtensionEndpointConfigurationImpl{
		customExtensionEndpointConfiguration: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
