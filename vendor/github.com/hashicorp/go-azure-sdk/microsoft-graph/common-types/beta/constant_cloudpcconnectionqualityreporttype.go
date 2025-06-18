package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectionQualityReportType string

const (
	CloudPCConnectionQualityReportType_RegionalConnectionQualityInsightsReport CloudPCConnectionQualityReportType = "regionalConnectionQualityInsightsReport"
	CloudPCConnectionQualityReportType_RegionalConnectionQualityTrendReport    CloudPCConnectionQualityReportType = "regionalConnectionQualityTrendReport"
	CloudPCConnectionQualityReportType_RemoteConnectionQualityReport           CloudPCConnectionQualityReportType = "remoteConnectionQualityReport"
)

func PossibleValuesForCloudPCConnectionQualityReportType() []string {
	return []string{
		string(CloudPCConnectionQualityReportType_RegionalConnectionQualityInsightsReport),
		string(CloudPCConnectionQualityReportType_RegionalConnectionQualityTrendReport),
		string(CloudPCConnectionQualityReportType_RemoteConnectionQualityReport),
	}
}

func (s *CloudPCConnectionQualityReportType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectionQualityReportType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectionQualityReportType(input string) (*CloudPCConnectionQualityReportType, error) {
	vals := map[string]CloudPCConnectionQualityReportType{
		"regionalconnectionqualityinsightsreport": CloudPCConnectionQualityReportType_RegionalConnectionQualityInsightsReport,
		"regionalconnectionqualitytrendreport":    CloudPCConnectionQualityReportType_RegionalConnectionQualityTrendReport,
		"remoteconnectionqualityreport":           CloudPCConnectionQualityReportType_RemoteConnectionQualityReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectionQualityReportType(input)
	return &out, nil
}
