package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHealthIssueType string

const (
	SecurityHealthIssueType_Global SecurityHealthIssueType = "global"
	SecurityHealthIssueType_Sensor SecurityHealthIssueType = "sensor"
)

func PossibleValuesForSecurityHealthIssueType() []string {
	return []string{
		string(SecurityHealthIssueType_Global),
		string(SecurityHealthIssueType_Sensor),
	}
}

func (s *SecurityHealthIssueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHealthIssueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHealthIssueType(input string) (*SecurityHealthIssueType, error) {
	vals := map[string]SecurityHealthIssueType{
		"global": SecurityHealthIssueType_Global,
		"sensor": SecurityHealthIssueType_Sensor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHealthIssueType(input)
	return &out, nil
}
