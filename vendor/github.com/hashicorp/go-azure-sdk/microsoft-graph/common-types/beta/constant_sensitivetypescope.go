package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitiveTypeScope string

const (
	SensitiveTypeScope_FullDocument    SensitiveTypeScope = "fullDocument"
	SensitiveTypeScope_PartialDocument SensitiveTypeScope = "partialDocument"
)

func PossibleValuesForSensitiveTypeScope() []string {
	return []string{
		string(SensitiveTypeScope_FullDocument),
		string(SensitiveTypeScope_PartialDocument),
	}
}

func (s *SensitiveTypeScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitiveTypeScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitiveTypeScope(input string) (*SensitiveTypeScope, error) {
	vals := map[string]SensitiveTypeScope{
		"fulldocument":    SensitiveTypeScope_FullDocument,
		"partialdocument": SensitiveTypeScope_PartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitiveTypeScope(input)
	return &out, nil
}
