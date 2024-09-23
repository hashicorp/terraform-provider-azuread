package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = IosGeneralDeviceConfiguration{}

type IosGeneralDeviceConfiguration struct {
	// Indicates whether or not to allow account modification when the device is in supervised mode.
	AccountBlockModification *bool `json:"accountBlockModification,omitempty"`

	// Indicates whether or not to allow activation lock when the device is in the supervised mode.
	ActivationLockAllowWhenSupervised *bool `json:"activationLockAllowWhenSupervised,omitempty"`

	// Indicates whether or not to allow AirDrop when the device is in supervised mode.
	AirDropBlocked *bool `json:"airDropBlocked,omitempty"`

	// Indicates whether or not to cause AirDrop to be considered an unmanaged drop target (iOS 9.0 and later).
	AirDropForceUnmanagedDropTarget *bool `json:"airDropForceUnmanagedDropTarget,omitempty"`

	// Indicates whether or not to enforce all devices receiving AirPlay requests from this device to use a pairing
	// password.
	AirPlayForcePairingPasswordForOutgoingRequests *bool `json:"airPlayForcePairingPasswordForOutgoingRequests,omitempty"`

	// Indicates whether or not to block the automatic downloading of apps purchased on other devices when the device is in
	// supervised mode (iOS 9.0 and later).
	AppStoreBlockAutomaticDownloads *bool `json:"appStoreBlockAutomaticDownloads,omitempty"`

	// Indicates whether or not to block the user from making in app purchases.
	AppStoreBlockInAppPurchases *bool `json:"appStoreBlockInAppPurchases,omitempty"`

	// Indicates whether or not to block the App Store app, not restricting installation through Host apps. Applies to
	// supervised mode only (iOS 9.0 and later).
	AppStoreBlockUIAppInstallation *bool `json:"appStoreBlockUIAppInstallation,omitempty"`

	// Indicates whether or not to block the user from using the App Store. Requires a supervised device for iOS 13 and
	// later.
	AppStoreBlocked *bool `json:"appStoreBlocked,omitempty"`

	// Indicates whether or not to require a password when using the app store.
	AppStoreRequirePassword *bool `json:"appStoreRequirePassword,omitempty"`

	// Indicates whether or not to block the user from using News when the device is in supervised mode (iOS 9.0 and later).
	AppleNewsBlocked *bool `json:"appleNewsBlocked,omitempty"`

	// Indicates whether or not to allow Apple Watch pairing when the device is in supervised mode (iOS 9.0 and later).
	AppleWatchBlockPairing *bool `json:"appleWatchBlockPairing,omitempty"`

	// Indicates whether or not to force a paired Apple Watch to use Wrist Detection (iOS 8.2 and later).
	AppleWatchForceWristDetection *bool `json:"appleWatchForceWristDetection,omitempty"`

	// Gets or sets the list of iOS apps allowed to autonomously enter Single App Mode. Supervised only. iOS 7.0 and later.
	// This collection can contain a maximum of 500 elements.
	AppsSingleAppModeList *[]AppListItem `json:"appsSingleAppModeList,omitempty"`

	// List of apps in the visibility list (either visible/launchable apps list or hidden/unlaunchable apps list, controlled
	// by AppsVisibilityListType) (iOS 9.3 and later). This collection can contain a maximum of 10000 elements.
	AppsVisibilityList *[]AppListItem `json:"appsVisibilityList,omitempty"`

	// Possible values of the compliance app list.
	AppsVisibilityListType *AppListType `json:"appsVisibilityListType,omitempty"`

	// Indicates whether or not to allow modification of Bluetooth settings when the device is in supervised mode (iOS 10.0
	// and later).
	BluetoothBlockModification *bool `json:"bluetoothBlockModification,omitempty"`

	// Indicates whether or not to block the user from accessing the camera of the device. Requires a supervised device for
	// iOS 13 and later.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`

	// Indicates whether or not to block data roaming.
	CellularBlockDataRoaming *bool `json:"cellularBlockDataRoaming,omitempty"`

	// Indicates whether or not to block global background fetch while roaming.
	CellularBlockGlobalBackgroundFetchWhileRoaming *bool `json:"cellularBlockGlobalBackgroundFetchWhileRoaming,omitempty"`

	// Indicates whether or not to allow changes to cellular app data usage settings when the device is in supervised mode.
	CellularBlockPerAppDataModification *bool `json:"cellularBlockPerAppDataModification,omitempty"`

	// Indicates whether or not to block Personal Hotspot.
	CellularBlockPersonalHotspot *bool `json:"cellularBlockPersonalHotspot,omitempty"`

	// Indicates whether or not to block voice roaming.
	CellularBlockVoiceRoaming *bool `json:"cellularBlockVoiceRoaming,omitempty"`

	// Indicates whether or not to block untrusted TLS certificates.
	CertificatesBlockUntrustedTlsCertificates *bool `json:"certificatesBlockUntrustedTlsCertificates,omitempty"`

	// Indicates whether or not to allow remote screen observation by Classroom app when the device is in supervised mode
	// (iOS 9.3 and later).
	ClassroomAppBlockRemoteScreenObservation *bool `json:"classroomAppBlockRemoteScreenObservation,omitempty"`

	// Indicates whether or not to automatically give permission to the teacher of a managed course on the Classroom app to
	// view a student's screen without prompting when the device is in supervised mode.
	ClassroomAppForceUnpromptedScreenObservation *bool `json:"classroomAppForceUnpromptedScreenObservation,omitempty"`

	// Possible values of the compliance app list.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`

	// List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection
	// can contain a maximum of 10000 elements.
	CompliantAppsList *[]AppListItem `json:"compliantAppsList,omitempty"`

	// Indicates whether or not to block the user from installing configuration profiles and certificates interactively when
	// the device is in supervised mode.
	ConfigurationProfileBlockChanges *bool `json:"configurationProfileBlockChanges,omitempty"`

	// Indicates whether or not to block definition lookup when the device is in supervised mode (iOS 8.1.3 and later ).
	DefinitionLookupBlocked *bool `json:"definitionLookupBlocked,omitempty"`

	// Indicates whether or not to allow the user to enables restrictions in the device settings when the device is in
	// supervised mode.
	DeviceBlockEnableRestrictions *bool `json:"deviceBlockEnableRestrictions,omitempty"`

	// Indicates whether or not to allow the use of the 'Erase all content and settings' option on the device when the
	// device is in supervised mode.
	DeviceBlockEraseContentAndSettings *bool `json:"deviceBlockEraseContentAndSettings,omitempty"`

	// Indicates whether or not to allow device name modification when the device is in supervised mode (iOS 9.0 and later).
	DeviceBlockNameModification *bool `json:"deviceBlockNameModification,omitempty"`

	// Indicates whether or not to block diagnostic data submission.
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`

	// Indicates whether or not to allow diagnostics submission settings modification when the device is in supervised mode
	// (iOS 9.3.2 and later).
	DiagnosticDataBlockSubmissionModification *bool `json:"diagnosticDataBlockSubmissionModification,omitempty"`

	// Indicates whether or not to block the user from viewing managed documents in unmanaged apps.
	DocumentsBlockManagedDocumentsInUnmanagedApps *bool `json:"documentsBlockManagedDocumentsInUnmanagedApps,omitempty"`

	// Indicates whether or not to block the user from viewing unmanaged documents in managed apps.
	DocumentsBlockUnmanagedDocumentsInManagedApps *bool `json:"documentsBlockUnmanagedDocumentsInManagedApps,omitempty"`

	// An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
	EmailInDomainSuffixes *[]string `json:"emailInDomainSuffixes,omitempty"`

	// Indicates whether or not to block the user from trusting an enterprise app.
	EnterpriseAppBlockTrust *bool `json:"enterpriseAppBlockTrust,omitempty"`

	// [Deprecated] Configuring this setting and setting the value to 'true' has no effect on the device.
	EnterpriseAppBlockTrustModification *bool `json:"enterpriseAppBlockTrustModification,omitempty"`

	// Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
	FaceTimeBlocked *bool `json:"faceTimeBlocked,omitempty"`

	// Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
	FindMyFriendsBlocked *bool `json:"findMyFriendsBlocked,omitempty"`

	// Indicates whether or not to block the user from using Game Center when the device is in supervised mode.
	GameCenterBlocked *bool `json:"gameCenterBlocked,omitempty"`

	// Indicates whether or not to block the user from having friends in Game Center. Requires a supervised device for iOS
	// 13 and later.
	GamingBlockGameCenterFriends *bool `json:"gamingBlockGameCenterFriends,omitempty"`

	// Indicates whether or not to block the user from using multiplayer gaming. Requires a supervised device for iOS 13 and
	// later.
	GamingBlockMultiplayer *bool `json:"gamingBlockMultiplayer,omitempty"`

	// indicates whether or not to allow host pairing to control the devices an iOS device can pair with when the iOS device
	// is in supervised mode.
	HostPairingBlocked *bool `json:"hostPairingBlocked,omitempty"`

	// Indicates whether or not to block the user from downloading media from the iBookstore that has been tagged as
	// erotica.
	IBooksStoreBlockErotica *bool `json:"iBooksStoreBlockErotica,omitempty"`

	// Indicates whether or not to block the user from using the iBooks Store when the device is in supervised mode.
	IBooksStoreBlocked *bool `json:"iBooksStoreBlocked,omitempty"`

	// Indicates whether or not to block the user from continuing work they started on iOS device to another iOS or macOS
	// device.
	ICloudBlockActivityContinuation *bool `json:"iCloudBlockActivityContinuation,omitempty"`

	// Indicates whether or not to block iCloud backup. Requires a supervised device for iOS 13 and later.
	ICloudBlockBackup *bool `json:"iCloudBlockBackup,omitempty"`

	// Indicates whether or not to block iCloud document sync. Requires a supervised device for iOS 13 and later.
	ICloudBlockDocumentSync *bool `json:"iCloudBlockDocumentSync,omitempty"`

	// Indicates whether or not to block Managed Apps Cloud Sync.
	ICloudBlockManagedAppsSync *bool `json:"iCloudBlockManagedAppsSync,omitempty"`

	// Indicates whether or not to block iCloud Photo Library.
	ICloudBlockPhotoLibrary *bool `json:"iCloudBlockPhotoLibrary,omitempty"`

	// Indicates whether or not to block iCloud Photo Stream Sync.
	ICloudBlockPhotoStreamSync *bool `json:"iCloudBlockPhotoStreamSync,omitempty"`

	// Indicates whether or not to block Shared Photo Stream.
	ICloudBlockSharedPhotoStream *bool `json:"iCloudBlockSharedPhotoStream,omitempty"`

	// Indicates whether or not to require backups to iCloud be encrypted.
	ICloudRequireEncryptedBackup *bool `json:"iCloudRequireEncryptedBackup,omitempty"`

	// Indicates whether or not to block the user from accessing explicit content in iTunes and the App Store. Requires a
	// supervised device for iOS 13 and later.
	ITunesBlockExplicitContent *bool `json:"iTunesBlockExplicitContent,omitempty"`

	// Indicates whether or not to block Music service and revert Music app to classic mode when the device is in supervised
	// mode (iOS 9.3 and later and macOS 10.12 and later).
	ITunesBlockMusicService *bool `json:"iTunesBlockMusicService,omitempty"`

	// Indicates whether or not to block the user from using iTunes Radio when the device is in supervised mode (iOS 9.3 and
	// later).
	ITunesBlockRadio *bool `json:"iTunesBlockRadio,omitempty"`

	// Indicates whether or not to block keyboard auto-correction when the device is in supervised mode (iOS 8.1.3 and
	// later).
	KeyboardBlockAutoCorrect *bool `json:"keyboardBlockAutoCorrect,omitempty"`

	// Indicates whether or not to block the user from using dictation input when the device is in supervised mode.
	KeyboardBlockDictation *bool `json:"keyboardBlockDictation,omitempty"`

	// Indicates whether or not to block predictive keyboards when device is in supervised mode (iOS 8.1.3 and later).
	KeyboardBlockPredictive *bool `json:"keyboardBlockPredictive,omitempty"`

	// Indicates whether or not to block keyboard shortcuts when the device is in supervised mode (iOS 9.0 and later).
	KeyboardBlockShortcuts *bool `json:"keyboardBlockShortcuts,omitempty"`

	// Indicates whether or not to block keyboard spell-checking when the device is in supervised mode (iOS 8.1.3 and
	// later).
	KeyboardBlockSpellCheck *bool `json:"keyboardBlockSpellCheck,omitempty"`

	// Indicates whether or not to allow assistive speak while in kiosk mode.
	KioskModeAllowAssistiveSpeak *bool `json:"kioskModeAllowAssistiveSpeak,omitempty"`

	// Indicates whether or not to allow access to the Assistive Touch Settings while in kiosk mode.
	KioskModeAllowAssistiveTouchSettings *bool `json:"kioskModeAllowAssistiveTouchSettings,omitempty"`

	// Indicates whether or not to allow device auto lock while in kiosk mode. This property's functionality is redundant
	// with the OS default and is deprecated. Use KioskModeBlockAutoLock instead.
	KioskModeAllowAutoLock *bool `json:"kioskModeAllowAutoLock,omitempty"`

	// Indicates whether or not to allow access to the Color Inversion Settings while in kiosk mode.
	KioskModeAllowColorInversionSettings *bool `json:"kioskModeAllowColorInversionSettings,omitempty"`

	// Indicates whether or not to allow use of the ringer switch while in kiosk mode. This property's functionality is
	// redundant with the OS default and is deprecated. Use KioskModeBlockRingerSwitch instead.
	KioskModeAllowRingerSwitch *bool `json:"kioskModeAllowRingerSwitch,omitempty"`

	// Indicates whether or not to allow screen rotation while in kiosk mode. This property's functionality is redundant
	// with the OS default and is deprecated. Use KioskModeBlockScreenRotation instead.
	KioskModeAllowScreenRotation *bool `json:"kioskModeAllowScreenRotation,omitempty"`

	// Indicates whether or not to allow use of the sleep button while in kiosk mode. This property's functionality is
	// redundant with the OS default and is deprecated. Use KioskModeBlockSleepButton instead.
	KioskModeAllowSleepButton *bool `json:"kioskModeAllowSleepButton,omitempty"`

	// Indicates whether or not to allow use of the touchscreen while in kiosk mode. This property's functionality is
	// redundant with the OS default and is deprecated. Use KioskModeBlockTouchscreen instead.
	KioskModeAllowTouchscreen *bool `json:"kioskModeAllowTouchscreen,omitempty"`

	// Indicates whether or not to allow access to the voice over settings while in kiosk mode.
	KioskModeAllowVoiceOverSettings *bool `json:"kioskModeAllowVoiceOverSettings,omitempty"`

	// Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is
	// redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
	KioskModeAllowVolumeButtons *bool `json:"kioskModeAllowVolumeButtons,omitempty"`

	// Indicates whether or not to allow access to the zoom settings while in kiosk mode.
	KioskModeAllowZoomSettings *bool `json:"kioskModeAllowZoomSettings,omitempty"`

	// URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
	KioskModeAppStoreUrl nullable.Type[string] `json:"kioskModeAppStoreUrl,omitempty"`

	// ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
	KioskModeBuiltInAppId nullable.Type[string] `json:"kioskModeBuiltInAppId,omitempty"`

	// Managed app id of the app to use for kiosk mode. If KioskModeManagedAppId is specified then KioskModeAppStoreUrl will
	// be ignored.
	KioskModeManagedAppId nullable.Type[string] `json:"kioskModeManagedAppId,omitempty"`

	// Indicates whether or not to require assistive touch while in kiosk mode.
	KioskModeRequireAssistiveTouch *bool `json:"kioskModeRequireAssistiveTouch,omitempty"`

	// Indicates whether or not to require color inversion while in kiosk mode.
	KioskModeRequireColorInversion *bool `json:"kioskModeRequireColorInversion,omitempty"`

	// Indicates whether or not to require mono audio while in kiosk mode.
	KioskModeRequireMonoAudio *bool `json:"kioskModeRequireMonoAudio,omitempty"`

	// Indicates whether or not to require voice over while in kiosk mode.
	KioskModeRequireVoiceOver *bool `json:"kioskModeRequireVoiceOver,omitempty"`

	// Indicates whether or not to require zoom while in kiosk mode.
	KioskModeRequireZoom *bool `json:"kioskModeRequireZoom,omitempty"`

	// Indicates whether or not to block the user from using control center on the lock screen.
	LockScreenBlockControlCenter *bool `json:"lockScreenBlockControlCenter,omitempty"`

	// Indicates whether or not to block the user from using the notification view on the lock screen.
	LockScreenBlockNotificationView *bool `json:"lockScreenBlockNotificationView,omitempty"`

	// Indicates whether or not to block the user from using passbook when the device is locked.
	LockScreenBlockPassbook *bool `json:"lockScreenBlockPassbook,omitempty"`

	// Indicates whether or not to block the user from using the Today View on the lock screen.
	LockScreenBlockTodayView *bool `json:"lockScreenBlockTodayView,omitempty"`

	// Apps rating as in media content
	MediaContentRatingApps *RatingAppsType `json:"mediaContentRatingApps,omitempty"`

	// Media content rating settings for Australia
	MediaContentRatingAustralia *MediaContentRatingAustralia `json:"mediaContentRatingAustralia,omitempty"`

	// Media content rating settings for Canada
	MediaContentRatingCanada *MediaContentRatingCanada `json:"mediaContentRatingCanada,omitempty"`

	// Media content rating settings for France
	MediaContentRatingFrance *MediaContentRatingFrance `json:"mediaContentRatingFrance,omitempty"`

	// Media content rating settings for Germany
	MediaContentRatingGermany *MediaContentRatingGermany `json:"mediaContentRatingGermany,omitempty"`

	// Media content rating settings for Ireland
	MediaContentRatingIreland *MediaContentRatingIreland `json:"mediaContentRatingIreland,omitempty"`

	// Media content rating settings for Japan
	MediaContentRatingJapan *MediaContentRatingJapan `json:"mediaContentRatingJapan,omitempty"`

	// Media content rating settings for New Zealand
	MediaContentRatingNewZealand *MediaContentRatingNewZealand `json:"mediaContentRatingNewZealand,omitempty"`

	// Media content rating settings for United Kingdom
	MediaContentRatingUnitedKingdom *MediaContentRatingUnitedKingdom `json:"mediaContentRatingUnitedKingdom,omitempty"`

	// Media content rating settings for United States
	MediaContentRatingUnitedStates *MediaContentRatingUnitedStates `json:"mediaContentRatingUnitedStates,omitempty"`

	// Indicates whether or not to block the user from using the Messages app on the supervised device.
	MessagesBlocked *bool `json:"messagesBlocked,omitempty"`

	// List of managed apps and the network rules that applies to them. This collection can contain a maximum of 1000
	// elements.
	NetworkUsageRules *[]IosNetworkUsageRule `json:"networkUsageRules,omitempty"`

	// Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
	NotificationsBlockSettingsModification *bool `json:"notificationsBlockSettingsModification,omitempty"`

	// Block modification of registered Touch ID fingerprints when in supervised mode.
	PasscodeBlockFingerprintModification *bool `json:"passcodeBlockFingerprintModification,omitempty"`

	// Indicates whether or not to block fingerprint unlock.
	PasscodeBlockFingerprintUnlock *bool `json:"passcodeBlockFingerprintUnlock,omitempty"`

	// Indicates whether or not to allow passcode modification on the supervised device (iOS 9.0 and later).
	PasscodeBlockModification *bool `json:"passcodeBlockModification,omitempty"`

	// Indicates whether or not to block simple passcodes.
	PasscodeBlockSimple *bool `json:"passcodeBlockSimple,omitempty"`

	// Number of days before the passcode expires. Valid values 1 to 65535
	PasscodeExpirationDays nullable.Type[int64] `json:"passcodeExpirationDays,omitempty"`

	// Number of character sets a passcode must contain. Valid values 0 to 4
	PasscodeMinimumCharacterSetCount nullable.Type[int64] `json:"passcodeMinimumCharacterSetCount,omitempty"`

	// Minimum length of passcode. Valid values 4 to 14
	PasscodeMinimumLength nullable.Type[int64] `json:"passcodeMinimumLength,omitempty"`

	// Minutes of inactivity before a passcode is required.
	PasscodeMinutesOfInactivityBeforeLock nullable.Type[int64] `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`

	// Minutes of inactivity before the screen times out.
	PasscodeMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Number of previous passcodes to block. Valid values 1 to 24
	PasscodePreviousPasscodeBlockCount nullable.Type[int64] `json:"passcodePreviousPasscodeBlockCount,omitempty"`

	// Indicates whether or not to require a passcode.
	PasscodeRequired *bool `json:"passcodeRequired,omitempty"`

	// Possible values of required passwords.
	PasscodeRequiredType *RequiredPasswordType `json:"passcodeRequiredType,omitempty"`

	// Number of sign in failures allowed before wiping the device. Valid values 2 to 11
	PasscodeSignInFailureCountBeforeWipe nullable.Type[int64] `json:"passcodeSignInFailureCountBeforeWipe,omitempty"`

	// Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
	PodcastsBlocked *bool `json:"podcastsBlocked,omitempty"`

	// Indicates whether or not to block the user from using Auto fill in Safari. Requires a supervised device for iOS 13
	// and later.
	SafariBlockAutofill *bool `json:"safariBlockAutofill,omitempty"`

	// Indicates whether or not to block JavaScript in Safari.
	SafariBlockJavaScript *bool `json:"safariBlockJavaScript,omitempty"`

	// Indicates whether or not to block popups in Safari.
	SafariBlockPopups *bool `json:"safariBlockPopups,omitempty"`

	// Indicates whether or not to block the user from using Safari. Requires a supervised device for iOS 13 and later.
	SafariBlocked *bool `json:"safariBlocked,omitempty"`

	// Web Browser Cookie Settings.
	SafariCookieSettings *WebBrowserCookieSettings `json:"safariCookieSettings,omitempty"`

	// URLs matching the patterns listed here will be considered managed.
	SafariManagedDomains *[]string `json:"safariManagedDomains,omitempty"`

	// Users can save passwords in Safari only from URLs matching the patterns listed here. Applies to devices in supervised
	// mode (iOS 9.3 and later).
	SafariPasswordAutoFillDomains *[]string `json:"safariPasswordAutoFillDomains,omitempty"`

	// Indicates whether or not to require fraud warning in Safari.
	SafariRequireFraudWarning *bool `json:"safariRequireFraudWarning,omitempty"`

	// Indicates whether or not to block the user from taking Screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
	SiriBlockUserGeneratedContent *bool `json:"siriBlockUserGeneratedContent,omitempty"`

	// Indicates whether or not to block the user from using Siri.
	SiriBlocked *bool `json:"siriBlocked,omitempty"`

	// Indicates whether or not to block the user from using Siri when locked.
	SiriBlockedWhenLocked *bool `json:"siriBlockedWhenLocked,omitempty"`

	// Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
	SiriRequireProfanityFilter *bool `json:"siriRequireProfanityFilter,omitempty"`

	// Indicates whether or not to block Spotlight search from returning internet results on supervised device.
	SpotlightBlockInternetResults *bool `json:"spotlightBlockInternetResults,omitempty"`

	// Indicates whether or not to block voice dialing.
	VoiceDialingBlocked *bool `json:"voiceDialingBlocked,omitempty"`

	// Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
	WallpaperBlockModification *bool `json:"wallpaperBlockModification,omitempty"`

	// Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device
	// is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+
	// should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
	WiFiConnectOnlyToConfiguredNetworks *bool `json:"wiFiConnectOnlyToConfiguredNetworks,omitempty"`

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

func (s IosGeneralDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s IosGeneralDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosGeneralDeviceConfiguration{}

func (s IosGeneralDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper IosGeneralDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosGeneralDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosGeneralDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosGeneralDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosGeneralDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}
