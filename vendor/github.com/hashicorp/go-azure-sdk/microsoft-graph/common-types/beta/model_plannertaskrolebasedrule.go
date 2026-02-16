package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskRoleBasedRule struct {
	// Default rule that applies when a property or action-specific rule is not provided. Possible values are: Allow, Block
	DefaultRule nullable.Type[string] `json:"defaultRule,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Rules for specific properties and actions.
	PropertyRule *PlannerTaskPropertyRule `json:"propertyRule,omitempty"`

	// The role these rules apply to.
	Role PlannerTaskConfigurationRoleBase `json:"role"`
}

var _ json.Unmarshaler = &PlannerTaskRoleBasedRule{}

func (s *PlannerTaskRoleBasedRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DefaultRule  nullable.Type[string]    `json:"defaultRule,omitempty"`
		ODataId      *string                  `json:"@odata.id,omitempty"`
		ODataType    *string                  `json:"@odata.type,omitempty"`
		PropertyRule *PlannerTaskPropertyRule `json:"propertyRule,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DefaultRule = decoded.DefaultRule
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.PropertyRule = decoded.PropertyRule

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlannerTaskRoleBasedRule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["role"]; ok {
		impl, err := UnmarshalPlannerTaskConfigurationRoleBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Role' for 'PlannerTaskRoleBasedRule': %+v", err)
		}
		s.Role = impl
	}

	return nil
}
