package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreArtifactsBulkRequestStatus string

const (
	RestoreArtifactsBulkRequestStatus_Active              RestoreArtifactsBulkRequestStatus = "active"
	RestoreArtifactsBulkRequestStatus_Completed           RestoreArtifactsBulkRequestStatus = "completed"
	RestoreArtifactsBulkRequestStatus_CompletedWithErrors RestoreArtifactsBulkRequestStatus = "completedWithErrors"
	RestoreArtifactsBulkRequestStatus_Unknown             RestoreArtifactsBulkRequestStatus = "unknown"
)

func PossibleValuesForRestoreArtifactsBulkRequestStatus() []string {
	return []string{
		string(RestoreArtifactsBulkRequestStatus_Active),
		string(RestoreArtifactsBulkRequestStatus_Completed),
		string(RestoreArtifactsBulkRequestStatus_CompletedWithErrors),
		string(RestoreArtifactsBulkRequestStatus_Unknown),
	}
}

func (s *RestoreArtifactsBulkRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestoreArtifactsBulkRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestoreArtifactsBulkRequestStatus(input string) (*RestoreArtifactsBulkRequestStatus, error) {
	vals := map[string]RestoreArtifactsBulkRequestStatus{
		"active":              RestoreArtifactsBulkRequestStatus_Active,
		"completed":           RestoreArtifactsBulkRequestStatus_Completed,
		"completedwitherrors": RestoreArtifactsBulkRequestStatus_CompletedWithErrors,
		"unknown":             RestoreArtifactsBulkRequestStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreArtifactsBulkRequestStatus(input)
	return &out, nil
}
