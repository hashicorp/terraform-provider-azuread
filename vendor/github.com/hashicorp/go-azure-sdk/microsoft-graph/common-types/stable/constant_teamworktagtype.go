package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkTagType string

const (
	TeamworkTagType_Standard TeamworkTagType = "standard"
)

func PossibleValuesForTeamworkTagType() []string {
	return []string{
		string(TeamworkTagType_Standard),
	}
}

func (s *TeamworkTagType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkTagType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkTagType(input string) (*TeamworkTagType, error) {
	vals := map[string]TeamworkTagType{
		"standard": TeamworkTagType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkTagType(input)
	return &out, nil
}
