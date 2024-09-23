package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppDetection interface {
	Win32LobAppDetection() BaseWin32LobAppDetectionImpl
}

var _ Win32LobAppDetection = BaseWin32LobAppDetectionImpl{}

type BaseWin32LobAppDetectionImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWin32LobAppDetectionImpl) Win32LobAppDetection() BaseWin32LobAppDetectionImpl {
	return s
}

var _ Win32LobAppDetection = RawWin32LobAppDetectionImpl{}

// RawWin32LobAppDetectionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWin32LobAppDetectionImpl struct {
	win32LobAppDetection BaseWin32LobAppDetectionImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawWin32LobAppDetectionImpl) Win32LobAppDetection() BaseWin32LobAppDetectionImpl {
	return s.win32LobAppDetection
}

func UnmarshalWin32LobAppDetectionImplementation(input []byte) (Win32LobAppDetection, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppDetection into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppFileSystemDetection") {
		var out Win32LobAppFileSystemDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppFileSystemDetection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppPowerShellScriptDetection") {
		var out Win32LobAppPowerShellScriptDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppPowerShellScriptDetection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppProductCodeDetection") {
		var out Win32LobAppProductCodeDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppProductCodeDetection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppRegistryDetection") {
		var out Win32LobAppRegistryDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppRegistryDetection: %+v", err)
		}
		return out, nil
	}

	var parent BaseWin32LobAppDetectionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWin32LobAppDetectionImpl: %+v", err)
	}

	return RawWin32LobAppDetectionImpl{
		win32LobAppDetection: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
