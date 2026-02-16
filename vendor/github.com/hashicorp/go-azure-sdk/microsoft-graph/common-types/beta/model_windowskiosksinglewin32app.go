package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskAppConfiguration = WindowsKioskSingleWin32App{}

type WindowsKioskSingleWin32App struct {
	Win32App *WindowsKioskWin32App `json:"win32App,omitempty"`

	// Fields inherited from WindowsKioskAppConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsKioskSingleWin32App) WindowsKioskAppConfiguration() BaseWindowsKioskAppConfigurationImpl {
	return BaseWindowsKioskAppConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskSingleWin32App{}

func (s WindowsKioskSingleWin32App) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskSingleWin32App
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskSingleWin32App: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskSingleWin32App: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskSingleWin32App"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskSingleWin32App: %+v", err)
	}

	return encoded, nil
}
