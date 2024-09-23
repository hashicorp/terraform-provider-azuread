package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CourseStatus string

const (
	CourseStatus_Completed  CourseStatus = "completed"
	CourseStatus_InProgress CourseStatus = "inProgress"
	CourseStatus_NotStarted CourseStatus = "notStarted"
)

func PossibleValuesForCourseStatus() []string {
	return []string{
		string(CourseStatus_Completed),
		string(CourseStatus_InProgress),
		string(CourseStatus_NotStarted),
	}
}

func (s *CourseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCourseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCourseStatus(input string) (*CourseStatus, error) {
	vals := map[string]CourseStatus{
		"completed":  CourseStatus_Completed,
		"inprogress": CourseStatus_InProgress,
		"notstarted": CourseStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CourseStatus(input)
	return &out, nil
}
