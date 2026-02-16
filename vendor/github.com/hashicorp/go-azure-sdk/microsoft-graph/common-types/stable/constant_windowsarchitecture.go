package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsArchitecture string

const (
	WindowsArchitecture_Arm     WindowsArchitecture = "arm"
	WindowsArchitecture_Neutral WindowsArchitecture = "neutral"
	WindowsArchitecture_None    WindowsArchitecture = "none"
	WindowsArchitecture_X64     WindowsArchitecture = "x64"
	WindowsArchitecture_X86     WindowsArchitecture = "x86"
)

func PossibleValuesForWindowsArchitecture() []string {
	return []string{
		string(WindowsArchitecture_Arm),
		string(WindowsArchitecture_Neutral),
		string(WindowsArchitecture_None),
		string(WindowsArchitecture_X64),
		string(WindowsArchitecture_X86),
	}
}

func (s *WindowsArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsArchitecture(input string) (*WindowsArchitecture, error) {
	vals := map[string]WindowsArchitecture{
		"arm":     WindowsArchitecture_Arm,
		"neutral": WindowsArchitecture_Neutral,
		"none":    WindowsArchitecture_None,
		"x64":     WindowsArchitecture_X64,
		"x86":     WindowsArchitecture_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsArchitecture(input)
	return &out, nil
}
