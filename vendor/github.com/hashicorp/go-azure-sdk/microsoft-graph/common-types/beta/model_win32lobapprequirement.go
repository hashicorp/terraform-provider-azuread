package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppRequirement interface {
	Win32LobAppRequirement() BaseWin32LobAppRequirementImpl
}

var _ Win32LobAppRequirement = BaseWin32LobAppRequirementImpl{}

type BaseWin32LobAppRequirementImpl struct {
	// The detection value
	DetectionValue nullable.Type[string] `json:"detectionValue,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains properties for detection operator.
	Operator *Win32LobAppDetectionOperator `json:"operator,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseWin32LobAppRequirementImpl) Win32LobAppRequirement() BaseWin32LobAppRequirementImpl {
	return s
}

var _ Win32LobAppRequirement = RawWin32LobAppRequirementImpl{}

// RawWin32LobAppRequirementImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWin32LobAppRequirementImpl struct {
	win32LobAppRequirement BaseWin32LobAppRequirementImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawWin32LobAppRequirementImpl) Win32LobAppRequirement() BaseWin32LobAppRequirementImpl {
	return s.win32LobAppRequirement
}

func UnmarshalWin32LobAppRequirementImplementation(input []byte) (Win32LobAppRequirement, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppRequirement into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppFileSystemRequirement") {
		var out Win32LobAppFileSystemRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppFileSystemRequirement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppPowerShellScriptRequirement") {
		var out Win32LobAppPowerShellScriptRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppPowerShellScriptRequirement: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.win32LobAppRegistryRequirement") {
		var out Win32LobAppRegistryRequirement
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Win32LobAppRegistryRequirement: %+v", err)
		}
		return out, nil
	}

	var parent BaseWin32LobAppRequirementImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWin32LobAppRequirementImpl: %+v", err)
	}

	return RawWin32LobAppRequirementImpl{
		win32LobAppRequirement: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
