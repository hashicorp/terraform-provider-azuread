package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = AndroidGeneralDeviceConfiguration{}

type AndroidGeneralDeviceConfiguration struct {
	// Indicates whether or not to block clipboard sharing to copy and paste between applications.
	AppsBlockClipboardSharing *bool `json:"appsBlockClipboardSharing,omitempty"`

	// Indicates whether or not to block copy and paste within applications.
	AppsBlockCopyPaste *bool `json:"appsBlockCopyPaste,omitempty"`

	// Indicates whether or not to block the YouTube app.
	AppsBlockYouTube *bool `json:"appsBlockYouTube,omitempty"`

	// List of apps to be hidden on the KNOX device. This collection can contain a maximum of 500 elements.
	AppsHideList *[]AppListItem `json:"appsHideList,omitempty"`

	// List of apps which can be installed on the KNOX device. This collection can contain a maximum of 500 elements.
	AppsInstallAllowList *[]AppListItem `json:"appsInstallAllowList,omitempty"`

	// List of apps which are blocked from being launched on the KNOX device. This collection can contain a maximum of 500
	// elements.
	AppsLaunchBlockList *[]AppListItem `json:"appsLaunchBlockList,omitempty"`

	// Indicates whether or not to block Bluetooth.
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`

	// Indicates whether or not to block the use of the camera.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`

	// Indicates whether or not to block data roaming.
	CellularBlockDataRoaming *bool `json:"cellularBlockDataRoaming,omitempty"`

	// Indicates whether or not to block SMS/MMS messaging.
	CellularBlockMessaging *bool `json:"cellularBlockMessaging,omitempty"`

	// Indicates whether or not to block voice roaming.
	CellularBlockVoiceRoaming *bool `json:"cellularBlockVoiceRoaming,omitempty"`

	// Indicates whether or not to block syncing Wi-Fi tethering.
	CellularBlockWiFiTethering *bool `json:"cellularBlockWiFiTethering,omitempty"`

	// Possible values of the compliance app list.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`

	// List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection
	// can contain a maximum of 10000 elements.
	CompliantAppsList *[]AppListItem `json:"compliantAppsList,omitempty"`

	// Indicates whether or not to block changing date and time while in KNOX Mode.
	DateAndTimeBlockChanges *bool `json:"dateAndTimeBlockChanges,omitempty"`

	// Indicates whether or not to allow device sharing mode.
	DeviceSharingAllowed *bool `json:"deviceSharingAllowed,omitempty"`

	// Indicates whether or not to block diagnostic data submission.
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`

	// Indicates whether or not to block user performing a factory reset.
	FactoryResetBlocked *bool `json:"factoryResetBlocked,omitempty"`

	// Indicates whether or not to block Google account auto sync.
	GoogleAccountBlockAutoSync *bool `json:"googleAccountBlockAutoSync,omitempty"`

	// Indicates whether or not to block the Google Play store.
	GooglePlayStoreBlocked *bool `json:"googlePlayStoreBlocked,omitempty"`

	// A list of apps that will be allowed to run when the device is in Kiosk Mode. This collection can contain a maximum of
	// 500 elements.
	KioskModeApps *[]AppListItem `json:"kioskModeApps,omitempty"`

	// Indicates whether or not to block the screen sleep button while in Kiosk Mode.
	KioskModeBlockSleepButton *bool `json:"kioskModeBlockSleepButton,omitempty"`

	// Indicates whether or not to block the volume buttons while in Kiosk Mode.
	KioskModeBlockVolumeButtons *bool `json:"kioskModeBlockVolumeButtons,omitempty"`

	// Indicates whether or not to block location services.
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`

	// Indicates whether or not to block Near-Field Communication.
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`

	// Indicates whether or not to block fingerprint unlock.
	PasswordBlockFingerprintUnlock *bool `json:"passwordBlockFingerprintUnlock,omitempty"`

	// Indicates whether or not to block Smart Lock and other trust agents.
	PasswordBlockTrustAgents *bool `json:"passwordBlockTrustAgents,omitempty"`

	// Number of days before the password expires. Valid values 1 to 365
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Minimum length of passwords. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Number of previous passwords to block. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Indicates whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Android required password type.
	PasswordRequiredType *AndroidRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Number of sign in failures allowed before factory reset. Valid values 1 to 16
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// Indicates whether or not to block powering off the device.
	PowerOffBlocked *bool `json:"powerOffBlocked,omitempty"`

	// The password complexity types that can be set on Android. One of: NONE, LOW, MEDIUM, HIGH. This is an API targeted to
	// Android 11+.
	RequiredPasswordComplexity *AndroidRequiredPasswordComplexity `json:"requiredPasswordComplexity,omitempty"`

	// Indicates whether or not to block screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Require the Android Verify apps feature is turned on.
	SecurityRequireVerifyApps *bool `json:"securityRequireVerifyApps,omitempty"`

	// Indicates whether or not to block Google Backup.
	StorageBlockGoogleBackup *bool `json:"storageBlockGoogleBackup,omitempty"`

	// Indicates whether or not to block removable storage usage.
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`

	// Indicates whether or not to require device encryption.
	StorageRequireDeviceEncryption *bool `json:"storageRequireDeviceEncryption,omitempty"`

	// Indicates whether or not to require removable storage encryption.
	StorageRequireRemovableStorageEncryption *bool `json:"storageRequireRemovableStorageEncryption,omitempty"`

	// Indicates whether or not to block the use of the Voice Assistant.
	VoiceAssistantBlocked *bool `json:"voiceAssistantBlocked,omitempty"`

	// Indicates whether or not to block voice dialing.
	VoiceDialingBlocked *bool `json:"voiceDialingBlocked,omitempty"`

	// Indicates whether or not to block the web browser's auto fill feature.
	WebBrowserBlockAutofill *bool `json:"webBrowserBlockAutofill,omitempty"`

	// Indicates whether or not to block JavaScript within the web browser.
	WebBrowserBlockJavaScript *bool `json:"webBrowserBlockJavaScript,omitempty"`

	// Indicates whether or not to block popups within the web browser.
	WebBrowserBlockPopups *bool `json:"webBrowserBlockPopups,omitempty"`

	// Indicates whether or not to block the web browser.
	WebBrowserBlocked *bool `json:"webBrowserBlocked,omitempty"`

	// Web Browser Cookie Settings.
	WebBrowserCookieSettings *WebBrowserCookieSettings `json:"webBrowserCookieSettings,omitempty"`

	// Indicates whether or not to block syncing Wi-Fi.
	WiFiBlocked *bool `json:"wiFiBlocked,omitempty"`

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

func (s AndroidGeneralDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s AndroidGeneralDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidGeneralDeviceConfiguration{}

func (s AndroidGeneralDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidGeneralDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidGeneralDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidGeneralDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidGeneralDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidGeneralDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidGeneralDeviceConfiguration{}

func (s *AndroidGeneralDeviceConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AppsBlockClipboardSharing                      *bool                                        `json:"appsBlockClipboardSharing,omitempty"`
		AppsBlockCopyPaste                             *bool                                        `json:"appsBlockCopyPaste,omitempty"`
		AppsBlockYouTube                               *bool                                        `json:"appsBlockYouTube,omitempty"`
		BluetoothBlocked                               *bool                                        `json:"bluetoothBlocked,omitempty"`
		CameraBlocked                                  *bool                                        `json:"cameraBlocked,omitempty"`
		CellularBlockDataRoaming                       *bool                                        `json:"cellularBlockDataRoaming,omitempty"`
		CellularBlockMessaging                         *bool                                        `json:"cellularBlockMessaging,omitempty"`
		CellularBlockVoiceRoaming                      *bool                                        `json:"cellularBlockVoiceRoaming,omitempty"`
		CellularBlockWiFiTethering                     *bool                                        `json:"cellularBlockWiFiTethering,omitempty"`
		CompliantAppListType                           *AppListType                                 `json:"compliantAppListType,omitempty"`
		DateAndTimeBlockChanges                        *bool                                        `json:"dateAndTimeBlockChanges,omitempty"`
		DeviceSharingAllowed                           *bool                                        `json:"deviceSharingAllowed,omitempty"`
		DiagnosticDataBlockSubmission                  *bool                                        `json:"diagnosticDataBlockSubmission,omitempty"`
		FactoryResetBlocked                            *bool                                        `json:"factoryResetBlocked,omitempty"`
		GoogleAccountBlockAutoSync                     *bool                                        `json:"googleAccountBlockAutoSync,omitempty"`
		GooglePlayStoreBlocked                         *bool                                        `json:"googlePlayStoreBlocked,omitempty"`
		KioskModeBlockSleepButton                      *bool                                        `json:"kioskModeBlockSleepButton,omitempty"`
		KioskModeBlockVolumeButtons                    *bool                                        `json:"kioskModeBlockVolumeButtons,omitempty"`
		LocationServicesBlocked                        *bool                                        `json:"locationServicesBlocked,omitempty"`
		NfcBlocked                                     *bool                                        `json:"nfcBlocked,omitempty"`
		PasswordBlockFingerprintUnlock                 *bool                                        `json:"passwordBlockFingerprintUnlock,omitempty"`
		PasswordBlockTrustAgents                       *bool                                        `json:"passwordBlockTrustAgents,omitempty"`
		PasswordExpirationDays                         nullable.Type[int64]                         `json:"passwordExpirationDays,omitempty"`
		PasswordMinimumLength                          nullable.Type[int64]                         `json:"passwordMinimumLength,omitempty"`
		PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64]                         `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasswordPreviousPasswordBlockCount             nullable.Type[int64]                         `json:"passwordPreviousPasswordBlockCount,omitempty"`
		PasswordRequired                               *bool                                        `json:"passwordRequired,omitempty"`
		PasswordRequiredType                           *AndroidRequiredPasswordType                 `json:"passwordRequiredType,omitempty"`
		PasswordSignInFailureCountBeforeFactoryReset   nullable.Type[int64]                         `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
		PowerOffBlocked                                *bool                                        `json:"powerOffBlocked,omitempty"`
		RequiredPasswordComplexity                     *AndroidRequiredPasswordComplexity           `json:"requiredPasswordComplexity,omitempty"`
		ScreenCaptureBlocked                           *bool                                        `json:"screenCaptureBlocked,omitempty"`
		SecurityRequireVerifyApps                      *bool                                        `json:"securityRequireVerifyApps,omitempty"`
		StorageBlockGoogleBackup                       *bool                                        `json:"storageBlockGoogleBackup,omitempty"`
		StorageBlockRemovableStorage                   *bool                                        `json:"storageBlockRemovableStorage,omitempty"`
		StorageRequireDeviceEncryption                 *bool                                        `json:"storageRequireDeviceEncryption,omitempty"`
		StorageRequireRemovableStorageEncryption       *bool                                        `json:"storageRequireRemovableStorageEncryption,omitempty"`
		VoiceAssistantBlocked                          *bool                                        `json:"voiceAssistantBlocked,omitempty"`
		VoiceDialingBlocked                            *bool                                        `json:"voiceDialingBlocked,omitempty"`
		WebBrowserBlockAutofill                        *bool                                        `json:"webBrowserBlockAutofill,omitempty"`
		WebBrowserBlockJavaScript                      *bool                                        `json:"webBrowserBlockJavaScript,omitempty"`
		WebBrowserBlockPopups                          *bool                                        `json:"webBrowserBlockPopups,omitempty"`
		WebBrowserBlocked                              *bool                                        `json:"webBrowserBlocked,omitempty"`
		WebBrowserCookieSettings                       *WebBrowserCookieSettings                    `json:"webBrowserCookieSettings,omitempty"`
		WiFiBlocked                                    *bool                                        `json:"wiFiBlocked,omitempty"`
		Assignments                                    *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                                *string                                      `json:"createdDateTime,omitempty"`
		Description                                    nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode    *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition     *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion     *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                           *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                 *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                    *string                                      `json:"displayName,omitempty"`
		GroupAssignments                               *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                           *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                              *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                             *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                   *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                        *int64                                       `json:"version,omitempty"`
		Id                                             *string                                      `json:"id,omitempty"`
		ODataId                                        *string                                      `json:"@odata.id,omitempty"`
		ODataType                                      *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AppsBlockClipboardSharing = decoded.AppsBlockClipboardSharing
	s.AppsBlockCopyPaste = decoded.AppsBlockCopyPaste
	s.AppsBlockYouTube = decoded.AppsBlockYouTube
	s.BluetoothBlocked = decoded.BluetoothBlocked
	s.CameraBlocked = decoded.CameraBlocked
	s.CellularBlockDataRoaming = decoded.CellularBlockDataRoaming
	s.CellularBlockMessaging = decoded.CellularBlockMessaging
	s.CellularBlockVoiceRoaming = decoded.CellularBlockVoiceRoaming
	s.CellularBlockWiFiTethering = decoded.CellularBlockWiFiTethering
	s.CompliantAppListType = decoded.CompliantAppListType
	s.DateAndTimeBlockChanges = decoded.DateAndTimeBlockChanges
	s.DeviceSharingAllowed = decoded.DeviceSharingAllowed
	s.DiagnosticDataBlockSubmission = decoded.DiagnosticDataBlockSubmission
	s.FactoryResetBlocked = decoded.FactoryResetBlocked
	s.GoogleAccountBlockAutoSync = decoded.GoogleAccountBlockAutoSync
	s.GooglePlayStoreBlocked = decoded.GooglePlayStoreBlocked
	s.KioskModeBlockSleepButton = decoded.KioskModeBlockSleepButton
	s.KioskModeBlockVolumeButtons = decoded.KioskModeBlockVolumeButtons
	s.LocationServicesBlocked = decoded.LocationServicesBlocked
	s.NfcBlocked = decoded.NfcBlocked
	s.PasswordBlockFingerprintUnlock = decoded.PasswordBlockFingerprintUnlock
	s.PasswordBlockTrustAgents = decoded.PasswordBlockTrustAgents
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinutesOfInactivityBeforeScreenTimeout = decoded.PasswordMinutesOfInactivityBeforeScreenTimeout
	s.PasswordPreviousPasswordBlockCount = decoded.PasswordPreviousPasswordBlockCount
	s.PasswordRequired = decoded.PasswordRequired
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PasswordSignInFailureCountBeforeFactoryReset = decoded.PasswordSignInFailureCountBeforeFactoryReset
	s.PowerOffBlocked = decoded.PowerOffBlocked
	s.RequiredPasswordComplexity = decoded.RequiredPasswordComplexity
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.SecurityRequireVerifyApps = decoded.SecurityRequireVerifyApps
	s.StorageBlockGoogleBackup = decoded.StorageBlockGoogleBackup
	s.StorageBlockRemovableStorage = decoded.StorageBlockRemovableStorage
	s.StorageRequireDeviceEncryption = decoded.StorageRequireDeviceEncryption
	s.StorageRequireRemovableStorageEncryption = decoded.StorageRequireRemovableStorageEncryption
	s.VoiceAssistantBlocked = decoded.VoiceAssistantBlocked
	s.VoiceDialingBlocked = decoded.VoiceDialingBlocked
	s.WebBrowserBlockAutofill = decoded.WebBrowserBlockAutofill
	s.WebBrowserBlockJavaScript = decoded.WebBrowserBlockJavaScript
	s.WebBrowserBlockPopups = decoded.WebBrowserBlockPopups
	s.WebBrowserBlocked = decoded.WebBrowserBlocked
	s.WebBrowserCookieSettings = decoded.WebBrowserCookieSettings
	s.WiFiBlocked = decoded.WiFiBlocked
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
		return fmt.Errorf("unmarshaling AndroidGeneralDeviceConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appsHideList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppsHideList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppsHideList' for 'AndroidGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppsHideList = &output
	}

	if v, ok := temp["appsInstallAllowList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppsInstallAllowList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppsInstallAllowList' for 'AndroidGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppsInstallAllowList = &output
	}

	if v, ok := temp["appsLaunchBlockList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppsLaunchBlockList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppsLaunchBlockList' for 'AndroidGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppsLaunchBlockList = &output
	}

	if v, ok := temp["compliantAppsList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CompliantAppsList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CompliantAppsList' for 'AndroidGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CompliantAppsList = &output
	}

	if v, ok := temp["kioskModeApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling KioskModeApps into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'KioskModeApps' for 'AndroidGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.KioskModeApps = &output
	}

	return nil
}
