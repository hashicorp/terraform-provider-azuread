package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAdministrationAssignmentType string

const (
	TeamsAdministrationAssignmentType_Direct TeamsAdministrationAssignmentType = "direct"
	TeamsAdministrationAssignmentType_Group  TeamsAdministrationAssignmentType = "group"
)

func PossibleValuesForTeamsAdministrationAssignmentType() []string {
	return []string{
		string(TeamsAdministrationAssignmentType_Direct),
		string(TeamsAdministrationAssignmentType_Group),
	}
}

func (s *TeamsAdministrationAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAdministrationAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAdministrationAssignmentType(input string) (*TeamsAdministrationAssignmentType, error) {
	vals := map[string]TeamsAdministrationAssignmentType{
		"direct": TeamsAdministrationAssignmentType_Direct,
		"group":  TeamsAdministrationAssignmentType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAdministrationAssignmentType(input)
	return &out, nil
}
