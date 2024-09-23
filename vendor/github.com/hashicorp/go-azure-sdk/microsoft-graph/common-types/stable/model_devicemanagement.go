package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagement{}

type DeviceManagement struct {
	// Apple push notification certificate.
	ApplePushNotificationCertificate *ApplePushNotificationCertificate `json:"applePushNotificationCertificate,omitempty"`

	// The Audit Events
	AuditEvents *[]AuditEvent `json:"auditEvents,omitempty"`

	// The list of Compliance Management Partners configured by the tenant.
	ComplianceManagementPartners *[]ComplianceManagementPartner `json:"complianceManagementPartners,omitempty"`

	// The Exchange on premises conditional access settings. On premises conditional access will require devices to be both
	// enrolled and compliant for mail access
	ConditionalAccessSettings *OnPremisesConditionalAccessSettings `json:"conditionalAccessSettings,omitempty"`

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

	// The device configuration device state summary for this account.
	DeviceConfigurationDeviceStateSummaries *DeviceConfigurationDeviceStateSummary `json:"deviceConfigurationDeviceStateSummaries,omitempty"`

	// The device configurations.
	DeviceConfigurations *[]DeviceConfiguration `json:"deviceConfigurations,omitempty"`

	// The list of device enrollment configurations
	DeviceEnrollmentConfigurations *[]DeviceEnrollmentConfiguration `json:"deviceEnrollmentConfigurations,omitempty"`

	// The list of Device Management Partners configured by the tenant.
	DeviceManagementPartners *[]DeviceManagementPartner `json:"deviceManagementPartners,omitempty"`

	// Device protection overview.
	DeviceProtectionOverview *DeviceProtectionOverview `json:"deviceProtectionOverview,omitempty"`

	// The list of Exchange Connectors configured by the tenant.
	ExchangeConnectors *[]DeviceManagementExchangeConnector `json:"exchangeConnectors,omitempty"`

	// Collection of imported Windows autopilot devices.
	ImportedWindowsAutopilotDeviceIdentities *[]ImportedWindowsAutopilotDeviceIdentity `json:"importedWindowsAutopilotDeviceIdentities,omitempty"`

	// Intune Account Id for given tenant
	IntuneAccountId *string `json:"intuneAccountId,omitempty"`

	// intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as
	// the end user web portal.
	IntuneBrand *IntuneBrand `json:"intuneBrand,omitempty"`

	// The IOS software update installation statuses for this account.
	IosUpdateStatuses *[]IosUpdateDeviceStatus `json:"iosUpdateStatuses,omitempty"`

	// Device overview
	ManagedDeviceOverview *ManagedDeviceOverview `json:"managedDeviceOverview,omitempty"`

	// The list of managed devices.
	ManagedDevices *[]ManagedDevice `json:"managedDevices,omitempty"`

	// The collection property of MobileAppTroubleshootingEvent.
	MobileAppTroubleshootingEvents *[]MobileAppTroubleshootingEvent `json:"mobileAppTroubleshootingEvents,omitempty"`

	// The list of Mobile threat Defense connectors configured by the tenant.
	MobileThreatDefenseConnectors *[]MobileThreatDefenseConnector `json:"mobileThreatDefenseConnectors,omitempty"`

	// The Notification Message Templates.
	NotificationMessageTemplates *[]NotificationMessageTemplate `json:"notificationMessageTemplates,omitempty"`

	// The remote assist partners.
	RemoteAssistancePartners *[]RemoteAssistancePartner `json:"remoteAssistancePartners,omitempty"`

	// Reports singleton
	Reports *DeviceManagementReports `json:"reports,omitempty"`

	// The Resource Operations.
	ResourceOperations *[]ResourceOperation `json:"resourceOperations,omitempty"`

	// The Role Assignments.
	RoleAssignments *[]DeviceAndAppManagementRoleAssignment `json:"roleAssignments,omitempty"`

	// The Role Definitions.
	RoleDefinitions *[]RoleDefinition `json:"roleDefinitions,omitempty"`

	// Account level settings.
	Settings *DeviceManagementSettings `json:"settings,omitempty"`

	// The software update status summary.
	SoftwareUpdateStatusSummary *SoftwareUpdateStatusSummary `json:"softwareUpdateStatusSummary,omitempty"`

	// Tenant mobile device management subscription state.
	SubscriptionState *DeviceManagementSubscriptionState `json:"subscriptionState,omitempty"`

	// The telecom expense management partners.
	TelecomExpenseManagementPartners *[]TelecomExpenseManagementPartner `json:"telecomExpenseManagementPartners,omitempty"`

	// The terms and conditions associated with device management of the company.
	TermsAndConditions *[]TermsAndConditions `json:"termsAndConditions,omitempty"`

	// The list of troubleshooting events for the tenant.
	TroubleshootingEvents *[]DeviceManagementTroubleshootingEvent `json:"troubleshootingEvents,omitempty"`

	// User experience analytics appHealth Application Performance
	UserExperienceAnalyticsAppHealthApplicationPerformance *[]UserExperienceAnalyticsAppHealthApplicationPerformance `json:"userExperienceAnalyticsAppHealthApplicationPerformance,omitempty"`

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

	// User experience analytics categories
	UserExperienceAnalyticsCategories *[]UserExperienceAnalyticsCategory `json:"userExperienceAnalyticsCategories,omitempty"`

	// User experience analytics device performance
	UserExperienceAnalyticsDevicePerformance *[]UserExperienceAnalyticsDevicePerformance `json:"userExperienceAnalyticsDevicePerformance,omitempty"`

	// User experience analytics device scores
	UserExperienceAnalyticsDeviceScores *[]UserExperienceAnalyticsDeviceScores `json:"userExperienceAnalyticsDeviceScores,omitempty"`

	// User experience analytics device Startup History
	UserExperienceAnalyticsDeviceStartupHistory *[]UserExperienceAnalyticsDeviceStartupHistory `json:"userExperienceAnalyticsDeviceStartupHistory,omitempty"`

	// User experience analytics device Startup Process Performance
	UserExperienceAnalyticsDeviceStartupProcessPerformance *[]UserExperienceAnalyticsDeviceStartupProcessPerformance `json:"userExperienceAnalyticsDeviceStartupProcessPerformance,omitempty"`

	// User experience analytics device Startup Processes
	UserExperienceAnalyticsDeviceStartupProcesses *[]UserExperienceAnalyticsDeviceStartupProcess `json:"userExperienceAnalyticsDeviceStartupProcesses,omitempty"`

	// User experience analytics metric history
	UserExperienceAnalyticsMetricHistory *[]UserExperienceAnalyticsMetricHistory `json:"userExperienceAnalyticsMetricHistory,omitempty"`

	// User experience analytics model scores
	UserExperienceAnalyticsModelScores *[]UserExperienceAnalyticsModelScores `json:"userExperienceAnalyticsModelScores,omitempty"`

	// User experience analytics overview
	UserExperienceAnalyticsOverview *UserExperienceAnalyticsOverview `json:"userExperienceAnalyticsOverview,omitempty"`

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

	// Virtual endpoint
	VirtualEndpoint *VirtualEndpoint `json:"virtualEndpoint,omitempty"`

	// The Windows autopilot device identities contained collection.
	WindowsAutopilotDeviceIdentities *[]WindowsAutopilotDeviceIdentity `json:"windowsAutopilotDeviceIdentities,omitempty"`

	// The windows information protection app learning summaries.
	WindowsInformationProtectionAppLearningSummaries *[]WindowsInformationProtectionAppLearningSummary `json:"windowsInformationProtectionAppLearningSummaries,omitempty"`

	// The windows information protection network learning summaries.
	WindowsInformationProtectionNetworkLearningSummaries *[]WindowsInformationProtectionNetworkLearningSummary `json:"windowsInformationProtectionNetworkLearningSummaries,omitempty"`

	// The list of affected malware in the tenant.
	WindowsMalwareInformation *[]WindowsMalwareInformation `json:"windowsMalwareInformation,omitempty"`

	// Malware overview for windows devices.
	WindowsMalwareOverview *WindowsMalwareOverview `json:"windowsMalwareOverview,omitempty"`

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
		ApplePushNotificationCertificate                                           *ApplePushNotificationCertificate                                     `json:"applePushNotificationCertificate,omitempty"`
		AuditEvents                                                                *[]AuditEvent                                                         `json:"auditEvents,omitempty"`
		ComplianceManagementPartners                                               *[]ComplianceManagementPartner                                        `json:"complianceManagementPartners,omitempty"`
		ConditionalAccessSettings                                                  *OnPremisesConditionalAccessSettings                                  `json:"conditionalAccessSettings,omitempty"`
		DetectedApps                                                               *[]DetectedApp                                                        `json:"detectedApps,omitempty"`
		DeviceCategories                                                           *[]DeviceCategory                                                     `json:"deviceCategories,omitempty"`
		DeviceCompliancePolicyDeviceStateSummary                                   *DeviceCompliancePolicyDeviceStateSummary                             `json:"deviceCompliancePolicyDeviceStateSummary,omitempty"`
		DeviceCompliancePolicySettingStateSummaries                                *[]DeviceCompliancePolicySettingStateSummary                          `json:"deviceCompliancePolicySettingStateSummaries,omitempty"`
		DeviceConfigurationDeviceStateSummaries                                    *DeviceConfigurationDeviceStateSummary                                `json:"deviceConfigurationDeviceStateSummaries,omitempty"`
		DeviceManagementPartners                                                   *[]DeviceManagementPartner                                            `json:"deviceManagementPartners,omitempty"`
		DeviceProtectionOverview                                                   *DeviceProtectionOverview                                             `json:"deviceProtectionOverview,omitempty"`
		ExchangeConnectors                                                         *[]DeviceManagementExchangeConnector                                  `json:"exchangeConnectors,omitempty"`
		ImportedWindowsAutopilotDeviceIdentities                                   *[]ImportedWindowsAutopilotDeviceIdentity                             `json:"importedWindowsAutopilotDeviceIdentities,omitempty"`
		IntuneAccountId                                                            *string                                                               `json:"intuneAccountId,omitempty"`
		IntuneBrand                                                                *IntuneBrand                                                          `json:"intuneBrand,omitempty"`
		IosUpdateStatuses                                                          *[]IosUpdateDeviceStatus                                              `json:"iosUpdateStatuses,omitempty"`
		ManagedDeviceOverview                                                      *ManagedDeviceOverview                                                `json:"managedDeviceOverview,omitempty"`
		ManagedDevices                                                             *[]ManagedDevice                                                      `json:"managedDevices,omitempty"`
		MobileAppTroubleshootingEvents                                             *[]MobileAppTroubleshootingEvent                                      `json:"mobileAppTroubleshootingEvents,omitempty"`
		MobileThreatDefenseConnectors                                              *[]MobileThreatDefenseConnector                                       `json:"mobileThreatDefenseConnectors,omitempty"`
		NotificationMessageTemplates                                               *[]NotificationMessageTemplate                                        `json:"notificationMessageTemplates,omitempty"`
		RemoteAssistancePartners                                                   *[]RemoteAssistancePartner                                            `json:"remoteAssistancePartners,omitempty"`
		Reports                                                                    *DeviceManagementReports                                              `json:"reports,omitempty"`
		ResourceOperations                                                         *[]ResourceOperation                                                  `json:"resourceOperations,omitempty"`
		RoleAssignments                                                            *[]DeviceAndAppManagementRoleAssignment                               `json:"roleAssignments,omitempty"`
		Settings                                                                   *DeviceManagementSettings                                             `json:"settings,omitempty"`
		SoftwareUpdateStatusSummary                                                *SoftwareUpdateStatusSummary                                          `json:"softwareUpdateStatusSummary,omitempty"`
		SubscriptionState                                                          *DeviceManagementSubscriptionState                                    `json:"subscriptionState,omitempty"`
		TelecomExpenseManagementPartners                                           *[]TelecomExpenseManagementPartner                                    `json:"telecomExpenseManagementPartners,omitempty"`
		TermsAndConditions                                                         *[]TermsAndConditions                                                 `json:"termsAndConditions,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformance                     *[]UserExperienceAnalyticsAppHealthApplicationPerformance             `json:"userExperienceAnalyticsAppHealthApplicationPerformance,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails  *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDetails  `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId *[]UserExperienceAnalyticsAppHealthAppPerformanceByAppVersionDeviceId `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId,omitempty"`
		UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion          *[]UserExperienceAnalyticsAppHealthAppPerformanceByOSVersion          `json:"userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion,omitempty"`
		UserExperienceAnalyticsAppHealthDeviceModelPerformance                     *[]UserExperienceAnalyticsAppHealthDeviceModelPerformance             `json:"userExperienceAnalyticsAppHealthDeviceModelPerformance,omitempty"`
		UserExperienceAnalyticsAppHealthDevicePerformance                          *[]UserExperienceAnalyticsAppHealthDevicePerformance                  `json:"userExperienceAnalyticsAppHealthDevicePerformance,omitempty"`
		UserExperienceAnalyticsAppHealthDevicePerformanceDetails                   *[]UserExperienceAnalyticsAppHealthDevicePerformanceDetails           `json:"userExperienceAnalyticsAppHealthDevicePerformanceDetails,omitempty"`
		UserExperienceAnalyticsAppHealthOSVersionPerformance                       *[]UserExperienceAnalyticsAppHealthOSVersionPerformance               `json:"userExperienceAnalyticsAppHealthOSVersionPerformance,omitempty"`
		UserExperienceAnalyticsAppHealthOverview                                   *UserExperienceAnalyticsCategory                                      `json:"userExperienceAnalyticsAppHealthOverview,omitempty"`
		UserExperienceAnalyticsBaselines                                           *[]UserExperienceAnalyticsBaseline                                    `json:"userExperienceAnalyticsBaselines,omitempty"`
		UserExperienceAnalyticsCategories                                          *[]UserExperienceAnalyticsCategory                                    `json:"userExperienceAnalyticsCategories,omitempty"`
		UserExperienceAnalyticsDevicePerformance                                   *[]UserExperienceAnalyticsDevicePerformance                           `json:"userExperienceAnalyticsDevicePerformance,omitempty"`
		UserExperienceAnalyticsDeviceScores                                        *[]UserExperienceAnalyticsDeviceScores                                `json:"userExperienceAnalyticsDeviceScores,omitempty"`
		UserExperienceAnalyticsDeviceStartupHistory                                *[]UserExperienceAnalyticsDeviceStartupHistory                        `json:"userExperienceAnalyticsDeviceStartupHistory,omitempty"`
		UserExperienceAnalyticsDeviceStartupProcessPerformance                     *[]UserExperienceAnalyticsDeviceStartupProcessPerformance             `json:"userExperienceAnalyticsDeviceStartupProcessPerformance,omitempty"`
		UserExperienceAnalyticsDeviceStartupProcesses                              *[]UserExperienceAnalyticsDeviceStartupProcess                        `json:"userExperienceAnalyticsDeviceStartupProcesses,omitempty"`
		UserExperienceAnalyticsMetricHistory                                       *[]UserExperienceAnalyticsMetricHistory                               `json:"userExperienceAnalyticsMetricHistory,omitempty"`
		UserExperienceAnalyticsModelScores                                         *[]UserExperienceAnalyticsModelScores                                 `json:"userExperienceAnalyticsModelScores,omitempty"`
		UserExperienceAnalyticsOverview                                            *UserExperienceAnalyticsOverview                                      `json:"userExperienceAnalyticsOverview,omitempty"`
		UserExperienceAnalyticsScoreHistory                                        *[]UserExperienceAnalyticsScoreHistory                                `json:"userExperienceAnalyticsScoreHistory,omitempty"`
		UserExperienceAnalyticsSettings                                            *UserExperienceAnalyticsSettings                                      `json:"userExperienceAnalyticsSettings,omitempty"`
		UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric             *UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric       `json:"userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric,omitempty"`
		UserExperienceAnalyticsWorkFromAnywhereMetrics                             *[]UserExperienceAnalyticsWorkFromAnywhereMetric                      `json:"userExperienceAnalyticsWorkFromAnywhereMetrics,omitempty"`
		UserExperienceAnalyticsWorkFromAnywhereModelPerformance                    *[]UserExperienceAnalyticsWorkFromAnywhereModelPerformance            `json:"userExperienceAnalyticsWorkFromAnywhereModelPerformance,omitempty"`
		VirtualEndpoint                                                            *VirtualEndpoint                                                      `json:"virtualEndpoint,omitempty"`
		WindowsAutopilotDeviceIdentities                                           *[]WindowsAutopilotDeviceIdentity                                     `json:"windowsAutopilotDeviceIdentities,omitempty"`
		WindowsInformationProtectionAppLearningSummaries                           *[]WindowsInformationProtectionAppLearningSummary                     `json:"windowsInformationProtectionAppLearningSummaries,omitempty"`
		WindowsInformationProtectionNetworkLearningSummaries                       *[]WindowsInformationProtectionNetworkLearningSummary                 `json:"windowsInformationProtectionNetworkLearningSummaries,omitempty"`
		WindowsMalwareInformation                                                  *[]WindowsMalwareInformation                                          `json:"windowsMalwareInformation,omitempty"`
		WindowsMalwareOverview                                                     *WindowsMalwareOverview                                               `json:"windowsMalwareOverview,omitempty"`
		Id                                                                         *string                                                               `json:"id,omitempty"`
		ODataId                                                                    *string                                                               `json:"@odata.id,omitempty"`
		ODataType                                                                  *string                                                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplePushNotificationCertificate = decoded.ApplePushNotificationCertificate
	s.AuditEvents = decoded.AuditEvents
	s.ComplianceManagementPartners = decoded.ComplianceManagementPartners
	s.ConditionalAccessSettings = decoded.ConditionalAccessSettings
	s.DetectedApps = decoded.DetectedApps
	s.DeviceCategories = decoded.DeviceCategories
	s.DeviceCompliancePolicyDeviceStateSummary = decoded.DeviceCompliancePolicyDeviceStateSummary
	s.DeviceCompliancePolicySettingStateSummaries = decoded.DeviceCompliancePolicySettingStateSummaries
	s.DeviceConfigurationDeviceStateSummaries = decoded.DeviceConfigurationDeviceStateSummaries
	s.DeviceManagementPartners = decoded.DeviceManagementPartners
	s.DeviceProtectionOverview = decoded.DeviceProtectionOverview
	s.ExchangeConnectors = decoded.ExchangeConnectors
	s.ImportedWindowsAutopilotDeviceIdentities = decoded.ImportedWindowsAutopilotDeviceIdentities
	s.IntuneAccountId = decoded.IntuneAccountId
	s.IntuneBrand = decoded.IntuneBrand
	s.IosUpdateStatuses = decoded.IosUpdateStatuses
	s.ManagedDeviceOverview = decoded.ManagedDeviceOverview
	s.ManagedDevices = decoded.ManagedDevices
	s.MobileAppTroubleshootingEvents = decoded.MobileAppTroubleshootingEvents
	s.MobileThreatDefenseConnectors = decoded.MobileThreatDefenseConnectors
	s.NotificationMessageTemplates = decoded.NotificationMessageTemplates
	s.RemoteAssistancePartners = decoded.RemoteAssistancePartners
	s.Reports = decoded.Reports
	s.ResourceOperations = decoded.ResourceOperations
	s.RoleAssignments = decoded.RoleAssignments
	s.Settings = decoded.Settings
	s.SoftwareUpdateStatusSummary = decoded.SoftwareUpdateStatusSummary
	s.SubscriptionState = decoded.SubscriptionState
	s.TelecomExpenseManagementPartners = decoded.TelecomExpenseManagementPartners
	s.TermsAndConditions = decoded.TermsAndConditions
	s.UserExperienceAnalyticsAppHealthApplicationPerformance = decoded.UserExperienceAnalyticsAppHealthApplicationPerformance
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId
	s.UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion = decoded.UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion
	s.UserExperienceAnalyticsAppHealthDeviceModelPerformance = decoded.UserExperienceAnalyticsAppHealthDeviceModelPerformance
	s.UserExperienceAnalyticsAppHealthDevicePerformance = decoded.UserExperienceAnalyticsAppHealthDevicePerformance
	s.UserExperienceAnalyticsAppHealthDevicePerformanceDetails = decoded.UserExperienceAnalyticsAppHealthDevicePerformanceDetails
	s.UserExperienceAnalyticsAppHealthOSVersionPerformance = decoded.UserExperienceAnalyticsAppHealthOSVersionPerformance
	s.UserExperienceAnalyticsAppHealthOverview = decoded.UserExperienceAnalyticsAppHealthOverview
	s.UserExperienceAnalyticsBaselines = decoded.UserExperienceAnalyticsBaselines
	s.UserExperienceAnalyticsCategories = decoded.UserExperienceAnalyticsCategories
	s.UserExperienceAnalyticsDevicePerformance = decoded.UserExperienceAnalyticsDevicePerformance
	s.UserExperienceAnalyticsDeviceScores = decoded.UserExperienceAnalyticsDeviceScores
	s.UserExperienceAnalyticsDeviceStartupHistory = decoded.UserExperienceAnalyticsDeviceStartupHistory
	s.UserExperienceAnalyticsDeviceStartupProcessPerformance = decoded.UserExperienceAnalyticsDeviceStartupProcessPerformance
	s.UserExperienceAnalyticsDeviceStartupProcesses = decoded.UserExperienceAnalyticsDeviceStartupProcesses
	s.UserExperienceAnalyticsMetricHistory = decoded.UserExperienceAnalyticsMetricHistory
	s.UserExperienceAnalyticsModelScores = decoded.UserExperienceAnalyticsModelScores
	s.UserExperienceAnalyticsOverview = decoded.UserExperienceAnalyticsOverview
	s.UserExperienceAnalyticsScoreHistory = decoded.UserExperienceAnalyticsScoreHistory
	s.UserExperienceAnalyticsSettings = decoded.UserExperienceAnalyticsSettings
	s.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric = decoded.UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric
	s.UserExperienceAnalyticsWorkFromAnywhereMetrics = decoded.UserExperienceAnalyticsWorkFromAnywhereMetrics
	s.UserExperienceAnalyticsWorkFromAnywhereModelPerformance = decoded.UserExperienceAnalyticsWorkFromAnywhereModelPerformance
	s.VirtualEndpoint = decoded.VirtualEndpoint
	s.WindowsAutopilotDeviceIdentities = decoded.WindowsAutopilotDeviceIdentities
	s.WindowsInformationProtectionAppLearningSummaries = decoded.WindowsInformationProtectionAppLearningSummaries
	s.WindowsInformationProtectionNetworkLearningSummaries = decoded.WindowsInformationProtectionNetworkLearningSummaries
	s.WindowsMalwareInformation = decoded.WindowsMalwareInformation
	s.WindowsMalwareOverview = decoded.WindowsMalwareOverview
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DeviceManagement into map[string]json.RawMessage: %+v", err)
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

	return nil
}
