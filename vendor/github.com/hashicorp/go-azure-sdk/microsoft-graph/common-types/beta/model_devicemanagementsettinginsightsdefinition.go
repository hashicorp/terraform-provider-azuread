package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingInsightsDefinition struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting definition id that is being referred to a setting.
	SettingDefinitionId nullable.Type[string] `json:"settingDefinitionId,omitempty"`

	// Data Insights Target Value
	SettingInsight DeviceManagementConfigurationSettingValue `json:"settingInsight"`
}

var _ json.Unmarshaler = &DeviceManagementSettingInsightsDefinition{}

func (s *DeviceManagementSettingInsightsDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId             *string               `json:"@odata.id,omitempty"`
		ODataType           *string               `json:"@odata.type,omitempty"`
		SettingDefinitionId nullable.Type[string] `json:"settingDefinitionId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingDefinitionId = decoded.SettingDefinitionId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementSettingInsightsDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["settingInsight"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSettingValueImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SettingInsight' for 'DeviceManagementSettingInsightsDefinition': %+v", err)
		}
		s.SettingInsight = impl
	}

	return nil
}
