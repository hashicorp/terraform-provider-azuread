package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityAlertResolvedReason string

const (
	PartnerSecuritySecurityAlertResolvedReason_Fraud      PartnerSecuritySecurityAlertResolvedReason = "fraud"
	PartnerSecuritySecurityAlertResolvedReason_Ignore     PartnerSecuritySecurityAlertResolvedReason = "ignore"
	PartnerSecuritySecurityAlertResolvedReason_Legitimate PartnerSecuritySecurityAlertResolvedReason = "legitimate"
)

func PossibleValuesForPartnerSecuritySecurityAlertResolvedReason() []string {
	return []string{
		string(PartnerSecuritySecurityAlertResolvedReason_Fraud),
		string(PartnerSecuritySecurityAlertResolvedReason_Ignore),
		string(PartnerSecuritySecurityAlertResolvedReason_Legitimate),
	}
}

func (s *PartnerSecuritySecurityAlertResolvedReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecuritySecurityAlertResolvedReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecuritySecurityAlertResolvedReason(input string) (*PartnerSecuritySecurityAlertResolvedReason, error) {
	vals := map[string]PartnerSecuritySecurityAlertResolvedReason{
		"fraud":      PartnerSecuritySecurityAlertResolvedReason_Fraud,
		"ignore":     PartnerSecuritySecurityAlertResolvedReason_Ignore,
		"legitimate": PartnerSecuritySecurityAlertResolvedReason_Legitimate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecuritySecurityAlertResolvedReason(input)
	return &out, nil
}
