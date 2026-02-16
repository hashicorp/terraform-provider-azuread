package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionState string

const (
	ProvisionState_NotProvisioned         ProvisionState = "notProvisioned"
	ProvisionState_ProvisioningCompleted  ProvisionState = "provisioningCompleted"
	ProvisionState_ProvisioningFailed     ProvisionState = "provisioningFailed"
	ProvisionState_ProvisioningInProgress ProvisionState = "provisioningInProgress"
)

func PossibleValuesForProvisionState() []string {
	return []string{
		string(ProvisionState_NotProvisioned),
		string(ProvisionState_ProvisioningCompleted),
		string(ProvisionState_ProvisioningFailed),
		string(ProvisionState_ProvisioningInProgress),
	}
}

func (s *ProvisionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisionState(input string) (*ProvisionState, error) {
	vals := map[string]ProvisionState{
		"notprovisioned":         ProvisionState_NotProvisioned,
		"provisioningcompleted":  ProvisionState_ProvisioningCompleted,
		"provisioningfailed":     ProvisionState_ProvisioningFailed,
		"provisioninginprogress": ProvisionState_ProvisioningInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisionState(input)
	return &out, nil
}
