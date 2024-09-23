package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationResult string

const (
	OperationResult_Failure OperationResult = "failure"
	OperationResult_Success OperationResult = "success"
	OperationResult_Timeout OperationResult = "timeout"
)

func PossibleValuesForOperationResult() []string {
	return []string{
		string(OperationResult_Failure),
		string(OperationResult_Success),
		string(OperationResult_Timeout),
	}
}

func (s *OperationResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationResult(input string) (*OperationResult, error) {
	vals := map[string]OperationResult{
		"failure": OperationResult_Failure,
		"success": OperationResult_Success,
		"timeout": OperationResult_Timeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationResult(input)
	return &out, nil
}
