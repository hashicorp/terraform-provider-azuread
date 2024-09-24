package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAsyncOperationStatus string

const (
	TeamsAsyncOperationStatus_Failed     TeamsAsyncOperationStatus = "failed"
	TeamsAsyncOperationStatus_InProgress TeamsAsyncOperationStatus = "inProgress"
	TeamsAsyncOperationStatus_Invalid    TeamsAsyncOperationStatus = "invalid"
	TeamsAsyncOperationStatus_NotStarted TeamsAsyncOperationStatus = "notStarted"
	TeamsAsyncOperationStatus_Succeeded  TeamsAsyncOperationStatus = "succeeded"
)

func PossibleValuesForTeamsAsyncOperationStatus() []string {
	return []string{
		string(TeamsAsyncOperationStatus_Failed),
		string(TeamsAsyncOperationStatus_InProgress),
		string(TeamsAsyncOperationStatus_Invalid),
		string(TeamsAsyncOperationStatus_NotStarted),
		string(TeamsAsyncOperationStatus_Succeeded),
	}
}

func (s *TeamsAsyncOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAsyncOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAsyncOperationStatus(input string) (*TeamsAsyncOperationStatus, error) {
	vals := map[string]TeamsAsyncOperationStatus{
		"failed":     TeamsAsyncOperationStatus_Failed,
		"inprogress": TeamsAsyncOperationStatus_InProgress,
		"invalid":    TeamsAsyncOperationStatus_Invalid,
		"notstarted": TeamsAsyncOperationStatus_NotStarted,
		"succeeded":  TeamsAsyncOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationStatus(input)
	return &out, nil
}
