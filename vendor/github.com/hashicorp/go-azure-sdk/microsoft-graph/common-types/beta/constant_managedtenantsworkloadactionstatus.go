package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionStatus string

const (
	ManagedTenantsWorkloadActionStatus_Completed  ManagedTenantsWorkloadActionStatus = "completed"
	ManagedTenantsWorkloadActionStatus_Error      ManagedTenantsWorkloadActionStatus = "error"
	ManagedTenantsWorkloadActionStatus_InProgress ManagedTenantsWorkloadActionStatus = "inProgress"
	ManagedTenantsWorkloadActionStatus_TimeOut    ManagedTenantsWorkloadActionStatus = "timeOut"
	ManagedTenantsWorkloadActionStatus_ToAddress  ManagedTenantsWorkloadActionStatus = "toAddress"
)

func PossibleValuesForManagedTenantsWorkloadActionStatus() []string {
	return []string{
		string(ManagedTenantsWorkloadActionStatus_Completed),
		string(ManagedTenantsWorkloadActionStatus_Error),
		string(ManagedTenantsWorkloadActionStatus_InProgress),
		string(ManagedTenantsWorkloadActionStatus_TimeOut),
		string(ManagedTenantsWorkloadActionStatus_ToAddress),
	}
}

func (s *ManagedTenantsWorkloadActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsWorkloadActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsWorkloadActionStatus(input string) (*ManagedTenantsWorkloadActionStatus, error) {
	vals := map[string]ManagedTenantsWorkloadActionStatus{
		"completed":  ManagedTenantsWorkloadActionStatus_Completed,
		"error":      ManagedTenantsWorkloadActionStatus_Error,
		"inprogress": ManagedTenantsWorkloadActionStatus_InProgress,
		"timeout":    ManagedTenantsWorkloadActionStatus_TimeOut,
		"toaddress":  ManagedTenantsWorkloadActionStatus_ToAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsWorkloadActionStatus(input)
	return &out, nil
}
