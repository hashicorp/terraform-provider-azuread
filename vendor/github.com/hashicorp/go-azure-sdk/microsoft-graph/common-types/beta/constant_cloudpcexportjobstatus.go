package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJobStatus string

const (
	CloudPCExportJobStatus_Completed  CloudPCExportJobStatus = "completed"
	CloudPCExportJobStatus_Failed     CloudPCExportJobStatus = "failed"
	CloudPCExportJobStatus_InProgress CloudPCExportJobStatus = "inProgress"
	CloudPCExportJobStatus_NotStarted CloudPCExportJobStatus = "notStarted"
)

func PossibleValuesForCloudPCExportJobStatus() []string {
	return []string{
		string(CloudPCExportJobStatus_Completed),
		string(CloudPCExportJobStatus_Failed),
		string(CloudPCExportJobStatus_InProgress),
		string(CloudPCExportJobStatus_NotStarted),
	}
}

func (s *CloudPCExportJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCExportJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCExportJobStatus(input string) (*CloudPCExportJobStatus, error) {
	vals := map[string]CloudPCExportJobStatus{
		"completed":  CloudPCExportJobStatus_Completed,
		"failed":     CloudPCExportJobStatus_Failed,
		"inprogress": CloudPCExportJobStatus_InProgress,
		"notstarted": CloudPCExportJobStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCExportJobStatus(input)
	return &out, nil
}
