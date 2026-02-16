package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskUser interface {
	WindowsKioskUser() BaseWindowsKioskUserImpl
}

var _ WindowsKioskUser = BaseWindowsKioskUserImpl{}

type BaseWindowsKioskUserImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWindowsKioskUserImpl) WindowsKioskUser() BaseWindowsKioskUserImpl {
	return s
}

var _ WindowsKioskUser = RawWindowsKioskUserImpl{}

// RawWindowsKioskUserImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsKioskUserImpl struct {
	windowsKioskUser BaseWindowsKioskUserImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawWindowsKioskUserImpl) WindowsKioskUser() BaseWindowsKioskUserImpl {
	return s.windowsKioskUser
}

func UnmarshalWindowsKioskUserImplementation(input []byte) (WindowsKioskUser, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsKioskUser into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskActiveDirectoryGroup") {
		var out WindowsKioskActiveDirectoryGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskActiveDirectoryGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskAutologon") {
		var out WindowsKioskAutologon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskAutologon: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskAzureADGroup") {
		var out WindowsKioskAzureADGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskAzureADGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskAzureADUser") {
		var out WindowsKioskAzureADUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskAzureADUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskLocalGroup") {
		var out WindowsKioskLocalGroup
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskLocalGroup: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskLocalUser") {
		var out WindowsKioskLocalUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskLocalUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsKioskVisitor") {
		var out WindowsKioskVisitor
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsKioskVisitor: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsKioskUserImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsKioskUserImpl: %+v", err)
	}

	return RawWindowsKioskUserImpl{
		windowsKioskUser: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
