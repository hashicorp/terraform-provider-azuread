package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionApp interface {
	WindowsInformationProtectionApp() BaseWindowsInformationProtectionAppImpl
}

var _ WindowsInformationProtectionApp = BaseWindowsInformationProtectionAppImpl{}

type BaseWindowsInformationProtectionAppImpl struct {
	// If true, app is denied protection or exemption.
	Denied *bool `json:"denied,omitempty"`

	// The app's description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// App display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The product name.
	ProductName nullable.Type[string] `json:"productName,omitempty"`

	// The publisher name
	PublisherName nullable.Type[string] `json:"publisherName,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsInformationProtectionAppImpl) WindowsInformationProtectionApp() BaseWindowsInformationProtectionAppImpl {
	return s
}

var _ WindowsInformationProtectionApp = RawWindowsInformationProtectionAppImpl{}

// RawWindowsInformationProtectionAppImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsInformationProtectionAppImpl struct {
	windowsInformationProtectionApp BaseWindowsInformationProtectionAppImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawWindowsInformationProtectionAppImpl) WindowsInformationProtectionApp() BaseWindowsInformationProtectionAppImpl {
	return s.windowsInformationProtectionApp
}

func UnmarshalWindowsInformationProtectionAppImplementation(input []byte) (WindowsInformationProtectionApp, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsInformationProtectionApp into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionDesktopApp") {
		var out WindowsInformationProtectionDesktopApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionDesktopApp: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsInformationProtectionStoreApp") {
		var out WindowsInformationProtectionStoreApp
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsInformationProtectionStoreApp: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsInformationProtectionAppImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsInformationProtectionAppImpl: %+v", err)
	}

	return RawWindowsInformationProtectionAppImpl{
		windowsInformationProtectionApp: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
