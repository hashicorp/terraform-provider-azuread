package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskAppBase = WindowsKioskWin32App{}

type WindowsKioskWin32App struct {
	// This is the classicapppath to be used by v4 Win32 app while in Kiosk Mode
	ClassicAppPath *string `json:"classicAppPath,omitempty"`

	// Edge kiosk (url) for Edge kiosk mode
	EdgeKiosk nullable.Type[string] `json:"edgeKiosk,omitempty"`

	// Edge kiosk idle timeout in minutes for Edge kiosk mode. Valid values 0 to 1440
	EdgeKioskIdleTimeoutMinutes nullable.Type[int64] `json:"edgeKioskIdleTimeoutMinutes,omitempty"`

	// Edge kiosk type
	EdgeKioskType *WindowsEdgeKioskType `json:"edgeKioskType,omitempty"`

	// Edge first run flag for Edge kiosk mode
	EdgeNoFirstRun *bool `json:"edgeNoFirstRun,omitempty"`

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

func (s WindowsKioskWin32App) WindowsKioskAppBase() BaseWindowsKioskAppBaseImpl {
	return BaseWindowsKioskAppBaseImpl{
		AppType:             s.AppType,
		AutoLaunch:          s.AutoLaunch,
		Name:                s.Name,
		ODataId:             s.ODataId,
		ODataType:           s.ODataType,
		StartLayoutTileSize: s.StartLayoutTileSize,
	}
}

var _ json.Marshaler = WindowsKioskWin32App{}

func (s WindowsKioskWin32App) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskWin32App
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskWin32App: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskWin32App: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskWin32App"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskWin32App: %+v", err)
	}

	return encoded, nil
}
