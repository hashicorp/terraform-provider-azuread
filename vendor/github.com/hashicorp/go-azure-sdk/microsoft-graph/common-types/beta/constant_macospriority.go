package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPriority string

const (
	MacOSPriority_High MacOSPriority = "high"
	MacOSPriority_Low  MacOSPriority = "low"
)

func PossibleValuesForMacOSPriority() []string {
	return []string{
		string(MacOSPriority_High),
		string(MacOSPriority_Low),
	}
}

func (s *MacOSPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPriority(input string) (*MacOSPriority, error) {
	vals := map[string]MacOSPriority{
		"high": MacOSPriority_High,
		"low":  MacOSPriority_Low,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPriority(input)
	return &out, nil
}
