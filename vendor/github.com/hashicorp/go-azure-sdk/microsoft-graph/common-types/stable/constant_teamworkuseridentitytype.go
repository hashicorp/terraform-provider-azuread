package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkUserIdentityType string

const (
	TeamworkUserIdentityType_AadUser                      TeamworkUserIdentityType = "aadUser"
	TeamworkUserIdentityType_AnonymousGuest               TeamworkUserIdentityType = "anonymousGuest"
	TeamworkUserIdentityType_EmailUser                    TeamworkUserIdentityType = "emailUser"
	TeamworkUserIdentityType_FederatedUser                TeamworkUserIdentityType = "federatedUser"
	TeamworkUserIdentityType_OnPremiseAadUser             TeamworkUserIdentityType = "onPremiseAadUser"
	TeamworkUserIdentityType_PersonalMicrosoftAccountUser TeamworkUserIdentityType = "personalMicrosoftAccountUser"
	TeamworkUserIdentityType_PhoneUser                    TeamworkUserIdentityType = "phoneUser"
	TeamworkUserIdentityType_SkypeUser                    TeamworkUserIdentityType = "skypeUser"
)

func PossibleValuesForTeamworkUserIdentityType() []string {
	return []string{
		string(TeamworkUserIdentityType_AadUser),
		string(TeamworkUserIdentityType_AnonymousGuest),
		string(TeamworkUserIdentityType_EmailUser),
		string(TeamworkUserIdentityType_FederatedUser),
		string(TeamworkUserIdentityType_OnPremiseAadUser),
		string(TeamworkUserIdentityType_PersonalMicrosoftAccountUser),
		string(TeamworkUserIdentityType_PhoneUser),
		string(TeamworkUserIdentityType_SkypeUser),
	}
}

func (s *TeamworkUserIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkUserIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkUserIdentityType(input string) (*TeamworkUserIdentityType, error) {
	vals := map[string]TeamworkUserIdentityType{
		"aaduser":                      TeamworkUserIdentityType_AadUser,
		"anonymousguest":               TeamworkUserIdentityType_AnonymousGuest,
		"emailuser":                    TeamworkUserIdentityType_EmailUser,
		"federateduser":                TeamworkUserIdentityType_FederatedUser,
		"onpremiseaaduser":             TeamworkUserIdentityType_OnPremiseAadUser,
		"personalmicrosoftaccountuser": TeamworkUserIdentityType_PersonalMicrosoftAccountUser,
		"phoneuser":                    TeamworkUserIdentityType_PhoneUser,
		"skypeuser":                    TeamworkUserIdentityType_SkypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkUserIdentityType(input)
	return &out, nil
}
