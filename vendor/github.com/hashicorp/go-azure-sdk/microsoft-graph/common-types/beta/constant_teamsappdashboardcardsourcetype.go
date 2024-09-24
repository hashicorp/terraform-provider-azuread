package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDashboardCardSourceType string

const (
	TeamsAppDashboardCardSourceType_Bot TeamsAppDashboardCardSourceType = "bot"
)

func PossibleValuesForTeamsAppDashboardCardSourceType() []string {
	return []string{
		string(TeamsAppDashboardCardSourceType_Bot),
	}
}

func (s *TeamsAppDashboardCardSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppDashboardCardSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppDashboardCardSourceType(input string) (*TeamsAppDashboardCardSourceType, error) {
	vals := map[string]TeamsAppDashboardCardSourceType{
		"bot": TeamsAppDashboardCardSourceType_Bot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDashboardCardSourceType(input)
	return &out, nil
}
