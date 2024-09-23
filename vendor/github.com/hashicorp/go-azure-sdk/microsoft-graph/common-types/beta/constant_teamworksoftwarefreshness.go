package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSoftwareFreshness string

const (
	TeamworkSoftwareFreshness_Latest          TeamworkSoftwareFreshness = "latest"
	TeamworkSoftwareFreshness_Unknown         TeamworkSoftwareFreshness = "unknown"
	TeamworkSoftwareFreshness_UpdateAvailable TeamworkSoftwareFreshness = "updateAvailable"
)

func PossibleValuesForTeamworkSoftwareFreshness() []string {
	return []string{
		string(TeamworkSoftwareFreshness_Latest),
		string(TeamworkSoftwareFreshness_Unknown),
		string(TeamworkSoftwareFreshness_UpdateAvailable),
	}
}

func (s *TeamworkSoftwareFreshness) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkSoftwareFreshness(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkSoftwareFreshness(input string) (*TeamworkSoftwareFreshness, error) {
	vals := map[string]TeamworkSoftwareFreshness{
		"latest":          TeamworkSoftwareFreshness_Latest,
		"unknown":         TeamworkSoftwareFreshness_Unknown,
		"updateavailable": TeamworkSoftwareFreshness_UpdateAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkSoftwareFreshness(input)
	return &out, nil
}
