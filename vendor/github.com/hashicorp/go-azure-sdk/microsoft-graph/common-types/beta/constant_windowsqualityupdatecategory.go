package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsQualityUpdateCategory string

const (
	WindowsQualityUpdateCategory_All         WindowsQualityUpdateCategory = "all"
	WindowsQualityUpdateCategory_NonSecurity WindowsQualityUpdateCategory = "nonSecurity"
	WindowsQualityUpdateCategory_Security    WindowsQualityUpdateCategory = "security"
)

func PossibleValuesForWindowsQualityUpdateCategory() []string {
	return []string{
		string(WindowsQualityUpdateCategory_All),
		string(WindowsQualityUpdateCategory_NonSecurity),
		string(WindowsQualityUpdateCategory_Security),
	}
}

func (s *WindowsQualityUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsQualityUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsQualityUpdateCategory(input string) (*WindowsQualityUpdateCategory, error) {
	vals := map[string]WindowsQualityUpdateCategory{
		"all":         WindowsQualityUpdateCategory_All,
		"nonsecurity": WindowsQualityUpdateCategory_NonSecurity,
		"security":    WindowsQualityUpdateCategory_Security,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsQualityUpdateCategory(input)
	return &out, nil
}
