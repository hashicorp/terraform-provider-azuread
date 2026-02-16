package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPfxIntendedPurpose string

const (
	UserPfxIntendedPurpose_SmimeEncryption UserPfxIntendedPurpose = "smimeEncryption"
	UserPfxIntendedPurpose_SmimeSigning    UserPfxIntendedPurpose = "smimeSigning"
	UserPfxIntendedPurpose_Unassigned      UserPfxIntendedPurpose = "unassigned"
	UserPfxIntendedPurpose_Vpn             UserPfxIntendedPurpose = "vpn"
	UserPfxIntendedPurpose_Wifi            UserPfxIntendedPurpose = "wifi"
)

func PossibleValuesForUserPfxIntendedPurpose() []string {
	return []string{
		string(UserPfxIntendedPurpose_SmimeEncryption),
		string(UserPfxIntendedPurpose_SmimeSigning),
		string(UserPfxIntendedPurpose_Unassigned),
		string(UserPfxIntendedPurpose_Vpn),
		string(UserPfxIntendedPurpose_Wifi),
	}
}

func (s *UserPfxIntendedPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserPfxIntendedPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserPfxIntendedPurpose(input string) (*UserPfxIntendedPurpose, error) {
	vals := map[string]UserPfxIntendedPurpose{
		"smimeencryption": UserPfxIntendedPurpose_SmimeEncryption,
		"smimesigning":    UserPfxIntendedPurpose_SmimeSigning,
		"unassigned":      UserPfxIntendedPurpose_Unassigned,
		"vpn":             UserPfxIntendedPurpose_Vpn,
		"wifi":            UserPfxIntendedPurpose_Wifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPfxIntendedPurpose(input)
	return &out, nil
}
