package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingValueDefinition = DeviceManagementConfigurationStringSettingValueDefinition{}

type DeviceManagementConfigurationStringSettingValueDefinition struct {
	// Supported file types for this setting.
	FileTypes *[]string `json:"fileTypes,omitempty"`

	// Pre-defined format of the string. Possible values are: none, email, guid, ip, base64, url, version, xml, date, time,
	// binary, regEx, json, dateTime, surfaceHub.
	Format *DeviceManagementConfigurationStringFormat `json:"format,omitempty"`

	// Regular expression or any xml or json schema that the input string should match
	InputValidationSchema nullable.Type[string] `json:"inputValidationSchema,omitempty"`

	// Specifies whether the setting needs to be treated as a secret. Settings marked as yes will be encrypted in transit
	// and at rest and will be displayed as asterisks when represented in the UX.
	IsSecret nullable.Type[bool] `json:"isSecret,omitempty"`

	// Maximum length of string
	MaximumLength nullable.Type[int64] `json:"maximumLength,omitempty"`

	// Minimum length of string
	MinimumLength nullable.Type[int64] `json:"minimumLength,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingValueDefinition

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationStringSettingValueDefinition) DeviceManagementConfigurationSettingValueDefinition() BaseDeviceManagementConfigurationSettingValueDefinitionImpl {
	return BaseDeviceManagementConfigurationSettingValueDefinitionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationStringSettingValueDefinition{}

func (s DeviceManagementConfigurationStringSettingValueDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationStringSettingValueDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationStringSettingValueDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationStringSettingValueDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationStringSettingValueDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationStringSettingValueDefinition: %+v", err)
	}

	return encoded, nil
}
