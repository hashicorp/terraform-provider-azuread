package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesSafeguardCategory string

const (
	WindowsUpdatesSafeguardCategory_LikelyIssues WindowsUpdatesSafeguardCategory = "likelyIssues"
)

func PossibleValuesForWindowsUpdatesSafeguardCategory() []string {
	return []string{
		string(WindowsUpdatesSafeguardCategory_LikelyIssues),
	}
}

func (s *WindowsUpdatesSafeguardCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesSafeguardCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesSafeguardCategory(input string) (*WindowsUpdatesSafeguardCategory, error) {
	vals := map[string]WindowsUpdatesSafeguardCategory{
		"likelyissues": WindowsUpdatesSafeguardCategory_LikelyIssues,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesSafeguardCategory(input)
	return &out, nil
}
