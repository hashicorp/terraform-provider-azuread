package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesRemediationType string

const (
	WindowsUpdatesRemediationType_InPlaceUpgrade WindowsUpdatesRemediationType = "inPlaceUpgrade"
)

func PossibleValuesForWindowsUpdatesRemediationType() []string {
	return []string{
		string(WindowsUpdatesRemediationType_InPlaceUpgrade),
	}
}

func (s *WindowsUpdatesRemediationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesRemediationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesRemediationType(input string) (*WindowsUpdatesRemediationType, error) {
	vals := map[string]WindowsUpdatesRemediationType{
		"inplaceupgrade": WindowsUpdatesRemediationType_InPlaceUpgrade,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesRemediationType(input)
	return &out, nil
}
