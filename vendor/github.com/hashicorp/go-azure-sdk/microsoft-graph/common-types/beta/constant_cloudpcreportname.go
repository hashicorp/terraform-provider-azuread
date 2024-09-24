package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReportName string

const (
	CloudPCReportName_ActionStatusReport                          CloudPCReportName = "actionStatusReport"
	CloudPCReportName_CloudPCUsageCategoryReports                 CloudPCReportName = "cloudPcUsageCategoryReports"
	CloudPCReportName_CrossRegionDisasterRecoveryReport           CloudPCReportName = "crossRegionDisasterRecoveryReport"
	CloudPCReportName_DailyAggregatedRemoteConnectionReports      CloudPCReportName = "dailyAggregatedRemoteConnectionReports"
	CloudPCReportName_FrontlineLicenseUsageRealTimeReport         CloudPCReportName = "frontlineLicenseUsageRealTimeReport"
	CloudPCReportName_FrontlineLicenseUsageReport                 CloudPCReportName = "frontlineLicenseUsageReport"
	CloudPCReportName_InaccessibleCloudPCReports                  CloudPCReportName = "inaccessibleCloudPcReports"
	CloudPCReportName_InaccessibleCloudPCTrendReport              CloudPCReportName = "inaccessibleCloudPcTrendReport"
	CloudPCReportName_NoLicenseAvailableConnectivityFailureReport CloudPCReportName = "noLicenseAvailableConnectivityFailureReport"
	CloudPCReportName_PerformanceTrendReport                      CloudPCReportName = "performanceTrendReport"
	CloudPCReportName_RawRemoteConnectionReports                  CloudPCReportName = "rawRemoteConnectionReports"
	CloudPCReportName_RemoteConnectionHistoricalReports           CloudPCReportName = "remoteConnectionHistoricalReports"
	CloudPCReportName_RemoteConnectionQualityReports              CloudPCReportName = "remoteConnectionQualityReports"
	CloudPCReportName_SharedUseLicenseUsageRealTimeReport         CloudPCReportName = "sharedUseLicenseUsageRealTimeReport"
	CloudPCReportName_SharedUseLicenseUsageReport                 CloudPCReportName = "sharedUseLicenseUsageReport"
	CloudPCReportName_TotalAggregatedRemoteConnectionReports      CloudPCReportName = "totalAggregatedRemoteConnectionReports"
)

func PossibleValuesForCloudPCReportName() []string {
	return []string{
		string(CloudPCReportName_ActionStatusReport),
		string(CloudPCReportName_CloudPCUsageCategoryReports),
		string(CloudPCReportName_CrossRegionDisasterRecoveryReport),
		string(CloudPCReportName_DailyAggregatedRemoteConnectionReports),
		string(CloudPCReportName_FrontlineLicenseUsageRealTimeReport),
		string(CloudPCReportName_FrontlineLicenseUsageReport),
		string(CloudPCReportName_InaccessibleCloudPCReports),
		string(CloudPCReportName_InaccessibleCloudPCTrendReport),
		string(CloudPCReportName_NoLicenseAvailableConnectivityFailureReport),
		string(CloudPCReportName_PerformanceTrendReport),
		string(CloudPCReportName_RawRemoteConnectionReports),
		string(CloudPCReportName_RemoteConnectionHistoricalReports),
		string(CloudPCReportName_RemoteConnectionQualityReports),
		string(CloudPCReportName_SharedUseLicenseUsageRealTimeReport),
		string(CloudPCReportName_SharedUseLicenseUsageReport),
		string(CloudPCReportName_TotalAggregatedRemoteConnectionReports),
	}
}

func (s *CloudPCReportName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCReportName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCReportName(input string) (*CloudPCReportName, error) {
	vals := map[string]CloudPCReportName{
		"actionstatusreport":                          CloudPCReportName_ActionStatusReport,
		"cloudpcusagecategoryreports":                 CloudPCReportName_CloudPCUsageCategoryReports,
		"crossregiondisasterrecoveryreport":           CloudPCReportName_CrossRegionDisasterRecoveryReport,
		"dailyaggregatedremoteconnectionreports":      CloudPCReportName_DailyAggregatedRemoteConnectionReports,
		"frontlinelicenseusagerealtimereport":         CloudPCReportName_FrontlineLicenseUsageRealTimeReport,
		"frontlinelicenseusagereport":                 CloudPCReportName_FrontlineLicenseUsageReport,
		"inaccessiblecloudpcreports":                  CloudPCReportName_InaccessibleCloudPCReports,
		"inaccessiblecloudpctrendreport":              CloudPCReportName_InaccessibleCloudPCTrendReport,
		"nolicenseavailableconnectivityfailurereport": CloudPCReportName_NoLicenseAvailableConnectivityFailureReport,
		"performancetrendreport":                      CloudPCReportName_PerformanceTrendReport,
		"rawremoteconnectionreports":                  CloudPCReportName_RawRemoteConnectionReports,
		"remoteconnectionhistoricalreports":           CloudPCReportName_RemoteConnectionHistoricalReports,
		"remoteconnectionqualityreports":              CloudPCReportName_RemoteConnectionQualityReports,
		"shareduselicenseusagerealtimereport":         CloudPCReportName_SharedUseLicenseUsageRealTimeReport,
		"shareduselicenseusagereport":                 CloudPCReportName_SharedUseLicenseUsageReport,
		"totalaggregatedremoteconnectionreports":      CloudPCReportName_TotalAggregatedRemoteConnectionReports,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCReportName(input)
	return &out, nil
}
