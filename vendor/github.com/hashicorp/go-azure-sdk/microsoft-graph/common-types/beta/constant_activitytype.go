package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityType string

const (
	ActivityType_ServicePrincipal ActivityType = "servicePrincipal"
	ActivityType_Signin           ActivityType = "signin"
	ActivityType_User             ActivityType = "user"
)

func PossibleValuesForActivityType() []string {
	return []string{
		string(ActivityType_ServicePrincipal),
		string(ActivityType_Signin),
		string(ActivityType_User),
	}
}

func (s *ActivityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivityType(input string) (*ActivityType, error) {
	vals := map[string]ActivityType{
		"serviceprincipal": ActivityType_ServicePrincipal,
		"signin":           ActivityType_Signin,
		"user":             ActivityType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivityType(input)
	return &out, nil
}
