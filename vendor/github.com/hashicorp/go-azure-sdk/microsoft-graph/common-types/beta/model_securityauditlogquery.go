package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityAuditLogQuery{}

type SecurityAuditLogQuery struct {
	// The administrative units tagged to an audit log record.
	AdministrativeUnitIdFilters *[]string `json:"administrativeUnitIdFilters,omitempty"`

	// The display name of the saved audit log query.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The end date of the date range in the query.
	FilterEndDateTime nullable.Type[string] `json:"filterEndDateTime,omitempty"`

	// The start date of the date range in the query.
	FilterStartDateTime nullable.Type[string] `json:"filterStartDateTime,omitempty"`

	// The IP address of the device that was used when the activity was logged.
	IPAddressFilters *[]string `json:"ipAddressFilters,omitempty"`

	// Free text field to search non-indexed properties of the audit log.
	KeywordFilter nullable.Type[string] `json:"keywordFilter,omitempty"`

	// For SharePoint and OneDrive for Business activity, the full path name of the file or folder accessed by the user. For
	// Exchange admin audit logging, the name of the object that was modified by the cmdlet.
	ObjectIdFilters *[]string `json:"objectIdFilters,omitempty"`

	// The name of the user or admin activity. For a description of the most common operations/activities, see Search the
	// audit log in the Office 365 Protection Center.
	OperationFilters *[]string `json:"operationFilters,omitempty"`

	// The type of operation indicated by the record. The possible values are: exchangeAdmin, exchangeItem,
	// exchangeItemGroup, sharePoint, syntheticProbe, sharePointFileOperation, oneDrive, azureActiveDirectory,
	// azureActiveDirectoryAccountLogon, dataCenterSecurityCmdlet, complianceDLPSharePoint, sway, complianceDLPExchange,
	// sharePointSharingOperation, azureActiveDirectoryStsLogon, skypeForBusinessPSTNUsage, skypeForBusinessUsersBlocked,
	// securityComplianceCenterEOPCmdlet, exchangeAggregatedOperation, powerBIAudit, crm, yammer, skypeForBusinessCmdlets,
	// discovery, microsoftTeams, threatIntelligence, mailSubmission, microsoftFlow, aeD, microsoftStream,
	// complianceDLPSharePointClassification, threatFinder, project, sharePointListOperation, sharePointCommentOperation,
	// dataGovernance, kaizala, securityComplianceAlerts, threatIntelligenceUrl, securityComplianceInsights, mipLabel,
	// workplaceAnalytics, powerAppsApp, powerAppsPlan, threatIntelligenceAtpContent, labelContentExplorer, teamsHealthcare,
	// exchangeItemAggregated, hygieneEvent, dataInsightsRestApiAudit, informationBarrierPolicyApplication,
	// sharePointListItemOperation, sharePointContentTypeOperation, sharePointFieldOperation, microsoftTeamsAdmin, hrSignal,
	// microsoftTeamsDevice, microsoftTeamsAnalytics, informationWorkerProtection, campaign, dlpEndpoint, airInvestigation,
	// quarantine, microsoftForms, applicationAudit, complianceSupervisionExchange, customerKeyServiceEncryption,
	// officeNative, mipAutoLabelSharePointItem, mipAutoLabelSharePointPolicyLocation, microsoftTeamsShifts, secureScore,
	// mipAutoLabelExchangeItem, cortanaBriefing, search, wdatpAlerts, powerPlatformAdminDlp, powerPlatformAdminEnvironment,
	// mdatpAudit, sensitivityLabelPolicyMatch, sensitivityLabelAction, sensitivityLabeledFileAction, attackSim,
	// airManualInvestigation, securityComplianceRBAC, userTraining, airAdminActionInvestigation, mstic,
	// physicalBadgingSignal, teamsEasyApprovals, aipDiscover, aipSensitivityLabelAction, aipProtectionAction,
	// aipFileDeleted, aipHeartBeat, mcasAlerts, onPremisesFileShareScannerDlp, onPremisesSharePointScannerDlp,
	// exchangeSearch, sharePointSearch, privacyDataMinimization, labelAnalyticsAggregate, myAnalyticsSettings,
	// securityComplianceUserChange, complianceDLPExchangeClassification, complianceDLPEndpoint, mipExactDataMatch,
	// msdeResponseActions, msdeGeneralSettings, msdeIndicatorsSettings, ms365DCustomDetection, msdeRolesSettings,
	// mapgAlerts, mapgPolicy, mapgRemediation, privacyRemediationAction, privacyDigestEmail,
	// mipAutoLabelSimulationProgress, mipAutoLabelSimulationCompletion, mipAutoLabelProgressFeedback,
	// dlpSensitiveInformationType, mipAutoLabelSimulationStatistics, largeContentMetadata, microsoft365Group,
	// cdpMlInferencingResult, filteringMailMetadata, cdpClassificationMailItem, cdpClassificationDocument,
	// officeScriptsRunAction, filteringPostMailDeliveryAction, cdpUnifiedFeedback, tenantAllowBlockList,
	// consumptionResource, healthcareSignal, dlpImportResult, cdpCompliancePolicyExecution, multiStageDisposition,
	// privacyDataMatch, filteringDocMetadata, filteringEmailFeatures, powerBIDlp, filteringUrlInfo,
	// filteringAttachmentInfo, coreReportingSettings, complianceConnector, powerPlatformLockboxResourceAccessRequest,
	// powerPlatformLockboxResourceCommand, cdpPredictiveCodingLabel, cdpCompliancePolicyUserFeedback,
	// webpageActivityEndpoint, omePortal, cmImprovementActionChange, filteringUrlClick, mipLabelAnalyticsAuditRecord,
	// filteringEntityEvent, filteringRuleHits, filteringMailSubmission, labelExplorer, microsoftManagedServicePlatform,
	// powerPlatformServiceActivity, scorePlatformGenericAuditRecord, filteringTimeTravelDocMetadata, alert, alertStatus,
	// alertIncident, incidentStatus, case, caseInvestigation, recordsManagement, privacyRemediation, dataShareOperation,
	// cdpDlpSensitive, ehrConnector, filteringMailGradingResult, publicFolder, privacyTenantAuditHistoryRecord,
	// aipScannerDiscoverEvent, eduDataLakeDownloadOperation, m365ComplianceConnector, microsoftGraphDataConnectOperation,
	// microsoftPurview, filteringEmailContentFeatures, powerPagesSite, powerAppsResource, plannerPlan, plannerCopyPlan,
	// plannerTask, plannerRoster, plannerPlanList, plannerTaskList, plannerTenantSettings, projectForTheWebProject,
	// projectForTheWebTask, projectForTheWebRoadmap, projectForTheWebRoadmapItem, projectForTheWebProjectSettings,
	// projectForTheWebRoadmapSettings, quarantineMetadata, microsoftTodoAudit, timeTravelFilteringDocMetadata,
	// teamsQuarantineMetadata, sharePointAppPermissionOperation, microsoftTeamsSensitivityLabelAction,
	// filteringTeamsMetadata, filteringTeamsUrlInfo, filteringTeamsPostDeliveryAction, mdcAssessments,
	// mdcRegulatoryComplianceStandards, mdcRegulatoryComplianceControls, mdcRegulatoryComplianceAssessments,
	// mdcSecurityConnectors, mdaDataSecuritySignal, vivaGoals, filteringRuntimeInfo, attackSimAdmin,
	// microsoftGraphDataConnectConsent, filteringAtpDetonationInfo, privacyPortal, managedTenants,
	// unifiedSimulationMatchedItem, unifiedSimulationSummary, updateQuarantineMetadata, ms365DSuppressionRule,
	// purviewDataMapOperation, filteringUrlPostClickAction, irmUserDefinedDetectionSignal, teamsUpdates,
	// plannerRosterSensitivityLabel, ms365DIncident, filteringDelistingMetadata,
	// complianceDLPSharePointClassificationExtended, microsoftDefenderForIdentityAudit, supervisoryReviewDayXInsight,
	// defenderExpertsforXDRAdmin, cdpEdgeBlockedMessage, hostedRpa, cdpContentExplorerAggregateRecord,
	// cdpHygieneAttachmentInfo, cdpHygieneSummary, cdpPostMailDeliveryAction, cdpEmailFeatures, cdpHygieneUrlInfo,
	// cdpUrlClick, cdpPackageManagerHygieneEvent, filteringDocScan, timeTravelFilteringDocScan, mapgOnboard,
	// unknownFutureValue.
	RecordTypeFilters *[]SecurityAuditLogRecordType `json:"recordTypeFilters,omitempty"`

	// An individual audit log record.
	Records *[]SecurityAuditLogRecord `json:"records,omitempty"`

	ServiceFilters *[]string `json:"serviceFilters,omitempty"`

	// Describes the current status of the query. The possible values are: notStarted, running, succeeded, failed,
	// cancelled, unknownFutureValue.
	Status *SecurityAuditLogQueryStatus `json:"status,omitempty"`

	// The UPN (user principal name) of the user who performed the action (specified in the operation property) that
	// resulted in the record being logged; for example, myname@mydomain_name.
	UserPrincipalNameFilters *[]string `json:"userPrincipalNameFilters,omitempty"`

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

func (s SecurityAuditLogQuery) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityAuditLogQuery{}

func (s SecurityAuditLogQuery) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAuditLogQuery
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAuditLogQuery: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAuditLogQuery: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.auditLogQuery"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAuditLogQuery: %+v", err)
	}

	return encoded, nil
}
