package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityAuditLogRecord{}

type SecurityAuditLogRecord struct {
	// The administrative units tagged to an audit log record.
	AdministrativeUnits *[]string `json:"administrativeUnits,omitempty"`

	// A JSON object that contains the actual audit log data.
	AuditData SecurityAuditData `json:"auditData"`

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
	AuditLogRecordType *SecurityAuditLogRecordType `json:"auditLogRecordType,omitempty"`

	// The IP address of the device used when the activity was logged. The IP address is displayed in either an IPv4 or IPv6
	// address format.
	ClientIp nullable.Type[string] `json:"clientIp,omitempty"`

	// The date and time in UTC when the user performed the activity.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// For Exchange admin audit logging, the name of the object modified by the cmdlet. For SharePoint activity, the full
	// URL path name of the file or folder accessed by a user. For Microsoft Entra activity, the name of the user account
	// that was modified.
	ObjectId nullable.Type[string] `json:"objectId,omitempty"`

	// The name of the user or admin activity.
	Operation nullable.Type[string] `json:"operation,omitempty"`

	// The GUID for your organization.
	OrganizationId nullable.Type[string] `json:"organizationId,omitempty"`

	// The Microsoft 365 service where the activity occurred.
	Service nullable.Type[string] `json:"service,omitempty"`

	// The user who performed the action (specified in the Operation property) that resulted in the record being logged.
	// Audit records for activity performed by system accounts (such as SHAREPOINT/system or NT AUTHORITY/SYSTEM) are also
	// included in the audit log. Another common value for the UserId property is app@sharepoint. It indicates that the
	// 'user' who performed the activity was an application with the necessary permissions in SharePoint to perform
	// organization-wide actions (such as searching a SharePoint site or OneDrive account) on behalf of a user, admin, or
	// service.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// UPN of the user who performed the action.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// The type of user that performed the operation. The possible values are: regular, reserved, admin, dcAdmin, system,
	// application, servicePrincipal, customPolicy, systemPolicy, partnerTechnician, guest, unknownFutureValue.
	UserType *SecurityAuditLogUserType `json:"userType,omitempty"`

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

func (s SecurityAuditLogRecord) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityAuditLogRecord{}

func (s SecurityAuditLogRecord) MarshalJSON() ([]byte, error) {
	type wrapper SecurityAuditLogRecord
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityAuditLogRecord: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityAuditLogRecord: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.auditLogRecord"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityAuditLogRecord: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityAuditLogRecord{}

func (s *SecurityAuditLogRecord) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AdministrativeUnits *[]string                   `json:"administrativeUnits,omitempty"`
		AuditLogRecordType  *SecurityAuditLogRecordType `json:"auditLogRecordType,omitempty"`
		ClientIp            nullable.Type[string]       `json:"clientIp,omitempty"`
		CreatedDateTime     nullable.Type[string]       `json:"createdDateTime,omitempty"`
		ObjectId            nullable.Type[string]       `json:"objectId,omitempty"`
		Operation           nullable.Type[string]       `json:"operation,omitempty"`
		OrganizationId      nullable.Type[string]       `json:"organizationId,omitempty"`
		Service             nullable.Type[string]       `json:"service,omitempty"`
		UserId              nullable.Type[string]       `json:"userId,omitempty"`
		UserPrincipalName   nullable.Type[string]       `json:"userPrincipalName,omitempty"`
		UserType            *SecurityAuditLogUserType   `json:"userType,omitempty"`
		Id                  *string                     `json:"id,omitempty"`
		ODataId             *string                     `json:"@odata.id,omitempty"`
		ODataType           *string                     `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AdministrativeUnits = decoded.AdministrativeUnits
	s.AuditLogRecordType = decoded.AuditLogRecordType
	s.ClientIp = decoded.ClientIp
	s.CreatedDateTime = decoded.CreatedDateTime
	s.ObjectId = decoded.ObjectId
	s.Operation = decoded.Operation
	s.OrganizationId = decoded.OrganizationId
	s.Service = decoded.Service
	s.UserId = decoded.UserId
	s.UserPrincipalName = decoded.UserPrincipalName
	s.UserType = decoded.UserType
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityAuditLogRecord into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["auditData"]; ok {
		impl, err := UnmarshalSecurityAuditDataImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AuditData' for 'SecurityAuditLogRecord': %+v", err)
		}
		s.AuditData = impl
	}

	return nil
}
