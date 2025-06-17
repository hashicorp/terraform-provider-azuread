package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAdministrationAssignmentCategory string

const (
	TeamsAdministrationAssignmentCategory_Alternate TeamsAdministrationAssignmentCategory = "alternate"
	TeamsAdministrationAssignmentCategory_Primary   TeamsAdministrationAssignmentCategory = "primary"
	TeamsAdministrationAssignmentCategory_Private   TeamsAdministrationAssignmentCategory = "private"
)

func PossibleValuesForTeamsAdministrationAssignmentCategory() []string {
	return []string{
		string(TeamsAdministrationAssignmentCategory_Alternate),
		string(TeamsAdministrationAssignmentCategory_Primary),
		string(TeamsAdministrationAssignmentCategory_Private),
	}
}

func (s *TeamsAdministrationAssignmentCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAdministrationAssignmentCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAdministrationAssignmentCategory(input string) (*TeamsAdministrationAssignmentCategory, error) {
	vals := map[string]TeamsAdministrationAssignmentCategory{
		"alternate": TeamsAdministrationAssignmentCategory_Alternate,
		"primary":   TeamsAdministrationAssignmentCategory_Primary,
		"private":   TeamsAdministrationAssignmentCategory_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAdministrationAssignmentCategory(input)
	return &out, nil
}
