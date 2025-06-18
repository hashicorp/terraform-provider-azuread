package beta

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
	ProtectionRuleStatus_DeleteRequested     ProtectionRuleStatus = "deleteRequested"
	ProtectionRuleStatus_Draft               ProtectionRuleStatus = "draft"
	ProtectionRuleStatus_UpdateRequested     ProtectionRuleStatus = "updateRequested"
)

func PossibleValuesForProtectionRuleStatus() []string {
	return []string{
		string(ProtectionRuleStatus_Active),
		string(ProtectionRuleStatus_Completed),
		string(ProtectionRuleStatus_CompletedWithErrors),
		string(ProtectionRuleStatus_DeleteRequested),
		string(ProtectionRuleStatus_Draft),
		string(ProtectionRuleStatus_UpdateRequested),
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
		"deleterequested":     ProtectionRuleStatus_DeleteRequested,
		"draft":               ProtectionRuleStatus_Draft,
		"updaterequested":     ProtectionRuleStatus_UpdateRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionRuleStatus(input)
	return &out, nil
}
