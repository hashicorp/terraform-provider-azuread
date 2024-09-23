package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterProcessingState string

const (
	PrinterProcessingState_Idle       PrinterProcessingState = "idle"
	PrinterProcessingState_Processing PrinterProcessingState = "processing"
	PrinterProcessingState_Stopped    PrinterProcessingState = "stopped"
	PrinterProcessingState_Unknown    PrinterProcessingState = "unknown"
)

func PossibleValuesForPrinterProcessingState() []string {
	return []string{
		string(PrinterProcessingState_Idle),
		string(PrinterProcessingState_Processing),
		string(PrinterProcessingState_Stopped),
		string(PrinterProcessingState_Unknown),
	}
}

func (s *PrinterProcessingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterProcessingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterProcessingState(input string) (*PrinterProcessingState, error) {
	vals := map[string]PrinterProcessingState{
		"idle":       PrinterProcessingState_Idle,
		"processing": PrinterProcessingState_Processing,
		"stopped":    PrinterProcessingState_Stopped,
		"unknown":    PrinterProcessingState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterProcessingState(input)
	return &out, nil
}
