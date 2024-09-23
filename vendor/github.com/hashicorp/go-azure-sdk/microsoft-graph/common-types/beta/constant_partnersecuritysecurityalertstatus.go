package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityAlertStatus string

const (
	PartnerSecuritySecurityAlertStatus_Active        PartnerSecuritySecurityAlertStatus = "active"
	PartnerSecuritySecurityAlertStatus_Investigating PartnerSecuritySecurityAlertStatus = "investigating"
	PartnerSecuritySecurityAlertStatus_Resolved      PartnerSecuritySecurityAlertStatus = "resolved"
)

func PossibleValuesForPartnerSecuritySecurityAlertStatus() []string {
	return []string{
		string(PartnerSecuritySecurityAlertStatus_Active),
		string(PartnerSecuritySecurityAlertStatus_Investigating),
		string(PartnerSecuritySecurityAlertStatus_Resolved),
	}
}

func (s *PartnerSecuritySecurityAlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecuritySecurityAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecuritySecurityAlertStatus(input string) (*PartnerSecuritySecurityAlertStatus, error) {
	vals := map[string]PartnerSecuritySecurityAlertStatus{
		"active":        PartnerSecuritySecurityAlertStatus_Active,
		"investigating": PartnerSecuritySecurityAlertStatus_Investigating,
		"resolved":      PartnerSecuritySecurityAlertStatus_Resolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecuritySecurityAlertStatus(input)
	return &out, nil
}
