package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionAuthenticationConfiguration interface {
	CustomExtensionAuthenticationConfiguration() BaseCustomExtensionAuthenticationConfigurationImpl
}

var _ CustomExtensionAuthenticationConfiguration = BaseCustomExtensionAuthenticationConfigurationImpl{}

type BaseCustomExtensionAuthenticationConfigurationImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseCustomExtensionAuthenticationConfigurationImpl) CustomExtensionAuthenticationConfiguration() BaseCustomExtensionAuthenticationConfigurationImpl {
	return s
}

var _ CustomExtensionAuthenticationConfiguration = RawCustomExtensionAuthenticationConfigurationImpl{}

// RawCustomExtensionAuthenticationConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawCustomExtensionAuthenticationConfigurationImpl struct {
	customExtensionAuthenticationConfiguration BaseCustomExtensionAuthenticationConfigurationImpl
	Type                                       string
	Values                                     map[string]interface{}
}

func (s RawCustomExtensionAuthenticationConfigurationImpl) CustomExtensionAuthenticationConfiguration() BaseCustomExtensionAuthenticationConfigurationImpl {
	return s.customExtensionAuthenticationConfiguration
}

func UnmarshalCustomExtensionAuthenticationConfigurationImplementation(input []byte) (CustomExtensionAuthenticationConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomExtensionAuthenticationConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAdPopTokenAuthentication") {
		var out AzureAdPopTokenAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAdPopTokenAuthentication: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.azureAdTokenAuthentication") {
		var out AzureAdTokenAuthentication
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureAdTokenAuthentication: %+v", err)
		}
		return out, nil
	}

	var parent BaseCustomExtensionAuthenticationConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseCustomExtensionAuthenticationConfigurationImpl: %+v", err)
	}

	return RawCustomExtensionAuthenticationConfigurationImpl{
		customExtensionAuthenticationConfiguration: parent,
		Type:   value,
		Values: temp,
	}, nil

}
