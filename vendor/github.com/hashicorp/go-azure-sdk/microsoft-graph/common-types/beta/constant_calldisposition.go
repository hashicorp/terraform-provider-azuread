package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallDisposition string

const (
	CallDisposition_Default          CallDisposition = "default"
	CallDisposition_Forward          CallDisposition = "forward"
	CallDisposition_SimultaneousRing CallDisposition = "simultaneousRing"
)

func PossibleValuesForCallDisposition() []string {
	return []string{
		string(CallDisposition_Default),
		string(CallDisposition_Forward),
		string(CallDisposition_SimultaneousRing),
	}
}

func (s *CallDisposition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallDisposition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallDisposition(input string) (*CallDisposition, error) {
	vals := map[string]CallDisposition{
		"default":          CallDisposition_Default,
		"forward":          CallDisposition_Forward,
		"simultaneousring": CallDisposition_SimultaneousRing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallDisposition(input)
	return &out, nil
}
