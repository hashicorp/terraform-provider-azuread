package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintOperationProcessingState string

const (
	PrintOperationProcessingState_Failed     PrintOperationProcessingState = "failed"
	PrintOperationProcessingState_NotStarted PrintOperationProcessingState = "notStarted"
	PrintOperationProcessingState_Running    PrintOperationProcessingState = "running"
	PrintOperationProcessingState_Succeeded  PrintOperationProcessingState = "succeeded"
)

func PossibleValuesForPrintOperationProcessingState() []string {
	return []string{
		string(PrintOperationProcessingState_Failed),
		string(PrintOperationProcessingState_NotStarted),
		string(PrintOperationProcessingState_Running),
		string(PrintOperationProcessingState_Succeeded),
	}
}

func (s *PrintOperationProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintOperationProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintOperationProcessingState(input string) (*PrintOperationProcessingState, error) {
	vals := map[string]PrintOperationProcessingState{
		"failed":     PrintOperationProcessingState_Failed,
		"notstarted": PrintOperationProcessingState_NotStarted,
		"running":    PrintOperationProcessingState_Running,
		"succeeded":  PrintOperationProcessingState_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintOperationProcessingState(input)
	return &out, nil
}
