package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsUserConfigurationAccountType string

const (
	TeamsUserConfigurationAccountType_Guest           TeamsUserConfigurationAccountType = "guest"
	TeamsUserConfigurationAccountType_ResourceAccount TeamsUserConfigurationAccountType = "resourceAccount"
	TeamsUserConfigurationAccountType_SfbOnPremUser   TeamsUserConfigurationAccountType = "sfbOnPremUser"
	TeamsUserConfigurationAccountType_Unknown         TeamsUserConfigurationAccountType = "unknown"
	TeamsUserConfigurationAccountType_User            TeamsUserConfigurationAccountType = "user"
)

func PossibleValuesForTeamsUserConfigurationAccountType() []string {
	return []string{
		string(TeamsUserConfigurationAccountType_Guest),
		string(TeamsUserConfigurationAccountType_ResourceAccount),
		string(TeamsUserConfigurationAccountType_SfbOnPremUser),
		string(TeamsUserConfigurationAccountType_Unknown),
		string(TeamsUserConfigurationAccountType_User),
	}
}

func (s *TeamsUserConfigurationAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsUserConfigurationAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsUserConfigurationAccountType(input string) (*TeamsUserConfigurationAccountType, error) {
	vals := map[string]TeamsUserConfigurationAccountType{
		"guest":           TeamsUserConfigurationAccountType_Guest,
		"resourceaccount": TeamsUserConfigurationAccountType_ResourceAccount,
		"sfbonpremuser":   TeamsUserConfigurationAccountType_SfbOnPremUser,
		"unknown":         TeamsUserConfigurationAccountType_Unknown,
		"user":            TeamsUserConfigurationAccountType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsUserConfigurationAccountType(input)
	return &out, nil
}
