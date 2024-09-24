package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenSyncStatus string

const (
	VppTokenSyncStatus_Completed  VppTokenSyncStatus = "completed"
	VppTokenSyncStatus_Failed     VppTokenSyncStatus = "failed"
	VppTokenSyncStatus_InProgress VppTokenSyncStatus = "inProgress"
	VppTokenSyncStatus_None       VppTokenSyncStatus = "none"
)

func PossibleValuesForVppTokenSyncStatus() []string {
	return []string{
		string(VppTokenSyncStatus_Completed),
		string(VppTokenSyncStatus_Failed),
		string(VppTokenSyncStatus_InProgress),
		string(VppTokenSyncStatus_None),
	}
}

func (s *VppTokenSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenSyncStatus(input string) (*VppTokenSyncStatus, error) {
	vals := map[string]VppTokenSyncStatus{
		"completed":  VppTokenSyncStatus_Completed,
		"failed":     VppTokenSyncStatus_Failed,
		"inprogress": VppTokenSyncStatus_InProgress,
		"none":       VppTokenSyncStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenSyncStatus(input)
	return &out, nil
}
