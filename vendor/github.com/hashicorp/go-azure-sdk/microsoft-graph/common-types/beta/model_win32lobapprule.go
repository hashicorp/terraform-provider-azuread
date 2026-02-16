package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRule interface {
	Win32LobAppRule() BaseWin32LobAppRuleImpl
}

var _ Win32LobAppRule = BaseWin32LobAppRuleImpl{}

type BaseWin32LobAppRuleImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains rule types for Win32 LOB apps.
	RuleType *Win32LobAppRuleType `json:"ruleType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWin32LobAppRuleImpl) Win32LobAppRule() BaseWin32LobAppRuleImpl {
	return s
}

var _ Win32LobAppRule = RawWin32LobAppRuleImpl{}

// RawWin32LobAppRuleImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWin32LobAppRuleImpl struct {
	win32LobAppRule BaseWin32LobAppRuleImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawWin32LobAppRuleImpl) Win32LobAppRule() BaseWin32LobAppRuleImpl {
	return s.win32LobAppRule
}

func UnmarshalWin32LobAppRuleImplementation(input []byte) (Win32LobAppRule, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppRule into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppFileSystemRule") {
		var out Win32LobAppFileSystemRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppFileSystemRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppPowerShellScriptRule") {
		var out Win32LobAppPowerShellScriptRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppPowerShellScriptRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppProductCodeRule") {
		var out Win32LobAppProductCodeRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppProductCodeRule: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppRegistryRule") {
		var out Win32LobAppRegistryRule
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppRegistryRule: %+v", err)
		}
		return out, nil
	}

	var parent BaseWin32LobAppRuleImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWin32LobAppRuleImpl: %+v", err)
	}

	return RawWin32LobAppRuleImpl{
		win32LobAppRule: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
