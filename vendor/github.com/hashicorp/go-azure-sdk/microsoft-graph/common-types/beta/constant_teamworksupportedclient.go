package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSupportedClient string

const (
	TeamworkSupportedClient_SkypeDefaultAndTeams TeamworkSupportedClient = "skypeDefaultAndTeams"
	TeamworkSupportedClient_SkypeOnly            TeamworkSupportedClient = "skypeOnly"
	TeamworkSupportedClient_TeamsDefaultAndSkype TeamworkSupportedClient = "teamsDefaultAndSkype"
	TeamworkSupportedClient_TeamsOnly            TeamworkSupportedClient = "teamsOnly"
	TeamworkSupportedClient_Unknown              TeamworkSupportedClient = "unknown"
)

func PossibleValuesForTeamworkSupportedClient() []string {
	return []string{
		string(TeamworkSupportedClient_SkypeDefaultAndTeams),
		string(TeamworkSupportedClient_SkypeOnly),
		string(TeamworkSupportedClient_TeamsDefaultAndSkype),
		string(TeamworkSupportedClient_TeamsOnly),
		string(TeamworkSupportedClient_Unknown),
	}
}

func (s *TeamworkSupportedClient) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkSupportedClient(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkSupportedClient(input string) (*TeamworkSupportedClient, error) {
	vals := map[string]TeamworkSupportedClient{
		"skypedefaultandteams": TeamworkSupportedClient_SkypeDefaultAndTeams,
		"skypeonly":            TeamworkSupportedClient_SkypeOnly,
		"teamsdefaultandskype": TeamworkSupportedClient_TeamsDefaultAndSkype,
		"teamsonly":            TeamworkSupportedClient_TeamsOnly,
		"unknown":              TeamworkSupportedClient_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkSupportedClient(input)
	return &out, nil
}
