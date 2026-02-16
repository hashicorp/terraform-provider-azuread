package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementReportFileFormat string

const (
	DeviceManagementReportFileFormat_Csv  DeviceManagementReportFileFormat = "csv"
	DeviceManagementReportFileFormat_Json DeviceManagementReportFileFormat = "json"
	DeviceManagementReportFileFormat_Pdf  DeviceManagementReportFileFormat = "pdf"
)

func PossibleValuesForDeviceManagementReportFileFormat() []string {
	return []string{
		string(DeviceManagementReportFileFormat_Csv),
		string(DeviceManagementReportFileFormat_Json),
		string(DeviceManagementReportFileFormat_Pdf),
	}
}

func (s *DeviceManagementReportFileFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementReportFileFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementReportFileFormat(input string) (*DeviceManagementReportFileFormat, error) {
	vals := map[string]DeviceManagementReportFileFormat{
		"csv":  DeviceManagementReportFileFormat_Csv,
		"json": DeviceManagementReportFileFormat_Json,
		"pdf":  DeviceManagementReportFileFormat_Pdf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementReportFileFormat(input)
	return &out, nil
}
