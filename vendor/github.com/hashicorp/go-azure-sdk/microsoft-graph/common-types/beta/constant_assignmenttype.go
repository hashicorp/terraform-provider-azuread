package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentType string

const (
	AssignmentType_PeerRecommended AssignmentType = "peerRecommended"
	AssignmentType_Recommended     AssignmentType = "recommended"
	AssignmentType_Required        AssignmentType = "required"
)

func PossibleValuesForAssignmentType() []string {
	return []string{
		string(AssignmentType_PeerRecommended),
		string(AssignmentType_Recommended),
		string(AssignmentType_Required),
	}
}

func (s *AssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentType(input string) (*AssignmentType, error) {
	vals := map[string]AssignmentType{
		"peerrecommended": AssignmentType_PeerRecommended,
		"recommended":     AssignmentType_Recommended,
		"required":        AssignmentType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentType(input)
	return &out, nil
}
