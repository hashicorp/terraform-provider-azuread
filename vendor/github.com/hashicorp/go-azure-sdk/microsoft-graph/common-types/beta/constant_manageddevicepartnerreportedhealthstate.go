package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDevicePartnerReportedHealthState string

const (
	ManagedDevicePartnerReportedHealthState_Activated      ManagedDevicePartnerReportedHealthState = "activated"
	ManagedDevicePartnerReportedHealthState_Compromised    ManagedDevicePartnerReportedHealthState = "compromised"
	ManagedDevicePartnerReportedHealthState_Deactivated    ManagedDevicePartnerReportedHealthState = "deactivated"
	ManagedDevicePartnerReportedHealthState_HighSeverity   ManagedDevicePartnerReportedHealthState = "highSeverity"
	ManagedDevicePartnerReportedHealthState_LowSeverity    ManagedDevicePartnerReportedHealthState = "lowSeverity"
	ManagedDevicePartnerReportedHealthState_MediumSeverity ManagedDevicePartnerReportedHealthState = "mediumSeverity"
	ManagedDevicePartnerReportedHealthState_Misconfigured  ManagedDevicePartnerReportedHealthState = "misconfigured"
	ManagedDevicePartnerReportedHealthState_Secured        ManagedDevicePartnerReportedHealthState = "secured"
	ManagedDevicePartnerReportedHealthState_Unknown        ManagedDevicePartnerReportedHealthState = "unknown"
	ManagedDevicePartnerReportedHealthState_Unresponsive   ManagedDevicePartnerReportedHealthState = "unresponsive"
)

func PossibleValuesForManagedDevicePartnerReportedHealthState() []string {
	return []string{
		string(ManagedDevicePartnerReportedHealthState_Activated),
		string(ManagedDevicePartnerReportedHealthState_Compromised),
		string(ManagedDevicePartnerReportedHealthState_Deactivated),
		string(ManagedDevicePartnerReportedHealthState_HighSeverity),
		string(ManagedDevicePartnerReportedHealthState_LowSeverity),
		string(ManagedDevicePartnerReportedHealthState_MediumSeverity),
		string(ManagedDevicePartnerReportedHealthState_Misconfigured),
		string(ManagedDevicePartnerReportedHealthState_Secured),
		string(ManagedDevicePartnerReportedHealthState_Unknown),
		string(ManagedDevicePartnerReportedHealthState_Unresponsive),
	}
}

func (s *ManagedDevicePartnerReportedHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDevicePartnerReportedHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDevicePartnerReportedHealthState(input string) (*ManagedDevicePartnerReportedHealthState, error) {
	vals := map[string]ManagedDevicePartnerReportedHealthState{
		"activated":      ManagedDevicePartnerReportedHealthState_Activated,
		"compromised":    ManagedDevicePartnerReportedHealthState_Compromised,
		"deactivated":    ManagedDevicePartnerReportedHealthState_Deactivated,
		"highseverity":   ManagedDevicePartnerReportedHealthState_HighSeverity,
		"lowseverity":    ManagedDevicePartnerReportedHealthState_LowSeverity,
		"mediumseverity": ManagedDevicePartnerReportedHealthState_MediumSeverity,
		"misconfigured":  ManagedDevicePartnerReportedHealthState_Misconfigured,
		"secured":        ManagedDevicePartnerReportedHealthState_Secured,
		"unknown":        ManagedDevicePartnerReportedHealthState_Unknown,
		"unresponsive":   ManagedDevicePartnerReportedHealthState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDevicePartnerReportedHealthState(input)
	return &out, nil
}
