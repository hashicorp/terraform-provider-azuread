package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleMode string

const (
	RuleMode_Audit           RuleMode = "audit"
	RuleMode_AuditAndNotify  RuleMode = "auditAndNotify"
	RuleMode_Enforce         RuleMode = "enforce"
	RuleMode_PendingDeletion RuleMode = "pendingDeletion"
	RuleMode_Test            RuleMode = "test"
)

func PossibleValuesForRuleMode() []string {
	return []string{
		string(RuleMode_Audit),
		string(RuleMode_AuditAndNotify),
		string(RuleMode_Enforce),
		string(RuleMode_PendingDeletion),
		string(RuleMode_Test),
	}
}

func (s *RuleMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRuleMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRuleMode(input string) (*RuleMode, error) {
	vals := map[string]RuleMode{
		"audit":           RuleMode_Audit,
		"auditandnotify":  RuleMode_AuditAndNotify,
		"enforce":         RuleMode_Enforce,
		"pendingdeletion": RuleMode_PendingDeletion,
		"test":            RuleMode_Test,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RuleMode(input)
	return &out, nil
}
