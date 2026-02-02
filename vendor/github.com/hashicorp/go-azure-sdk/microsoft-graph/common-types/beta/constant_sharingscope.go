package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingScope string

const (
	SharingScope_Anonymous      SharingScope = "anonymous"
	SharingScope_Anyone         SharingScope = "anyone"
	SharingScope_Organization   SharingScope = "organization"
	SharingScope_SpecificPeople SharingScope = "specificPeople"
	SharingScope_Users          SharingScope = "users"
)

func PossibleValuesForSharingScope() []string {
	return []string{
		string(SharingScope_Anonymous),
		string(SharingScope_Anyone),
		string(SharingScope_Organization),
		string(SharingScope_SpecificPeople),
		string(SharingScope_Users),
	}
}

func (s *SharingScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharingScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharingScope(input string) (*SharingScope, error) {
	vals := map[string]SharingScope{
		"anonymous":      SharingScope_Anonymous,
		"anyone":         SharingScope_Anyone,
		"organization":   SharingScope_Organization,
		"specificpeople": SharingScope_SpecificPeople,
		"users":          SharingScope_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharingScope(input)
	return &out, nil
}
