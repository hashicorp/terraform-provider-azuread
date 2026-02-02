package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunAsAccountType string

const (
	RunAsAccountType_System RunAsAccountType = "system"
	RunAsAccountType_User   RunAsAccountType = "user"
)

func PossibleValuesForRunAsAccountType() []string {
	return []string{
		string(RunAsAccountType_System),
		string(RunAsAccountType_User),
	}
}

func (s *RunAsAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRunAsAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRunAsAccountType(input string) (*RunAsAccountType, error) {
	vals := map[string]RunAsAccountType{
		"system": RunAsAccountType_System,
		"user":   RunAsAccountType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunAsAccountType(input)
	return &out, nil
}
