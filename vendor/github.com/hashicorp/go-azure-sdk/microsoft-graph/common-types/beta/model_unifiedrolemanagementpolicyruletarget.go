package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyRuleTarget struct {
	// The type of caller that's the target of the policy rule. Allowed values are: None, Admin, EndUser.
	Caller nullable.Type[string] `json:"caller,omitempty"`

	// The list of role settings that are enforced and cannot be overridden by child scopes. Use All for all settings.
	EnforcedSettings *[]string `json:"enforcedSettings,omitempty"`

	// The list of role settings that can be inherited by child scopes. Use All for all settings.
	InheritableSettings *[]string `json:"inheritableSettings,omitempty"`

	// The role assignment type that's the target of policy rule. Allowed values are: Eligibility, Assignment.
	Level nullable.Type[string] `json:"level,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The role management operations that are the target of the policy rule. Allowed values are: All, Activate, Deactivate,
	// Assign, Update, Remove, Extend, Renew.
	Operations *[]string `json:"operations,omitempty"`

	TargetObjects *[]DirectoryObject `json:"targetObjects,omitempty"`

	// List of OData IDs for `TargetObjects` to bind to this entity
	TargetObjects_ODataBind *[]string `json:"targetObjects@odata.bind,omitempty"`
}

var _ json.Unmarshaler = &UnifiedRoleManagementPolicyRuleTarget{}

func (s *UnifiedRoleManagementPolicyRuleTarget) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Caller                  nullable.Type[string] `json:"caller,omitempty"`
		EnforcedSettings        *[]string             `json:"enforcedSettings,omitempty"`
		InheritableSettings     *[]string             `json:"inheritableSettings,omitempty"`
		Level                   nullable.Type[string] `json:"level,omitempty"`
		ODataId                 *string               `json:"@odata.id,omitempty"`
		ODataType               *string               `json:"@odata.type,omitempty"`
		Operations              *[]string             `json:"operations,omitempty"`
		TargetObjects_ODataBind *[]string             `json:"targetObjects@odata.bind,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Caller = decoded.Caller
	s.EnforcedSettings = decoded.EnforcedSettings
	s.InheritableSettings = decoded.InheritableSettings
	s.Level = decoded.Level
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Operations = decoded.Operations
	s.TargetObjects_ODataBind = decoded.TargetObjects_ODataBind

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UnifiedRoleManagementPolicyRuleTarget into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["targetObjects"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TargetObjects into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TargetObjects' for 'UnifiedRoleManagementPolicyRuleTarget': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TargetObjects = &output
	}

	return nil
}
