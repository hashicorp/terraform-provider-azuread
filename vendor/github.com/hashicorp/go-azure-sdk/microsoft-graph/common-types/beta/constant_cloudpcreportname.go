package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReportName string

const (
	CloudPCReportName_ActionStatusReport                          CloudPCReportName = "actionStatusReport"
	CloudPCReportName_BulkActionStatusReport                      CloudPCReportName = "bulkActionStatusReport"
	CloudPCReportName_CloudPCInsightReport                        CloudPCReportName = "cloudPcInsightReport"
	CloudPCReportName_CloudPCUsageCategoryReports                 CloudPCReportName = "cloudPcUsageCategoryReports"
	CloudPCReportName_CrossRegionDisasterRecoveryReport           CloudPCReportName = "crossRegionDisasterRecoveryReport"
	CloudPCReportName_DailyAggregatedRemoteConnectionReports      CloudPCReportName = "dailyAggregatedRemoteConnectionReports"
	CloudPCReportName_FrontlineLicenseHourlyUsageReport           CloudPCReportName = "frontlineLicenseHourlyUsageReport"
	CloudPCReportName_FrontlineLicenseUsageRealTimeReport         CloudPCReportName = "frontlineLicenseUsageRealTimeReport"
	CloudPCReportName_FrontlineLicenseUsageReport                 CloudPCReportName = "frontlineLicenseUsageReport"
	CloudPCReportName_FrontlineRealtimeUserConnectionsReport      CloudPCReportName = "frontlineRealtimeUserConnectionsReport"
	CloudPCReportName_InaccessibleCloudPCReports                  CloudPCReportName = "inaccessibleCloudPcReports"
	CloudPCReportName_InaccessibleCloudPCTrendReport              CloudPCReportName = "inaccessibleCloudPcTrendReport"
	CloudPCReportName_NoLicenseAvailableConnectivityFailureReport CloudPCReportName = "noLicenseAvailableConnectivityFailureReport"
	CloudPCReportName_PerformanceTrendReport                      CloudPCReportName = "performanceTrendReport"
	CloudPCReportName_RawRemoteConnectionReports                  CloudPCReportName = "rawRemoteConnectionReports"
	CloudPCReportName_RegionalConnectionQualityInsightsReport     CloudPCReportName = "regionalConnectionQualityInsightsReport"
	CloudPCReportName_RegionalConnectionQualityTrendReport        CloudPCReportName = "regionalConnectionQualityTrendReport"
	CloudPCReportName_RegionalInaccessibleCloudPCTrendReport      CloudPCReportName = "regionalInaccessibleCloudPcTrendReport"
	CloudPCReportName_RemoteConnectionHistoricalReports           CloudPCReportName = "remoteConnectionHistoricalReports"
	CloudPCReportName_RemoteConnectionQualityReport               CloudPCReportName = "remoteConnectionQualityReport"
	CloudPCReportName_RemoteConnectionQualityReports              CloudPCReportName = "remoteConnectionQualityReports"
	CloudPCReportName_TotalAggregatedRemoteConnectionReports      CloudPCReportName = "totalAggregatedRemoteConnectionReports"
	CloudPCReportName_TroubleshootDetailsReport                   CloudPCReportName = "troubleshootDetailsReport"
	CloudPCReportName_TroubleshootIssueCountReport                CloudPCReportName = "troubleshootIssueCountReport"
	CloudPCReportName_TroubleshootRegionalReport                  CloudPCReportName = "troubleshootRegionalReport"
	CloudPCReportName_TroubleshootTrendCountReport                CloudPCReportName = "troubleshootTrendCountReport"
)

func PossibleValuesForCloudPCReportName() []string {
	return []string{
		string(CloudPCReportName_ActionStatusReport),
		string(CloudPCReportName_BulkActionStatusReport),
		string(CloudPCReportName_CloudPCInsightReport),
		string(CloudPCReportName_CloudPCUsageCategoryReports),
		string(CloudPCReportName_CrossRegionDisasterRecoveryReport),
		string(CloudPCReportName_DailyAggregatedRemoteConnectionReports),
		string(CloudPCReportName_FrontlineLicenseHourlyUsageReport),
		string(CloudPCReportName_FrontlineLicenseUsageRealTimeReport),
		string(CloudPCReportName_FrontlineLicenseUsageReport),
		string(CloudPCReportName_FrontlineRealtimeUserConnectionsReport),
		string(CloudPCReportName_InaccessibleCloudPCReports),
		string(CloudPCReportName_InaccessibleCloudPCTrendReport),
		string(CloudPCReportName_NoLicenseAvailableConnectivityFailureReport),
		string(CloudPCReportName_PerformanceTrendReport),
		string(CloudPCReportName_RawRemoteConnectionReports),
		string(CloudPCReportName_RegionalConnectionQualityInsightsReport),
		string(CloudPCReportName_RegionalConnectionQualityTrendReport),
		string(CloudPCReportName_RegionalInaccessibleCloudPCTrendReport),
		string(CloudPCReportName_RemoteConnectionHistoricalReports),
		string(CloudPCReportName_RemoteConnectionQualityReport),
		string(CloudPCReportName_RemoteConnectionQualityReports),
		string(CloudPCReportName_TotalAggregatedRemoteConnectionReports),
		string(CloudPCReportName_TroubleshootDetailsReport),
		string(CloudPCReportName_TroubleshootIssueCountReport),
		string(CloudPCReportName_TroubleshootRegionalReport),
		string(CloudPCReportName_TroubleshootTrendCountReport),
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
		"bulkactionstatusreport":                      CloudPCReportName_BulkActionStatusReport,
		"cloudpcinsightreport":                        CloudPCReportName_CloudPCInsightReport,
		"cloudpcusagecategoryreports":                 CloudPCReportName_CloudPCUsageCategoryReports,
		"crossregiondisasterrecoveryreport":           CloudPCReportName_CrossRegionDisasterRecoveryReport,
		"dailyaggregatedremoteconnectionreports":      CloudPCReportName_DailyAggregatedRemoteConnectionReports,
		"frontlinelicensehourlyusagereport":           CloudPCReportName_FrontlineLicenseHourlyUsageReport,
		"frontlinelicenseusagerealtimereport":         CloudPCReportName_FrontlineLicenseUsageRealTimeReport,
		"frontlinelicenseusagereport":                 CloudPCReportName_FrontlineLicenseUsageReport,
		"frontlinerealtimeuserconnectionsreport":      CloudPCReportName_FrontlineRealtimeUserConnectionsReport,
		"inaccessiblecloudpcreports":                  CloudPCReportName_InaccessibleCloudPCReports,
		"inaccessiblecloudpctrendreport":              CloudPCReportName_InaccessibleCloudPCTrendReport,
		"nolicenseavailableconnectivityfailurereport": CloudPCReportName_NoLicenseAvailableConnectivityFailureReport,
		"performancetrendreport":                      CloudPCReportName_PerformanceTrendReport,
		"rawremoteconnectionreports":                  CloudPCReportName_RawRemoteConnectionReports,
		"regionalconnectionqualityinsightsreport":     CloudPCReportName_RegionalConnectionQualityInsightsReport,
		"regionalconnectionqualitytrendreport":        CloudPCReportName_RegionalConnectionQualityTrendReport,
		"regionalinaccessiblecloudpctrendreport":      CloudPCReportName_RegionalInaccessibleCloudPCTrendReport,
		"remoteconnectionhistoricalreports":           CloudPCReportName_RemoteConnectionHistoricalReports,
		"remoteconnectionqualityreport":               CloudPCReportName_RemoteConnectionQualityReport,
		"remoteconnectionqualityreports":              CloudPCReportName_RemoteConnectionQualityReports,
		"totalaggregatedremoteconnectionreports":      CloudPCReportName_TotalAggregatedRemoteConnectionReports,
		"troubleshootdetailsreport":                   CloudPCReportName_TroubleshootDetailsReport,
		"troubleshootissuecountreport":                CloudPCReportName_TroubleshootIssueCountReport,
		"troubleshootregionalreport":                  CloudPCReportName_TroubleshootRegionalReport,
		"troubleshoottrendcountreport":                CloudPCReportName_TroubleshootTrendCountReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCReportName(input)
	return &out, nil
}
