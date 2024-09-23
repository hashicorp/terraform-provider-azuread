package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDefaultRecordBehavior string

const (
	SecurityDefaultRecordBehavior_StartLocked   SecurityDefaultRecordBehavior = "startLocked"
	SecurityDefaultRecordBehavior_StartUnlocked SecurityDefaultRecordBehavior = "startUnlocked"
)

func PossibleValuesForSecurityDefaultRecordBehavior() []string {
	return []string{
		string(SecurityDefaultRecordBehavior_StartLocked),
		string(SecurityDefaultRecordBehavior_StartUnlocked),
	}
}

func (s *SecurityDefaultRecordBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDefaultRecordBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDefaultRecordBehavior(input string) (*SecurityDefaultRecordBehavior, error) {
	vals := map[string]SecurityDefaultRecordBehavior{
		"startlocked":   SecurityDefaultRecordBehavior_StartLocked,
		"startunlocked": SecurityDefaultRecordBehavior_StartUnlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDefaultRecordBehavior(input)
	return &out, nil
}
