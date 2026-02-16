package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRule = Win32LobAppPowerShellScriptRule{}

type Win32LobAppPowerShellScriptRule struct {
	// The script output comparison value. Do not specify a value if the rule is used for detection.
	ComparisonValue nullable.Type[string] `json:"comparisonValue,omitempty"`

	// The display name for the rule. Do not specify this value if the rule is used for detection.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// A value indicating whether a signature check is enforced.
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`

	// Contains all supported Powershell Script output detection type.
	OperationType *Win32LobAppPowerShellScriptRuleOperationType `json:"operationType,omitempty"`

	// Contains properties for detection operator.
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`

	// A value indicating whether the script should run as 32-bit.
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`

	// The execution context of the script. Do not specify this value if the rule is used for detection. Script detection
	// rules will run in the same context as the associated app install context. Possible values are: system, user.
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`

	// The base64-encoded script content.
	ScriptContent nullable.Type[string] `json:"scriptContent,omitempty"`

	// Fields inherited from Win32LobAppRule

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Contains rule types for Win32 LOB apps.
	RuleType *Win32LobAppRuleType `json:"ruleType,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Win32LobAppPowerShellScriptRule) Win32LobAppRule() BaseWin32LobAppRuleImpl {
	return BaseWin32LobAppRuleImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RuleType:  s.RuleType,
	}
}

var _ json.Marshaler = Win32LobAppPowerShellScriptRule{}

func (s Win32LobAppPowerShellScriptRule) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppPowerShellScriptRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppPowerShellScriptRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppPowerShellScriptRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppPowerShellScriptRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppPowerShellScriptRule: %+v", err)
	}

	return encoded, nil
}
