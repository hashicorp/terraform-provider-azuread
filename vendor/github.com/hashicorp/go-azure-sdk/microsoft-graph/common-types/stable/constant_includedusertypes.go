package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludedUserTypes string

const (
	IncludedUserTypes_All    IncludedUserTypes = "all"
	IncludedUserTypes_Guest  IncludedUserTypes = "guest"
	IncludedUserTypes_Member IncludedUserTypes = "member"
)

func PossibleValuesForIncludedUserTypes() []string {
	return []string{
		string(IncludedUserTypes_All),
		string(IncludedUserTypes_Guest),
		string(IncludedUserTypes_Member),
	}
}

func (s *IncludedUserTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncludedUserTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncludedUserTypes(input string) (*IncludedUserTypes, error) {
	vals := map[string]IncludedUserTypes{
		"all":    IncludedUserTypes_All,
		"guest":  IncludedUserTypes_Guest,
		"member": IncludedUserTypes_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncludedUserTypes(input)
	return &out, nil
}
