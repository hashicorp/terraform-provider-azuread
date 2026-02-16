package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Enablement string

const (
	Enablement_Disabled      Enablement = "disabled"
	Enablement_Enabled       Enablement = "enabled"
	Enablement_NotConfigured Enablement = "notConfigured"
)

func PossibleValuesForEnablement() []string {
	return []string{
		string(Enablement_Disabled),
		string(Enablement_Enabled),
		string(Enablement_NotConfigured),
	}
}

func (s *Enablement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnablement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnablement(input string) (*Enablement, error) {
	vals := map[string]Enablement{
		"disabled":      Enablement_Disabled,
		"enabled":       Enablement_Enabled,
		"notconfigured": Enablement_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Enablement(input)
	return &out, nil
}
