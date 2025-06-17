package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRecommendationReportType string

const (
	CloudPCRecommendationReportType_CloudPCUsageCategoryReports CloudPCRecommendationReportType = "cloudPcUsageCategoryReports"
)

func PossibleValuesForCloudPCRecommendationReportType() []string {
	return []string{
		string(CloudPCRecommendationReportType_CloudPCUsageCategoryReports),
	}
}

func (s *CloudPCRecommendationReportType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRecommendationReportType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRecommendationReportType(input string) (*CloudPCRecommendationReportType, error) {
	vals := map[string]CloudPCRecommendationReportType{
		"cloudpcusagecategoryreports": CloudPCRecommendationReportType_CloudPCUsageCategoryReports,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRecommendationReportType(input)
	return &out, nil
}
