package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstallIntent string

const (
	InstallIntent_Available                  InstallIntent = "available"
	InstallIntent_AvailableWithoutEnrollment InstallIntent = "availableWithoutEnrollment"
	InstallIntent_Required                   InstallIntent = "required"
	InstallIntent_Uninstall                  InstallIntent = "uninstall"
)

func PossibleValuesForInstallIntent() []string {
	return []string{
		string(InstallIntent_Available),
		string(InstallIntent_AvailableWithoutEnrollment),
		string(InstallIntent_Required),
		string(InstallIntent_Uninstall),
	}
}

func (s *InstallIntent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstallIntent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstallIntent(input string) (*InstallIntent, error) {
	vals := map[string]InstallIntent{
		"available":                  InstallIntent_Available,
		"availablewithoutenrollment": InstallIntent_AvailableWithoutEnrollment,
		"required":                   InstallIntent_Required,
		"uninstall":                  InstallIntent_Uninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstallIntent(input)
	return &out, nil
}
