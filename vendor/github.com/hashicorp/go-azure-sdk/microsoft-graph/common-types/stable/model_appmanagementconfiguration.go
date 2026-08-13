package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppManagementConfiguration interface {
	AppManagementConfiguration() BaseAppManagementConfigurationImpl
}

var _ AppManagementConfiguration = BaseAppManagementConfigurationImpl{}

type BaseAppManagementConfigurationImpl struct {
	// Collection of keyCredential restrictions settings to be applied to an application or service principal.
	KeyCredentials *[]KeyCredentialConfiguration `json:"keyCredentials,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Collection of password restrictions settings to be applied to an application or service principal.
	PasswordCredentials *[]PasswordCredentialConfiguration `json:"passwordCredentials,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAppManagementConfigurationImpl) AppManagementConfiguration() BaseAppManagementConfigurationImpl {
	return s
}

var _ AppManagementConfiguration = RawAppManagementConfigurationImpl{}

// RawAppManagementConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawAppManagementConfigurationImpl struct {
	appManagementConfiguration BaseAppManagementConfigurationImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawAppManagementConfigurationImpl) AppManagementConfiguration() BaseAppManagementConfigurationImpl {
	return s.appManagementConfiguration
}

func (s RawAppManagementConfigurationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalAppManagementConfigurationImplementation(input []byte) (AppManagementConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AppManagementConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.appManagementApplicationConfiguration") {
		var out AppManagementApplicationConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppManagementApplicationConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.appManagementServicePrincipalConfiguration") {
		var out AppManagementServicePrincipalConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AppManagementServicePrincipalConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.customAppManagementConfiguration") {
		var out CustomAppManagementConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomAppManagementConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseAppManagementConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAppManagementConfigurationImpl: %+v", err)
	}

	return RawAppManagementConfigurationImpl{
		appManagementConfiguration: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
