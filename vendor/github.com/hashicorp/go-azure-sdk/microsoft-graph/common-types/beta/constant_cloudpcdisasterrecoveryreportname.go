package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDisasterRecoveryReportName string

const (
	CloudPCDisasterRecoveryReportName_CrossRegionDisasterRecoveryReport CloudPCDisasterRecoveryReportName = "crossRegionDisasterRecoveryReport"
	CloudPCDisasterRecoveryReportName_DisasterRecoveryReport            CloudPCDisasterRecoveryReportName = "disasterRecoveryReport"
)

func PossibleValuesForCloudPCDisasterRecoveryReportName() []string {
	return []string{
		string(CloudPCDisasterRecoveryReportName_CrossRegionDisasterRecoveryReport),
		string(CloudPCDisasterRecoveryReportName_DisasterRecoveryReport),
	}
}

func (s *CloudPCDisasterRecoveryReportName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDisasterRecoveryReportName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDisasterRecoveryReportName(input string) (*CloudPCDisasterRecoveryReportName, error) {
	vals := map[string]CloudPCDisasterRecoveryReportName{
		"crossregiondisasterrecoveryreport": CloudPCDisasterRecoveryReportName_CrossRegionDisasterRecoveryReport,
		"disasterrecoveryreport":            CloudPCDisasterRecoveryReportName_DisasterRecoveryReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDisasterRecoveryReportName(input)
	return &out, nil
}
