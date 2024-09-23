package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagement{}

type DeviceManagement struct {
	// The date & time when tenant data moved between scaleunits.
	AccountMoveCompletionDateTime *string `json:"accountMoveCompletionDateTime,omitempty"`

	// Admin consent information.
	AdminConsent *AdminConsent `json:"adminConsent,omitempty"`

	// The summary state of ATP onboarding state for this account.
	AdvancedThreatProtectionOnboardingStateSummary *AdvancedThreatProtectionOnboardingStateSummary `json:"advancedThreatProtectionOnboardingStateSummary,omitempty"`

	// Android device owner enrollment profile entities.
	AndroidDeviceOwnerEnrollmentProfiles *[]AndroidDeviceOwnerEnrollmentProfile `json:"androidDeviceOwnerEnrollmentProfiles,omitempty"`

	// Android for Work app configuration schema entities.
	AndroidForWorkAppConfigurationSchemas *[]AndroidForWorkAppConfigurationSchema `json:"androidForWorkAppConfigurationSchemas,omitempty"`

	// Android for Work enrollment profile entities.
	AndroidForWorkEnrollmentProfiles *[]AndroidForWorkEnrollmentProfile `json:"androidForWorkEnrollmentProfiles,omitempty"`

	// The singleton Android for Work settings entity.
	AndroidForWorkSettings *AndroidForWorkSettings `json:"androidForWorkSettings,omitempty"`

	// The singleton Android managed store account enterprise settings entity.
	AndroidManagedStoreAccountEnterpriseSettings *AndroidManagedStoreAccountEnterpriseSettings `json:"androidManagedStoreAccountEnterpriseSettings,omitempty"`

	// Android Enterprise app configuration schema entities.
	AndroidManagedStoreAppConfigurationSchemas *[]AndroidManagedStoreAppConfigurationSchema `json:"androidManagedStoreAppConfigurationSchemas,omitempty"`

	// Apple push notification certificate.
	ApplePushNotificationCertificate *ApplePushNotificationCertificate `json:"applePushNotificationCertificate,omitempty"`

	// Apple user initiated enrollment profiles
	AppleUserInitiatedEnrollmentProfiles *[]AppleUserInitiatedEnrollmentProfile `json:"appleUserInitiatedEnrollmentProfiles,omitempty"`

	// The list of assignment filters
	AssignmentFilters *[]DeviceAndAppManagementAssignmentFilter `json:"assignmentFilters,omitempty"`

	// The Audit Events
	AuditEvents *[]AuditEvent `json:"auditEvents,omitempty"`

	// The list of autopilot events for the tenant.
	AutopilotEvents *[]DeviceManagementAutopilotEvent `json:"autopilotEvents,omitempty"`

	// The Cart To Class Associations.
	CartToClassAssociations *[]CartToClassAssociation `json:"cartToClassAssociations,omitempty"`

	// The available categories
	Categories *[]DeviceManagementSettingCategory `json:"categories,omitempty"`

	// Collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
	CertificateConnectorDetails *[]CertificateConnectorDetails `json:"certificateConnectorDetails,omitempty"`

	// Collection of ChromeOSOnboardingSettings settings associated with account.
	ChromeOSOnboardingSettings *[]ChromeOSOnboardingSettings `json:"chromeOSOnboardingSettings,omitempty"`

	// The list of CloudPC Connectivity Issue.
	CloudPCConnectivityIssues *[]CloudPCConnectivityIssue `json:"cloudPCConnectivityIssues,omitempty"`

	// The list of co-managed devices report
	ComanagedDevices *[]ManagedDevice `json:"comanagedDevices,omitempty"`

	// The list of co-management eligible devices report
	ComanagementEligibleDevices *[]ComanagementEligibleDevice `json:"comanagementEligibleDevices,omitempty"`

	// List of all compliance categories
	ComplianceCategories *[]DeviceManagementConfigurationCategory `json:"complianceCategories,omitempty"`

	// The list of Compliance Management Partners configured by the tenant.
	ComplianceManagementPartners *[]ComplianceManagementPartner `json:"complianceManagementPartners,omitempty"`

	// List of all compliance policies
	CompliancePolicies *[]DeviceManagementCompliancePolicy `json:"compliancePolicies,omitempty"`

	// List of all ComplianceSettings
	ComplianceSettings *[]DeviceManagementConfigurationSettingDefinition `json:"complianceSettings,omitempty"`

	// The Exchange on premises conditional access settings. On premises conditional access will require devices to be both
	// enrolled and compliant for mail access
	ConditionalAccessSettings *OnPremisesConditionalAccessSettings `json:"conditionalAccessSettings,omitempty"`

	// A list of ConfigManagerCollection
	ConfigManagerCollections *[]ConfigManagerCollection `json:"configManagerCollections,omitempty"`

	// List of all Configuration Categories
	ConfigurationCategories *[]DeviceManagementConfigurationCategory `json:"configurationCategories,omitempty"`

	// List of all Configuration policies
	ConfigurationPolicies *[]DeviceManagementConfigurationPolicy `json:"configurationPolicies,omitempty"`

	// List of all templates
	ConfigurationPolicyTemplates *[]DeviceManagementConfigurationPolicyTemplate `json:"configurationPolicyTemplates,omitempty"`

	// List of all ConfigurationSettings
	ConfigurationSettings *[]DeviceManagementConfigurationSettingDefinition `json:"configurationSettings,omitempty"`

	// The list of connector status for the tenant.
	ConnectorStatus *[]ConnectorStatusDetails `json:"connectorStatus,omitempty"`

	// A configuration entity for MEM features that utilize Data Processor Service for Windows (DPSW) data.
	DataProcessorServiceForWindowsFeaturesOnboarding *DataProcessorServiceForWindowsFeaturesOnboarding `json:"dataProcessorServiceForWindowsFeaturesOnboarding,omitempty"`

	// Data sharing consents.
	DataSharingConsents *[]DataSharingConsent `json:"dataSharingConsents,omitempty"`

	// This collections of multiple DEP tokens per-tenant.
	DepOnboardingSettings *[]DepOnboardingSetting `json:"depOnboardingSettings,omitempty"`

	// Collection of Derived credential settings associated with account.
	DerivedCredentials *[]DeviceManagementDerivedCredentialSettings `json:"derivedCredentials,omitempty"`

	// The list of detected apps associated with a device.
	DetectedApps *[]DetectedApp `json:"detectedApps,omitempty"`

	// The list of device categories with the tenant.
	DeviceCategories *[]DeviceCategory `json:"deviceCategories,omitempty"`

	// The device compliance policies.
	DeviceCompliancePolicies *[]DeviceCompliancePolicy `json:"deviceCompliancePolicies,omitempty"`

	// The device compliance state summary for this account.
	DeviceCompliancePolicyDeviceStateSummary *DeviceCompliancePolicyDeviceStateSummary `json:"deviceCompliancePolicyDeviceStateSummary,omitempty"`

	// The summary states of compliance policy settings for this account.
	DeviceCompliancePolicySettingStateSummaries *[]DeviceCompliancePolicySettingStateSummary `json:"deviceCompliancePolicySettingStateSummaries,omitempty"`

	// The last requested time of device compliance reporting for this account. This property is read-only.
	DeviceComplianceReportSummarizationDateTime *string `json:"deviceComplianceReportSummarizationDateTime,omitempty"`

	// The list of device compliance scripts associated with the tenant.
	DeviceComplianceScripts *[]DeviceComplianceScript `json:"deviceComplianceScripts,omitempty"`

	// Summary of policies in conflict state for this account.
	DeviceConfigurationConflictSummary *[]DeviceConfigurationConflictSummary `json:"deviceConfigurationConflictSummary,omitempty"`

	// The device configuration device state summary for this account.
	DeviceConfigurationDeviceStateSummaries *DeviceConfigurationDeviceStateSummary `json:"deviceConfigurationDeviceStateSummaries,omitempty"`

	// Profile Id of the object.
	DeviceConfigurationProfiles *[]DeviceConfigurationProfile `json:"deviceConfigurationProfiles,omitempty"`

	// Restricted apps violations for this account.
	DeviceConfigurationRestrictedAppsViolations *[]RestrictedAppsViolation `json:"deviceConfigurationRestrictedAppsViolations,omitempty"`

	// The device configuration user state summary for this account.
	DeviceConfigurationUserStateSummaries *DeviceConfigurationUserStateSummary `json:"deviceConfigurationUserStateSummaries,omitempty"`

	// The device configurations.
	DeviceConfigurations *[]DeviceConfiguration `json:"deviceConfigurations,omitempty"`

	// Summary of all certificates for all devices.
	DeviceConfigurationsAllManagedDeviceCertificateStates *[]ManagedAllDeviceCertificateState `json:"deviceConfigurationsAllManagedDeviceCertificateStates,omitempty"`

	// The list of device custom attribute shell scripts associated with the tenant.
	DeviceCustomAttributeShellScripts *[]DeviceCustomAttributeShellScript `json:"deviceCustomAttributeShellScripts,omitempty"`

	// The list of device enrollment configurations
	DeviceEnrollmentConfigurations *[]DeviceEnrollmentConfiguration `json:"deviceEnrollmentConfigurations,omitempty"`

	// The list of device health scripts associated with the tenant.
	DeviceHealthScripts *[]DeviceHealthScript `json:"deviceHealthScripts,omitempty"`

	// The list of Device Management Partners configured by the tenant.
	DeviceManagementPartners *[]DeviceManagementPartner `json:"deviceManagementPartners,omitempty"`

	// The list of device management scripts associated with the tenant.
	DeviceManagementScripts *[]DeviceManagementScript `json:"deviceManagementScripts,omitempty"`

	// Device protection overview.
	DeviceProtectionOverview *DeviceProtectionOverview `json:"deviceProtectionOverview,omitempty"`

	// The list of device shell scripts associated with the tenant.
	DeviceShellScripts *[]DeviceShellScript `json:"deviceShellScripts,omitempty"`

	// A list of connector objects.
	DomainJoinConnectors *[]DeviceManagementDomainJoinConnector `json:"domainJoinConnectors,omitempty"`

	// List of elevation requests
	ElevationRequests *[]PrivilegeManagementElevationRequest `json:"elevationRequests,omitempty"`

	// The embedded SIM activation code pools created by this account.
	EmbeddedSIMActivationCodePools *[]EmbeddedSIMActivationCodePool `json:"embeddedSIMActivationCodePools,omitempty"`

	// Endpoint privilege management (EPM) tenant provisioning status contains tenant level license and onboarding state
	// information.
	EndpointPrivilegeManagementProvisioningStatus *EndpointPrivilegeManagementProvisioningStatus `json:"endpointPrivilegeManagementProvisioningStatus,omitempty"`

	// The list of Exchange Connectors configured by the tenant.
	ExchangeConnectors *[]DeviceManagementExchangeConnector `json:"exchangeConnectors,omitempty"`

	// The list of Exchange On Premisis policies configured by the tenant.
	ExchangeOnPremisesPolicies *[]DeviceManagementExchangeOnPremisesPolicy `json:"exchangeOnPremisesPolicies,omitempty"`

	// The policy which controls mobile device access to Exchange On Premises
	ExchangeOnPremisesPolicy *DeviceManagementExchangeOnPremisesPolicy `json:"exchangeOnPremisesPolicy,omitempty"`

	// The available group policy categories for this account.
	GroupPolicyCategories *[]GroupPolicyCategory `json:"groupPolicyCategories,omitempty"`

	// The group policy configurations created by this account.
	GroupPolicyConfigurations *[]GroupPolicyConfiguration `json:"groupPolicyConfigurations,omitempty"`

	// The available group policy definition files for this account.
	GroupPolicyDefinitionFiles *[]GroupPolicyDefinitionFile `json:"groupPolicyDefinitionFiles,omitempty"`

	// The available group policy definitions for this account.
	GroupPolicyDefinitions *[]GroupPolicyDefinition `json:"groupPolicyDefinitions,omitempty"`

	// A list of Group Policy migration reports.
	GroupPolicyMigrationReports *[]GroupPolicyMigrationReport `json:"groupPolicyMigrationReports,omitempty"`

	// A list of Group Policy Object files uploaded.
	GroupPolicyObjectFiles *[]GroupPolicyObjectFile `json:"groupPolicyObjectFiles,omitempty"`

	// The available group policy uploaded definition files for this account.
	GroupPolicyUploadedDefinitionFiles *[]GroupPolicyUploadedDefinitionFile `json:"groupPolicyUploadedDefinitionFiles,omitempty"`

	// BIOS configuration and other settings provides customers the ability to configure hardware/bios settings on the
	// enrolled Windows 10/11 Entra ID joined devices by uploading a configuration file generated with their OEM tool (e.g.
	// Dell Command tool). A BIOS configuration policy can be assigned to multiple devices, allowing admins to remotely
	// control a device's hardware properties (e.g. enable Secure Boot) from the Intune Portal. Supported for Dell only at
	// this time.
	HardwareConfigurations *[]HardwareConfiguration `json:"hardwareConfigurations,omitempty"`

	// Device BIOS password information for devices with managed BIOS and firmware configuration, which provides device
	// serial number, list of previous passwords, and current password.
	HardwarePasswordDetails *[]HardwarePasswordDetail `json:"hardwarePasswordDetails,omitempty"`

	// Intune will provide customer the ability to configure hardware/bios settings on the enrolled windows 10 Azure Active
	// Directory joined devices. Starting from June, 2024 (Intune Release 2406), this type will no longer be supported and
	// will be marked as deprecated
	HardwarePasswordInfo *[]HardwarePasswordInfo `json:"hardwarePasswordInfo,omitempty"`

	// The imported device identities.
	ImportedDeviceIdentities *[]ImportedDeviceIdentity `json:"importedDeviceIdentities,omitempty"`

	// Collection of imported Windows autopilot devices.
	ImportedWindowsAutopilotDeviceIdentities *[]ImportedWindowsAutopilotDeviceIdentity `json:"importedWindowsAutopilotDeviceIdentities,omitempty"`

	// The device management intents
	Intents *[]DeviceManagementIntent `json:"intents,omitempty"`

	// Intune Account ID for given tenant
	IntuneAccountId *string `json:"intuneAccountId,omitempty"`

	// intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as
	// the end user web portal.
	IntuneBrand *IntuneBrand `json:"intuneBrand,omitempty"`

	// Intune branding profiles targeted to AAD groups
	IntuneBrandingProfiles *[]IntuneBrandingProfile `json:"intuneBrandingProfiles,omitempty"`

	// The IOS software update installation statuses for this account.
	IosUpdateStatuses *[]IosUpdateDeviceStatus `json:"iosUpdateStatuses,omitempty"`

	// The last modified time of reporting for this account. This property is read-only.
	LastReportAggregationDateTime *string `json:"lastReportAggregationDateTime,omitempty"`

	// The property to enable Non-MDM managed legacy PC management for this account. This property is read-only.
	LegacyPcManangementEnabled *bool `json:"legacyPcManangementEnabled,omitempty"`

	// The MacOS software update account summaries for this account.
	MacOSSoftwareUpdateAccountSummaries *[]MacOSSoftwareUpdateAccountSummary `json:"macOSSoftwareUpdateAccountSummaries,omitempty"`

	// Device cleanup rule V2
	ManagedDeviceCleanupRules *[]ManagedDeviceCleanupRule `json:"managedDeviceCleanupRules,omitempty"`

	// Device cleanup rule
	ManagedDeviceCleanupSettings *ManagedDeviceCleanupSettings `json:"managedDeviceCleanupSettings,omitempty"`

	// Encryption report for devices in this account
	ManagedDeviceEncryptionStates *[]ManagedDeviceEncryptionState `json:"managedDeviceEncryptionStates,omitempty"`

	// Device overview
	ManagedDeviceOverview *ManagedDeviceOverview `json:"managedDeviceOverview,omitempty"`

	// A list of ManagedDeviceWindowsOperatingSystemImages
	ManagedDeviceWindowsOSImages *[]ManagedDeviceWindowsOperatingSystemImage `json:"managedDeviceWindowsOSImages,omitempty"`

	// The list of managed devices.
	ManagedDevices *[]ManagedDevice `json:"managedDevices,omitempty"`

	// Maximum number of DEP tokens allowed per-tenant.
	MaximumDepTokens *int64 `json:"maximumDepTokens,omitempty"`

	// Collection of MicrosoftTunnelConfiguration settings associated with account.
	MicrosoftTunnelConfigurations *[]MicrosoftTunnelConfiguration `json:"microsoftTunnelConfigurations,omitempty"`

	// Collection of MicrosoftTunnelHealthThreshold settings associated with account.
	MicrosoftTunnelHealthThresholds *[]MicrosoftTunnelHealthThreshold `json:"microsoftTunnelHealthThresholds,omitempty"`

	// Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
	MicrosoftTunnelServerLogCollectionResponses *[]MicrosoftTunnelServerLogCollectionResponse `json:"microsoftTunnelServerLogCollectionResponses,omitempty"`

	// Collection of MicrosoftTunnelSite settings associated with account.
	MicrosoftTunnelSites *[]MicrosoftTunnelSite `json:"microsoftTunnelSites,omitempty"`

	// The collection property of MobileAppTroubleshootingEvent.
	MobileAppTroubleshootingEvents *[]MobileAppTroubleshootingEvent `json:"mobileAppTroubleshootingEvents,omitempty"`

	// The list of Mobile threat Defense connectors configured by the tenant.
	MobileThreatDefenseConnectors *[]MobileThreatDefenseConnector `json:"mobileThreatDefenseConnectors,omitempty"`

	Monitoring *DeviceManagementMonitoring `json:"monitoring,omitempty"`

	// The collection of Ndes connectors for this account.
	NdesConnectors *[]NdesConnector `json:"ndesConnectors,omitempty"`

	// The Notification Message Templates.
	NotificationMessageTemplates *[]NotificationMessageTemplate `json:"notificationMessageTemplates,omitempty"`

	// The Operation Approval Policies
	OperationApprovalPolicies *[]OperationApprovalPolicy `json:"operationApprovalPolicies,omitempty"`

	// The Operation Approval Requests
	OperationApprovalRequests *[]OperationApprovalRequest `json:"operationApprovalRequests,omitempty"`

	// The endpoint privilege management elevation event entity contains elevation details.
	PrivilegeManagementElevations *[]PrivilegeManagementElevation `json:"privilegeManagementElevations,omitempty"`

	// The list of device remote action audits with the tenant.
	RemoteActionAudits *[]RemoteActionAudit `json:"remoteActionAudits,omitempty"`

	// The remote assist partners.
	RemoteAssistancePartners *[]RemoteAssistancePartner `json:"remoteAssistancePartners,omitempty"`

	// The remote assistance settings singleton
	RemoteAssistanceSettings *RemoteAssistanceSettings `json:"remoteAssistanceSettings,omitempty"`

	// Reports singleton
	Reports *DeviceManagementReports `json:"reports,omitempty"`

	// Collection of resource access settings associated with account.
	ResourceAccessProfiles *[]DeviceManagementResourceAccessProfileBase `json:"resourceAccessProfiles,omitempty"`

	// The Resource Operations.
	ResourceOperations *[]ResourceOperation `json:"resourceOperations,omitempty"`

	// List of all reusable settings that can be referred in a policy
	ReusablePolicySettings *[]DeviceManagementReusablePolicySetting `json:"reusablePolicySettings,omitempty"`

	// List of all reusable settings
	ReusableSettings *[]DeviceManagementConfigurationSettingDefinition `json:"reusableSettings,omitempty"`

	// The Role Assignments.
	RoleAssignments *[]DeviceAndAppManagementRoleAssignment `json:"roleAssignments,omitempty"`

	// The Role Definitions.
	RoleDefinitions *[]RoleDefinition `json:"roleDefinitions,omitempty"`

	// The Role Scope Tags.
	RoleScopeTags *[]RoleScopeTag `json:"roleScopeTags,omitempty"`

	// A list of ServiceNowConnections
	ServiceNowConnections *[]ServiceNowConnection `json:"serviceNowConnections,omitempty"`

	// The device management intent setting definitions
	SettingDefinitions *[]DeviceManagementSettingDefinition `json:"settingDefinitions,omitempty"`

	// Account level settings.
	Settings *DeviceManagementSettings `json:"settings,omitempty"`

	// The software update status summary.
	SoftwareUpdateStatusSummary *SoftwareUpdateStatusSummary `json:"softwareUpdateStatusSummary,omitempty"`

	// Tenant mobile device management subscription state.
	SubscriptionState *DeviceManagementSubscriptionState `json:"subscriptionState,omitempty"`

	// Tenant mobile device management subscriptions.
	Subscriptions *DeviceManagementSubscriptions `json:"subscriptions,omitempty"`

	// The telecom expense management partners.
	TelecomExpenseManagementPartners *[]TelecomExpenseManagementPartner `json:"telecomExpenseManagementPartners,omitempty"`

	// List of setting insights in a template
	TemplateInsights *[]DeviceManagementTemplateInsightsDefinition `json:"templateInsights,omitempty"`

	// List of all TemplateSettings
	TemplateSettings *[]DeviceManagementConfigurationSettingTemplate `json:"templateSettings,omitempty"`

	// The available templates
	Templates *[]DeviceManagementTemplate `json:"templates,omitempty"`

	// TenantAttach RBAC Enablement
	TenantAttachRBAC *TenantAttachRBAC `json:"tenantAttachRBAC,omitempty"`

	// The terms and conditions associated with device management of the company.
	TermsAndConditions *[]TermsAndConditions `json:"termsAndConditions,omitempty"`

	// The list of troubleshooting events for the tenant.
	TroubleshootingEvents *[]DeviceManagementTroubleshootingEvent `json:"troubleshootingEvents,omitempty"`

	// When enabled, users assigned as administrators via Role Assignment Memberships do not require an assigned Intune
	// license. Prior to this, only Intune licensed users were granted permissions with an Intune role unless they were
	// assigned a role via Azure Active Directory. You are limited to 350 unlicensed direct members for each AAD security
	// group in a role assignment, but you can assign multiple AAD security groups to a role if you need to support more
	// than 350 unlicensed administrators. Licensed administrators are unaffected, do not have to be direct members, nor
	// does the 350 member limit apply. This property is read-only.
	UnlicensedAdminstratorsEnabled *bool `json:"unlicensedAdminstratorsEnabled,omitempty"`

	// The user experience analytics anomaly entity contains anomaly details.
	UserExperienceAnalyticsAnomaly *[]UserExperienceAnalyticsAnomaly `json:"userExperienceAnalyticsAnomaly,omitempty"`

	// The user experience analytics anomaly correlation group overview entity contains the information for each correlation
	// group of an anomaly.
	UserExperienceAnalyticsAnomalyCorrelationGroupOverview *[]UserExperienceAnalyticsAnomalyCorrelationGroupOverview `json:"userExperienceAnalyticsAnomalyCorrelationGroupOverview,omitempty"`

	// The user experience analytics anomaly entity contains device details.
	UserExperienceAnalyticsAnomalyDevice *[]UserExperienceAnalyticsAnomalyDevice `json:"userExperienceAnalyticsAnomalyDevice,omitempty"`

	// The user experience analytics anomaly severity overview entity contains the count information for each severity of
	// anomaly.
	UserExperienceAnalyticsAnomalySeverityOverview *UserExperienceAnalyticsAnomalySeverityOverview `json:"userExperienceAnalyticsAnomalySeverityOverview,omitempty"`

	// User experience analytics appHealth Application Performance
	UserExperienceAnalyticsAppHealthApplicationPerformance *[]UserExperienceAnalyticsAppHealthApplicationPerformance `json:"userExperienceAnalyticsAppHealthApplicationPerformance,omitempty"`

	// User experience analytics appHealth Application Performance by App Version
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion,omitempty"`

	// User experience analytics appHealth Application Performance by App Version details
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails,omitempty"`

	// User experience analytics appHealth Application Performance by App Version Device Id
	UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId,omitempty"`

	// User experience analytics appHealth Application Performance by OS Version
	UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion *[]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion,omitempty"`

	// User experience analytics appHealth Model Performance
	UserExperienceAnalyticsAppHealthDeviceModelPerformance *[]UserExperienceAnalyticsAppHealthDeviceModelPerformance `json:"userExperienceAnalyticsAppHealthDeviceModelPerformance,omitempty"`

	// User experience analytics appHealth Device Performance
	UserExperienceAnalyticsAppHealthDevicePerformance *[]UserExperienceAnalyticsAppHealthDevicePerformance `json:"userExperienceAnalyticsAppHealthDevicePerformance,omitempty"`

	// User experience analytics device performance details
	UserExperienceAnalyticsAppHealthDevicePerformanceDetails *[]UserExperienceAnalyticsAppHealthDevicePerformanceDetails `json:"userExperienceAnalyticsAppHealthDevicePerformanceDetails,omitempty"`

	// User experience analytics appHealth OS version Performance
	UserExperienceAnalyticsAppHealthOSVersionPerformance *[]UserExperienceAnalyticsAppHealthOSVersionPerformance `json:"userExperienceAnalyticsAppHealthOSVersionPerformance,omitempty"`

	// User experience analytics appHealth overview
	UserExperienceAnalyticsAppHealthOverview *UserExperienceAnalyticsCategory `json:"userExperienceAnalyticsAppHealthOverview,omitempty"`

	// User experience analytics baselines
	UserExperienceAnalyticsBaselines *[]UserExperienceAnalyticsBaseline `json:"userExperienceAnalyticsBaselines,omitempty"`

	// User Experience Analytics Battery Health App Impact
	UserExperienceAnalyticsBatteryHealthAppImpact *[]UserExperienceAnalyticsBatteryHealthAppImpact `json:"userExperienceAnalyticsBatteryHealthAppImpact,omitempty"`

	// User Experience Analytics Battery Health Capacity Details
	UserExperienceAnalyticsBatteryHealthCapacityDetails *UserExperienceAnalyticsBatteryHealthCapacityDetails `json:"userExperienceAnalyticsBatteryHealthCapacityDetails,omitempty"`

	// User Experience Analytics Battery Health Device App Impact
	UserExperienceAnalyticsBatteryHealthDeviceAppImpact *[]UserExperienceAnalyticsBatteryHealthDeviceAppImpact `json:"userExperienceAnalyticsBatteryHealthDeviceAppImpact,omitempty"`

	// User Experience Analytics Battery Health Device Performance
	UserExperienceAnalyticsBatteryHealthDevicePerformance *[]UserExperienceAnalyticsBatteryHealthDevicePerformance `json:"userExperienceAnalyticsBatteryHealthDevicePerformance,omitempty"`

	// User Experience Analytics Battery Health Device Runtime History
	UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory *[]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory `json:"userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory,omitempty"`

	// User Experience Analytics Battery Health Model Performance
	UserExperienceAnalyticsBatteryHealthModelPerformance *[]UserExperienceAnalyticsBatteryHealthModelPerformance `json:"userExperienceAnalyticsBatteryHealthModelPerformance,omitempty"`

	// User Experience Analytics Battery Health Os Performance
	UserExperienceAnalyticsBatteryHealthOsPerformance *[]UserExperienceAnalyticsBatteryHealthOsPerformance `json:"userExperienceAnalyticsBatteryHealthOsPerformance,omitempty"`

	// User Experience Analytics Battery Health Runtime Details
	UserExperienceAnalyticsBatteryHealthRuntimeDetails *UserExperienceAnalyticsBatteryHealthRuntimeDetails `json:"userExperienceAnalyticsBatteryHealthRuntimeDetails,omitempty"`

	// User experience analytics categories
	UserExperienceAnalyticsCategories *[]UserExperienceAnalyticsCategory `json:"userExperienceAnalyticsCategories,omitempty"`

	// User experience analytics device metric history
	UserExperienceAnalyticsDeviceMetricHistory *[]UserExperienceAnalyticsMetricHistory `json:"userExperienceAnalyticsDeviceMetricHistory,omitempty"`

	// User experience analytics device performance
	UserExperienceAnalyticsDevicePerformance *[]UserExperienceAnalyticsDevicePerformance `json:"userExperienceAnalyticsDevicePerformance,omitempty"`

	// The user experience analytics device scope entity endpoint to trigger on the service to either START or STOP
	// computing metrics data based on a device scope configuration.
	UserExperienceAnalyticsDeviceScope *UserExperienceAnalyticsDeviceScope `json:"userExperienceAnalyticsDeviceScope,omitempty"`

	// The user experience analytics device scope entity contains device scope configuration use to apply filtering on the
	// endpoint analytics reports.
	UserExperienceAnalyticsDeviceScopes *[]UserExperienceAnalyticsDeviceScope `json:"userExperienceAnalyticsDeviceScopes,omitempty"`

	// User experience analytics device scores
	UserExperienceAnalyticsDeviceScores *[]UserExperienceAnalyticsDeviceScores `json:"userExperienceAnalyticsDeviceScores,omitempty"`

	// User experience analytics device Startup History
	UserExperienceAnalyticsDeviceStartupHistory *[]UserExperienceAnalyticsDeviceStartupHistory `json:"userExperienceAnalyticsDeviceStartupHistory,omitempty"`

	// User experience analytics device Startup Process Performance
	UserExperienceAnalyticsDeviceStartupProcessPerformance *[]UserExperienceAnalyticsDeviceStartupProcessPerformance `json:"userExperienceAnalyticsDeviceStartupProcessPerformance,omitempty"`

	// User experience analytics device Startup Processes
	UserExperienceAnalyticsDeviceStartupProcesses *[]UserExperienceAnalyticsDeviceStartupProcess `json:"userExperienceAnalyticsDeviceStartupProcesses,omitempty"`

	// The user experience analytics device events entity contains NRT device timeline event details.
	UserExperienceAnalyticsDeviceTimelineEvent *[]UserExperienceAnalyticsDeviceTimelineEvent `json:"userExperienceAnalyticsDeviceTimelineEvent,omitempty"`

	// User experience analytics devices without cloud identity.
	UserExperienceAnalyticsDevicesWithoutCloudIdentity *[]UserExperienceAnalyticsDeviceWithoutCloudIdentity `json:"userExperienceAnalyticsDevicesWithoutCloudIdentity,omitempty"`

	// User experience analytics impacting process
	UserExperienceAnalyticsImpactingProcess *[]UserExperienceAnalyticsImpactingProcess `json:"userExperienceAnalyticsImpactingProcess,omitempty"`

	// User experience analytics metric history
	UserExperienceAnalyticsMetricHistory *[]UserExperienceAnalyticsMetricHistory `json:"userExperienceAnalyticsMetricHistory,omitempty"`

	// User experience analytics model scores
	UserExperienceAnalyticsModelScores *[]UserExperienceAnalyticsModelScores `json:"userExperienceAnalyticsModelScores,omitempty"`

	// User experience analytics devices not Windows Autopilot ready.
	UserExperienceAnalyticsNotAutopilotReadyDevice *[]UserExperienceAnalyticsNotAutopilotReadyDevice `json:"userExperienceAnalyticsNotAutopilotReadyDevice,omitempty"`

	// User experience analytics overview
	UserExperienceAnalyticsOverview *UserExperienceAnalyticsOverview `json:"userExperienceAnalyticsOverview,omitempty"`

	// User experience analytics remote connection
	UserExperienceAnalyticsRemoteConnection *[]UserExperienceAnalyticsRemoteConnection `json:"userExperienceAnalyticsRemoteConnection,omitempty"`

	// User experience analytics resource performance
	UserExperienceAnalyticsResourcePerformance *[]UserExperienceAnalyticsResourcePerformance `json:"userExperienceAnalyticsResourcePerformance,omitempty"`

	// User experience analytics device Startup Score History
	UserExperienceAnalyticsScoreHistory *[]UserExperienceAnalyticsScoreHistory `json:"userExperienceAnalyticsScoreHistory,omitempty"`

	// User experience analytics device settings
	UserExperienceAnalyticsSettings *UserExperienceAnalyticsSettings `json:"userExperienceAnalyticsSettings,omitempty"`

	// User experience analytics work from anywhere hardware readiness metrics.
	UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric `json:"userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric,omitempty"`

	// User experience analytics work from anywhere metrics.
	UserExperienceAnalyticsWorkFromAnywhereMetrics *[]UserExperienceAnalyticsWorkFromAnywhereMetric `json:"userExperienceAnalyticsWorkFromAnywhereMetrics,omitempty"`

	// The user experience analytics work from anywhere model performance
	UserExperienceAnalyticsWorkFromAnywhereModelPerformance *[]UserExperienceAnalyticsWorkFromAnywhereModelPerformance `json:"userExperienceAnalyticsWorkFromAnywhereModelPerformance,omitempty"`

	// Collection of PFX certificates associated with a user.
	UserPfxCertificates *[]UserPFXCertificate `json:"userPfxCertificates,omitempty"`

	VirtualEndpoint *VirtualEndpoint `json:"virtualEndpoint,omitempty"`

	// Windows auto pilot deployment profiles
	WindowsAutopilotDeploymentProfiles *[]WindowsAutopilotDeploymentProfile `json:"windowsAutopilotDeploymentProfiles,omitempty"`

	// The Windows autopilot device identities contained collection.
	WindowsAutopilotDeviceIdentities *[]WindowsAutopilotDeviceIdentity `json:"windowsAutopilotDeviceIdentities,omitempty"`

	// The Windows autopilot account settings.
	WindowsAutopilotSettings *WindowsAutopilotSettings `json:"windowsAutopilotSettings,omitempty"`

	// A collection of windows driver update profiles
	WindowsDriverUpdateProfiles *[]WindowsDriverUpdateProfile `json:"windowsDriverUpdateProfiles,omitempty"`

	// A collection of windows feature update profiles
	WindowsFeatureUpdateProfiles *[]WindowsFeatureUpdateProfile `json:"windowsFeatureUpdateProfiles,omitempty"`

	// The windows information protection app learning summaries.
	WindowsInformationProtectionAppLearningSummaries *[]WindowsInformationProtectionAppLearningSummary `json:"windowsInformationProtectionAppLearningSummaries,omitempty"`

	// The windows information protection network learning summaries.
	WindowsInformationProtectionNetworkLearningSummaries *[]WindowsInformationProtectionNetworkLearningSummary `json:"windowsInformationProtectionNetworkLearningSummaries,omitempty"`

	// The list of affected malware in the tenant.
	WindowsMalwareInformation *[]WindowsMalwareInformation `json:"windowsMalwareInformation,omitempty"`

	// Malware overview for windows devices.
	WindowsMalwareOverview *WindowsMalwareOverview `json:"windowsMalwareOverview,omitempty"`

	// A collection of Windows quality update policies
	WindowsQualityUpdatePolicies *[]WindowsQualityUpdatePolicy `json:"windowsQualityUpdatePolicies,omitempty"`

	// A collection of windows quality update profiles
	WindowsQualityUpdateProfiles *[]WindowsQualityUpdateProfile `json:"windowsQualityUpdateProfiles,omitempty"`

	// A collection of windows update catalog items (fetaure updates item , quality updates item)
	WindowsUpdateCatalogItems *[]WindowsUpdateCatalogItem `json:"windowsUpdateCatalogItems,omitempty"`

	// The Collection of ZebraFotaArtifacts.
	ZebraFotaArtifacts *[]ZebraFotaArtifact `json:"zebraFotaArtifacts,omitempty"`

	// The singleton ZebraFotaConnector associated with account.
	ZebraFotaConnector *ZebraFotaConnector `json:"zebraFotaConnector,omitempty"`

	// Collection of ZebraFotaDeployments associated with account.
	ZebraFotaDeployments *[]ZebraFotaDeployment `json:"zebraFotaDeployments,omitempty"`

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

func (s DeviceManagement) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagement{}

func (s DeviceManagement) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagement
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagement: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagement: %+v", err)
	}

	delete(decoded, "deviceComplianceReportSummarizationDateTime")
	delete(decoded, "lastReportAggregationDateTime")
	delete(decoded, "legacyPcManangementEnabled")
	delete(decoded, "unlicensedAdminstratorsEnabled")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagement"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagement: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DeviceManagement{}

func (s *DeviceManagement) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountMoveCompletionDateTime                                              *string                                                               `json:"accountMoveCompletionDateTime,omitempty"`
		AdminConsent                                                               *AdminConsent                                                         `json:"adminConsent,omitempty"`
		AdvancedThreatProtectionOnboardingStateSummary                             *AdvancedThreatProtectionOnboardingStateSummary                       `json:"advancedThreatProtectionOnboardingStateSummary,omitempty"`
		AndroidDeviceOwnerEnrollmentProfiles                                       *[]AndroidDeviceOwnerEnrollmentProfile                                `json:"androidDeviceOwnerEnrollmentProfiles,omitempty"`
		AndroidForWorkAppConfigurationSchemas                                      *[]AndroidForWorkAppConfigurationSchema                               `json:"androidForWorkAppConfigurationSchemas,omitempty"`
		AndroidForWorkEnrollmentProfiles                                           *[]AndroidForWorkEnrollmentProfile                                    `json:"androidForWorkEnrollmentProfiles,omitempty"`
		AndroidForWorkSettings                                                     *AndroidForWorkSettings                                               `json:"androidForWorkSettings,omitempty"`
		AndroidManagedStoreAccountEnterpriseSettings                               *AndroidManagedStoreAccountEnterpriseSettings                         `json:"androidManagedStoreAccountEnterpriseSettings,omitempty"`
		AndroidManagedStoreAppConfigurationSchemas                                 *[]AndroidManagedStoreAppConfigurationSchema                          `json:"androidManagedStoreAppConfigurationSchemas,omitempty"`
		ApplePushNotificationCertificate                                           *ApplePushNotificationCertificate                                     `json:"applePushNotificationCertificate,omitempty"`
		AppleUserInitiatedEnrollmentProfiles                                       *[]AppleUserInitiatedEnrollmentProfile                                `json:"appleUserInitiatedEnrollmentProfiles,omitempty"`
		AuditEvents                                                                *[]AuditEvent                                                         `json:"auditEvents,omitempty"`
		AutopilotEvents                                                            *[]DeviceManagementAutopilotEvent                                     `json:"autopilotEvents,omitempty"`
		CartToClassAssociations                                                    *[]CartToClassAssociation                                             `json:"cartToClassAssociations,omitempty"`
		CertificateConnectorDetails                                                *[]CertificateConnectorDetails                                        `json:"certificateConnectorDetails,omitempty"`
		ChromeOSOnboardingSettings                                                 *[]ChromeOSOnboardingSettings                                         `json:"chromeOSOnboardingSettings,omitempty"`
		CloudPCConnectivityIssues                                                  *[]CloudPCConnectivityIssue                                           `json:"cloudPCConnectivityIssues,omitempty"`
		ComanagementEligibleDevices                                                *[]ComanagementEligibleDevice                                         `json:"comanagementEligibleDevices,omitempty"`
		ComplianceCategories                                                       *[]DeviceManagementConfigurationCategory                              `json:"complianceCategories,omitempty"`
		ComplianceManagementPartners                                               *[]ComplianceManagementPartner                                        `json:"complianceManagementPartners,omitempty"`
		CompliancePolicies                                                         *[]DeviceManagementCompliancePolicy                                   `json:"compliancePolicies,omitempty"`
		ConditionalAccessSettings                                                  *OnPremisesConditionalAccessSettings                                  `json:"conditionalAccessSettings,omitempty"`
		ConfigManagerCollections                                                   *[]ConfigManagerCollection                                            `json:"configManagerCollections,omitempty"`
		ConfigurationCategories                                                    *[]DeviceManagementConfigurationCategory                              `json:"configurationCategories,omitempty"`
		ConfigurationPolicies                                                      *[]DeviceManagementConfigurationPolicy                                `json:"configurationPolicies,omitempty"`
		ConfigurationPolicyTemplates                                               *[]DeviceManagementConfigurationPolicyTemplate                        `json:"configurationPolicyTemplates,omitempty"`
		ConnectorStatus                                                            *[]ConnectorStatusDetails                                             `json:"connectorStatus,omitempty"`
		DataProcessorServiceForWindowsFeaturesOnboarding                           *DataProcessorServiceForWindowsFeaturesOnboarding                     `json:"dataProcessorServiceForWindowsFeaturesOnboarding,omitempty"`
		DataSharingConsents                                                        *[]DataSharingConsent                                                 `json:"dataSharingConsents,omitempty"`
		DepOnboardingSettings                                                      *[]DepOnboardingSetting                                               `json:"depOnboardingSettings,omitempty"`
		DerivedCredentials                                                         *[]DeviceManagementDerivedCredentialSettings                          `json:"derivedCredentials,omitempty"`
		DetectedApps                                                               *[]DetectedApp                                                        `json:"detectedApps,omitempty"`
		DeviceCategories                                                           *[]DeviceCategory                                                     `json:"deviceCategories,omitempty"`
		DeviceCompliancePolicyDeviceStateSummary                                   *DeviceCompliancePolicyDeviceStateSummary                             `json:"deviceCompliancePolicyDeviceStateSummary,omitempty"`
		DeviceCompliancePolicySettingStateSummaries                                *[]DeviceCompliancePolicySettingStateSummary                          `json:"deviceCompliancePolicySettingStateSummaries,omitempty"`
		DeviceComplianceReportSummarizationDateTime                                *string                                                               `json:"deviceComplianceReportSummarizationDateTime,omitempty"`
		DeviceComplianceScripts                                                    *[]DeviceComplianceScript                                             `json:"deviceComplianceScripts,omitempty"`
		DeviceConfigurationConflictSummary                                         *[]DeviceConfigurationConflictSummary                                 `json:"deviceConfigurationConflictSummary,omitempty"`
		DeviceConfigurationDeviceStateSummaries                                    *DeviceConfigurationDeviceStateSummary                                `json:"deviceConfigurationDeviceStateSummaries,omitempty"`
		DeviceConfigurationProfiles                                                *[]DeviceConfigurationProfile                                         `json:"deviceConfigurationProfiles,omitempty"`
		DeviceConfigurationRestrictedAppsViolations                                *[]RestrictedAppsViolation                                            `json:"deviceConfigurationRestrictedAppsViolations,omitempty"`
		DeviceConfigurationUserStateSummaries                                      *DeviceConfigurationUserStateSummary                                  `json:"deviceConfigurationUserStateSummaries,omitempty"`
		DeviceConfigurationsAllManagedDeviceCertificateStates                      *[]ManagedAllDeviceCertificateState                                   `json:"deviceConfigurationsAllManagedDeviceCertificateStates,omitempty"`
		DeviceCustomAttributeShellScripts                                          *[]DeviceCustomAttributeShellScript                                   `json:"deviceCustomAttributeShellScripts,omitempty"`
		DeviceHealthScripts                                                        *[]DeviceHealthScript                                                 `json:"deviceHealthScripts,omitempty"`
		DeviceManagementPartners                                                   *[]DeviceManagementPartner                                            `json:"deviceManagementPartners,omitempty"`
		DeviceManagementScripts                                                    *[]DeviceManagementScript                                             `json:"deviceManagementScripts,omitempty"`
		DeviceProtectionOverview                                                   *DeviceProtectionOverview                                             `json:"deviceProtectionOverview,omitempty"`
		DeviceShellScripts                                                         *[]DeviceShellScript                                                  `json:"deviceShellScripts,omitempty"`
		DomainJoinConnectors                                                       *[]DeviceManagementDomainJoinConnector                                `json:"domainJoinConnectors,omitempty"`
		ElevationRequests                                                          *[]PrivilegeManagementElevationRequest                                `json:"elevationRequests,omitempty"`
		EmbeddedSIMActivationCodePools                                             *[]EmbeddedSIMActivationCodePool                                      `json:"embeddedSIMActivationCodePools,omitempty"`
		EndpointPrivilegeManagementProvisioningStatus                              *EndpointPrivilegeManagementProvisioningStatus                        `json:"endpointPrivilegeManagementProvisioningStatus,omitempty"`
		ExchangeConnectors                                                         *[]DeviceManagementExchangeConnector                                  `json:"exchangeConnectors,omitempty"`
		ExchangeOnPremisesPolicies                                                 *[]DeviceManagementExchangeOnPremisesPolicy                           `json:"exchangeOnPremisesPolicies,omitempty"`
		ExchangeOnPremisesPolicy                                                   *DeviceManagementExchangeOnPremisesPolicy                             `json:"exchangeOnPremisesPolicy,omitempty"`
		GroupPolicyCategories                                                      *[]GroupPolicyCategory                                                `json:"groupPolicyCategories,omitempty"`
		GroupPolicyConfigurations                                                  *[]GroupPolicyConfiguration                                           `json:"groupPolicyConfigurations,omitempty"`
		GroupPolicyDefinitions                                                     *[]GroupPolicyDefinition                                              `json:"groupPolicyDefinitions,omitempty"`
		GroupPolicyMigrationReports                                                *[]GroupPolicyMigrationReport                                         `json:"groupPolicyMigrationReports,omitempty"`
		GroupPolicyObjectFiles                                                     *[]GroupPolicyObjectFile                                              `json:"groupPolicyObjectFiles,omitempty"`
		GroupPolicyUploadedDefinitionFiles                                         *[]GroupPolicyUploadedDefinitionFile                                  `json:"groupPolicyUploadedDefinitionFiles,omitempty"`
		HardwareConfigurations                                                     *[]HardwareConfiguration                                              `json:"hardwareConfigurations,omitempty"`
		HardwarePasswordDetails                                                    *[]HardwarePasswordDetail                                             `json:"hardwarePasswordDetails,omitempty"`
		HardwarePasswordInfo                                                       *[]HardwarePasswordInfo                                               `json:"hardwarePasswordInfo,omitempty"`
		ImportedWindowsAutopilotDeviceIdentities                                   *[]ImportedWindowsAutopilotDeviceIdentity                             `json:"importedWindowsAutopilotDeviceIdentities,omitempty"`
		Intents                                                                    *[]DeviceManagementIntent                                             `json:"intents,omitempty"`
		IntuneAccountId                                                            *string                                                               `json:"intuneAccountId,omitempty"`
		IntuneBrand                                                                *IntuneBrand                                                          `json:"intuneBrand,omitempty"`
		IntuneBrandingProfiles                                                     *[]IntuneBrandingProfile                                              `json:"intuneBrandingProfiles,omitempty"`
		IosUpdateStatuses                                                          *[]IosUpdateDeviceStatus                                              `json:"iosUpdateStatuses,omitempty"`
		LastReportAggregationDateTime                                              *string                                                               `json:"lastReportAggregationDateTime,omitempty"`
		LegacyPcManangementEnabled                                                 *bool                                                                 `json:"legacyPcManangementEnabled,omitempty"`
		MacOSSoftwareUpdateAccountSummaries                                        *[]MacOSSoftwareUpdateAccountSummary                                  `json:"macOSSoftwareUpdateAccountSummaries,omitempty"`
		ManagedDeviceCleanupRules                                                  *[]ManagedDeviceCleanupRule                                           `json:"managedDeviceCleanupRules,omitempty"`
		ManagedDeviceCleanupSettings                                               *ManagedDeviceCleanupSettings                                         `json:"managedDeviceCleanupSettings,omitempty"`
		ManagedDeviceEncryptionStates                                              *[]ManagedDeviceEncryptionState                                       `json:"managedDeviceEncryptionStates,omitempty"`
		ManagedDeviceOverview                                                      *ManagedDeviceOverview                                                `json:"managedDeviceOverview,omitempty"`
		ManagedDeviceWindowsOSImages                                               *[]ManagedDeviceWindowsOperatingSystemImage                           `json:"managedDeviceWindowsOSImages,omitempty"`
		MaximumDepTokens                                                           *int64                                                                `json:"maximumDepTokens,omitempty"`
		MicrosoftTunnelConfigurations                                              *[]MicrosoftTunnelConfiguration                                       `json:"microsoftTunnelConfigurations,omitempty"`
		MicrosoftTunnelHealthThresholds                                            *[]MicrosoftTunnelHealthThreshold                                     `json:"microsoftTunnelHealthThresholds,omitempty"`
		MicrosoftTunnelServerLogCollectionResponses                                *[]MicrosoftTunnelServerLogCollectionResponse                         `json:"microsoftTunnelServerLogCollectionResponses,omitempty"`
		MicrosoftTunnelSites                                                       *[]MicrosoftTunnelSite                                                `json:"microsoftTunnelSites,omitempty"`
		MobileAppTroubleshootingEvents                                             *[]MobileAppTroubleshootingEvent                                      `json:"mobileAppTroubleshootingEvents,omitempty"`
		MobileThreatDefenseConnectors                                              *[]MobileThreatDefenseConnector                                       `json:"mobileThreatDefenseConnectors,omitempty"`
		Monitoring                                                                 *DeviceManagementMonitoring                                           `json:"monitoring,omitempty"`
		NdesConnectors                                                             *[]NdesConnector                                                      `json:"ndesConnectors,omitempty"`
		NotificationMessageTemplates                                               *[]NotificationMessageTemplate                                        `json:"notificationMessageTemplates,omitempty"`
		OperationApprovalPolicies                                                  *[]OperationApprovalPolicy                                            `json:"operationApprovalPolicies,omitempty"`
		OperationApprovalRequests                                                  *[]OperationApprovalRequest                                           `json:"operationApprovalRequests,omitempty"`
		PrivilegeManagementElevations                                              *[]PrivilegeManagementElevation                                       `json:"privilegeManagementElevations,omitempty"`
		RemoteActionAudits                                                         *[]RemoteActionAudit                                                  `json:"remoteActionAudits,omitempty"`
		RemoteAssistancePartners                                                   *[]RemoteAssistancePartner                                            `json:"remoteAssistancePartners,omitempty"`
		RemoteAssistanceSettings                                                   *RemoteAssistanceSettings                                             `json:"remoteAssistanceSettings,omitempty"`
		Reports                                                                    *DeviceManagementReports                                              `json:"reports,omitempty"`
		ResourceOperations                                                         *[]ResourceOperation                                                  `json:"resourceOperations,omitempty"`
		ReusablePolicySettings                                                     *[]DeviceManagementReusablePolicySetting                              `json:"reusablePolicySettings,omitempty"`
		RoleAssignments                                                            *[]DeviceAndAppManagementRoleAssignment                               `json:"roleAssignments,omitempty"`
		RoleScopeTags                                                              *[]RoleScopeTag                                                       `json:"roleScopeTags,omitempty"`
		ServiceNowConnections                                                      *[]ServiceNowConnection                                               `json:"serviceNowConnections,omitempty"`
		Settings                                                                   *DeviceManagementSettings                                             `json:"settings,omitempty"`
		SoftwareUpdateStatusSummary                                                *SoftwareUpdateStatusSummary                                          `json:"softwareUpdateStatusSummary,omitempty"`
		SubscriptionState                                                          *DeviceManagementSubscriptionState                                    `json:"subscriptionState,omitempty"`
		Subscriptions                                                              *DeviceManagementSubscriptions                                        `json:"subscriptions,omitempty"`
		TelecomExpenseManagementPartners                                           *[]TelecomExpenseManagementPartner                                    `json:"telecomExpenseManagementPartners,omitempty"`
		TemplateInsights                                                           *[]DeviceManagementTemplateInsightsDefinition                         `json:"templateInsights,omitempty"`
		TemplateSettings                                                           *[]DeviceManagementConfigurationSettingTemplate                       `json:"templateSettings,omitempty"`
		TenantAttachRBAC                                                           *TenantAttachRBAC                                                     `json:"tenantAttachRBAC,omitempty"`
		TermsAndConditions                                                         *[]TermsAndConditions                                                 `json:"termsAndConditions,omitempty"`
		UnlicensedAdminstratorsEnabled                                             *bool                                                                 `json:"unlicensedAdminstratorsEnabled,omitempty"`
		UserExperienceAnalyticsAnomaly                                             *[]UserExperienceAnalyticsAnomaly                                     `json:"userExperienceAnalyticsAnomaly,omitempty"`
		UserExperienceAnalyticsAnomalyCorrelationGroupOverview                     *[]UserExperienceAnalyticsAnomalyCorrelationGroupOverview             `json:"userExperienceAnalyticsAnomalyCorrelationGroupOverview,omitempty"`
		UserExperienceAnalyticsAnomalyDevice                                       *[]UserExperienceAnalyticsAnomalyDevice                               `json:"userExperienceAnalyticsAnomalyDevice,omitempty"`
		UserExperienceAnalyticsAnomalySeverityOverview                             *UserExperienceAnalyticsAnomalySeverityOverview                       `json:"userExperienceAnalyticsAnomalySeverityOverview,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformance                     *[]UserExperienceAnalyticsAppHealthApplicationPerformance             `json:"userExperienceAnalyticsAppHealthApplicationPerformance,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion         *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersion         `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails  *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails  `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion          *[]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion          `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion,omitempty"`
		UserExperienceAnalyticsAppHealthDeviceModelPerformance                     *[]UserExperienceAnalyticsAppHealthDeviceModelPerformance             `json:"userExperienceAnalyticsAppHealthDeviceModelPerformance,omitempty"`
		UserExperienceAnalyticsAppHealthDevicePerformance                          *[]UserExperienceAnalyticsAppHealthDevicePerformance                  `json:"userExperienceAnalyticsAppHealthDevicePerformance,omitempty"`
		UserExperienceAnalyticsAppHealthDevicePerformanceDetails                   *[]UserExperienceAnalyticsAppHealthDevicePerformanceDetails           `json:"userExperienceAnalyticsAppHealthDevicePerformanceDetails,omitempty"`
		UserExperienceAnalyticsAppHealthOSVersionPerformance                       *[]UserExperienceAnalyticsAppHealthOSVersionPerformance               `json:"userExperienceAnalyticsAppHealthOSVersionPerformance,omitempty"`
		UserExperienceAnalyticsAppHealthOverview                                   *UserExperienceAnalyticsCategory                                      `json:"userExperienceAnalyticsAppHealthOverview,omitempty"`
		UserExperienceAnalyticsBaselines                                           *[]UserExperienceAnalyticsBaseline                                    `json:"userExperienceAnalyticsBaselines,omitempty"`
		UserExperienceAnalyticsBatteryHealthAppImpact                              *[]UserExperienceAnalyticsBatteryHealthAppImpact                      `json:"userExperienceAnalyticsBatteryHealthAppImpact,omitempty"`
		UserExperienceAnalyticsBatteryHealthCapacityDetails                        *UserExperienceAnalyticsBatteryHealthCapacityDetails                  `json:"userExperienceAnalyticsBatteryHealthCapacityDetails,omitempty"`
		UserExperienceAnalyticsBatteryHealthDeviceAppImpact                        *[]UserExperienceAnalyticsBatteryHealthDeviceAppImpact                `json:"userExperienceAnalyticsBatteryHealthDeviceAppImpact,omitempty"`
		UserExperienceAnalyticsBatteryHealthDevicePerformance                      *[]UserExperienceAnalyticsBatteryHealthDevicePerformance              `json:"userExperienceAnalyticsBatteryHealthDevicePerformance,omitempty"`
		UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory                   *[]UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory           `json:"userExperienceAnalyticsBatteryHealthDeviceRuntimeHistory,omitempty"`
		UserExperienceAnalyticsBatteryHealthModelPerformance                       *[]UserExperienceAnalyticsBatteryHealthModelPerformance               `json:"userExperienceAnalyticsBatteryHealthModelPerformance,omitempty"`
		UserExperienceAnalyticsBatteryHealthOsPerformance                          *[]UserExperienceAnalyticsBatteryHealthOsPerformance                  `json:"userExperienceAnalyticsBatteryHealthOsPerformance,omitempty"`
		UserExperienceAnalyticsBatteryHealthRuntimeDetails                         *UserExperienceAnalyticsBatteryHealthRuntimeDetails                   `json:"userExperienceAnalyticsBatteryHealthRuntimeDetails,omitempty"`
		UserExperienceAnalyticsCategories                                          *[]UserExperienceAnalyticsCategory                                    `json:"userExperienceAnalyticsCategories,omitempty"`
		UserExperienceAnalyticsDeviceMetricHistory                                 *[]UserExperienceAnalyticsMetricHistory                               `json:"userExperienceAnalyticsDeviceMetricHistory,omitempty"`
		UserExperienceAnalyticsDevicePerformance                                   *[]UserExperienceAnalyticsDevicePerformance                           `json:"userExperienceAnalyticsDevicePerformance,omitempty"`
		UserExperienceAnalyticsDeviceScope                                         *UserExperienceAnalyticsDeviceScope                                   `json:"userExperienceAnalyticsDeviceScope,omitempty"`
		UserExperienceAnalyticsDeviceScopes                                        *[]UserExperienceAnalyticsDeviceScope                                 `json:"userExperienceAnalyticsDeviceScopes,omitempty"`
		UserExperienceAnalyticsDeviceScores                                        *[]UserExperienceAnalyticsDeviceScores                                `json:"userExperienceAnalyticsDeviceScores,omitempty"`
		UserExperienceAnalyticsDeviceStartupHistory                                *[]UserExperienceAnalyticsDeviceStartupHistory                        `json:"userExperienceAnalyticsDeviceStartupHistory,omitempty"`
		UserExperienceAnalyticsDeviceStartupProcessPerformance                     *[]UserExperienceAnalyticsDeviceStartupProcessPerformance             `json:"userExperienceAnalyticsDeviceStartupProcessPerformance,omitempty"`
		UserExperienceAnalyticsDeviceStartupProcesses                              *[]UserExperienceAnalyticsDeviceStartupProcess                        `json:"userExperienceAnalyticsDeviceStartupProcesses,omitempty"`
		UserExperienceAnalyticsDeviceTimelineEvent                                 *[]UserExperienceAnalyticsDeviceTimelineEvent                         `json:"userExperienceAnalyticsDeviceTimelineEvent,omitempty"`
		UserExperienceAnalyticsDevicesWithoutCloudIdentity                         *[]UserExperienceAnalyticsDeviceWithoutCloudIdentity                  `json:"userExperienceAnalyticsDevicesWithoutCloudIdentity,omitempty"`
		UserExperienceAnalyticsImpactingProcess                                    *[]UserExperienceAnalyticsImpactingProcess                            `json:"userExperienceAnalyticsImpactingProcess,omitempty"`
		UserExperienceAnalyticsMetricHistory                                       *[]UserExperienceAnalyticsMetricHistory                               `json:"userExperienceAnalyticsMetricHistory,omitempty"`
		UserExperienceAnalyticsModelScores                                         *[]UserExperienceAnalyticsModelScores                                 `json:"userExperienceAnalyticsModelScores,omitempty"`
		UserExperienceAnalyticsNotAutopilotReadyDevice                             *[]UserExperienceAnalyticsNotAutopilotReadyDevice                     `json:"userExperienceAnalyticsNotAutopilotReadyDevice,omitempty"`
		UserExperienceAnalyticsOverview                                            *UserExperienceAnalyticsOverview                                      `json:"userExperienceAnalyticsOverview,omitempty"`
		UserExperienceAnalyticsRemoteConnection                                    *[]UserExperienceAnalyticsRemoteConnection                            `json:"userExperienceAnalyticsRemoteConnection,omitempty"`
		UserExperienceAnalyticsResourcePerformance                                 *[]UserExperienceAnalyticsResourcePerformance                         `json:"userExperienceAnalyticsResourcePerformance,omitempty"`
		UserExperienceAnalyticsScoreHistory                                        *[]UserExperienceAnalyticsScoreHistory                                `json:"userExperienceAnalyticsScoreHistory,omitempty"`
		UserExperienceAnalyticsSettings                                            *UserExperienceAnalyticsSettings                                      `json:"userExperienceAnalyticsSettings,omitempty"`
		UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric             *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric       `json:"userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric,omitempty"`
		UserExperienceAnalyticsWorkFromAnywhereMetrics                             *[]UserExperienceAnalyticsWorkFromAnywhereMetric                      `json:"userExperienceAnalyticsWorkFromAnywhereMetrics,omitempty"`
		UserExperienceAnalyticsWorkFromAnywhereModelPerformance                    *[]UserExperienceAnalyticsWorkFromAnywhereModelPerformance            `json:"userExperienceAnalyticsWorkFromAnywhereModelPerformance,omitempty"`
		UserPfxCertificates                                                        *[]UserPFXCertificate                                                 `json:"userPfxCertificates,omitempty"`
		VirtualEndpoint                                                            *VirtualEndpoint                                                      `json:"virtualEndpoint,omitempty"`
		WindowsAutopilotDeviceIdentities                                           *[]WindowsAutopilotDeviceIdentity                                     `json:"windowsAutopilotDeviceIdentities,omitempty"`
		WindowsAutopilotSettings                                                   *WindowsAutopilotSettings                                             `json:"windowsAutopilotSettings,omitempty"`
		WindowsDriverUpdateProfiles                                                *[]WindowsDriverUpdateProfile                                         `json:"windowsDriverUpdateProfiles,omitempty"`
		WindowsFeatureUpdateProfiles                                               *[]WindowsFeatureUpdateProfile                                        `json:"windowsFeatureUpdateProfiles,omitempty"`
		WindowsInformationProtectionAppLearningSummaries                           *[]WindowsInformationProtectionAppLearningSummary                     `json:"windowsInformationProtectionAppLearningSummaries,omitempty"`
		WindowsInformationProtectionNetworkLearningSummaries                       *[]WindowsInformationProtectionNetworkLearningSummary                 `json:"windowsInformationProtectionNetworkLearningSummaries,omitempty"`
		WindowsMalwareInformation                                                  *[]WindowsMalwareInformation                                          `json:"windowsMalwareInformation,omitempty"`
		WindowsMalwareOverview                                                     *WindowsMalwareOverview                                               `json:"windowsMalwareOverview,omitempty"`
		WindowsQualityUpdatePolicies                                               *[]WindowsQualityUpdatePolicy                                         `json:"windowsQualityUpdatePolicies,omitempty"`
		WindowsQualityUpdateProfiles                                               *[]WindowsQualityUpdateProfile                                        `json:"windowsQualityUpdateProfiles,omitempty"`
		ZebraFotaArtifacts                                                         *[]ZebraFotaArtifact                                                  `json:"zebraFotaArtifacts,omitempty"`
		ZebraFotaConnector                                                         *ZebraFotaConnector                                                   `json:"zebraFotaConnector,omitempty"`
		ZebraFotaDeployments                                                       *[]ZebraFotaDeployment                                                `json:"zebraFotaDeployments,omitempty"`
		Id                                                                         *string                                                               `json:"id,omitempty"`
		ODataId                                                                    *string                                                               `json:"@odata.id,omitempty"`
		ODataType                                                                  *string                                                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountMoveCompletionDateTime = decoded.AccountMoveCompletionDateTime
	s.AdminConsent = decoded.AdminConsent
	s.AdvancedThreatProtectionOnboardingStateSummary = decoded.AdvancedThreatProtectionOnboardingStateSummary
	s.AndroidDeviceOwnerEnrollmentProfiles = decoded.AndroidDeviceOwnerEnrollmentProfiles
	s.AndroidForWorkAppConfigurationSchemas = decoded.AndroidForWorkAppConfigurationSchemas
	s.AndroidForWorkEnrollmentProfiles = decoded.AndroidForWorkEnrollmentProfiles
	s.AndroidForWorkSettings = decoded.AndroidForWorkSettings
	s.AndroidManagedStoreAccountEnterpriseSettings = decoded.AndroidManagedStoreAccountEnterpriseSettings
	s.AndroidManagedStoreAppConfigurationSchemas = decoded.AndroidManagedStoreAppConfigurationSchemas
	s.ApplePushNotificationCertificate = decoded.ApplePushNotificationCertificate
	s.AppleUserInitiatedEnrollmentProfiles = decoded.AppleUserInitiatedEnrollmentProfiles
	s.AuditEvents = decoded.AuditEvents
	s.AutopilotEvents = decoded.AutopilotEvents
	s.CartToClassAssociations = decoded.CartToClassAssociations
	s.CertificateConnectorDetails = decoded.CertificateConnectorDetails
	s.ChromeOSOnboardingSettings = decoded.ChromeOSOnboardingSettings
	s.CloudPCConnectivityIssues = decoded.CloudPCConnectivityIssues
	s.ComanagementEligibleDevices = decoded.ComanagementEligibleDevices
	s.ComplianceCategories = decoded.ComplianceCategories
	s.ComplianceManagementPartners = decoded.ComplianceManagementPartners
	s.CompliancePolicies = decoded.CompliancePolicies
	s.ConditionalAccessSettings = decoded.ConditionalAccessSettings
	s.ConfigManagerCollections = decoded.ConfigManagerCollections
	s.ConfigurationCategories = decoded.ConfigurationCategories
	s.ConfigurationPolicies = decoded.ConfigurationPolicies
	s.ConfigurationPolicyTemplates = decoded.ConfigurationPolicyTemplates
	s.ConnectorStatus = decoded.ConnectorStatus
	s.DataProcessorServiceForWindowsFeaturesOnboarding = decoded.DataProcessorServiceForWindowsFeaturesOnboarding
	s.DataSharingConsents = decoded.DataSharingConsents
	s.DepOnboardingSettings = decoded.DepOnboardingSettings
	s.DerivedCredentials = decoded.DerivedCredentials
	s.DetectedApps = decoded.DetectedApps
	s.DeviceCategories = decoded.DeviceCategories
	s.DeviceCompliancePolicyDeviceStateSummary = decoded.DeviceCompliancePolicyDeviceStateSummary
	s.DeviceCompliancePolicySettingStateSummaries = decoded.DeviceCompliancePolicySettingStateSummaries
	s.DeviceComplianceReportSummarizationDateTime = decoded.DeviceComplianceReportSummarizationDateTime
	s.DeviceComplianceScripts = decoded.DeviceComplianceScripts
	s.DeviceConfigurationConflictSummary = decoded.DeviceConfigurationConflictSummary
	s.DeviceConfigurationDeviceStateSummaries = decoded.DeviceConfigurationDeviceStateSummaries
	s.DeviceConfigurationProfiles = decoded.DeviceConfigurationProfiles
	s.DeviceConfigurationRestrictedAppsViolations = decoded.DeviceConfigurationRestrictedAppsViolations
	s.DeviceConfigurationUserStateSummaries = decoded.DeviceConfigurationUserStateSummaries
	s.DeviceConfigurationsAllManagedDeviceCertificateStates = decoded.DeviceConfigurationsAllManagedDeviceCertificateStates
	s.DeviceCustomAttributeShellScripts = decoded.DeviceCustomAttributeShellScripts
	s.DeviceHealthScripts = decoded.DeviceHealthScripts
	s.DeviceManagementPartners = decoded.DeviceManagementPartners
	s.DeviceManagementScripts = decoded.DeviceManagementScripts
	s.DeviceProtectionOverview = decoded.DeviceProtectionOverview
	s.DeviceShellScripts = decoded.DeviceShellScripts
	s.DomainJoinConnectors = decoded.DomainJoinConnectors
	s.ElevationRequests = decoded.ElevationRequests
	s.EmbeddedSIMActivationCodePools = decoded.EmbeddedSIMActivationCodePools
	s.EndpointPrivilegeManagementProvisioningStatus = decoded.EndpointPrivilegeManagementProvisioningStatus
	s.ExchangeConnectors = decoded.ExchangeConnectors
	s.ExchangeOnPremisesPolicies = decoded.ExchangeOnPremisesPolicies
	s.ExchangeOnPremisesPolicy = decoded.ExchangeOnPremisesPolicy
	s.GroupPolicyCategories = decoded.GroupPolicyCategories
	s.GroupPolicyConfigurations = decoded.GroupPolicyConfigurations
	s.GroupPolicyDefinitions = decoded.GroupPolicyDefinitions
	s.GroupPolicyMigrationReports = decoded.GroupPolicyMigrationReports
	s.GroupPolicyObjectFiles = decoded.GroupPolicyObjectFiles
	s.GroupPolicyUploadedDefinitionFiles = decoded.GroupPolicyUploadedDefinitionFiles
	s.HardwareConfigurations = decoded.HardwareConfigurations
	s.HardwarePasswordDetails = decoded.HardwarePasswordDetails
	s.HardwarePasswordInfo = decoded.HardwarePasswordInfo
	s.ImportedWindowsAutopilotDeviceIdentities = decoded.ImportedWindowsAutopilotDeviceIdentities
	s.Intents = decoded.Intents
	s.IntuneAccountId = decoded.IntuneAccountId
	s.IntuneBrand = decoded.IntuneBrand
	s.IntuneBrandingProfiles = decoded.IntuneBrandingProfiles
	s.IosUpdateStatuses = decoded.IosUpdateStatuses
	s.LastReportAggregationDateTime = decoded.LastReportAggregationDateTime
	s.LegacyPcManangementEnabled = decoded.LegacyPcManangementEnabled
	s.MacOSSoftwareUpdateAccountSummaries = decoded.MacOSSoftwareUpdateAccountSummaries
	s.ManagedDeviceCleanupRules = decoded.ManagedDeviceCleanupRules
	s.ManagedDeviceCleanupSettings = decoded.ManagedDeviceCleanupSettings
	s.ManagedDeviceEncryptionStates = decoded.ManagedDeviceEncryptionStates
	s.ManagedDeviceOverview = decoded.ManagedDeviceOverview
	s.ManagedDeviceWindowsOSImages = decoded.ManagedDeviceWindowsOSImages
	s.MaximumDepTokens = decoded.MaximumDepTokens
	s.MicrosoftTunnelConfigurations = decoded.MicrosoftTunnelConfigurations
	s.MicrosoftTunnelHealthThresholds = decoded.MicrosoftTunnelHealthThresholds
	s.MicrosoftTunnelServerLogCollectionResponses = decoded.MicrosoftTunnelServerLogCollectionResponses
	s.MicrosoftTunnelSites = decoded.MicrosoftTunnelSites
	s.MobileAppTroubleshootingEvents = decoded.MobileAppTroubleshootingEvents
	s.MobileThreatDefenseConnectors = decoded.MobileThreatDefenseConnectors
	s.Monitoring = decoded.Monitoring
	s.NdesConnectors = decoded.NdesConnectors
	s.NotificationMessageTemplates = decoded.NotificationMessageTemplates
	s.OperationApprovalPolicies = decoded.OperationApprovalPolicies
	s.OperationApprovalRequests = decoded.OperationApprovalRequests
	s.PrivilegeManagementElevations = decoded.PrivilegeManagementElevations
	s.RemoteActionAudits = decoded.RemoteActionAudits
	s.RemoteAssistancePartners = decoded.RemoteAssistancePartners
	s.RemoteAssistanceSettings = decoded.RemoteAssistanceSettings
	s.Reports = decoded.Reports
	s.ResourceOperations = decoded.ResourceOperations
	s.ReusablePolicySettings = decoded.ReusablePolicySettings
	s.RoleAssignments = decoded.RoleAssignments
	s.RoleScopeTags = decoded.RoleScopeTags
	s.ServiceNowConnections = decoded.ServiceNowConnections
	s.Settings = decoded.Settings
	s.SoftwareUpdateStatusSummary = decoded.SoftwareUpdateStatusSummary
	s.SubscriptionState = decoded.SubscriptionState
	s.Subscriptions = decoded.Subscriptions
	s.TelecomExpenseManagementPartners = decoded.TelecomExpenseManagementPartners
	s.TemplateInsights = decoded.TemplateInsights
	s.TemplateSettings = decoded.TemplateSettings
	s.TenantAttachRBAC = decoded.TenantAttachRBAC
	s.TermsAndConditions = decoded.TermsAndConditions
	s.UnlicensedAdminstratorsEnabled = decoded.UnlicensedAdminstratorsEnabled
	s.UserExperienceAnalyticsAnomaly = decoded.UserExperienceAnalyticsAnomaly
	s.UserExperienceAnalyticsAnomalyCorrelationGroupOverview = decoded.UserExperienceAnalyticsAnomalyCorrelationGroupOverview
	s.UserExperienceAnalyticsAnomalyDevice = decoded.UserExperienceAnalyticsAnomalyDevice
	s.UserExperienceAnalyticsAnomalySeverityOverview = decoded.UserExperienceAnalyticsAnomalySeverityOverview
	s.UserExperienceAnalyticsAppHealthApplicationPerformance = decoded.UserExperienceAnalyticsAppHealthApplicationPerformance
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersion
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
	s.UserExperienceAnalyticsAppHealthDeviceModelPerformance = decoded.UserExperienceAnalyticsAppHealthDeviceModelPerformance
	s.UserExperienceAnalyticsAppHealthDevicePerformance = decoded.UserExperienceAnalyticsAppHealthDevicePerformance
	s.UserExperienceAnalyticsAppHealthDevicePerformanceDetails = decoded.UserExperienceAnalyticsAppHealthDevicePerformanceDetails
	s.UserExperienceAnalyticsAppHealthOSVersionPerformance = decoded.UserExperienceAnalyticsAppHealthOSVersionPerformance
	s.UserExperienceAnalyticsAppHealthOverview = decoded.UserExperienceAnalyticsAppHealthOverview
	s.UserExperienceAnalyticsBaselines = decoded.UserExperienceAnalyticsBaselines
	s.UserExperienceAnalyticsBatteryHealthAppImpact = decoded.UserExperienceAnalyticsBatteryHealthAppImpact
	s.UserExperienceAnalyticsBatteryHealthCapacityDetails = decoded.UserExperienceAnalyticsBatteryHealthCapacityDetails
	s.UserExperienceAnalyticsBatteryHealthDeviceAppImpact = decoded.UserExperienceAnalyticsBatteryHealthDeviceAppImpact
	s.UserExperienceAnalyticsBatteryHealthDevicePerformance = decoded.UserExperienceAnalyticsBatteryHealthDevicePerformance
	s.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory = decoded.UserExperienceAnalyticsBatteryHealthDeviceRuntimeHistory
	s.UserExperienceAnalyticsBatteryHealthModelPerformance = decoded.UserExperienceAnalyticsBatteryHealthModelPerformance
	s.UserExperienceAnalyticsBatteryHealthOsPerformance = decoded.UserExperienceAnalyticsBatteryHealthOsPerformance
	s.UserExperienceAnalyticsBatteryHealthRuntimeDetails = decoded.UserExperienceAnalyticsBatteryHealthRuntimeDetails
	s.UserExperienceAnalyticsCategories = decoded.UserExperienceAnalyticsCategories
	s.UserExperienceAnalyticsDeviceMetricHistory = decoded.UserExperienceAnalyticsDeviceMetricHistory
	s.UserExperienceAnalyticsDevicePerformance = decoded.UserExperienceAnalyticsDevicePerformance
	s.UserExperienceAnalyticsDeviceScope = decoded.UserExperienceAnalyticsDeviceScope
	s.UserExperienceAnalyticsDeviceScopes = decoded.UserExperienceAnalyticsDeviceScopes
	s.UserExperienceAnalyticsDeviceScores = decoded.UserExperienceAnalyticsDeviceScores
	s.UserExperienceAnalyticsDeviceStartupHistory = decoded.UserExperienceAnalyticsDeviceStartupHistory
	s.UserExperienceAnalyticsDeviceStartupProcessPerformance = decoded.UserExperienceAnalyticsDeviceStartupProcessPerformance
	s.UserExperienceAnalyticsDeviceStartupProcesses = decoded.UserExperienceAnalyticsDeviceStartupProcesses
	s.UserExperienceAnalyticsDeviceTimelineEvent = decoded.UserExperienceAnalyticsDeviceTimelineEvent
	s.UserExperienceAnalyticsDevicesWithoutCloudIdentity = decoded.UserExperienceAnalyticsDevicesWithoutCloudIdentity
	s.UserExperienceAnalyticsImpactingProcess = decoded.UserExperienceAnalyticsImpactingProcess
	s.UserExperienceAnalyticsMetricHistory = decoded.UserExperienceAnalyticsMetricHistory
	s.UserExperienceAnalyticsModelScores = decoded.UserExperienceAnalyticsModelScores
	s.UserExperienceAnalyticsNotAutopilotReadyDevice = decoded.UserExperienceAnalyticsNotAutopilotReadyDevice
	s.UserExperienceAnalyticsOverview = decoded.UserExperienceAnalyticsOverview
	s.UserExperienceAnalyticsRemoteConnection = decoded.UserExperienceAnalyticsRemoteConnection
	s.UserExperienceAnalyticsResourcePerformance = decoded.UserExperienceAnalyticsResourcePerformance
	s.UserExperienceAnalyticsScoreHistory = decoded.UserExperienceAnalyticsScoreHistory
	s.UserExperienceAnalyticsSettings = decoded.UserExperienceAnalyticsSettings
	s.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric = decoded.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
	s.UserExperienceAnalyticsWorkFromAnywhereMetrics = decoded.UserExperienceAnalyticsWorkFromAnywhereMetrics
	s.UserExperienceAnalyticsWorkFromAnywhereModelPerformance = decoded.UserExperienceAnalyticsWorkFromAnywhereModelPerformance
	s.UserPfxCertificates = decoded.UserPfxCertificates
	s.VirtualEndpoint = decoded.VirtualEndpoint
	s.WindowsAutopilotDeviceIdentities = decoded.WindowsAutopilotDeviceIdentities
	s.WindowsAutopilotSettings = decoded.WindowsAutopilotSettings
	s.WindowsDriverUpdateProfiles = decoded.WindowsDriverUpdateProfiles
	s.WindowsFeatureUpdateProfiles = decoded.WindowsFeatureUpdateProfiles
	s.WindowsInformationProtectionAppLearningSummaries = decoded.WindowsInformationProtectionAppLearningSummaries
	s.WindowsInformationProtectionNetworkLearningSummaries = decoded.WindowsInformationProtectionNetworkLearningSummaries
	s.WindowsMalwareInformation = decoded.WindowsMalwareInformation
	s.WindowsMalwareOverview = decoded.WindowsMalwareOverview
	s.WindowsQualityUpdatePolicies = decoded.WindowsQualityUpdatePolicies
	s.WindowsQualityUpdateProfiles = decoded.WindowsQualityUpdateProfiles
	s.ZebraFotaArtifacts = decoded.ZebraFotaArtifacts
	s.ZebraFotaConnector = decoded.ZebraFotaConnector
	s.ZebraFotaDeployments = decoded.ZebraFotaDeployments
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagement into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignmentFilters"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AssignmentFilters into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceAndAppManagementAssignmentFilter, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceAndAppManagementAssignmentFilterImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AssignmentFilters' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AssignmentFilters = &output
	}

	if v, ok := temp["categories"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Categories into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingCategory, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingCategoryImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Categories' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Categories = &output
	}

	if v, ok := temp["comanagedDevices"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ComanagedDevices into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedDevice, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedDeviceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ComanagedDevices' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ComanagedDevices = &output
	}

	if v, ok := temp["complianceSettings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ComplianceSettings into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ComplianceSettings' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ComplianceSettings = &output
	}

	if v, ok := temp["configurationSettings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ConfigurationSettings into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ConfigurationSettings' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ConfigurationSettings = &output
	}

	if v, ok := temp["deviceCompliancePolicies"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeviceCompliancePolicies into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceCompliancePolicy, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceCompliancePolicyImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeviceCompliancePolicies' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceCompliancePolicies = &output
	}

	if v, ok := temp["deviceConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeviceConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceConfiguration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceConfigurationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeviceConfigurations' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceConfigurations = &output
	}

	if v, ok := temp["deviceEnrollmentConfigurations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling DeviceEnrollmentConfigurations into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceEnrollmentConfiguration, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceEnrollmentConfigurationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'DeviceEnrollmentConfigurations' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.DeviceEnrollmentConfigurations = &output
	}

	if v, ok := temp["groupPolicyDefinitionFiles"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling GroupPolicyDefinitionFiles into list []json.RawMessage: %+v", err)
		}

		output := make([]GroupPolicyDefinitionFile, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGroupPolicyDefinitionFileImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'GroupPolicyDefinitionFiles' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.GroupPolicyDefinitionFiles = &output
	}

	if v, ok := temp["importedDeviceIdentities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ImportedDeviceIdentities into list []json.RawMessage: %+v", err)
		}

		output := make([]ImportedDeviceIdentity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalImportedDeviceIdentityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ImportedDeviceIdentities' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ImportedDeviceIdentities = &output
	}

	if v, ok := temp["managedDevices"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ManagedDevices into list []json.RawMessage: %+v", err)
		}

		output := make([]ManagedDevice, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalManagedDeviceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ManagedDevices' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ManagedDevices = &output
	}

	if v, ok := temp["resourceAccessProfiles"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ResourceAccessProfiles into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementResourceAccessProfileBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementResourceAccessProfileBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ResourceAccessProfiles' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ResourceAccessProfiles = &output
	}

	if v, ok := temp["reusableSettings"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ReusableSettings into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementConfigurationSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementConfigurationSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ReusableSettings' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ReusableSettings = &output
	}

	if v, ok := temp["roleDefinitions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RoleDefinitions into list []json.RawMessage: %+v", err)
		}

		output := make([]RoleDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRoleDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RoleDefinitions' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RoleDefinitions = &output
	}

	if v, ok := temp["settingDefinitions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling SettingDefinitions into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementSettingDefinition, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementSettingDefinitionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'SettingDefinitions' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.SettingDefinitions = &output
	}

	if v, ok := temp["templates"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Templates into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementTemplate, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementTemplateImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Templates' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Templates = &output
	}

	if v, ok := temp["troubleshootingEvents"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TroubleshootingEvents into list []json.RawMessage: %+v", err)
		}

		output := make([]DeviceManagementTroubleshootingEvent, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDeviceManagementTroubleshootingEventImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TroubleshootingEvents' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TroubleshootingEvents = &output
	}

	if v, ok := temp["windowsAutopilotDeploymentProfiles"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling WindowsAutopilotDeploymentProfiles into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsAutopilotDeploymentProfile, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsAutopilotDeploymentProfileImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'WindowsAutopilotDeploymentProfiles' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.WindowsAutopilotDeploymentProfiles = &output
	}

	if v, ok := temp["windowsUpdateCatalogItems"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling WindowsUpdateCatalogItems into list []json.RawMessage: %+v", err)
		}

		output := make([]WindowsUpdateCatalogItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalWindowsUpdateCatalogItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'WindowsUpdateCatalogItems' for 'DeviceManagement': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.WindowsUpdateCatalogItems = &output
	}

	return nil
}
