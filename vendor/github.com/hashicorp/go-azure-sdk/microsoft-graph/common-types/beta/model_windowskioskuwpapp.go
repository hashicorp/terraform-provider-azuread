package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskAppBase = WindowsKioskUWPApp{}

type WindowsKioskUWPApp struct {
	// This references an Intune App that will be target to the same assignments as Kiosk configuration
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// This is the only Application User Model ID (AUMID) that will be available to launch use while in Kiosk Mode
	AppUserModelId *string `json:"appUserModelId,omitempty"`

	// This references an contained App from an Intune App
	ContainedAppId nullable.Type[string] `json:"containedAppId,omitempty"`

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

func (s WindowsKioskUWPApp) WindowsKioskAppBase() BaseWindowsKioskAppBaseImpl {
	return BaseWindowsKioskAppBaseImpl{
		AppType:             s.AppType,
		AutoLaunch:          s.AutoLaunch,
		Name:                s.Name,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartLayoutTileSize: s.StartLayoutTileSize,
	}
}

var _ json.Marshaler = WindowsKioskUWPApp{}

func (s WindowsKioskUWPApp) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskUWPApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskUWPApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskUWPApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskUWPApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskUWPApp: %+v", err)
	}

	return encoded, nil
}
