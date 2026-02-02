package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedLobbyAdmitterRoles string

const (
	AllowedLobbyAdmitterRoles_OrganizerAndCoOrganizers              AllowedLobbyAdmitterRoles = "organizerAndCoOrganizers"
	AllowedLobbyAdmitterRoles_OrganizerAndCoOrganizersAndPresenters AllowedLobbyAdmitterRoles = "organizerAndCoOrganizersAndPresenters"
)

func PossibleValuesForAllowedLobbyAdmitterRoles() []string {
	return []string{
		string(AllowedLobbyAdmitterRoles_OrganizerAndCoOrganizers),
		string(AllowedLobbyAdmitterRoles_OrganizerAndCoOrganizersAndPresenters),
	}
}

func (s *AllowedLobbyAdmitterRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedLobbyAdmitterRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedLobbyAdmitterRoles(input string) (*AllowedLobbyAdmitterRoles, error) {
	vals := map[string]AllowedLobbyAdmitterRoles{
		"organizerandcoorganizers":              AllowedLobbyAdmitterRoles_OrganizerAndCoOrganizers,
		"organizerandcoorganizersandpresenters": AllowedLobbyAdmitterRoles_OrganizerAndCoOrganizersAndPresenters,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedLobbyAdmitterRoles(input)
	return &out, nil
}
