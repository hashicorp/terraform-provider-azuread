package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCExportJob{}

type CloudPCExportJob struct {
	// The date and time when the export job expires.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// The status of the export job. The possible values are: notStarted, inProgress, completed, unknownFutureValue.
	// Read-only.
	ExportJobStatus *CloudPCExportJobStatus `json:"exportJobStatus,omitempty"`

	// The storage account URL of the exported report. It can be used to download the file.
	ExportUrl nullable.Type[string] `json:"exportUrl,omitempty"`

	// The filter applied on the report.
	Filter nullable.Type[string] `json:"filter,omitempty"`

	// The format of the exported report.
	Format nullable.Type[string] `json:"format,omitempty"`

	// The report name. The possible values are: remoteConnectionHistoricalReports, dailyAggregatedRemoteConnectionReports,
	// totalAggregatedRemoteConnectionReports, sharedUseLicenseUsageReport, sharedUseLicenseUsageRealTimeReport,
	// unknownFutureValue, noLicenseAvailableConnectivityFailureReport, frontlineLicenseUsageReport,
	// frontlineLicenseUsageRealTimeReport, remoteConnectionQualityReports, inaccessibleCloudPcReports, actionStatusReport,
	// rawRemoteConnectionReports, cloudPcUsageCategoryReports, crossRegionDisasterRecoveryReport,
	// regionalConnectionQualityTrendReport, regionalConnectionQualityInsightsReport, remoteConnectionQualityReport,
	// bulkActionStatusReport, cloudPcInsightReport, regionalInaccessibleCloudPcTrendReport, troubleshootDetailsReport,
	// troubleshootTrendCountReport, troubleshootRegionalReport, troubleshootIssueCountReport. Use the Prefer:
	// include-unknown-enum-members request header to get the following values in this evolvable enum:
	// noLicenseAvailableConnectivityFailureReport, frontlineLicenseUsageReport, frontlineLicenseUsageRealTimeReport,
	// remoteConnectionQualityReports, inaccessibleCloudPcReports, rawRemoteConnectionReports, cloudPcUsageCategoryReports,
	// crossRegionDisasterRecoveryReport, cloudPcInsightReport, regionalInaccessibleCloudPcTrendReport,,
	// troubleshootDetailsReport, troubleshootTrendCountReport, troubleshootRegionalReport, troubleshootIssueCountReport.
	ReportName *CloudPCReportName `json:"reportName,omitempty"`

	// The date and time when the export job was requested.
	RequestDateTime nullable.Type[string] `json:"requestDateTime,omitempty"`

	// The selected columns of the report.
	Select *[]string `json:"select,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s CloudPCExportJob) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCExportJob{}

func (s CloudPCExportJob) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCExportJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCExportJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCExportJob: %+v", err)
	}

	delete(decoded, "exportJobStatus")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcExportJob"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCExportJob: %+v", err)
	}

	return encoded, nil
}
