package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = WindowsPhone81GeneralConfiguration{}

type WindowsPhone81GeneralConfiguration struct {
	// Value indicating whether this policy only applies to Windows Phone 8.1. This property is read-only.
	ApplyOnlyToWindowsPhone81 *bool `json:"applyOnlyToWindowsPhone81,omitempty"`

	// Indicates whether or not to block copy paste.
	AppsBlockCopyPaste *bool `json:"appsBlockCopyPaste,omitempty"`

	// Indicates whether or not to block bluetooth.
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`

	// Indicates whether or not to block camera.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`

	// Indicates whether or not to block Wi-Fi tethering. Has no impact if Wi-Fi is blocked.
	CellularBlockWifiTethering *bool `json:"cellularBlockWifiTethering,omitempty"`

	// Possible values of the compliance app list.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`

	// List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection
	// can contain a maximum of 10000 elements.
	CompliantAppsList *[]AppListItem `json:"compliantAppsList,omitempty"`

	// Indicates whether or not to block diagnostic data submission.
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`

	// Indicates whether or not to block custom email accounts.
	EmailBlockAddingAccounts *bool `json:"emailBlockAddingAccounts,omitempty"`

	// Indicates whether or not to block location services.
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`

	// Indicates whether or not to block using a Microsoft Account.
	MicrosoftAccountBlocked *bool `json:"microsoftAccountBlocked,omitempty"`

	// Indicates whether or not to block Near-Field Communication.
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`

	// Indicates whether or not to block syncing the calendar.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// Number of days before the password expires.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Number of character sets a password must contain.
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// Minimum length of passwords.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity before screen timeout.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Number of previous passwords to block. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Indicates whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Number of sign in failures allowed before factory reset.
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// Indicates whether or not to block screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Indicates whether or not to block removable storage.
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`

	// Indicates whether or not to require encryption.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`

	// Indicates whether or not to block the web browser.
	WebBrowserBlocked *bool `json:"webBrowserBlocked,omitempty"`

	// Indicates whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
	WifiBlockAutomaticConnectHotspots *bool `json:"wifiBlockAutomaticConnectHotspots,omitempty"`

	// Indicates whether or not to block Wi-Fi hotspot reporting. Has no impact if Wi-Fi is blocked.
	WifiBlockHotspotReporting *bool `json:"wifiBlockHotspotReporting,omitempty"`

	// Indicates whether or not to block Wi-Fi.
	WifiBlocked *bool `json:"wifiBlocked,omitempty"`

	// Indicates whether or not to block the Windows Store.
	WindowsStoreBlocked *bool `json:"windowsStoreBlocked,omitempty"`

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

func (s WindowsPhone81GeneralConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s WindowsPhone81GeneralConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = WindowsPhone81GeneralConfiguration{}

func (s WindowsPhone81GeneralConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper WindowsPhone81GeneralConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WindowsPhone81GeneralConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsPhone81GeneralConfiguration: %+v", err)
	}

	delete(decoded, "applyOnlyToWindowsPhone81")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsPhone81GeneralConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WindowsPhone81GeneralConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &WindowsPhone81GeneralConfiguration{}

func (s *WindowsPhone81GeneralConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplyOnlyToWindowsPhone81                      *bool                                        `json:"applyOnlyToWindowsPhone81,omitempty"`
		AppsBlockCopyPaste                             *bool                                        `json:"appsBlockCopyPaste,omitempty"`
		BluetoothBlocked                               *bool                                        `json:"bluetoothBlocked,omitempty"`
		CameraBlocked                                  *bool                                        `json:"cameraBlocked,omitempty"`
		CellularBlockWifiTethering                     *bool                                        `json:"cellularBlockWifiTethering,omitempty"`
		CompliantAppListType                           *AppListType                                 `json:"compliantAppListType,omitempty"`
		DiagnosticDataBlockSubmission                  *bool                                        `json:"diagnosticDataBlockSubmission,omitempty"`
		EmailBlockAddingAccounts                       *bool                                        `json:"emailBlockAddingAccounts,omitempty"`
		LocationServicesBlocked                        *bool                                        `json:"locationServicesBlocked,omitempty"`
		MicrosoftAccountBlocked                        *bool                                        `json:"microsoftAccountBlocked,omitempty"`
		NfcBlocked                                     *bool                                        `json:"nfcBlocked,omitempty"`
		PasswordBlockSimple                            *bool                                        `json:"passwordBlockSimple,omitempty"`
		PasswordExpirationDays                         nullable.Type[int64]                         `json:"passwordExpirationDays,omitempty"`
		PasswordMinimumCharacterSetCount               nullable.Type[int64]                         `json:"passwordMinimumCharacterSetCount,omitempty"`
		PasswordMinimumLength                          nullable.Type[int64]                         `json:"passwordMinimumLength,omitempty"`
		PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64]                         `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasswordPreviousPasswordBlockCount             nullable.Type[int64]                         `json:"passwordPreviousPasswordBlockCount,omitempty"`
		PasswordRequired                               *bool                                        `json:"passwordRequired,omitempty"`
		PasswordRequiredType                           *RequiredPasswordType                        `json:"passwordRequiredType,omitempty"`
		PasswordSignInFailureCountBeforeFactoryReset   nullable.Type[int64]                         `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
		ScreenCaptureBlocked                           *bool                                        `json:"screenCaptureBlocked,omitempty"`
		StorageBlockRemovableStorage                   *bool                                        `json:"storageBlockRemovableStorage,omitempty"`
		StorageRequireEncryption                       *bool                                        `json:"storageRequireEncryption,omitempty"`
		WebBrowserBlocked                              *bool                                        `json:"webBrowserBlocked,omitempty"`
		WifiBlockAutomaticConnectHotspots              *bool                                        `json:"wifiBlockAutomaticConnectHotspots,omitempty"`
		WifiBlockHotspotReporting                      *bool                                        `json:"wifiBlockHotspotReporting,omitempty"`
		WifiBlocked                                    *bool                                        `json:"wifiBlocked,omitempty"`
		WindowsStoreBlocked                            *bool                                        `json:"windowsStoreBlocked,omitempty"`
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

	s.ApplyOnlyToWindowsPhone81 = decoded.ApplyOnlyToWindowsPhone81
	s.AppsBlockCopyPaste = decoded.AppsBlockCopyPaste
	s.BluetoothBlocked = decoded.BluetoothBlocked
	s.CameraBlocked = decoded.CameraBlocked
	s.CellularBlockWifiTethering = decoded.CellularBlockWifiTethering
	s.CompliantAppListType = decoded.CompliantAppListType
	s.DiagnosticDataBlockSubmission = decoded.DiagnosticDataBlockSubmission
	s.EmailBlockAddingAccounts = decoded.EmailBlockAddingAccounts
	s.LocationServicesBlocked = decoded.LocationServicesBlocked
	s.MicrosoftAccountBlocked = decoded.MicrosoftAccountBlocked
	s.NfcBlocked = decoded.NfcBlocked
	s.PasswordBlockSimple = decoded.PasswordBlockSimple
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMinimumCharacterSetCount = decoded.PasswordMinimumCharacterSetCount
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinutesOfInactivityBeforeScreenTimeout = decoded.PasswordMinutesOfInactivityBeforeScreenTimeout
	s.PasswordPreviousPasswordBlockCount = decoded.PasswordPreviousPasswordBlockCount
	s.PasswordRequired = decoded.PasswordRequired
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PasswordSignInFailureCountBeforeFactoryReset = decoded.PasswordSignInFailureCountBeforeFactoryReset
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.StorageBlockRemovableStorage = decoded.StorageBlockRemovableStorage
	s.StorageRequireEncryption = decoded.StorageRequireEncryption
	s.WebBrowserBlocked = decoded.WebBrowserBlocked
	s.WifiBlockAutomaticConnectHotspots = decoded.WifiBlockAutomaticConnectHotspots
	s.WifiBlockHotspotReporting = decoded.WifiBlockHotspotReporting
	s.WifiBlocked = decoded.WifiBlocked
	s.WindowsStoreBlocked = decoded.WindowsStoreBlocked
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
		return fmt.Errorf("unmarshaling WindowsPhone81GeneralConfiguration into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'CompliantAppsList' for 'WindowsPhone81GeneralConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CompliantAppsList = &output
	}

	return nil
}
