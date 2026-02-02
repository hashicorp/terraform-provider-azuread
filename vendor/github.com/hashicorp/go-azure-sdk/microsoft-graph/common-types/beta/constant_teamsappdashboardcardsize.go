package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDashboardCardSize string

const (
	TeamsAppDashboardCardSize_Large  TeamsAppDashboardCardSize = "large"
	TeamsAppDashboardCardSize_Medium TeamsAppDashboardCardSize = "medium"
)

func PossibleValuesForTeamsAppDashboardCardSize() []string {
	return []string{
		string(TeamsAppDashboardCardSize_Large),
		string(TeamsAppDashboardCardSize_Medium),
	}
}

func (s *TeamsAppDashboardCardSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDashboardCardSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDashboardCardSize(input string) (*TeamsAppDashboardCardSize, error) {
	vals := map[string]TeamsAppDashboardCardSize{
		"large":  TeamsAppDashboardCardSize_Large,
		"medium": TeamsAppDashboardCardSize_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDashboardCardSize(input)
	return &out, nil
}
