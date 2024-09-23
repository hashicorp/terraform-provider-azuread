package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResultantAppState string

const (
	ResultantAppState_Failed          ResultantAppState = "failed"
	ResultantAppState_Installed       ResultantAppState = "installed"
	ResultantAppState_NotApplicable   ResultantAppState = "notApplicable"
	ResultantAppState_NotInstalled    ResultantAppState = "notInstalled"
	ResultantAppState_PendingInstall  ResultantAppState = "pendingInstall"
	ResultantAppState_UninstallFailed ResultantAppState = "uninstallFailed"
	ResultantAppState_Unknown         ResultantAppState = "unknown"
)

func PossibleValuesForResultantAppState() []string {
	return []string{
		string(ResultantAppState_Failed),
		string(ResultantAppState_Installed),
		string(ResultantAppState_NotApplicable),
		string(ResultantAppState_NotInstalled),
		string(ResultantAppState_PendingInstall),
		string(ResultantAppState_UninstallFailed),
		string(ResultantAppState_Unknown),
	}
}

func (s *ResultantAppState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResultantAppState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResultantAppState(input string) (*ResultantAppState, error) {
	vals := map[string]ResultantAppState{
		"failed":          ResultantAppState_Failed,
		"installed":       ResultantAppState_Installed,
		"notapplicable":   ResultantAppState_NotApplicable,
		"notinstalled":    ResultantAppState_NotInstalled,
		"pendinginstall":  ResultantAppState_PendingInstall,
		"uninstallfailed": ResultantAppState_UninstallFailed,
		"unknown":         ResultantAppState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResultantAppState(input)
	return &out, nil
}
