package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateNotificationDisplayOption string

const (
	WindowsUpdateNotificationDisplayOption_DefaultNotifications    WindowsUpdateNotificationDisplayOption = "defaultNotifications"
	WindowsUpdateNotificationDisplayOption_DisableAllNotifications WindowsUpdateNotificationDisplayOption = "disableAllNotifications"
	WindowsUpdateNotificationDisplayOption_NotConfigured           WindowsUpdateNotificationDisplayOption = "notConfigured"
	WindowsUpdateNotificationDisplayOption_RestartWarningsOnly     WindowsUpdateNotificationDisplayOption = "restartWarningsOnly"
)

func PossibleValuesForWindowsUpdateNotificationDisplayOption() []string {
	return []string{
		string(WindowsUpdateNotificationDisplayOption_DefaultNotifications),
		string(WindowsUpdateNotificationDisplayOption_DisableAllNotifications),
		string(WindowsUpdateNotificationDisplayOption_NotConfigured),
		string(WindowsUpdateNotificationDisplayOption_RestartWarningsOnly),
	}
}

func (s *WindowsUpdateNotificationDisplayOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateNotificationDisplayOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateNotificationDisplayOption(input string) (*WindowsUpdateNotificationDisplayOption, error) {
	vals := map[string]WindowsUpdateNotificationDisplayOption{
		"defaultnotifications":    WindowsUpdateNotificationDisplayOption_DefaultNotifications,
		"disableallnotifications": WindowsUpdateNotificationDisplayOption_DisableAllNotifications,
		"notconfigured":           WindowsUpdateNotificationDisplayOption_NotConfigured,
		"restartwarningsonly":     WindowsUpdateNotificationDisplayOption_RestartWarningsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateNotificationDisplayOption(input)
	return &out, nil
}
