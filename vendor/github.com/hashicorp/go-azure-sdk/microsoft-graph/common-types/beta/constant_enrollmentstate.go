package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentState string

const (
	EnrollmentState_Blocked      EnrollmentState = "blocked"
	EnrollmentState_Enrolled     EnrollmentState = "enrolled"
	EnrollmentState_Failed       EnrollmentState = "failed"
	EnrollmentState_NotContacted EnrollmentState = "notContacted"
	EnrollmentState_PendingReset EnrollmentState = "pendingReset"
	EnrollmentState_Unknown      EnrollmentState = "unknown"
)

func PossibleValuesForEnrollmentState() []string {
	return []string{
		string(EnrollmentState_Blocked),
		string(EnrollmentState_Enrolled),
		string(EnrollmentState_Failed),
		string(EnrollmentState_NotContacted),
		string(EnrollmentState_PendingReset),
		string(EnrollmentState_Unknown),
	}
}

func (s *EnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentState(input string) (*EnrollmentState, error) {
	vals := map[string]EnrollmentState{
		"blocked":      EnrollmentState_Blocked,
		"enrolled":     EnrollmentState_Enrolled,
		"failed":       EnrollmentState_Failed,
		"notcontacted": EnrollmentState_NotContacted,
		"pendingreset": EnrollmentState_PendingReset,
		"unknown":      EnrollmentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentState(input)
	return &out, nil
}
