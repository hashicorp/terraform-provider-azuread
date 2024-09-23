package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows81GeneralConfiguration{}

type Windows81GeneralConfiguration struct {
	// Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a
	// Microsoft account.
	AccountsBlockAddingNonMicrosoftAccountEmail *bool `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`

	// Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
	ApplyOnlyToWindows81 *bool `json:"applyOnlyToWindows81,omitempty"`

	// Indicates whether or not to block auto fill.
	BrowserBlockAutofill *bool `json:"browserBlockAutofill,omitempty"`

	// Indicates whether or not to block automatic detection of Intranet sites.
	BrowserBlockAutomaticDetectionOfIntranetSites *bool `json:"browserBlockAutomaticDetectionOfIntranetSites,omitempty"`

	// Indicates whether or not to block enterprise mode access.
	BrowserBlockEnterpriseModeAccess *bool `json:"browserBlockEnterpriseModeAccess,omitempty"`

	// Indicates whether or not to Block the user from using JavaScript.
	BrowserBlockJavaScript *bool `json:"browserBlockJavaScript,omitempty"`

	// Indicates whether or not to block plug-ins.
	BrowserBlockPlugins *bool `json:"browserBlockPlugins,omitempty"`

	// Indicates whether or not to block popups.
	BrowserBlockPopups *bool `json:"browserBlockPopups,omitempty"`

	// Indicates whether or not to Block the user from sending the do not track header.
	BrowserBlockSendingDoNotTrackHeader *bool `json:"browserBlockSendingDoNotTrackHeader,omitempty"`

	// Indicates whether or not to block a single word entry on Intranet sites.
	BrowserBlockSingleWordEntryOnIntranetSites *bool `json:"browserBlockSingleWordEntryOnIntranetSites,omitempty"`

	// The enterprise mode site list location. Could be a local file, local network or http location.
	BrowserEnterpriseModeSiteListLocation nullable.Type[string] `json:"browserEnterpriseModeSiteListLocation,omitempty"`

	// Possible values for internet site security level.
	BrowserInternetSecurityLevel *InternetSiteSecurityLevel `json:"browserInternetSecurityLevel,omitempty"`

	// Possible values for site security level.
	BrowserIntranetSecurityLevel *SiteSecurityLevel `json:"browserIntranetSecurityLevel,omitempty"`

	// The logging report location.
	BrowserLoggingReportLocation nullable.Type[string] `json:"browserLoggingReportLocation,omitempty"`

	// Indicates whether or not to require a firewall.
	BrowserRequireFirewall *bool `json:"browserRequireFirewall,omitempty"`

	// Indicates whether or not to require fraud warning.
	BrowserRequireFraudWarning *bool `json:"browserRequireFraudWarning,omitempty"`

	// Indicates whether or not to require high security for restricted sites.
	BrowserRequireHighSecurityForRestrictedSites *bool `json:"browserRequireHighSecurityForRestrictedSites,omitempty"`

	// Indicates whether or not to require the user to use the smart screen filter.
	BrowserRequireSmartScreen *bool `json:"browserRequireSmartScreen,omitempty"`

	// Possible values for site security level.
	BrowserTrustedSitesSecurityLevel *SiteSecurityLevel `json:"browserTrustedSitesSecurityLevel,omitempty"`

	// Indicates whether or not to block data roaming.
	CellularBlockDataRoaming *bool `json:"cellularBlockDataRoaming,omitempty"`

	// Indicates whether or not to block diagnostic data submission.
	DiagnosticsBlockDataSubmission *bool `json:"diagnosticsBlockDataSubmission,omitempty"`

	// Indicates whether or not to Block the user from using a pictures password and pin.
	PasswordBlockPicturePasswordAndPin *bool `json:"passwordBlockPicturePasswordAndPin,omitempty"`

	// Password expiration in days.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// The minimum password length.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// The minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// The number of previous passwords to prevent re-use of. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// The number of sign in failures before factory reset.
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// Indicates whether or not to require encryption on a mobile device.
	StorageRequireDeviceEncryption *bool `json:"storageRequireDeviceEncryption,omitempty"`

	// Indicates whether or not to require automatic updates.
	UpdatesRequireAutomaticUpdates *bool `json:"updatesRequireAutomaticUpdates,omitempty"`

	// Possible values for Windows user account control settings.
	UserAccountControlSettings *WindowsUserAccountControlSettings `json:"userAccountControlSettings,omitempty"`

	// The work folders url.
	WorkFoldersUrl nullable.Type[string] `json:"workFoldersUrl,omitempty"`

	// Fields inherited from DeviceConfiguration

	// The list of assignments for the device configuration profile.
	Assignments *[]DeviceConfigurationAssignment `json:"assignments,omitempty"`

	// DateTime the object was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Admin provided description of the Device Configuration.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Device Configuration Setting State Device Summary
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary `json:"deviceSettingStateSummaries,omitempty"`

	// Device Configuration devices status overview
	DeviceStatusOverview *DeviceConfigurationDeviceOverview `json:"deviceStatusOverview,omitempty"`

	// Device configuration installation status by device.
	DeviceStatuses *[]DeviceConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`

	// Admin provided name of the device configuration.
	DisplayName *string `json:"displayName,omitempty"`

	// DateTime the object was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

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

func (s Windows81GeneralConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
	return BaseDeviceConfigurationImpl{
		Assignments:                 s.Assignments,
		CreatedDateTime:             s.CreatedDateTime,
		Description:                 s.Description,
		DeviceSettingStateSummaries: s.DeviceSettingStateSummaries,
		DeviceStatusOverview:        s.DeviceStatusOverview,
		DeviceStatuses:              s.DeviceStatuses,
		DisplayName:                 s.DisplayName,
		LastModifiedDateTime:        s.LastModifiedDateTime,
		UserStatusOverview:          s.UserStatusOverview,
		UserStatuses:                s.UserStatuses,
		Version:                     s.Version,
		Id:                          s.Id,
		ODataId:                     s.ODataId,
		ODataType:                   s.ODataType,
	}
}

func (s Windows81GeneralConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows81GeneralConfiguration{}

func (s Windows81GeneralConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows81GeneralConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows81GeneralConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows81GeneralConfiguration: %+v", err)
	}

	delete(decoded, "applyOnlyToWindows81")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows81GeneralConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows81GeneralConfiguration: %+v", err)
	}

	return encoded, nil
}
