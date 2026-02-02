package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementReportStatus string

const (
	DeviceManagementReportStatus_Completed  DeviceManagementReportStatus = "completed"
	DeviceManagementReportStatus_Failed     DeviceManagementReportStatus = "failed"
	DeviceManagementReportStatus_InProgress DeviceManagementReportStatus = "inProgress"
	DeviceManagementReportStatus_NotStarted DeviceManagementReportStatus = "notStarted"
	DeviceManagementReportStatus_Unknown    DeviceManagementReportStatus = "unknown"
)

func PossibleValuesForDeviceManagementReportStatus() []string {
	return []string{
		string(DeviceManagementReportStatus_Completed),
		string(DeviceManagementReportStatus_Failed),
		string(DeviceManagementReportStatus_InProgress),
		string(DeviceManagementReportStatus_NotStarted),
		string(DeviceManagementReportStatus_Unknown),
	}
}

func (s *DeviceManagementReportStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementReportStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementReportStatus(input string) (*DeviceManagementReportStatus, error) {
	vals := map[string]DeviceManagementReportStatus{
		"completed":  DeviceManagementReportStatus_Completed,
		"failed":     DeviceManagementReportStatus_Failed,
		"inprogress": DeviceManagementReportStatus_InProgress,
		"notstarted": DeviceManagementReportStatus_NotStarted,
		"unknown":    DeviceManagementReportStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementReportStatus(input)
	return &out, nil
}
