package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdateCategory string

const (
	WindowsUpdatesUpdateCategory_Driver  WindowsUpdatesUpdateCategory = "driver"
	WindowsUpdatesUpdateCategory_Feature WindowsUpdatesUpdateCategory = "feature"
	WindowsUpdatesUpdateCategory_Quality WindowsUpdatesUpdateCategory = "quality"
)

func PossibleValuesForWindowsUpdatesUpdateCategory() []string {
	return []string{
		string(WindowsUpdatesUpdateCategory_Driver),
		string(WindowsUpdatesUpdateCategory_Feature),
		string(WindowsUpdatesUpdateCategory_Quality),
	}
}

func (s *WindowsUpdatesUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesUpdateCategory(input string) (*WindowsUpdatesUpdateCategory, error) {
	vals := map[string]WindowsUpdatesUpdateCategory{
		"driver":  WindowsUpdatesUpdateCategory_Driver,
		"feature": WindowsUpdatesUpdateCategory_Feature,
		"quality": WindowsUpdatesUpdateCategory_Quality,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesUpdateCategory(input)
	return &out, nil
}
