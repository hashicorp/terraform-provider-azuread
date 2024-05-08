// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"
)

type RoleManagementPolicyRuleId struct {
	RuleId string
}

func NewRoleManagementPolicyRuleID(ruleId string) *RoleManagementPolicyRuleId {
	return &RoleManagementPolicyRuleId{
		RuleId: ruleId,
	}
}

func ParseRoleManagementPolicyRuleID(input string) (*RoleManagementPolicyRuleId, error) {
	parts := strings.Split(input, "_")
	if len(parts) == 3 && parts[0] == "Notification" {
		return nil, fmt.Errorf("parsing RoleManagementPolicyRuleId: invalid format")
	} else if len(parts) == 4 && parts[0] != "Notification" {
		return nil, fmt.Errorf("parsing RoleManagementPolicyRuleId: invalid format")
	}

	return &RoleManagementPolicyRuleId{
		RuleId: input,
	}, nil
}

func (id *RoleManagementPolicyRuleId) ID() string {
	return id.RuleId
}

func (id *RoleManagementPolicyRuleId) String() string {
	return fmt.Sprintf("Role Management Policy Assignment ID: %s", id.RuleId)
}
