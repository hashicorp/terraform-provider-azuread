package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesEnrollmentState string

const (
	WindowsUpdatesEnrollmentState_Enrolled           WindowsUpdatesEnrollmentState = "enrolled"
	WindowsUpdatesEnrollmentState_EnrolledWithPolicy WindowsUpdatesEnrollmentState = "enrolledWithPolicy"
	WindowsUpdatesEnrollmentState_Enrolling          WindowsUpdatesEnrollmentState = "enrolling"
	WindowsUpdatesEnrollmentState_NotEnrolled        WindowsUpdatesEnrollmentState = "notEnrolled"
	WindowsUpdatesEnrollmentState_Unenrolling        WindowsUpdatesEnrollmentState = "unenrolling"
)

func PossibleValuesForWindowsUpdatesEnrollmentState() []string {
	return []string{
		string(WindowsUpdatesEnrollmentState_Enrolled),
		string(WindowsUpdatesEnrollmentState_EnrolledWithPolicy),
		string(WindowsUpdatesEnrollmentState_Enrolling),
		string(WindowsUpdatesEnrollmentState_NotEnrolled),
		string(WindowsUpdatesEnrollmentState_Unenrolling),
	}
}

func (s *WindowsUpdatesEnrollmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesEnrollmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesEnrollmentState(input string) (*WindowsUpdatesEnrollmentState, error) {
	vals := map[string]WindowsUpdatesEnrollmentState{
		"enrolled":           WindowsUpdatesEnrollmentState_Enrolled,
		"enrolledwithpolicy": WindowsUpdatesEnrollmentState_EnrolledWithPolicy,
		"enrolling":          WindowsUpdatesEnrollmentState_Enrolling,
		"notenrolled":        WindowsUpdatesEnrollmentState_NotEnrolled,
		"unenrolling":        WindowsUpdatesEnrollmentState_Unenrolling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesEnrollmentState(input)
	return &out, nil
}
