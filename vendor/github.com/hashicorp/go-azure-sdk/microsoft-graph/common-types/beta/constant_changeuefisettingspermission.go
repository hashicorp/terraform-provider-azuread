package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeUefiSettingsPermission string

const (
	ChangeUefiSettingsPermission_None              ChangeUefiSettingsPermission = "none"
	ChangeUefiSettingsPermission_NotConfiguredOnly ChangeUefiSettingsPermission = "notConfiguredOnly"
)

func PossibleValuesForChangeUefiSettingsPermission() []string {
	return []string{
		string(ChangeUefiSettingsPermission_None),
		string(ChangeUefiSettingsPermission_NotConfiguredOnly),
	}
}

func (s *ChangeUefiSettingsPermission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChangeUefiSettingsPermission(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChangeUefiSettingsPermission(input string) (*ChangeUefiSettingsPermission, error) {
	vals := map[string]ChangeUefiSettingsPermission{
		"none":              ChangeUefiSettingsPermission_None,
		"notconfiguredonly": ChangeUefiSettingsPermission_NotConfiguredOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChangeUefiSettingsPermission(input)
	return &out, nil
}
