package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoRestartNotificationDismissalMethod string

const (
	AutoRestartNotificationDismissalMethod_Automatic     AutoRestartNotificationDismissalMethod = "automatic"
	AutoRestartNotificationDismissalMethod_NotConfigured AutoRestartNotificationDismissalMethod = "notConfigured"
	AutoRestartNotificationDismissalMethod_User          AutoRestartNotificationDismissalMethod = "user"
)

func PossibleValuesForAutoRestartNotificationDismissalMethod() []string {
	return []string{
		string(AutoRestartNotificationDismissalMethod_Automatic),
		string(AutoRestartNotificationDismissalMethod_NotConfigured),
		string(AutoRestartNotificationDismissalMethod_User),
	}
}

func (s *AutoRestartNotificationDismissalMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoRestartNotificationDismissalMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoRestartNotificationDismissalMethod(input string) (*AutoRestartNotificationDismissalMethod, error) {
	vals := map[string]AutoRestartNotificationDismissalMethod{
		"automatic":     AutoRestartNotificationDismissalMethod_Automatic,
		"notconfigured": AutoRestartNotificationDismissalMethod_NotConfigured,
		"user":          AutoRestartNotificationDismissalMethod_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoRestartNotificationDismissalMethod(input)
	return &out, nil
}
