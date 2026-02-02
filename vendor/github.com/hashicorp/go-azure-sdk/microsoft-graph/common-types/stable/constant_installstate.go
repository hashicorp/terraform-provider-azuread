package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstallState string

const (
	InstallState_Failed          InstallState = "failed"
	InstallState_Installed       InstallState = "installed"
	InstallState_NotApplicable   InstallState = "notApplicable"
	InstallState_NotInstalled    InstallState = "notInstalled"
	InstallState_UninstallFailed InstallState = "uninstallFailed"
	InstallState_Unknown         InstallState = "unknown"
)

func PossibleValuesForInstallState() []string {
	return []string{
		string(InstallState_Failed),
		string(InstallState_Installed),
		string(InstallState_NotApplicable),
		string(InstallState_NotInstalled),
		string(InstallState_UninstallFailed),
		string(InstallState_Unknown),
	}
}

func (s *InstallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInstallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInstallState(input string) (*InstallState, error) {
	vals := map[string]InstallState{
		"failed":          InstallState_Failed,
		"installed":       InstallState_Installed,
		"notapplicable":   InstallState_NotApplicable,
		"notinstalled":    InstallState_NotInstalled,
		"uninstallfailed": InstallState_UninstallFailed,
		"unknown":         InstallState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InstallState(input)
	return &out, nil
}
