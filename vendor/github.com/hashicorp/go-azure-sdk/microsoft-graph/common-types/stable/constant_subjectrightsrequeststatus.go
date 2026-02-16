package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStatus string

const (
	SubjectRightsRequestStatus_Active SubjectRightsRequestStatus = "active"
	SubjectRightsRequestStatus_Closed SubjectRightsRequestStatus = "closed"
)

func PossibleValuesForSubjectRightsRequestStatus() []string {
	return []string{
		string(SubjectRightsRequestStatus_Active),
		string(SubjectRightsRequestStatus_Closed),
	}
}

func (s *SubjectRightsRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestStatus(input string) (*SubjectRightsRequestStatus, error) {
	vals := map[string]SubjectRightsRequestStatus{
		"active": SubjectRightsRequestStatus_Active,
		"closed": SubjectRightsRequestStatus_Closed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStatus(input)
	return &out, nil
}
