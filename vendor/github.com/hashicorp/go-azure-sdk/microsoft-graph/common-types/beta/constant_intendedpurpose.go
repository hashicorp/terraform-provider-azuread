package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntendedPurpose string

const (
	IntendedPurpose_SmimeEncryption IntendedPurpose = "smimeEncryption"
	IntendedPurpose_SmimeSigning    IntendedPurpose = "smimeSigning"
	IntendedPurpose_Unassigned      IntendedPurpose = "unassigned"
	IntendedPurpose_Vpn             IntendedPurpose = "vpn"
	IntendedPurpose_Wifi            IntendedPurpose = "wifi"
)

func PossibleValuesForIntendedPurpose() []string {
	return []string{
		string(IntendedPurpose_SmimeEncryption),
		string(IntendedPurpose_SmimeSigning),
		string(IntendedPurpose_Unassigned),
		string(IntendedPurpose_Vpn),
		string(IntendedPurpose_Wifi),
	}
}

func (s *IntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntendedPurpose(input string) (*IntendedPurpose, error) {
	vals := map[string]IntendedPurpose{
		"smimeencryption": IntendedPurpose_SmimeEncryption,
		"smimesigning":    IntendedPurpose_SmimeSigning,
		"unassigned":      IntendedPurpose_Unassigned,
		"vpn":             IntendedPurpose_Vpn,
		"wifi":            IntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntendedPurpose(input)
	return &out, nil
}
