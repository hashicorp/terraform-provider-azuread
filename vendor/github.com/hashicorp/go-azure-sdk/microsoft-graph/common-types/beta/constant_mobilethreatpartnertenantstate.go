package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileThreatPartnerTenantState string

const (
	MobileThreatPartnerTenantState_Available    MobileThreatPartnerTenantState = "available"
	MobileThreatPartnerTenantState_Enabled      MobileThreatPartnerTenantState = "enabled"
	MobileThreatPartnerTenantState_Error        MobileThreatPartnerTenantState = "error"
	MobileThreatPartnerTenantState_NotSetUp     MobileThreatPartnerTenantState = "notSetUp"
	MobileThreatPartnerTenantState_Unavailable  MobileThreatPartnerTenantState = "unavailable"
	MobileThreatPartnerTenantState_Unresponsive MobileThreatPartnerTenantState = "unresponsive"
)

func PossibleValuesForMobileThreatPartnerTenantState() []string {
	return []string{
		string(MobileThreatPartnerTenantState_Available),
		string(MobileThreatPartnerTenantState_Enabled),
		string(MobileThreatPartnerTenantState_Error),
		string(MobileThreatPartnerTenantState_NotSetUp),
		string(MobileThreatPartnerTenantState_Unavailable),
		string(MobileThreatPartnerTenantState_Unresponsive),
	}
}

func (s *MobileThreatPartnerTenantState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileThreatPartnerTenantState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileThreatPartnerTenantState(input string) (*MobileThreatPartnerTenantState, error) {
	vals := map[string]MobileThreatPartnerTenantState{
		"available":    MobileThreatPartnerTenantState_Available,
		"enabled":      MobileThreatPartnerTenantState_Enabled,
		"error":        MobileThreatPartnerTenantState_Error,
		"notsetup":     MobileThreatPartnerTenantState_NotSetUp,
		"unavailable":  MobileThreatPartnerTenantState_Unavailable,
		"unresponsive": MobileThreatPartnerTenantState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileThreatPartnerTenantState(input)
	return &out, nil
}
