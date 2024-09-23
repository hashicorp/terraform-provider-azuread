package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAlertStatus string

const (
	ManagedTenantsAlertStatus_Dismissed  ManagedTenantsAlertStatus = "dismissed"
	ManagedTenantsAlertStatus_InProgress ManagedTenantsAlertStatus = "inProgress"
	ManagedTenantsAlertStatus_NewAlert   ManagedTenantsAlertStatus = "newAlert"
	ManagedTenantsAlertStatus_Resolved   ManagedTenantsAlertStatus = "resolved"
	ManagedTenantsAlertStatus_Unknown    ManagedTenantsAlertStatus = "unknown"
)

func PossibleValuesForManagedTenantsAlertStatus() []string {
	return []string{
		string(ManagedTenantsAlertStatus_Dismissed),
		string(ManagedTenantsAlertStatus_InProgress),
		string(ManagedTenantsAlertStatus_NewAlert),
		string(ManagedTenantsAlertStatus_Resolved),
		string(ManagedTenantsAlertStatus_Unknown),
	}
}

func (s *ManagedTenantsAlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsAlertStatus(input string) (*ManagedTenantsAlertStatus, error) {
	vals := map[string]ManagedTenantsAlertStatus{
		"dismissed":  ManagedTenantsAlertStatus_Dismissed,
		"inprogress": ManagedTenantsAlertStatus_InProgress,
		"newalert":   ManagedTenantsAlertStatus_NewAlert,
		"resolved":   ManagedTenantsAlertStatus_Resolved,
		"unknown":    ManagedTenantsAlertStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsAlertStatus(input)
	return &out, nil
}
