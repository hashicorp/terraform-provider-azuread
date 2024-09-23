package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationOptionDefinition struct {
	// List of Settings that depends on this option
	DependedOnBy *[]DeviceManagementConfigurationSettingDependedOnBy `json:"dependedOnBy,omitempty"`

	// List of dependent settings for this option
	DependentOn *[]DeviceManagementConfigurationDependentOn `json:"dependentOn,omitempty"`

	// Description of the option
	Description nullable.Type[string] `json:"description,omitempty"`

	// Friendly name of the option
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Help text of the option
	HelpText nullable.Type[string] `json:"helpText,omitempty"`

	// Identifier of option
	ItemId nullable.Type[string] `json:"itemId,omitempty"`

	// Name of the option
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Value of the option
	OptionValue DeviceManagementConfigurationSettingValue `json:"optionValue"`
}

var _ json.Unmarshaler = &DeviceManagementConfigurationOptionDefinition{}

func (s *DeviceManagementConfigurationOptionDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DependedOnBy *[]DeviceManagementConfigurationSettingDependedOnBy `json:"dependedOnBy,omitempty"`
		DependentOn  *[]DeviceManagementConfigurationDependentOn         `json:"dependentOn,omitempty"`
		Description  nullable.Type[string]                               `json:"description,omitempty"`
		DisplayName  nullable.Type[string]                               `json:"displayName,omitempty"`
		HelpText     nullable.Type[string]                               `json:"helpText,omitempty"`
		ItemId       nullable.Type[string]                               `json:"itemId,omitempty"`
		Name         nullable.Type[string]                               `json:"name,omitempty"`
		ODataId      *string                                             `json:"@odata.id,omitempty"`
		ODataType    *string                                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DependedOnBy = decoded.DependedOnBy
	s.DependentOn = decoded.DependentOn
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.HelpText = decoded.HelpText
	s.ItemId = decoded.ItemId
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagementConfigurationOptionDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["optionValue"]; ok {
		impl, err := UnmarshalDeviceManagementConfigurationSettingValueImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'OptionValue' for 'DeviceManagementConfigurationOptionDefinition': %+v", err)
		}
		s.OptionValue = impl
	}

	return nil
}
