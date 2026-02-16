package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntuneBrand struct {
	// Collection of blocked actions on the company portal as per platform and device ownership types.
	CompanyPortalBlockedActions *[]CompanyPortalBlockedAction `json:"companyPortalBlockedActions,omitempty"`

	// Email address of the person/organization responsible for IT support.
	ContactITEmailAddress nullable.Type[string] `json:"contactITEmailAddress,omitempty"`

	// Name of the person/organization responsible for IT support.
	ContactITName nullable.Type[string] `json:"contactITName,omitempty"`

	// Text comments regarding the person/organization responsible for IT support.
	ContactITNotes nullable.Type[string] `json:"contactITNotes,omitempty"`

	// Phone number of the person/organization responsible for IT support.
	ContactITPhoneNumber nullable.Type[string] `json:"contactITPhoneNumber,omitempty"`

	// The custom privacy message used to explain what the organization can see and do on managed devices.
	CustomCanSeePrivacyMessage nullable.Type[string] `json:"customCanSeePrivacyMessage,omitempty"`

	// The custom privacy message used to explain what the organization can’t see or do on managed devices.
	CustomCantSeePrivacyMessage nullable.Type[string] `json:"customCantSeePrivacyMessage,omitempty"`

	// The custom privacy message used to explain what the organization can’t see or do on managed devices.
	CustomPrivacyMessage nullable.Type[string] `json:"customPrivacyMessage,omitempty"`

	// Logo image displayed in Company Portal apps which have a dark background behind the logo.
	DarkBackgroundLogo *MimeContent `json:"darkBackgroundLogo,omitempty"`

	// Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and
	// issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
	DisableClientTelemetry *bool `json:"disableClientTelemetry,omitempty"`

	// Boolean that indicates if Device Category Selection will be shown in Company Portal
	DisableDeviceCategorySelection *bool `json:"disableDeviceCategorySelection,omitempty"`

	// Company/organization name that is displayed to end users.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Options available for enrollment flow customization
	EnrollmentAvailability *EnrollmentAvailabilityOptions `json:"enrollmentAvailability,omitempty"`

	// Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned
	// devices.
	IsFactoryResetDisabled *bool `json:"isFactoryResetDisabled,omitempty"`

	// Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned
	// devices.
	IsRemoveDeviceDisabled *bool `json:"isRemoveDeviceDisabled,omitempty"`

	// Customized image displayed in Company Portal app landing page
	LandingPageCustomizedImage *MimeContent `json:"landingPageCustomizedImage,omitempty"`

	// Logo image displayed in Company Portal apps which have a light background behind the logo.
	LightBackgroundLogo *MimeContent `json:"lightBackgroundLogo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Display name of the company/organization’s IT helpdesk site.
	OnlineSupportSiteName nullable.Type[string] `json:"onlineSupportSiteName,omitempty"`

	// URL to the company/organization’s IT helpdesk site.
	OnlineSupportSiteUrl nullable.Type[string] `json:"onlineSupportSiteUrl,omitempty"`

	// URL to the company/organization’s privacy policy.
	PrivacyUrl nullable.Type[string] `json:"privacyUrl,omitempty"`

	// List of scope tags assigned to the default branding profile
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal
	// to corporate
	SendDeviceOwnershipChangePushNotification *bool `json:"sendDeviceOwnershipChangePushNotification,omitempty"`

	// Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
	ShowAzureADEnterpriseApps *bool `json:"showAzureADEnterpriseApps,omitempty"`

	// Boolean that indicates if ConfigurationManagerApps will be shown in Company Portal
	ShowConfigurationManagerApps *bool `json:"showConfigurationManagerApps,omitempty"`

	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
	ShowDisplayNameNextToLogo *bool `json:"showDisplayNameNextToLogo,omitempty"`

	// Boolean that represents whether the administrator-supplied logo images are shown or not shown.
	ShowLogo *bool `json:"showLogo,omitempty"`

	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image.
	ShowNameNextToLogo *bool `json:"showNameNextToLogo,omitempty"`

	// Boolean that indicates if Office WebApps will be shown in Company Portal
	ShowOfficeWebApps *bool `json:"showOfficeWebApps,omitempty"`

	// Primary theme color used in the Company Portal applications and web portal.
	ThemeColor *RgbColor `json:"themeColor,omitempty"`
}
