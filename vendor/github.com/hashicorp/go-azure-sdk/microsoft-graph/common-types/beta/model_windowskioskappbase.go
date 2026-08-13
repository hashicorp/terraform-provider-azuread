package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppBase interface {
	WindowsKioskAppBase() BaseWindowsKioskAppBaseImpl
}

var _ WindowsKioskAppBase = BaseWindowsKioskAppBaseImpl{}

type BaseWindowsKioskAppBaseImpl struct {
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

func (s BaseWindowsKioskAppBaseImpl) WindowsKioskAppBase() BaseWindowsKioskAppBaseImpl {
	return s
}

var _ WindowsKioskAppBase = RawWindowsKioskAppBaseImpl{}

// RawWindowsKioskAppBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawWindowsKioskAppBaseImpl struct {
	windowsKioskAppBase BaseWindowsKioskAppBaseImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawWindowsKioskAppBaseImpl) WindowsKioskAppBase() BaseWindowsKioskAppBaseImpl {
	return s.windowsKioskAppBase
}

func (s RawWindowsKioskAppBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalWindowsKioskAppBaseImplementation(input []byte) (WindowsKioskAppBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskAppBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskDesktopApp") {
		var out WindowsKioskDesktopApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskDesktopApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskUWPApp") {
		var out WindowsKioskUWPApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskUWPApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskWin32App") {
		var out WindowsKioskWin32App
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskWin32App: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsKioskAppBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsKioskAppBaseImpl: %+v", err)
	}

	return RawWindowsKioskAppBaseImpl{
		windowsKioskAppBase: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
