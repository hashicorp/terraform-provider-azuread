package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPortStatus string

const (
	SecurityHostPortStatus_Closed   SecurityHostPortStatus = "closed"
	SecurityHostPortStatus_Filtered SecurityHostPortStatus = "filtered"
	SecurityHostPortStatus_Open     SecurityHostPortStatus = "open"
)

func PossibleValuesForSecurityHostPortStatus() []string {
	return []string{
		string(SecurityHostPortStatus_Closed),
		string(SecurityHostPortStatus_Filtered),
		string(SecurityHostPortStatus_Open),
	}
}

func (s *SecurityHostPortStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostPortStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostPortStatus(input string) (*SecurityHostPortStatus, error) {
	vals := map[string]SecurityHostPortStatus{
		"closed":   SecurityHostPortStatus_Closed,
		"filtered": SecurityHostPortStatus_Filtered,
		"open":     SecurityHostPortStatus_Open,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostPortStatus(input)
	return &out, nil
}
