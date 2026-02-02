package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityAlertSeverity string

const (
	PartnerSecuritySecurityAlertSeverity_High          PartnerSecuritySecurityAlertSeverity = "high"
	PartnerSecuritySecurityAlertSeverity_Informational PartnerSecuritySecurityAlertSeverity = "informational"
	PartnerSecuritySecurityAlertSeverity_Low           PartnerSecuritySecurityAlertSeverity = "low"
	PartnerSecuritySecurityAlertSeverity_Medium        PartnerSecuritySecurityAlertSeverity = "medium"
)

func PossibleValuesForPartnerSecuritySecurityAlertSeverity() []string {
	return []string{
		string(PartnerSecuritySecurityAlertSeverity_High),
		string(PartnerSecuritySecurityAlertSeverity_Informational),
		string(PartnerSecuritySecurityAlertSeverity_Low),
		string(PartnerSecuritySecurityAlertSeverity_Medium),
	}
}

func (s *PartnerSecuritySecurityAlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecuritySecurityAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecuritySecurityAlertSeverity(input string) (*PartnerSecuritySecurityAlertSeverity, error) {
	vals := map[string]PartnerSecuritySecurityAlertSeverity{
		"high":          PartnerSecuritySecurityAlertSeverity_High,
		"informational": PartnerSecuritySecurityAlertSeverity_Informational,
		"low":           PartnerSecuritySecurityAlertSeverity_Low,
		"medium":        PartnerSecuritySecurityAlertSeverity_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecuritySecurityAlertSeverity(input)
	return &out, nil
}
