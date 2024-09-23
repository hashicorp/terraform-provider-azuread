package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionSource string

const (
	ActionSource_Automatic   ActionSource = "automatic"
	ActionSource_Default     ActionSource = "default"
	ActionSource_Manual      ActionSource = "manual"
	ActionSource_Recommended ActionSource = "recommended"
)

func PossibleValuesForActionSource() []string {
	return []string{
		string(ActionSource_Automatic),
		string(ActionSource_Default),
		string(ActionSource_Manual),
		string(ActionSource_Recommended),
	}
}

func (s *ActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActionSource(input string) (*ActionSource, error) {
	vals := map[string]ActionSource{
		"automatic":   ActionSource_Automatic,
		"default":     ActionSource_Default,
		"manual":      ActionSource_Manual,
		"recommended": ActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionSource(input)
	return &out, nil
}
