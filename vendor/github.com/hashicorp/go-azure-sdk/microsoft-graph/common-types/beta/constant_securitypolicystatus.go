package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPolicyStatus string

const (
	SecurityPolicyStatus_Error   SecurityPolicyStatus = "error"
	SecurityPolicyStatus_Pending SecurityPolicyStatus = "pending"
	SecurityPolicyStatus_Success SecurityPolicyStatus = "success"
)

func PossibleValuesForSecurityPolicyStatus() []string {
	return []string{
		string(SecurityPolicyStatus_Error),
		string(SecurityPolicyStatus_Pending),
		string(SecurityPolicyStatus_Success),
	}
}

func (s *SecurityPolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityPolicyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityPolicyStatus(input string) (*SecurityPolicyStatus, error) {
	vals := map[string]SecurityPolicyStatus{
		"error":   SecurityPolicyStatus_Error,
		"pending": SecurityPolicyStatus_Pending,
		"success": SecurityPolicyStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityPolicyStatus(input)
	return &out, nil
}
