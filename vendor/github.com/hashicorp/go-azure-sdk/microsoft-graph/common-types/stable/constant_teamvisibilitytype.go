package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamVisibilityType string

const (
	TeamVisibilityType_HiddenMembership TeamVisibilityType = "hiddenMembership"
	TeamVisibilityType_Private          TeamVisibilityType = "private"
	TeamVisibilityType_Public           TeamVisibilityType = "public"
)

func PossibleValuesForTeamVisibilityType() []string {
	return []string{
		string(TeamVisibilityType_HiddenMembership),
		string(TeamVisibilityType_Private),
		string(TeamVisibilityType_Public),
	}
}

func (s *TeamVisibilityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamVisibilityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamVisibilityType(input string) (*TeamVisibilityType, error) {
	vals := map[string]TeamVisibilityType{
		"hiddenmembership": TeamVisibilityType_HiddenMembership,
		"private":          TeamVisibilityType_Private,
		"public":           TeamVisibilityType_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamVisibilityType(input)
	return &out, nil
}
