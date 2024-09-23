package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyOperationStatus string

const (
	GroupPolicyOperationStatus_Failed     GroupPolicyOperationStatus = "failed"
	GroupPolicyOperationStatus_InProgress GroupPolicyOperationStatus = "inProgress"
	GroupPolicyOperationStatus_Success    GroupPolicyOperationStatus = "success"
	GroupPolicyOperationStatus_Unknown    GroupPolicyOperationStatus = "unknown"
)

func PossibleValuesForGroupPolicyOperationStatus() []string {
	return []string{
		string(GroupPolicyOperationStatus_Failed),
		string(GroupPolicyOperationStatus_InProgress),
		string(GroupPolicyOperationStatus_Success),
		string(GroupPolicyOperationStatus_Unknown),
	}
}

func (s *GroupPolicyOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyOperationStatus(input string) (*GroupPolicyOperationStatus, error) {
	vals := map[string]GroupPolicyOperationStatus{
		"failed":     GroupPolicyOperationStatus_Failed,
		"inprogress": GroupPolicyOperationStatus_InProgress,
		"success":    GroupPolicyOperationStatus_Success,
		"unknown":    GroupPolicyOperationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyOperationStatus(input)
	return &out, nil
}
