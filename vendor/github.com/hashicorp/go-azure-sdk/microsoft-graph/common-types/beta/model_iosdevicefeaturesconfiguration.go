package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AppleDeviceFeaturesConfigurationBase = IosDeviceFeaturesConfiguration{}

type IosDeviceFeaturesConfiguration struct {
	// Asset tag information for the device, displayed on the login window and lock screen.
	AssetTagTemplate nullable.Type[string] `json:"assetTagTemplate,omitempty"`

	// Gets or sets iOS Web Content Filter settings, supervised mode only
	ContentFilterSettings IosWebContentFilterBase `json:"contentFilterSettings"`

	// A list of app and folders to appear on the Home Screen Dock. This collection can contain a maximum of 500 elements.
	HomeScreenDockIcons *[]IosHomeScreenItem `json:"homeScreenDockIcons,omitempty"`

	// Gets or sets the number of rows to render when configuring iOS home screen layout settings. If this value is
	// configured, homeScreenGridWidth must be configured as well.
	HomeScreenGridHeight nullable.Type[int64] `json:"homeScreenGridHeight,omitempty"`

	// Gets or sets the number of columns to render when configuring iOS home screen layout settings. If this value is
	// configured, homeScreenGridHeight must be configured as well.
	HomeScreenGridWidth nullable.Type[int64] `json:"homeScreenGridWidth,omitempty"`

	// A list of pages on the Home Screen. This collection can contain a maximum of 500 elements.
	HomeScreenPages *[]IosHomeScreenPage `json:"homeScreenPages,omitempty"`

	// Identity Certificate for the renewal of Kerberos ticket used in single sign-on settings.
	IdentityCertificateForClientAuthentication *IosCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`

	// Gets or sets a single sign-on extension profile.
	IosSingleSignOnExtension *IosSingleSignOnExtension `json:"iosSingleSignOnExtension,omitempty"`

	// A footnote displayed on the login window and lock screen. Available in iOS 9.3.1 and later.
	LockScreenFootnote nullable.Type[string] `json:"lockScreenFootnote,omitempty"`

	// Notification settings for each bundle id. Applicable to devices in supervised mode only (iOS 9.3 and later). This
	// collection can contain a maximum of 500 elements.
	NotificationSettings *[]IosNotificationSettings `json:"notificationSettings,omitempty"`

	// Gets or sets a single sign-on extension profile. Deprecated: use IOSSingleSignOnExtension instead.
	SingleSignOnExtension SingleSignOnExtension `json:"singleSignOnExtension"`

	// PKINIT Certificate for the authentication with single sign-on extension settings.
	SingleSignOnExtensionPkinitCertificate *IosCertificateProfileBase `json:"singleSignOnExtensionPkinitCertificate,omitempty"`

	// The Kerberos login settings that enable apps on receiving devices to authenticate smoothly.
	SingleSignOnSettings *IosSingleSignOnSettings `json:"singleSignOnSettings,omitempty"`

	// An enum type for wallpaper display location specifier.
	WallpaperDisplayLocation *IosWallpaperDisplayLocation `json:"wallpaperDisplayLocation,omitempty"`

	// A wallpaper image must be in either PNG or JPEG format. It requires a supervised device with iOS 8 or later version.
	WallpaperImage *MimeContent `json:"wallpaperImage,omitempty"`

	// Fields inherited from AppleDeviceFeaturesConfigurationBase

	// An array of AirPrint printers that should always be shown. This collection can contain a maximum of 500 elements.
	AirPrintDestinations *[]AirPrintDestination `json:"airPrintDestinations,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The device mode applicability rule for this Policy.
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`

	// The OS edition applicability for this Policy.
	DeviceManagementApplicabilityRuleOsEdition *DeviceManagementApplicabilityRuleOsEdition `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`

	// The OS version applicability rule for this Policy.
	DeviceManagementApplicabilityRuleOsVersion *DeviceManagementApplicabilityRuleOsVersion `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// The list of group assignments for the device configuration profile.
	GroupAssignments *[]DeviceConfigurationGroupAssignment `json:"groupAssignments,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// List of Scope Tags for this Entity instance.
	RoleScopeTagIds *[]string `json:"roleScopeTagIds,omitempty"`

	// Indicates whether or not the underlying Device Configuration supports the assignment of scope tags. Assigning to the
	// ScopeTags property is not allowed when this value is false and entities will not be visible to scoped users. This
	// occurs for Legacy policies created in Silverlight and can be resolved by deleting and recreating the policy in the
	// Azure Portal. This property is read-only.
	SupportsScopeTags *bool `json:"supportsScopeTags,omitempty"`

	// Device Configuration users status overview
	UserStatusOverview *DeviceConfigurationUserOverview `json:"userStatusOverview,omitempty"`

	// Device configuration installation status by user.
	UserStatuses *[]DeviceConfigurationUserStatus `json:"userStatuses,omitempty"`

	// Version of the device configuration.
	Version *int64 `json:"version,omitempty"`

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

func (s IosDeviceFeaturesConfiguration) AppleDeviceFeaturesConfigurationBase() BaseAppleDeviceFeaturesConfigurationBaseImpl {
	return BaseAppleDeviceFeaturesConfigurationBaseImpl{
		AirPrintDestinations: s.AirPrintDestinations,
		Assignments:          s.Assignments,
		CreatedDateTime:      s.CreatedDateTime,
		Description:          s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s IosDeviceFeaturesConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:     s.Assignments,
		CreatedDateTime: s.CreatedDateTime,
		Description:     s.Description,
		DeviceManagementApplicabilityRuleDeviceMode: s.DeviceManagementApplicabilityRuleDeviceMode,
		DeviceManagementApplicabilityRuleOsEdition:  s.DeviceManagementApplicabilityRuleOsEdition,
		DeviceManagementApplicabilityRuleOsVersion:  s.DeviceManagementApplicabilityRuleOsVersion,
		DeviceSettingStateSummaries:                 s.DeviceSettingStateSummaries,
		DeviceStatusOverview:                        s.DeviceStatusOverview,
		DeviceStatuses:                              s.DeviceStatuses,
		DisplayName:                                 s.DisplayName,
		GroupAssignments:                            s.GroupAssignments,
		LastModifiedDateTime:                        s.LastModifiedDateTime,
		RoleScopeTagIds:                             s.RoleScopeTagIds,
		SupportsScopeTags:                           s.SupportsScopeTags,
		UserStatusOverview:                          s.UserStatusOverview,
		UserStatuses:                                s.UserStatuses,
		Version:                                     s.Version,
		Id:                                          s.Id,
		ODataId:                                     s.ODataId,
		ODataType:                                   s.ODataType,
	}
}

func (s IosDeviceFeaturesConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosDeviceFeaturesConfiguration{}

func (s IosDeviceFeaturesConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IosDeviceFeaturesConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosDeviceFeaturesConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosDeviceFeaturesConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosDeviceFeaturesConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosDeviceFeaturesConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &IosDeviceFeaturesConfiguration{}

func (s *IosDeviceFeaturesConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AssetTagTemplate                            nullable.Type[string]                        `json:"assetTagTemplate,omitempty"`
		HomeScreenGridHeight                        nullable.Type[int64]                         `json:"homeScreenGridHeight,omitempty"`
		HomeScreenGridWidth                         nullable.Type[int64]                         `json:"homeScreenGridWidth,omitempty"`
		HomeScreenPages                             *[]IosHomeScreenPage                         `json:"homeScreenPages,omitempty"`
		LockScreenFootnote                          nullable.Type[string]                        `json:"lockScreenFootnote,omitempty"`
		NotificationSettings                        *[]IosNotificationSettings                   `json:"notificationSettings,omitempty"`
		SingleSignOnSettings                        *IosSingleSignOnSettings                     `json:"singleSignOnSettings,omitempty"`
		WallpaperDisplayLocation                    *IosWallpaperDisplayLocation                 `json:"wallpaperDisplayLocation,omitempty"`
		WallpaperImage                              *MimeContent                                 `json:"wallpaperImage,omitempty"`
		AirPrintDestinations                        *[]AirPrintDestination                       `json:"airPrintDestinations,omitempty"`
		Assignments                                 *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                             *string                                      `json:"createdDateTime,omitempty"`
		Description                                 nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                        *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                              *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                 *string                                      `json:"displayName,omitempty"`
		GroupAssignments                            *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                        *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                             *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                           *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                          *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                     *int64                                       `json:"version,omitempty"`
		Id                                          *string                                      `json:"id,omitempty"`
		ODataId                                     *string                                      `json:"@odata.id,omitempty"`
		ODataType                                   *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AssetTagTemplate = decoded.AssetTagTemplate
	s.HomeScreenGridHeight = decoded.HomeScreenGridHeight
	s.HomeScreenGridWidth = decoded.HomeScreenGridWidth
	s.HomeScreenPages = decoded.HomeScreenPages
	s.LockScreenFootnote = decoded.LockScreenFootnote
	s.NotificationSettings = decoded.NotificationSettings
	s.SingleSignOnSettings = decoded.SingleSignOnSettings
	s.WallpaperDisplayLocation = decoded.WallpaperDisplayLocation
	s.WallpaperImage = decoded.WallpaperImage
	s.AirPrintDestinations = decoded.AirPrintDestinations
	s.Assignments = decoded.Assignments
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Description = decoded.Description
	s.DeviceManagementApplicabilityRuleDeviceMode = decoded.DeviceManagementApplicabilityRuleDeviceMode
	s.DeviceManagementApplicabilityRuleOsEdition = decoded.DeviceManagementApplicabilityRuleOsEdition
	s.DeviceManagementApplicabilityRuleOsVersion = decoded.DeviceManagementApplicabilityRuleOsVersion
	s.DeviceSettingStateSummaries = decoded.DeviceSettingStateSummaries
	s.DeviceStatusOverview = decoded.DeviceStatusOverview
	s.DeviceStatuses = decoded.DeviceStatuses
	s.DisplayName = decoded.DisplayName
	s.GroupAssignments = decoded.GroupAssignments
	s.Id = decoded.Id
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RoleScopeTagIds = decoded.RoleScopeTagIds
	s.SupportsScopeTags = decoded.SupportsScopeTags
	s.UserStatusOverview = decoded.UserStatusOverview
	s.UserStatuses = decoded.UserStatuses
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IosDeviceFeaturesConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["contentFilterSettings"]; ok {
		impl, err := UnmarshalIosWebContentFilterBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ContentFilterSettings' for 'IosDeviceFeaturesConfiguration': %+v", err)
		}
		s.ContentFilterSettings = impl
	}

	if v, ok := temp["homeScreenDockIcons"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling HomeScreenDockIcons into list []json.RawMessage: %+v", err)
		}

		output := make([]IosHomeScreenItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIosHomeScreenItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'HomeScreenDockIcons' for 'IosDeviceFeaturesConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.HomeScreenDockIcons = &output
	}

	if v, ok := temp["identityCertificateForClientAuthentication"]; ok {
		impl, err := UnmarshalIosCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IdentityCertificateForClientAuthentication' for 'IosDeviceFeaturesConfiguration': %+v", err)
		}
		s.IdentityCertificateForClientAuthentication = &impl
	}

	if v, ok := temp["iosSingleSignOnExtension"]; ok {
		impl, err := UnmarshalIosSingleSignOnExtensionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'IosSingleSignOnExtension' for 'IosDeviceFeaturesConfiguration': %+v", err)
		}
		s.IosSingleSignOnExtension = &impl
	}

	if v, ok := temp["singleSignOnExtension"]; ok {
		impl, err := UnmarshalSingleSignOnExtensionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SingleSignOnExtension' for 'IosDeviceFeaturesConfiguration': %+v", err)
		}
		s.SingleSignOnExtension = impl
	}

	if v, ok := temp["singleSignOnExtensionPkinitCertificate"]; ok {
		impl, err := UnmarshalIosCertificateProfileBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SingleSignOnExtensionPkinitCertificate' for 'IosDeviceFeaturesConfiguration': %+v", err)
		}
		s.SingleSignOnExtensionPkinitCertificate = &impl
	}

	return nil
}
