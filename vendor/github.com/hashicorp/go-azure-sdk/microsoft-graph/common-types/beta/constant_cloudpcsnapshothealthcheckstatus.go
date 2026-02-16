package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSnapshotHealthCheckStatus string

const (
	CloudPCSnapshotHealthCheckStatus_Healthy   CloudPCSnapshotHealthCheckStatus = "healthy"
	CloudPCSnapshotHealthCheckStatus_Unhealthy CloudPCSnapshotHealthCheckStatus = "unhealthy"
	CloudPCSnapshotHealthCheckStatus_Unknown   CloudPCSnapshotHealthCheckStatus = "unknown"
)

func PossibleValuesForCloudPCSnapshotHealthCheckStatus() []string {
	return []string{
		string(CloudPCSnapshotHealthCheckStatus_Healthy),
		string(CloudPCSnapshotHealthCheckStatus_Unhealthy),
		string(CloudPCSnapshotHealthCheckStatus_Unknown),
	}
}

func (s *CloudPCSnapshotHealthCheckStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSnapshotHealthCheckStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSnapshotHealthCheckStatus(input string) (*CloudPCSnapshotHealthCheckStatus, error) {
	vals := map[string]CloudPCSnapshotHealthCheckStatus{
		"healthy":   CloudPCSnapshotHealthCheckStatus_Healthy,
		"unhealthy": CloudPCSnapshotHealthCheckStatus_Unhealthy,
		"unknown":   CloudPCSnapshotHealthCheckStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSnapshotHealthCheckStatus(input)
	return &out, nil
}
