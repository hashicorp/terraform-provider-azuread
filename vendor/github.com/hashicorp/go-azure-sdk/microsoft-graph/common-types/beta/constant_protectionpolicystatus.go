package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionPolicyStatus string

const (
	ProtectionPolicyStatus_Active           ProtectionPolicyStatus = "active"
	ProtectionPolicyStatus_ActiveWithErrors ProtectionPolicyStatus = "activeWithErrors"
	ProtectionPolicyStatus_Inactive         ProtectionPolicyStatus = "inactive"
	ProtectionPolicyStatus_Updating         ProtectionPolicyStatus = "updating"
)

func PossibleValuesForProtectionPolicyStatus() []string {
	return []string{
		string(ProtectionPolicyStatus_Active),
		string(ProtectionPolicyStatus_ActiveWithErrors),
		string(ProtectionPolicyStatus_Inactive),
		string(ProtectionPolicyStatus_Updating),
	}
}

func (s *ProtectionPolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionPolicyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionPolicyStatus(input string) (*ProtectionPolicyStatus, error) {
	vals := map[string]ProtectionPolicyStatus{
		"active":           ProtectionPolicyStatus_Active,
		"activewitherrors": ProtectionPolicyStatus_ActiveWithErrors,
		"inactive":         ProtectionPolicyStatus_Inactive,
		"updating":         ProtectionPolicyStatus_Updating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionPolicyStatus(input)
	return &out, nil
}
