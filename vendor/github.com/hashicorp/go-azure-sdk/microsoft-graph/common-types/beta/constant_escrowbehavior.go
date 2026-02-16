package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EscrowBehavior string

const (
	EscrowBehavior_Default                                EscrowBehavior = "Default"
	EscrowBehavior_IgnoreLookupReferenceResolutionFailure EscrowBehavior = "IgnoreLookupReferenceResolutionFailure"
)

func PossibleValuesForEscrowBehavior() []string {
	return []string{
		string(EscrowBehavior_Default),
		string(EscrowBehavior_IgnoreLookupReferenceResolutionFailure),
	}
}

func (s *EscrowBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEscrowBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEscrowBehavior(input string) (*EscrowBehavior, error) {
	vals := map[string]EscrowBehavior{
		"default":                                EscrowBehavior_Default,
		"ignorelookupreferenceresolutionfailure": EscrowBehavior_IgnoreLookupReferenceResolutionFailure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EscrowBehavior(input)
	return &out, nil
}
