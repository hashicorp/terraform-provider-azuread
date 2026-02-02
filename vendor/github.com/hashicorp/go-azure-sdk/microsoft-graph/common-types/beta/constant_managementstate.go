package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementState string

const (
	ManagementState_DeletePending  ManagementState = "deletePending"
	ManagementState_Discovered     ManagementState = "discovered"
	ManagementState_Managed        ManagementState = "managed"
	ManagementState_RetireCanceled ManagementState = "retireCanceled"
	ManagementState_RetireFailed   ManagementState = "retireFailed"
	ManagementState_RetireIssued   ManagementState = "retireIssued"
	ManagementState_RetirePending  ManagementState = "retirePending"
	ManagementState_Unhealthy      ManagementState = "unhealthy"
	ManagementState_WipeCanceled   ManagementState = "wipeCanceled"
	ManagementState_WipeFailed     ManagementState = "wipeFailed"
	ManagementState_WipeIssued     ManagementState = "wipeIssued"
	ManagementState_WipePending    ManagementState = "wipePending"
)

func PossibleValuesForManagementState() []string {
	return []string{
		string(ManagementState_DeletePending),
		string(ManagementState_Discovered),
		string(ManagementState_Managed),
		string(ManagementState_RetireCanceled),
		string(ManagementState_RetireFailed),
		string(ManagementState_RetireIssued),
		string(ManagementState_RetirePending),
		string(ManagementState_Unhealthy),
		string(ManagementState_WipeCanceled),
		string(ManagementState_WipeFailed),
		string(ManagementState_WipeIssued),
		string(ManagementState_WipePending),
	}
}

func (s *ManagementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagementState(input string) (*ManagementState, error) {
	vals := map[string]ManagementState{
		"deletepending":  ManagementState_DeletePending,
		"discovered":     ManagementState_Discovered,
		"managed":        ManagementState_Managed,
		"retirecanceled": ManagementState_RetireCanceled,
		"retirefailed":   ManagementState_RetireFailed,
		"retireissued":   ManagementState_RetireIssued,
		"retirepending":  ManagementState_RetirePending,
		"unhealthy":      ManagementState_Unhealthy,
		"wipecanceled":   ManagementState_WipeCanceled,
		"wipefailed":     ManagementState_WipeFailed,
		"wipeissued":     ManagementState_WipeIssued,
		"wipepending":    ManagementState_WipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementState(input)
	return &out, nil
}
