package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppInstallationScopes string

const (
	TeamsAppInstallationScopes_GroupChat TeamsAppInstallationScopes = "groupChat"
	TeamsAppInstallationScopes_Personal  TeamsAppInstallationScopes = "personal"
	TeamsAppInstallationScopes_Team      TeamsAppInstallationScopes = "team"
)

func PossibleValuesForTeamsAppInstallationScopes() []string {
	return []string{
		string(TeamsAppInstallationScopes_GroupChat),
		string(TeamsAppInstallationScopes_Personal),
		string(TeamsAppInstallationScopes_Team),
	}
}

func (s *TeamsAppInstallationScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppInstallationScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppInstallationScopes(input string) (*TeamsAppInstallationScopes, error) {
	vals := map[string]TeamsAppInstallationScopes{
		"groupchat": TeamsAppInstallationScopes_GroupChat,
		"personal":  TeamsAppInstallationScopes_Personal,
		"team":      TeamsAppInstallationScopes_Team,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppInstallationScopes(input)
	return &out, nil
}
