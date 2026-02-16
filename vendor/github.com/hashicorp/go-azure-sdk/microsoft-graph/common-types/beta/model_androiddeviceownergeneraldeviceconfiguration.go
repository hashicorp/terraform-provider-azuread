package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = AndroidDeviceOwnerGeneralDeviceConfiguration{}

type AndroidDeviceOwnerGeneralDeviceConfiguration struct {
	// Indicates whether or not adding or removing accounts is disabled.
	AccountsBlockModification nullable.Type[bool] `json:"accountsBlockModification,omitempty"`

	// Specifies the list of managed apps with app details and its associated delegated scope(s). This collection can
	// contain a maximum of 500 elements.
	AndroidDeviceOwnerDelegatedScopeAppSettings *[]AndroidDeviceOwnerDelegatedScopeAppSetting `json:"androidDeviceOwnerDelegatedScopeAppSettings,omitempty"`

	// Indicates whether or not the user is allowed to enable to unknown sources setting.
	AppsAllowInstallFromUnknownSources nullable.Type[bool] `json:"appsAllowInstallFromUnknownSources,omitempty"`

	// Indicates the value of the app auto update policy. Possible values are: notConfigured, userChoice, never, wiFiOnly,
	// always.
	AppsAutoUpdatePolicy *AndroidDeviceOwnerAppAutoUpdatePolicyType `json:"appsAutoUpdatePolicy,omitempty"`

	// Indicates the permission policy for requests for runtime permissions if one is not defined for the app specifically.
	// Possible values are: deviceDefault, prompt, autoGrant, autoDeny.
	AppsDefaultPermissionPolicy *AndroidDeviceOwnerDefaultAppPermissionPolicyType `json:"appsDefaultPermissionPolicy,omitempty"`

	// Whether or not to recommend all apps skip any first-time-use hints they may have added.
	AppsRecommendSkippingFirstUseHints nullable.Type[bool] `json:"appsRecommendSkippingFirstUseHints,omitempty"`

	// A list of managed apps that will have their data cleared during a global sign-out in AAD shared device mode. This
	// collection can contain a maximum of 500 elements.
	AzureAdSharedDeviceDataClearApps *[]AppListItem `json:"azureAdSharedDeviceDataClearApps,omitempty"`

	// Indicates whether or not to block a user from configuring bluetooth.
	BluetoothBlockConfiguration nullable.Type[bool] `json:"bluetoothBlockConfiguration,omitempty"`

	// Indicates whether or not to block a user from sharing contacts via bluetooth.
	BluetoothBlockContactSharing nullable.Type[bool] `json:"bluetoothBlockContactSharing,omitempty"`

	// Indicates whether or not to disable the use of the camera.
	CameraBlocked nullable.Type[bool] `json:"cameraBlocked,omitempty"`

	// Indicates whether or not to block Wi-Fi tethering.
	CellularBlockWiFiTethering nullable.Type[bool] `json:"cellularBlockWiFiTethering,omitempty"`

	// Indicates whether or not to block users from any certificate credential configuration.
	CertificateCredentialConfigurationDisabled nullable.Type[bool] `json:"certificateCredentialConfigurationDisabled,omitempty"`

	// Indicates whether or not text copied from one profile (personal or work) can be pasted in the other.
	CrossProfilePoliciesAllowCopyPaste nullable.Type[bool] `json:"crossProfilePoliciesAllowCopyPaste,omitempty"`

	// Indicates whether data from one profile (personal or work) can be shared with apps in the other profile. Possible
	// values are: notConfigured, crossProfileDataSharingBlocked, dataSharingFromWorkToPersonalBlocked,
	// crossProfileDataSharingAllowed, unkownFutureValue.
	CrossProfilePoliciesAllowDataSharing *AndroidDeviceOwnerCrossProfileDataSharing `json:"crossProfilePoliciesAllowDataSharing,omitempty"`

	// Indicates whether or not contacts stored in work profile are shown in personal profile contact searches/incoming
	// calls.
	CrossProfilePoliciesShowWorkContactsInPersonalProfile nullable.Type[bool] `json:"crossProfilePoliciesShowWorkContactsInPersonalProfile,omitempty"`

	// Indicates whether or not to block a user from data roaming.
	DataRoamingBlocked nullable.Type[bool] `json:"dataRoamingBlocked,omitempty"`

	// Indicates whether or not to block the user from manually changing the date or time on the device
	DateTimeConfigurationBlocked nullable.Type[bool] `json:"dateTimeConfigurationBlocked,omitempty"`

	// Represents the customized detailed help text provided to users when they attempt to modify managed settings on their
	// device.
	DetailedHelpText *AndroidDeviceOwnerUserFacingMessage `json:"detailedHelpText,omitempty"`

	// Indicates the location setting configuration for fully managed devices (COBO) and corporate owned devices with a work
	// profile (COPE). Possible values are: notConfigured, disabled, unknownFutureValue.
	DeviceLocationMode *AndroidDeviceOwnerLocationMode `json:"deviceLocationMode,omitempty"`

	// Represents the customized lock screen message provided to users when they attempt to modify managed settings on their
	// device.
	DeviceOwnerLockScreenMessage *AndroidDeviceOwnerUserFacingMessage `json:"deviceOwnerLockScreenMessage,omitempty"`

	// Android Device Owner Enrollment Profile types.
	EnrollmentProfile *AndroidDeviceOwnerEnrollmentProfileType `json:"enrollmentProfile,omitempty"`

	// Indicates whether or not the factory reset option in settings is disabled.
	FactoryResetBlocked nullable.Type[bool] `json:"factoryResetBlocked,omitempty"`

	// List of Google account emails that will be required to authenticate after a device is factory reset before it can be
	// set up.
	FactoryResetDeviceAdministratorEmails *[]string `json:"factoryResetDeviceAdministratorEmails,omitempty"`

	// Proxy is set up directly with host, port and excluded hosts.
	GlobalProxy AndroidDeviceOwnerGlobalProxy `json:"globalProxy"`

	// Indicates whether or not google accounts will be blocked.
	GoogleAccountsBlocked nullable.Type[bool] `json:"googleAccountsBlocked,omitempty"`

	// Indicates whether a user can access the device's Settings app while in Kiosk Mode.
	KioskCustomizationDeviceSettingsBlocked nullable.Type[bool] `json:"kioskCustomizationDeviceSettingsBlocked,omitempty"`

	// Whether the power menu is shown when a user long presses the Power button of a device in Kiosk Mode.
	KioskCustomizationPowerButtonActionsBlocked nullable.Type[bool] `json:"kioskCustomizationPowerButtonActionsBlocked,omitempty"`

	// Indicates whether system info and notifications are disabled in Kiosk Mode. Possible values are: notConfigured,
	// notificationsAndSystemInfoEnabled, systemInfoOnly.
	KioskCustomizationStatusBar *AndroidDeviceOwnerKioskCustomizationStatusBar `json:"kioskCustomizationStatusBar,omitempty"`

	// Indicates whether system error dialogs for crashed or unresponsive apps are shown in Kiosk Mode.
	KioskCustomizationSystemErrorWarnings nullable.Type[bool] `json:"kioskCustomizationSystemErrorWarnings,omitempty"`

	// Indicates which navigation features are enabled in Kiosk Mode. Possible values are: notConfigured, navigationEnabled,
	// homeButtonOnly.
	KioskCustomizationSystemNavigation *AndroidDeviceOwnerKioskCustomizationSystemNavigation `json:"kioskCustomizationSystemNavigation,omitempty"`

	// Whether or not to enable app ordering in Kiosk Mode.
	KioskModeAppOrderEnabled nullable.Type[bool] `json:"kioskModeAppOrderEnabled,omitempty"`

	// The ordering of items on Kiosk Mode Managed Home Screen. This collection can contain a maximum of 500 elements.
	KioskModeAppPositions *[]AndroidDeviceOwnerKioskModeAppPositionItem `json:"kioskModeAppPositions,omitempty"`

	// A list of managed apps that will be shown when the device is in Kiosk Mode. This collection can contain a maximum of
	// 500 elements.
	KioskModeApps *[]AppListItem `json:"kioskModeApps,omitempty"`

	// Whether or not to alphabetize applications within a folder in Kiosk Mode.
	KioskModeAppsInFolderOrderedByName nullable.Type[bool] `json:"kioskModeAppsInFolderOrderedByName,omitempty"`

	// Whether or not to allow a user to configure Bluetooth settings in Kiosk Mode.
	KioskModeBluetoothConfigurationEnabled nullable.Type[bool] `json:"kioskModeBluetoothConfigurationEnabled,omitempty"`

	// Whether or not to allow a user to easy access to the debug menu in Kiosk Mode.
	KioskModeDebugMenuEasyAccessEnabled nullable.Type[bool] `json:"kioskModeDebugMenuEasyAccessEnabled,omitempty"`

	// Exit code to allow a user to escape from Kiosk Mode when the device is in Kiosk Mode.
	KioskModeExitCode nullable.Type[string] `json:"kioskModeExitCode,omitempty"`

	// Whether or not to allow a user to use the flashlight in Kiosk Mode.
	KioskModeFlashlightConfigurationEnabled nullable.Type[bool] `json:"kioskModeFlashlightConfigurationEnabled,omitempty"`

	// Folder icon configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, darkSquare,
	// darkCircle, lightSquare, lightCircle.
	KioskModeFolderIcon *AndroidDeviceOwnerKioskModeFolderIcon `json:"kioskModeFolderIcon,omitempty"`

	// Number of rows for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
	KioskModeGridHeight nullable.Type[int64] `json:"kioskModeGridHeight,omitempty"`

	// Number of columns for Managed Home Screen grid with app ordering enabled in Kiosk Mode. Valid values 1 to 9999999
	KioskModeGridWidth nullable.Type[int64] `json:"kioskModeGridWidth,omitempty"`

	// Icon size configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, smallest, small,
	// regular, large, largest.
	KioskModeIconSize *AndroidDeviceOwnerKioskModeIconSize `json:"kioskModeIconSize,omitempty"`

	// Whether or not to lock home screen to the end user in Kiosk Mode.
	KioskModeLockHomeScreen nullable.Type[bool] `json:"kioskModeLockHomeScreen,omitempty"`

	// A list of managed folders for a device in Kiosk Mode. This collection can contain a maximum of 500 elements.
	KioskModeManagedFolders *[]AndroidDeviceOwnerKioskModeManagedFolder `json:"kioskModeManagedFolders,omitempty"`

	// Whether or not to automatically sign-out of MHS and Shared device mode applications after inactive for Managed Home
	// Screen.
	KioskModeManagedHomeScreenAutoSignout nullable.Type[bool] `json:"kioskModeManagedHomeScreenAutoSignout,omitempty"`

	// Number of seconds to give user notice before automatically signing them out for Managed Home Screen. Valid values 0
	// to 9999999
	KioskModeManagedHomeScreenInactiveSignOutDelayInSeconds nullable.Type[int64] `json:"kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds,omitempty"`

	// Number of seconds device is inactive before automatically signing user out for Managed Home Screen. Valid values 0 to
	// 9999999
	KioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds nullable.Type[int64] `json:"kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds,omitempty"`

	// Complexity of PIN for sign-in session for Managed Home Screen. Possible values are: notConfigured, simple, complex.
	KioskModeManagedHomeScreenPinComplexity *KioskModeManagedHomeScreenPinComplexity `json:"kioskModeManagedHomeScreenPinComplexity,omitempty"`

	// Whether or not require user to set a PIN for sign-in session for Managed Home Screen.
	KioskModeManagedHomeScreenPinRequired nullable.Type[bool] `json:"kioskModeManagedHomeScreenPinRequired,omitempty"`

	// Whether or not required user to enter session PIN if screensaver has appeared for Managed Home Screen.
	KioskModeManagedHomeScreenPinRequiredToResume nullable.Type[bool] `json:"kioskModeManagedHomeScreenPinRequiredToResume,omitempty"`

	// Custom URL background for sign-in screen for Managed Home Screen.
	KioskModeManagedHomeScreenSignInBackground nullable.Type[string] `json:"kioskModeManagedHomeScreenSignInBackground,omitempty"`

	// Custom URL branding logo for sign-in screen and session pin page for Managed Home Screen.
	KioskModeManagedHomeScreenSignInBrandingLogo nullable.Type[string] `json:"kioskModeManagedHomeScreenSignInBrandingLogo,omitempty"`

	// Whether or not show sign-in screen for Managed Home Screen.
	KioskModeManagedHomeScreenSignInEnabled nullable.Type[bool] `json:"kioskModeManagedHomeScreenSignInEnabled,omitempty"`

	// Whether or not to display the Managed Settings entry point on the managed home screen in Kiosk Mode.
	KioskModeManagedSettingsEntryDisabled nullable.Type[bool] `json:"kioskModeManagedSettingsEntryDisabled,omitempty"`

	// Whether or not to allow a user to change the media volume in Kiosk Mode.
	KioskModeMediaVolumeConfigurationEnabled nullable.Type[bool] `json:"kioskModeMediaVolumeConfigurationEnabled,omitempty"`

	// Screen orientation configuration for managed home screen in Kiosk Mode. Possible values are: notConfigured, portrait,
	// landscape, autoRotate.
	KioskModeScreenOrientation *AndroidDeviceOwnerKioskModeScreenOrientation `json:"kioskModeScreenOrientation,omitempty"`

	// Whether or not to enable screen saver mode or not in Kiosk Mode.
	KioskModeScreenSaverConfigurationEnabled nullable.Type[bool] `json:"kioskModeScreenSaverConfigurationEnabled,omitempty"`

	// Whether or not the device screen should show the screen saver if audio/video is playing in Kiosk Mode.
	KioskModeScreenSaverDetectMediaDisabled nullable.Type[bool] `json:"kioskModeScreenSaverDetectMediaDisabled,omitempty"`

	// The number of seconds that the device will display the screen saver for in Kiosk Mode. Valid values 0 to 9999999
	KioskModeScreenSaverDisplayTimeInSeconds nullable.Type[int64] `json:"kioskModeScreenSaverDisplayTimeInSeconds,omitempty"`

	// URL for an image that will be the device's screen saver in Kiosk Mode.
	KioskModeScreenSaverImageUrl nullable.Type[string] `json:"kioskModeScreenSaverImageUrl,omitempty"`

	// The number of seconds the device needs to be inactive for before the screen saver is shown in Kiosk Mode. Valid
	// values 1 to 9999999
	KioskModeScreenSaverStartDelayInSeconds nullable.Type[int64] `json:"kioskModeScreenSaverStartDelayInSeconds,omitempty"`

	// Whether or not to display application notification badges in Kiosk Mode.
	KioskModeShowAppNotificationBadge nullable.Type[bool] `json:"kioskModeShowAppNotificationBadge,omitempty"`

	// Whether or not to allow a user to access basic device information.
	KioskModeShowDeviceInfo nullable.Type[bool] `json:"kioskModeShowDeviceInfo,omitempty"`

	// Whether or not to use single app kiosk mode or multi-app kiosk mode. Possible values are: notConfigured,
	// singleAppMode, multiAppMode.
	KioskModeUseManagedHomeScreenApp *KioskModeType `json:"kioskModeUseManagedHomeScreenApp,omitempty"`

	// Whether or not to display a virtual home button when the device is in Kiosk Mode.
	KioskModeVirtualHomeButtonEnabled nullable.Type[bool] `json:"kioskModeVirtualHomeButtonEnabled,omitempty"`

	// Indicates whether the virtual home button is a swipe up home button or a floating home button. Possible values are:
	// notConfigured, swipeUp, floating.
	KioskModeVirtualHomeButtonType *AndroidDeviceOwnerVirtualHomeButtonType `json:"kioskModeVirtualHomeButtonType,omitempty"`

	// URL to a publicly accessible image to use for the wallpaper when the device is in Kiosk Mode.
	KioskModeWallpaperUrl nullable.Type[string] `json:"kioskModeWallpaperUrl,omitempty"`

	// Whether or not to allow a user to configure Wi-Fi settings in Kiosk Mode.
	KioskModeWiFiConfigurationEnabled nullable.Type[bool] `json:"kioskModeWiFiConfigurationEnabled,omitempty"`

	// The restricted set of WIFI SSIDs available for the user to configure in Kiosk Mode. This collection can contain a
	// maximum of 500 elements.
	KioskModeWifiAllowedSsids *[]string `json:"kioskModeWifiAllowedSsids,omitempty"`

	// Indicates whether or not LocateDevice for devices with lost mode (COBO, COPE) is enabled.
	LocateDeviceLostModeEnabled nullable.Type[bool] `json:"locateDeviceLostModeEnabled,omitempty"`

	// Indicates whether or not LocateDevice for userless (COSU) devices is disabled.
	LocateDeviceUserlessDisabled nullable.Type[bool] `json:"locateDeviceUserlessDisabled,omitempty"`

	// Indicates whether or not to block unmuting the microphone on the device.
	MicrophoneForceMute nullable.Type[bool] `json:"microphoneForceMute,omitempty"`

	// Indicates whether or not to you want configure Microsoft Launcher.
	MicrosoftLauncherConfigurationEnabled nullable.Type[bool] `json:"microsoftLauncherConfigurationEnabled,omitempty"`

	// Indicates whether or not the user can modify the wallpaper to personalize their device.
	MicrosoftLauncherCustomWallpaperAllowUserModification nullable.Type[bool] `json:"microsoftLauncherCustomWallpaperAllowUserModification,omitempty"`

	// Indicates whether or not to configure the wallpaper on the targeted devices.
	MicrosoftLauncherCustomWallpaperEnabled nullable.Type[bool] `json:"microsoftLauncherCustomWallpaperEnabled,omitempty"`

	// Indicates the URL for the image file to use as the wallpaper on the targeted devices.
	MicrosoftLauncherCustomWallpaperImageUrl nullable.Type[string] `json:"microsoftLauncherCustomWallpaperImageUrl,omitempty"`

	// Indicates whether or not the user can modify the device dock configuration on the device.
	MicrosoftLauncherDockPresenceAllowUserModification nullable.Type[bool] `json:"microsoftLauncherDockPresenceAllowUserModification,omitempty"`

	// Indicates whether or not you want to configure the device dock. Possible values are: notConfigured, show, hide,
	// disabled.
	MicrosoftLauncherDockPresenceConfiguration *MicrosoftLauncherDockPresence `json:"microsoftLauncherDockPresenceConfiguration,omitempty"`

	// Indicates whether or not the user can modify the launcher feed on the device.
	MicrosoftLauncherFeedAllowUserModification nullable.Type[bool] `json:"microsoftLauncherFeedAllowUserModification,omitempty"`

	// Indicates whether or not you want to enable the launcher feed on the device.
	MicrosoftLauncherFeedEnabled nullable.Type[bool] `json:"microsoftLauncherFeedEnabled,omitempty"`

	// Indicates the search bar placement configuration on the device. Possible values are: notConfigured, top, bottom,
	// hide.
	MicrosoftLauncherSearchBarPlacementConfiguration *MicrosoftLauncherSearchBarPlacement `json:"microsoftLauncherSearchBarPlacementConfiguration,omitempty"`

	// Indicates whether or not the device will allow connecting to a temporary network connection at boot time.
	NetworkEscapeHatchAllowed nullable.Type[bool] `json:"networkEscapeHatchAllowed,omitempty"`

	// Indicates whether or not to block NFC outgoing beam.
	NfcBlockOutgoingBeam nullable.Type[bool] `json:"nfcBlockOutgoingBeam,omitempty"`

	// Indicates whether or not the keyguard is disabled.
	PasswordBlockKeyguard nullable.Type[bool] `json:"passwordBlockKeyguard,omitempty"`

	// List of device keyguard features to block. This collection can contain a maximum of 11 elements.
	PasswordBlockKeyguardFeatures *[]AndroidKeyguardFeature `json:"passwordBlockKeyguardFeatures,omitempty"`

	// Indicates the amount of time that a password can be set for before it expires and a new password will be required.
	// Valid values 1 to 365
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// Indicates the minimum length of the password required on the device. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// Indicates the minimum number of letter characters required for device password. Valid values 1 to 16
	PasswordMinimumLetterCharacters nullable.Type[int64] `json:"passwordMinimumLetterCharacters,omitempty"`

	// Indicates the minimum number of lower case characters required for device password. Valid values 1 to 16
	PasswordMinimumLowerCaseCharacters nullable.Type[int64] `json:"passwordMinimumLowerCaseCharacters,omitempty"`

	// Indicates the minimum number of non-letter characters required for device password. Valid values 1 to 16
	PasswordMinimumNonLetterCharacters nullable.Type[int64] `json:"passwordMinimumNonLetterCharacters,omitempty"`

	// Indicates the minimum number of numeric characters required for device password. Valid values 1 to 16
	PasswordMinimumNumericCharacters nullable.Type[int64] `json:"passwordMinimumNumericCharacters,omitempty"`

	// Indicates the minimum number of symbol characters required for device password. Valid values 1 to 16
	PasswordMinimumSymbolCharacters nullable.Type[int64] `json:"passwordMinimumSymbolCharacters,omitempty"`

	// Indicates the minimum number of upper case letter characters required for device password. Valid values 1 to 16
	PasswordMinimumUpperCaseCharacters nullable.Type[int64] `json:"passwordMinimumUpperCaseCharacters,omitempty"`

	// Minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// Indicates the length of password history, where the user will not be able to enter a new password that is the same as
	// any password in the history. Valid values 0 to 24
	PasswordPreviousPasswordCountToBlock nullable.Type[int64] `json:"passwordPreviousPasswordCountToBlock,omitempty"`

	// Indicates the timeout period after which a device must be unlocked using a form of strong authentication. Possible
	// values are: deviceDefault, daily, unkownFutureValue.
	PasswordRequireUnlock *AndroidDeviceOwnerRequiredPasswordUnlock `json:"passwordRequireUnlock,omitempty"`

	// Indicates the minimum password quality required on the device. Possible values are: deviceDefault, required, numeric,
	// numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric, customPassword.
	PasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// Indicates the number of times a user can enter an incorrect password before the device is wiped. Valid values 4 to 11
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// Indicates whether the user can install apps from unknown sources on the personal profile.
	PersonalProfileAppsAllowInstallFromUnknownSources nullable.Type[bool] `json:"personalProfileAppsAllowInstallFromUnknownSources,omitempty"`

	// Indicates whether to disable the use of the camera on the personal profile.
	PersonalProfileCameraBlocked nullable.Type[bool] `json:"personalProfileCameraBlocked,omitempty"`

	// Policy applied to applications in the personal profile. This collection can contain a maximum of 500 elements.
	PersonalProfilePersonalApplications *[]AppListItem `json:"personalProfilePersonalApplications,omitempty"`

	// Used together with PersonalProfilePersonalApplications to control how apps in the personal profile are allowed or
	// blocked. Possible values are: notConfigured, blockedApps, allowedApps.
	PersonalProfilePlayStoreMode *PersonalProfilePersonalPlayStoreMode `json:"personalProfilePlayStoreMode,omitempty"`

	// Indicates whether to disable the capability to take screenshots on the personal profile.
	PersonalProfileScreenCaptureBlocked nullable.Type[bool] `json:"personalProfileScreenCaptureBlocked,omitempty"`

	// Indicates the Play Store mode of the device. Possible values are: notConfigured, allowList, blockList.
	PlayStoreMode *AndroidDeviceOwnerPlayStoreMode `json:"playStoreMode,omitempty"`

	// Indicates whether or not to disable the capability to take screenshots.
	ScreenCaptureBlocked nullable.Type[bool] `json:"screenCaptureBlocked,omitempty"`

	// Represents the security common criteria mode enabled provided to users when they attempt to modify managed settings
	// on their device.
	SecurityCommonCriteriaModeEnabled nullable.Type[bool] `json:"securityCommonCriteriaModeEnabled,omitempty"`

	// Indicates whether or not the user is allowed to access developer settings like developer options and safe boot on the
	// device.
	SecurityDeveloperSettingsEnabled nullable.Type[bool] `json:"securityDeveloperSettingsEnabled,omitempty"`

	// Indicates whether or not verify apps is required.
	SecurityRequireVerifyApps nullable.Type[bool] `json:"securityRequireVerifyApps,omitempty"`

	// Indicates whether or not location sharing is disabled for fully managed devices (COBO), and corporate owned devices
	// with a work profile (COPE)
	ShareDeviceLocationDisabled nullable.Type[bool] `json:"shareDeviceLocationDisabled,omitempty"`

	// Represents the customized short help text provided to users when they attempt to modify managed settings on their
	// device.
	ShortHelpText *AndroidDeviceOwnerUserFacingMessage `json:"shortHelpText,omitempty"`

	// Indicates whether or the status bar is disabled, including notifications, quick settings and other screen overlays.
	StatusBarBlocked nullable.Type[bool] `json:"statusBarBlocked,omitempty"`

	// List of modes in which the device's display will stay powered-on. This collection can contain a maximum of 4
	// elements.
	StayOnModes *[]AndroidDeviceOwnerBatteryPluggedMode `json:"stayOnModes,omitempty"`

	// Indicates whether or not to allow USB mass storage.
	StorageAllowUsb nullable.Type[bool] `json:"storageAllowUsb,omitempty"`

	// Indicates whether or not to block external media.
	StorageBlockExternalMedia nullable.Type[bool] `json:"storageBlockExternalMedia,omitempty"`

	// Indicates whether or not to block USB file transfer.
	StorageBlockUsbFileTransfer nullable.Type[bool] `json:"storageBlockUsbFileTransfer,omitempty"`

	// Indicates the annually repeating time periods during which system updates are postponed. This collection can contain
	// a maximum of 500 elements.
	SystemUpdateFreezePeriods *[]AndroidDeviceOwnerSystemUpdateFreezePeriod `json:"systemUpdateFreezePeriods,omitempty"`

	// The type of system update configuration. Possible values are: deviceDefault, postpone, windowed, automatic.
	SystemUpdateInstallType *AndroidDeviceOwnerSystemUpdateInstallType `json:"systemUpdateInstallType,omitempty"`

	// Indicates the number of minutes after midnight that the system update window ends. Valid values 0 to 1440
	SystemUpdateWindowEndMinutesAfterMidnight nullable.Type[int64] `json:"systemUpdateWindowEndMinutesAfterMidnight,omitempty"`

	// Indicates the number of minutes after midnight that the system update window starts. Valid values 0 to 1440
	SystemUpdateWindowStartMinutesAfterMidnight nullable.Type[int64] `json:"systemUpdateWindowStartMinutesAfterMidnight,omitempty"`

	// Whether or not to block Android system prompt windows, like toasts, phone activities, and system alerts.
	SystemWindowsBlocked nullable.Type[bool] `json:"systemWindowsBlocked,omitempty"`

	// Indicates whether or not adding users and profiles is disabled.
	UsersBlockAdd nullable.Type[bool] `json:"usersBlockAdd,omitempty"`

	// Indicates whether or not to disable removing other users from the device.
	UsersBlockRemove nullable.Type[bool] `json:"usersBlockRemove,omitempty"`

	// Indicates whether or not adjusting the master volume is disabled.
	VolumeBlockAdjustment nullable.Type[bool] `json:"volumeBlockAdjustment,omitempty"`

	// If an always on VPN package name is specified, whether or not to lock network traffic when that VPN is disconnected.
	VpnAlwaysOnLockdownMode nullable.Type[bool] `json:"vpnAlwaysOnLockdownMode,omitempty"`

	// Android app package name for app that will handle an always-on VPN connection.
	VpnAlwaysOnPackageIdentifier nullable.Type[string] `json:"vpnAlwaysOnPackageIdentifier,omitempty"`

	// Indicates whether or not to block the user from editing the wifi connection settings.
	WifiBlockEditConfigurations nullable.Type[bool] `json:"wifiBlockEditConfigurations,omitempty"`

	// Indicates whether or not to block the user from editing just the networks defined by the policy.
	WifiBlockEditPolicyDefinedConfigurations nullable.Type[bool] `json:"wifiBlockEditPolicyDefinedConfigurations,omitempty"`

	// Indicates the number of days that a work profile password can be set before it expires and a new password will be
	// required. Valid values 1 to 365
	WorkProfilePasswordExpirationDays nullable.Type[int64] `json:"workProfilePasswordExpirationDays,omitempty"`

	// Indicates the minimum length of the work profile password. Valid values 4 to 16
	WorkProfilePasswordMinimumLength nullable.Type[int64] `json:"workProfilePasswordMinimumLength,omitempty"`

	// Indicates the minimum number of letter characters required for the work profile password. Valid values 1 to 16
	WorkProfilePasswordMinimumLetterCharacters nullable.Type[int64] `json:"workProfilePasswordMinimumLetterCharacters,omitempty"`

	// Indicates the minimum number of lower-case characters required for the work profile password. Valid values 1 to 16
	WorkProfilePasswordMinimumLowerCaseCharacters nullable.Type[int64] `json:"workProfilePasswordMinimumLowerCaseCharacters,omitempty"`

	// Indicates the minimum number of non-letter characters required for the work profile password. Valid values 1 to 16
	WorkProfilePasswordMinimumNonLetterCharacters nullable.Type[int64] `json:"workProfilePasswordMinimumNonLetterCharacters,omitempty"`

	// Indicates the minimum number of numeric characters required for the work profile password. Valid values 1 to 16
	WorkProfilePasswordMinimumNumericCharacters nullable.Type[int64] `json:"workProfilePasswordMinimumNumericCharacters,omitempty"`

	// Indicates the minimum number of symbol characters required for the work profile password. Valid values 1 to 16
	WorkProfilePasswordMinimumSymbolCharacters nullable.Type[int64] `json:"workProfilePasswordMinimumSymbolCharacters,omitempty"`

	// Indicates the minimum number of upper-case letter characters required for the work profile password. Valid values 1
	// to 16
	WorkProfilePasswordMinimumUpperCaseCharacters nullable.Type[int64] `json:"workProfilePasswordMinimumUpperCaseCharacters,omitempty"`

	// Indicates the length of the work profile password history, where the user will not be able to enter a new password
	// that is the same as any password in the history. Valid values 0 to 24
	WorkProfilePasswordPreviousPasswordCountToBlock nullable.Type[int64] `json:"workProfilePasswordPreviousPasswordCountToBlock,omitempty"`

	// Indicates the timeout period after which a work profile must be unlocked using a form of strong authentication.
	// Possible values are: deviceDefault, daily, unkownFutureValue.
	WorkProfilePasswordRequireUnlock *AndroidDeviceOwnerRequiredPasswordUnlock `json:"workProfilePasswordRequireUnlock,omitempty"`

	// Indicates the minimum password quality required on the work profile password. Possible values are: deviceDefault,
	// required, numeric, numericComplex, alphabetic, alphanumeric, alphanumericWithSymbols, lowSecurityBiometric,
	// customPassword.
	WorkProfilePasswordRequiredType *AndroidDeviceOwnerRequiredPasswordType `json:"workProfilePasswordRequiredType,omitempty"`

	// Indicates the number of times a user can enter an incorrect work profile password before the device is wiped. Valid
	// values 4 to 11
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`

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

func (s AndroidDeviceOwnerGeneralDeviceConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s AndroidDeviceOwnerGeneralDeviceConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceOwnerGeneralDeviceConfiguration{}

func (s AndroidDeviceOwnerGeneralDeviceConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceOwnerGeneralDeviceConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceOwnerGeneralDeviceConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceOwnerGeneralDeviceConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceOwnerGeneralDeviceConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceOwnerGeneralDeviceConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AndroidDeviceOwnerGeneralDeviceConfiguration{}

func (s *AndroidDeviceOwnerGeneralDeviceConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountsBlockModification                                nullable.Type[bool]                                   `json:"accountsBlockModification,omitempty"`
		AndroidDeviceOwnerDelegatedScopeAppSettings              *[]AndroidDeviceOwnerDelegatedScopeAppSetting         `json:"androidDeviceOwnerDelegatedScopeAppSettings,omitempty"`
		AppsAllowInstallFromUnknownSources                       nullable.Type[bool]                                   `json:"appsAllowInstallFromUnknownSources,omitempty"`
		AppsAutoUpdatePolicy                                     *AndroidDeviceOwnerAppAutoUpdatePolicyType            `json:"appsAutoUpdatePolicy,omitempty"`
		AppsDefaultPermissionPolicy                              *AndroidDeviceOwnerDefaultAppPermissionPolicyType     `json:"appsDefaultPermissionPolicy,omitempty"`
		AppsRecommendSkippingFirstUseHints                       nullable.Type[bool]                                   `json:"appsRecommendSkippingFirstUseHints,omitempty"`
		BluetoothBlockConfiguration                              nullable.Type[bool]                                   `json:"bluetoothBlockConfiguration,omitempty"`
		BluetoothBlockContactSharing                             nullable.Type[bool]                                   `json:"bluetoothBlockContactSharing,omitempty"`
		CameraBlocked                                            nullable.Type[bool]                                   `json:"cameraBlocked,omitempty"`
		CellularBlockWiFiTethering                               nullable.Type[bool]                                   `json:"cellularBlockWiFiTethering,omitempty"`
		CertificateCredentialConfigurationDisabled               nullable.Type[bool]                                   `json:"certificateCredentialConfigurationDisabled,omitempty"`
		CrossProfilePoliciesAllowCopyPaste                       nullable.Type[bool]                                   `json:"crossProfilePoliciesAllowCopyPaste,omitempty"`
		CrossProfilePoliciesAllowDataSharing                     *AndroidDeviceOwnerCrossProfileDataSharing            `json:"crossProfilePoliciesAllowDataSharing,omitempty"`
		CrossProfilePoliciesShowWorkContactsInPersonalProfile    nullable.Type[bool]                                   `json:"crossProfilePoliciesShowWorkContactsInPersonalProfile,omitempty"`
		DataRoamingBlocked                                       nullable.Type[bool]                                   `json:"dataRoamingBlocked,omitempty"`
		DateTimeConfigurationBlocked                             nullable.Type[bool]                                   `json:"dateTimeConfigurationBlocked,omitempty"`
		DetailedHelpText                                         *AndroidDeviceOwnerUserFacingMessage                  `json:"detailedHelpText,omitempty"`
		DeviceLocationMode                                       *AndroidDeviceOwnerLocationMode                       `json:"deviceLocationMode,omitempty"`
		DeviceOwnerLockScreenMessage                             *AndroidDeviceOwnerUserFacingMessage                  `json:"deviceOwnerLockScreenMessage,omitempty"`
		EnrollmentProfile                                        *AndroidDeviceOwnerEnrollmentProfileType              `json:"enrollmentProfile,omitempty"`
		FactoryResetBlocked                                      nullable.Type[bool]                                   `json:"factoryResetBlocked,omitempty"`
		FactoryResetDeviceAdministratorEmails                    *[]string                                             `json:"factoryResetDeviceAdministratorEmails,omitempty"`
		GoogleAccountsBlocked                                    nullable.Type[bool]                                   `json:"googleAccountsBlocked,omitempty"`
		KioskCustomizationDeviceSettingsBlocked                  nullable.Type[bool]                                   `json:"kioskCustomizationDeviceSettingsBlocked,omitempty"`
		KioskCustomizationPowerButtonActionsBlocked              nullable.Type[bool]                                   `json:"kioskCustomizationPowerButtonActionsBlocked,omitempty"`
		KioskCustomizationStatusBar                              *AndroidDeviceOwnerKioskCustomizationStatusBar        `json:"kioskCustomizationStatusBar,omitempty"`
		KioskCustomizationSystemErrorWarnings                    nullable.Type[bool]                                   `json:"kioskCustomizationSystemErrorWarnings,omitempty"`
		KioskCustomizationSystemNavigation                       *AndroidDeviceOwnerKioskCustomizationSystemNavigation `json:"kioskCustomizationSystemNavigation,omitempty"`
		KioskModeAppOrderEnabled                                 nullable.Type[bool]                                   `json:"kioskModeAppOrderEnabled,omitempty"`
		KioskModeAppPositions                                    *[]AndroidDeviceOwnerKioskModeAppPositionItem         `json:"kioskModeAppPositions,omitempty"`
		KioskModeAppsInFolderOrderedByName                       nullable.Type[bool]                                   `json:"kioskModeAppsInFolderOrderedByName,omitempty"`
		KioskModeBluetoothConfigurationEnabled                   nullable.Type[bool]                                   `json:"kioskModeBluetoothConfigurationEnabled,omitempty"`
		KioskModeDebugMenuEasyAccessEnabled                      nullable.Type[bool]                                   `json:"kioskModeDebugMenuEasyAccessEnabled,omitempty"`
		KioskModeExitCode                                        nullable.Type[string]                                 `json:"kioskModeExitCode,omitempty"`
		KioskModeFlashlightConfigurationEnabled                  nullable.Type[bool]                                   `json:"kioskModeFlashlightConfigurationEnabled,omitempty"`
		KioskModeFolderIcon                                      *AndroidDeviceOwnerKioskModeFolderIcon                `json:"kioskModeFolderIcon,omitempty"`
		KioskModeGridHeight                                      nullable.Type[int64]                                  `json:"kioskModeGridHeight,omitempty"`
		KioskModeGridWidth                                       nullable.Type[int64]                                  `json:"kioskModeGridWidth,omitempty"`
		KioskModeIconSize                                        *AndroidDeviceOwnerKioskModeIconSize                  `json:"kioskModeIconSize,omitempty"`
		KioskModeLockHomeScreen                                  nullable.Type[bool]                                   `json:"kioskModeLockHomeScreen,omitempty"`
		KioskModeManagedFolders                                  *[]AndroidDeviceOwnerKioskModeManagedFolder           `json:"kioskModeManagedFolders,omitempty"`
		KioskModeManagedHomeScreenAutoSignout                    nullable.Type[bool]                                   `json:"kioskModeManagedHomeScreenAutoSignout,omitempty"`
		KioskModeManagedHomeScreenInactiveSignOutDelayInSeconds  nullable.Type[int64]                                  `json:"kioskModeManagedHomeScreenInactiveSignOutDelayInSeconds,omitempty"`
		KioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds nullable.Type[int64]                                  `json:"kioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds,omitempty"`
		KioskModeManagedHomeScreenPinComplexity                  *KioskModeManagedHomeScreenPinComplexity              `json:"kioskModeManagedHomeScreenPinComplexity,omitempty"`
		KioskModeManagedHomeScreenPinRequired                    nullable.Type[bool]                                   `json:"kioskModeManagedHomeScreenPinRequired,omitempty"`
		KioskModeManagedHomeScreenPinRequiredToResume            nullable.Type[bool]                                   `json:"kioskModeManagedHomeScreenPinRequiredToResume,omitempty"`
		KioskModeManagedHomeScreenSignInBackground               nullable.Type[string]                                 `json:"kioskModeManagedHomeScreenSignInBackground,omitempty"`
		KioskModeManagedHomeScreenSignInBrandingLogo             nullable.Type[string]                                 `json:"kioskModeManagedHomeScreenSignInBrandingLogo,omitempty"`
		KioskModeManagedHomeScreenSignInEnabled                  nullable.Type[bool]                                   `json:"kioskModeManagedHomeScreenSignInEnabled,omitempty"`
		KioskModeManagedSettingsEntryDisabled                    nullable.Type[bool]                                   `json:"kioskModeManagedSettingsEntryDisabled,omitempty"`
		KioskModeMediaVolumeConfigurationEnabled                 nullable.Type[bool]                                   `json:"kioskModeMediaVolumeConfigurationEnabled,omitempty"`
		KioskModeScreenOrientation                               *AndroidDeviceOwnerKioskModeScreenOrientation         `json:"kioskModeScreenOrientation,omitempty"`
		KioskModeScreenSaverConfigurationEnabled                 nullable.Type[bool]                                   `json:"kioskModeScreenSaverConfigurationEnabled,omitempty"`
		KioskModeScreenSaverDetectMediaDisabled                  nullable.Type[bool]                                   `json:"kioskModeScreenSaverDetectMediaDisabled,omitempty"`
		KioskModeScreenSaverDisplayTimeInSeconds                 nullable.Type[int64]                                  `json:"kioskModeScreenSaverDisplayTimeInSeconds,omitempty"`
		KioskModeScreenSaverImageUrl                             nullable.Type[string]                                 `json:"kioskModeScreenSaverImageUrl,omitempty"`
		KioskModeScreenSaverStartDelayInSeconds                  nullable.Type[int64]                                  `json:"kioskModeScreenSaverStartDelayInSeconds,omitempty"`
		KioskModeShowAppNotificationBadge                        nullable.Type[bool]                                   `json:"kioskModeShowAppNotificationBadge,omitempty"`
		KioskModeShowDeviceInfo                                  nullable.Type[bool]                                   `json:"kioskModeShowDeviceInfo,omitempty"`
		KioskModeUseManagedHomeScreenApp                         *KioskModeType                                        `json:"kioskModeUseManagedHomeScreenApp,omitempty"`
		KioskModeVirtualHomeButtonEnabled                        nullable.Type[bool]                                   `json:"kioskModeVirtualHomeButtonEnabled,omitempty"`
		KioskModeVirtualHomeButtonType                           *AndroidDeviceOwnerVirtualHomeButtonType              `json:"kioskModeVirtualHomeButtonType,omitempty"`
		KioskModeWallpaperUrl                                    nullable.Type[string]                                 `json:"kioskModeWallpaperUrl,omitempty"`
		KioskModeWiFiConfigurationEnabled                        nullable.Type[bool]                                   `json:"kioskModeWiFiConfigurationEnabled,omitempty"`
		KioskModeWifiAllowedSsids                                *[]string                                             `json:"kioskModeWifiAllowedSsids,omitempty"`
		LocateDeviceLostModeEnabled                              nullable.Type[bool]                                   `json:"locateDeviceLostModeEnabled,omitempty"`
		LocateDeviceUserlessDisabled                             nullable.Type[bool]                                   `json:"locateDeviceUserlessDisabled,omitempty"`
		MicrophoneForceMute                                      nullable.Type[bool]                                   `json:"microphoneForceMute,omitempty"`
		MicrosoftLauncherConfigurationEnabled                    nullable.Type[bool]                                   `json:"microsoftLauncherConfigurationEnabled,omitempty"`
		MicrosoftLauncherCustomWallpaperAllowUserModification    nullable.Type[bool]                                   `json:"microsoftLauncherCustomWallpaperAllowUserModification,omitempty"`
		MicrosoftLauncherCustomWallpaperEnabled                  nullable.Type[bool]                                   `json:"microsoftLauncherCustomWallpaperEnabled,omitempty"`
		MicrosoftLauncherCustomWallpaperImageUrl                 nullable.Type[string]                                 `json:"microsoftLauncherCustomWallpaperImageUrl,omitempty"`
		MicrosoftLauncherDockPresenceAllowUserModification       nullable.Type[bool]                                   `json:"microsoftLauncherDockPresenceAllowUserModification,omitempty"`
		MicrosoftLauncherDockPresenceConfiguration               *MicrosoftLauncherDockPresence                        `json:"microsoftLauncherDockPresenceConfiguration,omitempty"`
		MicrosoftLauncherFeedAllowUserModification               nullable.Type[bool]                                   `json:"microsoftLauncherFeedAllowUserModification,omitempty"`
		MicrosoftLauncherFeedEnabled                             nullable.Type[bool]                                   `json:"microsoftLauncherFeedEnabled,omitempty"`
		MicrosoftLauncherSearchBarPlacementConfiguration         *MicrosoftLauncherSearchBarPlacement                  `json:"microsoftLauncherSearchBarPlacementConfiguration,omitempty"`
		NetworkEscapeHatchAllowed                                nullable.Type[bool]                                   `json:"networkEscapeHatchAllowed,omitempty"`
		NfcBlockOutgoingBeam                                     nullable.Type[bool]                                   `json:"nfcBlockOutgoingBeam,omitempty"`
		PasswordBlockKeyguard                                    nullable.Type[bool]                                   `json:"passwordBlockKeyguard,omitempty"`
		PasswordBlockKeyguardFeatures                            *[]AndroidKeyguardFeature                             `json:"passwordBlockKeyguardFeatures,omitempty"`
		PasswordExpirationDays                                   nullable.Type[int64]                                  `json:"passwordExpirationDays,omitempty"`
		PasswordMinimumLength                                    nullable.Type[int64]                                  `json:"passwordMinimumLength,omitempty"`
		PasswordMinimumLetterCharacters                          nullable.Type[int64]                                  `json:"passwordMinimumLetterCharacters,omitempty"`
		PasswordMinimumLowerCaseCharacters                       nullable.Type[int64]                                  `json:"passwordMinimumLowerCaseCharacters,omitempty"`
		PasswordMinimumNonLetterCharacters                       nullable.Type[int64]                                  `json:"passwordMinimumNonLetterCharacters,omitempty"`
		PasswordMinimumNumericCharacters                         nullable.Type[int64]                                  `json:"passwordMinimumNumericCharacters,omitempty"`
		PasswordMinimumSymbolCharacters                          nullable.Type[int64]                                  `json:"passwordMinimumSymbolCharacters,omitempty"`
		PasswordMinimumUpperCaseCharacters                       nullable.Type[int64]                                  `json:"passwordMinimumUpperCaseCharacters,omitempty"`
		PasswordMinutesOfInactivityBeforeScreenTimeout           nullable.Type[int64]                                  `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasswordPreviousPasswordCountToBlock                     nullable.Type[int64]                                  `json:"passwordPreviousPasswordCountToBlock,omitempty"`
		PasswordRequireUnlock                                    *AndroidDeviceOwnerRequiredPasswordUnlock             `json:"passwordRequireUnlock,omitempty"`
		PasswordRequiredType                                     *AndroidDeviceOwnerRequiredPasswordType               `json:"passwordRequiredType,omitempty"`
		PasswordSignInFailureCountBeforeFactoryReset             nullable.Type[int64]                                  `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
		PersonalProfileAppsAllowInstallFromUnknownSources        nullable.Type[bool]                                   `json:"personalProfileAppsAllowInstallFromUnknownSources,omitempty"`
		PersonalProfileCameraBlocked                             nullable.Type[bool]                                   `json:"personalProfileCameraBlocked,omitempty"`
		PersonalProfilePlayStoreMode                             *PersonalProfilePersonalPlayStoreMode                 `json:"personalProfilePlayStoreMode,omitempty"`
		PersonalProfileScreenCaptureBlocked                      nullable.Type[bool]                                   `json:"personalProfileScreenCaptureBlocked,omitempty"`
		PlayStoreMode                                            *AndroidDeviceOwnerPlayStoreMode                      `json:"playStoreMode,omitempty"`
		ScreenCaptureBlocked                                     nullable.Type[bool]                                   `json:"screenCaptureBlocked,omitempty"`
		SecurityCommonCriteriaModeEnabled                        nullable.Type[bool]                                   `json:"securityCommonCriteriaModeEnabled,omitempty"`
		SecurityDeveloperSettingsEnabled                         nullable.Type[bool]                                   `json:"securityDeveloperSettingsEnabled,omitempty"`
		SecurityRequireVerifyApps                                nullable.Type[bool]                                   `json:"securityRequireVerifyApps,omitempty"`
		ShareDeviceLocationDisabled                              nullable.Type[bool]                                   `json:"shareDeviceLocationDisabled,omitempty"`
		ShortHelpText                                            *AndroidDeviceOwnerUserFacingMessage                  `json:"shortHelpText,omitempty"`
		StatusBarBlocked                                         nullable.Type[bool]                                   `json:"statusBarBlocked,omitempty"`
		StayOnModes                                              *[]AndroidDeviceOwnerBatteryPluggedMode               `json:"stayOnModes,omitempty"`
		StorageAllowUsb                                          nullable.Type[bool]                                   `json:"storageAllowUsb,omitempty"`
		StorageBlockExternalMedia                                nullable.Type[bool]                                   `json:"storageBlockExternalMedia,omitempty"`
		StorageBlockUsbFileTransfer                              nullable.Type[bool]                                   `json:"storageBlockUsbFileTransfer,omitempty"`
		SystemUpdateFreezePeriods                                *[]AndroidDeviceOwnerSystemUpdateFreezePeriod         `json:"systemUpdateFreezePeriods,omitempty"`
		SystemUpdateInstallType                                  *AndroidDeviceOwnerSystemUpdateInstallType            `json:"systemUpdateInstallType,omitempty"`
		SystemUpdateWindowEndMinutesAfterMidnight                nullable.Type[int64]                                  `json:"systemUpdateWindowEndMinutesAfterMidnight,omitempty"`
		SystemUpdateWindowStartMinutesAfterMidnight              nullable.Type[int64]                                  `json:"systemUpdateWindowStartMinutesAfterMidnight,omitempty"`
		SystemWindowsBlocked                                     nullable.Type[bool]                                   `json:"systemWindowsBlocked,omitempty"`
		UsersBlockAdd                                            nullable.Type[bool]                                   `json:"usersBlockAdd,omitempty"`
		UsersBlockRemove                                         nullable.Type[bool]                                   `json:"usersBlockRemove,omitempty"`
		VolumeBlockAdjustment                                    nullable.Type[bool]                                   `json:"volumeBlockAdjustment,omitempty"`
		VpnAlwaysOnLockdownMode                                  nullable.Type[bool]                                   `json:"vpnAlwaysOnLockdownMode,omitempty"`
		VpnAlwaysOnPackageIdentifier                             nullable.Type[string]                                 `json:"vpnAlwaysOnPackageIdentifier,omitempty"`
		WifiBlockEditConfigurations                              nullable.Type[bool]                                   `json:"wifiBlockEditConfigurations,omitempty"`
		WifiBlockEditPolicyDefinedConfigurations                 nullable.Type[bool]                                   `json:"wifiBlockEditPolicyDefinedConfigurations,omitempty"`
		WorkProfilePasswordExpirationDays                        nullable.Type[int64]                                  `json:"workProfilePasswordExpirationDays,omitempty"`
		WorkProfilePasswordMinimumLength                         nullable.Type[int64]                                  `json:"workProfilePasswordMinimumLength,omitempty"`
		WorkProfilePasswordMinimumLetterCharacters               nullable.Type[int64]                                  `json:"workProfilePasswordMinimumLetterCharacters,omitempty"`
		WorkProfilePasswordMinimumLowerCaseCharacters            nullable.Type[int64]                                  `json:"workProfilePasswordMinimumLowerCaseCharacters,omitempty"`
		WorkProfilePasswordMinimumNonLetterCharacters            nullable.Type[int64]                                  `json:"workProfilePasswordMinimumNonLetterCharacters,omitempty"`
		WorkProfilePasswordMinimumNumericCharacters              nullable.Type[int64]                                  `json:"workProfilePasswordMinimumNumericCharacters,omitempty"`
		WorkProfilePasswordMinimumSymbolCharacters               nullable.Type[int64]                                  `json:"workProfilePasswordMinimumSymbolCharacters,omitempty"`
		WorkProfilePasswordMinimumUpperCaseCharacters            nullable.Type[int64]                                  `json:"workProfilePasswordMinimumUpperCaseCharacters,omitempty"`
		WorkProfilePasswordPreviousPasswordCountToBlock          nullable.Type[int64]                                  `json:"workProfilePasswordPreviousPasswordCountToBlock,omitempty"`
		WorkProfilePasswordRequireUnlock                         *AndroidDeviceOwnerRequiredPasswordUnlock             `json:"workProfilePasswordRequireUnlock,omitempty"`
		WorkProfilePasswordRequiredType                          *AndroidDeviceOwnerRequiredPasswordType               `json:"workProfilePasswordRequiredType,omitempty"`
		WorkProfilePasswordSignInFailureCountBeforeFactoryReset  nullable.Type[int64]                                  `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`
		Assignments                                              *[]DeviceConfigurationAssignment                      `json:"assignments,omitempty"`
		CreatedDateTime                                          *string                                               `json:"createdDateTime,omitempty"`
		Description                                              nullable.Type[string]                                 `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode              *DeviceManagementApplicabilityRuleDeviceMode          `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition               *DeviceManagementApplicabilityRuleOsEdition           `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion               *DeviceManagementApplicabilityRuleOsVersion           `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                              *[]SettingStateDeviceSummary                          `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                                     *DeviceConfigurationDeviceOverview                    `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                           *[]DeviceConfigurationDeviceStatus                    `json:"deviceStatuses,omitempty"`
		DisplayName                                              *string                                               `json:"displayName,omitempty"`
		GroupAssignments                                         *[]DeviceConfigurationGroupAssignment                 `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                                     *string                                               `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                          *[]string                                             `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                                        *bool                                                 `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                                       *DeviceConfigurationUserOverview                      `json:"userStatusOverview,omitempty"`
		UserStatuses                                             *[]DeviceConfigurationUserStatus                      `json:"userStatuses,omitempty"`
		Version                                                  *int64                                                `json:"version,omitempty"`
		Id                                                       *string                                               `json:"id,omitempty"`
		ODataId                                                  *string                                               `json:"@odata.id,omitempty"`
		ODataType                                                *string                                               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountsBlockModification = decoded.AccountsBlockModification
	s.AndroidDeviceOwnerDelegatedScopeAppSettings = decoded.AndroidDeviceOwnerDelegatedScopeAppSettings
	s.AppsAllowInstallFromUnknownSources = decoded.AppsAllowInstallFromUnknownSources
	s.AppsAutoUpdatePolicy = decoded.AppsAutoUpdatePolicy
	s.AppsDefaultPermissionPolicy = decoded.AppsDefaultPermissionPolicy
	s.AppsRecommendSkippingFirstUseHints = decoded.AppsRecommendSkippingFirstUseHints
	s.BluetoothBlockConfiguration = decoded.BluetoothBlockConfiguration
	s.BluetoothBlockContactSharing = decoded.BluetoothBlockContactSharing
	s.CameraBlocked = decoded.CameraBlocked
	s.CellularBlockWiFiTethering = decoded.CellularBlockWiFiTethering
	s.CertificateCredentialConfigurationDisabled = decoded.CertificateCredentialConfigurationDisabled
	s.CrossProfilePoliciesAllowCopyPaste = decoded.CrossProfilePoliciesAllowCopyPaste
	s.CrossProfilePoliciesAllowDataSharing = decoded.CrossProfilePoliciesAllowDataSharing
	s.CrossProfilePoliciesShowWorkContactsInPersonalProfile = decoded.CrossProfilePoliciesShowWorkContactsInPersonalProfile
	s.DataRoamingBlocked = decoded.DataRoamingBlocked
	s.DateTimeConfigurationBlocked = decoded.DateTimeConfigurationBlocked
	s.DetailedHelpText = decoded.DetailedHelpText
	s.DeviceLocationMode = decoded.DeviceLocationMode
	s.DeviceOwnerLockScreenMessage = decoded.DeviceOwnerLockScreenMessage
	s.EnrollmentProfile = decoded.EnrollmentProfile
	s.FactoryResetBlocked = decoded.FactoryResetBlocked
	s.FactoryResetDeviceAdministratorEmails = decoded.FactoryResetDeviceAdministratorEmails
	s.GoogleAccountsBlocked = decoded.GoogleAccountsBlocked
	s.KioskCustomizationDeviceSettingsBlocked = decoded.KioskCustomizationDeviceSettingsBlocked
	s.KioskCustomizationPowerButtonActionsBlocked = decoded.KioskCustomizationPowerButtonActionsBlocked
	s.KioskCustomizationStatusBar = decoded.KioskCustomizationStatusBar
	s.KioskCustomizationSystemErrorWarnings = decoded.KioskCustomizationSystemErrorWarnings
	s.KioskCustomizationSystemNavigation = decoded.KioskCustomizationSystemNavigation
	s.KioskModeAppOrderEnabled = decoded.KioskModeAppOrderEnabled
	s.KioskModeAppPositions = decoded.KioskModeAppPositions
	s.KioskModeAppsInFolderOrderedByName = decoded.KioskModeAppsInFolderOrderedByName
	s.KioskModeBluetoothConfigurationEnabled = decoded.KioskModeBluetoothConfigurationEnabled
	s.KioskModeDebugMenuEasyAccessEnabled = decoded.KioskModeDebugMenuEasyAccessEnabled
	s.KioskModeExitCode = decoded.KioskModeExitCode
	s.KioskModeFlashlightConfigurationEnabled = decoded.KioskModeFlashlightConfigurationEnabled
	s.KioskModeFolderIcon = decoded.KioskModeFolderIcon
	s.KioskModeGridHeight = decoded.KioskModeGridHeight
	s.KioskModeGridWidth = decoded.KioskModeGridWidth
	s.KioskModeIconSize = decoded.KioskModeIconSize
	s.KioskModeLockHomeScreen = decoded.KioskModeLockHomeScreen
	s.KioskModeManagedFolders = decoded.KioskModeManagedFolders
	s.KioskModeManagedHomeScreenAutoSignout = decoded.KioskModeManagedHomeScreenAutoSignout
	s.KioskModeManagedHomeScreenInactiveSignOutDelayInSeconds = decoded.KioskModeManagedHomeScreenInactiveSignOutDelayInSeconds
	s.KioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds = decoded.KioskModeManagedHomeScreenInactiveSignOutNoticeInSeconds
	s.KioskModeManagedHomeScreenPinComplexity = decoded.KioskModeManagedHomeScreenPinComplexity
	s.KioskModeManagedHomeScreenPinRequired = decoded.KioskModeManagedHomeScreenPinRequired
	s.KioskModeManagedHomeScreenPinRequiredToResume = decoded.KioskModeManagedHomeScreenPinRequiredToResume
	s.KioskModeManagedHomeScreenSignInBackground = decoded.KioskModeManagedHomeScreenSignInBackground
	s.KioskModeManagedHomeScreenSignInBrandingLogo = decoded.KioskModeManagedHomeScreenSignInBrandingLogo
	s.KioskModeManagedHomeScreenSignInEnabled = decoded.KioskModeManagedHomeScreenSignInEnabled
	s.KioskModeManagedSettingsEntryDisabled = decoded.KioskModeManagedSettingsEntryDisabled
	s.KioskModeMediaVolumeConfigurationEnabled = decoded.KioskModeMediaVolumeConfigurationEnabled
	s.KioskModeScreenOrientation = decoded.KioskModeScreenOrientation
	s.KioskModeScreenSaverConfigurationEnabled = decoded.KioskModeScreenSaverConfigurationEnabled
	s.KioskModeScreenSaverDetectMediaDisabled = decoded.KioskModeScreenSaverDetectMediaDisabled
	s.KioskModeScreenSaverDisplayTimeInSeconds = decoded.KioskModeScreenSaverDisplayTimeInSeconds
	s.KioskModeScreenSaverImageUrl = decoded.KioskModeScreenSaverImageUrl
	s.KioskModeScreenSaverStartDelayInSeconds = decoded.KioskModeScreenSaverStartDelayInSeconds
	s.KioskModeShowAppNotificationBadge = decoded.KioskModeShowAppNotificationBadge
	s.KioskModeShowDeviceInfo = decoded.KioskModeShowDeviceInfo
	s.KioskModeUseManagedHomeScreenApp = decoded.KioskModeUseManagedHomeScreenApp
	s.KioskModeVirtualHomeButtonEnabled = decoded.KioskModeVirtualHomeButtonEnabled
	s.KioskModeVirtualHomeButtonType = decoded.KioskModeVirtualHomeButtonType
	s.KioskModeWallpaperUrl = decoded.KioskModeWallpaperUrl
	s.KioskModeWiFiConfigurationEnabled = decoded.KioskModeWiFiConfigurationEnabled
	s.KioskModeWifiAllowedSsids = decoded.KioskModeWifiAllowedSsids
	s.LocateDeviceLostModeEnabled = decoded.LocateDeviceLostModeEnabled
	s.LocateDeviceUserlessDisabled = decoded.LocateDeviceUserlessDisabled
	s.MicrophoneForceMute = decoded.MicrophoneForceMute
	s.MicrosoftLauncherConfigurationEnabled = decoded.MicrosoftLauncherConfigurationEnabled
	s.MicrosoftLauncherCustomWallpaperAllowUserModification = decoded.MicrosoftLauncherCustomWallpaperAllowUserModification
	s.MicrosoftLauncherCustomWallpaperEnabled = decoded.MicrosoftLauncherCustomWallpaperEnabled
	s.MicrosoftLauncherCustomWallpaperImageUrl = decoded.MicrosoftLauncherCustomWallpaperImageUrl
	s.MicrosoftLauncherDockPresenceAllowUserModification = decoded.MicrosoftLauncherDockPresenceAllowUserModification
	s.MicrosoftLauncherDockPresenceConfiguration = decoded.MicrosoftLauncherDockPresenceConfiguration
	s.MicrosoftLauncherFeedAllowUserModification = decoded.MicrosoftLauncherFeedAllowUserModification
	s.MicrosoftLauncherFeedEnabled = decoded.MicrosoftLauncherFeedEnabled
	s.MicrosoftLauncherSearchBarPlacementConfiguration = decoded.MicrosoftLauncherSearchBarPlacementConfiguration
	s.NetworkEscapeHatchAllowed = decoded.NetworkEscapeHatchAllowed
	s.NfcBlockOutgoingBeam = decoded.NfcBlockOutgoingBeam
	s.PasswordBlockKeyguard = decoded.PasswordBlockKeyguard
	s.PasswordBlockKeyguardFeatures = decoded.PasswordBlockKeyguardFeatures
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinimumLetterCharacters = decoded.PasswordMinimumLetterCharacters
	s.PasswordMinimumLowerCaseCharacters = decoded.PasswordMinimumLowerCaseCharacters
	s.PasswordMinimumNonLetterCharacters = decoded.PasswordMinimumNonLetterCharacters
	s.PasswordMinimumNumericCharacters = decoded.PasswordMinimumNumericCharacters
	s.PasswordMinimumSymbolCharacters = decoded.PasswordMinimumSymbolCharacters
	s.PasswordMinimumUpperCaseCharacters = decoded.PasswordMinimumUpperCaseCharacters
	s.PasswordMinutesOfInactivityBeforeScreenTimeout = decoded.PasswordMinutesOfInactivityBeforeScreenTimeout
	s.PasswordPreviousPasswordCountToBlock = decoded.PasswordPreviousPasswordCountToBlock
	s.PasswordRequireUnlock = decoded.PasswordRequireUnlock
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PasswordSignInFailureCountBeforeFactoryReset = decoded.PasswordSignInFailureCountBeforeFactoryReset
	s.PersonalProfileAppsAllowInstallFromUnknownSources = decoded.PersonalProfileAppsAllowInstallFromUnknownSources
	s.PersonalProfileCameraBlocked = decoded.PersonalProfileCameraBlocked
	s.PersonalProfilePlayStoreMode = decoded.PersonalProfilePlayStoreMode
	s.PersonalProfileScreenCaptureBlocked = decoded.PersonalProfileScreenCaptureBlocked
	s.PlayStoreMode = decoded.PlayStoreMode
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.SecurityCommonCriteriaModeEnabled = decoded.SecurityCommonCriteriaModeEnabled
	s.SecurityDeveloperSettingsEnabled = decoded.SecurityDeveloperSettingsEnabled
	s.SecurityRequireVerifyApps = decoded.SecurityRequireVerifyApps
	s.ShareDeviceLocationDisabled = decoded.ShareDeviceLocationDisabled
	s.ShortHelpText = decoded.ShortHelpText
	s.StatusBarBlocked = decoded.StatusBarBlocked
	s.StayOnModes = decoded.StayOnModes
	s.StorageAllowUsb = decoded.StorageAllowUsb
	s.StorageBlockExternalMedia = decoded.StorageBlockExternalMedia
	s.StorageBlockUsbFileTransfer = decoded.StorageBlockUsbFileTransfer
	s.SystemUpdateFreezePeriods = decoded.SystemUpdateFreezePeriods
	s.SystemUpdateInstallType = decoded.SystemUpdateInstallType
	s.SystemUpdateWindowEndMinutesAfterMidnight = decoded.SystemUpdateWindowEndMinutesAfterMidnight
	s.SystemUpdateWindowStartMinutesAfterMidnight = decoded.SystemUpdateWindowStartMinutesAfterMidnight
	s.SystemWindowsBlocked = decoded.SystemWindowsBlocked
	s.UsersBlockAdd = decoded.UsersBlockAdd
	s.UsersBlockRemove = decoded.UsersBlockRemove
	s.VolumeBlockAdjustment = decoded.VolumeBlockAdjustment
	s.VpnAlwaysOnLockdownMode = decoded.VpnAlwaysOnLockdownMode
	s.VpnAlwaysOnPackageIdentifier = decoded.VpnAlwaysOnPackageIdentifier
	s.WifiBlockEditConfigurations = decoded.WifiBlockEditConfigurations
	s.WifiBlockEditPolicyDefinedConfigurations = decoded.WifiBlockEditPolicyDefinedConfigurations
	s.WorkProfilePasswordExpirationDays = decoded.WorkProfilePasswordExpirationDays
	s.WorkProfilePasswordMinimumLength = decoded.WorkProfilePasswordMinimumLength
	s.WorkProfilePasswordMinimumLetterCharacters = decoded.WorkProfilePasswordMinimumLetterCharacters
	s.WorkProfilePasswordMinimumLowerCaseCharacters = decoded.WorkProfilePasswordMinimumLowerCaseCharacters
	s.WorkProfilePasswordMinimumNonLetterCharacters = decoded.WorkProfilePasswordMinimumNonLetterCharacters
	s.WorkProfilePasswordMinimumNumericCharacters = decoded.WorkProfilePasswordMinimumNumericCharacters
	s.WorkProfilePasswordMinimumSymbolCharacters = decoded.WorkProfilePasswordMinimumSymbolCharacters
	s.WorkProfilePasswordMinimumUpperCaseCharacters = decoded.WorkProfilePasswordMinimumUpperCaseCharacters
	s.WorkProfilePasswordPreviousPasswordCountToBlock = decoded.WorkProfilePasswordPreviousPasswordCountToBlock
	s.WorkProfilePasswordRequireUnlock = decoded.WorkProfilePasswordRequireUnlock
	s.WorkProfilePasswordRequiredType = decoded.WorkProfilePasswordRequiredType
	s.WorkProfilePasswordSignInFailureCountBeforeFactoryReset = decoded.WorkProfilePasswordSignInFailureCountBeforeFactoryReset
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
		return fmt.Errorf("unmarshaling AndroidDeviceOwnerGeneralDeviceConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["azureAdSharedDeviceDataClearApps"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling AzureAdSharedDeviceDataClearApps into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'AzureAdSharedDeviceDataClearApps' for 'AndroidDeviceOwnerGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.AzureAdSharedDeviceDataClearApps = &output
	}

	if v, ok := temp["globalProxy"]; ok {
		impl, err := UnmarshalAndroidDeviceOwnerGlobalProxyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'GlobalProxy' for 'AndroidDeviceOwnerGeneralDeviceConfiguration': %+v", err)
		}
		s.GlobalProxy = impl
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
				return fmt.Errorf("unmarshaling index %d field 'KioskModeApps' for 'AndroidDeviceOwnerGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.KioskModeApps = &output
	}

	if v, ok := temp["personalProfilePersonalApplications"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling PersonalProfilePersonalApplications into list []json.RawMessage: %+v", err)
		}

		output := make([]AppListItem, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalAppListItemImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'PersonalProfilePersonalApplications' for 'AndroidDeviceOwnerGeneralDeviceConfiguration': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.PersonalProfilePersonalApplications = &output
	}

	return nil
}
