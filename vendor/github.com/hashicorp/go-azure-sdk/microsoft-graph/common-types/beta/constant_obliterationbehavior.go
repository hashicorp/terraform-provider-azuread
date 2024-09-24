package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObliterationBehavior string

const (
	ObliterationBehavior_Always                ObliterationBehavior = "always"
	ObliterationBehavior_Default               ObliterationBehavior = "default"
	ObliterationBehavior_DoNotObliterate       ObliterationBehavior = "doNotObliterate"
	ObliterationBehavior_ObliterateWithWarning ObliterationBehavior = "obliterateWithWarning"
)

func PossibleValuesForObliterationBehavior() []string {
	return []string{
		string(ObliterationBehavior_Always),
		string(ObliterationBehavior_Default),
		string(ObliterationBehavior_DoNotObliterate),
		string(ObliterationBehavior_ObliterateWithWarning),
	}
}

func (s *ObliterationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseObliterationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseObliterationBehavior(input string) (*ObliterationBehavior, error) {
	vals := map[string]ObliterationBehavior{
		"always":                ObliterationBehavior_Always,
		"default":               ObliterationBehavior_Default,
		"donotobliterate":       ObliterationBehavior_DoNotObliterate,
		"obliteratewithwarning": ObliterationBehavior_ObliterateWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObliterationBehavior(input)
	return &out, nil
}
