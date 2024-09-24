package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateAudience string

const (
	TeamTemplateAudience_Organization TeamTemplateAudience = "organization"
	TeamTemplateAudience_Public       TeamTemplateAudience = "public"
	TeamTemplateAudience_User         TeamTemplateAudience = "user"
)

func PossibleValuesForTeamTemplateAudience() []string {
	return []string{
		string(TeamTemplateAudience_Organization),
		string(TeamTemplateAudience_Public),
		string(TeamTemplateAudience_User),
	}
}

func (s *TeamTemplateAudience) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamTemplateAudience(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamTemplateAudience(input string) (*TeamTemplateAudience, error) {
	vals := map[string]TeamTemplateAudience{
		"organization": TeamTemplateAudience_Organization,
		"public":       TeamTemplateAudience_Public,
		"user":         TeamTemplateAudience_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamTemplateAudience(input)
	return &out, nil
}
