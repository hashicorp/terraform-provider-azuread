package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InitiatorType string

const (
	InitiatorType_Application InitiatorType = "application"
	InitiatorType_System      InitiatorType = "system"
	InitiatorType_User        InitiatorType = "user"
)

func PossibleValuesForInitiatorType() []string {
	return []string{
		string(InitiatorType_Application),
		string(InitiatorType_System),
		string(InitiatorType_User),
	}
}

func (s *InitiatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInitiatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInitiatorType(input string) (*InitiatorType, error) {
	vals := map[string]InitiatorType{
		"application": InitiatorType_Application,
		"system":      InitiatorType_System,
		"user":        InitiatorType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InitiatorType(input)
	return &out, nil
}
