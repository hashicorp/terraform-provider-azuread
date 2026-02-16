package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecutionMode string

const (
	ExecutionMode_EvaluateInline  ExecutionMode = "evaluateInline"
	ExecutionMode_EvaluateOffline ExecutionMode = "evaluateOffline"
)

func PossibleValuesForExecutionMode() []string {
	return []string{
		string(ExecutionMode_EvaluateInline),
		string(ExecutionMode_EvaluateOffline),
	}
}

func (s *ExecutionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExecutionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExecutionMode(input string) (*ExecutionMode, error) {
	vals := map[string]ExecutionMode{
		"evaluateinline":  ExecutionMode_EvaluateInline,
		"evaluateoffline": ExecutionMode_EvaluateOffline,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExecutionMode(input)
	return &out, nil
}
