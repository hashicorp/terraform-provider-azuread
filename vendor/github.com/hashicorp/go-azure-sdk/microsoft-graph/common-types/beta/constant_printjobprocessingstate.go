package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobProcessingState string

const (
	PrintJobProcessingState_Aborted    PrintJobProcessingState = "aborted"
	PrintJobProcessingState_Canceled   PrintJobProcessingState = "canceled"
	PrintJobProcessingState_Completed  PrintJobProcessingState = "completed"
	PrintJobProcessingState_Paused     PrintJobProcessingState = "paused"
	PrintJobProcessingState_Pending    PrintJobProcessingState = "pending"
	PrintJobProcessingState_Processing PrintJobProcessingState = "processing"
	PrintJobProcessingState_Stopped    PrintJobProcessingState = "stopped"
	PrintJobProcessingState_Unknown    PrintJobProcessingState = "unknown"
)

func PossibleValuesForPrintJobProcessingState() []string {
	return []string{
		string(PrintJobProcessingState_Aborted),
		string(PrintJobProcessingState_Canceled),
		string(PrintJobProcessingState_Completed),
		string(PrintJobProcessingState_Paused),
		string(PrintJobProcessingState_Pending),
		string(PrintJobProcessingState_Processing),
		string(PrintJobProcessingState_Stopped),
		string(PrintJobProcessingState_Unknown),
	}
}

func (s *PrintJobProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintJobProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintJobProcessingState(input string) (*PrintJobProcessingState, error) {
	vals := map[string]PrintJobProcessingState{
		"aborted":    PrintJobProcessingState_Aborted,
		"canceled":   PrintJobProcessingState_Canceled,
		"completed":  PrintJobProcessingState_Completed,
		"paused":     PrintJobProcessingState_Paused,
		"pending":    PrintJobProcessingState_Pending,
		"processing": PrintJobProcessingState_Processing,
		"stopped":    PrintJobProcessingState_Stopped,
		"unknown":    PrintJobProcessingState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintJobProcessingState(input)
	return &out, nil
}
