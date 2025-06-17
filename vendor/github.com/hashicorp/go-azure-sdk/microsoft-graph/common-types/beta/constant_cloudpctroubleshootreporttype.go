package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCTroubleshootReportType string

const (
	CloudPCTroubleshootReportType_TroubleshootDetailsReport    CloudPCTroubleshootReportType = "troubleshootDetailsReport"
	CloudPCTroubleshootReportType_TroubleshootIssueCountReport CloudPCTroubleshootReportType = "troubleshootIssueCountReport"
	CloudPCTroubleshootReportType_TroubleshootRegionalReport   CloudPCTroubleshootReportType = "troubleshootRegionalReport"
	CloudPCTroubleshootReportType_TroubleshootTrendCountReport CloudPCTroubleshootReportType = "troubleshootTrendCountReport"
)

func PossibleValuesForCloudPCTroubleshootReportType() []string {
	return []string{
		string(CloudPCTroubleshootReportType_TroubleshootDetailsReport),
		string(CloudPCTroubleshootReportType_TroubleshootIssueCountReport),
		string(CloudPCTroubleshootReportType_TroubleshootRegionalReport),
		string(CloudPCTroubleshootReportType_TroubleshootTrendCountReport),
	}
}

func (s *CloudPCTroubleshootReportType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCTroubleshootReportType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCTroubleshootReportType(input string) (*CloudPCTroubleshootReportType, error) {
	vals := map[string]CloudPCTroubleshootReportType{
		"troubleshootdetailsreport":    CloudPCTroubleshootReportType_TroubleshootDetailsReport,
		"troubleshootissuecountreport": CloudPCTroubleshootReportType_TroubleshootIssueCountReport,
		"troubleshootregionalreport":   CloudPCTroubleshootReportType_TroubleshootRegionalReport,
		"troubleshoottrendcountreport": CloudPCTroubleshootReportType_TroubleshootTrendCountReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCTroubleshootReportType(input)
	return &out, nil
}
