package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceStatus string

const (
	ComplianceStatus_Compliant     ComplianceStatus = "compliant"
	ComplianceStatus_Conflict      ComplianceStatus = "conflict"
	ComplianceStatus_Error         ComplianceStatus = "error"
	ComplianceStatus_NonCompliant  ComplianceStatus = "nonCompliant"
	ComplianceStatus_NotApplicable ComplianceStatus = "notApplicable"
	ComplianceStatus_NotAssigned   ComplianceStatus = "notAssigned"
	ComplianceStatus_Remediated    ComplianceStatus = "remediated"
	ComplianceStatus_Unknown       ComplianceStatus = "unknown"
)

func PossibleValuesForComplianceStatus() []string {
	return []string{
		string(ComplianceStatus_Compliant),
		string(ComplianceStatus_Conflict),
		string(ComplianceStatus_Error),
		string(ComplianceStatus_NonCompliant),
		string(ComplianceStatus_NotApplicable),
		string(ComplianceStatus_NotAssigned),
		string(ComplianceStatus_Remediated),
		string(ComplianceStatus_Unknown),
	}
}

func (s *ComplianceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComplianceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComplianceStatus(input string) (*ComplianceStatus, error) {
	vals := map[string]ComplianceStatus{
		"compliant":     ComplianceStatus_Compliant,
		"conflict":      ComplianceStatus_Conflict,
		"error":         ComplianceStatus_Error,
		"noncompliant":  ComplianceStatus_NonCompliant,
		"notapplicable": ComplianceStatus_NotApplicable,
		"notassigned":   ComplianceStatus_NotAssigned,
		"remediated":    ComplianceStatus_Remediated,
		"unknown":       ComplianceStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceStatus(input)
	return &out, nil
}
