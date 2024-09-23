package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppNotification string

const (
	WinGetAppNotification_HideAll    WinGetAppNotification = "hideAll"
	WinGetAppNotification_ShowAll    WinGetAppNotification = "showAll"
	WinGetAppNotification_ShowReboot WinGetAppNotification = "showReboot"
)

func PossibleValuesForWinGetAppNotification() []string {
	return []string{
		string(WinGetAppNotification_HideAll),
		string(WinGetAppNotification_ShowAll),
		string(WinGetAppNotification_ShowReboot),
	}
}

func (s *WinGetAppNotification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWinGetAppNotification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWinGetAppNotification(input string) (*WinGetAppNotification, error) {
	vals := map[string]WinGetAppNotification{
		"hideall":    WinGetAppNotification_HideAll,
		"showall":    WinGetAppNotification_ShowAll,
		"showreboot": WinGetAppNotification_ShowReboot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WinGetAppNotification(input)
	return &out, nil
}
