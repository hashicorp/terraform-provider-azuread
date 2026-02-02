package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
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

	// Indicates whether or not keychain storage of username and password for Airprint is blocked (iOS 11.0 and later).
	AirPrintBlockCredentialsStorage *bool `json:"airPrintBlockCredentialsStorage,omitempty"`

	// Indicates whether or not AirPrint is blocked (iOS 11.0 and later).
	AirPrintBlocked *bool `json:"airPrintBlocked,omitempty"`

	// Indicates whether or not iBeacon discovery of AirPrint printers is blocked. This prevents spurious AirPrint Bluetooth
	// beacons from phishing for network traffic (iOS 11.0 and later).
	AirPrintBlockiBeaconDiscovery *bool `json:"airPrintBlockiBeaconDiscovery,omitempty"`

	// Indicates if trusted certificates are required for TLS printing communication (iOS 11.0 and later).
	AirPrintForceTrustedTLS *bool `json:"airPrintForceTrustedTLS,omitempty"`

	// Prevents a user from adding any App Clips and removes any existing App Clips on the device.
	AppClipsBlocked *bool `json:"appClipsBlocked,omitempty"`

	// Indicates if the removal of apps is allowed.
	AppRemovalBlocked *bool `json:"appRemovalBlocked,omitempty"`

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

	// Limits Apple personalized advertising when true. Available in iOS 14 and later.
	ApplePersonalizedAdsBlocked *bool `json:"applePersonalizedAdsBlocked,omitempty"`

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

	// Indicates whether or not to force user authentication before autofilling passwords and credit card information in
	// Safari and other apps on supervised devices.
	AutoFillForceAuthentication *bool `json:"autoFillForceAuthentication,omitempty"`

	// Blocks users from unlocking their device with Apple Watch. Available for devices running iOS and iPadOS versions 14.5
	// and later.
	AutoUnlockBlocked *bool `json:"autoUnlockBlocked,omitempty"`

	// Indicates whether or not the removal of system apps from the device is blocked on a supervised device (iOS 11.0 and
	// later).
	BlockSystemAppRemoval *bool `json:"blockSystemAppRemoval,omitempty"`

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

	// Indicates whether or not to block the user from modifying the personal hotspot setting (iOS 12.2 or later).
	CellularBlockPersonalHotspotModification *bool `json:"cellularBlockPersonalHotspotModification,omitempty"`

	// Indicates whether or not to allow users to change the settings of the cellular plan on a supervised device.
	CellularBlockPlanModification *bool `json:"cellularBlockPlanModification,omitempty"`

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

	// Indicates whether or not to automatically give permission to the teacher's requests, without prompting the student,
	// when the device is in supervised mode.
	ClassroomForceAutomaticallyJoinClasses *bool `json:"classroomForceAutomaticallyJoinClasses,omitempty"`

	// Indicates whether a student enrolled in an unmanaged course via Classroom will request permission from the teacher
	// when attempting to leave the course (iOS 11.3 and later).
	ClassroomForceRequestPermissionToLeaveClasses *bool `json:"classroomForceRequestPermissionToLeaveClasses,omitempty"`

	// Indicates whether or not to allow the teacher to lock apps or the device without prompting the student. Supervised
	// only.
	ClassroomForceUnpromptedAppAndDeviceLock *bool `json:"classroomForceUnpromptedAppAndDeviceLock,omitempty"`

	// Possible values of the compliance app list.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`

	// List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection
	// can contain a maximum of 10000 elements.
	CompliantAppsList *[]AppListItem `json:"compliantAppsList,omitempty"`

	// Indicates whether or not to block the user from installing configuration profiles and certificates interactively when
	// the device is in supervised mode.
	ConfigurationProfileBlockChanges *bool `json:"configurationProfileBlockChanges,omitempty"`

	// Indicates whether or not managed apps can write contacts to unmanaged contacts accounts (iOS 12.0 and later).
	ContactsAllowManagedToUnmanagedWrite *bool `json:"contactsAllowManagedToUnmanagedWrite,omitempty"`

	// Indicates whether or not unmanaged apps can read from managed contacts accounts (iOS 12.0 or later).
	ContactsAllowUnmanagedToManagedRead *bool `json:"contactsAllowUnmanagedToManagedRead,omitempty"`

	// Indicates whether or not to block the continuous path keyboard when the device is supervised (iOS 13 or later).
	ContinuousPathKeyboardBlocked *bool `json:"continuousPathKeyboardBlocked,omitempty"`

	// Indicates whether or not the Date and Time 'Set Automatically' feature is enabled and cannot be turned off by the
	// user (iOS 12.0 and later).
	DateAndTimeForceSetAutomatically *bool `json:"dateAndTimeForceSetAutomatically,omitempty"`

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

	// Indicates whether or not Enterprise book back up is blocked.
	EnterpriseBookBlockBackup *bool `json:"enterpriseBookBlockBackup,omitempty"`

	// Indicates whether or not Enterprise book notes and highlights sync is blocked.
	EnterpriseBookBlockMetadataSync *bool `json:"enterpriseBookBlockMetadataSync,omitempty"`

	// Indicates whether or not to allow the addition or removal of cellular plans on the eSIM of a supervised device.
	EsimBlockModification *bool `json:"esimBlockModification,omitempty"`

	// Indicates whether or not to block the user from using FaceTime. Requires a supervised device for iOS 13 and later.
	FaceTimeBlocked *bool `json:"faceTimeBlocked,omitempty"`

	// Indicates if devices can access files or other resources on a network server using the Server Message Block (SMB)
	// protocol. Available for devices running iOS and iPadOS, versions 13.0 and later.
	FilesNetworkDriveAccessBlocked *bool `json:"filesNetworkDriveAccessBlocked,omitempty"`

	// Indicates if sevices with access can connect to and open files on a USB drive. Available for devices running iOS and
	// iPadOS, versions 13.0 and later.
	FilesUsbDriveAccessBlocked *bool `json:"filesUsbDriveAccessBlocked,omitempty"`

	// Indicates whether or not to block Find My Device when the device is supervised (iOS 13 or later).
	FindMyDeviceInFindMyAppBlocked *bool `json:"findMyDeviceInFindMyAppBlocked,omitempty"`

	// Indicates whether or not to block changes to Find My Friends when the device is in supervised mode.
	FindMyFriendsBlocked *bool `json:"findMyFriendsBlocked,omitempty"`

	// Indicates whether or not to block Find My Friends when the device is supervised (iOS 13 or later).
	FindMyFriendsInFindMyAppBlocked *bool `json:"findMyFriendsInFindMyAppBlocked,omitempty"`

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

	// iCloud private relay is an iCloud+ service that prevents networks and servers from monitoring a person's activity
	// across the internet. By blocking iCloud private relay, Apple will not encrypt the traffic leaving the device.
	// Available for devices running iOS 15 and later.
	ICloudPrivateRelayBlocked *bool `json:"iCloudPrivateRelayBlocked,omitempty"`

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

	// Indicates whether or not to block the iTunes app. Requires a supervised device for iOS 13 and later.
	ITunesBlocked *bool `json:"iTunesBlocked,omitempty"`

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

	// Indicates whether or not iCloud keychain synchronization is blocked. Requires a supervised device for iOS 13 and
	// later.
	KeychainBlockCloudSync *bool `json:"keychainBlockCloudSync,omitempty"`

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

	// Indicates whether or not to allow the user to toggle voice control in kiosk mode.
	KioskModeAllowVoiceControlModification *bool `json:"kioskModeAllowVoiceControlModification,omitempty"`

	// Indicates whether or not to allow access to the voice over settings while in kiosk mode.
	KioskModeAllowVoiceOverSettings *bool `json:"kioskModeAllowVoiceOverSettings,omitempty"`

	// Indicates whether or not to allow use of the volume buttons while in kiosk mode. This property's functionality is
	// redundant with the OS default and is deprecated. Use KioskModeBlockVolumeButtons instead.
	KioskModeAllowVolumeButtons *bool `json:"kioskModeAllowVolumeButtons,omitempty"`

	// Indicates whether or not to allow access to the zoom settings while in kiosk mode.
	KioskModeAllowZoomSettings *bool `json:"kioskModeAllowZoomSettings,omitempty"`

	// URL in the app store to the app to use for kiosk mode. Use if KioskModeManagedAppId is not known.
	KioskModeAppStoreUrl nullable.Type[string] `json:"kioskModeAppStoreUrl,omitempty"`

	// App source options for iOS kiosk mode.
	KioskModeAppType *IosKioskModeAppType `json:"kioskModeAppType,omitempty"`

	// Indicates whether or not to block device auto lock while in kiosk mode.
	KioskModeBlockAutoLock *bool `json:"kioskModeBlockAutoLock,omitempty"`

	// Indicates whether or not to block use of the ringer switch while in kiosk mode.
	KioskModeBlockRingerSwitch *bool `json:"kioskModeBlockRingerSwitch,omitempty"`

	// Indicates whether or not to block screen rotation while in kiosk mode.
	KioskModeBlockScreenRotation *bool `json:"kioskModeBlockScreenRotation,omitempty"`

	// Indicates whether or not to block use of the sleep button while in kiosk mode.
	KioskModeBlockSleepButton *bool `json:"kioskModeBlockSleepButton,omitempty"`

	// Indicates whether or not to block use of the touchscreen while in kiosk mode.
	KioskModeBlockTouchscreen *bool `json:"kioskModeBlockTouchscreen,omitempty"`

	// Indicates whether or not to block the volume buttons while in Kiosk Mode.
	KioskModeBlockVolumeButtons *bool `json:"kioskModeBlockVolumeButtons,omitempty"`

	// ID for built-in apps to use for kiosk mode. Used when KioskModeManagedAppId and KioskModeAppStoreUrl are not set.
	KioskModeBuiltInAppId nullable.Type[string] `json:"kioskModeBuiltInAppId,omitempty"`

	// Indicates whether or not to enable voice control in kiosk mode.
	KioskModeEnableVoiceControl *bool `json:"kioskModeEnableVoiceControl,omitempty"`

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

	// Open-in management controls how people share data between unmanaged and managed apps. Setting this to true enforces
	// copy/paste restrictions based on how you configured Block viewing corporate documents in unmanaged apps and Block
	// viewing non-corporate documents in corporate apps.
	ManagedPasteboardRequired *bool `json:"managedPasteboardRequired,omitempty"`

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

	// Disable NFC to prevent devices from pairing with other NFC-enabled devices. Available for iOS/iPadOS devices running
	// 14.2 and later.
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`

	// Indicates whether or not to allow notifications settings modification (iOS 9.3 and later).
	NotificationsBlockSettingsModification *bool `json:"notificationsBlockSettingsModification,omitempty"`

	// Disables connections to Siri servers so that users can’t use Siri to dictate text. Available for devices running
	// iOS and iPadOS versions 14.5 and later.
	OnDeviceOnlyDictationForced *bool `json:"onDeviceOnlyDictationForced,omitempty"`

	// When set to TRUE, the setting disables connections to Siri servers so that users can’t use Siri to translate text.
	// When set to FALSE, the setting allows connections to to Siri servers to users can use Siri to translate text.
	// Available for devices running iOS and iPadOS versions 15.0 and later.
	OnDeviceOnlyTranslationForced *bool `json:"onDeviceOnlyTranslationForced,omitempty"`

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

	// Indicates whether or not to block sharing passwords with the AirDrop passwords feature iOS 12.0 and later).
	PasswordBlockAirDropSharing *bool `json:"passwordBlockAirDropSharing,omitempty"`

	// Indicates if the AutoFill passwords feature is allowed (iOS 12.0 and later).
	PasswordBlockAutoFill *bool `json:"passwordBlockAutoFill,omitempty"`

	// Indicates whether or not to block requesting passwords from nearby devices (iOS 12.0 and later).
	PasswordBlockProximityRequests *bool `json:"passwordBlockProximityRequests,omitempty"`

	// Indicates whether or not over-the-air PKI updates are blocked. Setting this restriction to false does not disable CRL
	// and OCSP checks (iOS 7.0 and later).
	PkiBlockOTAUpdates *bool `json:"pkiBlockOTAUpdates,omitempty"`

	// Indicates whether or not to block the user from using podcasts on the supervised device (iOS 8.0 and later).
	PodcastsBlocked *bool `json:"podcastsBlocked,omitempty"`

	// Indicates if ad tracking is limited.(iOS 7.0 and later).
	PrivacyForceLimitAdTracking *bool `json:"privacyForceLimitAdTracking,omitempty"`

	// Indicates whether or not to enable the prompt to setup nearby devices with a supervised device.
	ProximityBlockSetupToNewDevice *bool `json:"proximityBlockSetupToNewDevice,omitempty"`

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

	// Indicates whether or not to block temporary sessions on Shared iPads (iOS 13.4 or later).
	SharedDeviceBlockTemporarySessions *bool `json:"sharedDeviceBlockTemporarySessions,omitempty"`

	// Indicates whether or not to block Siri from querying user-generated content when used on a supervised device.
	SiriBlockUserGeneratedContent *bool `json:"siriBlockUserGeneratedContent,omitempty"`

	// Indicates whether or not to block the user from using Siri.
	SiriBlocked *bool `json:"siriBlocked,omitempty"`

	// Indicates whether or not to block the user from using Siri when locked.
	SiriBlockedWhenLocked *bool `json:"siriBlockedWhenLocked,omitempty"`

	// Indicates whether or not to prevent Siri from dictating, or speaking profane language on supervised device.
	SiriRequireProfanityFilter *bool `json:"siriRequireProfanityFilter,omitempty"`

	// Sets how many days a software update will be delyed for a supervised device. Valid values 0 to 90
	SoftwareUpdatesEnforcedDelayInDays nullable.Type[int64] `json:"softwareUpdatesEnforcedDelayInDays,omitempty"`

	// Indicates whether or not to delay user visibility of software updates when the device is in supervised mode.
	SoftwareUpdatesForceDelayed *bool `json:"softwareUpdatesForceDelayed,omitempty"`

	// Indicates whether or not to block Spotlight search from returning internet results on supervised device.
	SpotlightBlockInternetResults *bool `json:"spotlightBlockInternetResults,omitempty"`

	// Allow users to boot devices into recovery mode with unpaired devices. Available for devices running iOS and iPadOS
	// versions 14.5 and later.
	UnpairedExternalBootToRecoveryAllowed *bool `json:"unpairedExternalBootToRecoveryAllowed,omitempty"`

	// Indicates if connecting to USB accessories while the device is locked is allowed (iOS 11.4.1 and later).
	UsbRestrictedModeBlocked *bool `json:"usbRestrictedModeBlocked,omitempty"`

	// Indicates whether or not to block voice dialing.
	VoiceDialingBlocked *bool `json:"voiceDialingBlocked,omitempty"`

	// Indicates whether or not the creation of VPN configurations is blocked (iOS 11.0 and later).
	VpnBlockCreation *bool `json:"vpnBlockCreation,omitempty"`

	// Indicates whether or not to allow wallpaper modification on supervised device (iOS 9.0 and later) .
	WallpaperBlockModification *bool `json:"wallpaperBlockModification,omitempty"`

	// Indicates whether or not to force the device to use only Wi-Fi networks from configuration profiles when the device
	// is in supervised mode. Available for devices running iOS and iPadOS versions 14.4 and earlier. Devices running 14.5+
	// should use the setting, 'WiFiConnectToAllowedNetworksOnlyForced.
	WiFiConnectOnlyToConfiguredNetworks *bool `json:"wiFiConnectOnlyToConfiguredNetworks,omitempty"`

	// Require devices to use Wi-Fi networks set up via configuration profiles. Available for devices running iOS and iPadOS
	// versions 14.5 and later.
	WiFiConnectToAllowedNetworksOnlyForced *bool `json:"wiFiConnectToAllowedNetworksOnlyForced,omitempty"`

	// Indicates whether or not Wi-Fi remains on, even when device is in airplane mode. Available for devices running iOS
	// and iPadOS, versions 13.0 and later.
	WifiPowerOnForced *bool `json:"wifiPowerOnForced,omitempty"`

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

func (s IosGeneralDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

var _ json.Unmarshaler = &IosGeneralDeviceConfiguration{}

func (s *IosGeneralDeviceConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountBlockModification                       *bool                                        `json:"accountBlockModification,omitempty"`
		ActivationLockAllowWhenSupervised              *bool                                        `json:"activationLockAllowWhenSupervised,omitempty"`
		AirDropBlocked                                 *bool                                        `json:"airDropBlocked,omitempty"`
		AirDropForceUnmanagedDropTarget                *bool                                        `json:"airDropForceUnmanagedDropTarget,omitempty"`
		AirPlayForcePairingPasswordForOutgoingRequests *bool                                        `json:"airPlayForcePairingPasswordForOutgoingRequests,omitempty"`
		AirPrintBlockCredentialsStorage                *bool                                        `json:"airPrintBlockCredentialsStorage,omitempty"`
		AirPrintBlocked                                *bool                                        `json:"airPrintBlocked,omitempty"`
		AirPrintBlockiBeaconDiscovery                  *bool                                        `json:"airPrintBlockiBeaconDiscovery,omitempty"`
		AirPrintForceTrustedTLS                        *bool                                        `json:"airPrintForceTrustedTLS,omitempty"`
		AppClipsBlocked                                *bool                                        `json:"appClipsBlocked,omitempty"`
		AppRemovalBlocked                              *bool                                        `json:"appRemovalBlocked,omitempty"`
		AppStoreBlockAutomaticDownloads                *bool                                        `json:"appStoreBlockAutomaticDownloads,omitempty"`
		AppStoreBlockInAppPurchases                    *bool                                        `json:"appStoreBlockInAppPurchases,omitempty"`
		AppStoreBlockUIAppInstallation                 *bool                                        `json:"appStoreBlockUIAppInstallation,omitempty"`
		AppStoreBlocked                                *bool                                        `json:"appStoreBlocked,omitempty"`
		AppStoreRequirePassword                        *bool                                        `json:"appStoreRequirePassword,omitempty"`
		AppleNewsBlocked                               *bool                                        `json:"appleNewsBlocked,omitempty"`
		ApplePersonalizedAdsBlocked                    *bool                                        `json:"applePersonalizedAdsBlocked,omitempty"`
		AppleWatchBlockPairing                         *bool                                        `json:"appleWatchBlockPairing,omitempty"`
		AppleWatchForceWristDetection                  *bool                                        `json:"appleWatchForceWristDetection,omitempty"`
		AppsVisibilityListType                         *AppListType                                 `json:"appsVisibilityListType,omitempty"`
		AutoFillForceAuthentication                    *bool                                        `json:"autoFillForceAuthentication,omitempty"`
		AutoUnlockBlocked                              *bool                                        `json:"autoUnlockBlocked,omitempty"`
		BlockSystemAppRemoval                          *bool                                        `json:"blockSystemAppRemoval,omitempty"`
		BluetoothBlockModification                     *bool                                        `json:"bluetoothBlockModification,omitempty"`
		CameraBlocked                                  *bool                                        `json:"cameraBlocked,omitempty"`
		CellularBlockDataRoaming                       *bool                                        `json:"cellularBlockDataRoaming,omitempty"`
		CellularBlockGlobalBackgroundFetchWhileRoaming *bool                                        `json:"cellularBlockGlobalBackgroundFetchWhileRoaming,omitempty"`
		CellularBlockPerAppDataModification            *bool                                        `json:"cellularBlockPerAppDataModification,omitempty"`
		CellularBlockPersonalHotspot                   *bool                                        `json:"cellularBlockPersonalHotspot,omitempty"`
		CellularBlockPersonalHotspotModification       *bool                                        `json:"cellularBlockPersonalHotspotModification,omitempty"`
		CellularBlockPlanModification                  *bool                                        `json:"cellularBlockPlanModification,omitempty"`
		CellularBlockVoiceRoaming                      *bool                                        `json:"cellularBlockVoiceRoaming,omitempty"`
		CertificatesBlockUntrustedTlsCertificates      *bool                                        `json:"certificatesBlockUntrustedTlsCertificates,omitempty"`
		ClassroomAppBlockRemoteScreenObservation       *bool                                        `json:"classroomAppBlockRemoteScreenObservation,omitempty"`
		ClassroomAppForceUnpromptedScreenObservation   *bool                                        `json:"classroomAppForceUnpromptedScreenObservation,omitempty"`
		ClassroomForceAutomaticallyJoinClasses         *bool                                        `json:"classroomForceAutomaticallyJoinClasses,omitempty"`
		ClassroomForceRequestPermissionToLeaveClasses  *bool                                        `json:"classroomForceRequestPermissionToLeaveClasses,omitempty"`
		ClassroomForceUnpromptedAppAndDeviceLock       *bool                                        `json:"classroomForceUnpromptedAppAndDeviceLock,omitempty"`
		CompliantAppListType                           *AppListType                                 `json:"compliantAppListType,omitempty"`
		ConfigurationProfileBlockChanges               *bool                                        `json:"configurationProfileBlockChanges,omitempty"`
		ContactsAllowManagedToUnmanagedWrite           *bool                                        `json:"contactsAllowManagedToUnmanagedWrite,omitempty"`
		ContactsAllowUnmanagedToManagedRead            *bool                                        `json:"contactsAllowUnmanagedToManagedRead,omitempty"`
		ContinuousPathKeyboardBlocked                  *bool                                        `json:"continuousPathKeyboardBlocked,omitempty"`
		DateAndTimeForceSetAutomatically               *bool                                        `json:"dateAndTimeForceSetAutomatically,omitempty"`
		DefinitionLookupBlocked                        *bool                                        `json:"definitionLookupBlocked,omitempty"`
		DeviceBlockEnableRestrictions                  *bool                                        `json:"deviceBlockEnableRestrictions,omitempty"`
		DeviceBlockEraseContentAndSettings             *bool                                        `json:"deviceBlockEraseContentAndSettings,omitempty"`
		DeviceBlockNameModification                    *bool                                        `json:"deviceBlockNameModification,omitempty"`
		DiagnosticDataBlockSubmission                  *bool                                        `json:"diagnosticDataBlockSubmission,omitempty"`
		DiagnosticDataBlockSubmissionModification      *bool                                        `json:"diagnosticDataBlockSubmissionModification,omitempty"`
		DocumentsBlockManagedDocumentsInUnmanagedApps  *bool                                        `json:"documentsBlockManagedDocumentsInUnmanagedApps,omitempty"`
		DocumentsBlockUnmanagedDocumentsInManagedApps  *bool                                        `json:"documentsBlockUnmanagedDocumentsInManagedApps,omitempty"`
		EmailInDomainSuffixes                          *[]string                                    `json:"emailInDomainSuffixes,omitempty"`
		EnterpriseAppBlockTrust                        *bool                                        `json:"enterpriseAppBlockTrust,omitempty"`
		EnterpriseAppBlockTrustModification            *bool                                        `json:"enterpriseAppBlockTrustModification,omitempty"`
		EnterpriseBookBlockBackup                      *bool                                        `json:"enterpriseBookBlockBackup,omitempty"`
		EnterpriseBookBlockMetadataSync                *bool                                        `json:"enterpriseBookBlockMetadataSync,omitempty"`
		EsimBlockModification                          *bool                                        `json:"esimBlockModification,omitempty"`
		FaceTimeBlocked                                *bool                                        `json:"faceTimeBlocked,omitempty"`
		FilesNetworkDriveAccessBlocked                 *bool                                        `json:"filesNetworkDriveAccessBlocked,omitempty"`
		FilesUsbDriveAccessBlocked                     *bool                                        `json:"filesUsbDriveAccessBlocked,omitempty"`
		FindMyDeviceInFindMyAppBlocked                 *bool                                        `json:"findMyDeviceInFindMyAppBlocked,omitempty"`
		FindMyFriendsBlocked                           *bool                                        `json:"findMyFriendsBlocked,omitempty"`
		FindMyFriendsInFindMyAppBlocked                *bool                                        `json:"findMyFriendsInFindMyAppBlocked,omitempty"`
		GameCenterBlocked                              *bool                                        `json:"gameCenterBlocked,omitempty"`
		GamingBlockGameCenterFriends                   *bool                                        `json:"gamingBlockGameCenterFriends,omitempty"`
		GamingBlockMultiplayer                         *bool                                        `json:"gamingBlockMultiplayer,omitempty"`
		HostPairingBlocked                             *bool                                        `json:"hostPairingBlocked,omitempty"`
		IBooksStoreBlockErotica                        *bool                                        `json:"iBooksStoreBlockErotica,omitempty"`
		IBooksStoreBlocked                             *bool                                        `json:"iBooksStoreBlocked,omitempty"`
		ICloudBlockActivityContinuation                *bool                                        `json:"iCloudBlockActivityContinuation,omitempty"`
		ICloudBlockBackup                              *bool                                        `json:"iCloudBlockBackup,omitempty"`
		ICloudBlockDocumentSync                        *bool                                        `json:"iCloudBlockDocumentSync,omitempty"`
		ICloudBlockManagedAppsSync                     *bool                                        `json:"iCloudBlockManagedAppsSync,omitempty"`
		ICloudBlockPhotoLibrary                        *bool                                        `json:"iCloudBlockPhotoLibrary,omitempty"`
		ICloudBlockPhotoStreamSync                     *bool                                        `json:"iCloudBlockPhotoStreamSync,omitempty"`
		ICloudBlockSharedPhotoStream                   *bool                                        `json:"iCloudBlockSharedPhotoStream,omitempty"`
		ICloudPrivateRelayBlocked                      *bool                                        `json:"iCloudPrivateRelayBlocked,omitempty"`
		ICloudRequireEncryptedBackup                   *bool                                        `json:"iCloudRequireEncryptedBackup,omitempty"`
		ITunesBlockExplicitContent                     *bool                                        `json:"iTunesBlockExplicitContent,omitempty"`
		ITunesBlockMusicService                        *bool                                        `json:"iTunesBlockMusicService,omitempty"`
		ITunesBlockRadio                               *bool                                        `json:"iTunesBlockRadio,omitempty"`
		ITunesBlocked                                  *bool                                        `json:"iTunesBlocked,omitempty"`
		KeyboardBlockAutoCorrect                       *bool                                        `json:"keyboardBlockAutoCorrect,omitempty"`
		KeyboardBlockDictation                         *bool                                        `json:"keyboardBlockDictation,omitempty"`
		KeyboardBlockPredictive                        *bool                                        `json:"keyboardBlockPredictive,omitempty"`
		KeyboardBlockShortcuts                         *bool                                        `json:"keyboardBlockShortcuts,omitempty"`
		KeyboardBlockSpellCheck                        *bool                                        `json:"keyboardBlockSpellCheck,omitempty"`
		KeychainBlockCloudSync                         *bool                                        `json:"keychainBlockCloudSync,omitempty"`
		KioskModeAllowAssistiveSpeak                   *bool                                        `json:"kioskModeAllowAssistiveSpeak,omitempty"`
		KioskModeAllowAssistiveTouchSettings           *bool                                        `json:"kioskModeAllowAssistiveTouchSettings,omitempty"`
		KioskModeAllowAutoLock                         *bool                                        `json:"kioskModeAllowAutoLock,omitempty"`
		KioskModeAllowColorInversionSettings           *bool                                        `json:"kioskModeAllowColorInversionSettings,omitempty"`
		KioskModeAllowRingerSwitch                     *bool                                        `json:"kioskModeAllowRingerSwitch,omitempty"`
		KioskModeAllowScreenRotation                   *bool                                        `json:"kioskModeAllowScreenRotation,omitempty"`
		KioskModeAllowSleepButton                      *bool                                        `json:"kioskModeAllowSleepButton,omitempty"`
		KioskModeAllowTouchscreen                      *bool                                        `json:"kioskModeAllowTouchscreen,omitempty"`
		KioskModeAllowVoiceControlModification         *bool                                        `json:"kioskModeAllowVoiceControlModification,omitempty"`
		KioskModeAllowVoiceOverSettings                *bool                                        `json:"kioskModeAllowVoiceOverSettings,omitempty"`
		KioskModeAllowVolumeButtons                    *bool                                        `json:"kioskModeAllowVolumeButtons,omitempty"`
		KioskModeAllowZoomSettings                     *bool                                        `json:"kioskModeAllowZoomSettings,omitempty"`
		KioskModeAppStoreUrl                           nullable.Type[string]                        `json:"kioskModeAppStoreUrl,omitempty"`
		KioskModeAppType                               *IosKioskModeAppType                         `json:"kioskModeAppType,omitempty"`
		KioskModeBlockAutoLock                         *bool                                        `json:"kioskModeBlockAutoLock,omitempty"`
		KioskModeBlockRingerSwitch                     *bool                                        `json:"kioskModeBlockRingerSwitch,omitempty"`
		KioskModeBlockScreenRotation                   *bool                                        `json:"kioskModeBlockScreenRotation,omitempty"`
		KioskModeBlockSleepButton                      *bool                                        `json:"kioskModeBlockSleepButton,omitempty"`
		KioskModeBlockTouchscreen                      *bool                                        `json:"kioskModeBlockTouchscreen,omitempty"`
		KioskModeBlockVolumeButtons                    *bool                                        `json:"kioskModeBlockVolumeButtons,omitempty"`
		KioskModeBuiltInAppId                          nullable.Type[string]                        `json:"kioskModeBuiltInAppId,omitempty"`
		KioskModeEnableVoiceControl                    *bool                                        `json:"kioskModeEnableVoiceControl,omitempty"`
		KioskModeManagedAppId                          nullable.Type[string]                        `json:"kioskModeManagedAppId,omitempty"`
		KioskModeRequireAssistiveTouch                 *bool                                        `json:"kioskModeRequireAssistiveTouch,omitempty"`
		KioskModeRequireColorInversion                 *bool                                        `json:"kioskModeRequireColorInversion,omitempty"`
		KioskModeRequireMonoAudio                      *bool                                        `json:"kioskModeRequireMonoAudio,omitempty"`
		KioskModeRequireVoiceOver                      *bool                                        `json:"kioskModeRequireVoiceOver,omitempty"`
		KioskModeRequireZoom                           *bool                                        `json:"kioskModeRequireZoom,omitempty"`
		LockScreenBlockControlCenter                   *bool                                        `json:"lockScreenBlockControlCenter,omitempty"`
		LockScreenBlockNotificationView                *bool                                        `json:"lockScreenBlockNotificationView,omitempty"`
		LockScreenBlockPassbook                        *bool                                        `json:"lockScreenBlockPassbook,omitempty"`
		LockScreenBlockTodayView                       *bool                                        `json:"lockScreenBlockTodayView,omitempty"`
		ManagedPasteboardRequired                      *bool                                        `json:"managedPasteboardRequired,omitempty"`
		MediaContentRatingApps                         *RatingAppsType                              `json:"mediaContentRatingApps,omitempty"`
		MediaContentRatingAustralia                    *MediaContentRatingAustralia                 `json:"mediaContentRatingAustralia,omitempty"`
		MediaContentRatingCanada                       *MediaContentRatingCanada                    `json:"mediaContentRatingCanada,omitempty"`
		MediaContentRatingFrance                       *MediaContentRatingFrance                    `json:"mediaContentRatingFrance,omitempty"`
		MediaContentRatingGermany                      *MediaContentRatingGermany                   `json:"mediaContentRatingGermany,omitempty"`
		MediaContentRatingIreland                      *MediaContentRatingIreland                   `json:"mediaContentRatingIreland,omitempty"`
		MediaContentRatingJapan                        *MediaContentRatingJapan                     `json:"mediaContentRatingJapan,omitempty"`
		MediaContentRatingNewZealand                   *MediaContentRatingNewZealand                `json:"mediaContentRatingNewZealand,omitempty"`
		MediaContentRatingUnitedKingdom                *MediaContentRatingUnitedKingdom             `json:"mediaContentRatingUnitedKingdom,omitempty"`
		MediaContentRatingUnitedStates                 *MediaContentRatingUnitedStates              `json:"mediaContentRatingUnitedStates,omitempty"`
		MessagesBlocked                                *bool                                        `json:"messagesBlocked,omitempty"`
		NetworkUsageRules                              *[]IosNetworkUsageRule                       `json:"networkUsageRules,omitempty"`
		NfcBlocked                                     *bool                                        `json:"nfcBlocked,omitempty"`
		NotificationsBlockSettingsModification         *bool                                        `json:"notificationsBlockSettingsModification,omitempty"`
		OnDeviceOnlyDictationForced                    *bool                                        `json:"onDeviceOnlyDictationForced,omitempty"`
		OnDeviceOnlyTranslationForced                  *bool                                        `json:"onDeviceOnlyTranslationForced,omitempty"`
		PasscodeBlockFingerprintModification           *bool                                        `json:"passcodeBlockFingerprintModification,omitempty"`
		PasscodeBlockFingerprintUnlock                 *bool                                        `json:"passcodeBlockFingerprintUnlock,omitempty"`
		PasscodeBlockModification                      *bool                                        `json:"passcodeBlockModification,omitempty"`
		PasscodeBlockSimple                            *bool                                        `json:"passcodeBlockSimple,omitempty"`
		PasscodeExpirationDays                         nullable.Type[int64]                         `json:"passcodeExpirationDays,omitempty"`
		PasscodeMinimumCharacterSetCount               nullable.Type[int64]                         `json:"passcodeMinimumCharacterSetCount,omitempty"`
		PasscodeMinimumLength                          nullable.Type[int64]                         `json:"passcodeMinimumLength,omitempty"`
		PasscodeMinutesOfInactivityBeforeLock          nullable.Type[int64]                         `json:"passcodeMinutesOfInactivityBeforeLock,omitempty"`
		PasscodeMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64]                         `json:"passcodeMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasscodePreviousPasscodeBlockCount             nullable.Type[int64]                         `json:"passcodePreviousPasscodeBlockCount,omitempty"`
		PasscodeRequired                               *bool                                        `json:"passcodeRequired,omitempty"`
		PasscodeRequiredType                           *RequiredPasswordType                        `json:"passcodeRequiredType,omitempty"`
		PasscodeSignInFailureCountBeforeWipe           nullable.Type[int64]                         `json:"passcodeSignInFailureCountBeforeWipe,omitempty"`
		PasswordBlockAirDropSharing                    *bool                                        `json:"passwordBlockAirDropSharing,omitempty"`
		PasswordBlockAutoFill                          *bool                                        `json:"passwordBlockAutoFill,omitempty"`
		PasswordBlockProximityRequests                 *bool                                        `json:"passwordBlockProximityRequests,omitempty"`
		PkiBlockOTAUpdates                             *bool                                        `json:"pkiBlockOTAUpdates,omitempty"`
		PodcastsBlocked                                *bool                                        `json:"podcastsBlocked,omitempty"`
		PrivacyForceLimitAdTracking                    *bool                                        `json:"privacyForceLimitAdTracking,omitempty"`
		ProximityBlockSetupToNewDevice                 *bool                                        `json:"proximityBlockSetupToNewDevice,omitempty"`
		SafariBlockAutofill                            *bool                                        `json:"safariBlockAutofill,omitempty"`
		SafariBlockJavaScript                          *bool                                        `json:"safariBlockJavaScript,omitempty"`
		SafariBlockPopups                              *bool                                        `json:"safariBlockPopups,omitempty"`
		SafariBlocked                                  *bool                                        `json:"safariBlocked,omitempty"`
		SafariCookieSettings                           *WebBrowserCookieSettings                    `json:"safariCookieSettings,omitempty"`
		SafariManagedDomains                           *[]string                                    `json:"safariManagedDomains,omitempty"`
		SafariPasswordAutoFillDomains                  *[]string                                    `json:"safariPasswordAutoFillDomains,omitempty"`
		SafariRequireFraudWarning                      *bool                                        `json:"safariRequireFraudWarning,omitempty"`
		ScreenCaptureBlocked                           *bool                                        `json:"screenCaptureBlocked,omitempty"`
		SharedDeviceBlockTemporarySessions             *bool                                        `json:"sharedDeviceBlockTemporarySessions,omitempty"`
		SiriBlockUserGeneratedContent                  *bool                                        `json:"siriBlockUserGeneratedContent,omitempty"`
		SiriBlocked                                    *bool                                        `json:"siriBlocked,omitempty"`
		SiriBlockedWhenLocked                          *bool                                        `json:"siriBlockedWhenLocked,omitempty"`
		SiriRequireProfanityFilter                     *bool                                        `json:"siriRequireProfanityFilter,omitempty"`
		SoftwareUpdatesEnforcedDelayInDays             nullable.Type[int64]                         `json:"softwareUpdatesEnforcedDelayInDays,omitempty"`
		SoftwareUpdatesForceDelayed                    *bool                                        `json:"softwareUpdatesForceDelayed,omitempty"`
		SpotlightBlockInternetResults                  *bool                                        `json:"spotlightBlockInternetResults,omitempty"`
		UnpairedExternalBootToRecoveryAllowed          *bool                                        `json:"unpairedExternalBootToRecoveryAllowed,omitempty"`
		UsbRestrictedModeBlocked                       *bool                                        `json:"usbRestrictedModeBlocked,omitempty"`
		VoiceDialingBlocked                            *bool                                        `json:"voiceDialingBlocked,omitempty"`
		VpnBlockCreation                               *bool                                        `json:"vpnBlockCreation,omitempty"`
		WallpaperBlockModification                     *bool                                        `json:"wallpaperBlockModification,omitempty"`
		WiFiConnectOnlyToConfiguredNetworks            *bool                                        `json:"wiFiConnectOnlyToConfiguredNetworks,omitempty"`
		WiFiConnectToAllowedNetworksOnlyForced         *bool                                        `json:"wiFiConnectToAllowedNetworksOnlyForced,omitempty"`
		WifiPowerOnForced                              *bool                                        `json:"wifiPowerOnForced,omitempty"`
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

	s.AccountBlockModification = decoded.AccountBlockModification
	s.ActivationLockAllowWhenSupervised = decoded.ActivationLockAllowWhenSupervised
	s.AirDropBlocked = decoded.AirDropBlocked
	s.AirDropForceUnmanagedDropTarget = decoded.AirDropForceUnmanagedDropTarget
	s.AirPlayForcePairingPasswordForOutgoingRequests = decoded.AirPlayForcePairingPasswordForOutgoingRequests
	s.AirPrintBlockCredentialsStorage = decoded.AirPrintBlockCredentialsStorage
	s.AirPrintBlocked = decoded.AirPrintBlocked
	s.AirPrintBlockiBeaconDiscovery = decoded.AirPrintBlockiBeaconDiscovery
	s.AirPrintForceTrustedTLS = decoded.AirPrintForceTrustedTLS
	s.AppClipsBlocked = decoded.AppClipsBlocked
	s.AppRemovalBlocked = decoded.AppRemovalBlocked
	s.AppStoreBlockAutomaticDownloads = decoded.AppStoreBlockAutomaticDownloads
	s.AppStoreBlockInAppPurchases = decoded.AppStoreBlockInAppPurchases
	s.AppStoreBlockUIAppInstallation = decoded.AppStoreBlockUIAppInstallation
	s.AppStoreBlocked = decoded.AppStoreBlocked
	s.AppStoreRequirePassword = decoded.AppStoreRequirePassword
	s.AppleNewsBlocked = decoded.AppleNewsBlocked
	s.ApplePersonalizedAdsBlocked = decoded.ApplePersonalizedAdsBlocked
	s.AppleWatchBlockPairing = decoded.AppleWatchBlockPairing
	s.AppleWatchForceWristDetection = decoded.AppleWatchForceWristDetection
	s.AppsVisibilityListType = decoded.AppsVisibilityListType
	s.AutoFillForceAuthentication = decoded.AutoFillForceAuthentication
	s.AutoUnlockBlocked = decoded.AutoUnlockBlocked
	s.BlockSystemAppRemoval = decoded.BlockSystemAppRemoval
	s.BluetoothBlockModification = decoded.BluetoothBlockModification
	s.CameraBlocked = decoded.CameraBlocked
	s.CellularBlockDataRoaming = decoded.CellularBlockDataRoaming
	s.CellularBlockGlobalBackgroundFetchWhileRoaming = decoded.CellularBlockGlobalBackgroundFetchWhileRoaming
	s.CellularBlockPerAppDataModification = decoded.CellularBlockPerAppDataModification
	s.CellularBlockPersonalHotspot = decoded.CellularBlockPersonalHotspot
	s.CellularBlockPersonalHotspotModification = decoded.CellularBlockPersonalHotspotModification
	s.CellularBlockPlanModification = decoded.CellularBlockPlanModification
	s.CellularBlockVoiceRoaming = decoded.CellularBlockVoiceRoaming
	s.CertificatesBlockUntrustedTlsCertificates = decoded.CertificatesBlockUntrustedTlsCertificates
	s.ClassroomAppBlockRemoteScreenObservation = decoded.ClassroomAppBlockRemoteScreenObservation
	s.ClassroomAppForceUnpromptedScreenObservation = decoded.ClassroomAppForceUnpromptedScreenObservation
	s.ClassroomForceAutomaticallyJoinClasses = decoded.ClassroomForceAutomaticallyJoinClasses
	s.ClassroomForceRequestPermissionToLeaveClasses = decoded.ClassroomForceRequestPermissionToLeaveClasses
	s.ClassroomForceUnpromptedAppAndDeviceLock = decoded.ClassroomForceUnpromptedAppAndDeviceLock
	s.CompliantAppListType = decoded.CompliantAppListType
	s.ConfigurationProfileBlockChanges = decoded.ConfigurationProfileBlockChanges
	s.ContactsAllowManagedToUnmanagedWrite = decoded.ContactsAllowManagedToUnmanagedWrite
	s.ContactsAllowUnmanagedToManagedRead = decoded.ContactsAllowUnmanagedToManagedRead
	s.ContinuousPathKeyboardBlocked = decoded.ContinuousPathKeyboardBlocked
	s.DateAndTimeForceSetAutomatically = decoded.DateAndTimeForceSetAutomatically
	s.DefinitionLookupBlocked = decoded.DefinitionLookupBlocked
	s.DeviceBlockEnableRestrictions = decoded.DeviceBlockEnableRestrictions
	s.DeviceBlockEraseContentAndSettings = decoded.DeviceBlockEraseContentAndSettings
	s.DeviceBlockNameModification = decoded.DeviceBlockNameModification
	s.DiagnosticDataBlockSubmission = decoded.DiagnosticDataBlockSubmission
	s.DiagnosticDataBlockSubmissionModification = decoded.DiagnosticDataBlockSubmissionModification
	s.DocumentsBlockManagedDocumentsInUnmanagedApps = decoded.DocumentsBlockManagedDocumentsInUnmanagedApps
	s.DocumentsBlockUnmanagedDocumentsInManagedApps = decoded.DocumentsBlockUnmanagedDocumentsInManagedApps
	s.EmailInDomainSuffixes = decoded.EmailInDomainSuffixes
	s.EnterpriseAppBlockTrust = decoded.EnterpriseAppBlockTrust
	s.EnterpriseAppBlockTrustModification = decoded.EnterpriseAppBlockTrustModification
	s.EnterpriseBookBlockBackup = decoded.EnterpriseBookBlockBackup
	s.EnterpriseBookBlockMetadataSync = decoded.EnterpriseBookBlockMetadataSync
	s.EsimBlockModification = decoded.EsimBlockModification
	s.FaceTimeBlocked = decoded.FaceTimeBlocked
	s.FilesNetworkDriveAccessBlocked = decoded.FilesNetworkDriveAccessBlocked
	s.FilesUsbDriveAccessBlocked = decoded.FilesUsbDriveAccessBlocked
	s.FindMyDeviceInFindMyAppBlocked = decoded.FindMyDeviceInFindMyAppBlocked
	s.FindMyFriendsBlocked = decoded.FindMyFriendsBlocked
	s.FindMyFriendsInFindMyAppBlocked = decoded.FindMyFriendsInFindMyAppBlocked
	s.GameCenterBlocked = decoded.GameCenterBlocked
	s.GamingBlockGameCenterFriends = decoded.GamingBlockGameCenterFriends
	s.GamingBlockMultiplayer = decoded.GamingBlockMultiplayer
	s.HostPairingBlocked = decoded.HostPairingBlocked
	s.IBooksStoreBlockErotica = decoded.IBooksStoreBlockErotica
	s.IBooksStoreBlocked = decoded.IBooksStoreBlocked
	s.ICloudBlockActivityContinuation = decoded.ICloudBlockActivityContinuation
	s.ICloudBlockBackup = decoded.ICloudBlockBackup
	s.ICloudBlockDocumentSync = decoded.ICloudBlockDocumentSync
	s.ICloudBlockManagedAppsSync = decoded.ICloudBlockManagedAppsSync
	s.ICloudBlockPhotoLibrary = decoded.ICloudBlockPhotoLibrary
	s.ICloudBlockPhotoStreamSync = decoded.ICloudBlockPhotoStreamSync
	s.ICloudBlockSharedPhotoStream = decoded.ICloudBlockSharedPhotoStream
	s.ICloudPrivateRelayBlocked = decoded.ICloudPrivateRelayBlocked
	s.ICloudRequireEncryptedBackup = decoded.ICloudRequireEncryptedBackup
	s.ITunesBlockExplicitContent = decoded.ITunesBlockExplicitContent
	s.ITunesBlockMusicService = decoded.ITunesBlockMusicService
	s.ITunesBlockRadio = decoded.ITunesBlockRadio
	s.ITunesBlocked = decoded.ITunesBlocked
	s.KeyboardBlockAutoCorrect = decoded.KeyboardBlockAutoCorrect
	s.KeyboardBlockDictation = decoded.KeyboardBlockDictation
	s.KeyboardBlockPredictive = decoded.KeyboardBlockPredictive
	s.KeyboardBlockShortcuts = decoded.KeyboardBlockShortcuts
	s.KeyboardBlockSpellCheck = decoded.KeyboardBlockSpellCheck
	s.KeychainBlockCloudSync = decoded.KeychainBlockCloudSync
	s.KioskModeAllowAssistiveSpeak = decoded.KioskModeAllowAssistiveSpeak
	s.KioskModeAllowAssistiveTouchSettings = decoded.KioskModeAllowAssistiveTouchSettings
	s.KioskModeAllowAutoLock = decoded.KioskModeAllowAutoLock
	s.KioskModeAllowColorInversionSettings = decoded.KioskModeAllowColorInversionSettings
	s.KioskModeAllowRingerSwitch = decoded.KioskModeAllowRingerSwitch
	s.KioskModeAllowScreenRotation = decoded.KioskModeAllowScreenRotation
	s.KioskModeAllowSleepButton = decoded.KioskModeAllowSleepButton
	s.KioskModeAllowTouchscreen = decoded.KioskModeAllowTouchscreen
	s.KioskModeAllowVoiceControlModification = decoded.KioskModeAllowVoiceControlModification
	s.KioskModeAllowVoiceOverSettings = decoded.KioskModeAllowVoiceOverSettings
	s.KioskModeAllowVolumeButtons = decoded.KioskModeAllowVolumeButtons
	s.KioskModeAllowZoomSettings = decoded.KioskModeAllowZoomSettings
	s.KioskModeAppStoreUrl = decoded.KioskModeAppStoreUrl
	s.KioskModeAppType = decoded.KioskModeAppType
	s.KioskModeBlockAutoLock = decoded.KioskModeBlockAutoLock
	s.KioskModeBlockRingerSwitch = decoded.KioskModeBlockRingerSwitch
	s.KioskModeBlockScreenRotation = decoded.KioskModeBlockScreenRotation
	s.KioskModeBlockSleepButton = decoded.KioskModeBlockSleepButton
	s.KioskModeBlockTouchscreen = decoded.KioskModeBlockTouchscreen
	s.KioskModeBlockVolumeButtons = decoded.KioskModeBlockVolumeButtons
	s.KioskModeBuiltInAppId = decoded.KioskModeBuiltInAppId
	s.KioskModeEnableVoiceControl = decoded.KioskModeEnableVoiceControl
	s.KioskModeManagedAppId = decoded.KioskModeManagedAppId
	s.KioskModeRequireAssistiveTouch = decoded.KioskModeRequireAssistiveTouch
	s.KioskModeRequireColorInversion = decoded.KioskModeRequireColorInversion
	s.KioskModeRequireMonoAudio = decoded.KioskModeRequireMonoAudio
	s.KioskModeRequireVoiceOver = decoded.KioskModeRequireVoiceOver
	s.KioskModeRequireZoom = decoded.KioskModeRequireZoom
	s.LockScreenBlockControlCenter = decoded.LockScreenBlockControlCenter
	s.LockScreenBlockNotificationView = decoded.LockScreenBlockNotificationView
	s.LockScreenBlockPassbook = decoded.LockScreenBlockPassbook
	s.LockScreenBlockTodayView = decoded.LockScreenBlockTodayView
	s.ManagedPasteboardRequired = decoded.ManagedPasteboardRequired
	s.MediaContentRatingApps = decoded.MediaContentRatingApps
	s.MediaContentRatingAustralia = decoded.MediaContentRatingAustralia
	s.MediaContentRatingCanada = decoded.MediaContentRatingCanada
	s.MediaContentRatingFrance = decoded.MediaContentRatingFrance
	s.MediaContentRatingGermany = decoded.MediaContentRatingGermany
	s.MediaContentRatingIreland = decoded.MediaContentRatingIreland
	s.MediaContentRatingJapan = decoded.MediaContentRatingJapan
	s.MediaContentRatingNewZealand = decoded.MediaContentRatingNewZealand
	s.MediaContentRatingUnitedKingdom = decoded.MediaContentRatingUnitedKingdom
	s.MediaContentRatingUnitedStates = decoded.MediaContentRatingUnitedStates
	s.MessagesBlocked = decoded.MessagesBlocked
	s.NetworkUsageRules = decoded.NetworkUsageRules
	s.NfcBlocked = decoded.NfcBlocked
	s.NotificationsBlockSettingsModification = decoded.NotificationsBlockSettingsModification
	s.OnDeviceOnlyDictationForced = decoded.OnDeviceOnlyDictationForced
	s.OnDeviceOnlyTranslationForced = decoded.OnDeviceOnlyTranslationForced
	s.PasscodeBlockFingerprintModification = decoded.PasscodeBlockFingerprintModification
	s.PasscodeBlockFingerprintUnlock = decoded.PasscodeBlockFingerprintUnlock
	s.PasscodeBlockModification = decoded.PasscodeBlockModification
	s.PasscodeBlockSimple = decoded.PasscodeBlockSimple
	s.PasscodeExpirationDays = decoded.PasscodeExpirationDays
	s.PasscodeMinimumCharacterSetCount = decoded.PasscodeMinimumCharacterSetCount
	s.PasscodeMinimumLength = decoded.PasscodeMinimumLength
	s.PasscodeMinutesOfInactivityBeforeLock = decoded.PasscodeMinutesOfInactivityBeforeLock
	s.PasscodeMinutesOfInactivityBeforeScreenTimeout = decoded.PasscodeMinutesOfInactivityBeforeScreenTimeout
	s.PasscodePreviousPasscodeBlockCount = decoded.PasscodePreviousPasscodeBlockCount
	s.PasscodeRequired = decoded.PasscodeRequired
	s.PasscodeRequiredType = decoded.PasscodeRequiredType
	s.PasscodeSignInFailureCountBeforeWipe = decoded.PasscodeSignInFailureCountBeforeWipe
	s.PasswordBlockAirDropSharing = decoded.PasswordBlockAirDropSharing
	s.PasswordBlockAutoFill = decoded.PasswordBlockAutoFill
	s.PasswordBlockProximityRequests = decoded.PasswordBlockProximityRequests
	s.PkiBlockOTAUpdates = decoded.PkiBlockOTAUpdates
	s.PodcastsBlocked = decoded.PodcastsBlocked
	s.PrivacyForceLimitAdTracking = decoded.PrivacyForceLimitAdTracking
	s.ProximityBlockSetupToNewDevice = decoded.ProximityBlockSetupToNewDevice
	s.SafariBlockAutofill = decoded.SafariBlockAutofill
	s.SafariBlockJavaScript = decoded.SafariBlockJavaScript
	s.SafariBlockPopups = decoded.SafariBlockPopups
	s.SafariBlocked = decoded.SafariBlocked
	s.SafariCookieSettings = decoded.SafariCookieSettings
	s.SafariManagedDomains = decoded.SafariManagedDomains
	s.SafariPasswordAutoFillDomains = decoded.SafariPasswordAutoFillDomains
	s.SafariRequireFraudWarning = decoded.SafariRequireFraudWarning
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.SharedDeviceBlockTemporarySessions = decoded.SharedDeviceBlockTemporarySessions
	s.SiriBlockUserGeneratedContent = decoded.SiriBlockUserGeneratedContent
	s.SiriBlocked = decoded.SiriBlocked
	s.SiriBlockedWhenLocked = decoded.SiriBlockedWhenLocked
	s.SiriRequireProfanityFilter = decoded.SiriRequireProfanityFilter
	s.SoftwareUpdatesEnforcedDelayInDays = decoded.SoftwareUpdatesEnforcedDelayInDays
	s.SoftwareUpdatesForceDelayed = decoded.SoftwareUpdatesForceDelayed
	s.SpotlightBlockInternetResults = decoded.SpotlightBlockInternetResults
	s.UnpairedExternalBootToRecoveryAllowed = decoded.UnpairedExternalBootToRecoveryAllowed
	s.UsbRestrictedModeBlocked = decoded.UsbRestrictedModeBlocked
	s.VoiceDialingBlocked = decoded.VoiceDialingBlocked
	s.VpnBlockCreation = decoded.VpnBlockCreation
	s.WallpaperBlockModification = decoded.WallpaperBlockModification
	s.WiFiConnectOnlyToConfiguredNetworks = decoded.WiFiConnectOnlyToConfiguredNetworks
	s.WiFiConnectToAllowedNetworksOnlyForced = decoded.WiFiConnectToAllowedNetworksOnlyForced
	s.WifiPowerOnForced = decoded.WifiPowerOnForced
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
		return fmt.Errorf("unmarshaling IosGeneralDeviceConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appsSingleAppModeList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppsSingleAppModeList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppsSingleAppModeList' for 'IosGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppsSingleAppModeList = &output
	}

	if v, ok := temp["appsVisibilityList"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AppsVisibilityList into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AppsVisibilityList' for 'IosGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AppsVisibilityList = &output
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
				return fmt.Errorf("unmarshaling index %d field 'CompliantAppsList' for 'IosGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CompliantAppsList = &output
	}

	return nil
}
