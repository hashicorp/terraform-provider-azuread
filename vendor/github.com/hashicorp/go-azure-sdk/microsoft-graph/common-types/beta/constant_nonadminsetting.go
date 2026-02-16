package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NonAdminSetting string

const (
	NonAdminSetting_False NonAdminSetting = "false"
	NonAdminSetting_True  NonAdminSetting = "true"
)

func PossibleValuesForNonAdminSetting() []string {
	return []string{
		string(NonAdminSetting_False),
		string(NonAdminSetting_True),
	}
}

func (s *NonAdminSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNonAdminSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNonAdminSetting(input string) (*NonAdminSetting, error) {
	vals := map[string]NonAdminSetting{
		"false": NonAdminSetting_False,
		"true":  NonAdminSetting_True,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NonAdminSetting(input)
	return &out, nil
}
