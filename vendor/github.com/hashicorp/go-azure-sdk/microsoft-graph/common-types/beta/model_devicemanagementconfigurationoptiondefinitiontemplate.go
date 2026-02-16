package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationOptionDefinitionTemplate struct {
	// Option Children
	Children *[]DeviceManagementConfigurationSettingInstanceTemplate `json:"children,omitempty"`

	// Option ItemId
	ItemId nullable.Type[string] `json:"itemId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &DeviceManagementConfigurationOptionDefinitionTemplate{}

func (s *DeviceManagementConfigurationOptionDefinitionTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ItemId    nullable.Type[string] `json:"itemId,omitempty"`
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ItemId = decoded.ItemId
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationOptionDefinitionTemplate into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'Children' for 'DeviceManagementConfigurationOptionDefinitionTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Children = &output
	}

	return nil
}
