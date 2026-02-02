package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppLogUploadState string

const (
	ManagedAppLogUploadState_Completed      ManagedAppLogUploadState = "completed"
	ManagedAppLogUploadState_DeclinedByUser ManagedAppLogUploadState = "declinedByUser"
	ManagedAppLogUploadState_Failed         ManagedAppLogUploadState = "failed"
	ManagedAppLogUploadState_InProgress     ManagedAppLogUploadState = "inProgress"
	ManagedAppLogUploadState_Pending        ManagedAppLogUploadState = "pending"
	ManagedAppLogUploadState_TimedOut       ManagedAppLogUploadState = "timedOut"
)

func PossibleValuesForManagedAppLogUploadState() []string {
	return []string{
		string(ManagedAppLogUploadState_Completed),
		string(ManagedAppLogUploadState_DeclinedByUser),
		string(ManagedAppLogUploadState_Failed),
		string(ManagedAppLogUploadState_InProgress),
		string(ManagedAppLogUploadState_Pending),
		string(ManagedAppLogUploadState_TimedOut),
	}
}

func (s *ManagedAppLogUploadState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppLogUploadState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppLogUploadState(input string) (*ManagedAppLogUploadState, error) {
	vals := map[string]ManagedAppLogUploadState{
		"completed":      ManagedAppLogUploadState_Completed,
		"declinedbyuser": ManagedAppLogUploadState_DeclinedByUser,
		"failed":         ManagedAppLogUploadState_Failed,
		"inprogress":     ManagedAppLogUploadState_InProgress,
		"pending":        ManagedAppLogUploadState_Pending,
		"timedout":       ManagedAppLogUploadState_TimedOut,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppLogUploadState(input)
	return &out, nil
}
