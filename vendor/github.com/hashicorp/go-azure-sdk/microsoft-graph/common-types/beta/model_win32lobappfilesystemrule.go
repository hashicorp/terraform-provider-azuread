package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRule = Win32LobAppFileSystemRule{}

type Win32LobAppFileSystemRule struct {
	// A value indicating whether to expand environment variables in the 32-bit context on 64-bit systems.
	Check32BitOn64System *bool `json:"check32BitOn64System,omitempty"`

	// The file or folder comparison value.
	ComparisonValue nullable.Type[string] `json:"comparisonValue,omitempty"`

	// The file or folder name to look up.
	FileOrFolderName nullable.Type[string] `json:"fileOrFolderName,omitempty"`

	// A list of possible operations for rules used to make determinations about an application based on files or folders.
	// Unless noted, can be used with either detection or requirement rules.
	OperationType *Win32LobAppFileSystemOperationType `json:"operationType,omitempty"`

	// Contains properties for detection operator.
	Operator *Win32LobAppRuleOperator `json:"operator,omitempty"`

	// The file or folder path to look up.
	Path nullable.Type[string] `json:"path,omitempty"`

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

func (s Win32LobAppFileSystemRule) Win32LobAppRule() BaseWin32LobAppRuleImpl {
	return BaseWin32LobAppRuleImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RuleType:  s.RuleType,
	}
}

var _ json.Marshaler = Win32LobAppFileSystemRule{}

func (s Win32LobAppFileSystemRule) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppFileSystemRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppFileSystemRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppFileSystemRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppFileSystemRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppFileSystemRule: %+v", err)
	}

	return encoded, nil
}
