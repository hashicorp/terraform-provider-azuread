package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = MacOSGeneralDeviceConfiguration{}

type MacOSGeneralDeviceConfiguration struct {
	// When TRUE, activation lock is allowed when the devices is in the supervised mode. When FALSE, activation lock is not
	// allowed. Default is false.
	ActivationLockWhenSupervisedAllowed *bool `json:"activationLockWhenSupervisedAllowed,omitempty"`

	// Yes prevents users from adding friends to Game Center. Available for devices running macOS versions 10.13 and later.
	AddingGameCenterFriendsBlocked *bool `json:"addingGameCenterFriendsBlocked,omitempty"`

	// Indicates whether or not to allow AirDrop.
	AirDropBlocked *bool `json:"airDropBlocked,omitempty"`

	// Indicates whether or to block users from unlocking their Mac with Apple Watch.
	AppleWatchBlockAutoUnlock *bool `json:"appleWatchBlockAutoUnlock,omitempty"`

	// Indicates whether or not to block the user from accessing the camera of the device.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`

	// Indicates whether or not to allow remote screen observation by Classroom app. Requires MDM enrollment via Apple
	// School Manager or Apple Business Manager.
	ClassroomAppBlockRemoteScreenObservation *bool `json:"classroomAppBlockRemoteScreenObservation,omitempty"`

	// Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to
	// view a student's screen without prompting. Requires MDM enrollment via Apple School Manager or Apple Business
	// Manager.
	ClassroomAppForceUnpromptedScreenObservation *bool `json:"classroomAppForceUnpromptedScreenObservation,omitempty"`

	// Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student.
	// Requires MDM enrollment via Apple School Manager or Apple Business Manager.
	ClassroomForceAutomaticallyJoinClasses *bool `json:"classroomForceAutomaticallyJoinClasses,omitempty"`

	// Indicates whether a student enrolled in an unmanaged course via Classroom will be required to request permission from
	// the teacher when attempting to leave the course. Requires MDM enrollment via Apple School Manager or Apple Business
	// Manager.
	ClassroomForceRequestPermissionToLeaveClasses *bool `json:"classroomForceRequestPermissionToLeaveClasses,omitempty"`

	// Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Requires MDM
	// enrollment via Apple School Manager or Apple Business Manager.
	ClassroomForceUnpromptedAppAndDeviceLock *bool `json:"classroomForceUnpromptedAppAndDeviceLock,omitempty"`

	// Possible values of the compliance app list.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`

	// List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection
	// can contain a maximum of 10000 elements.
	CompliantAppsList *[]AppListItem `json:"compliantAppsList,omitempty"`

	// Indicates whether or not to allow content caching.
	ContentCachingBlocked *bool `json:"contentCachingBlocked,omitempty"`

	// Indicates whether or not to block definition lookup.
	DefinitionLookupBlocked *bool `json:"definitionLookupBlocked,omitempty"`

	// An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
	EmailInDomainSuffixes *[]string `json:"emailInDomainSuffixes,omitempty"`

	// TRUE disables the reset option on supervised devices. FALSE enables the reset option on supervised devices. Available
	// for devices running macOS versions 12.0 and later.
	EraseContentAndSettingsBlocked *bool `json:"eraseContentAndSettingsBlocked,omitempty"`

	// Yes disables Game Center, and the Game Center icon is removed from the Home screen. Available for devices running
	// macOS versions 10.13 and later.
	GameCenterBlocked *bool `json:"gameCenterBlocked,omitempty"`

	// Indicates whether or not to block the user from continuing work that they started on a MacOS device on another iOS or
	// MacOS device (MacOS 10.15 or later).
	ICloudBlockActivityContinuation *bool `json:"iCloudBlockActivityContinuation,omitempty"`

	// Indicates whether or not to block iCloud from syncing contacts.
	ICloudBlockAddressBook *bool `json:"iCloudBlockAddressBook,omitempty"`

	// Indicates whether or not to block iCloud from syncing bookmarks.
	ICloudBlockBookmarks *bool `json:"iCloudBlockBookmarks,omitempty"`

	// Indicates whether or not to block iCloud from syncing calendars.
	ICloudBlockCalendar *bool `json:"iCloudBlockCalendar,omitempty"`

	// Indicates whether or not to block iCloud document sync.
	ICloudBlockDocumentSync *bool `json:"iCloudBlockDocumentSync,omitempty"`

	// Indicates whether or not to block iCloud from syncing mail.
	ICloudBlockMail *bool `json:"iCloudBlockMail,omitempty"`

	// Indicates whether or not to block iCloud from syncing notes.
	ICloudBlockNotes *bool `json:"iCloudBlockNotes,omitempty"`

	// Indicates whether or not to block iCloud Photo Library.
	ICloudBlockPhotoLibrary *bool `json:"iCloudBlockPhotoLibrary,omitempty"`

	// Indicates whether or not to block iCloud from syncing reminders.
	ICloudBlockReminders *bool `json:"iCloudBlockReminders,omitempty"`

	// When TRUE the synchronization of cloud desktop and documents is blocked. When FALSE, synchronization of the cloud
	// desktop and documents are allowed. Available for devices running macOS 10.12.4 and later.
	ICloudDesktopAndDocumentsBlocked *bool `json:"iCloudDesktopAndDocumentsBlocked,omitempty"`

	// iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity
	// across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device.
	// Available for devices running macOS 12 and later.
	ICloudPrivateRelayBlocked *bool `json:"iCloudPrivateRelayBlocked,omitempty"`

	// Indicates whether or not to block files from being transferred using iTunes.
	ITunesBlockFileSharing *bool `json:"iTunesBlockFileSharing,omitempty"`

	// Indicates whether or not to block Music service and revert Music app to classic mode.
	ITunesBlockMusicService *bool `json:"iTunesBlockMusicService,omitempty"`

	// Indicates whether or not to block the user from using dictation input.
	KeyboardBlockDictation *bool `json:"keyboardBlockDictation,omitempty"`

	// Indicates whether or not iCloud keychain synchronization is blocked (macOS 10.12 and later).
	KeychainBlockCloudSync *bool `json:"keychainBlockCloudSync,omitempty"`

	// TRUE prevents multiplayer gaming when using Game Center. FALSE allows multiplayer gaming when using Game Center.
	// Available for devices running macOS versions 10.13 and later.
	MultiplayerGamingBlocked *bool `json:"multiplayerGamingBlocked,omitempty"`

	// Indicates whether or not to block sharing passwords with the AirDrop passwords feature.
	PasswordBlockAirDropSharing *bool `json:"passwordBlockAirDropSharing,omitempty"`

	// Indicates whether or not to block the AutoFill Passwords feature.
	PasswordBlockAutoFill *bool `json:"passwordBlockAutoFill,omitempty"`

	// Indicates whether or not to block fingerprint unlock.
	PasswordBlockFingerprintUnlock *bool `json:"passwordBlockFingerprintUnlock,omitempty"`

	// Indicates whether or not to allow passcode modification.
	PasswordBlockModification *bool `json:"passwordBlockModification,omitempty"`

	// Indicates whether or not to block requesting passwords from nearby devices.
	PasswordBlockProximityRequests *bool `json:"passwordBlockProximityRequests,omitempty"`

	// Block simple passwords.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// Number of days before the password expires.
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// The number of allowed failed attempts to enter the passcode at the device's lock screen. Valid values 2 to 11
	PasswordMaximumAttemptCount nullable.Type[int64] `json:"passwordMaximumAttemptCount,omitempty"`

	// Number of character sets a password must contain. Valid values 0 to 4
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// Minimum length of passwords.
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Minutes of inactivity required before a password is required.
	PasswordMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`

	// Minutes of inactivity required before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// The number of minutes before the login is reset after the maximum number of unsuccessful login attempts is reached.
	PasswordMinutesUntilFailedLoginReset nullable.Type[int64] `json:"passwordMinutesUntilFailedLoginReset,omitempty"`

	// Number of previous passwords to block.
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// List of privacy preference policy controls. This collection can contain a maximum of 10000 elements.
	PrivacyAccessControls *[]MacOSPrivacyAccessControlItem `json:"privacyAccessControls,omitempty"`

	// Indicates whether or not to block the user from using Auto fill in Safari.
	SafariBlockAutofill *bool `json:"safariBlockAutofill,omitempty"`

	// Indicates whether or not to block the user from taking Screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Specify the number of days (1-90) to delay visibility of major OS software updates. Available for devices running
	// macOS versions 11.3 and later. Valid values 0 to 90
	SoftwareUpdateMajorOSDeferredInstallDelayInDays nullable.Type[int64] `json:"softwareUpdateMajorOSDeferredInstallDelayInDays,omitempty"`

	// Specify the number of days (1-90) to delay visibility of minor OS software updates. Available for devices running
	// macOS versions 11.3 and later. Valid values 0 to 90
	SoftwareUpdateMinorOSDeferredInstallDelayInDays nullable.Type[int64] `json:"softwareUpdateMinorOSDeferredInstallDelayInDays,omitempty"`

	// Specify the number of days (1-90) to delay visibility of non-OS software updates. Available for devices running macOS
	// versions 11.3 and later. Valid values 0 to 90
	SoftwareUpdateNonOSDeferredInstallDelayInDays nullable.Type[int64] `json:"softwareUpdateNonOSDeferredInstallDelayInDays,omitempty"`

	// Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
	SoftwareUpdatesEnforcedDelayInDays nullable.Type[int64] `json:"softwareUpdatesEnforcedDelayInDays,omitempty"`

	// Indicates whether or not to block Spotlight from returning any results from an Internet search.
	SpotlightBlockInternetResults *bool `json:"spotlightBlockInternetResults,omitempty"`

	// Maximum hours after which the user must enter their password to unlock the device instead of using Touch ID.
	// Available for devices running macOS 12 and later. Valid values 0 to 2147483647
	TouchIdTimeoutInHours nullable.Type[int64] `json:"touchIdTimeoutInHours,omitempty"`

	// Determines whether to delay OS and/or app updates for macOS. Possible values are: none, delayOSUpdateVisibility,
	// delayAppUpdateVisibility, unknownFutureValue, delayMajorOsUpdateVisibility.
	UpdateDelayPolicy *MacOSSoftwareUpdateDelayPolicy `json:"updateDelayPolicy,omitempty"`

	// TRUE prevents the wallpaper from being changed. FALSE allows the wallpaper to be changed. Available for devices
	// running macOS versions 10.13 and later.
	WallpaperModificationBlocked *bool `json:"wallpaperModificationBlocked,omitempty"`

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

func (s MacOSGeneralDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s MacOSGeneralDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MacOSGeneralDeviceConfiguration{}

func (s MacOSGeneralDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MacOSGeneralDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MacOSGeneralDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MacOSGeneralDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.macOSGeneralDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MacOSGeneralDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MacOSGeneralDeviceConfiguration{}

func (s *MacOSGeneralDeviceConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ActivationLockWhenSupervisedAllowed             *bool                                        `json:"activationLockWhenSupervisedAllowed,omitempty"`
		AddingGameCenterFriendsBlocked                  *bool                                        `json:"addingGameCenterFriendsBlocked,omitempty"`
		AirDropBlocked                                  *bool                                        `json:"airDropBlocked,omitempty"`
		AppleWatchBlockAutoUnlock                       *bool                                        `json:"appleWatchBlockAutoUnlock,omitempty"`
		CameraBlocked                                   *bool                                        `json:"cameraBlocked,omitempty"`
		ClassroomAppBlockRemoteScreenObservation        *bool                                        `json:"classroomAppBlockRemoteScreenObservation,omitempty"`
		ClassroomAppForceUnpromptedScreenObservation    *bool                                        `json:"classroomAppForceUnpromptedScreenObservation,omitempty"`
		ClassroomForceAutomaticallyJoinClasses          *bool                                        `json:"classroomForceAutomaticallyJoinClasses,omitempty"`
		ClassroomForceRequestPermissionToLeaveClasses   *bool                                        `json:"classroomForceRequestPermissionToLeaveClasses,omitempty"`
		ClassroomForceUnpromptedAppAndDeviceLock        *bool                                        `json:"classroomForceUnpromptedAppAndDeviceLock,omitempty"`
		CompliantAppListType                            *AppListType                                 `json:"compliantAppListType,omitempty"`
		ContentCachingBlocked                           *bool                                        `json:"contentCachingBlocked,omitempty"`
		DefinitionLookupBlocked                         *bool                                        `json:"definitionLookupBlocked,omitempty"`
		EmailInDomainSuffixes                           *[]string                                    `json:"emailInDomainSuffixes,omitempty"`
		EraseContentAndSettingsBlocked                  *bool                                        `json:"eraseContentAndSettingsBlocked,omitempty"`
		GameCenterBlocked                               *bool                                        `json:"gameCenterBlocked,omitempty"`
		ICloudBlockActivityContinuation                 *bool                                        `json:"iCloudBlockActivityContinuation,omitempty"`
		ICloudBlockAddressBook                          *bool                                        `json:"iCloudBlockAddressBook,omitempty"`
		ICloudBlockBookmarks                            *bool                                        `json:"iCloudBlockBookmarks,omitempty"`
		ICloudBlockCalendar                             *bool                                        `json:"iCloudBlockCalendar,omitempty"`
		ICloudBlockDocumentSync                         *bool                                        `json:"iCloudBlockDocumentSync,omitempty"`
		ICloudBlockMail                                 *bool                                        `json:"iCloudBlockMail,omitempty"`
		ICloudBlockNotes                                *bool                                        `json:"iCloudBlockNotes,omitempty"`
		ICloudBlockPhotoLibrary                         *bool                                        `json:"iCloudBlockPhotoLibrary,omitempty"`
		ICloudBlockReminders                            *bool                                        `json:"iCloudBlockReminders,omitempty"`
		ICloudDesktopAndDocumentsBlocked                *bool                                        `json:"iCloudDesktopAndDocumentsBlocked,omitempty"`
		ICloudPrivateRelayBlocked                       *bool                                        `json:"iCloudPrivateRelayBlocked,omitempty"`
		ITunesBlockFileSharing                          *bool                                        `json:"iTunesBlockFileSharing,omitempty"`
		ITunesBlockMusicService                         *bool                                        `json:"iTunesBlockMusicService,omitempty"`
		KeyboardBlockDictation                          *bool                                        `json:"keyboardBlockDictation,omitempty"`
		KeychainBlockCloudSync                          *bool                                        `json:"keychainBlockCloudSync,omitempty"`
		MultiplayerGamingBlocked                        *bool                                        `json:"multiplayerGamingBlocked,omitempty"`
		PasswordBlockAirDropSharing                     *bool                                        `json:"passwordBlockAirDropSharing,omitempty"`
		PasswordBlockAutoFill                           *bool                                        `json:"passwordBlockAutoFill,omitempty"`
		PasswordBlockFingerprintUnlock                  *bool                                        `json:"passwordBlockFingerprintUnlock,omitempty"`
		PasswordBlockModification                       *bool                                        `json:"passwordBlockModification,omitempty"`
		PasswordBlockProximityRequests                  *bool                                        `json:"passwordBlockProximityRequests,omitempty"`
		PasswordBlockSimple                             *bool                                        `json:"passwordBlockSimple,omitempty"`
		PasswordExpirationDays                          nullable.Type[int64]                         `json:"passwordExpirationDays,omitempty"`
		PasswordMaximumAttemptCount                     nullable.Type[int64]                         `json:"passwordMaximumAttemptCount,omitempty"`
		PasswordMinimumCharacterSetCount                nullable.Type[int64]                         `json:"passwordMinimumCharacterSetCount,omitempty"`
		PasswordMinimumLength                           nullable.Type[int64]                         `json:"passwordMinimumLength,omitempty"`
		PasswordMinutesOfInactivityBeforeLock           nullable.Type[int64]                         `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
		PasswordMinutesOfInactivityBeforeScreenTimeout  nullable.Type[int64]                         `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasswordMinutesUntilFailedLoginReset            nullable.Type[int64]                         `json:"passwordMinutesUntilFailedLoginReset,omitempty"`
		PasswordPreviousPasswordBlockCount              nullable.Type[int64]                         `json:"passwordPreviousPasswordBlockCount,omitempty"`
		PasswordRequired                                *bool                                        `json:"passwordRequired,omitempty"`
		PasswordRequiredType                            *RequiredPasswordType                        `json:"passwordRequiredType,omitempty"`
		PrivacyAccessControls                           *[]MacOSPrivacyAccessControlItem             `json:"privacyAccessControls,omitempty"`
		SafariBlockAutofill                             *bool                                        `json:"safariBlockAutofill,omitempty"`
		ScreenCaptureBlocked                            *bool                                        `json:"screenCaptureBlocked,omitempty"`
		SoftwareUpdateMajorOSDeferredInstallDelayInDays nullable.Type[int64]                         `json:"softwareUpdateMajorOSDeferredInstallDelayInDays,omitempty"`
		SoftwareUpdateMinorOSDeferredInstallDelayInDays nullable.Type[int64]                         `json:"softwareUpdateMinorOSDeferredInstallDelayInDays,omitempty"`
		SoftwareUpdateNonOSDeferredInstallDelayInDays   nullable.Type[int64]                         `json:"softwareUpdateNonOSDeferredInstallDelayInDays,omitempty"`
		SoftwareUpdatesEnforcedDelayInDays              nullable.Type[int64]                         `json:"softwareUpdatesEnforcedDelayInDays,omitempty"`
		SpotlightBlockInternetResults                   *bool                                        `json:"spotlightBlockInternetResults,omitempty"`
		TouchIdTimeoutInHours                           nullable.Type[int64]                         `json:"touchIdTimeoutInHours,omitempty"`
		UpdateDelayPolicy                               *MacOSSoftwareUpdateDelayPolicy              `json:"updateDelayPolicy,omitempty"`
		WallpaperModificationBlocked                    *bool                                        `json:"wallpaperModificationBlocked,omitempty"`
		Assignments                                     *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                                 *string                                      `json:"createdDateTime,omitempty"`
		Description                                     nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode     *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition      *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion      *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                     *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                            *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                  *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                     *string                                      `json:"displayName,omitempty"`
		GroupAssignments                                *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                            *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                 *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                               *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                              *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                    *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                         *int64                                       `json:"version,omitempty"`
		Id                                              *string                                      `json:"id,omitempty"`
		ODataId                                         *string                                      `json:"@odata.id,omitempty"`
		ODataType                                       *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ActivationLockWhenSupervisedAllowed = decoded.ActivationLockWhenSupervisedAllowed
	s.AddingGameCenterFriendsBlocked = decoded.AddingGameCenterFriendsBlocked
	s.AirDropBlocked = decoded.AirDropBlocked
	s.AppleWatchBlockAutoUnlock = decoded.AppleWatchBlockAutoUnlock
	s.CameraBlocked = decoded.CameraBlocked
	s.ClassroomAppBlockRemoteScreenObservation = decoded.ClassroomAppBlockRemoteScreenObservation
	s.ClassroomAppForceUnpromptedScreenObservation = decoded.ClassroomAppForceUnpromptedScreenObservation
	s.ClassroomForceAutomaticallyJoinClasses = decoded.ClassroomForceAutomaticallyJoinClasses
	s.ClassroomForceRequestPermissionToLeaveClasses = decoded.ClassroomForceRequestPermissionToLeaveClasses
	s.ClassroomForceUnpromptedAppAndDeviceLock = decoded.ClassroomForceUnpromptedAppAndDeviceLock
	s.CompliantAppListType = decoded.CompliantAppListType
	s.ContentCachingBlocked = decoded.ContentCachingBlocked
	s.DefinitionLookupBlocked = decoded.DefinitionLookupBlocked
	s.EmailInDomainSuffixes = decoded.EmailInDomainSuffixes
	s.EraseContentAndSettingsBlocked = decoded.EraseContentAndSettingsBlocked
	s.GameCenterBlocked = decoded.GameCenterBlocked
	s.ICloudBlockActivityContinuation = decoded.ICloudBlockActivityContinuation
	s.ICloudBlockAddressBook = decoded.ICloudBlockAddressBook
	s.ICloudBlockBookmarks = decoded.ICloudBlockBookmarks
	s.ICloudBlockCalendar = decoded.ICloudBlockCalendar
	s.ICloudBlockDocumentSync = decoded.ICloudBlockDocumentSync
	s.ICloudBlockMail = decoded.ICloudBlockMail
	s.ICloudBlockNotes = decoded.ICloudBlockNotes
	s.ICloudBlockPhotoLibrary = decoded.ICloudBlockPhotoLibrary
	s.ICloudBlockReminders = decoded.ICloudBlockReminders
	s.ICloudDesktopAndDocumentsBlocked = decoded.ICloudDesktopAndDocumentsBlocked
	s.ICloudPrivateRelayBlocked = decoded.ICloudPrivateRelayBlocked
	s.ITunesBlockFileSharing = decoded.ITunesBlockFileSharing
	s.ITunesBlockMusicService = decoded.ITunesBlockMusicService
	s.KeyboardBlockDictation = decoded.KeyboardBlockDictation
	s.KeychainBlockCloudSync = decoded.KeychainBlockCloudSync
	s.MultiplayerGamingBlocked = decoded.MultiplayerGamingBlocked
	s.PasswordBlockAirDropSharing = decoded.PasswordBlockAirDropSharing
	s.PasswordBlockAutoFill = decoded.PasswordBlockAutoFill
	s.PasswordBlockFingerprintUnlock = decoded.PasswordBlockFingerprintUnlock
	s.PasswordBlockModification = decoded.PasswordBlockModification
	s.PasswordBlockProximityRequests = decoded.PasswordBlockProximityRequests
	s.PasswordBlockSimple = decoded.PasswordBlockSimple
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMaximumAttemptCount = decoded.PasswordMaximumAttemptCount
	s.PasswordMinimumCharacterSetCount = decoded.PasswordMinimumCharacterSetCount
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinutesOfInactivityBeforeLock = decoded.PasswordMinutesOfInactivityBeforeLock
	s.PasswordMinutesOfInactivityBeforeScreenTimeout = decoded.PasswordMinutesOfInactivityBeforeScreenTimeout
	s.PasswordMinutesUntilFailedLoginReset = decoded.PasswordMinutesUntilFailedLoginReset
	s.PasswordPreviousPasswordBlockCount = decoded.PasswordPreviousPasswordBlockCount
	s.PasswordRequired = decoded.PasswordRequired
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PrivacyAccessControls = decoded.PrivacyAccessControls
	s.SafariBlockAutofill = decoded.SafariBlockAutofill
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.SoftwareUpdateMajorOSDeferredInstallDelayInDays = decoded.SoftwareUpdateMajorOSDeferredInstallDelayInDays
	s.SoftwareUpdateMinorOSDeferredInstallDelayInDays = decoded.SoftwareUpdateMinorOSDeferredInstallDelayInDays
	s.SoftwareUpdateNonOSDeferredInstallDelayInDays = decoded.SoftwareUpdateNonOSDeferredInstallDelayInDays
	s.SoftwareUpdatesEnforcedDelayInDays = decoded.SoftwareUpdatesEnforcedDelayInDays
	s.SpotlightBlockInternetResults = decoded.SpotlightBlockInternetResults
	s.TouchIdTimeoutInHours = decoded.TouchIdTimeoutInHours
	s.UpdateDelayPolicy = decoded.UpdateDelayPolicy
	s.WallpaperModificationBlocked = decoded.WallpaperModificationBlocked
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
		return fmt.Errorf("unmarshaling MacOSGeneralDeviceConfiguration into map[string]json.RawMessage: %+v", err)
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
				return fmt.Errorf("unmarshaling index %d field 'CompliantAppsList' for 'MacOSGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CompliantAppsList = &output
	}

	return nil
}
