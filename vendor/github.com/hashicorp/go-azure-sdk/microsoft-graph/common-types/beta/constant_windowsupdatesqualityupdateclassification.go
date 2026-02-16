package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesQualityUpdateClassification string

const (
	WindowsUpdatesQualityUpdateClassification_All         WindowsUpdatesQualityUpdateClassification = "all"
	WindowsUpdatesQualityUpdateClassification_NonSecurity WindowsUpdatesQualityUpdateClassification = "nonSecurity"
	WindowsUpdatesQualityUpdateClassification_Security    WindowsUpdatesQualityUpdateClassification = "security"
)

func PossibleValuesForWindowsUpdatesQualityUpdateClassification() []string {
	return []string{
		string(WindowsUpdatesQualityUpdateClassification_All),
		string(WindowsUpdatesQualityUpdateClassification_NonSecurity),
		string(WindowsUpdatesQualityUpdateClassification_Security),
	}
}

func (s *WindowsUpdatesQualityUpdateClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesQualityUpdateClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesQualityUpdateClassification(input string) (*WindowsUpdatesQualityUpdateClassification, error) {
	vals := map[string]WindowsUpdatesQualityUpdateClassification{
		"all":         WindowsUpdatesQualityUpdateClassification_All,
		"nonsecurity": WindowsUpdatesQualityUpdateClassification_NonSecurity,
		"security":    WindowsUpdatesQualityUpdateClassification_Security,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesQualityUpdateClassification(input)
	return &out, nil
}
