package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTaskExecutionResult string

const (
	SynchronizationTaskExecutionResult_EntryLevelErrors SynchronizationTaskExecutionResult = "EntryLevelErrors"
	SynchronizationTaskExecutionResult_Failed           SynchronizationTaskExecutionResult = "Failed"
	SynchronizationTaskExecutionResult_Succeeded        SynchronizationTaskExecutionResult = "Succeeded"
)

func PossibleValuesForSynchronizationTaskExecutionResult() []string {
	return []string{
		string(SynchronizationTaskExecutionResult_EntryLevelErrors),
		string(SynchronizationTaskExecutionResult_Failed),
		string(SynchronizationTaskExecutionResult_Succeeded),
	}
}

func (s *SynchronizationTaskExecutionResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationTaskExecutionResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationTaskExecutionResult(input string) (*SynchronizationTaskExecutionResult, error) {
	vals := map[string]SynchronizationTaskExecutionResult{
		"entrylevelerrors": SynchronizationTaskExecutionResult_EntryLevelErrors,
		"failed":           SynchronizationTaskExecutionResult_Failed,
		"succeeded":        SynchronizationTaskExecutionResult_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationTaskExecutionResult(input)
	return &out, nil
}
