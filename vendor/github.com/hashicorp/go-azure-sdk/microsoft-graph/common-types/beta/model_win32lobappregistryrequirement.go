package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRequirement = Win32LobAppRegistryRequirement{}

type Win32LobAppRegistryRequirement struct {
	// A value indicating whether this registry path is for checking 32-bit app on 64-bit system
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`

	// Contains all supported registry data detection type.
	DetectionType *Win32LobAppRegistryDetectionType `json:"detectionType,omitempty"`

	// The registry key path to detect Win32 Line of Business (LoB) app
	KeyPath nullable.Type[string] `json:"keyPath,omitempty"`

	// The registry value name
	ValueName nullable.Type[string] `json:"valueName,omitempty"`

	// Fields inherited from Win32LobAppRequirement

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

func (s Win32LobAppRegistryRequirement) Win32LobAppRequirement() BaseWin32LobAppRequirementImpl {
	return BaseWin32LobAppRequirementImpl{
		DetectionValue: s.DetectionValue,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		Operator:       s.Operator,
	}
}

var _ json.Marshaler = Win32LobAppRegistryRequirement{}

func (s Win32LobAppRegistryRequirement) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppRegistryRequirement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppRegistryRequirement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppRegistryRequirement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppRegistryRequirement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppRegistryRequirement: %+v", err)
	}

	return encoded, nil
}
