package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRule = Win32LobAppRegistryRule{}

type Win32LobAppRegistryRule struct {
	// A value indicating whether to search the 32-bit registry on 64-bit systems.
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`

	// The registry comparison value.
	ComparisonValue nullable.Type[string] `json:"comparisonValue,omitempty"`

	// The full path of the registry entry containing the value to detect.
	KeyPath nullable.Type[string] `json:"keyPath,omitempty"`

	// A list of possible operations for rules used to make determinations about an application based on registry keys or
	// values. Unless noted, the values can be used with either detection or requirement rules.
	OperationType *Win32LobAppRegistryRuleOperationType `json:"operationType,omitempty"`

	// Contains properties for detection operator.
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`

	// The name of the registry value to detect.
	ValueName nullable.Type[string] `json:"valueName,omitempty"`

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

func (s Win32LobAppRegistryRule) Win32LobAppRule() BaseWin32LobAppRuleImpl {
	return BaseWin32LobAppRuleImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RuleType:  s.RuleType,
	}
}

var _ json.Marshaler = Win32LobAppRegistryRule{}

func (s Win32LobAppRegistryRule) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppRegistryRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppRegistryRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppRegistryRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppRegistryRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppRegistryRule: %+v", err)
	}

	return encoded, nil
}
