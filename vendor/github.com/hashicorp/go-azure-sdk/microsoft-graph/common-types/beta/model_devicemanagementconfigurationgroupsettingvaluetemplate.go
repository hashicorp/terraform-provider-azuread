package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationGroupSettingValueTemplate struct {
	// Group setting value children
	Children *[]DeviceManagementConfigurationSettingInstanceTemplate `json:"children,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Setting Value Template Id
	SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`
}

var _ json.Unmarshaler = &DeviceManagementConfigurationGroupSettingValueTemplate{}

func (s *DeviceManagementConfigurationGroupSettingValueTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId                *string `json:"@odata.id,omitempty"`
		ODataType              *string `json:"@odata.type,omitempty"`
		SettingValueTemplateId *string `json:"settingValueTemplateId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.SettingValueTemplateId = decoded.SettingValueTemplateId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationGroupSettingValueTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["children"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Children into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingInstanceTemplate, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingInstanceTemplateImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Children' for 'DeviceManagementConfigurationGroupSettingValueTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Children = &output
	}

	return nil
}
