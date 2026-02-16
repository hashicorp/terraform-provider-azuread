package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationResult string

const (
	AssignmentFilterEvaluationResult_Failure      AssignmentFilterEvaluationResult = "failure"
	AssignmentFilterEvaluationResult_Inconclusive AssignmentFilterEvaluationResult = "inconclusive"
	AssignmentFilterEvaluationResult_Match        AssignmentFilterEvaluationResult = "match"
	AssignmentFilterEvaluationResult_NotEvaluated AssignmentFilterEvaluationResult = "notEvaluated"
	AssignmentFilterEvaluationResult_NotMatch     AssignmentFilterEvaluationResult = "notMatch"
	AssignmentFilterEvaluationResult_Unknown      AssignmentFilterEvaluationResult = "unknown"
)

func PossibleValuesForAssignmentFilterEvaluationResult() []string {
	return []string{
		string(AssignmentFilterEvaluationResult_Failure),
		string(AssignmentFilterEvaluationResult_Inconclusive),
		string(AssignmentFilterEvaluationResult_Match),
		string(AssignmentFilterEvaluationResult_NotEvaluated),
		string(AssignmentFilterEvaluationResult_NotMatch),
		string(AssignmentFilterEvaluationResult_Unknown),
	}
}

func (s *AssignmentFilterEvaluationResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterEvaluationResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterEvaluationResult(input string) (*AssignmentFilterEvaluationResult, error) {
	vals := map[string]AssignmentFilterEvaluationResult{
		"failure":      AssignmentFilterEvaluationResult_Failure,
		"inconclusive": AssignmentFilterEvaluationResult_Inconclusive,
		"match":        AssignmentFilterEvaluationResult_Match,
		"notevaluated": AssignmentFilterEvaluationResult_NotEvaluated,
		"notmatch":     AssignmentFilterEvaluationResult_NotMatch,
		"unknown":      AssignmentFilterEvaluationResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationResult(input)
	return &out, nil
}
