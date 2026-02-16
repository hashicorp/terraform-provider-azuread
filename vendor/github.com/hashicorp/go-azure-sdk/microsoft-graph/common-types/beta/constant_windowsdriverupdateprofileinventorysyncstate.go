package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfileInventorySyncState string

const (
	WindowsDriverUpdateProfileInventorySyncState_Failure WindowsDriverUpdateProfileInventorySyncState = "failure"
	WindowsDriverUpdateProfileInventorySyncState_Pending WindowsDriverUpdateProfileInventorySyncState = "pending"
	WindowsDriverUpdateProfileInventorySyncState_Success WindowsDriverUpdateProfileInventorySyncState = "success"
)

func PossibleValuesForWindowsDriverUpdateProfileInventorySyncState() []string {
	return []string{
		string(WindowsDriverUpdateProfileInventorySyncState_Failure),
		string(WindowsDriverUpdateProfileInventorySyncState_Pending),
		string(WindowsDriverUpdateProfileInventorySyncState_Success),
	}
}

func (s *WindowsDriverUpdateProfileInventorySyncState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDriverUpdateProfileInventorySyncState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDriverUpdateProfileInventorySyncState(input string) (*WindowsDriverUpdateProfileInventorySyncState, error) {
	vals := map[string]WindowsDriverUpdateProfileInventorySyncState{
		"failure": WindowsDriverUpdateProfileInventorySyncState_Failure,
		"pending": WindowsDriverUpdateProfileInventorySyncState_Pending,
		"success": WindowsDriverUpdateProfileInventorySyncState_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDriverUpdateProfileInventorySyncState(input)
	return &out, nil
}
