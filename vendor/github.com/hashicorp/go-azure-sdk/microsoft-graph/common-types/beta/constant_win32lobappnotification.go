package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppNotification string

const (
	Win32LobAppNotification_HideAll    Win32LobAppNotification = "hideAll"
	Win32LobAppNotification_ShowAll    Win32LobAppNotification = "showAll"
	Win32LobAppNotification_ShowReboot Win32LobAppNotification = "showReboot"
)

func PossibleValuesForWin32LobAppNotification() []string {
	return []string{
		string(Win32LobAppNotification_HideAll),
		string(Win32LobAppNotification_ShowAll),
		string(Win32LobAppNotification_ShowReboot),
	}
}

func (s *Win32LobAppNotification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppNotification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppNotification(input string) (*Win32LobAppNotification, error) {
	vals := map[string]Win32LobAppNotification{
		"hideall":    Win32LobAppNotification_HideAll,
		"showall":    Win32LobAppNotification_ShowAll,
		"showreboot": Win32LobAppNotification_ShowReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppNotification(input)
	return &out, nil
}
