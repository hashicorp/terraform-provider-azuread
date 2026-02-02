package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCFrontlineReportType string

const (
	CloudPCFrontlineReportType_ConnectedUserRealtimeReport                 CloudPCFrontlineReportType = "connectedUserRealtimeReport"
	CloudPCFrontlineReportType_LicenseHourlyUsageReport                    CloudPCFrontlineReportType = "licenseHourlyUsageReport"
	CloudPCFrontlineReportType_LicenseUsageRealTimeReport                  CloudPCFrontlineReportType = "licenseUsageRealTimeReport"
	CloudPCFrontlineReportType_LicenseUsageReport                          CloudPCFrontlineReportType = "licenseUsageReport"
	CloudPCFrontlineReportType_NoLicenseAvailableConnectivityFailureReport CloudPCFrontlineReportType = "noLicenseAvailableConnectivityFailureReport"
)

func PossibleValuesForCloudPCFrontlineReportType() []string {
	return []string{
		string(CloudPCFrontlineReportType_ConnectedUserRealtimeReport),
		string(CloudPCFrontlineReportType_LicenseHourlyUsageReport),
		string(CloudPCFrontlineReportType_LicenseUsageRealTimeReport),
		string(CloudPCFrontlineReportType_LicenseUsageReport),
		string(CloudPCFrontlineReportType_NoLicenseAvailableConnectivityFailureReport),
	}
}

func (s *CloudPCFrontlineReportType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCFrontlineReportType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCFrontlineReportType(input string) (*CloudPCFrontlineReportType, error) {
	vals := map[string]CloudPCFrontlineReportType{
		"connecteduserrealtimereport":                 CloudPCFrontlineReportType_ConnectedUserRealtimeReport,
		"licensehourlyusagereport":                    CloudPCFrontlineReportType_LicenseHourlyUsageReport,
		"licenseusagerealtimereport":                  CloudPCFrontlineReportType_LicenseUsageRealTimeReport,
		"licenseusagereport":                          CloudPCFrontlineReportType_LicenseUsageReport,
		"nolicenseavailableconnectivityfailurereport": CloudPCFrontlineReportType_NoLicenseAvailableConnectivityFailureReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCFrontlineReportType(input)
	return &out, nil
}
