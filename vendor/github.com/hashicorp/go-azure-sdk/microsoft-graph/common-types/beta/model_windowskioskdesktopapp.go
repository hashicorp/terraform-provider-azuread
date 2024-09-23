package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskAppBase = WindowsKioskDesktopApp{}

type WindowsKioskDesktopApp struct {
	// Define the DesktopApplicationID of the app
	DesktopApplicationId nullable.Type[string] `json:"desktopApplicationId,omitempty"`

	// Define the DesktopApplicationLinkPath of the app
	DesktopApplicationLinkPath nullable.Type[string] `json:"desktopApplicationLinkPath,omitempty"`

	// Define the path of a desktop app
	Path *string `json:"path,omitempty"`

	// Fields inherited from WindowsKioskAppBase

	// The type of Windows kiosk app.
	AppType *WindowsKioskAppType `json:"appType,omitempty"`

	// Allow the app to be auto-launched in multi-app kiosk mode
	AutoLaunch *bool `json:"autoLaunch,omitempty"`

	// Represents the friendly name of an app
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The tile size of Windows app in the start layout.
	StartLayoutTileSize *WindowsAppStartLayoutTileSize `json:"startLayoutTileSize,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsKioskDesktopApp) WindowsKioskAppBase() BaseWindowsKioskAppBaseImpl {
	return BaseWindowsKioskAppBaseImpl{
		AppType:             s.AppType,
		AutoLaunch:          s.AutoLaunch,
		Name:                s.Name,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartLayoutTileSize: s.StartLayoutTileSize,
	}
}

var _ json.Marshaler = WindowsKioskDesktopApp{}

func (s WindowsKioskDesktopApp) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskDesktopApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskDesktopApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskDesktopApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskDesktopApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskDesktopApp: %+v", err)
	}

	return encoded, nil
}
