package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenState string

const (
	VppTokenState_AssignedToExternalMDM VppTokenState = "assignedToExternalMDM"
	VppTokenState_Expired               VppTokenState = "expired"
	VppTokenState_Invalid               VppTokenState = "invalid"
	VppTokenState_Unknown               VppTokenState = "unknown"
	VppTokenState_Valid                 VppTokenState = "valid"
)

func PossibleValuesForVppTokenState() []string {
	return []string{
		string(VppTokenState_AssignedToExternalMDM),
		string(VppTokenState_Expired),
		string(VppTokenState_Invalid),
		string(VppTokenState_Unknown),
		string(VppTokenState_Valid),
	}
}

func (s *VppTokenState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenState(input string) (*VppTokenState, error) {
	vals := map[string]VppTokenState{
		"assignedtoexternalmdm": VppTokenState_AssignedToExternalMDM,
		"expired":               VppTokenState_Expired,
		"invalid":               VppTokenState_Invalid,
		"unknown":               VppTokenState_Unknown,
		"valid":                 VppTokenState_Valid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenState(input)
	return &out, nil
}
