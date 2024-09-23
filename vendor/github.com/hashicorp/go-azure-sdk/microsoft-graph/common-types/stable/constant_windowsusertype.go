package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUserType string

const (
	WindowsUserType_Administrator WindowsUserType = "administrator"
	WindowsUserType_Standard      WindowsUserType = "standard"
)

func PossibleValuesForWindowsUserType() []string {
	return []string{
		string(WindowsUserType_Administrator),
		string(WindowsUserType_Standard),
	}
}

func (s *WindowsUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUserType(input string) (*WindowsUserType, error) {
	vals := map[string]WindowsUserType{
		"administrator": WindowsUserType_Administrator,
		"standard":      WindowsUserType_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUserType(input)
	return &out, nil
}
