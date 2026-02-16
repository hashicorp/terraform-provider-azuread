package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationStatusCode string

const (
	SynchronizationStatusCode_Active        SynchronizationStatusCode = "Active"
	SynchronizationStatusCode_NotConfigured SynchronizationStatusCode = "NotConfigured"
	SynchronizationStatusCode_NotRun        SynchronizationStatusCode = "NotRun"
	SynchronizationStatusCode_Paused        SynchronizationStatusCode = "Paused"
	SynchronizationStatusCode_Quarantine    SynchronizationStatusCode = "Quarantine"
)

func PossibleValuesForSynchronizationStatusCode() []string {
	return []string{
		string(SynchronizationStatusCode_Active),
		string(SynchronizationStatusCode_NotConfigured),
		string(SynchronizationStatusCode_NotRun),
		string(SynchronizationStatusCode_Paused),
		string(SynchronizationStatusCode_Quarantine),
	}
}

func (s *SynchronizationStatusCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationStatusCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationStatusCode(input string) (*SynchronizationStatusCode, error) {
	vals := map[string]SynchronizationStatusCode{
		"active":        SynchronizationStatusCode_Active,
		"notconfigured": SynchronizationStatusCode_NotConfigured,
		"notrun":        SynchronizationStatusCode_NotRun,
		"paused":        SynchronizationStatusCode_Paused,
		"quarantine":    SynchronizationStatusCode_Quarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationStatusCode(input)
	return &out, nil
}
