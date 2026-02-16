package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDetectionStatus string

const (
	SecurityDetectionStatus_Blocked   SecurityDetectionStatus = "blocked"
	SecurityDetectionStatus_Detected  SecurityDetectionStatus = "detected"
	SecurityDetectionStatus_Prevented SecurityDetectionStatus = "prevented"
)

func PossibleValuesForSecurityDetectionStatus() []string {
	return []string{
		string(SecurityDetectionStatus_Blocked),
		string(SecurityDetectionStatus_Detected),
		string(SecurityDetectionStatus_Prevented),
	}
}

func (s *SecurityDetectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDetectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDetectionStatus(input string) (*SecurityDetectionStatus, error) {
	vals := map[string]SecurityDetectionStatus{
		"blocked":   SecurityDetectionStatus_Blocked,
		"detected":  SecurityDetectionStatus_Detected,
		"prevented": SecurityDetectionStatus_Prevented,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDetectionStatus(input)
	return &out, nil
}
