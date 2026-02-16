package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsState string

const (
	RestrictedAppsState_NotApprovedApps RestrictedAppsState = "notApprovedApps"
	RestrictedAppsState_ProhibitedApps  RestrictedAppsState = "prohibitedApps"
)

func PossibleValuesForRestrictedAppsState() []string {
	return []string{
		string(RestrictedAppsState_NotApprovedApps),
		string(RestrictedAppsState_ProhibitedApps),
	}
}

func (s *RestrictedAppsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestrictedAppsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestrictedAppsState(input string) (*RestrictedAppsState, error) {
	vals := map[string]RestrictedAppsState{
		"notapprovedapps": RestrictedAppsState_NotApprovedApps,
		"prohibitedapps":  RestrictedAppsState_ProhibitedApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictedAppsState(input)
	return &out, nil
}
