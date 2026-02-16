package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionStatus string

const (
	ManagedTenantsManagementActionStatus_Completed                          ManagedTenantsManagementActionStatus = "completed"
	ManagedTenantsManagementActionStatus_Error                              ManagedTenantsManagementActionStatus = "error"
	ManagedTenantsManagementActionStatus_InProgress                         ManagedTenantsManagementActionStatus = "inProgress"
	ManagedTenantsManagementActionStatus_Planned                            ManagedTenantsManagementActionStatus = "planned"
	ManagedTenantsManagementActionStatus_ResolvedBy3rdParty                 ManagedTenantsManagementActionStatus = "resolvedBy3rdParty"
	ManagedTenantsManagementActionStatus_ResolvedThroughAlternateMitigation ManagedTenantsManagementActionStatus = "resolvedThroughAlternateMitigation"
	ManagedTenantsManagementActionStatus_RiskAccepted                       ManagedTenantsManagementActionStatus = "riskAccepted"
	ManagedTenantsManagementActionStatus_TimeOut                            ManagedTenantsManagementActionStatus = "timeOut"
	ManagedTenantsManagementActionStatus_ToAddress                          ManagedTenantsManagementActionStatus = "toAddress"
)

func PossibleValuesForManagedTenantsManagementActionStatus() []string {
	return []string{
		string(ManagedTenantsManagementActionStatus_Completed),
		string(ManagedTenantsManagementActionStatus_Error),
		string(ManagedTenantsManagementActionStatus_InProgress),
		string(ManagedTenantsManagementActionStatus_Planned),
		string(ManagedTenantsManagementActionStatus_ResolvedBy3rdParty),
		string(ManagedTenantsManagementActionStatus_ResolvedThroughAlternateMitigation),
		string(ManagedTenantsManagementActionStatus_RiskAccepted),
		string(ManagedTenantsManagementActionStatus_TimeOut),
		string(ManagedTenantsManagementActionStatus_ToAddress),
	}
}

func (s *ManagedTenantsManagementActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementActionStatus(input string) (*ManagedTenantsManagementActionStatus, error) {
	vals := map[string]ManagedTenantsManagementActionStatus{
		"completed":                          ManagedTenantsManagementActionStatus_Completed,
		"error":                              ManagedTenantsManagementActionStatus_Error,
		"inprogress":                         ManagedTenantsManagementActionStatus_InProgress,
		"planned":                            ManagedTenantsManagementActionStatus_Planned,
		"resolvedby3rdparty":                 ManagedTenantsManagementActionStatus_ResolvedBy3rdParty,
		"resolvedthroughalternatemitigation": ManagedTenantsManagementActionStatus_ResolvedThroughAlternateMitigation,
		"riskaccepted":                       ManagedTenantsManagementActionStatus_RiskAccepted,
		"timeout":                            ManagedTenantsManagementActionStatus_TimeOut,
		"toaddress":                          ManagedTenantsManagementActionStatus_ToAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementActionStatus(input)
	return &out, nil
}
