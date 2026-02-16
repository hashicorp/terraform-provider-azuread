package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyRuleTargetOperations string

const (
	UnifiedRoleManagementPolicyRuleTargetOperations_Activate   UnifiedRoleManagementPolicyRuleTargetOperations = "activate"
	UnifiedRoleManagementPolicyRuleTargetOperations_All        UnifiedRoleManagementPolicyRuleTargetOperations = "all"
	UnifiedRoleManagementPolicyRuleTargetOperations_Assign     UnifiedRoleManagementPolicyRuleTargetOperations = "assign"
	UnifiedRoleManagementPolicyRuleTargetOperations_Deactivate UnifiedRoleManagementPolicyRuleTargetOperations = "deactivate"
	UnifiedRoleManagementPolicyRuleTargetOperations_Extend     UnifiedRoleManagementPolicyRuleTargetOperations = "extend"
	UnifiedRoleManagementPolicyRuleTargetOperations_Remove     UnifiedRoleManagementPolicyRuleTargetOperations = "remove"
	UnifiedRoleManagementPolicyRuleTargetOperations_Renew      UnifiedRoleManagementPolicyRuleTargetOperations = "renew"
	UnifiedRoleManagementPolicyRuleTargetOperations_Update     UnifiedRoleManagementPolicyRuleTargetOperations = "update"
)

func PossibleValuesForUnifiedRoleManagementPolicyRuleTargetOperations() []string {
	return []string{
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Activate),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_All),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Assign),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Deactivate),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Extend),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Remove),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Renew),
		string(UnifiedRoleManagementPolicyRuleTargetOperations_Update),
	}
}

func (s *UnifiedRoleManagementPolicyRuleTargetOperations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnifiedRoleManagementPolicyRuleTargetOperations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnifiedRoleManagementPolicyRuleTargetOperations(input string) (*UnifiedRoleManagementPolicyRuleTargetOperations, error) {
	vals := map[string]UnifiedRoleManagementPolicyRuleTargetOperations{
		"activate":   UnifiedRoleManagementPolicyRuleTargetOperations_Activate,
		"all":        UnifiedRoleManagementPolicyRuleTargetOperations_All,
		"assign":     UnifiedRoleManagementPolicyRuleTargetOperations_Assign,
		"deactivate": UnifiedRoleManagementPolicyRuleTargetOperations_Deactivate,
		"extend":     UnifiedRoleManagementPolicyRuleTargetOperations_Extend,
		"remove":     UnifiedRoleManagementPolicyRuleTargetOperations_Remove,
		"renew":      UnifiedRoleManagementPolicyRuleTargetOperations_Renew,
		"update":     UnifiedRoleManagementPolicyRuleTargetOperations_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnifiedRoleManagementPolicyRuleTargetOperations(input)
	return &out, nil
}
