package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecurityComplianceStatus string

const (
	PartnerSecurityComplianceStatus_Compliant    PartnerSecurityComplianceStatus = "compliant"
	PartnerSecurityComplianceStatus_Noncomplaint PartnerSecurityComplianceStatus = "noncomplaint"
)

func PossibleValuesForPartnerSecurityComplianceStatus() []string {
	return []string{
		string(PartnerSecurityComplianceStatus_Compliant),
		string(PartnerSecurityComplianceStatus_Noncomplaint),
	}
}

func (s *PartnerSecurityComplianceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecurityComplianceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecurityComplianceStatus(input string) (*PartnerSecurityComplianceStatus, error) {
	vals := map[string]PartnerSecurityComplianceStatus{
		"compliant":    PartnerSecurityComplianceStatus_Compliant,
		"noncomplaint": PartnerSecurityComplianceStatus_Noncomplaint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecurityComplianceStatus(input)
	return &out, nil
}
