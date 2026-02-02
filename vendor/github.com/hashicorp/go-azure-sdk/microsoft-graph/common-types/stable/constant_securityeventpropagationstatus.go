package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventPropagationStatus string

const (
	SecurityEventPropagationStatus_Failed       SecurityEventPropagationStatus = "failed"
	SecurityEventPropagationStatus_InProcessing SecurityEventPropagationStatus = "inProcessing"
	SecurityEventPropagationStatus_None         SecurityEventPropagationStatus = "none"
	SecurityEventPropagationStatus_Success      SecurityEventPropagationStatus = "success"
)

func PossibleValuesForSecurityEventPropagationStatus() []string {
	return []string{
		string(SecurityEventPropagationStatus_Failed),
		string(SecurityEventPropagationStatus_InProcessing),
		string(SecurityEventPropagationStatus_None),
		string(SecurityEventPropagationStatus_Success),
	}
}

func (s *SecurityEventPropagationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEventPropagationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEventPropagationStatus(input string) (*SecurityEventPropagationStatus, error) {
	vals := map[string]SecurityEventPropagationStatus{
		"failed":       SecurityEventPropagationStatus_Failed,
		"inprocessing": SecurityEventPropagationStatus_InProcessing,
		"none":         SecurityEventPropagationStatus_None,
		"success":      SecurityEventPropagationStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventPropagationStatus(input)
	return &out, nil
}
