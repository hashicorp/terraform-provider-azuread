package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppRule = Win32LobAppProductCodeRule{}

type Win32LobAppProductCodeRule struct {
	// The product code of the app.
	ProductCode nullable.Type[string] `json:"productCode,omitempty"`

	// The product version comparison value.
	ProductVersion nullable.Type[string] `json:"productVersion,omitempty"`

	// Contains properties for detection operator.
	ProductVersionOperator *Win32LobAppRuleOperator `json:"productVersionOperator,omitempty"`

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

func (s Win32LobAppProductCodeRule) Win32LobAppRule() BaseWin32LobAppRuleImpl {
	return BaseWin32LobAppRuleImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		RuleType:  s.RuleType,
	}
}

var _ json.Marshaler = Win32LobAppProductCodeRule{}

func (s Win32LobAppProductCodeRule) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppProductCodeRule
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppProductCodeRule: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppProductCodeRule: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppProductCodeRule"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppProductCodeRule: %+v", err)
	}

	return encoded, nil
}
