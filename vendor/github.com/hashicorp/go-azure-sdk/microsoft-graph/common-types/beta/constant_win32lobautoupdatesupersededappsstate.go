package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAutoUpdateSupersededAppsState string

const (
	Win32LobAutoUpdateSupersededAppsState_Enabled       Win32LobAutoUpdateSupersededAppsState = "enabled"
	Win32LobAutoUpdateSupersededAppsState_NotConfigured Win32LobAutoUpdateSupersededAppsState = "notConfigured"
)

func PossibleValuesForWin32LobAutoUpdateSupersededAppsState() []string {
	return []string{
		string(Win32LobAutoUpdateSupersededAppsState_Enabled),
		string(Win32LobAutoUpdateSupersededAppsState_NotConfigured),
	}
}

func (s *Win32LobAutoUpdateSupersededAppsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAutoUpdateSupersededAppsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAutoUpdateSupersededAppsState(input string) (*Win32LobAutoUpdateSupersededAppsState, error) {
	vals := map[string]Win32LobAutoUpdateSupersededAppsState{
		"enabled":       Win32LobAutoUpdateSupersededAppsState_Enabled,
		"notconfigured": Win32LobAutoUpdateSupersededAppsState_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAutoUpdateSupersededAppsState(input)
	return &out, nil
}
