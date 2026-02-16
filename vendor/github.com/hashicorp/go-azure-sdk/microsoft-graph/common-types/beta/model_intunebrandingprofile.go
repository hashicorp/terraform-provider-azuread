package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = IntuneBrandingProfile{}

type IntuneBrandingProfile struct {
	// The list of group assignments for the branding profile
	Assignments *[]IntuneBrandingProfileAssignment `json:"assignments,omitempty"`

	// Collection of blocked actions on the company portal as per platform and device ownership types.
	CompanyPortalBlockedActions *[]CompanyPortalBlockedAction `json:"companyPortalBlockedActions,omitempty"`

	// E-mail address of the person/organization responsible for IT support
	ContactITEmailAddress nullable.Type[string] `json:"contactITEmailAddress,omitempty"`

	// Name of the person/organization responsible for IT support
	ContactITName nullable.Type[string] `json:"contactITName,omitempty"`

	// Text comments regarding the person/organization responsible for IT support
	ContactITNotes nullable.Type[string] `json:"contactITNotes,omitempty"`

	// Phone number of the person/organization responsible for IT support
	ContactITPhoneNumber nullable.Type[string] `json:"contactITPhoneNumber,omitempty"`

	// Time when the BrandingProfile was created
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Text comments regarding what the admin has access to on the device
	CustomCanSeePrivacyMessage nullable.Type[string] `json:"customCanSeePrivacyMessage,omitempty"`

	// Text comments regarding what the admin doesn't have access to on the device
	CustomCantSeePrivacyMessage nullable.Type[string] `json:"customCantSeePrivacyMessage,omitempty"`

	// Text comments regarding what the admin doesn't have access to on the device
	CustomPrivacyMessage nullable.Type[string] `json:"customPrivacyMessage,omitempty"`

	// Applies to telemetry sent from all clients to the Intune service. When disabled, all proactive troubleshooting and
	// issue warnings within the client are turned off, and telemetry settings appear inactive or hidden to the device user.
	DisableClientTelemetry *bool `json:"disableClientTelemetry,omitempty"`

	// Boolean that indicates if Device Category Selection will be shown in Company Portal
	DisableDeviceCategorySelection *bool `json:"disableDeviceCategorySelection,omitempty"`

	// Company/organization name that is displayed to end users
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Options available for enrollment flow customization
	EnrollmentAvailability *EnrollmentAvailabilityOptions `json:"enrollmentAvailability,omitempty"`

	// Boolean that represents whether the profile is used as default or not
	IsDefaultProfile *bool `json:"isDefaultProfile,omitempty"`

	// Boolean that represents whether the adminsistrator has disabled the 'Factory Reset' action on corporate owned
	// devices.
	IsFactoryResetDisabled *bool `json:"isFactoryResetDisabled,omitempty"`

	// Boolean that represents whether the adminsistrator has disabled the 'Remove Device' action on corporate owned
	// devices.
	IsRemoveDeviceDisabled *bool `json:"isRemoveDeviceDisabled,omitempty"`

	// Customized image displayed in Company Portal apps landing page
	LandingPageCustomizedImage *MimeContent `json:"landingPageCustomizedImage,omitempty"`

	// Time when the BrandingProfile was last modified
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// Logo image displayed in Company Portal apps which have a light background behind the logo
	LightBackgroundLogo *MimeContent `json:"lightBackgroundLogo,omitempty"`

	// Display name of the company/organization’s IT helpdesk site
	OnlineSupportSiteName nullable.Type[string] `json:"onlineSupportSiteName,omitempty"`

	// URL to the company/organization’s IT helpdesk site
	OnlineSupportSiteUrl nullable.Type[string] `json:"onlineSupportSiteUrl,omitempty"`

	// URL to the company/organization’s privacy policy
	PrivacyUrl nullable.Type[string] `json:"privacyUrl,omitempty"`

	// Description of the profile
	ProfileDescription nullable.Type[string] `json:"profileDescription,omitempty"`

	// Name of the profile
	ProfileName nullable.Type[string] `json:"profileName,omitempty"`

	// List of scope tags assigned to the branding profile
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Boolean that indicates if a push notification is sent to users when their device ownership type changes from personal
	// to corporate
	SendDeviceOwnershipChangePushNotification *bool `json:"sendDeviceOwnershipChangePushNotification,omitempty"`

	// Boolean that indicates if AzureAD Enterprise Apps will be shown in Company Portal
	ShowAzureADEnterpriseApps *bool `json:"showAzureADEnterpriseApps,omitempty"`

	// Boolean that indicates if Configuration Manager Apps will be shown in Company Portal
	ShowConfigurationManagerApps *bool `json:"showConfigurationManagerApps,omitempty"`

	// Boolean that represents whether the administrator-supplied display name will be shown next to the logo image or not
	ShowDisplayNameNextToLogo *bool `json:"showDisplayNameNextToLogo,omitempty"`

	// Boolean that represents whether the administrator-supplied logo images are shown or not
	ShowLogo *bool `json:"showLogo,omitempty"`

	// Boolean that indicates if Office WebApps will be shown in Company Portal
	ShowOfficeWebApps *bool `json:"showOfficeWebApps,omitempty"`

	// Primary theme color used in the Company Portal applications and web portal
	ThemeColor *RgbColor `json:"themeColor,omitempty"`

	// Logo image displayed in Company Portal apps which have a theme color background behind the logo
	ThemeColorLogo *MimeContent `json:"themeColorLogo,omitempty"`

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

func (s IntuneBrandingProfile) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IntuneBrandingProfile{}

func (s IntuneBrandingProfile) MarshalJSON() ([]byte, error) {
	type wrapper IntuneBrandingProfile
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IntuneBrandingProfile: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IntuneBrandingProfile: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.intuneBrandingProfile"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IntuneBrandingProfile: %+v", err)
	}

	return encoded, nil
}
