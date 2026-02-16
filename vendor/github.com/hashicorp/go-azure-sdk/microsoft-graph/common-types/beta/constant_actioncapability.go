package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionCapability string

const (
	ActionCapability_Disabled ActionCapability = "disabled"
	ActionCapability_Enabled  ActionCapability = "enabled"
)

func PossibleValuesForActionCapability() []string {
	return []string{
		string(ActionCapability_Disabled),
		string(ActionCapability_Enabled),
	}
}

func (s *ActionCapability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActionCapability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActionCapability(input string) (*ActionCapability, error) {
	vals := map[string]ActionCapability{
		"disabled": ActionCapability_Disabled,
		"enabled":  ActionCapability_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionCapability(input)
	return &out, nil
}
