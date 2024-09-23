package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceManagementConfigurationSettingValueDefinition = DeviceManagementConfigurationIntegerSettingValueDefinition{}

type DeviceManagementConfigurationIntegerSettingValueDefinition struct {
	// Maximum allowed value of the integer
	MaximumValue nullable.Type[int64] `json:"maximumValue,omitempty"`

	// Minimum allowed value of the integer
	MinimumValue nullable.Type[int64] `json:"minimumValue,omitempty"`

	// Fields inherited from DeviceManagementConfigurationSettingValueDefinition

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DeviceManagementConfigurationIntegerSettingValueDefinition) DeviceManagementConfigurationSettingValueDefinition() BaseDeviceManagementConfigurationSettingValueDefinitionImpl {
	return BaseDeviceManagementConfigurationSettingValueDefinitionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementConfigurationIntegerSettingValueDefinition{}

func (s DeviceManagementConfigurationIntegerSettingValueDefinition) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementConfigurationIntegerSettingValueDefinition
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementConfigurationIntegerSettingValueDefinition: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementConfigurationIntegerSettingValueDefinition: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementConfigurationIntegerSettingValueDefinition"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementConfigurationIntegerSettingValueDefinition: %+v", err)
	}

	return encoded, nil
}
