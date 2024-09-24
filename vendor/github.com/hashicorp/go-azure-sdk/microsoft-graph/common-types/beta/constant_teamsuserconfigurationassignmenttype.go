package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsUserConfigurationAssignmentType string

const (
	TeamsUserConfigurationAssignmentType_Direct TeamsUserConfigurationAssignmentType = "direct"
	TeamsUserConfigurationAssignmentType_Group  TeamsUserConfigurationAssignmentType = "group"
)

func PossibleValuesForTeamsUserConfigurationAssignmentType() []string {
	return []string{
		string(TeamsUserConfigurationAssignmentType_Direct),
		string(TeamsUserConfigurationAssignmentType_Group),
	}
}

func (s *TeamsUserConfigurationAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsUserConfigurationAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsUserConfigurationAssignmentType(input string) (*TeamsUserConfigurationAssignmentType, error) {
	vals := map[string]TeamsUserConfigurationAssignmentType{
		"direct": TeamsUserConfigurationAssignmentType_Direct,
		"group":  TeamsUserConfigurationAssignmentType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsUserConfigurationAssignmentType(input)
	return &out, nil
}
