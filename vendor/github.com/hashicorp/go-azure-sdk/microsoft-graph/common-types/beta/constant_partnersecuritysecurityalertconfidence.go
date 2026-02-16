package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityAlertConfidence string

const (
	PartnerSecuritySecurityAlertConfidence_High   PartnerSecuritySecurityAlertConfidence = "high"
	PartnerSecuritySecurityAlertConfidence_Low    PartnerSecuritySecurityAlertConfidence = "low"
	PartnerSecuritySecurityAlertConfidence_Medium PartnerSecuritySecurityAlertConfidence = "medium"
)

func PossibleValuesForPartnerSecuritySecurityAlertConfidence() []string {
	return []string{
		string(PartnerSecuritySecurityAlertConfidence_High),
		string(PartnerSecuritySecurityAlertConfidence_Low),
		string(PartnerSecuritySecurityAlertConfidence_Medium),
	}
}

func (s *PartnerSecuritySecurityAlertConfidence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecuritySecurityAlertConfidence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecuritySecurityAlertConfidence(input string) (*PartnerSecuritySecurityAlertConfidence, error) {
	vals := map[string]PartnerSecuritySecurityAlertConfidence{
		"high":   PartnerSecuritySecurityAlertConfidence_High,
		"low":    PartnerSecuritySecurityAlertConfidence_Low,
		"medium": PartnerSecuritySecurityAlertConfidence_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecuritySecurityAlertConfidence(input)
	return &out, nil
}
