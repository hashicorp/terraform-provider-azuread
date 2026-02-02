package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRequirement = Win32LobAppPowerShellScriptRequirement{}

type Win32LobAppPowerShellScriptRequirement struct {
	// Contains all supported Powershell Script output detection type.
	DetectionType *Win32LobAppPowerShellScriptDetectionType `json:"detectionType,omitempty"`

	// The unique display name for this rule
	DisplayName *string `json:"displayName,omitempty"`

	// A value indicating whether signature check is enforced
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`

	// A value indicating whether this script should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`

	// Indicates the type of execution context the app runs in.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`

	// The base64 encoded script content to detect Win32 Line of Business (LoB) app
	ScriptContent *string `json:"scriptContent,omitempty"`

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

func (s Win32LobAppPowerShellScriptRequirement) Win32LobAppRequirement() BaseWin32LobAppRequirementImpl {
	return BaseWin32LobAppRequirementImpl{
		DetectionValue: s.DetectionValue,
		ODataId:        s.ODataId,
		ODataType:      s.ODataType,
		Operator:       s.Operator,
	}
}

var _ json.Marshaler = Win32LobAppPowerShellScriptRequirement{}

func (s Win32LobAppPowerShellScriptRequirement) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppPowerShellScriptRequirement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppPowerShellScriptRequirement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppPowerShellScriptRequirement: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppPowerShellScriptRequirement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppPowerShellScriptRequirement: %+v", err)
	}

	return encoded, nil
}
