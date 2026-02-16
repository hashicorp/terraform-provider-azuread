package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepTokenType string

const (
	DepTokenType_AppleSchoolManager DepTokenType = "appleSchoolManager"
	DepTokenType_Dep                DepTokenType = "dep"
	DepTokenType_None               DepTokenType = "none"
)

func PossibleValuesForDepTokenType() []string {
	return []string{
		string(DepTokenType_AppleSchoolManager),
		string(DepTokenType_Dep),
		string(DepTokenType_None),
	}
}

func (s *DepTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDepTokenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDepTokenType(input string) (*DepTokenType, error) {
	vals := map[string]DepTokenType{
		"appleschoolmanager": DepTokenType_AppleSchoolManager,
		"dep":                DepTokenType_Dep,
		"none":               DepTokenType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DepTokenType(input)
	return &out, nil
}
