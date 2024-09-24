package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHuntingRuleRunStatus string

const (
	SecurityHuntingRuleRunStatus_Completed       SecurityHuntingRuleRunStatus = "completed"
	SecurityHuntingRuleRunStatus_Failed          SecurityHuntingRuleRunStatus = "failed"
	SecurityHuntingRuleRunStatus_PartiallyFailed SecurityHuntingRuleRunStatus = "partiallyFailed"
	SecurityHuntingRuleRunStatus_Running         SecurityHuntingRuleRunStatus = "running"
)

func PossibleValuesForSecurityHuntingRuleRunStatus() []string {
	return []string{
		string(SecurityHuntingRuleRunStatus_Completed),
		string(SecurityHuntingRuleRunStatus_Failed),
		string(SecurityHuntingRuleRunStatus_PartiallyFailed),
		string(SecurityHuntingRuleRunStatus_Running),
	}
}

func (s *SecurityHuntingRuleRunStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHuntingRuleRunStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHuntingRuleRunStatus(input string) (*SecurityHuntingRuleRunStatus, error) {
	vals := map[string]SecurityHuntingRuleRunStatus{
		"completed":       SecurityHuntingRuleRunStatus_Completed,
		"failed":          SecurityHuntingRuleRunStatus_Failed,
		"partiallyfailed": SecurityHuntingRuleRunStatus_PartiallyFailed,
		"running":         SecurityHuntingRuleRunStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHuntingRuleRunStatus(input)
	return &out, nil
}
