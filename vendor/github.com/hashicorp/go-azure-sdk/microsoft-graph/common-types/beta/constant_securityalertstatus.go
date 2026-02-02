package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertStatus string

const (
	SecurityAlertStatus_InProgress SecurityAlertStatus = "inProgress"
	SecurityAlertStatus_New        SecurityAlertStatus = "new"
	SecurityAlertStatus_Resolved   SecurityAlertStatus = "resolved"
	SecurityAlertStatus_Unknown    SecurityAlertStatus = "unknown"
)

func PossibleValuesForSecurityAlertStatus() []string {
	return []string{
		string(SecurityAlertStatus_InProgress),
		string(SecurityAlertStatus_New),
		string(SecurityAlertStatus_Resolved),
		string(SecurityAlertStatus_Unknown),
	}
}

func (s *SecurityAlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertStatus(input string) (*SecurityAlertStatus, error) {
	vals := map[string]SecurityAlertStatus{
		"inprogress": SecurityAlertStatus_InProgress,
		"new":        SecurityAlertStatus_New,
		"resolved":   SecurityAlertStatus_Resolved,
		"unknown":    SecurityAlertStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertStatus(input)
	return &out, nil
}
