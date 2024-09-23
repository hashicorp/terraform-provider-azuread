package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPerformanceReportName string

const (
	CloudPCPerformanceReportName_PerformanceTrendReport CloudPCPerformanceReportName = "performanceTrendReport"
)

func PossibleValuesForCloudPCPerformanceReportName() []string {
	return []string{
		string(CloudPCPerformanceReportName_PerformanceTrendReport),
	}
}

func (s *CloudPCPerformanceReportName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPerformanceReportName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPerformanceReportName(input string) (*CloudPCPerformanceReportName, error) {
	vals := map[string]CloudPCPerformanceReportName{
		"performancetrendreport": CloudPCPerformanceReportName_PerformanceTrendReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPerformanceReportName(input)
	return &out, nil
}
