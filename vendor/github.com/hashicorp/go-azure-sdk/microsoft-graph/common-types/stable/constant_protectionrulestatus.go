package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionRuleStatus string

const (
	ProtectionRuleStatus_Active              ProtectionRuleStatus = "active"
	ProtectionRuleStatus_Completed           ProtectionRuleStatus = "completed"
	ProtectionRuleStatus_CompletedWithErrors ProtectionRuleStatus = "completedWithErrors"
	ProtectionRuleStatus_Draft               ProtectionRuleStatus = "draft"
)

func PossibleValuesForProtectionRuleStatus() []string {
	return []string{
		string(ProtectionRuleStatus_Active),
		string(ProtectionRuleStatus_Completed),
		string(ProtectionRuleStatus_CompletedWithErrors),
		string(ProtectionRuleStatus_Draft),
	}
}

func (s *ProtectionRuleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionRuleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionRuleStatus(input string) (*ProtectionRuleStatus, error) {
	vals := map[string]ProtectionRuleStatus{
		"active":              ProtectionRuleStatus_Active,
		"completed":           ProtectionRuleStatus_Completed,
		"completedwitherrors": ProtectionRuleStatus_CompletedWithErrors,
		"draft":               ProtectionRuleStatus_Draft,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionRuleStatus(input)
	return &out, nil
}
