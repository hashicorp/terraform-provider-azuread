package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettings struct {
	// The property to determine if Android device administrator enrollment is enabled for this account.
	AndroidDeviceAdministratorEnrollmentEnabled *bool `json:"androidDeviceAdministratorEnrollmentEnabled,omitempty"`

	// Provider type for Derived Credentials.
	DerivedCredentialProvider *DerivedCredentialProviderType `json:"derivedCredentialProvider,omitempty"`

	// The Derived Credential Provider self-service URI.
	DerivedCredentialUrl nullable.Type[string] `json:"derivedCredentialUrl,omitempty"`

	// The number of days a device is allowed to go without checking in to remain compliant.
	DeviceComplianceCheckinThresholdDays *int64 `json:"deviceComplianceCheckinThresholdDays,omitempty"`

	// When the device does not check in for specified number of days, the company data might be removed and the device will
	// not be under management. Valid values 30 to 270
	DeviceInactivityBeforeRetirementInDay *int64 `json:"deviceInactivityBeforeRetirementInDay,omitempty"`

	// Determines whether the autopilot diagnostic feature is enabled or not.
	EnableAutopilotDiagnostics *bool `json:"enableAutopilotDiagnostics,omitempty"`

	// Determines whether the device group membership report feature is enabled or not.
	EnableDeviceGroupMembershipReport *bool `json:"enableDeviceGroupMembershipReport,omitempty"`

	// Determines whether the enhanced troubleshooting UX is enabled or not.
	EnableEnhancedTroubleshootingExperience *bool `json:"enableEnhancedTroubleshootingExperience,omitempty"`

	// Determines whether the log collection feature should be available for use.
	EnableLogCollection *bool `json:"enableLogCollection,omitempty"`

	// Is feature enabled or not for enhanced jailbreak detection.
	EnhancedJailBreak *bool `json:"enhancedJailBreak,omitempty"`

	// The property to determine whether to ignore unsupported compliance settings on certian models of devices.
	IgnoreDevicesForUnsupportedSettingsEnabled *bool `json:"ignoreDevicesForUnsupportedSettingsEnabled,omitempty"`

	// Is feature enabled or not for scheduled action for rule.
	IsScheduledActionEnabled *bool `json:"isScheduledActionEnabled,omitempty"`

	// The property to determine if M365 App log collection is enabled for account. When TRUE it indicates that M365 app log
	// collection is enabled for account. When FALSE it indicates that M365 app log collection is disabled for account.
	// Default value is FALSE
	M365AppDiagnosticsEnabled *bool `json:"m365AppDiagnosticsEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Device should be noncompliant when there is no compliance policy targeted when this is true
	SecureByDefault *bool `json:"secureByDefault,omitempty"`
}
