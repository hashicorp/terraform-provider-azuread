package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHelloForBusinessPinUsage string

const (
	WindowsHelloForBusinessPinUsage_Allowed    WindowsHelloForBusinessPinUsage = "allowed"
	WindowsHelloForBusinessPinUsage_Disallowed WindowsHelloForBusinessPinUsage = "disallowed"
	WindowsHelloForBusinessPinUsage_Required   WindowsHelloForBusinessPinUsage = "required"
)

func PossibleValuesForWindowsHelloForBusinessPinUsage() []string {
	return []string{
		string(WindowsHelloForBusinessPinUsage_Allowed),
		string(WindowsHelloForBusinessPinUsage_Disallowed),
		string(WindowsHelloForBusinessPinUsage_Required),
	}
}

func (s *WindowsHelloForBusinessPinUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsHelloForBusinessPinUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsHelloForBusinessPinUsage(input string) (*WindowsHelloForBusinessPinUsage, error) {
	vals := map[string]WindowsHelloForBusinessPinUsage{
		"allowed":    WindowsHelloForBusinessPinUsage_Allowed,
		"disallowed": WindowsHelloForBusinessPinUsage_Disallowed,
		"required":   WindowsHelloForBusinessPinUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHelloForBusinessPinUsage(input)
	return &out, nil
}
