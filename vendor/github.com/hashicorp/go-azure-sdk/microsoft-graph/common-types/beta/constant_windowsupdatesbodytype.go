package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesBodyType string

const (
	WindowsUpdatesBodyType_Html WindowsUpdatesBodyType = "html"
	WindowsUpdatesBodyType_Text WindowsUpdatesBodyType = "text"
)

func PossibleValuesForWindowsUpdatesBodyType() []string {
	return []string{
		string(WindowsUpdatesBodyType_Html),
		string(WindowsUpdatesBodyType_Text),
	}
}

func (s *WindowsUpdatesBodyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesBodyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesBodyType(input string) (*WindowsUpdatesBodyType, error) {
	vals := map[string]WindowsUpdatesBodyType{
		"html": WindowsUpdatesBodyType_Html,
		"text": WindowsUpdatesBodyType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesBodyType(input)
	return &out, nil
}
