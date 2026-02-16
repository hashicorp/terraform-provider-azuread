package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesCveSeverityLevel string

const (
	WindowsUpdatesCveSeverityLevel_Critical  WindowsUpdatesCveSeverityLevel = "critical"
	WindowsUpdatesCveSeverityLevel_Important WindowsUpdatesCveSeverityLevel = "important"
	WindowsUpdatesCveSeverityLevel_Moderate  WindowsUpdatesCveSeverityLevel = "moderate"
)

func PossibleValuesForWindowsUpdatesCveSeverityLevel() []string {
	return []string{
		string(WindowsUpdatesCveSeverityLevel_Critical),
		string(WindowsUpdatesCveSeverityLevel_Important),
		string(WindowsUpdatesCveSeverityLevel_Moderate),
	}
}

func (s *WindowsUpdatesCveSeverityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesCveSeverityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesCveSeverityLevel(input string) (*WindowsUpdatesCveSeverityLevel, error) {
	vals := map[string]WindowsUpdatesCveSeverityLevel{
		"critical":  WindowsUpdatesCveSeverityLevel_Critical,
		"important": WindowsUpdatesCveSeverityLevel_Important,
		"moderate":  WindowsUpdatesCveSeverityLevel_Moderate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesCveSeverityLevel(input)
	return &out, nil
}
