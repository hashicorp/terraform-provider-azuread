package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintTaskProcessingState string

const (
	PrintTaskProcessingState_Aborted    PrintTaskProcessingState = "aborted"
	PrintTaskProcessingState_Completed  PrintTaskProcessingState = "completed"
	PrintTaskProcessingState_Pending    PrintTaskProcessingState = "pending"
	PrintTaskProcessingState_Processing PrintTaskProcessingState = "processing"
)

func PossibleValuesForPrintTaskProcessingState() []string {
	return []string{
		string(PrintTaskProcessingState_Aborted),
		string(PrintTaskProcessingState_Completed),
		string(PrintTaskProcessingState_Pending),
		string(PrintTaskProcessingState_Processing),
	}
}

func (s *PrintTaskProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintTaskProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintTaskProcessingState(input string) (*PrintTaskProcessingState, error) {
	vals := map[string]PrintTaskProcessingState{
		"aborted":    PrintTaskProcessingState_Aborted,
		"completed":  PrintTaskProcessingState_Completed,
		"pending":    PrintTaskProcessingState_Pending,
		"processing": PrintTaskProcessingState_Processing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintTaskProcessingState(input)
	return &out, nil
}
