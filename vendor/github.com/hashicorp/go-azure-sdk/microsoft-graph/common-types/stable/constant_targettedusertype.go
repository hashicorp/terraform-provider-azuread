package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargettedUserType string

const (
	TargettedUserType_AllUsers    TargettedUserType = "allUsers"
	TargettedUserType_Clicked     TargettedUserType = "clicked"
	TargettedUserType_Compromised TargettedUserType = "compromised"
	TargettedUserType_Unknown     TargettedUserType = "unknown"
)

func PossibleValuesForTargettedUserType() []string {
	return []string{
		string(TargettedUserType_AllUsers),
		string(TargettedUserType_Clicked),
		string(TargettedUserType_Compromised),
		string(TargettedUserType_Unknown),
	}
}

func (s *TargettedUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargettedUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargettedUserType(input string) (*TargettedUserType, error) {
	vals := map[string]TargettedUserType{
		"allusers":    TargettedUserType_AllUsers,
		"clicked":     TargettedUserType_Clicked,
		"compromised": TargettedUserType_Compromised,
		"unknown":     TargettedUserType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargettedUserType(input)
	return &out, nil
}
