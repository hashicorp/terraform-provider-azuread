package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LobbyBypassScope string

const (
	LobbyBypassScope_Everyone                    LobbyBypassScope = "everyone"
	LobbyBypassScope_Invited                     LobbyBypassScope = "invited"
	LobbyBypassScope_Organization                LobbyBypassScope = "organization"
	LobbyBypassScope_OrganizationAndFederated    LobbyBypassScope = "organizationAndFederated"
	LobbyBypassScope_OrganizationExcludingGuests LobbyBypassScope = "organizationExcludingGuests"
	LobbyBypassScope_Organizer                   LobbyBypassScope = "organizer"
)

func PossibleValuesForLobbyBypassScope() []string {
	return []string{
		string(LobbyBypassScope_Everyone),
		string(LobbyBypassScope_Invited),
		string(LobbyBypassScope_Organization),
		string(LobbyBypassScope_OrganizationAndFederated),
		string(LobbyBypassScope_OrganizationExcludingGuests),
		string(LobbyBypassScope_Organizer),
	}
}

func (s *LobbyBypassScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLobbyBypassScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLobbyBypassScope(input string) (*LobbyBypassScope, error) {
	vals := map[string]LobbyBypassScope{
		"everyone":                    LobbyBypassScope_Everyone,
		"invited":                     LobbyBypassScope_Invited,
		"organization":                LobbyBypassScope_Organization,
		"organizationandfederated":    LobbyBypassScope_OrganizationAndFederated,
		"organizationexcludingguests": LobbyBypassScope_OrganizationExcludingGuests,
		"organizer":                   LobbyBypassScope_Organizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LobbyBypassScope(input)
	return &out, nil
}
