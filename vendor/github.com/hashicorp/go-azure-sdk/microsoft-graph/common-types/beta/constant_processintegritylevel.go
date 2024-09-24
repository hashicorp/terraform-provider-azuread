package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessIntegrityLevel string

const (
	ProcessIntegrityLevel_High      ProcessIntegrityLevel = "high"
	ProcessIntegrityLevel_Low       ProcessIntegrityLevel = "low"
	ProcessIntegrityLevel_Medium    ProcessIntegrityLevel = "medium"
	ProcessIntegrityLevel_System    ProcessIntegrityLevel = "system"
	ProcessIntegrityLevel_Unknown   ProcessIntegrityLevel = "unknown"
	ProcessIntegrityLevel_Untrusted ProcessIntegrityLevel = "untrusted"
)

func PossibleValuesForProcessIntegrityLevel() []string {
	return []string{
		string(ProcessIntegrityLevel_High),
		string(ProcessIntegrityLevel_Low),
		string(ProcessIntegrityLevel_Medium),
		string(ProcessIntegrityLevel_System),
		string(ProcessIntegrityLevel_Unknown),
		string(ProcessIntegrityLevel_Untrusted),
	}
}

func (s *ProcessIntegrityLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProcessIntegrityLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProcessIntegrityLevel(input string) (*ProcessIntegrityLevel, error) {
	vals := map[string]ProcessIntegrityLevel{
		"high":      ProcessIntegrityLevel_High,
		"low":       ProcessIntegrityLevel_Low,
		"medium":    ProcessIntegrityLevel_Medium,
		"system":    ProcessIntegrityLevel_System,
		"unknown":   ProcessIntegrityLevel_Unknown,
		"untrusted": ProcessIntegrityLevel_Untrusted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProcessIntegrityLevel(input)
	return &out, nil
}
