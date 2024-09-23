package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateClassification string

const (
	WindowsQualityUpdateClassification_All         WindowsQualityUpdateClassification = "all"
	WindowsQualityUpdateClassification_NonSecurity WindowsQualityUpdateClassification = "nonSecurity"
	WindowsQualityUpdateClassification_Security    WindowsQualityUpdateClassification = "security"
)

func PossibleValuesForWindowsQualityUpdateClassification() []string {
	return []string{
		string(WindowsQualityUpdateClassification_All),
		string(WindowsQualityUpdateClassification_NonSecurity),
		string(WindowsQualityUpdateClassification_Security),
	}
}

func (s *WindowsQualityUpdateClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdateClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdateClassification(input string) (*WindowsQualityUpdateClassification, error) {
	vals := map[string]WindowsQualityUpdateClassification{
		"all":         WindowsQualityUpdateClassification_All,
		"nonsecurity": WindowsQualityUpdateClassification_NonSecurity,
		"security":    WindowsQualityUpdateClassification_Security,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdateClassification(input)
	return &out, nil
}
