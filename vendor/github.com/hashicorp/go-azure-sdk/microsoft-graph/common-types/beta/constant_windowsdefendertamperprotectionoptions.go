package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderTamperProtectionOptions string

const (
	WindowsDefenderTamperProtectionOptions_Disable       WindowsDefenderTamperProtectionOptions = "disable"
	WindowsDefenderTamperProtectionOptions_Enable        WindowsDefenderTamperProtectionOptions = "enable"
	WindowsDefenderTamperProtectionOptions_NotConfigured WindowsDefenderTamperProtectionOptions = "notConfigured"
)

func PossibleValuesForWindowsDefenderTamperProtectionOptions() []string {
	return []string{
		string(WindowsDefenderTamperProtectionOptions_Disable),
		string(WindowsDefenderTamperProtectionOptions_Enable),
		string(WindowsDefenderTamperProtectionOptions_NotConfigured),
	}
}

func (s *WindowsDefenderTamperProtectionOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDefenderTamperProtectionOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDefenderTamperProtectionOptions(input string) (*WindowsDefenderTamperProtectionOptions, error) {
	vals := map[string]WindowsDefenderTamperProtectionOptions{
		"disable":       WindowsDefenderTamperProtectionOptions_Disable,
		"enable":        WindowsDefenderTamperProtectionOptions_Enable,
		"notconfigured": WindowsDefenderTamperProtectionOptions_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDefenderTamperProtectionOptions(input)
	return &out, nil
}
