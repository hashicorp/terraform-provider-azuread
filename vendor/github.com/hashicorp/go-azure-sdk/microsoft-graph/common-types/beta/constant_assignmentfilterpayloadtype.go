package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterPayloadType string

const (
	AssignmentFilterPayloadType_EnrollmentRestrictions AssignmentFilterPayloadType = "enrollmentRestrictions"
	AssignmentFilterPayloadType_NotSet                 AssignmentFilterPayloadType = "notSet"
)

func PossibleValuesForAssignmentFilterPayloadType() []string {
	return []string{
		string(AssignmentFilterPayloadType_EnrollmentRestrictions),
		string(AssignmentFilterPayloadType_NotSet),
	}
}

func (s *AssignmentFilterPayloadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterPayloadType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterPayloadType(input string) (*AssignmentFilterPayloadType, error) {
	vals := map[string]AssignmentFilterPayloadType{
		"enrollmentrestrictions": AssignmentFilterPayloadType_EnrollmentRestrictions,
		"notset":                 AssignmentFilterPayloadType_NotSet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterPayloadType(input)
	return &out, nil
}
