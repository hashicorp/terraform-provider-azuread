package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotUserlessEnrollmentStatus string

const (
	WindowsAutopilotUserlessEnrollmentStatus_Allowed WindowsAutopilotUserlessEnrollmentStatus = "allowed"
	WindowsAutopilotUserlessEnrollmentStatus_Blocked WindowsAutopilotUserlessEnrollmentStatus = "blocked"
	WindowsAutopilotUserlessEnrollmentStatus_Unknown WindowsAutopilotUserlessEnrollmentStatus = "unknown"
)

func PossibleValuesForWindowsAutopilotUserlessEnrollmentStatus() []string {
	return []string{
		string(WindowsAutopilotUserlessEnrollmentStatus_Allowed),
		string(WindowsAutopilotUserlessEnrollmentStatus_Blocked),
		string(WindowsAutopilotUserlessEnrollmentStatus_Unknown),
	}
}

func (s *WindowsAutopilotUserlessEnrollmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotUserlessEnrollmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotUserlessEnrollmentStatus(input string) (*WindowsAutopilotUserlessEnrollmentStatus, error) {
	vals := map[string]WindowsAutopilotUserlessEnrollmentStatus{
		"allowed": WindowsAutopilotUserlessEnrollmentStatus_Allowed,
		"blocked": WindowsAutopilotUserlessEnrollmentStatus_Blocked,
		"unknown": WindowsAutopilotUserlessEnrollmentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotUserlessEnrollmentStatus(input)
	return &out, nil
}
