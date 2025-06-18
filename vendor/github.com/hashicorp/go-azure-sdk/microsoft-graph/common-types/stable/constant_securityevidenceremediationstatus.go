package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEvidenceRemediationStatus string

const (
	SecurityEvidenceRemediationStatus_Active              SecurityEvidenceRemediationStatus = "active"
	SecurityEvidenceRemediationStatus_Blocked             SecurityEvidenceRemediationStatus = "blocked"
	SecurityEvidenceRemediationStatus_Declined            SecurityEvidenceRemediationStatus = "declined"
	SecurityEvidenceRemediationStatus_None                SecurityEvidenceRemediationStatus = "none"
	SecurityEvidenceRemediationStatus_NotFound            SecurityEvidenceRemediationStatus = "notFound"
	SecurityEvidenceRemediationStatus_PartiallyRemediated SecurityEvidenceRemediationStatus = "partiallyRemediated"
	SecurityEvidenceRemediationStatus_PendingApproval     SecurityEvidenceRemediationStatus = "pendingApproval"
	SecurityEvidenceRemediationStatus_Prevented           SecurityEvidenceRemediationStatus = "prevented"
	SecurityEvidenceRemediationStatus_Remediated          SecurityEvidenceRemediationStatus = "remediated"
	SecurityEvidenceRemediationStatus_Running             SecurityEvidenceRemediationStatus = "running"
	SecurityEvidenceRemediationStatus_Unremediated        SecurityEvidenceRemediationStatus = "unremediated"
)

func PossibleValuesForSecurityEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityEvidenceRemediationStatus_Active),
		string(SecurityEvidenceRemediationStatus_Blocked),
		string(SecurityEvidenceRemediationStatus_Declined),
		string(SecurityEvidenceRemediationStatus_None),
		string(SecurityEvidenceRemediationStatus_NotFound),
		string(SecurityEvidenceRemediationStatus_PartiallyRemediated),
		string(SecurityEvidenceRemediationStatus_PendingApproval),
		string(SecurityEvidenceRemediationStatus_Prevented),
		string(SecurityEvidenceRemediationStatus_Remediated),
		string(SecurityEvidenceRemediationStatus_Running),
		string(SecurityEvidenceRemediationStatus_Unremediated),
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
		"active":              SecurityEvidenceRemediationStatus_Active,
		"blocked":             SecurityEvidenceRemediationStatus_Blocked,
		"declined":            SecurityEvidenceRemediationStatus_Declined,
		"none":                SecurityEvidenceRemediationStatus_None,
		"notfound":            SecurityEvidenceRemediationStatus_NotFound,
		"partiallyremediated": SecurityEvidenceRemediationStatus_PartiallyRemediated,
		"pendingapproval":     SecurityEvidenceRemediationStatus_PendingApproval,
		"prevented":           SecurityEvidenceRemediationStatus_Prevented,
		"remediated":          SecurityEvidenceRemediationStatus_Remediated,
		"running":             SecurityEvidenceRemediationStatus_Running,
		"unremediated":        SecurityEvidenceRemediationStatus_Unremediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEvidenceRemediationStatus(input)
	return &out, nil
}
