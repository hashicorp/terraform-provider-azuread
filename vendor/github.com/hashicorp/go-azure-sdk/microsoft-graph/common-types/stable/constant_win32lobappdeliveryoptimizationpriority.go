package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppDeliveryOptimizationPriority string

const (
	Win32LobAppDeliveryOptimizationPriority_Foreground    Win32LobAppDeliveryOptimizationPriority = "foreground"
	Win32LobAppDeliveryOptimizationPriority_NotConfigured Win32LobAppDeliveryOptimizationPriority = "notConfigured"
)

func PossibleValuesForWin32LobAppDeliveryOptimizationPriority() []string {
	return []string{
		string(Win32LobAppDeliveryOptimizationPriority_Foreground),
		string(Win32LobAppDeliveryOptimizationPriority_NotConfigured),
	}
}

func (s *Win32LobAppDeliveryOptimizationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppDeliveryOptimizationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppDeliveryOptimizationPriority(input string) (*Win32LobAppDeliveryOptimizationPriority, error) {
	vals := map[string]Win32LobAppDeliveryOptimizationPriority{
		"foreground":    Win32LobAppDeliveryOptimizationPriority_Foreground,
		"notconfigured": Win32LobAppDeliveryOptimizationPriority_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppDeliveryOptimizationPriority(input)
	return &out, nil
}
