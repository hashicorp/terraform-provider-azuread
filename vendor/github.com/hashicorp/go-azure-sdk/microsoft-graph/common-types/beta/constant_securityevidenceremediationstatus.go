package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEvidenceRemediationStatus string

const (
	SecurityEvidenceRemediationStatus_Blocked    SecurityEvidenceRemediationStatus = "blocked"
	SecurityEvidenceRemediationStatus_None       SecurityEvidenceRemediationStatus = "none"
	SecurityEvidenceRemediationStatus_NotFound   SecurityEvidenceRemediationStatus = "notFound"
	SecurityEvidenceRemediationStatus_Prevented  SecurityEvidenceRemediationStatus = "prevented"
	SecurityEvidenceRemediationStatus_Remediated SecurityEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityEvidenceRemediationStatus_Blocked),
		string(SecurityEvidenceRemediationStatus_None),
		string(SecurityEvidenceRemediationStatus_NotFound),
		string(SecurityEvidenceRemediationStatus_Prevented),
		string(SecurityEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEvidenceRemediationStatus(input string) (*SecurityEvidenceRemediationStatus, error) {
	vals := map[string]SecurityEvidenceRemediationStatus{
		"blocked":    SecurityEvidenceRemediationStatus_Blocked,
		"none":       SecurityEvidenceRemediationStatus_None,
		"notfound":   SecurityEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityEvidenceRemediationStatus_Prevented,
		"remediated": SecurityEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEvidenceRemediationStatus(input)
	return &out, nil
}
