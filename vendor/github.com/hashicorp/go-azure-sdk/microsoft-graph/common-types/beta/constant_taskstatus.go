package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskStatus string

const (
	TaskStatus_Completed       TaskStatus = "completed"
	TaskStatus_Deferred        TaskStatus = "deferred"
	TaskStatus_InProgress      TaskStatus = "inProgress"
	TaskStatus_NotStarted      TaskStatus = "notStarted"
	TaskStatus_WaitingOnOthers TaskStatus = "waitingOnOthers"
)

func PossibleValuesForTaskStatus() []string {
	return []string{
		string(TaskStatus_Completed),
		string(TaskStatus_Deferred),
		string(TaskStatus_InProgress),
		string(TaskStatus_NotStarted),
		string(TaskStatus_WaitingOnOthers),
	}
}

func (s *TaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskStatus(input string) (*TaskStatus, error) {
	vals := map[string]TaskStatus{
		"completed":       TaskStatus_Completed,
		"deferred":        TaskStatus_Deferred,
		"inprogress":      TaskStatus_InProgress,
		"notstarted":      TaskStatus_NotStarted,
		"waitingonothers": TaskStatus_WaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskStatus(input)
	return &out, nil
}
