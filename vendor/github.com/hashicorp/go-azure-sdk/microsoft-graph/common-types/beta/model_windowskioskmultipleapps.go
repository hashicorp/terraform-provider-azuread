package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskAppConfiguration = WindowsKioskMultipleApps{}

type WindowsKioskMultipleApps struct {
	// This setting allows access to Downloads folder in file explorer.
	AllowAccessToDownloadsFolder *bool `json:"allowAccessToDownloadsFolder,omitempty"`

	// These are the only Windows Store Apps that will be available to launch from the Start menu. This collection can
	// contain a maximum of 128 elements.
	Apps *[]WindowsKioskAppBase `json:"apps,omitempty"`

	// This setting indicates that desktop apps are allowed. Default to true.
	DisallowDesktopApps *bool `json:"disallowDesktopApps,omitempty"`

	// This setting allows the admin to specify whether the Task Bar is shown or not.
	ShowTaskBar *bool `json:"showTaskBar,omitempty"`

	// Allows admins to override the default Start layout and prevents the user from changing it. The layout is modified by
	// specifying an XML file based on a layout modification schema. XML needs to be in Binary format.
	StartMenuLayoutXml nullable.Type[string] `json:"startMenuLayoutXml,omitempty"`

	// Fields inherited from WindowsKioskAppConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsKioskMultipleApps) WindowsKioskAppConfiguration() BaseWindowsKioskAppConfigurationImpl {
	return BaseWindowsKioskAppConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskMultipleApps{}

func (s WindowsKioskMultipleApps) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskMultipleApps
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskMultipleApps: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskMultipleApps: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskMultipleApps"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskMultipleApps: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsKioskMultipleApps{}

func (s *WindowsKioskMultipleApps) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AllowAccessToDownloadsFolder *bool                 `json:"allowAccessToDownloadsFolder,omitempty"`
		DisallowDesktopApps          *bool                 `json:"disallowDesktopApps,omitempty"`
		ShowTaskBar                  *bool                 `json:"showTaskBar,omitempty"`
		StartMenuLayoutXml           nullable.Type[string] `json:"startMenuLayoutXml,omitempty"`
		ODataId                      *string               `json:"@odata.id,omitempty"`
		ODataType                    *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AllowAccessToDownloadsFolder = decoded.AllowAccessToDownloadsFolder
	s.DisallowDesktopApps = decoded.DisallowDesktopApps
	s.ShowTaskBar = decoded.ShowTaskBar
	s.StartMenuLayoutXml = decoded.StartMenuLayoutXml
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsKioskMultipleApps into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["apps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Apps into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsKioskAppBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsKioskAppBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Apps' for 'WindowsKioskMultipleApps': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Apps = &output
	}

	return nil
}
