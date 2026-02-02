package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityDomain string

const (
	ActivityDomain_Personal     ActivityDomain = "personal"
	ActivityDomain_Unknown      ActivityDomain = "unknown"
	ActivityDomain_Unrestricted ActivityDomain = "unrestricted"
	ActivityDomain_Work         ActivityDomain = "work"
)

func PossibleValuesForActivityDomain() []string {
	return []string{
		string(ActivityDomain_Personal),
		string(ActivityDomain_Unknown),
		string(ActivityDomain_Unrestricted),
		string(ActivityDomain_Work),
	}
}

func (s *ActivityDomain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivityDomain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivityDomain(input string) (*ActivityDomain, error) {
	vals := map[string]ActivityDomain{
		"personal":     ActivityDomain_Personal,
		"unknown":      ActivityDomain_Unknown,
		"unrestricted": ActivityDomain_Unrestricted,
		"work":         ActivityDomain_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityDomain(input)
	return &out, nil
}
