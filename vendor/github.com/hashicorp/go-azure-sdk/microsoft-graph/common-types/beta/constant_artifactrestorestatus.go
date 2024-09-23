package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactRestoreStatus string

const (
	ArtifactRestoreStatus_Added      ArtifactRestoreStatus = "added"
	ArtifactRestoreStatus_Failed     ArtifactRestoreStatus = "failed"
	ArtifactRestoreStatus_InProgress ArtifactRestoreStatus = "inProgress"
	ArtifactRestoreStatus_Scheduled  ArtifactRestoreStatus = "scheduled"
	ArtifactRestoreStatus_Scheduling ArtifactRestoreStatus = "scheduling"
	ArtifactRestoreStatus_Succeeded  ArtifactRestoreStatus = "succeeded"
)

func PossibleValuesForArtifactRestoreStatus() []string {
	return []string{
		string(ArtifactRestoreStatus_Added),
		string(ArtifactRestoreStatus_Failed),
		string(ArtifactRestoreStatus_InProgress),
		string(ArtifactRestoreStatus_Scheduled),
		string(ArtifactRestoreStatus_Scheduling),
		string(ArtifactRestoreStatus_Succeeded),
	}
}

func (s *ArtifactRestoreStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseArtifactRestoreStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseArtifactRestoreStatus(input string) (*ArtifactRestoreStatus, error) {
	vals := map[string]ArtifactRestoreStatus{
		"added":      ArtifactRestoreStatus_Added,
		"failed":     ArtifactRestoreStatus_Failed,
		"inprogress": ArtifactRestoreStatus_InProgress,
		"scheduled":  ArtifactRestoreStatus_Scheduled,
		"scheduling": ArtifactRestoreStatus_Scheduling,
		"succeeded":  ArtifactRestoreStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArtifactRestoreStatus(input)
	return &out, nil
}
