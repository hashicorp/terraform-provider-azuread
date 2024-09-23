package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAuditLogRecordType string

const (
	SecurityAuditLogRecordType_AeD                                           SecurityAuditLogRecordType = "AeD"
	SecurityAuditLogRecordType_AipDiscover                                   SecurityAuditLogRecordType = "AipDiscover"
	SecurityAuditLogRecordType_AipFileDeleted                                SecurityAuditLogRecordType = "AipFileDeleted"
	SecurityAuditLogRecordType_AipHeartBeat                                  SecurityAuditLogRecordType = "AipHeartBeat"
	SecurityAuditLogRecordType_AipProtectionAction                           SecurityAuditLogRecordType = "AipProtectionAction"
	SecurityAuditLogRecordType_AipScannerDiscoverEvent                       SecurityAuditLogRecordType = "AipScannerDiscoverEvent"
	SecurityAuditLogRecordType_AipSensitivityLabelAction                     SecurityAuditLogRecordType = "AipSensitivityLabelAction"
	SecurityAuditLogRecordType_AirAdminActionInvestigation                   SecurityAuditLogRecordType = "AirAdminActionInvestigation"
	SecurityAuditLogRecordType_AirInvestigation                              SecurityAuditLogRecordType = "AirInvestigation"
	SecurityAuditLogRecordType_AirManualInvestigation                        SecurityAuditLogRecordType = "AirManualInvestigation"
	SecurityAuditLogRecordType_Alert                                         SecurityAuditLogRecordType = "Alert"
	SecurityAuditLogRecordType_AlertIncident                                 SecurityAuditLogRecordType = "AlertIncident"
	SecurityAuditLogRecordType_AlertStatus                                   SecurityAuditLogRecordType = "AlertStatus"
	SecurityAuditLogRecordType_ApplicationAudit                              SecurityAuditLogRecordType = "ApplicationAudit"
	SecurityAuditLogRecordType_AttackSim                                     SecurityAuditLogRecordType = "AttackSim"
	SecurityAuditLogRecordType_AttackSimAdmin                                SecurityAuditLogRecordType = "AttackSimAdmin"
	SecurityAuditLogRecordType_AzureActiveDirectory                          SecurityAuditLogRecordType = "AzureActiveDirectory"
	SecurityAuditLogRecordType_AzureActiveDirectoryAccountLogon              SecurityAuditLogRecordType = "AzureActiveDirectoryAccountLogon"
	SecurityAuditLogRecordType_AzureActiveDirectoryStsLogon                  SecurityAuditLogRecordType = "AzureActiveDirectoryStsLogon"
	SecurityAuditLogRecordType_CDPClassificationDocument                     SecurityAuditLogRecordType = "CDPClassificationDocument"
	SecurityAuditLogRecordType_CDPClassificationMailItem                     SecurityAuditLogRecordType = "CDPClassificationMailItem"
	SecurityAuditLogRecordType_CDPCompliancePolicyExecution                  SecurityAuditLogRecordType = "CDPCompliancePolicyExecution"
	SecurityAuditLogRecordType_CDPCompliancePolicyUserFeedback               SecurityAuditLogRecordType = "CDPCompliancePolicyUserFeedback"
	SecurityAuditLogRecordType_CDPEdgeBlockedMessage                         SecurityAuditLogRecordType = "CDPEdgeBlockedMessage"
	SecurityAuditLogRecordType_CDPEmailFeatures                              SecurityAuditLogRecordType = "CDPEmailFeatures"
	SecurityAuditLogRecordType_CDPHygieneAttachmentInfo                      SecurityAuditLogRecordType = "CDPHygieneAttachmentInfo"
	SecurityAuditLogRecordType_CDPHygieneSummary                             SecurityAuditLogRecordType = "CDPHygieneSummary"
	SecurityAuditLogRecordType_CDPHygieneUrlInfo                             SecurityAuditLogRecordType = "CDPHygieneUrlInfo"
	SecurityAuditLogRecordType_CDPMlInferencingResult                        SecurityAuditLogRecordType = "CDPMlInferencingResult"
	SecurityAuditLogRecordType_CDPPackageManagerHygieneEvent                 SecurityAuditLogRecordType = "CDPPackageManagerHygieneEvent"
	SecurityAuditLogRecordType_CDPPostMailDeliveryAction                     SecurityAuditLogRecordType = "CDPPostMailDeliveryAction"
	SecurityAuditLogRecordType_CDPPredictiveCodingLabel                      SecurityAuditLogRecordType = "CDPPredictiveCodingLabel"
	SecurityAuditLogRecordType_CDPUnifiedFeedback                            SecurityAuditLogRecordType = "CDPUnifiedFeedback"
	SecurityAuditLogRecordType_CDPUrlClick                                   SecurityAuditLogRecordType = "CDPUrlClick"
	SecurityAuditLogRecordType_CMImprovementActionChange                     SecurityAuditLogRecordType = "CMImprovementActionChange"
	SecurityAuditLogRecordType_CRM                                           SecurityAuditLogRecordType = "CRM"
	SecurityAuditLogRecordType_Campaign                                      SecurityAuditLogRecordType = "Campaign"
	SecurityAuditLogRecordType_Case                                          SecurityAuditLogRecordType = "Case"
	SecurityAuditLogRecordType_CaseInvestigation                             SecurityAuditLogRecordType = "CaseInvestigation"
	SecurityAuditLogRecordType_CdpColdCrawlStatus                            SecurityAuditLogRecordType = "CdpColdCrawlStatus"
	SecurityAuditLogRecordType_CdpContentExplorerAggregateRecord             SecurityAuditLogRecordType = "CdpContentExplorerAggregateRecord"
	SecurityAuditLogRecordType_CdpDlpSensitive                               SecurityAuditLogRecordType = "CdpDlpSensitive"
	SecurityAuditLogRecordType_ComplianceConnector                           SecurityAuditLogRecordType = "ComplianceConnector"
	SecurityAuditLogRecordType_ComplianceDLPEndpoint                         SecurityAuditLogRecordType = "ComplianceDLPEndpoint"
	SecurityAuditLogRecordType_ComplianceDLPExchange                         SecurityAuditLogRecordType = "ComplianceDLPExchange"
	SecurityAuditLogRecordType_ComplianceDLPExchangeClassification           SecurityAuditLogRecordType = "ComplianceDLPExchangeClassification"
	SecurityAuditLogRecordType_ComplianceDLPSharePoint                       SecurityAuditLogRecordType = "ComplianceDLPSharePoint"
	SecurityAuditLogRecordType_ComplianceDLPSharePointClassification         SecurityAuditLogRecordType = "ComplianceDLPSharePointClassification"
	SecurityAuditLogRecordType_ComplianceDLPSharePointClassificationExtended SecurityAuditLogRecordType = "ComplianceDLPSharePointClassificationExtended"
	SecurityAuditLogRecordType_ComplianceSupervisionExchange                 SecurityAuditLogRecordType = "ComplianceSupervisionExchange"
	SecurityAuditLogRecordType_ConsumptionResource                           SecurityAuditLogRecordType = "ConsumptionResource"
	SecurityAuditLogRecordType_CoreReportingSettings                         SecurityAuditLogRecordType = "CoreReportingSettings"
	SecurityAuditLogRecordType_CortanaBriefing                               SecurityAuditLogRecordType = "CortanaBriefing"
	SecurityAuditLogRecordType_CustomerKeyServiceEncryption                  SecurityAuditLogRecordType = "CustomerKeyServiceEncryption"
	SecurityAuditLogRecordType_DLPEndpoint                                   SecurityAuditLogRecordType = "DLPEndpoint"
	SecurityAuditLogRecordType_DataCenterSecurityCmdlet                      SecurityAuditLogRecordType = "DataCenterSecurityCmdlet"
	SecurityAuditLogRecordType_DataGovernance                                SecurityAuditLogRecordType = "DataGovernance"
	SecurityAuditLogRecordType_DataInsightsRestApiAudit                      SecurityAuditLogRecordType = "DataInsightsRestApiAudit"
	SecurityAuditLogRecordType_DataShareOperation                            SecurityAuditLogRecordType = "DataShareOperation"
	SecurityAuditLogRecordType_DefenderExpertsforXDRAdmin                    SecurityAuditLogRecordType = "DefenderExpertsforXDRAdmin"
	SecurityAuditLogRecordType_Discovery                                     SecurityAuditLogRecordType = "Discovery"
	SecurityAuditLogRecordType_DlpImportResult                               SecurityAuditLogRecordType = "DlpImportResult"
	SecurityAuditLogRecordType_DlpSensitiveInformationType                   SecurityAuditLogRecordType = "DlpSensitiveInformationType"
	SecurityAuditLogRecordType_EHRConnector                                  SecurityAuditLogRecordType = "EHRConnector"
	SecurityAuditLogRecordType_EduDataLakeDownloadOperation                  SecurityAuditLogRecordType = "EduDataLakeDownloadOperation"
	SecurityAuditLogRecordType_ExchangeAdmin                                 SecurityAuditLogRecordType = "ExchangeAdmin"
	SecurityAuditLogRecordType_ExchangeAggregatedOperation                   SecurityAuditLogRecordType = "ExchangeAggregatedOperation"
	SecurityAuditLogRecordType_ExchangeItem                                  SecurityAuditLogRecordType = "ExchangeItem"
	SecurityAuditLogRecordType_ExchangeItemAggregated                        SecurityAuditLogRecordType = "ExchangeItemAggregated"
	SecurityAuditLogRecordType_ExchangeItemGroup                             SecurityAuditLogRecordType = "ExchangeItemGroup"
	SecurityAuditLogRecordType_ExchangeSearch                                SecurityAuditLogRecordType = "ExchangeSearch"
	SecurityAuditLogRecordType_FilteringAtpDetonationInfo                    SecurityAuditLogRecordType = "FilteringAtpDetonationInfo"
	SecurityAuditLogRecordType_FilteringAttachmentInfo                       SecurityAuditLogRecordType = "FilteringAttachmentInfo"
	SecurityAuditLogRecordType_FilteringDelistingMetadata                    SecurityAuditLogRecordType = "FilteringDelistingMetadata"
	SecurityAuditLogRecordType_FilteringDocMetadata                          SecurityAuditLogRecordType = "FilteringDocMetadata"
	SecurityAuditLogRecordType_FilteringDocScan                              SecurityAuditLogRecordType = "FilteringDocScan"
	SecurityAuditLogRecordType_FilteringEmailContentFeatures                 SecurityAuditLogRecordType = "FilteringEmailContentFeatures"
	SecurityAuditLogRecordType_FilteringEmailFeatures                        SecurityAuditLogRecordType = "FilteringEmailFeatures"
	SecurityAuditLogRecordType_FilteringEntityEvent                          SecurityAuditLogRecordType = "FilteringEntityEvent"
	SecurityAuditLogRecordType_FilteringMailGradingResult                    SecurityAuditLogRecordType = "FilteringMailGradingResult"
	SecurityAuditLogRecordType_FilteringMailMetadata                         SecurityAuditLogRecordType = "FilteringMailMetadata"
	SecurityAuditLogRecordType_FilteringMailSubmission                       SecurityAuditLogRecordType = "FilteringMailSubmission"
	SecurityAuditLogRecordType_FilteringPostMailDeliveryAction               SecurityAuditLogRecordType = "FilteringPostMailDeliveryAction"
	SecurityAuditLogRecordType_FilteringRuleHits                             SecurityAuditLogRecordType = "FilteringRuleHits"
	SecurityAuditLogRecordType_FilteringRuntimeInfo                          SecurityAuditLogRecordType = "FilteringRuntimeInfo"
	SecurityAuditLogRecordType_FilteringTeamsMetadata                        SecurityAuditLogRecordType = "FilteringTeamsMetadata"
	SecurityAuditLogRecordType_FilteringTeamsPostDeliveryAction              SecurityAuditLogRecordType = "FilteringTeamsPostDeliveryAction"
	SecurityAuditLogRecordType_FilteringTeamsUrlInfo                         SecurityAuditLogRecordType = "FilteringTeamsUrlInfo"
	SecurityAuditLogRecordType_FilteringTimeTravelDocMetadata                SecurityAuditLogRecordType = "FilteringTimeTravelDocMetadata"
	SecurityAuditLogRecordType_FilteringUrlClick                             SecurityAuditLogRecordType = "FilteringUrlClick"
	SecurityAuditLogRecordType_FilteringUrlInfo                              SecurityAuditLogRecordType = "FilteringUrlInfo"
	SecurityAuditLogRecordType_FilteringUrlPostClickAction                   SecurityAuditLogRecordType = "FilteringUrlPostClickAction"
	SecurityAuditLogRecordType_HRSignal                                      SecurityAuditLogRecordType = "HRSignal"
	SecurityAuditLogRecordType_HealthcareSignal                              SecurityAuditLogRecordType = "HealthcareSignal"
	SecurityAuditLogRecordType_HostedRpa                                     SecurityAuditLogRecordType = "HostedRpa"
	SecurityAuditLogRecordType_HygieneEvent                                  SecurityAuditLogRecordType = "HygieneEvent"
	SecurityAuditLogRecordType_IncidentStatus                                SecurityAuditLogRecordType = "IncidentStatus"
	SecurityAuditLogRecordType_InformationBarrierPolicyApplication           SecurityAuditLogRecordType = "InformationBarrierPolicyApplication"
	SecurityAuditLogRecordType_InformationWorkerProtection                   SecurityAuditLogRecordType = "InformationWorkerProtection"
	SecurityAuditLogRecordType_IrmUserDefinedDetectionSignal                 SecurityAuditLogRecordType = "IrmUserDefinedDetectionSignal"
	SecurityAuditLogRecordType_Kaizala                                       SecurityAuditLogRecordType = "Kaizala"
	SecurityAuditLogRecordType_LabelAnalyticsAggregate                       SecurityAuditLogRecordType = "LabelAnalyticsAggregate"
	SecurityAuditLogRecordType_LabelContentExplorer                          SecurityAuditLogRecordType = "LabelContentExplorer"
	SecurityAuditLogRecordType_LabelExplorer                                 SecurityAuditLogRecordType = "LabelExplorer"
	SecurityAuditLogRecordType_LargeContentMetadata                          SecurityAuditLogRecordType = "LargeContentMetadata"
	SecurityAuditLogRecordType_M365ComplianceConnector                       SecurityAuditLogRecordType = "M365ComplianceConnector"
	SecurityAuditLogRecordType_M365DAAD                                      SecurityAuditLogRecordType = "M365DAAD"
	SecurityAuditLogRecordType_MAPGAlerts                                    SecurityAuditLogRecordType = "MAPGAlerts"
	SecurityAuditLogRecordType_MAPGOnboard                                   SecurityAuditLogRecordType = "MAPGOnboard"
	SecurityAuditLogRecordType_MAPGPolicy                                    SecurityAuditLogRecordType = "MAPGPolicy"
	SecurityAuditLogRecordType_MAPGRemediation                               SecurityAuditLogRecordType = "MAPGRemediation"
	SecurityAuditLogRecordType_MCASAlerts                                    SecurityAuditLogRecordType = "MCASAlerts"
	SecurityAuditLogRecordType_MDADataSecuritySignal                         SecurityAuditLogRecordType = "MDADataSecuritySignal"
	SecurityAuditLogRecordType_MDATPAudit                                    SecurityAuditLogRecordType = "MDATPAudit"
	SecurityAuditLogRecordType_MDCAssessments                                SecurityAuditLogRecordType = "MDCAssessments"
	SecurityAuditLogRecordType_MDCRegulatoryComplianceAssessments            SecurityAuditLogRecordType = "MDCRegulatoryComplianceAssessments"
	SecurityAuditLogRecordType_MDCRegulatoryComplianceControls               SecurityAuditLogRecordType = "MDCRegulatoryComplianceControls"
	SecurityAuditLogRecordType_MDCRegulatoryComplianceStandards              SecurityAuditLogRecordType = "MDCRegulatoryComplianceStandards"
	SecurityAuditLogRecordType_MDCSecurityConnectors                         SecurityAuditLogRecordType = "MDCSecurityConnectors"
	SecurityAuditLogRecordType_MIPLabel                                      SecurityAuditLogRecordType = "MIPLabel"
	SecurityAuditLogRecordType_MS365DCustomDetection                         SecurityAuditLogRecordType = "MS365DCustomDetection"
	SecurityAuditLogRecordType_MS365DIncident                                SecurityAuditLogRecordType = "MS365DIncident"
	SecurityAuditLogRecordType_MS365DSuppressionRule                         SecurityAuditLogRecordType = "MS365DSuppressionRule"
	SecurityAuditLogRecordType_MSDEGeneralSettings                           SecurityAuditLogRecordType = "MSDEGeneralSettings"
	SecurityAuditLogRecordType_MSDEIndicatorsSettings                        SecurityAuditLogRecordType = "MSDEIndicatorsSettings"
	SecurityAuditLogRecordType_MSDEResponseActions                           SecurityAuditLogRecordType = "MSDEResponseActions"
	SecurityAuditLogRecordType_MSDERolesSettings                             SecurityAuditLogRecordType = "MSDERolesSettings"
	SecurityAuditLogRecordType_MSTIC                                         SecurityAuditLogRecordType = "MSTIC"
	SecurityAuditLogRecordType_MailSubmission                                SecurityAuditLogRecordType = "MailSubmission"
	SecurityAuditLogRecordType_ManagedTenants                                SecurityAuditLogRecordType = "ManagedTenants"
	SecurityAuditLogRecordType_Microsoft365Group                             SecurityAuditLogRecordType = "Microsoft365Group"
	SecurityAuditLogRecordType_MicrosoftDefenderForIdentityAudit             SecurityAuditLogRecordType = "MicrosoftDefenderForIdentityAudit"
	SecurityAuditLogRecordType_MicrosoftFlow                                 SecurityAuditLogRecordType = "MicrosoftFlow"
	SecurityAuditLogRecordType_MicrosoftForms                                SecurityAuditLogRecordType = "MicrosoftForms"
	SecurityAuditLogRecordType_MicrosoftGraphDataConnectConsent              SecurityAuditLogRecordType = "MicrosoftGraphDataConnectConsent"
	SecurityAuditLogRecordType_MicrosoftGraphDataConnectOperation            SecurityAuditLogRecordType = "MicrosoftGraphDataConnectOperation"
	SecurityAuditLogRecordType_MicrosoftManagedServicePlatform               SecurityAuditLogRecordType = "MicrosoftManagedServicePlatform"
	SecurityAuditLogRecordType_MicrosoftPurview                              SecurityAuditLogRecordType = "MicrosoftPurview"
	SecurityAuditLogRecordType_MicrosoftStream                               SecurityAuditLogRecordType = "MicrosoftStream"
	SecurityAuditLogRecordType_MicrosoftTeams                                SecurityAuditLogRecordType = "MicrosoftTeams"
	SecurityAuditLogRecordType_MicrosoftTeamsAdmin                           SecurityAuditLogRecordType = "MicrosoftTeamsAdmin"
	SecurityAuditLogRecordType_MicrosoftTeamsAnalytics                       SecurityAuditLogRecordType = "MicrosoftTeamsAnalytics"
	SecurityAuditLogRecordType_MicrosoftTeamsDevice                          SecurityAuditLogRecordType = "MicrosoftTeamsDevice"
	SecurityAuditLogRecordType_MicrosoftTeamsSensitivityLabelAction          SecurityAuditLogRecordType = "MicrosoftTeamsSensitivityLabelAction"
	SecurityAuditLogRecordType_MicrosoftTeamsShifts                          SecurityAuditLogRecordType = "MicrosoftTeamsShifts"
	SecurityAuditLogRecordType_MicrosoftTodoAudit                            SecurityAuditLogRecordType = "MicrosoftTodoAudit"
	SecurityAuditLogRecordType_MipAutoLabelExchangeItem                      SecurityAuditLogRecordType = "MipAutoLabelExchangeItem"
	SecurityAuditLogRecordType_MipAutoLabelProgressFeedback                  SecurityAuditLogRecordType = "MipAutoLabelProgressFeedback"
	SecurityAuditLogRecordType_MipAutoLabelSharePointItem                    SecurityAuditLogRecordType = "MipAutoLabelSharePointItem"
	SecurityAuditLogRecordType_MipAutoLabelSharePointPolicyLocation          SecurityAuditLogRecordType = "MipAutoLabelSharePointPolicyLocation"
	SecurityAuditLogRecordType_MipAutoLabelSimulationCompletion              SecurityAuditLogRecordType = "MipAutoLabelSimulationCompletion"
	SecurityAuditLogRecordType_MipAutoLabelSimulationProgress                SecurityAuditLogRecordType = "MipAutoLabelSimulationProgress"
	SecurityAuditLogRecordType_MipAutoLabelSimulationStatistics              SecurityAuditLogRecordType = "MipAutoLabelSimulationStatistics"
	SecurityAuditLogRecordType_MipExactDataMatch                             SecurityAuditLogRecordType = "MipExactDataMatch"
	SecurityAuditLogRecordType_MipLabelAnalyticsAuditRecord                  SecurityAuditLogRecordType = "MipLabelAnalyticsAuditRecord"
	SecurityAuditLogRecordType_MultiStageDisposition                         SecurityAuditLogRecordType = "MultiStageDisposition"
	SecurityAuditLogRecordType_MyAnalyticsSettings                           SecurityAuditLogRecordType = "MyAnalyticsSettings"
	SecurityAuditLogRecordType_OMEPortal                                     SecurityAuditLogRecordType = "OMEPortal"
	SecurityAuditLogRecordType_OfficeNative                                  SecurityAuditLogRecordType = "OfficeNative"
	SecurityAuditLogRecordType_OfficeScriptsRunAction                        SecurityAuditLogRecordType = "OfficeScriptsRunAction"
	SecurityAuditLogRecordType_OnPremisesFileShareScannerDlp                 SecurityAuditLogRecordType = "OnPremisesFileShareScannerDlp"
	SecurityAuditLogRecordType_OnPremisesSharePointScannerDlp                SecurityAuditLogRecordType = "OnPremisesSharePointScannerDlp"
	SecurityAuditLogRecordType_OneDrive                                      SecurityAuditLogRecordType = "OneDrive"
	SecurityAuditLogRecordType_PhysicalBadgingSignal                         SecurityAuditLogRecordType = "PhysicalBadgingSignal"
	SecurityAuditLogRecordType_PlannerCopyPlan                               SecurityAuditLogRecordType = "PlannerCopyPlan"
	SecurityAuditLogRecordType_PlannerPlan                                   SecurityAuditLogRecordType = "PlannerPlan"
	SecurityAuditLogRecordType_PlannerPlanList                               SecurityAuditLogRecordType = "PlannerPlanList"
	SecurityAuditLogRecordType_PlannerRoster                                 SecurityAuditLogRecordType = "PlannerRoster"
	SecurityAuditLogRecordType_PlannerRosterSensitivityLabel                 SecurityAuditLogRecordType = "PlannerRosterSensitivityLabel"
	SecurityAuditLogRecordType_PlannerTask                                   SecurityAuditLogRecordType = "PlannerTask"
	SecurityAuditLogRecordType_PlannerTaskList                               SecurityAuditLogRecordType = "PlannerTaskList"
	SecurityAuditLogRecordType_PlannerTenantSettings                         SecurityAuditLogRecordType = "PlannerTenantSettings"
	SecurityAuditLogRecordType_PowerAppsApp                                  SecurityAuditLogRecordType = "PowerAppsApp"
	SecurityAuditLogRecordType_PowerAppsPlan                                 SecurityAuditLogRecordType = "PowerAppsPlan"
	SecurityAuditLogRecordType_PowerAppsResource                             SecurityAuditLogRecordType = "PowerAppsResource"
	SecurityAuditLogRecordType_PowerBIAudit                                  SecurityAuditLogRecordType = "PowerBIAudit"
	SecurityAuditLogRecordType_PowerBIDlp                                    SecurityAuditLogRecordType = "PowerBIDlp"
	SecurityAuditLogRecordType_PowerPagesSite                                SecurityAuditLogRecordType = "PowerPagesSite"
	SecurityAuditLogRecordType_PowerPlatformAdminDlp                         SecurityAuditLogRecordType = "PowerPlatformAdminDlp"
	SecurityAuditLogRecordType_PowerPlatformAdminEnvironment                 SecurityAuditLogRecordType = "PowerPlatformAdminEnvironment"
	SecurityAuditLogRecordType_PowerPlatformLockboxResourceAccessRequest     SecurityAuditLogRecordType = "PowerPlatformLockboxResourceAccessRequest"
	SecurityAuditLogRecordType_PowerPlatformLockboxResourceCommand           SecurityAuditLogRecordType = "PowerPlatformLockboxResourceCommand"
	SecurityAuditLogRecordType_PowerPlatformServiceActivity                  SecurityAuditLogRecordType = "PowerPlatformServiceActivity"
	SecurityAuditLogRecordType_PrivacyDataMatch                              SecurityAuditLogRecordType = "PrivacyDataMatch"
	SecurityAuditLogRecordType_PrivacyDataMinimization                       SecurityAuditLogRecordType = "PrivacyDataMinimization"
	SecurityAuditLogRecordType_PrivacyDigestEmail                            SecurityAuditLogRecordType = "PrivacyDigestEmail"
	SecurityAuditLogRecordType_PrivacyPortal                                 SecurityAuditLogRecordType = "PrivacyPortal"
	SecurityAuditLogRecordType_PrivacyRemediation                            SecurityAuditLogRecordType = "PrivacyRemediation"
	SecurityAuditLogRecordType_PrivacyRemediationAction                      SecurityAuditLogRecordType = "PrivacyRemediationAction"
	SecurityAuditLogRecordType_PrivacyTenantAuditHistoryRecord               SecurityAuditLogRecordType = "PrivacyTenantAuditHistoryRecord"
	SecurityAuditLogRecordType_Project                                       SecurityAuditLogRecordType = "Project"
	SecurityAuditLogRecordType_ProjectForTheWebProject                       SecurityAuditLogRecordType = "ProjectForTheWebProject"
	SecurityAuditLogRecordType_ProjectForTheWebProjectSettings               SecurityAuditLogRecordType = "ProjectForTheWebProjectSettings"
	SecurityAuditLogRecordType_ProjectForTheWebRoadmap                       SecurityAuditLogRecordType = "ProjectForTheWebRoadmap"
	SecurityAuditLogRecordType_ProjectForTheWebRoadmapItem                   SecurityAuditLogRecordType = "ProjectForTheWebRoadmapItem"
	SecurityAuditLogRecordType_ProjectForTheWebRoadmapSettings               SecurityAuditLogRecordType = "ProjectForTheWebRoadmapSettings"
	SecurityAuditLogRecordType_ProjectForTheWebTask                          SecurityAuditLogRecordType = "ProjectForTheWebTask"
	SecurityAuditLogRecordType_PublicFolder                                  SecurityAuditLogRecordType = "PublicFolder"
	SecurityAuditLogRecordType_PurviewDataMapOperation                       SecurityAuditLogRecordType = "PurviewDataMapOperation"
	SecurityAuditLogRecordType_Quarantine                                    SecurityAuditLogRecordType = "Quarantine"
	SecurityAuditLogRecordType_QuarantineMetadata                            SecurityAuditLogRecordType = "QuarantineMetadata"
	SecurityAuditLogRecordType_RecordsManagement                             SecurityAuditLogRecordType = "RecordsManagement"
	SecurityAuditLogRecordType_ScorePlatformGenericAuditRecord               SecurityAuditLogRecordType = "ScorePlatformGenericAuditRecord"
	SecurityAuditLogRecordType_Search                                        SecurityAuditLogRecordType = "Search"
	SecurityAuditLogRecordType_SecureScore                                   SecurityAuditLogRecordType = "SecureScore"
	SecurityAuditLogRecordType_SecurityComplianceAlerts                      SecurityAuditLogRecordType = "SecurityComplianceAlerts"
	SecurityAuditLogRecordType_SecurityComplianceCenterEOPCmdlet             SecurityAuditLogRecordType = "SecurityComplianceCenterEOPCmdlet"
	SecurityAuditLogRecordType_SecurityComplianceInsights                    SecurityAuditLogRecordType = "SecurityComplianceInsights"
	SecurityAuditLogRecordType_SecurityComplianceRBAC                        SecurityAuditLogRecordType = "SecurityComplianceRBAC"
	SecurityAuditLogRecordType_SecurityComplianceUserChange                  SecurityAuditLogRecordType = "SecurityComplianceUserChange"
	SecurityAuditLogRecordType_SensitivityLabelAction                        SecurityAuditLogRecordType = "SensitivityLabelAction"
	SecurityAuditLogRecordType_SensitivityLabelPolicyMatch                   SecurityAuditLogRecordType = "SensitivityLabelPolicyMatch"
	SecurityAuditLogRecordType_SensitivityLabeledFileAction                  SecurityAuditLogRecordType = "SensitivityLabeledFileAction"
	SecurityAuditLogRecordType_SharePoint                                    SecurityAuditLogRecordType = "SharePoint"
	SecurityAuditLogRecordType_SharePointAppPermissionOperation              SecurityAuditLogRecordType = "SharePointAppPermissionOperation"
	SecurityAuditLogRecordType_SharePointCommentOperation                    SecurityAuditLogRecordType = "SharePointCommentOperation"
	SecurityAuditLogRecordType_SharePointContentTypeOperation                SecurityAuditLogRecordType = "SharePointContentTypeOperation"
	SecurityAuditLogRecordType_SharePointFieldOperation                      SecurityAuditLogRecordType = "SharePointFieldOperation"
	SecurityAuditLogRecordType_SharePointFileOperation                       SecurityAuditLogRecordType = "SharePointFileOperation"
	SecurityAuditLogRecordType_SharePointListItemOperation                   SecurityAuditLogRecordType = "SharePointListItemOperation"
	SecurityAuditLogRecordType_SharePointListOperation                       SecurityAuditLogRecordType = "SharePointListOperation"
	SecurityAuditLogRecordType_SharePointSearch                              SecurityAuditLogRecordType = "SharePointSearch"
	SecurityAuditLogRecordType_SharePointSharingOperation                    SecurityAuditLogRecordType = "SharePointSharingOperation"
	SecurityAuditLogRecordType_SkypeForBusinessCmdlets                       SecurityAuditLogRecordType = "SkypeForBusinessCmdlets"
	SecurityAuditLogRecordType_SkypeForBusinessPSTNUsage                     SecurityAuditLogRecordType = "SkypeForBusinessPSTNUsage"
	SecurityAuditLogRecordType_SkypeForBusinessUsersBlocked                  SecurityAuditLogRecordType = "SkypeForBusinessUsersBlocked"
	SecurityAuditLogRecordType_SupervisoryReviewDayXInsight                  SecurityAuditLogRecordType = "SupervisoryReviewDayXInsight"
	SecurityAuditLogRecordType_Sway                                          SecurityAuditLogRecordType = "Sway"
	SecurityAuditLogRecordType_SyntheticProbe                                SecurityAuditLogRecordType = "SyntheticProbe"
	SecurityAuditLogRecordType_TeamsEasyApprovals                            SecurityAuditLogRecordType = "TeamsEasyApprovals"
	SecurityAuditLogRecordType_TeamsHealthcare                               SecurityAuditLogRecordType = "TeamsHealthcare"
	SecurityAuditLogRecordType_TeamsQuarantineMetadata                       SecurityAuditLogRecordType = "TeamsQuarantineMetadata"
	SecurityAuditLogRecordType_TeamsUpdates                                  SecurityAuditLogRecordType = "TeamsUpdates"
	SecurityAuditLogRecordType_TenantAllowBlockList                          SecurityAuditLogRecordType = "TenantAllowBlockList"
	SecurityAuditLogRecordType_ThreatFinder                                  SecurityAuditLogRecordType = "ThreatFinder"
	SecurityAuditLogRecordType_ThreatIntelligence                            SecurityAuditLogRecordType = "ThreatIntelligence"
	SecurityAuditLogRecordType_ThreatIntelligenceAtpContent                  SecurityAuditLogRecordType = "ThreatIntelligenceAtpContent"
	SecurityAuditLogRecordType_ThreatIntelligenceUrl                         SecurityAuditLogRecordType = "ThreatIntelligenceUrl"
	SecurityAuditLogRecordType_TimeTravelFilteringDocMetadata                SecurityAuditLogRecordType = "TimeTravelFilteringDocMetadata"
	SecurityAuditLogRecordType_TimeTravelFilteringDocScan                    SecurityAuditLogRecordType = "TimeTravelFilteringDocScan"
	SecurityAuditLogRecordType_UnifiedSimulationMatchedItem                  SecurityAuditLogRecordType = "UnifiedSimulationMatchedItem"
	SecurityAuditLogRecordType_UnifiedSimulationSummary                      SecurityAuditLogRecordType = "UnifiedSimulationSummary"
	SecurityAuditLogRecordType_UpdateQuarantineMetadata                      SecurityAuditLogRecordType = "UpdateQuarantineMetadata"
	SecurityAuditLogRecordType_UserTraining                                  SecurityAuditLogRecordType = "UserTraining"
	SecurityAuditLogRecordType_VfamCreatePolicy                              SecurityAuditLogRecordType = "VfamCreatePolicy"
	SecurityAuditLogRecordType_VfamDeletePolicy                              SecurityAuditLogRecordType = "VfamDeletePolicy"
	SecurityAuditLogRecordType_VfamUpdatePolicy                              SecurityAuditLogRecordType = "VfamUpdatePolicy"
	SecurityAuditLogRecordType_VivaGoals                                     SecurityAuditLogRecordType = "VivaGoals"
	SecurityAuditLogRecordType_WDATPAlerts                                   SecurityAuditLogRecordType = "WDATPAlerts"
	SecurityAuditLogRecordType_WebpageActivityEndpoint                       SecurityAuditLogRecordType = "WebpageActivityEndpoint"
	SecurityAuditLogRecordType_WorkplaceAnalytics                            SecurityAuditLogRecordType = "WorkplaceAnalytics"
	SecurityAuditLogRecordType_Yammer                                        SecurityAuditLogRecordType = "Yammer"
)

func PossibleValuesForSecurityAuditLogRecordType() []string {
	return []string{
		string(SecurityAuditLogRecordType_AeD),
		string(SecurityAuditLogRecordType_AipDiscover),
		string(SecurityAuditLogRecordType_AipFileDeleted),
		string(SecurityAuditLogRecordType_AipHeartBeat),
		string(SecurityAuditLogRecordType_AipProtectionAction),
		string(SecurityAuditLogRecordType_AipScannerDiscoverEvent),
		string(SecurityAuditLogRecordType_AipSensitivityLabelAction),
		string(SecurityAuditLogRecordType_AirAdminActionInvestigation),
		string(SecurityAuditLogRecordType_AirInvestigation),
		string(SecurityAuditLogRecordType_AirManualInvestigation),
		string(SecurityAuditLogRecordType_Alert),
		string(SecurityAuditLogRecordType_AlertIncident),
		string(SecurityAuditLogRecordType_AlertStatus),
		string(SecurityAuditLogRecordType_ApplicationAudit),
		string(SecurityAuditLogRecordType_AttackSim),
		string(SecurityAuditLogRecordType_AttackSimAdmin),
		string(SecurityAuditLogRecordType_AzureActiveDirectory),
		string(SecurityAuditLogRecordType_AzureActiveDirectoryAccountLogon),
		string(SecurityAuditLogRecordType_AzureActiveDirectoryStsLogon),
		string(SecurityAuditLogRecordType_CDPClassificationDocument),
		string(SecurityAuditLogRecordType_CDPClassificationMailItem),
		string(SecurityAuditLogRecordType_CDPCompliancePolicyExecution),
		string(SecurityAuditLogRecordType_CDPCompliancePolicyUserFeedback),
		string(SecurityAuditLogRecordType_CDPEdgeBlockedMessage),
		string(SecurityAuditLogRecordType_CDPEmailFeatures),
		string(SecurityAuditLogRecordType_CDPHygieneAttachmentInfo),
		string(SecurityAuditLogRecordType_CDPHygieneSummary),
		string(SecurityAuditLogRecordType_CDPHygieneUrlInfo),
		string(SecurityAuditLogRecordType_CDPMlInferencingResult),
		string(SecurityAuditLogRecordType_CDPPackageManagerHygieneEvent),
		string(SecurityAuditLogRecordType_CDPPostMailDeliveryAction),
		string(SecurityAuditLogRecordType_CDPPredictiveCodingLabel),
		string(SecurityAuditLogRecordType_CDPUnifiedFeedback),
		string(SecurityAuditLogRecordType_CDPUrlClick),
		string(SecurityAuditLogRecordType_CMImprovementActionChange),
		string(SecurityAuditLogRecordType_CRM),
		string(SecurityAuditLogRecordType_Campaign),
		string(SecurityAuditLogRecordType_Case),
		string(SecurityAuditLogRecordType_CaseInvestigation),
		string(SecurityAuditLogRecordType_CdpColdCrawlStatus),
		string(SecurityAuditLogRecordType_CdpContentExplorerAggregateRecord),
		string(SecurityAuditLogRecordType_CdpDlpSensitive),
		string(SecurityAuditLogRecordType_ComplianceConnector),
		string(SecurityAuditLogRecordType_ComplianceDLPEndpoint),
		string(SecurityAuditLogRecordType_ComplianceDLPExchange),
		string(SecurityAuditLogRecordType_ComplianceDLPExchangeClassification),
		string(SecurityAuditLogRecordType_ComplianceDLPSharePoint),
		string(SecurityAuditLogRecordType_ComplianceDLPSharePointClassification),
		string(SecurityAuditLogRecordType_ComplianceDLPSharePointClassificationExtended),
		string(SecurityAuditLogRecordType_ComplianceSupervisionExchange),
		string(SecurityAuditLogRecordType_ConsumptionResource),
		string(SecurityAuditLogRecordType_CoreReportingSettings),
		string(SecurityAuditLogRecordType_CortanaBriefing),
		string(SecurityAuditLogRecordType_CustomerKeyServiceEncryption),
		string(SecurityAuditLogRecordType_DLPEndpoint),
		string(SecurityAuditLogRecordType_DataCenterSecurityCmdlet),
		string(SecurityAuditLogRecordType_DataGovernance),
		string(SecurityAuditLogRecordType_DataInsightsRestApiAudit),
		string(SecurityAuditLogRecordType_DataShareOperation),
		string(SecurityAuditLogRecordType_DefenderExpertsforXDRAdmin),
		string(SecurityAuditLogRecordType_Discovery),
		string(SecurityAuditLogRecordType_DlpImportResult),
		string(SecurityAuditLogRecordType_DlpSensitiveInformationType),
		string(SecurityAuditLogRecordType_EHRConnector),
		string(SecurityAuditLogRecordType_EduDataLakeDownloadOperation),
		string(SecurityAuditLogRecordType_ExchangeAdmin),
		string(SecurityAuditLogRecordType_ExchangeAggregatedOperation),
		string(SecurityAuditLogRecordType_ExchangeItem),
		string(SecurityAuditLogRecordType_ExchangeItemAggregated),
		string(SecurityAuditLogRecordType_ExchangeItemGroup),
		string(SecurityAuditLogRecordType_ExchangeSearch),
		string(SecurityAuditLogRecordType_FilteringAtpDetonationInfo),
		string(SecurityAuditLogRecordType_FilteringAttachmentInfo),
		string(SecurityAuditLogRecordType_FilteringDelistingMetadata),
		string(SecurityAuditLogRecordType_FilteringDocMetadata),
		string(SecurityAuditLogRecordType_FilteringDocScan),
		string(SecurityAuditLogRecordType_FilteringEmailContentFeatures),
		string(SecurityAuditLogRecordType_FilteringEmailFeatures),
		string(SecurityAuditLogRecordType_FilteringEntityEvent),
		string(SecurityAuditLogRecordType_FilteringMailGradingResult),
		string(SecurityAuditLogRecordType_FilteringMailMetadata),
		string(SecurityAuditLogRecordType_FilteringMailSubmission),
		string(SecurityAuditLogRecordType_FilteringPostMailDeliveryAction),
		string(SecurityAuditLogRecordType_FilteringRuleHits),
		string(SecurityAuditLogRecordType_FilteringRuntimeInfo),
		string(SecurityAuditLogRecordType_FilteringTeamsMetadata),
		string(SecurityAuditLogRecordType_FilteringTeamsPostDeliveryAction),
		string(SecurityAuditLogRecordType_FilteringTeamsUrlInfo),
		string(SecurityAuditLogRecordType_FilteringTimeTravelDocMetadata),
		string(SecurityAuditLogRecordType_FilteringUrlClick),
		string(SecurityAuditLogRecordType_FilteringUrlInfo),
		string(SecurityAuditLogRecordType_FilteringUrlPostClickAction),
		string(SecurityAuditLogRecordType_HRSignal),
		string(SecurityAuditLogRecordType_HealthcareSignal),
		string(SecurityAuditLogRecordType_HostedRpa),
		string(SecurityAuditLogRecordType_HygieneEvent),
		string(SecurityAuditLogRecordType_IncidentStatus),
		string(SecurityAuditLogRecordType_InformationBarrierPolicyApplication),
		string(SecurityAuditLogRecordType_InformationWorkerProtection),
		string(SecurityAuditLogRecordType_IrmUserDefinedDetectionSignal),
		string(SecurityAuditLogRecordType_Kaizala),
		string(SecurityAuditLogRecordType_LabelAnalyticsAggregate),
		string(SecurityAuditLogRecordType_LabelContentExplorer),
		string(SecurityAuditLogRecordType_LabelExplorer),
		string(SecurityAuditLogRecordType_LargeContentMetadata),
		string(SecurityAuditLogRecordType_M365ComplianceConnector),
		string(SecurityAuditLogRecordType_M365DAAD),
		string(SecurityAuditLogRecordType_MAPGAlerts),
		string(SecurityAuditLogRecordType_MAPGOnboard),
		string(SecurityAuditLogRecordType_MAPGPolicy),
		string(SecurityAuditLogRecordType_MAPGRemediation),
		string(SecurityAuditLogRecordType_MCASAlerts),
		string(SecurityAuditLogRecordType_MDADataSecuritySignal),
		string(SecurityAuditLogRecordType_MDATPAudit),
		string(SecurityAuditLogRecordType_MDCAssessments),
		string(SecurityAuditLogRecordType_MDCRegulatoryComplianceAssessments),
		string(SecurityAuditLogRecordType_MDCRegulatoryComplianceControls),
		string(SecurityAuditLogRecordType_MDCRegulatoryComplianceStandards),
		string(SecurityAuditLogRecordType_MDCSecurityConnectors),
		string(SecurityAuditLogRecordType_MIPLabel),
		string(SecurityAuditLogRecordType_MS365DCustomDetection),
		string(SecurityAuditLogRecordType_MS365DIncident),
		string(SecurityAuditLogRecordType_MS365DSuppressionRule),
		string(SecurityAuditLogRecordType_MSDEGeneralSettings),
		string(SecurityAuditLogRecordType_MSDEIndicatorsSettings),
		string(SecurityAuditLogRecordType_MSDEResponseActions),
		string(SecurityAuditLogRecordType_MSDERolesSettings),
		string(SecurityAuditLogRecordType_MSTIC),
		string(SecurityAuditLogRecordType_MailSubmission),
		string(SecurityAuditLogRecordType_ManagedTenants),
		string(SecurityAuditLogRecordType_Microsoft365Group),
		string(SecurityAuditLogRecordType_MicrosoftDefenderForIdentityAudit),
		string(SecurityAuditLogRecordType_MicrosoftFlow),
		string(SecurityAuditLogRecordType_MicrosoftForms),
		string(SecurityAuditLogRecordType_MicrosoftGraphDataConnectConsent),
		string(SecurityAuditLogRecordType_MicrosoftGraphDataConnectOperation),
		string(SecurityAuditLogRecordType_MicrosoftManagedServicePlatform),
		string(SecurityAuditLogRecordType_MicrosoftPurview),
		string(SecurityAuditLogRecordType_MicrosoftStream),
		string(SecurityAuditLogRecordType_MicrosoftTeams),
		string(SecurityAuditLogRecordType_MicrosoftTeamsAdmin),
		string(SecurityAuditLogRecordType_MicrosoftTeamsAnalytics),
		string(SecurityAuditLogRecordType_MicrosoftTeamsDevice),
		string(SecurityAuditLogRecordType_MicrosoftTeamsSensitivityLabelAction),
		string(SecurityAuditLogRecordType_MicrosoftTeamsShifts),
		string(SecurityAuditLogRecordType_MicrosoftTodoAudit),
		string(SecurityAuditLogRecordType_MipAutoLabelExchangeItem),
		string(SecurityAuditLogRecordType_MipAutoLabelProgressFeedback),
		string(SecurityAuditLogRecordType_MipAutoLabelSharePointItem),
		string(SecurityAuditLogRecordType_MipAutoLabelSharePointPolicyLocation),
		string(SecurityAuditLogRecordType_MipAutoLabelSimulationCompletion),
		string(SecurityAuditLogRecordType_MipAutoLabelSimulationProgress),
		string(SecurityAuditLogRecordType_MipAutoLabelSimulationStatistics),
		string(SecurityAuditLogRecordType_MipExactDataMatch),
		string(SecurityAuditLogRecordType_MipLabelAnalyticsAuditRecord),
		string(SecurityAuditLogRecordType_MultiStageDisposition),
		string(SecurityAuditLogRecordType_MyAnalyticsSettings),
		string(SecurityAuditLogRecordType_OMEPortal),
		string(SecurityAuditLogRecordType_OfficeNative),
		string(SecurityAuditLogRecordType_OfficeScriptsRunAction),
		string(SecurityAuditLogRecordType_OnPremisesFileShareScannerDlp),
		string(SecurityAuditLogRecordType_OnPremisesSharePointScannerDlp),
		string(SecurityAuditLogRecordType_OneDrive),
		string(SecurityAuditLogRecordType_PhysicalBadgingSignal),
		string(SecurityAuditLogRecordType_PlannerCopyPlan),
		string(SecurityAuditLogRecordType_PlannerPlan),
		string(SecurityAuditLogRecordType_PlannerPlanList),
		string(SecurityAuditLogRecordType_PlannerRoster),
		string(SecurityAuditLogRecordType_PlannerRosterSensitivityLabel),
		string(SecurityAuditLogRecordType_PlannerTask),
		string(SecurityAuditLogRecordType_PlannerTaskList),
		string(SecurityAuditLogRecordType_PlannerTenantSettings),
		string(SecurityAuditLogRecordType_PowerAppsApp),
		string(SecurityAuditLogRecordType_PowerAppsPlan),
		string(SecurityAuditLogRecordType_PowerAppsResource),
		string(SecurityAuditLogRecordType_PowerBIAudit),
		string(SecurityAuditLogRecordType_PowerBIDlp),
		string(SecurityAuditLogRecordType_PowerPagesSite),
		string(SecurityAuditLogRecordType_PowerPlatformAdminDlp),
		string(SecurityAuditLogRecordType_PowerPlatformAdminEnvironment),
		string(SecurityAuditLogRecordType_PowerPlatformLockboxResourceAccessRequest),
		string(SecurityAuditLogRecordType_PowerPlatformLockboxResourceCommand),
		string(SecurityAuditLogRecordType_PowerPlatformServiceActivity),
		string(SecurityAuditLogRecordType_PrivacyDataMatch),
		string(SecurityAuditLogRecordType_PrivacyDataMinimization),
		string(SecurityAuditLogRecordType_PrivacyDigestEmail),
		string(SecurityAuditLogRecordType_PrivacyPortal),
		string(SecurityAuditLogRecordType_PrivacyRemediation),
		string(SecurityAuditLogRecordType_PrivacyRemediationAction),
		string(SecurityAuditLogRecordType_PrivacyTenantAuditHistoryRecord),
		string(SecurityAuditLogRecordType_Project),
		string(SecurityAuditLogRecordType_ProjectForTheWebProject),
		string(SecurityAuditLogRecordType_ProjectForTheWebProjectSettings),
		string(SecurityAuditLogRecordType_ProjectForTheWebRoadmap),
		string(SecurityAuditLogRecordType_ProjectForTheWebRoadmapItem),
		string(SecurityAuditLogRecordType_ProjectForTheWebRoadmapSettings),
		string(SecurityAuditLogRecordType_ProjectForTheWebTask),
		string(SecurityAuditLogRecordType_PublicFolder),
		string(SecurityAuditLogRecordType_PurviewDataMapOperation),
		string(SecurityAuditLogRecordType_Quarantine),
		string(SecurityAuditLogRecordType_QuarantineMetadata),
		string(SecurityAuditLogRecordType_RecordsManagement),
		string(SecurityAuditLogRecordType_ScorePlatformGenericAuditRecord),
		string(SecurityAuditLogRecordType_Search),
		string(SecurityAuditLogRecordType_SecureScore),
		string(SecurityAuditLogRecordType_SecurityComplianceAlerts),
		string(SecurityAuditLogRecordType_SecurityComplianceCenterEOPCmdlet),
		string(SecurityAuditLogRecordType_SecurityComplianceInsights),
		string(SecurityAuditLogRecordType_SecurityComplianceRBAC),
		string(SecurityAuditLogRecordType_SecurityComplianceUserChange),
		string(SecurityAuditLogRecordType_SensitivityLabelAction),
		string(SecurityAuditLogRecordType_SensitivityLabelPolicyMatch),
		string(SecurityAuditLogRecordType_SensitivityLabeledFileAction),
		string(SecurityAuditLogRecordType_SharePoint),
		string(SecurityAuditLogRecordType_SharePointAppPermissionOperation),
		string(SecurityAuditLogRecordType_SharePointCommentOperation),
		string(SecurityAuditLogRecordType_SharePointContentTypeOperation),
		string(SecurityAuditLogRecordType_SharePointFieldOperation),
		string(SecurityAuditLogRecordType_SharePointFileOperation),
		string(SecurityAuditLogRecordType_SharePointListItemOperation),
		string(SecurityAuditLogRecordType_SharePointListOperation),
		string(SecurityAuditLogRecordType_SharePointSearch),
		string(SecurityAuditLogRecordType_SharePointSharingOperation),
		string(SecurityAuditLogRecordType_SkypeForBusinessCmdlets),
		string(SecurityAuditLogRecordType_SkypeForBusinessPSTNUsage),
		string(SecurityAuditLogRecordType_SkypeForBusinessUsersBlocked),
		string(SecurityAuditLogRecordType_SupervisoryReviewDayXInsight),
		string(SecurityAuditLogRecordType_Sway),
		string(SecurityAuditLogRecordType_SyntheticProbe),
		string(SecurityAuditLogRecordType_TeamsEasyApprovals),
		string(SecurityAuditLogRecordType_TeamsHealthcare),
		string(SecurityAuditLogRecordType_TeamsQuarantineMetadata),
		string(SecurityAuditLogRecordType_TeamsUpdates),
		string(SecurityAuditLogRecordType_TenantAllowBlockList),
		string(SecurityAuditLogRecordType_ThreatFinder),
		string(SecurityAuditLogRecordType_ThreatIntelligence),
		string(SecurityAuditLogRecordType_ThreatIntelligenceAtpContent),
		string(SecurityAuditLogRecordType_ThreatIntelligenceUrl),
		string(SecurityAuditLogRecordType_TimeTravelFilteringDocMetadata),
		string(SecurityAuditLogRecordType_TimeTravelFilteringDocScan),
		string(SecurityAuditLogRecordType_UnifiedSimulationMatchedItem),
		string(SecurityAuditLogRecordType_UnifiedSimulationSummary),
		string(SecurityAuditLogRecordType_UpdateQuarantineMetadata),
		string(SecurityAuditLogRecordType_UserTraining),
		string(SecurityAuditLogRecordType_VfamCreatePolicy),
		string(SecurityAuditLogRecordType_VfamDeletePolicy),
		string(SecurityAuditLogRecordType_VfamUpdatePolicy),
		string(SecurityAuditLogRecordType_VivaGoals),
		string(SecurityAuditLogRecordType_WDATPAlerts),
		string(SecurityAuditLogRecordType_WebpageActivityEndpoint),
		string(SecurityAuditLogRecordType_WorkplaceAnalytics),
		string(SecurityAuditLogRecordType_Yammer),
	}
}

func (s *SecurityAuditLogRecordType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAuditLogRecordType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAuditLogRecordType(input string) (*SecurityAuditLogRecordType, error) {
	vals := map[string]SecurityAuditLogRecordType{
		"aed":                                           SecurityAuditLogRecordType_AeD,
		"aipdiscover":                                   SecurityAuditLogRecordType_AipDiscover,
		"aipfiledeleted":                                SecurityAuditLogRecordType_AipFileDeleted,
		"aipheartbeat":                                  SecurityAuditLogRecordType_AipHeartBeat,
		"aipprotectionaction":                           SecurityAuditLogRecordType_AipProtectionAction,
		"aipscannerdiscoverevent":                       SecurityAuditLogRecordType_AipScannerDiscoverEvent,
		"aipsensitivitylabelaction":                     SecurityAuditLogRecordType_AipSensitivityLabelAction,
		"airadminactioninvestigation":                   SecurityAuditLogRecordType_AirAdminActionInvestigation,
		"airinvestigation":                              SecurityAuditLogRecordType_AirInvestigation,
		"airmanualinvestigation":                        SecurityAuditLogRecordType_AirManualInvestigation,
		"alert":                                         SecurityAuditLogRecordType_Alert,
		"alertincident":                                 SecurityAuditLogRecordType_AlertIncident,
		"alertstatus":                                   SecurityAuditLogRecordType_AlertStatus,
		"applicationaudit":                              SecurityAuditLogRecordType_ApplicationAudit,
		"attacksim":                                     SecurityAuditLogRecordType_AttackSim,
		"attacksimadmin":                                SecurityAuditLogRecordType_AttackSimAdmin,
		"azureactivedirectory":                          SecurityAuditLogRecordType_AzureActiveDirectory,
		"azureactivedirectoryaccountlogon":              SecurityAuditLogRecordType_AzureActiveDirectoryAccountLogon,
		"azureactivedirectorystslogon":                  SecurityAuditLogRecordType_AzureActiveDirectoryStsLogon,
		"cdpclassificationdocument":                     SecurityAuditLogRecordType_CDPClassificationDocument,
		"cdpclassificationmailitem":                     SecurityAuditLogRecordType_CDPClassificationMailItem,
		"cdpcompliancepolicyexecution":                  SecurityAuditLogRecordType_CDPCompliancePolicyExecution,
		"cdpcompliancepolicyuserfeedback":               SecurityAuditLogRecordType_CDPCompliancePolicyUserFeedback,
		"cdpedgeblockedmessage":                         SecurityAuditLogRecordType_CDPEdgeBlockedMessage,
		"cdpemailfeatures":                              SecurityAuditLogRecordType_CDPEmailFeatures,
		"cdphygieneattachmentinfo":                      SecurityAuditLogRecordType_CDPHygieneAttachmentInfo,
		"cdphygienesummary":                             SecurityAuditLogRecordType_CDPHygieneSummary,
		"cdphygieneurlinfo":                             SecurityAuditLogRecordType_CDPHygieneUrlInfo,
		"cdpmlinferencingresult":                        SecurityAuditLogRecordType_CDPMlInferencingResult,
		"cdppackagemanagerhygieneevent":                 SecurityAuditLogRecordType_CDPPackageManagerHygieneEvent,
		"cdppostmaildeliveryaction":                     SecurityAuditLogRecordType_CDPPostMailDeliveryAction,
		"cdppredictivecodinglabel":                      SecurityAuditLogRecordType_CDPPredictiveCodingLabel,
		"cdpunifiedfeedback":                            SecurityAuditLogRecordType_CDPUnifiedFeedback,
		"cdpurlclick":                                   SecurityAuditLogRecordType_CDPUrlClick,
		"cmimprovementactionchange":                     SecurityAuditLogRecordType_CMImprovementActionChange,
		"crm":                                           SecurityAuditLogRecordType_CRM,
		"campaign":                                      SecurityAuditLogRecordType_Campaign,
		"case":                                          SecurityAuditLogRecordType_Case,
		"caseinvestigation":                             SecurityAuditLogRecordType_CaseInvestigation,
		"cdpcoldcrawlstatus":                            SecurityAuditLogRecordType_CdpColdCrawlStatus,
		"cdpcontentexploreraggregaterecord":             SecurityAuditLogRecordType_CdpContentExplorerAggregateRecord,
		"cdpdlpsensitive":                               SecurityAuditLogRecordType_CdpDlpSensitive,
		"complianceconnector":                           SecurityAuditLogRecordType_ComplianceConnector,
		"compliancedlpendpoint":                         SecurityAuditLogRecordType_ComplianceDLPEndpoint,
		"compliancedlpexchange":                         SecurityAuditLogRecordType_ComplianceDLPExchange,
		"compliancedlpexchangeclassification":           SecurityAuditLogRecordType_ComplianceDLPExchangeClassification,
		"compliancedlpsharepoint":                       SecurityAuditLogRecordType_ComplianceDLPSharePoint,
		"compliancedlpsharepointclassification":         SecurityAuditLogRecordType_ComplianceDLPSharePointClassification,
		"compliancedlpsharepointclassificationextended": SecurityAuditLogRecordType_ComplianceDLPSharePointClassificationExtended,
		"compliancesupervisionexchange":                 SecurityAuditLogRecordType_ComplianceSupervisionExchange,
		"consumptionresource":                           SecurityAuditLogRecordType_ConsumptionResource,
		"corereportingsettings":                         SecurityAuditLogRecordType_CoreReportingSettings,
		"cortanabriefing":                               SecurityAuditLogRecordType_CortanaBriefing,
		"customerkeyserviceencryption":                  SecurityAuditLogRecordType_CustomerKeyServiceEncryption,
		"dlpendpoint":                                   SecurityAuditLogRecordType_DLPEndpoint,
		"datacentersecuritycmdlet":                      SecurityAuditLogRecordType_DataCenterSecurityCmdlet,
		"datagovernance":                                SecurityAuditLogRecordType_DataGovernance,
		"datainsightsrestapiaudit":                      SecurityAuditLogRecordType_DataInsightsRestApiAudit,
		"datashareoperation":                            SecurityAuditLogRecordType_DataShareOperation,
		"defenderexpertsforxdradmin":                    SecurityAuditLogRecordType_DefenderExpertsforXDRAdmin,
		"discovery":                                     SecurityAuditLogRecordType_Discovery,
		"dlpimportresult":                               SecurityAuditLogRecordType_DlpImportResult,
		"dlpsensitiveinformationtype":                   SecurityAuditLogRecordType_DlpSensitiveInformationType,
		"ehrconnector":                                  SecurityAuditLogRecordType_EHRConnector,
		"edudatalakedownloadoperation":                  SecurityAuditLogRecordType_EduDataLakeDownloadOperation,
		"exchangeadmin":                                 SecurityAuditLogRecordType_ExchangeAdmin,
		"exchangeaggregatedoperation":                   SecurityAuditLogRecordType_ExchangeAggregatedOperation,
		"exchangeitem":                                  SecurityAuditLogRecordType_ExchangeItem,
		"exchangeitemaggregated":                        SecurityAuditLogRecordType_ExchangeItemAggregated,
		"exchangeitemgroup":                             SecurityAuditLogRecordType_ExchangeItemGroup,
		"exchangesearch":                                SecurityAuditLogRecordType_ExchangeSearch,
		"filteringatpdetonationinfo":                    SecurityAuditLogRecordType_FilteringAtpDetonationInfo,
		"filteringattachmentinfo":                       SecurityAuditLogRecordType_FilteringAttachmentInfo,
		"filteringdelistingmetadata":                    SecurityAuditLogRecordType_FilteringDelistingMetadata,
		"filteringdocmetadata":                          SecurityAuditLogRecordType_FilteringDocMetadata,
		"filteringdocscan":                              SecurityAuditLogRecordType_FilteringDocScan,
		"filteringemailcontentfeatures":                 SecurityAuditLogRecordType_FilteringEmailContentFeatures,
		"filteringemailfeatures":                        SecurityAuditLogRecordType_FilteringEmailFeatures,
		"filteringentityevent":                          SecurityAuditLogRecordType_FilteringEntityEvent,
		"filteringmailgradingresult":                    SecurityAuditLogRecordType_FilteringMailGradingResult,
		"filteringmailmetadata":                         SecurityAuditLogRecordType_FilteringMailMetadata,
		"filteringmailsubmission":                       SecurityAuditLogRecordType_FilteringMailSubmission,
		"filteringpostmaildeliveryaction":               SecurityAuditLogRecordType_FilteringPostMailDeliveryAction,
		"filteringrulehits":                             SecurityAuditLogRecordType_FilteringRuleHits,
		"filteringruntimeinfo":                          SecurityAuditLogRecordType_FilteringRuntimeInfo,
		"filteringteamsmetadata":                        SecurityAuditLogRecordType_FilteringTeamsMetadata,
		"filteringteamspostdeliveryaction":              SecurityAuditLogRecordType_FilteringTeamsPostDeliveryAction,
		"filteringteamsurlinfo":                         SecurityAuditLogRecordType_FilteringTeamsUrlInfo,
		"filteringtimetraveldocmetadata":                SecurityAuditLogRecordType_FilteringTimeTravelDocMetadata,
		"filteringurlclick":                             SecurityAuditLogRecordType_FilteringUrlClick,
		"filteringurlinfo":                              SecurityAuditLogRecordType_FilteringUrlInfo,
		"filteringurlpostclickaction":                   SecurityAuditLogRecordType_FilteringUrlPostClickAction,
		"hrsignal":                                      SecurityAuditLogRecordType_HRSignal,
		"healthcaresignal":                              SecurityAuditLogRecordType_HealthcareSignal,
		"hostedrpa":                                     SecurityAuditLogRecordType_HostedRpa,
		"hygieneevent":                                  SecurityAuditLogRecordType_HygieneEvent,
		"incidentstatus":                                SecurityAuditLogRecordType_IncidentStatus,
		"informationbarrierpolicyapplication":           SecurityAuditLogRecordType_InformationBarrierPolicyApplication,
		"informationworkerprotection":                   SecurityAuditLogRecordType_InformationWorkerProtection,
		"irmuserdefineddetectionsignal":                 SecurityAuditLogRecordType_IrmUserDefinedDetectionSignal,
		"kaizala":                                       SecurityAuditLogRecordType_Kaizala,
		"labelanalyticsaggregate":                       SecurityAuditLogRecordType_LabelAnalyticsAggregate,
		"labelcontentexplorer":                          SecurityAuditLogRecordType_LabelContentExplorer,
		"labelexplorer":                                 SecurityAuditLogRecordType_LabelExplorer,
		"largecontentmetadata":                          SecurityAuditLogRecordType_LargeContentMetadata,
		"m365complianceconnector":                       SecurityAuditLogRecordType_M365ComplianceConnector,
		"m365daad":                                      SecurityAuditLogRecordType_M365DAAD,
		"mapgalerts":                                    SecurityAuditLogRecordType_MAPGAlerts,
		"mapgonboard":                                   SecurityAuditLogRecordType_MAPGOnboard,
		"mapgpolicy":                                    SecurityAuditLogRecordType_MAPGPolicy,
		"mapgremediation":                               SecurityAuditLogRecordType_MAPGRemediation,
		"mcasalerts":                                    SecurityAuditLogRecordType_MCASAlerts,
		"mdadatasecuritysignal":                         SecurityAuditLogRecordType_MDADataSecuritySignal,
		"mdatpaudit":                                    SecurityAuditLogRecordType_MDATPAudit,
		"mdcassessments":                                SecurityAuditLogRecordType_MDCAssessments,
		"mdcregulatorycomplianceassessments":            SecurityAuditLogRecordType_MDCRegulatoryComplianceAssessments,
		"mdcregulatorycompliancecontrols":               SecurityAuditLogRecordType_MDCRegulatoryComplianceControls,
		"mdcregulatorycompliancestandards":              SecurityAuditLogRecordType_MDCRegulatoryComplianceStandards,
		"mdcsecurityconnectors":                         SecurityAuditLogRecordType_MDCSecurityConnectors,
		"miplabel":                                      SecurityAuditLogRecordType_MIPLabel,
		"ms365dcustomdetection":                         SecurityAuditLogRecordType_MS365DCustomDetection,
		"ms365dincident":                                SecurityAuditLogRecordType_MS365DIncident,
		"ms365dsuppressionrule":                         SecurityAuditLogRecordType_MS365DSuppressionRule,
		"msdegeneralsettings":                           SecurityAuditLogRecordType_MSDEGeneralSettings,
		"msdeindicatorssettings":                        SecurityAuditLogRecordType_MSDEIndicatorsSettings,
		"msderesponseactions":                           SecurityAuditLogRecordType_MSDEResponseActions,
		"msderolessettings":                             SecurityAuditLogRecordType_MSDERolesSettings,
		"mstic":                                         SecurityAuditLogRecordType_MSTIC,
		"mailsubmission":                                SecurityAuditLogRecordType_MailSubmission,
		"managedtenants":                                SecurityAuditLogRecordType_ManagedTenants,
		"microsoft365group":                             SecurityAuditLogRecordType_Microsoft365Group,
		"microsoftdefenderforidentityaudit":             SecurityAuditLogRecordType_MicrosoftDefenderForIdentityAudit,
		"microsoftflow":                                 SecurityAuditLogRecordType_MicrosoftFlow,
		"microsoftforms":                                SecurityAuditLogRecordType_MicrosoftForms,
		"microsoftgraphdataconnectconsent":              SecurityAuditLogRecordType_MicrosoftGraphDataConnectConsent,
		"microsoftgraphdataconnectoperation":            SecurityAuditLogRecordType_MicrosoftGraphDataConnectOperation,
		"microsoftmanagedserviceplatform":               SecurityAuditLogRecordType_MicrosoftManagedServicePlatform,
		"microsoftpurview":                              SecurityAuditLogRecordType_MicrosoftPurview,
		"microsoftstream":                               SecurityAuditLogRecordType_MicrosoftStream,
		"microsoftteams":                                SecurityAuditLogRecordType_MicrosoftTeams,
		"microsoftteamsadmin":                           SecurityAuditLogRecordType_MicrosoftTeamsAdmin,
		"microsoftteamsanalytics":                       SecurityAuditLogRecordType_MicrosoftTeamsAnalytics,
		"microsoftteamsdevice":                          SecurityAuditLogRecordType_MicrosoftTeamsDevice,
		"microsoftteamssensitivitylabelaction":          SecurityAuditLogRecordType_MicrosoftTeamsSensitivityLabelAction,
		"microsoftteamsshifts":                          SecurityAuditLogRecordType_MicrosoftTeamsShifts,
		"microsofttodoaudit":                            SecurityAuditLogRecordType_MicrosoftTodoAudit,
		"mipautolabelexchangeitem":                      SecurityAuditLogRecordType_MipAutoLabelExchangeItem,
		"mipautolabelprogressfeedback":                  SecurityAuditLogRecordType_MipAutoLabelProgressFeedback,
		"mipautolabelsharepointitem":                    SecurityAuditLogRecordType_MipAutoLabelSharePointItem,
		"mipautolabelsharepointpolicylocation":          SecurityAuditLogRecordType_MipAutoLabelSharePointPolicyLocation,
		"mipautolabelsimulationcompletion":              SecurityAuditLogRecordType_MipAutoLabelSimulationCompletion,
		"mipautolabelsimulationprogress":                SecurityAuditLogRecordType_MipAutoLabelSimulationProgress,
		"mipautolabelsimulationstatistics":              SecurityAuditLogRecordType_MipAutoLabelSimulationStatistics,
		"mipexactdatamatch":                             SecurityAuditLogRecordType_MipExactDataMatch,
		"miplabelanalyticsauditrecord":                  SecurityAuditLogRecordType_MipLabelAnalyticsAuditRecord,
		"multistagedisposition":                         SecurityAuditLogRecordType_MultiStageDisposition,
		"myanalyticssettings":                           SecurityAuditLogRecordType_MyAnalyticsSettings,
		"omeportal":                                     SecurityAuditLogRecordType_OMEPortal,
		"officenative":                                  SecurityAuditLogRecordType_OfficeNative,
		"officescriptsrunaction":                        SecurityAuditLogRecordType_OfficeScriptsRunAction,
		"onpremisesfilesharescannerdlp":                 SecurityAuditLogRecordType_OnPremisesFileShareScannerDlp,
		"onpremisessharepointscannerdlp":                SecurityAuditLogRecordType_OnPremisesSharePointScannerDlp,
		"onedrive":                                      SecurityAuditLogRecordType_OneDrive,
		"physicalbadgingsignal":                         SecurityAuditLogRecordType_PhysicalBadgingSignal,
		"plannercopyplan":                               SecurityAuditLogRecordType_PlannerCopyPlan,
		"plannerplan":                                   SecurityAuditLogRecordType_PlannerPlan,
		"plannerplanlist":                               SecurityAuditLogRecordType_PlannerPlanList,
		"plannerroster":                                 SecurityAuditLogRecordType_PlannerRoster,
		"plannerrostersensitivitylabel":                 SecurityAuditLogRecordType_PlannerRosterSensitivityLabel,
		"plannertask":                                   SecurityAuditLogRecordType_PlannerTask,
		"plannertasklist":                               SecurityAuditLogRecordType_PlannerTaskList,
		"plannertenantsettings":                         SecurityAuditLogRecordType_PlannerTenantSettings,
		"powerappsapp":                                  SecurityAuditLogRecordType_PowerAppsApp,
		"powerappsplan":                                 SecurityAuditLogRecordType_PowerAppsPlan,
		"powerappsresource":                             SecurityAuditLogRecordType_PowerAppsResource,
		"powerbiaudit":                                  SecurityAuditLogRecordType_PowerBIAudit,
		"powerbidlp":                                    SecurityAuditLogRecordType_PowerBIDlp,
		"powerpagessite":                                SecurityAuditLogRecordType_PowerPagesSite,
		"powerplatformadmindlp":                         SecurityAuditLogRecordType_PowerPlatformAdminDlp,
		"powerplatformadminenvironment":                 SecurityAuditLogRecordType_PowerPlatformAdminEnvironment,
		"powerplatformlockboxresourceaccessrequest":     SecurityAuditLogRecordType_PowerPlatformLockboxResourceAccessRequest,
		"powerplatformlockboxresourcecommand":           SecurityAuditLogRecordType_PowerPlatformLockboxResourceCommand,
		"powerplatformserviceactivity":                  SecurityAuditLogRecordType_PowerPlatformServiceActivity,
		"privacydatamatch":                              SecurityAuditLogRecordType_PrivacyDataMatch,
		"privacydataminimization":                       SecurityAuditLogRecordType_PrivacyDataMinimization,
		"privacydigestemail":                            SecurityAuditLogRecordType_PrivacyDigestEmail,
		"privacyportal":                                 SecurityAuditLogRecordType_PrivacyPortal,
		"privacyremediation":                            SecurityAuditLogRecordType_PrivacyRemediation,
		"privacyremediationaction":                      SecurityAuditLogRecordType_PrivacyRemediationAction,
		"privacytenantaudithistoryrecord":               SecurityAuditLogRecordType_PrivacyTenantAuditHistoryRecord,
		"project":                                       SecurityAuditLogRecordType_Project,
		"projectforthewebproject":                       SecurityAuditLogRecordType_ProjectForTheWebProject,
		"projectforthewebprojectsettings":               SecurityAuditLogRecordType_ProjectForTheWebProjectSettings,
		"projectforthewebroadmap":                       SecurityAuditLogRecordType_ProjectForTheWebRoadmap,
		"projectforthewebroadmapitem":                   SecurityAuditLogRecordType_ProjectForTheWebRoadmapItem,
		"projectforthewebroadmapsettings":               SecurityAuditLogRecordType_ProjectForTheWebRoadmapSettings,
		"projectforthewebtask":                          SecurityAuditLogRecordType_ProjectForTheWebTask,
		"publicfolder":                                  SecurityAuditLogRecordType_PublicFolder,
		"purviewdatamapoperation":                       SecurityAuditLogRecordType_PurviewDataMapOperation,
		"quarantine":                                    SecurityAuditLogRecordType_Quarantine,
		"quarantinemetadata":                            SecurityAuditLogRecordType_QuarantineMetadata,
		"recordsmanagement":                             SecurityAuditLogRecordType_RecordsManagement,
		"scoreplatformgenericauditrecord":               SecurityAuditLogRecordType_ScorePlatformGenericAuditRecord,
		"search":                                        SecurityAuditLogRecordType_Search,
		"securescore":                                   SecurityAuditLogRecordType_SecureScore,
		"securitycompliancealerts":                      SecurityAuditLogRecordType_SecurityComplianceAlerts,
		"securitycompliancecentereopcmdlet":             SecurityAuditLogRecordType_SecurityComplianceCenterEOPCmdlet,
		"securitycomplianceinsights":                    SecurityAuditLogRecordType_SecurityComplianceInsights,
		"securitycompliancerbac":                        SecurityAuditLogRecordType_SecurityComplianceRBAC,
		"securitycomplianceuserchange":                  SecurityAuditLogRecordType_SecurityComplianceUserChange,
		"sensitivitylabelaction":                        SecurityAuditLogRecordType_SensitivityLabelAction,
		"sensitivitylabelpolicymatch":                   SecurityAuditLogRecordType_SensitivityLabelPolicyMatch,
		"sensitivitylabeledfileaction":                  SecurityAuditLogRecordType_SensitivityLabeledFileAction,
		"sharepoint":                                    SecurityAuditLogRecordType_SharePoint,
		"sharepointapppermissionoperation":              SecurityAuditLogRecordType_SharePointAppPermissionOperation,
		"sharepointcommentoperation":                    SecurityAuditLogRecordType_SharePointCommentOperation,
		"sharepointcontenttypeoperation":                SecurityAuditLogRecordType_SharePointContentTypeOperation,
		"sharepointfieldoperation":                      SecurityAuditLogRecordType_SharePointFieldOperation,
		"sharepointfileoperation":                       SecurityAuditLogRecordType_SharePointFileOperation,
		"sharepointlistitemoperation":                   SecurityAuditLogRecordType_SharePointListItemOperation,
		"sharepointlistoperation":                       SecurityAuditLogRecordType_SharePointListOperation,
		"sharepointsearch":                              SecurityAuditLogRecordType_SharePointSearch,
		"sharepointsharingoperation":                    SecurityAuditLogRecordType_SharePointSharingOperation,
		"skypeforbusinesscmdlets":                       SecurityAuditLogRecordType_SkypeForBusinessCmdlets,
		"skypeforbusinesspstnusage":                     SecurityAuditLogRecordType_SkypeForBusinessPSTNUsage,
		"skypeforbusinessusersblocked":                  SecurityAuditLogRecordType_SkypeForBusinessUsersBlocked,
		"supervisoryreviewdayxinsight":                  SecurityAuditLogRecordType_SupervisoryReviewDayXInsight,
		"sway":                                          SecurityAuditLogRecordType_Sway,
		"syntheticprobe":                                SecurityAuditLogRecordType_SyntheticProbe,
		"teamseasyapprovals":                            SecurityAuditLogRecordType_TeamsEasyApprovals,
		"teamshealthcare":                               SecurityAuditLogRecordType_TeamsHealthcare,
		"teamsquarantinemetadata":                       SecurityAuditLogRecordType_TeamsQuarantineMetadata,
		"teamsupdates":                                  SecurityAuditLogRecordType_TeamsUpdates,
		"tenantallowblocklist":                          SecurityAuditLogRecordType_TenantAllowBlockList,
		"threatfinder":                                  SecurityAuditLogRecordType_ThreatFinder,
		"threatintelligence":                            SecurityAuditLogRecordType_ThreatIntelligence,
		"threatintelligenceatpcontent":                  SecurityAuditLogRecordType_ThreatIntelligenceAtpContent,
		"threatintelligenceurl":                         SecurityAuditLogRecordType_ThreatIntelligenceUrl,
		"timetravelfilteringdocmetadata":                SecurityAuditLogRecordType_TimeTravelFilteringDocMetadata,
		"timetravelfilteringdocscan":                    SecurityAuditLogRecordType_TimeTravelFilteringDocScan,
		"unifiedsimulationmatcheditem":                  SecurityAuditLogRecordType_UnifiedSimulationMatchedItem,
		"unifiedsimulationsummary":                      SecurityAuditLogRecordType_UnifiedSimulationSummary,
		"updatequarantinemetadata":                      SecurityAuditLogRecordType_UpdateQuarantineMetadata,
		"usertraining":                                  SecurityAuditLogRecordType_UserTraining,
		"vfamcreatepolicy":                              SecurityAuditLogRecordType_VfamCreatePolicy,
		"vfamdeletepolicy":                              SecurityAuditLogRecordType_VfamDeletePolicy,
		"vfamupdatepolicy":                              SecurityAuditLogRecordType_VfamUpdatePolicy,
		"vivagoals":                                     SecurityAuditLogRecordType_VivaGoals,
		"wdatpalerts":                                   SecurityAuditLogRecordType_WDATPAlerts,
		"webpageactivityendpoint":                       SecurityAuditLogRecordType_WebpageActivityEndpoint,
		"workplaceanalytics":                            SecurityAuditLogRecordType_WorkplaceAnalytics,
		"yammer":                                        SecurityAuditLogRecordType_Yammer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAuditLogRecordType(input)
	return &out, nil
}
