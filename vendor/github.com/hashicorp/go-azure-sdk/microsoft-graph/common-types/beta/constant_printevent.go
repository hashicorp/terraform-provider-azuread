package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintEvent string

const (
	PrintEvent_JobStarted PrintEvent = "jobStarted"
)

func PossibleValuesForPrintEvent() []string {
	return []string{
		string(PrintEvent_JobStarted),
	}
}

func (s *PrintEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintEvent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintEvent(input string) (*PrintEvent, error) {
	vals := map[string]PrintEvent{
		"jobstarted": PrintEvent_JobStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintEvent(input)
	return &out, nil
}
