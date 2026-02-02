package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppManagementRestrictionState string

const (
	AppManagementRestrictionState_Disabled AppManagementRestrictionState = "disabled"
	AppManagementRestrictionState_Enabled  AppManagementRestrictionState = "enabled"
)

func PossibleValuesForAppManagementRestrictionState() []string {
	return []string{
		string(AppManagementRestrictionState_Disabled),
		string(AppManagementRestrictionState_Enabled),
	}
}

func (s *AppManagementRestrictionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppManagementRestrictionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppManagementRestrictionState(input string) (*AppManagementRestrictionState, error) {
	vals := map[string]AppManagementRestrictionState{
		"disabled": AppManagementRestrictionState_Disabled,
		"enabled":  AppManagementRestrictionState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppManagementRestrictionState(input)
	return &out, nil
}
