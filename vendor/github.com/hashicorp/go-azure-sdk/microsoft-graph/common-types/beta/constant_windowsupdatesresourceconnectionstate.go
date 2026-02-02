package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesResourceConnectionState string

const (
	WindowsUpdatesResourceConnectionState_Connected     WindowsUpdatesResourceConnectionState = "connected"
	WindowsUpdatesResourceConnectionState_NotAuthorized WindowsUpdatesResourceConnectionState = "notAuthorized"
	WindowsUpdatesResourceConnectionState_NotFound      WindowsUpdatesResourceConnectionState = "notFound"
)

func PossibleValuesForWindowsUpdatesResourceConnectionState() []string {
	return []string{
		string(WindowsUpdatesResourceConnectionState_Connected),
		string(WindowsUpdatesResourceConnectionState_NotAuthorized),
		string(WindowsUpdatesResourceConnectionState_NotFound),
	}
}

func (s *WindowsUpdatesResourceConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesResourceConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesResourceConnectionState(input string) (*WindowsUpdatesResourceConnectionState, error) {
	vals := map[string]WindowsUpdatesResourceConnectionState{
		"connected":     WindowsUpdatesResourceConnectionState_Connected,
		"notauthorized": WindowsUpdatesResourceConnectionState_NotAuthorized,
		"notfound":      WindowsUpdatesResourceConnectionState_NotFound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesResourceConnectionState(input)
	return &out, nil
}
