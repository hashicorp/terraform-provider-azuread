package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WindowsKioskAppConfiguration = WindowsKioskSingleUWPApp{}

type WindowsKioskSingleUWPApp struct {
	UwpApp *WindowsKioskUWPApp `json:"uwpApp,omitempty"`

	// Fields inherited from WindowsKioskAppConfiguration

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s WindowsKioskSingleUWPApp) WindowsKioskAppConfiguration() BaseWindowsKioskAppConfigurationImpl {
	return BaseWindowsKioskAppConfigurationImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsKioskSingleUWPApp{}

func (s WindowsKioskSingleUWPApp) MarshalJSON() ([]byte, error) {
	type wrapper WindowsKioskSingleUWPApp
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsKioskSingleUWPApp: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskSingleUWPApp: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsKioskSingleUWPApp"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsKioskSingleUWPApp: %+v", err)
	}

	return encoded, nil
}
