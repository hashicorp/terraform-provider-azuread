package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAdministrationAccountType string

const (
	TeamsAdministrationAccountType_Guest           TeamsAdministrationAccountType = "guest"
	TeamsAdministrationAccountType_IneligibleUser  TeamsAdministrationAccountType = "ineligibleUser"
	TeamsAdministrationAccountType_ResourceAccount TeamsAdministrationAccountType = "resourceAccount"
	TeamsAdministrationAccountType_SfbOnPremUser   TeamsAdministrationAccountType = "sfbOnPremUser"
	TeamsAdministrationAccountType_Unknown         TeamsAdministrationAccountType = "unknown"
	TeamsAdministrationAccountType_User            TeamsAdministrationAccountType = "user"
)

func PossibleValuesForTeamsAdministrationAccountType() []string {
	return []string{
		string(TeamsAdministrationAccountType_Guest),
		string(TeamsAdministrationAccountType_IneligibleUser),
		string(TeamsAdministrationAccountType_ResourceAccount),
		string(TeamsAdministrationAccountType_SfbOnPremUser),
		string(TeamsAdministrationAccountType_Unknown),
		string(TeamsAdministrationAccountType_User),
	}
}

func (s *TeamsAdministrationAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAdministrationAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAdministrationAccountType(input string) (*TeamsAdministrationAccountType, error) {
	vals := map[string]TeamsAdministrationAccountType{
		"guest":           TeamsAdministrationAccountType_Guest,
		"ineligibleuser":  TeamsAdministrationAccountType_IneligibleUser,
		"resourceaccount": TeamsAdministrationAccountType_ResourceAccount,
		"sfbonpremuser":   TeamsAdministrationAccountType_SfbOnPremUser,
		"unknown":         TeamsAdministrationAccountType_Unknown,
		"user":            TeamsAdministrationAccountType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAdministrationAccountType(input)
	return &out, nil
}
