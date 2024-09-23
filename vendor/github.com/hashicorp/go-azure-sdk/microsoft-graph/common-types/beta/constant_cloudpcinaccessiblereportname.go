package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCInaccessibleReportName string

const (
	CloudPCInaccessibleReportName_InaccessibleCloudPCReports     CloudPCInaccessibleReportName = "inaccessibleCloudPcReports"
	CloudPCInaccessibleReportName_InaccessibleCloudPCTrendReport CloudPCInaccessibleReportName = "inaccessibleCloudPcTrendReport"
)

func PossibleValuesForCloudPCInaccessibleReportName() []string {
	return []string{
		string(CloudPCInaccessibleReportName_InaccessibleCloudPCReports),
		string(CloudPCInaccessibleReportName_InaccessibleCloudPCTrendReport),
	}
}

func (s *CloudPCInaccessibleReportName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCInaccessibleReportName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCInaccessibleReportName(input string) (*CloudPCInaccessibleReportName, error) {
	vals := map[string]CloudPCInaccessibleReportName{
		"inaccessiblecloudpcreports":     CloudPCInaccessibleReportName_InaccessibleCloudPCReports,
		"inaccessiblecloudpctrendreport": CloudPCInaccessibleReportName_InaccessibleCloudPCTrendReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCInaccessibleReportName(input)
	return &out, nil
}
