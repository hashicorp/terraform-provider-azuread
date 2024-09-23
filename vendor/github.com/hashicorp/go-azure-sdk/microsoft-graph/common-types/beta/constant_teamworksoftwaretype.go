package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkSoftwareType string

const (
	TeamworkSoftwareType_AdminAgent      TeamworkSoftwareType = "adminAgent"
	TeamworkSoftwareType_CompanyPortal   TeamworkSoftwareType = "companyPortal"
	TeamworkSoftwareType_Firmware        TeamworkSoftwareType = "firmware"
	TeamworkSoftwareType_OperatingSystem TeamworkSoftwareType = "operatingSystem"
	TeamworkSoftwareType_PartnerAgent    TeamworkSoftwareType = "partnerAgent"
	TeamworkSoftwareType_TeamsClient     TeamworkSoftwareType = "teamsClient"
)

func PossibleValuesForTeamworkSoftwareType() []string {
	return []string{
		string(TeamworkSoftwareType_AdminAgent),
		string(TeamworkSoftwareType_CompanyPortal),
		string(TeamworkSoftwareType_Firmware),
		string(TeamworkSoftwareType_OperatingSystem),
		string(TeamworkSoftwareType_PartnerAgent),
		string(TeamworkSoftwareType_TeamsClient),
	}
}

func (s *TeamworkSoftwareType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkSoftwareType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkSoftwareType(input string) (*TeamworkSoftwareType, error) {
	vals := map[string]TeamworkSoftwareType{
		"adminagent":      TeamworkSoftwareType_AdminAgent,
		"companyportal":   TeamworkSoftwareType_CompanyPortal,
		"firmware":        TeamworkSoftwareType_Firmware,
		"operatingsystem": TeamworkSoftwareType_OperatingSystem,
		"partneragent":    TeamworkSoftwareType_PartnerAgent,
		"teamsclient":     TeamworkSoftwareType_TeamsClient,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkSoftwareType(input)
	return &out, nil
}
