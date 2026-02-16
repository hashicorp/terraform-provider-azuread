package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationPermissionsRequired string

const (
	ApplicationPermissionsRequired_Administrator ApplicationPermissionsRequired = "administrator"
	ApplicationPermissionsRequired_Anonymous     ApplicationPermissionsRequired = "anonymous"
	ApplicationPermissionsRequired_Guest         ApplicationPermissionsRequired = "guest"
	ApplicationPermissionsRequired_System        ApplicationPermissionsRequired = "system"
	ApplicationPermissionsRequired_Unknown       ApplicationPermissionsRequired = "unknown"
	ApplicationPermissionsRequired_User          ApplicationPermissionsRequired = "user"
)

func PossibleValuesForApplicationPermissionsRequired() []string {
	return []string{
		string(ApplicationPermissionsRequired_Administrator),
		string(ApplicationPermissionsRequired_Anonymous),
		string(ApplicationPermissionsRequired_Guest),
		string(ApplicationPermissionsRequired_System),
		string(ApplicationPermissionsRequired_Unknown),
		string(ApplicationPermissionsRequired_User),
	}
}

func (s *ApplicationPermissionsRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApplicationPermissionsRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApplicationPermissionsRequired(input string) (*ApplicationPermissionsRequired, error) {
	vals := map[string]ApplicationPermissionsRequired{
		"administrator": ApplicationPermissionsRequired_Administrator,
		"anonymous":     ApplicationPermissionsRequired_Anonymous,
		"guest":         ApplicationPermissionsRequired_Guest,
		"system":        ApplicationPermissionsRequired_System,
		"unknown":       ApplicationPermissionsRequired_Unknown,
		"user":          ApplicationPermissionsRequired_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApplicationPermissionsRequired(input)
	return &out, nil
}
