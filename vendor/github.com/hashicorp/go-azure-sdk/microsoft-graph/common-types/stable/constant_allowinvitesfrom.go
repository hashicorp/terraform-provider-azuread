package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowInvitesFrom string

const (
	AllowInvitesFrom_AdminsAndGuestInviters           AllowInvitesFrom = "adminsAndGuestInviters"
	AllowInvitesFrom_AdminsGuestInvitersAndAllMembers AllowInvitesFrom = "adminsGuestInvitersAndAllMembers"
	AllowInvitesFrom_Everyone                         AllowInvitesFrom = "everyone"
	AllowInvitesFrom_None                             AllowInvitesFrom = "none"
)

func PossibleValuesForAllowInvitesFrom() []string {
	return []string{
		string(AllowInvitesFrom_AdminsAndGuestInviters),
		string(AllowInvitesFrom_AdminsGuestInvitersAndAllMembers),
		string(AllowInvitesFrom_Everyone),
		string(AllowInvitesFrom_None),
	}
}

func (s *AllowInvitesFrom) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowInvitesFrom(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowInvitesFrom(input string) (*AllowInvitesFrom, error) {
	vals := map[string]AllowInvitesFrom{
		"adminsandguestinviters":           AllowInvitesFrom_AdminsAndGuestInviters,
		"adminsguestinvitersandallmembers": AllowInvitesFrom_AdminsGuestInvitersAndAllMembers,
		"everyone":                         AllowInvitesFrom_Everyone,
		"none":                             AllowInvitesFrom_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowInvitesFrom(input)
	return &out, nil
}
