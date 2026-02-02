package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreSessionStatus string

const (
	RestoreSessionStatus_Activating         RestoreSessionStatus = "activating"
	RestoreSessionStatus_Active             RestoreSessionStatus = "active"
	RestoreSessionStatus_Completed          RestoreSessionStatus = "completed"
	RestoreSessionStatus_CompletedWithError RestoreSessionStatus = "completedWithError"
	RestoreSessionStatus_Draft              RestoreSessionStatus = "draft"
	RestoreSessionStatus_Failed             RestoreSessionStatus = "failed"
)

func PossibleValuesForRestoreSessionStatus() []string {
	return []string{
		string(RestoreSessionStatus_Activating),
		string(RestoreSessionStatus_Active),
		string(RestoreSessionStatus_Completed),
		string(RestoreSessionStatus_CompletedWithError),
		string(RestoreSessionStatus_Draft),
		string(RestoreSessionStatus_Failed),
	}
}

func (s *RestoreSessionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestoreSessionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestoreSessionStatus(input string) (*RestoreSessionStatus, error) {
	vals := map[string]RestoreSessionStatus{
		"activating":         RestoreSessionStatus_Activating,
		"active":             RestoreSessionStatus_Active,
		"completed":          RestoreSessionStatus_Completed,
		"completedwitherror": RestoreSessionStatus_CompletedWithError,
		"draft":              RestoreSessionStatus_Draft,
		"failed":             RestoreSessionStatus_Failed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreSessionStatus(input)
	return &out, nil
}
