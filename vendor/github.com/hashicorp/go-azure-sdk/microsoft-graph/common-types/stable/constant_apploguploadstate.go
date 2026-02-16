package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppLogUploadState string

const (
	AppLogUploadState_Completed AppLogUploadState = "completed"
	AppLogUploadState_Failed    AppLogUploadState = "failed"
	AppLogUploadState_Pending   AppLogUploadState = "pending"
)

func PossibleValuesForAppLogUploadState() []string {
	return []string{
		string(AppLogUploadState_Completed),
		string(AppLogUploadState_Failed),
		string(AppLogUploadState_Pending),
	}
}

func (s *AppLogUploadState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppLogUploadState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppLogUploadState(input string) (*AppLogUploadState, error) {
	vals := map[string]AppLogUploadState{
		"completed": AppLogUploadState_Completed,
		"failed":    AppLogUploadState_Failed,
		"pending":   AppLogUploadState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppLogUploadState(input)
	return &out, nil
}
