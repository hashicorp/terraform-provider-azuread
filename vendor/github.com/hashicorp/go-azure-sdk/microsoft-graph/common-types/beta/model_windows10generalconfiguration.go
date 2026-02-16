package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10GeneralConfiguration{}

type Windows10GeneralConfiguration struct {
	// Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a
	// Microsoft account.
	AccountsBlockAddingNonMicrosoftAccountEmail *bool `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`

	// Possible values of a property
	ActivateAppsWithVoice *Enablement `json:"activateAppsWithVoice,omitempty"`

	// Indicates whether or not to block the user from selecting an AntiTheft mode preference (Windows 10 Mobile only).
	AntiTheftModeBlocked *bool `json:"antiTheftModeBlocked,omitempty"`

	// This policy setting permits users to change installation options that typically are available only to system
	// administrators.
	AppManagementMSIAllowUserControlOverInstall *bool `json:"appManagementMSIAllowUserControlOverInstall,omitempty"`

	// This policy setting directs Windows Installer to use elevated permissions when it installs any program on the system.
	AppManagementMSIAlwaysInstallWithElevatedPrivileges *bool `json:"appManagementMSIAlwaysInstallWithElevatedPrivileges,omitempty"`

	// List of semi-colon delimited Package Family Names of Windows apps. Listed Windows apps are to be launched after
	// logon.​
	AppManagementPackageFamilyNamesToLaunchAfterLogOn *[]string `json:"appManagementPackageFamilyNamesToLaunchAfterLogOn,omitempty"`

	// State Management Setting.
	AppsAllowTrustedAppsSideloading *StateManagementSetting `json:"appsAllowTrustedAppsSideloading,omitempty"`

	// Indicates whether or not to disable the launch of all apps from Windows Store that came pre-installed or were
	// downloaded.
	AppsBlockWindowsStoreOriginatedApps *bool `json:"appsBlockWindowsStoreOriginatedApps,omitempty"`

	// Allows secondary authentication devices to work with Windows.
	AuthenticationAllowSecondaryDevice *bool `json:"authenticationAllowSecondaryDevice,omitempty"`

	// Specifies the preferred domain among available domains in the Azure AD tenant.
	AuthenticationPreferredAzureADTenantDomainName nullable.Type[string] `json:"authenticationPreferredAzureADTenantDomainName,omitempty"`

	// Possible values of a property
	AuthenticationWebSignIn *Enablement `json:"authenticationWebSignIn,omitempty"`

	// Specify a list of allowed Bluetooth services and profiles in hex formatted strings.
	BluetoothAllowedServices *[]string `json:"bluetoothAllowedServices,omitempty"`

	// Whether or not to Block the user from using bluetooth advertising.
	BluetoothBlockAdvertising *bool `json:"bluetoothBlockAdvertising,omitempty"`

	// Whether or not to Block the user from using bluetooth discoverable mode.
	BluetoothBlockDiscoverableMode *bool `json:"bluetoothBlockDiscoverableMode,omitempty"`

	// Whether or not to block specific bundled Bluetooth peripherals to automatically pair with the host device.
	BluetoothBlockPrePairing *bool `json:"bluetoothBlockPrePairing,omitempty"`

	// Whether or not to block the users from using Swift Pair and other proximity based scenarios.
	BluetoothBlockPromptedProximalConnections *bool `json:"bluetoothBlockPromptedProximalConnections,omitempty"`

	// Whether or not to Block the user from using bluetooth.
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`

	// Whether or not to Block the user from accessing the camera of the device.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`

	// Whether or not to Block the user from using data over cellular while roaming.
	CellularBlockDataWhenRoaming *bool `json:"cellularBlockDataWhenRoaming,omitempty"`

	// Whether or not to Block the user from using VPN over cellular.
	CellularBlockVpn *bool `json:"cellularBlockVpn,omitempty"`

	// Whether or not to Block the user from using VPN when roaming over cellular.
	CellularBlockVpnWhenRoaming *bool `json:"cellularBlockVpnWhenRoaming,omitempty"`

	// Possible values of the ConfigurationUsage list.
	CellularData *ConfigurationUsage `json:"cellularData,omitempty"`

	// Whether or not to Block the user from doing manual root certificate installation.
	CertificatesBlockManualRootCertificateInstallation *bool `json:"certificatesBlockManualRootCertificateInstallation,omitempty"`

	// Specifies the time zone to be applied to the device. This is the standard Windows name for the target time zone.
	ConfigureTimeZone nullable.Type[string] `json:"configureTimeZone,omitempty"`

	// Whether or not to block Connected Devices Service which enables discovery and connection to other devices, remote
	// messaging, remote app sessions and other cross-device experiences.
	ConnectedDevicesServiceBlocked *bool `json:"connectedDevicesServiceBlocked,omitempty"`

	// Whether or not to Block the user from using copy paste.
	CopyPasteBlocked *bool `json:"copyPasteBlocked,omitempty"`

	// Whether or not to Block the user from using Cortana.
	CortanaBlocked *bool `json:"cortanaBlocked,omitempty"`

	// Specify whether to allow or disallow the Federal Information Processing Standard (FIPS) policy.
	CryptographyAllowFipsAlgorithmPolicy *bool `json:"cryptographyAllowFipsAlgorithmPolicy,omitempty"`

	// This policy setting allows you to block direct memory access (DMA) for all hot pluggable PCI downstream ports until a
	// user logs into Windows.
	DataProtectionBlockDirectMemoryAccess nullable.Type[bool] `json:"dataProtectionBlockDirectMemoryAccess,omitempty"`

	// Whether or not to block end user access to Defender.
	DefenderBlockEndUserAccess *bool `json:"defenderBlockEndUserAccess,omitempty"`

	// Allows or disallows Windows Defender On Access Protection functionality.
	DefenderBlockOnAccessProtection *bool `json:"defenderBlockOnAccessProtection,omitempty"`

	// Possible values of Cloud Block Level
	DefenderCloudBlockLevel *DefenderCloudBlockLevelType `json:"defenderCloudBlockLevel,omitempty"`

	// Timeout extension for file scanning by the cloud. Valid values 0 to 50
	DefenderCloudExtendedTimeout nullable.Type[int64] `json:"defenderCloudExtendedTimeout,omitempty"`

	// Timeout extension for file scanning by the cloud. Valid values 0 to 50
	DefenderCloudExtendedTimeoutInSeconds nullable.Type[int64] `json:"defenderCloudExtendedTimeoutInSeconds,omitempty"`

	// Number of days before deleting quarantined malware. Valid values 0 to 90
	DefenderDaysBeforeDeletingQuarantinedMalware nullable.Type[int64] `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`

	// Gets or sets Defender’s actions to take on detected Malware per threat level.
	DefenderDetectedMalwareActions *DefenderDetectedMalwareActions `json:"defenderDetectedMalwareActions,omitempty"`

	// When blocked, catch-up scans for scheduled full scans will be turned off.
	DefenderDisableCatchupFullScan *bool `json:"defenderDisableCatchupFullScan,omitempty"`

	// When blocked, catch-up scans for scheduled quick scans will be turned off.
	DefenderDisableCatchupQuickScan *bool `json:"defenderDisableCatchupQuickScan,omitempty"`

	// File extensions to exclude from scans and real time protection.
	DefenderFileExtensionsToExclude *[]string `json:"defenderFileExtensionsToExclude,omitempty"`

	// Files and folder to exclude from scans and real time protection.
	DefenderFilesAndFoldersToExclude *[]string `json:"defenderFilesAndFoldersToExclude,omitempty"`

	// Possible values for monitoring file activity.
	DefenderMonitorFileActivity *DefenderMonitorFileActivity `json:"defenderMonitorFileActivity,omitempty"`

	// Gets or sets Defender’s action to take on Potentially Unwanted Application (PUA), which includes software with
	// behaviors of ad-injection, software bundling, persistent solicitation for payment or subscription, etc. Defender
	// alerts user when PUA is being downloaded or attempts to install itself. Added in Windows 10 for desktop. Possible
	// values are: deviceDefault, block, audit.
	DefenderPotentiallyUnwantedAppAction *DefenderPotentiallyUnwantedAppAction `json:"defenderPotentiallyUnwantedAppAction,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderPotentiallyUnwantedAppActionSetting *DefenderProtectionType `json:"defenderPotentiallyUnwantedAppActionSetting,omitempty"`

	// Processes to exclude from scans and real time protection.
	DefenderProcessesToExclude *[]string `json:"defenderProcessesToExclude,omitempty"`

	// Possible values for prompting user for samples submission.
	DefenderPromptForSampleSubmission *DefenderPromptForSampleSubmission `json:"defenderPromptForSampleSubmission,omitempty"`

	// Indicates whether or not to require behavior monitoring.
	DefenderRequireBehaviorMonitoring *bool `json:"defenderRequireBehaviorMonitoring,omitempty"`

	// Indicates whether or not to require cloud protection.
	DefenderRequireCloudProtection *bool `json:"defenderRequireCloudProtection,omitempty"`

	// Indicates whether or not to require network inspection system.
	DefenderRequireNetworkInspectionSystem *bool `json:"defenderRequireNetworkInspectionSystem,omitempty"`

	// Indicates whether or not to require real time monitoring.
	DefenderRequireRealTimeMonitoring *bool `json:"defenderRequireRealTimeMonitoring,omitempty"`

	// Indicates whether or not to scan archive files.
	DefenderScanArchiveFiles *bool `json:"defenderScanArchiveFiles,omitempty"`

	// Indicates whether or not to scan downloads.
	DefenderScanDownloads *bool `json:"defenderScanDownloads,omitempty"`

	// Indicates whether or not to scan incoming mail messages.
	DefenderScanIncomingMail *bool `json:"defenderScanIncomingMail,omitempty"`

	// Indicates whether or not to scan mapped network drives during full scan.
	DefenderScanMappedNetworkDrivesDuringFullScan *bool `json:"defenderScanMappedNetworkDrivesDuringFullScan,omitempty"`

	// Max CPU usage percentage during scan. Valid values 0 to 100
	DefenderScanMaxCpu nullable.Type[int64] `json:"defenderScanMaxCpu,omitempty"`

	// Indicates whether or not to scan files opened from a network folder.
	DefenderScanNetworkFiles *bool `json:"defenderScanNetworkFiles,omitempty"`

	// Indicates whether or not to scan removable drives during full scan.
	DefenderScanRemovableDrivesDuringFullScan *bool `json:"defenderScanRemovableDrivesDuringFullScan,omitempty"`

	// Indicates whether or not to scan scripts loaded in Internet Explorer browser.
	DefenderScanScriptsLoadedInInternetExplorer *bool `json:"defenderScanScriptsLoadedInInternetExplorer,omitempty"`

	// Possible values for system scan type.
	DefenderScanType *DefenderScanType `json:"defenderScanType,omitempty"`

	// When enabled, low CPU priority will be used during scheduled scans.
	DefenderScheduleScanEnableLowCpuPriority *bool `json:"defenderScheduleScanEnableLowCpuPriority,omitempty"`

	// The time to perform a daily quick scan.
	DefenderScheduledQuickScanTime nullable.Type[string] `json:"defenderScheduledQuickScanTime,omitempty"`

	// The defender time for the system scan.
	DefenderScheduledScanTime nullable.Type[string] `json:"defenderScheduledScanTime,omitempty"`

	// The signature update interval in hours. Specify 0 not to check. Valid values 0 to 24
	DefenderSignatureUpdateIntervalInHours nullable.Type[int64] `json:"defenderSignatureUpdateIntervalInHours,omitempty"`

	// Checks for the user consent level in Windows Defender to send data. Possible values are:
	// sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
	DefenderSubmitSamplesConsentType *DefenderSubmitSamplesConsentType `json:"defenderSubmitSamplesConsentType,omitempty"`

	// Possible values for a weekly schedule.
	DefenderSystemScanSchedule *WeeklySchedule `json:"defenderSystemScanSchedule,omitempty"`

	// State Management Setting.
	DeveloperUnlockSetting *StateManagementSetting `json:"developerUnlockSetting,omitempty"`

	// Indicates whether or not to Block the user from resetting their phone.
	DeviceManagementBlockFactoryResetOnMobile *bool `json:"deviceManagementBlockFactoryResetOnMobile,omitempty"`

	// Indicates whether or not to Block the user from doing manual un-enrollment from device management.
	DeviceManagementBlockManualUnenroll *bool `json:"deviceManagementBlockManualUnenroll,omitempty"`

	// Allow the device to send diagnostic and usage telemetry data, such as Watson.
	DiagnosticsDataSubmissionMode *DiagnosticDataSubmissionMode `json:"diagnosticsDataSubmissionMode,omitempty"`

	// List of legacy applications that have GDI DPI Scaling turned off.
	DisplayAppListWithGdiDPIScalingTurnedOff *[]string `json:"displayAppListWithGdiDPIScalingTurnedOff,omitempty"`

	// List of legacy applications that have GDI DPI Scaling turned on.
	DisplayAppListWithGdiDPIScalingTurnedOn *[]string `json:"displayAppListWithGdiDPIScalingTurnedOn,omitempty"`

	// Allow users to change Start pages on Edge. Use the EdgeHomepageUrls to specify the Start pages that the user would
	// see by default when they open Edge.
	EdgeAllowStartPagesModification *bool `json:"edgeAllowStartPagesModification,omitempty"`

	// Indicates whether or not to prevent access to about flags on Edge browser.
	EdgeBlockAccessToAboutFlags *bool `json:"edgeBlockAccessToAboutFlags,omitempty"`

	// Block the address bar dropdown functionality in Microsoft Edge. Disable this settings to minimize network connections
	// from Microsoft Edge to Microsoft services.
	EdgeBlockAddressBarDropdown *bool `json:"edgeBlockAddressBarDropdown,omitempty"`

	// Indicates whether or not to block auto fill.
	EdgeBlockAutofill *bool `json:"edgeBlockAutofill,omitempty"`

	// Block Microsoft compatibility list in Microsoft Edge. This list from Microsoft helps Edge properly display sites with
	// known compatibility issues.
	EdgeBlockCompatibilityList *bool `json:"edgeBlockCompatibilityList,omitempty"`

	// Indicates whether or not to block developer tools in the Edge browser.
	EdgeBlockDeveloperTools *bool `json:"edgeBlockDeveloperTools,omitempty"`

	// Indicates whether or not to Block the user from making changes to Favorites.
	EdgeBlockEditFavorites *bool `json:"edgeBlockEditFavorites,omitempty"`

	// Indicates whether or not to block extensions in the Edge browser.
	EdgeBlockExtensions *bool `json:"edgeBlockExtensions,omitempty"`

	// Allow or prevent Edge from entering the full screen mode.
	EdgeBlockFullScreenMode *bool `json:"edgeBlockFullScreenMode,omitempty"`

	// Indicates whether or not to block InPrivate browsing on corporate networks, in the Edge browser.
	EdgeBlockInPrivateBrowsing *bool `json:"edgeBlockInPrivateBrowsing,omitempty"`

	// Indicates whether or not to Block the user from using JavaScript.
	EdgeBlockJavaScript *bool `json:"edgeBlockJavaScript,omitempty"`

	// Block the collection of information by Microsoft for live tile creation when users pin a site to Start from Microsoft
	// Edge.
	EdgeBlockLiveTileDataCollection *bool `json:"edgeBlockLiveTileDataCollection,omitempty"`

	// Indicates whether or not to Block password manager.
	EdgeBlockPasswordManager *bool `json:"edgeBlockPasswordManager,omitempty"`

	// Indicates whether or not to block popups.
	EdgeBlockPopups *bool `json:"edgeBlockPopups,omitempty"`

	// Decide whether Microsoft Edge is prelaunched at Windows startup.
	EdgeBlockPrelaunch *bool `json:"edgeBlockPrelaunch,omitempty"`

	// Configure Edge to allow or block printing.
	EdgeBlockPrinting *bool `json:"edgeBlockPrinting,omitempty"`

	// Configure Edge to allow browsing history to be saved or to never save browsing history.
	EdgeBlockSavingHistory *bool `json:"edgeBlockSavingHistory,omitempty"`

	// Indicates whether or not to block the user from adding new search engine or changing the default search engine.
	EdgeBlockSearchEngineCustomization *bool `json:"edgeBlockSearchEngineCustomization,omitempty"`

	// Indicates whether or not to block the user from using the search suggestions in the address bar.
	EdgeBlockSearchSuggestions *bool `json:"edgeBlockSearchSuggestions,omitempty"`

	// Indicates whether or not to Block the user from sending the do not track header.
	EdgeBlockSendingDoNotTrackHeader *bool `json:"edgeBlockSendingDoNotTrackHeader,omitempty"`

	// Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. Note: the name of this
	// property is misleading; the property is obsolete, use EdgeSendIntranetTrafficToInternetExplorer instead.
	EdgeBlockSendingIntranetTrafficToInternetExplorer *bool `json:"edgeBlockSendingIntranetTrafficToInternetExplorer,omitempty"`

	// Indicates whether the user can sideload extensions.
	EdgeBlockSideloadingExtensions *bool `json:"edgeBlockSideloadingExtensions,omitempty"`

	// Configure whether Edge preloads the new tab page at Windows startup.
	EdgeBlockTabPreloading *bool `json:"edgeBlockTabPreloading,omitempty"`

	// Configure to load a blank page in Edge instead of the default New tab page and prevent users from changing it.
	EdgeBlockWebContentOnNewTabPage *bool `json:"edgeBlockWebContentOnNewTabPage,omitempty"`

	// Indicates whether or not to Block the user from using the Edge browser.
	EdgeBlocked *bool `json:"edgeBlocked,omitempty"`

	// Clear browsing data on exiting Microsoft Edge.
	EdgeClearBrowsingDataOnExit *bool `json:"edgeClearBrowsingDataOnExit,omitempty"`

	// Possible values to specify which cookies are allowed in Microsoft Edge.
	EdgeCookiePolicy *EdgeCookiePolicy `json:"edgeCookiePolicy,omitempty"`

	// Block the Microsoft web page that opens on the first use of Microsoft Edge. This policy allows enterprises, like
	// those enrolled in zero emissions configurations, to block this page.
	EdgeDisableFirstRunPage *bool `json:"edgeDisableFirstRunPage,omitempty"`

	// Indicates the enterprise mode site list location. Could be a local file, local network or http location.
	EdgeEnterpriseModeSiteListLocation nullable.Type[string] `json:"edgeEnterpriseModeSiteListLocation,omitempty"`

	// Generic visibility state.
	EdgeFavoritesBarVisibility *VisibilitySetting `json:"edgeFavoritesBarVisibility,omitempty"`

	// The location of the favorites list to provision. Could be a local file, local network or http location.
	EdgeFavoritesListLocation nullable.Type[string] `json:"edgeFavoritesListLocation,omitempty"`

	// The first run URL for when Edge browser is opened for the first time.
	EdgeFirstRunUrl nullable.Type[string] `json:"edgeFirstRunUrl,omitempty"`

	// Causes the Home button to either hide, load the default Start page, load a New tab page, or a custom URL
	EdgeHomeButtonConfiguration EdgeHomeButtonConfiguration `json:"edgeHomeButtonConfiguration"`

	// Enable the Home button configuration.
	EdgeHomeButtonConfigurationEnabled *bool `json:"edgeHomeButtonConfigurationEnabled,omitempty"`

	// The list of URLs for homepages shodwn on MDM-enrolled devices on Edge browser.
	EdgeHomepageUrls *[]string `json:"edgeHomepageUrls,omitempty"`

	// Specify how the Microsoft Edge settings are restricted based on kiosk mode.
	EdgeKioskModeRestriction *EdgeKioskModeRestrictionType `json:"edgeKioskModeRestriction,omitempty"`

	// Specifies the time in minutes from the last user activity before Microsoft Edge kiosk resets. Valid values are
	// 0-1440. The default is 5. 0 indicates no reset. Valid values 0 to 1440
	EdgeKioskResetAfterIdleTimeInMinutes nullable.Type[int64] `json:"edgeKioskResetAfterIdleTimeInMinutes,omitempty"`

	// Specify the page opened when new tabs are created.
	EdgeNewTabPageURL nullable.Type[string] `json:"edgeNewTabPageURL,omitempty"`

	// Possible values for the EdgeOpensWith setting.
	EdgeOpensWith *EdgeOpenOptions `json:"edgeOpensWith,omitempty"`

	// Allow or prevent users from overriding certificate errors.
	EdgePreventCertificateErrorOverride *bool `json:"edgePreventCertificateErrorOverride,omitempty"`

	// Indicates whether or not to Require the user to use the smart screen filter.
	EdgeRequireSmartScreen *bool `json:"edgeRequireSmartScreen,omitempty"`

	// Specify the list of package family names of browser extensions that are required and cannot be turned off by the
	// user.
	EdgeRequiredExtensionPackageFamilyNames *[]string `json:"edgeRequiredExtensionPackageFamilyNames,omitempty"`

	// Allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their
	// default search engine provided the AllowSearchEngineCustomization policy is not set.
	EdgeSearchEngine EdgeSearchEngineBase `json:"edgeSearchEngine"`

	// Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer.
	EdgeSendIntranetTrafficToInternetExplorer *bool `json:"edgeSendIntranetTrafficToInternetExplorer,omitempty"`

	// What message will be displayed by Edge before switching to Internet Explorer.
	EdgeShowMessageWhenOpeningInternetExplorerSites *InternetExplorerMessageSetting `json:"edgeShowMessageWhenOpeningInternetExplorerSites,omitempty"`

	// Enable favorites sync between Internet Explorer and Microsoft Edge. Additions, deletions, modifications and order
	// changes to favorites are shared between browsers.
	EdgeSyncFavoritesWithInternetExplorer *bool `json:"edgeSyncFavoritesWithInternetExplorer,omitempty"`

	// Type of browsing data sent to Microsoft 365 analytics
	EdgeTelemetryForMicrosoft365Analytics *EdgeTelemetryMode `json:"edgeTelemetryForMicrosoft365Analytics,omitempty"`

	// Allow users with administrative rights to delete all user data and settings using CTRL + Win + R at the device lock
	// screen so that the device can be automatically re-configured and re-enrolled into management.
	EnableAutomaticRedeployment *bool `json:"enableAutomaticRedeployment,omitempty"`

	// This setting allows you to specify battery charge level at which Energy Saver is turned on. While on battery, Energy
	// Saver is automatically turned on at (and below) the specified battery charge level. Valid input range (0-100). Valid
	// values 0 to 100
	EnergySaverOnBatteryThresholdPercentage nullable.Type[int64] `json:"energySaverOnBatteryThresholdPercentage,omitempty"`

	// This setting allows you to specify battery charge level at which Energy Saver is turned on. While plugged in, Energy
	// Saver is automatically turned on at (and below) the specified battery charge level. Valid input range (0-100). Valid
	// values 0 to 100
	EnergySaverPluggedInThresholdPercentage nullable.Type[int64] `json:"energySaverPluggedInThresholdPercentage,omitempty"`

	// Endpoint for discovering cloud printers.
	EnterpriseCloudPrintDiscoveryEndPoint nullable.Type[string] `json:"enterpriseCloudPrintDiscoveryEndPoint,omitempty"`

	// Maximum number of printers that should be queried from a discovery endpoint. This is a mobile only setting. Valid
	// values 1 to 65535
	EnterpriseCloudPrintDiscoveryMaxLimit nullable.Type[int64] `json:"enterpriseCloudPrintDiscoveryMaxLimit,omitempty"`

	// OAuth resource URI for printer discovery service as configured in Azure portal.
	EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier nullable.Type[string] `json:"enterpriseCloudPrintMopriaDiscoveryResourceIdentifier,omitempty"`

	// Authentication endpoint for acquiring OAuth tokens.
	EnterpriseCloudPrintOAuthAuthority nullable.Type[string] `json:"enterpriseCloudPrintOAuthAuthority,omitempty"`

	// GUID of a client application authorized to retrieve OAuth tokens from the OAuth Authority.
	EnterpriseCloudPrintOAuthClientIdentifier nullable.Type[string] `json:"enterpriseCloudPrintOAuthClientIdentifier,omitempty"`

	// OAuth resource URI for print service as configured in the Azure portal.
	EnterpriseCloudPrintResourceIdentifier nullable.Type[string] `json:"enterpriseCloudPrintResourceIdentifier,omitempty"`

	// Indicates whether or not to enable device discovery UX.
	ExperienceBlockDeviceDiscovery *bool `json:"experienceBlockDeviceDiscovery,omitempty"`

	// Indicates whether or not to allow the error dialog from displaying if no SIM card is detected.
	ExperienceBlockErrorDialogWhenNoSIM *bool `json:"experienceBlockErrorDialogWhenNoSIM,omitempty"`

	// Indicates whether or not to enable task switching on the device.
	ExperienceBlockTaskSwitcher *bool `json:"experienceBlockTaskSwitcher,omitempty"`

	// Allow(Not Configured) or prevent(Block) the syncing of Microsoft Edge Browser settings. Option to prevent syncing
	// across devices, but allow user override.
	ExperienceDoNotSyncBrowserSettings *BrowserSyncSetting `json:"experienceDoNotSyncBrowserSettings,omitempty"`

	// Possible values of a property
	FindMyFiles *Enablement `json:"findMyFiles,omitempty"`

	// Indicates whether or not to block DVR and broadcasting.
	GameDvrBlocked *bool `json:"gameDvrBlocked,omitempty"`

	// Values for the InkWorkspaceAccess setting.
	InkWorkspaceAccess *InkAccessSetting `json:"inkWorkspaceAccess,omitempty"`

	// State Management Setting.
	InkWorkspaceAccessState *StateManagementSetting `json:"inkWorkspaceAccessState,omitempty"`

	// Specify whether to show recommended app suggestions in the ink workspace.
	InkWorkspaceBlockSuggestedApps *bool `json:"inkWorkspaceBlockSuggestedApps,omitempty"`

	// Indicates whether or not to Block the user from using internet sharing.
	InternetSharingBlocked *bool `json:"internetSharingBlocked,omitempty"`

	// Indicates whether or not to Block the user from location services.
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`

	// Possible values of a property
	LockScreenActivateAppsWithVoice *Enablement `json:"lockScreenActivateAppsWithVoice,omitempty"`

	// Specify whether to show a user-configurable setting to control the screen timeout while on the lock screen of Windows
	// 10 Mobile devices. If this policy is set to Allow, the value set by lockScreenTimeoutInSeconds is ignored.
	LockScreenAllowTimeoutConfiguration *bool `json:"lockScreenAllowTimeoutConfiguration,omitempty"`

	// Indicates whether or not to block action center notifications over lock screen.
	LockScreenBlockActionCenterNotifications *bool `json:"lockScreenBlockActionCenterNotifications,omitempty"`

	// Indicates whether or not the user can interact with Cortana using speech while the system is locked.
	LockScreenBlockCortana *bool `json:"lockScreenBlockCortana,omitempty"`

	// Indicates whether to allow toast notifications above the device lock screen.
	LockScreenBlockToastNotifications *bool `json:"lockScreenBlockToastNotifications,omitempty"`

	// Set the duration (in seconds) from the screen locking to the screen turning off for Windows 10 Mobile devices.
	// Supported values are 11-1800. Valid values 11 to 1800
	LockScreenTimeoutInSeconds nullable.Type[int64] `json:"lockScreenTimeoutInSeconds,omitempty"`

	// Disables the ability to quickly switch between users that are logged on simultaneously without logging off.
	LogonBlockFastUserSwitching *bool `json:"logonBlockFastUserSwitching,omitempty"`

	// Indicates whether or not to block the MMS send/receive functionality on the device.
	MessagingBlockMMS *bool `json:"messagingBlockMMS,omitempty"`

	// Indicates whether or not to block the RCS send/receive functionality on the device.
	MessagingBlockRichCommunicationServices *bool `json:"messagingBlockRichCommunicationServices,omitempty"`

	// Indicates whether or not to block text message back up and restore and Messaging Everywhere.
	MessagingBlockSync *bool `json:"messagingBlockSync,omitempty"`

	// Indicates whether or not to Block Microsoft account settings sync.
	MicrosoftAccountBlockSettingsSync *bool `json:"microsoftAccountBlockSettingsSync,omitempty"`

	// Indicates whether or not to Block a Microsoft account.
	MicrosoftAccountBlocked *bool `json:"microsoftAccountBlocked,omitempty"`

	// Values for the SignInAssistantSettings.
	MicrosoftAccountSignInAssistantSettings *SignInAssistantOptions `json:"microsoftAccountSignInAssistantSettings,omitempty"`

	// If set, proxy settings will be applied to all processes and accounts in the device. Otherwise, it will be applied to
	// the user account that’s enrolled into MDM.
	NetworkProxyApplySettingsDeviceWide *bool `json:"networkProxyApplySettingsDeviceWide,omitempty"`

	// Address to the proxy auto-config (PAC) script you want to use.
	NetworkProxyAutomaticConfigurationUrl nullable.Type[string] `json:"networkProxyAutomaticConfigurationUrl,omitempty"`

	// Disable automatic detection of settings. If enabled, the system will try to find the path to a proxy auto-config
	// (PAC) script.
	NetworkProxyDisableAutoDetect *bool `json:"networkProxyDisableAutoDetect,omitempty"`

	// Specifies manual proxy server settings.
	NetworkProxyServer *Windows10NetworkProxyServer `json:"networkProxyServer,omitempty"`

	// Indicates whether or not to Block the user from using near field communication.
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`

	// Gets or sets a value allowing IT admins to prevent apps and features from working with files on OneDrive.
	OneDriveDisableFileSync *bool `json:"oneDriveDisableFileSync,omitempty"`

	// Specify whether PINs or passwords such as '1111' or '1234' are allowed. For Windows 10 desktops, it also controls the
	// use of picture passwords.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`

	// The password expiration in days. Valid values 0 to 730
	PasswordExpirationDays nullable.Type[int64] `json:"passwordExpirationDays,omitempty"`

	// This security setting determines the period of time (in days) that a password must be used before the user can change
	// it. Valid values 0 to 998
	PasswordMinimumAgeInDays nullable.Type[int64] `json:"passwordMinimumAgeInDays,omitempty"`

	// The number of character sets required in the password.
	PasswordMinimumCharacterSetCount nullable.Type[int64] `json:"passwordMinimumCharacterSetCount,omitempty"`

	// The minimum password length. Valid values 4 to 16
	PasswordMinimumLength nullable.Type[int64] `json:"passwordMinimumLength,omitempty"`

	// The minutes of inactivity before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout nullable.Type[int64] `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`

	// The number of previous passwords to prevent reuse of. Valid values 0 to 50
	PasswordPreviousPasswordBlockCount nullable.Type[int64] `json:"passwordPreviousPasswordBlockCount,omitempty"`

	// Indicates whether or not to require a password upon resuming from an idle state.
	PasswordRequireWhenResumeFromIdleState *bool `json:"passwordRequireWhenResumeFromIdleState,omitempty"`

	// Indicates whether or not to require the user to have a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`

	// Possible values of required passwords.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`

	// The number of sign in failures before factory reset. Valid values 0 to 999
	PasswordSignInFailureCountBeforeFactoryReset nullable.Type[int64] `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`

	// A http or https Url to a jpg, jpeg or png image that needs to be downloaded and used as the Desktop Image or a file
	// Url to a local image on the file system that needs to used as the Desktop Image.
	PersonalizationDesktopImageUrl nullable.Type[string] `json:"personalizationDesktopImageUrl,omitempty"`

	// A http or https Url to a jpg, jpeg or png image that neeeds to be downloaded and used as the Lock Screen Image or a
	// file Url to a local image on the file system that needs to be used as the Lock Screen Image.
	PersonalizationLockScreenImageUrl nullable.Type[string] `json:"personalizationLockScreenImageUrl,omitempty"`

	// Power action types
	PowerButtonActionOnBattery *PowerActionType `json:"powerButtonActionOnBattery,omitempty"`

	// Power action types
	PowerButtonActionPluggedIn *PowerActionType `json:"powerButtonActionPluggedIn,omitempty"`

	// Possible values of a property
	PowerHybridSleepOnBattery *Enablement `json:"powerHybridSleepOnBattery,omitempty"`

	// Possible values of a property
	PowerHybridSleepPluggedIn *Enablement `json:"powerHybridSleepPluggedIn,omitempty"`

	// Power action types
	PowerLidCloseActionOnBattery *PowerActionType `json:"powerLidCloseActionOnBattery,omitempty"`

	// Power action types
	PowerLidCloseActionPluggedIn *PowerActionType `json:"powerLidCloseActionPluggedIn,omitempty"`

	// Power action types
	PowerSleepButtonActionOnBattery *PowerActionType `json:"powerSleepButtonActionOnBattery,omitempty"`

	// Power action types
	PowerSleepButtonActionPluggedIn *PowerActionType `json:"powerSleepButtonActionPluggedIn,omitempty"`

	// Prevent user installation of additional printers from printers settings.
	PrinterBlockAddition *bool `json:"printerBlockAddition,omitempty"`

	// Name (network host name) of an installed printer.
	PrinterDefaultName nullable.Type[string] `json:"printerDefaultName,omitempty"`

	// Automatically provision printers based on their names (network host names).
	PrinterNames *[]string `json:"printerNames,omitempty"`

	// Indicates a list of applications with their access control levels over privacy data categories, and/or the default
	// access levels per category. This collection can contain a maximum of 500 elements.
	PrivacyAccessControls *[]WindowsPrivacyDataAccessControlItem `json:"privacyAccessControls,omitempty"`

	// State Management Setting.
	PrivacyAdvertisingId *StateManagementSetting `json:"privacyAdvertisingId,omitempty"`

	// Indicates whether or not to allow the automatic acceptance of the pairing and privacy user consent dialog when
	// launching apps.
	PrivacyAutoAcceptPairingAndConsentPrompts *bool `json:"privacyAutoAcceptPairingAndConsentPrompts,omitempty"`

	// Blocks the usage of cloud based speech services for Cortana, Dictation, or Store applications.
	PrivacyBlockActivityFeed *bool `json:"privacyBlockActivityFeed,omitempty"`

	// Indicates whether or not to block the usage of cloud based speech services for Cortana, Dictation, or Store
	// applications.
	PrivacyBlockInputPersonalization *bool `json:"privacyBlockInputPersonalization,omitempty"`

	// Blocks the shared experiences/discovery of recently used resources in task switcher etc.
	PrivacyBlockPublishUserActivities *bool `json:"privacyBlockPublishUserActivities,omitempty"`

	// This policy prevents the privacy experience from launching during user logon for new and upgraded users.​
	PrivacyDisableLaunchExperience *bool `json:"privacyDisableLaunchExperience,omitempty"`

	// Indicates whether or not to Block the user from reset protection mode.
	ResetProtectionModeBlocked *bool `json:"resetProtectionModeBlocked,omitempty"`

	// Specifies what level of safe search (filtering adult content) is required
	SafeSearchFilter *SafeSearchFilterType `json:"safeSearchFilter,omitempty"`

	// Indicates whether or not to Block the user from taking Screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`

	// Specifies if search can use diacritics.
	SearchBlockDiacritics *bool `json:"searchBlockDiacritics,omitempty"`

	// Indicates whether or not to block the web search.
	SearchBlockWebResults *bool `json:"searchBlockWebResults,omitempty"`

	// Specifies whether to use automatic language detection when indexing content and properties.
	SearchDisableAutoLanguageDetection *bool `json:"searchDisableAutoLanguageDetection,omitempty"`

	// Indicates whether or not to disable the search indexer backoff feature.
	SearchDisableIndexerBackoff *bool `json:"searchDisableIndexerBackoff,omitempty"`

	// Indicates whether or not to block indexing of WIP-protected items to prevent them from appearing in search results
	// for Cortana or Explorer.
	SearchDisableIndexingEncryptedItems *bool `json:"searchDisableIndexingEncryptedItems,omitempty"`

	// Indicates whether or not to allow users to add locations on removable drives to libraries and to be indexed.
	SearchDisableIndexingRemovableDrive *bool `json:"searchDisableIndexingRemovableDrive,omitempty"`

	// Specifies if search can use location information.
	SearchDisableLocation *bool `json:"searchDisableLocation,omitempty"`

	// Specifies if search can use location information.
	SearchDisableUseLocation *bool `json:"searchDisableUseLocation,omitempty"`

	// Specifies minimum amount of hard drive space on the same drive as the index location before indexing stops.
	SearchEnableAutomaticIndexSizeManangement *bool `json:"searchEnableAutomaticIndexSizeManangement,omitempty"`

	// Indicates whether or not to block remote queries of this computer’s index.
	SearchEnableRemoteQueries *bool `json:"searchEnableRemoteQueries,omitempty"`

	// Specify whether to allow automatic device encryption during OOBE when the device is Azure AD joined (desktop only).
	SecurityBlockAzureADJoinedDevicesAutoEncryption *bool `json:"securityBlockAzureADJoinedDevicesAutoEncryption,omitempty"`

	// Indicates whether or not to block access to Accounts in Settings app.
	SettingsBlockAccountsPage *bool `json:"settingsBlockAccountsPage,omitempty"`

	// Indicates whether or not to block the user from installing provisioning packages.
	SettingsBlockAddProvisioningPackage *bool `json:"settingsBlockAddProvisioningPackage,omitempty"`

	// Indicates whether or not to block access to Apps in Settings app.
	SettingsBlockAppsPage *bool `json:"settingsBlockAppsPage,omitempty"`

	// Indicates whether or not to block the user from changing the language settings.
	SettingsBlockChangeLanguage *bool `json:"settingsBlockChangeLanguage,omitempty"`

	// Indicates whether or not to block the user from changing power and sleep settings.
	SettingsBlockChangePowerSleep *bool `json:"settingsBlockChangePowerSleep,omitempty"`

	// Indicates whether or not to block the user from changing the region settings.
	SettingsBlockChangeRegion *bool `json:"settingsBlockChangeRegion,omitempty"`

	// Indicates whether or not to block the user from changing date and time settings.
	SettingsBlockChangeSystemTime *bool `json:"settingsBlockChangeSystemTime,omitempty"`

	// Indicates whether or not to block access to Devices in Settings app.
	SettingsBlockDevicesPage *bool `json:"settingsBlockDevicesPage,omitempty"`

	// Indicates whether or not to block access to Ease of Access in Settings app.
	SettingsBlockEaseOfAccessPage *bool `json:"settingsBlockEaseOfAccessPage,omitempty"`

	// Indicates whether or not to block the user from editing the device name.
	SettingsBlockEditDeviceName *bool `json:"settingsBlockEditDeviceName,omitempty"`

	// Indicates whether or not to block access to Gaming in Settings app.
	SettingsBlockGamingPage *bool `json:"settingsBlockGamingPage,omitempty"`

	// Indicates whether or not to block access to Network & Internet in Settings app.
	SettingsBlockNetworkInternetPage *bool `json:"settingsBlockNetworkInternetPage,omitempty"`

	// Indicates whether or not to block access to Personalization in Settings app.
	SettingsBlockPersonalizationPage *bool `json:"settingsBlockPersonalizationPage,omitempty"`

	// Indicates whether or not to block access to Privacy in Settings app.
	SettingsBlockPrivacyPage *bool `json:"settingsBlockPrivacyPage,omitempty"`

	// Indicates whether or not to block the runtime configuration agent from removing provisioning packages.
	SettingsBlockRemoveProvisioningPackage *bool `json:"settingsBlockRemoveProvisioningPackage,omitempty"`

	// Indicates whether or not to block access to Settings app.
	SettingsBlockSettingsApp *bool `json:"settingsBlockSettingsApp,omitempty"`

	// Indicates whether or not to block access to System in Settings app.
	SettingsBlockSystemPage *bool `json:"settingsBlockSystemPage,omitempty"`

	// Indicates whether or not to block access to Time & Language in Settings app.
	SettingsBlockTimeLanguagePage *bool `json:"settingsBlockTimeLanguagePage,omitempty"`

	// Indicates whether or not to block access to Update & Security in Settings app.
	SettingsBlockUpdateSecurityPage *bool `json:"settingsBlockUpdateSecurityPage,omitempty"`

	// Indicates whether or not to block multiple users of the same app to share data.
	SharedUserAppDataAllowed *bool `json:"sharedUserAppDataAllowed,omitempty"`

	// App Install control Setting
	SmartScreenAppInstallControl *AppInstallControlType `json:"smartScreenAppInstallControl,omitempty"`

	// Indicates whether or not users can override SmartScreen Filter warnings about potentially malicious websites.
	SmartScreenBlockPromptOverride *bool `json:"smartScreenBlockPromptOverride,omitempty"`

	// Indicates whether or not users can override the SmartScreen Filter warnings about downloading unverified files
	SmartScreenBlockPromptOverrideForFiles *bool `json:"smartScreenBlockPromptOverrideForFiles,omitempty"`

	// This property will be deprecated in July 2019 and will be replaced by property SmartScreenAppInstallControl. Allows
	// IT Admins to control whether users are allowed to install apps from places other than the Store.
	SmartScreenEnableAppInstallControl *bool `json:"smartScreenEnableAppInstallControl,omitempty"`

	// Indicates whether or not to block the user from unpinning apps from taskbar.
	StartBlockUnpinningAppsFromTaskbar *bool `json:"startBlockUnpinningAppsFromTaskbar,omitempty"`

	// Type of start menu app list visibility.
	StartMenuAppListVisibility *WindowsStartMenuAppListVisibilityType `json:"startMenuAppListVisibility,omitempty"`

	// Enabling this policy hides the change account setting from appearing in the user tile in the start menu.
	StartMenuHideChangeAccountSettings *bool `json:"startMenuHideChangeAccountSettings,omitempty"`

	// Enabling this policy hides the most used apps from appearing on the start menu and disables the corresponding toggle
	// in the Settings app.
	StartMenuHideFrequentlyUsedApps *bool `json:"startMenuHideFrequentlyUsedApps,omitempty"`

	// Enabling this policy hides hibernate from appearing in the power button in the start menu.
	StartMenuHideHibernate *bool `json:"startMenuHideHibernate,omitempty"`

	// Enabling this policy hides lock from appearing in the user tile in the start menu.
	StartMenuHideLock *bool `json:"startMenuHideLock,omitempty"`

	// Enabling this policy hides the power button from appearing in the start menu.
	StartMenuHidePowerButton *bool `json:"startMenuHidePowerButton,omitempty"`

	// Enabling this policy hides recent jump lists from appearing on the start menu/taskbar and disables the corresponding
	// toggle in the Settings app.
	StartMenuHideRecentJumpLists *bool `json:"startMenuHideRecentJumpLists,omitempty"`

	// Enabling this policy hides recently added apps from appearing on the start menu and disables the corresponding toggle
	// in the Settings app.
	StartMenuHideRecentlyAddedApps *bool `json:"startMenuHideRecentlyAddedApps,omitempty"`

	// Enabling this policy hides 'Restart/Update and Restart' from appearing in the power button in the start menu.
	StartMenuHideRestartOptions *bool `json:"startMenuHideRestartOptions,omitempty"`

	// Enabling this policy hides shut down/update and shut down from appearing in the power button in the start menu.
	StartMenuHideShutDown *bool `json:"startMenuHideShutDown,omitempty"`

	// Enabling this policy hides sign out from appearing in the user tile in the start menu.
	StartMenuHideSignOut *bool `json:"startMenuHideSignOut,omitempty"`

	// Enabling this policy hides sleep from appearing in the power button in the start menu.
	StartMenuHideSleep *bool `json:"startMenuHideSleep,omitempty"`

	// Enabling this policy hides switch account from appearing in the user tile in the start menu.
	StartMenuHideSwitchAccount *bool `json:"startMenuHideSwitchAccount,omitempty"`

	// Enabling this policy hides the user tile from appearing in the start menu.
	StartMenuHideUserTile *bool `json:"startMenuHideUserTile,omitempty"`

	// This policy setting allows you to import Edge assets to be used with startMenuLayoutXml policy. Start layout can
	// contain secondary tile from Edge app which looks for Edge local asset file. Edge local asset would not exist and
	// cause Edge secondary tile to appear empty in this case. This policy only gets applied when startMenuLayoutXml policy
	// is modified. The value should be a UTF-8 Base64 encoded byte array.
	StartMenuLayoutEdgeAssetsXml nullable.Type[string] `json:"startMenuLayoutEdgeAssetsXml,omitempty"`

	// Allows admins to override the default Start menu layout and prevents the user from changing it. The layout is
	// modified by specifying an XML file based on a layout modification schema. XML needs to be in a UTF8 encoded byte
	// array format.
	StartMenuLayoutXml nullable.Type[string] `json:"startMenuLayoutXml,omitempty"`

	// Type of display modes for the start menu.
	StartMenuMode *WindowsStartMenuModeType `json:"startMenuMode,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderDocuments *VisibilitySetting `json:"startMenuPinnedFolderDocuments,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderDownloads *VisibilitySetting `json:"startMenuPinnedFolderDownloads,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderFileExplorer *VisibilitySetting `json:"startMenuPinnedFolderFileExplorer,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderHomeGroup *VisibilitySetting `json:"startMenuPinnedFolderHomeGroup,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderMusic *VisibilitySetting `json:"startMenuPinnedFolderMusic,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderNetwork *VisibilitySetting `json:"startMenuPinnedFolderNetwork,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderPersonalFolder *VisibilitySetting `json:"startMenuPinnedFolderPersonalFolder,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderPictures *VisibilitySetting `json:"startMenuPinnedFolderPictures,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderSettings *VisibilitySetting `json:"startMenuPinnedFolderSettings,omitempty"`

	// Generic visibility state.
	StartMenuPinnedFolderVideos *VisibilitySetting `json:"startMenuPinnedFolderVideos,omitempty"`

	// Indicates whether or not to Block the user from using removable storage.
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`

	// Indicating whether or not to require encryption on a mobile device.
	StorageRequireMobileDeviceEncryption *bool `json:"storageRequireMobileDeviceEncryption,omitempty"`

	// Indicates whether application data is restricted to the system drive.
	StorageRestrictAppDataToSystemVolume *bool `json:"storageRestrictAppDataToSystemVolume,omitempty"`

	// Indicates whether the installation of applications is restricted to the system drive.
	StorageRestrictAppInstallToSystemVolume *bool `json:"storageRestrictAppInstallToSystemVolume,omitempty"`

	// Gets or sets the fully qualified domain name (FQDN) or IP address of a proxy server to forward Connected User
	// Experiences and Telemetry requests.
	SystemTelemetryProxyServer nullable.Type[string] `json:"systemTelemetryProxyServer,omitempty"`

	// Specify whether non-administrators can use Task Manager to end tasks.
	TaskManagerBlockEndTask *bool `json:"taskManagerBlockEndTask,omitempty"`

	// Whether the device is required to connect to the network.
	TenantLockdownRequireNetworkDuringOutOfBoxExperience *bool `json:"tenantLockdownRequireNetworkDuringOutOfBoxExperience,omitempty"`

	// Indicates whether or not to uninstall a fixed list of built-in Windows apps.
	UninstallBuiltInApps *bool `json:"uninstallBuiltInApps,omitempty"`

	// Indicates whether or not to Block the user from USB connection.
	UsbBlocked *bool `json:"usbBlocked,omitempty"`

	// Indicates whether or not to Block the user from voice recording.
	VoiceRecordingBlocked *bool `json:"voiceRecordingBlocked,omitempty"`

	// Indicates whether or not user's localhost IP address is displayed while making phone calls using the WebRTC
	WebRtcBlockLocalhostIPAddress *bool `json:"webRtcBlockLocalhostIpAddress,omitempty"`

	// Indicating whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
	WiFiBlockAutomaticConnectHotspots *bool `json:"wiFiBlockAutomaticConnectHotspots,omitempty"`

	// Indicates whether or not to Block the user from using Wi-Fi manual configuration.
	WiFiBlockManualConfiguration *bool `json:"wiFiBlockManualConfiguration,omitempty"`

	// Indicates whether or not to Block the user from using Wi-Fi.
	WiFiBlocked *bool `json:"wiFiBlocked,omitempty"`

	// Specify how often devices scan for Wi-Fi networks. Supported values are 1-500, where 100 = default, and 500 = low
	// frequency. Valid values 1 to 500
	WiFiScanInterval nullable.Type[int64] `json:"wiFiScanInterval,omitempty"`

	// Windows 10 force update schedule for Apps.
	Windows10AppsForceUpdateSchedule *Windows10AppsForceUpdateSchedule `json:"windows10AppsForceUpdateSchedule,omitempty"`

	// Allows IT admins to block experiences that are typically for consumers only, such as Start suggestions, Membership
	// notifications, Post-OOBE app install and redirect tiles.
	WindowsSpotlightBlockConsumerSpecificFeatures *bool `json:"windowsSpotlightBlockConsumerSpecificFeatures,omitempty"`

	// Block suggestions from Microsoft that show after each OS clean install, upgrade or in an on-going basis to introduce
	// users to what is new or changed
	WindowsSpotlightBlockOnActionCenter *bool `json:"windowsSpotlightBlockOnActionCenter,omitempty"`

	// Block personalized content in Windows spotlight based on user’s device usage.
	WindowsSpotlightBlockTailoredExperiences *bool `json:"windowsSpotlightBlockTailoredExperiences,omitempty"`

	// Block third party content delivered via Windows Spotlight
	WindowsSpotlightBlockThirdPartyNotifications *bool `json:"windowsSpotlightBlockThirdPartyNotifications,omitempty"`

	// Block Windows Spotlight Windows welcome experience
	WindowsSpotlightBlockWelcomeExperience *bool `json:"windowsSpotlightBlockWelcomeExperience,omitempty"`

	// Allows IT admins to turn off the popup of Windows Tips.
	WindowsSpotlightBlockWindowsTips *bool `json:"windowsSpotlightBlockWindowsTips,omitempty"`

	// Allows IT admins to turn off all Windows Spotlight features
	WindowsSpotlightBlocked *bool `json:"windowsSpotlightBlocked,omitempty"`

	// Allows IT admind to set a predefined default search engine for MDM-Controlled devices
	WindowsSpotlightConfigureOnLockScreen *WindowsSpotlightEnablementSettings `json:"windowsSpotlightConfigureOnLockScreen,omitempty"`

	// Indicates whether or not to block automatic update of apps from Windows Store.
	WindowsStoreBlockAutoUpdate *bool `json:"windowsStoreBlockAutoUpdate,omitempty"`

	// Indicates whether or not to Block the user from using the Windows store.
	WindowsStoreBlocked *bool `json:"windowsStoreBlocked,omitempty"`

	// Indicates whether or not to enable Private Store Only.
	WindowsStoreEnablePrivateStoreOnly *bool `json:"windowsStoreEnablePrivateStoreOnly,omitempty"`

	// Indicates whether or not to allow other devices from discovering this PC for projection.
	WirelessDisplayBlockProjectionToThisDevice *bool `json:"wirelessDisplayBlockProjectionToThisDevice,omitempty"`

	// Indicates whether or not to allow user input from wireless display receiver.
	WirelessDisplayBlockUserInputFromReceiver *bool `json:"wirelessDisplayBlockUserInputFromReceiver,omitempty"`

	// Indicates whether or not to require a PIN for new devices to initiate pairing.
	WirelessDisplayRequirePinForPairing *bool `json:"wirelessDisplayRequirePinForPairing,omitempty"`

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

func (s Windows10GeneralConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s Windows10GeneralConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10GeneralConfiguration{}

func (s Windows10GeneralConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10GeneralConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10GeneralConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10GeneralConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10GeneralConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10GeneralConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Windows10GeneralConfiguration{}

func (s *Windows10GeneralConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountsBlockAddingNonMicrosoftAccountEmail           *bool                                        `json:"accountsBlockAddingNonMicrosoftAccountEmail,omitempty"`
		ActivateAppsWithVoice                                 *Enablement                                  `json:"activateAppsWithVoice,omitempty"`
		AntiTheftModeBlocked                                  *bool                                        `json:"antiTheftModeBlocked,omitempty"`
		AppManagementMSIAllowUserControlOverInstall           *bool                                        `json:"appManagementMSIAllowUserControlOverInstall,omitempty"`
		AppManagementMSIAlwaysInstallWithElevatedPrivileges   *bool                                        `json:"appManagementMSIAlwaysInstallWithElevatedPrivileges,omitempty"`
		AppManagementPackageFamilyNamesToLaunchAfterLogOn     *[]string                                    `json:"appManagementPackageFamilyNamesToLaunchAfterLogOn,omitempty"`
		AppsAllowTrustedAppsSideloading                       *StateManagementSetting                      `json:"appsAllowTrustedAppsSideloading,omitempty"`
		AppsBlockWindowsStoreOriginatedApps                   *bool                                        `json:"appsBlockWindowsStoreOriginatedApps,omitempty"`
		AuthenticationAllowSecondaryDevice                    *bool                                        `json:"authenticationAllowSecondaryDevice,omitempty"`
		AuthenticationPreferredAzureADTenantDomainName        nullable.Type[string]                        `json:"authenticationPreferredAzureADTenantDomainName,omitempty"`
		AuthenticationWebSignIn                               *Enablement                                  `json:"authenticationWebSignIn,omitempty"`
		BluetoothAllowedServices                              *[]string                                    `json:"bluetoothAllowedServices,omitempty"`
		BluetoothBlockAdvertising                             *bool                                        `json:"bluetoothBlockAdvertising,omitempty"`
		BluetoothBlockDiscoverableMode                        *bool                                        `json:"bluetoothBlockDiscoverableMode,omitempty"`
		BluetoothBlockPrePairing                              *bool                                        `json:"bluetoothBlockPrePairing,omitempty"`
		BluetoothBlockPromptedProximalConnections             *bool                                        `json:"bluetoothBlockPromptedProximalConnections,omitempty"`
		BluetoothBlocked                                      *bool                                        `json:"bluetoothBlocked,omitempty"`
		CameraBlocked                                         *bool                                        `json:"cameraBlocked,omitempty"`
		CellularBlockDataWhenRoaming                          *bool                                        `json:"cellularBlockDataWhenRoaming,omitempty"`
		CellularBlockVpn                                      *bool                                        `json:"cellularBlockVpn,omitempty"`
		CellularBlockVpnWhenRoaming                           *bool                                        `json:"cellularBlockVpnWhenRoaming,omitempty"`
		CellularData                                          *ConfigurationUsage                          `json:"cellularData,omitempty"`
		CertificatesBlockManualRootCertificateInstallation    *bool                                        `json:"certificatesBlockManualRootCertificateInstallation,omitempty"`
		ConfigureTimeZone                                     nullable.Type[string]                        `json:"configureTimeZone,omitempty"`
		ConnectedDevicesServiceBlocked                        *bool                                        `json:"connectedDevicesServiceBlocked,omitempty"`
		CopyPasteBlocked                                      *bool                                        `json:"copyPasteBlocked,omitempty"`
		CortanaBlocked                                        *bool                                        `json:"cortanaBlocked,omitempty"`
		CryptographyAllowFipsAlgorithmPolicy                  *bool                                        `json:"cryptographyAllowFipsAlgorithmPolicy,omitempty"`
		DataProtectionBlockDirectMemoryAccess                 nullable.Type[bool]                          `json:"dataProtectionBlockDirectMemoryAccess,omitempty"`
		DefenderBlockEndUserAccess                            *bool                                        `json:"defenderBlockEndUserAccess,omitempty"`
		DefenderBlockOnAccessProtection                       *bool                                        `json:"defenderBlockOnAccessProtection,omitempty"`
		DefenderCloudBlockLevel                               *DefenderCloudBlockLevelType                 `json:"defenderCloudBlockLevel,omitempty"`
		DefenderCloudExtendedTimeout                          nullable.Type[int64]                         `json:"defenderCloudExtendedTimeout,omitempty"`
		DefenderCloudExtendedTimeoutInSeconds                 nullable.Type[int64]                         `json:"defenderCloudExtendedTimeoutInSeconds,omitempty"`
		DefenderDaysBeforeDeletingQuarantinedMalware          nullable.Type[int64]                         `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`
		DefenderDetectedMalwareActions                        *DefenderDetectedMalwareActions              `json:"defenderDetectedMalwareActions,omitempty"`
		DefenderDisableCatchupFullScan                        *bool                                        `json:"defenderDisableCatchupFullScan,omitempty"`
		DefenderDisableCatchupQuickScan                       *bool                                        `json:"defenderDisableCatchupQuickScan,omitempty"`
		DefenderFileExtensionsToExclude                       *[]string                                    `json:"defenderFileExtensionsToExclude,omitempty"`
		DefenderFilesAndFoldersToExclude                      *[]string                                    `json:"defenderFilesAndFoldersToExclude,omitempty"`
		DefenderMonitorFileActivity                           *DefenderMonitorFileActivity                 `json:"defenderMonitorFileActivity,omitempty"`
		DefenderPotentiallyUnwantedAppAction                  *DefenderPotentiallyUnwantedAppAction        `json:"defenderPotentiallyUnwantedAppAction,omitempty"`
		DefenderPotentiallyUnwantedAppActionSetting           *DefenderProtectionType                      `json:"defenderPotentiallyUnwantedAppActionSetting,omitempty"`
		DefenderProcessesToExclude                            *[]string                                    `json:"defenderProcessesToExclude,omitempty"`
		DefenderPromptForSampleSubmission                     *DefenderPromptForSampleSubmission           `json:"defenderPromptForSampleSubmission,omitempty"`
		DefenderRequireBehaviorMonitoring                     *bool                                        `json:"defenderRequireBehaviorMonitoring,omitempty"`
		DefenderRequireCloudProtection                        *bool                                        `json:"defenderRequireCloudProtection,omitempty"`
		DefenderRequireNetworkInspectionSystem                *bool                                        `json:"defenderRequireNetworkInspectionSystem,omitempty"`
		DefenderRequireRealTimeMonitoring                     *bool                                        `json:"defenderRequireRealTimeMonitoring,omitempty"`
		DefenderScanArchiveFiles                              *bool                                        `json:"defenderScanArchiveFiles,omitempty"`
		DefenderScanDownloads                                 *bool                                        `json:"defenderScanDownloads,omitempty"`
		DefenderScanIncomingMail                              *bool                                        `json:"defenderScanIncomingMail,omitempty"`
		DefenderScanMappedNetworkDrivesDuringFullScan         *bool                                        `json:"defenderScanMappedNetworkDrivesDuringFullScan,omitempty"`
		DefenderScanMaxCpu                                    nullable.Type[int64]                         `json:"defenderScanMaxCpu,omitempty"`
		DefenderScanNetworkFiles                              *bool                                        `json:"defenderScanNetworkFiles,omitempty"`
		DefenderScanRemovableDrivesDuringFullScan             *bool                                        `json:"defenderScanRemovableDrivesDuringFullScan,omitempty"`
		DefenderScanScriptsLoadedInInternetExplorer           *bool                                        `json:"defenderScanScriptsLoadedInInternetExplorer,omitempty"`
		DefenderScanType                                      *DefenderScanType                            `json:"defenderScanType,omitempty"`
		DefenderScheduleScanEnableLowCpuPriority              *bool                                        `json:"defenderScheduleScanEnableLowCpuPriority,omitempty"`
		DefenderScheduledQuickScanTime                        nullable.Type[string]                        `json:"defenderScheduledQuickScanTime,omitempty"`
		DefenderScheduledScanTime                             nullable.Type[string]                        `json:"defenderScheduledScanTime,omitempty"`
		DefenderSignatureUpdateIntervalInHours                nullable.Type[int64]                         `json:"defenderSignatureUpdateIntervalInHours,omitempty"`
		DefenderSubmitSamplesConsentType                      *DefenderSubmitSamplesConsentType            `json:"defenderSubmitSamplesConsentType,omitempty"`
		DefenderSystemScanSchedule                            *WeeklySchedule                              `json:"defenderSystemScanSchedule,omitempty"`
		DeveloperUnlockSetting                                *StateManagementSetting                      `json:"developerUnlockSetting,omitempty"`
		DeviceManagementBlockFactoryResetOnMobile             *bool                                        `json:"deviceManagementBlockFactoryResetOnMobile,omitempty"`
		DeviceManagementBlockManualUnenroll                   *bool                                        `json:"deviceManagementBlockManualUnenroll,omitempty"`
		DiagnosticsDataSubmissionMode                         *DiagnosticDataSubmissionMode                `json:"diagnosticsDataSubmissionMode,omitempty"`
		DisplayAppListWithGdiDPIScalingTurnedOff              *[]string                                    `json:"displayAppListWithGdiDPIScalingTurnedOff,omitempty"`
		DisplayAppListWithGdiDPIScalingTurnedOn               *[]string                                    `json:"displayAppListWithGdiDPIScalingTurnedOn,omitempty"`
		EdgeAllowStartPagesModification                       *bool                                        `json:"edgeAllowStartPagesModification,omitempty"`
		EdgeBlockAccessToAboutFlags                           *bool                                        `json:"edgeBlockAccessToAboutFlags,omitempty"`
		EdgeBlockAddressBarDropdown                           *bool                                        `json:"edgeBlockAddressBarDropdown,omitempty"`
		EdgeBlockAutofill                                     *bool                                        `json:"edgeBlockAutofill,omitempty"`
		EdgeBlockCompatibilityList                            *bool                                        `json:"edgeBlockCompatibilityList,omitempty"`
		EdgeBlockDeveloperTools                               *bool                                        `json:"edgeBlockDeveloperTools,omitempty"`
		EdgeBlockEditFavorites                                *bool                                        `json:"edgeBlockEditFavorites,omitempty"`
		EdgeBlockExtensions                                   *bool                                        `json:"edgeBlockExtensions,omitempty"`
		EdgeBlockFullScreenMode                               *bool                                        `json:"edgeBlockFullScreenMode,omitempty"`
		EdgeBlockInPrivateBrowsing                            *bool                                        `json:"edgeBlockInPrivateBrowsing,omitempty"`
		EdgeBlockJavaScript                                   *bool                                        `json:"edgeBlockJavaScript,omitempty"`
		EdgeBlockLiveTileDataCollection                       *bool                                        `json:"edgeBlockLiveTileDataCollection,omitempty"`
		EdgeBlockPasswordManager                              *bool                                        `json:"edgeBlockPasswordManager,omitempty"`
		EdgeBlockPopups                                       *bool                                        `json:"edgeBlockPopups,omitempty"`
		EdgeBlockPrelaunch                                    *bool                                        `json:"edgeBlockPrelaunch,omitempty"`
		EdgeBlockPrinting                                     *bool                                        `json:"edgeBlockPrinting,omitempty"`
		EdgeBlockSavingHistory                                *bool                                        `json:"edgeBlockSavingHistory,omitempty"`
		EdgeBlockSearchEngineCustomization                    *bool                                        `json:"edgeBlockSearchEngineCustomization,omitempty"`
		EdgeBlockSearchSuggestions                            *bool                                        `json:"edgeBlockSearchSuggestions,omitempty"`
		EdgeBlockSendingDoNotTrackHeader                      *bool                                        `json:"edgeBlockSendingDoNotTrackHeader,omitempty"`
		EdgeBlockSendingIntranetTrafficToInternetExplorer     *bool                                        `json:"edgeBlockSendingIntranetTrafficToInternetExplorer,omitempty"`
		EdgeBlockSideloadingExtensions                        *bool                                        `json:"edgeBlockSideloadingExtensions,omitempty"`
		EdgeBlockTabPreloading                                *bool                                        `json:"edgeBlockTabPreloading,omitempty"`
		EdgeBlockWebContentOnNewTabPage                       *bool                                        `json:"edgeBlockWebContentOnNewTabPage,omitempty"`
		EdgeBlocked                                           *bool                                        `json:"edgeBlocked,omitempty"`
		EdgeClearBrowsingDataOnExit                           *bool                                        `json:"edgeClearBrowsingDataOnExit,omitempty"`
		EdgeCookiePolicy                                      *EdgeCookiePolicy                            `json:"edgeCookiePolicy,omitempty"`
		EdgeDisableFirstRunPage                               *bool                                        `json:"edgeDisableFirstRunPage,omitempty"`
		EdgeEnterpriseModeSiteListLocation                    nullable.Type[string]                        `json:"edgeEnterpriseModeSiteListLocation,omitempty"`
		EdgeFavoritesBarVisibility                            *VisibilitySetting                           `json:"edgeFavoritesBarVisibility,omitempty"`
		EdgeFavoritesListLocation                             nullable.Type[string]                        `json:"edgeFavoritesListLocation,omitempty"`
		EdgeFirstRunUrl                                       nullable.Type[string]                        `json:"edgeFirstRunUrl,omitempty"`
		EdgeHomeButtonConfigurationEnabled                    *bool                                        `json:"edgeHomeButtonConfigurationEnabled,omitempty"`
		EdgeHomepageUrls                                      *[]string                                    `json:"edgeHomepageUrls,omitempty"`
		EdgeKioskModeRestriction                              *EdgeKioskModeRestrictionType                `json:"edgeKioskModeRestriction,omitempty"`
		EdgeKioskResetAfterIdleTimeInMinutes                  nullable.Type[int64]                         `json:"edgeKioskResetAfterIdleTimeInMinutes,omitempty"`
		EdgeNewTabPageURL                                     nullable.Type[string]                        `json:"edgeNewTabPageURL,omitempty"`
		EdgeOpensWith                                         *EdgeOpenOptions                             `json:"edgeOpensWith,omitempty"`
		EdgePreventCertificateErrorOverride                   *bool                                        `json:"edgePreventCertificateErrorOverride,omitempty"`
		EdgeRequireSmartScreen                                *bool                                        `json:"edgeRequireSmartScreen,omitempty"`
		EdgeRequiredExtensionPackageFamilyNames               *[]string                                    `json:"edgeRequiredExtensionPackageFamilyNames,omitempty"`
		EdgeSendIntranetTrafficToInternetExplorer             *bool                                        `json:"edgeSendIntranetTrafficToInternetExplorer,omitempty"`
		EdgeShowMessageWhenOpeningInternetExplorerSites       *InternetExplorerMessageSetting              `json:"edgeShowMessageWhenOpeningInternetExplorerSites,omitempty"`
		EdgeSyncFavoritesWithInternetExplorer                 *bool                                        `json:"edgeSyncFavoritesWithInternetExplorer,omitempty"`
		EdgeTelemetryForMicrosoft365Analytics                 *EdgeTelemetryMode                           `json:"edgeTelemetryForMicrosoft365Analytics,omitempty"`
		EnableAutomaticRedeployment                           *bool                                        `json:"enableAutomaticRedeployment,omitempty"`
		EnergySaverOnBatteryThresholdPercentage               nullable.Type[int64]                         `json:"energySaverOnBatteryThresholdPercentage,omitempty"`
		EnergySaverPluggedInThresholdPercentage               nullable.Type[int64]                         `json:"energySaverPluggedInThresholdPercentage,omitempty"`
		EnterpriseCloudPrintDiscoveryEndPoint                 nullable.Type[string]                        `json:"enterpriseCloudPrintDiscoveryEndPoint,omitempty"`
		EnterpriseCloudPrintDiscoveryMaxLimit                 nullable.Type[int64]                         `json:"enterpriseCloudPrintDiscoveryMaxLimit,omitempty"`
		EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier nullable.Type[string]                        `json:"enterpriseCloudPrintMopriaDiscoveryResourceIdentifier,omitempty"`
		EnterpriseCloudPrintOAuthAuthority                    nullable.Type[string]                        `json:"enterpriseCloudPrintOAuthAuthority,omitempty"`
		EnterpriseCloudPrintOAuthClientIdentifier             nullable.Type[string]                        `json:"enterpriseCloudPrintOAuthClientIdentifier,omitempty"`
		EnterpriseCloudPrintResourceIdentifier                nullable.Type[string]                        `json:"enterpriseCloudPrintResourceIdentifier,omitempty"`
		ExperienceBlockDeviceDiscovery                        *bool                                        `json:"experienceBlockDeviceDiscovery,omitempty"`
		ExperienceBlockErrorDialogWhenNoSIM                   *bool                                        `json:"experienceBlockErrorDialogWhenNoSIM,omitempty"`
		ExperienceBlockTaskSwitcher                           *bool                                        `json:"experienceBlockTaskSwitcher,omitempty"`
		ExperienceDoNotSyncBrowserSettings                    *BrowserSyncSetting                          `json:"experienceDoNotSyncBrowserSettings,omitempty"`
		FindMyFiles                                           *Enablement                                  `json:"findMyFiles,omitempty"`
		GameDvrBlocked                                        *bool                                        `json:"gameDvrBlocked,omitempty"`
		InkWorkspaceAccess                                    *InkAccessSetting                            `json:"inkWorkspaceAccess,omitempty"`
		InkWorkspaceAccessState                               *StateManagementSetting                      `json:"inkWorkspaceAccessState,omitempty"`
		InkWorkspaceBlockSuggestedApps                        *bool                                        `json:"inkWorkspaceBlockSuggestedApps,omitempty"`
		InternetSharingBlocked                                *bool                                        `json:"internetSharingBlocked,omitempty"`
		LocationServicesBlocked                               *bool                                        `json:"locationServicesBlocked,omitempty"`
		LockScreenActivateAppsWithVoice                       *Enablement                                  `json:"lockScreenActivateAppsWithVoice,omitempty"`
		LockScreenAllowTimeoutConfiguration                   *bool                                        `json:"lockScreenAllowTimeoutConfiguration,omitempty"`
		LockScreenBlockActionCenterNotifications              *bool                                        `json:"lockScreenBlockActionCenterNotifications,omitempty"`
		LockScreenBlockCortana                                *bool                                        `json:"lockScreenBlockCortana,omitempty"`
		LockScreenBlockToastNotifications                     *bool                                        `json:"lockScreenBlockToastNotifications,omitempty"`
		LockScreenTimeoutInSeconds                            nullable.Type[int64]                         `json:"lockScreenTimeoutInSeconds,omitempty"`
		LogonBlockFastUserSwitching                           *bool                                        `json:"logonBlockFastUserSwitching,omitempty"`
		MessagingBlockMMS                                     *bool                                        `json:"messagingBlockMMS,omitempty"`
		MessagingBlockRichCommunicationServices               *bool                                        `json:"messagingBlockRichCommunicationServices,omitempty"`
		MessagingBlockSync                                    *bool                                        `json:"messagingBlockSync,omitempty"`
		MicrosoftAccountBlockSettingsSync                     *bool                                        `json:"microsoftAccountBlockSettingsSync,omitempty"`
		MicrosoftAccountBlocked                               *bool                                        `json:"microsoftAccountBlocked,omitempty"`
		MicrosoftAccountSignInAssistantSettings               *SignInAssistantOptions                      `json:"microsoftAccountSignInAssistantSettings,omitempty"`
		NetworkProxyApplySettingsDeviceWide                   *bool                                        `json:"networkProxyApplySettingsDeviceWide,omitempty"`
		NetworkProxyAutomaticConfigurationUrl                 nullable.Type[string]                        `json:"networkProxyAutomaticConfigurationUrl,omitempty"`
		NetworkProxyDisableAutoDetect                         *bool                                        `json:"networkProxyDisableAutoDetect,omitempty"`
		NetworkProxyServer                                    *Windows10NetworkProxyServer                 `json:"networkProxyServer,omitempty"`
		NfcBlocked                                            *bool                                        `json:"nfcBlocked,omitempty"`
		OneDriveDisableFileSync                               *bool                                        `json:"oneDriveDisableFileSync,omitempty"`
		PasswordBlockSimple                                   *bool                                        `json:"passwordBlockSimple,omitempty"`
		PasswordExpirationDays                                nullable.Type[int64]                         `json:"passwordExpirationDays,omitempty"`
		PasswordMinimumAgeInDays                              nullable.Type[int64]                         `json:"passwordMinimumAgeInDays,omitempty"`
		PasswordMinimumCharacterSetCount                      nullable.Type[int64]                         `json:"passwordMinimumCharacterSetCount,omitempty"`
		PasswordMinimumLength                                 nullable.Type[int64]                         `json:"passwordMinimumLength,omitempty"`
		PasswordMinutesOfInactivityBeforeScreenTimeout        nullable.Type[int64]                         `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
		PasswordPreviousPasswordBlockCount                    nullable.Type[int64]                         `json:"passwordPreviousPasswordBlockCount,omitempty"`
		PasswordRequireWhenResumeFromIdleState                *bool                                        `json:"passwordRequireWhenResumeFromIdleState,omitempty"`
		PasswordRequired                                      *bool                                        `json:"passwordRequired,omitempty"`
		PasswordRequiredType                                  *RequiredPasswordType                        `json:"passwordRequiredType,omitempty"`
		PasswordSignInFailureCountBeforeFactoryReset          nullable.Type[int64]                         `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
		PersonalizationDesktopImageUrl                        nullable.Type[string]                        `json:"personalizationDesktopImageUrl,omitempty"`
		PersonalizationLockScreenImageUrl                     nullable.Type[string]                        `json:"personalizationLockScreenImageUrl,omitempty"`
		PowerButtonActionOnBattery                            *PowerActionType                             `json:"powerButtonActionOnBattery,omitempty"`
		PowerButtonActionPluggedIn                            *PowerActionType                             `json:"powerButtonActionPluggedIn,omitempty"`
		PowerHybridSleepOnBattery                             *Enablement                                  `json:"powerHybridSleepOnBattery,omitempty"`
		PowerHybridSleepPluggedIn                             *Enablement                                  `json:"powerHybridSleepPluggedIn,omitempty"`
		PowerLidCloseActionOnBattery                          *PowerActionType                             `json:"powerLidCloseActionOnBattery,omitempty"`
		PowerLidCloseActionPluggedIn                          *PowerActionType                             `json:"powerLidCloseActionPluggedIn,omitempty"`
		PowerSleepButtonActionOnBattery                       *PowerActionType                             `json:"powerSleepButtonActionOnBattery,omitempty"`
		PowerSleepButtonActionPluggedIn                       *PowerActionType                             `json:"powerSleepButtonActionPluggedIn,omitempty"`
		PrinterBlockAddition                                  *bool                                        `json:"printerBlockAddition,omitempty"`
		PrinterDefaultName                                    nullable.Type[string]                        `json:"printerDefaultName,omitempty"`
		PrinterNames                                          *[]string                                    `json:"printerNames,omitempty"`
		PrivacyAccessControls                                 *[]WindowsPrivacyDataAccessControlItem       `json:"privacyAccessControls,omitempty"`
		PrivacyAdvertisingId                                  *StateManagementSetting                      `json:"privacyAdvertisingId,omitempty"`
		PrivacyAutoAcceptPairingAndConsentPrompts             *bool                                        `json:"privacyAutoAcceptPairingAndConsentPrompts,omitempty"`
		PrivacyBlockActivityFeed                              *bool                                        `json:"privacyBlockActivityFeed,omitempty"`
		PrivacyBlockInputPersonalization                      *bool                                        `json:"privacyBlockInputPersonalization,omitempty"`
		PrivacyBlockPublishUserActivities                     *bool                                        `json:"privacyBlockPublishUserActivities,omitempty"`
		PrivacyDisableLaunchExperience                        *bool                                        `json:"privacyDisableLaunchExperience,omitempty"`
		ResetProtectionModeBlocked                            *bool                                        `json:"resetProtectionModeBlocked,omitempty"`
		SafeSearchFilter                                      *SafeSearchFilterType                        `json:"safeSearchFilter,omitempty"`
		ScreenCaptureBlocked                                  *bool                                        `json:"screenCaptureBlocked,omitempty"`
		SearchBlockDiacritics                                 *bool                                        `json:"searchBlockDiacritics,omitempty"`
		SearchBlockWebResults                                 *bool                                        `json:"searchBlockWebResults,omitempty"`
		SearchDisableAutoLanguageDetection                    *bool                                        `json:"searchDisableAutoLanguageDetection,omitempty"`
		SearchDisableIndexerBackoff                           *bool                                        `json:"searchDisableIndexerBackoff,omitempty"`
		SearchDisableIndexingEncryptedItems                   *bool                                        `json:"searchDisableIndexingEncryptedItems,omitempty"`
		SearchDisableIndexingRemovableDrive                   *bool                                        `json:"searchDisableIndexingRemovableDrive,omitempty"`
		SearchDisableLocation                                 *bool                                        `json:"searchDisableLocation,omitempty"`
		SearchDisableUseLocation                              *bool                                        `json:"searchDisableUseLocation,omitempty"`
		SearchEnableAutomaticIndexSizeManangement             *bool                                        `json:"searchEnableAutomaticIndexSizeManangement,omitempty"`
		SearchEnableRemoteQueries                             *bool                                        `json:"searchEnableRemoteQueries,omitempty"`
		SecurityBlockAzureADJoinedDevicesAutoEncryption       *bool                                        `json:"securityBlockAzureADJoinedDevicesAutoEncryption,omitempty"`
		SettingsBlockAccountsPage                             *bool                                        `json:"settingsBlockAccountsPage,omitempty"`
		SettingsBlockAddProvisioningPackage                   *bool                                        `json:"settingsBlockAddProvisioningPackage,omitempty"`
		SettingsBlockAppsPage                                 *bool                                        `json:"settingsBlockAppsPage,omitempty"`
		SettingsBlockChangeLanguage                           *bool                                        `json:"settingsBlockChangeLanguage,omitempty"`
		SettingsBlockChangePowerSleep                         *bool                                        `json:"settingsBlockChangePowerSleep,omitempty"`
		SettingsBlockChangeRegion                             *bool                                        `json:"settingsBlockChangeRegion,omitempty"`
		SettingsBlockChangeSystemTime                         *bool                                        `json:"settingsBlockChangeSystemTime,omitempty"`
		SettingsBlockDevicesPage                              *bool                                        `json:"settingsBlockDevicesPage,omitempty"`
		SettingsBlockEaseOfAccessPage                         *bool                                        `json:"settingsBlockEaseOfAccessPage,omitempty"`
		SettingsBlockEditDeviceName                           *bool                                        `json:"settingsBlockEditDeviceName,omitempty"`
		SettingsBlockGamingPage                               *bool                                        `json:"settingsBlockGamingPage,omitempty"`
		SettingsBlockNetworkInternetPage                      *bool                                        `json:"settingsBlockNetworkInternetPage,omitempty"`
		SettingsBlockPersonalizationPage                      *bool                                        `json:"settingsBlockPersonalizationPage,omitempty"`
		SettingsBlockPrivacyPage                              *bool                                        `json:"settingsBlockPrivacyPage,omitempty"`
		SettingsBlockRemoveProvisioningPackage                *bool                                        `json:"settingsBlockRemoveProvisioningPackage,omitempty"`
		SettingsBlockSettingsApp                              *bool                                        `json:"settingsBlockSettingsApp,omitempty"`
		SettingsBlockSystemPage                               *bool                                        `json:"settingsBlockSystemPage,omitempty"`
		SettingsBlockTimeLanguagePage                         *bool                                        `json:"settingsBlockTimeLanguagePage,omitempty"`
		SettingsBlockUpdateSecurityPage                       *bool                                        `json:"settingsBlockUpdateSecurityPage,omitempty"`
		SharedUserAppDataAllowed                              *bool                                        `json:"sharedUserAppDataAllowed,omitempty"`
		SmartScreenAppInstallControl                          *AppInstallControlType                       `json:"smartScreenAppInstallControl,omitempty"`
		SmartScreenBlockPromptOverride                        *bool                                        `json:"smartScreenBlockPromptOverride,omitempty"`
		SmartScreenBlockPromptOverrideForFiles                *bool                                        `json:"smartScreenBlockPromptOverrideForFiles,omitempty"`
		SmartScreenEnableAppInstallControl                    *bool                                        `json:"smartScreenEnableAppInstallControl,omitempty"`
		StartBlockUnpinningAppsFromTaskbar                    *bool                                        `json:"startBlockUnpinningAppsFromTaskbar,omitempty"`
		StartMenuAppListVisibility                            *WindowsStartMenuAppListVisibilityType       `json:"startMenuAppListVisibility,omitempty"`
		StartMenuHideChangeAccountSettings                    *bool                                        `json:"startMenuHideChangeAccountSettings,omitempty"`
		StartMenuHideFrequentlyUsedApps                       *bool                                        `json:"startMenuHideFrequentlyUsedApps,omitempty"`
		StartMenuHideHibernate                                *bool                                        `json:"startMenuHideHibernate,omitempty"`
		StartMenuHideLock                                     *bool                                        `json:"startMenuHideLock,omitempty"`
		StartMenuHidePowerButton                              *bool                                        `json:"startMenuHidePowerButton,omitempty"`
		StartMenuHideRecentJumpLists                          *bool                                        `json:"startMenuHideRecentJumpLists,omitempty"`
		StartMenuHideRecentlyAddedApps                        *bool                                        `json:"startMenuHideRecentlyAddedApps,omitempty"`
		StartMenuHideRestartOptions                           *bool                                        `json:"startMenuHideRestartOptions,omitempty"`
		StartMenuHideShutDown                                 *bool                                        `json:"startMenuHideShutDown,omitempty"`
		StartMenuHideSignOut                                  *bool                                        `json:"startMenuHideSignOut,omitempty"`
		StartMenuHideSleep                                    *bool                                        `json:"startMenuHideSleep,omitempty"`
		StartMenuHideSwitchAccount                            *bool                                        `json:"startMenuHideSwitchAccount,omitempty"`
		StartMenuHideUserTile                                 *bool                                        `json:"startMenuHideUserTile,omitempty"`
		StartMenuLayoutEdgeAssetsXml                          nullable.Type[string]                        `json:"startMenuLayoutEdgeAssetsXml,omitempty"`
		StartMenuLayoutXml                                    nullable.Type[string]                        `json:"startMenuLayoutXml,omitempty"`
		StartMenuMode                                         *WindowsStartMenuModeType                    `json:"startMenuMode,omitempty"`
		StartMenuPinnedFolderDocuments                        *VisibilitySetting                           `json:"startMenuPinnedFolderDocuments,omitempty"`
		StartMenuPinnedFolderDownloads                        *VisibilitySetting                           `json:"startMenuPinnedFolderDownloads,omitempty"`
		StartMenuPinnedFolderFileExplorer                     *VisibilitySetting                           `json:"startMenuPinnedFolderFileExplorer,omitempty"`
		StartMenuPinnedFolderHomeGroup                        *VisibilitySetting                           `json:"startMenuPinnedFolderHomeGroup,omitempty"`
		StartMenuPinnedFolderMusic                            *VisibilitySetting                           `json:"startMenuPinnedFolderMusic,omitempty"`
		StartMenuPinnedFolderNetwork                          *VisibilitySetting                           `json:"startMenuPinnedFolderNetwork,omitempty"`
		StartMenuPinnedFolderPersonalFolder                   *VisibilitySetting                           `json:"startMenuPinnedFolderPersonalFolder,omitempty"`
		StartMenuPinnedFolderPictures                         *VisibilitySetting                           `json:"startMenuPinnedFolderPictures,omitempty"`
		StartMenuPinnedFolderSettings                         *VisibilitySetting                           `json:"startMenuPinnedFolderSettings,omitempty"`
		StartMenuPinnedFolderVideos                           *VisibilitySetting                           `json:"startMenuPinnedFolderVideos,omitempty"`
		StorageBlockRemovableStorage                          *bool                                        `json:"storageBlockRemovableStorage,omitempty"`
		StorageRequireMobileDeviceEncryption                  *bool                                        `json:"storageRequireMobileDeviceEncryption,omitempty"`
		StorageRestrictAppDataToSystemVolume                  *bool                                        `json:"storageRestrictAppDataToSystemVolume,omitempty"`
		StorageRestrictAppInstallToSystemVolume               *bool                                        `json:"storageRestrictAppInstallToSystemVolume,omitempty"`
		SystemTelemetryProxyServer                            nullable.Type[string]                        `json:"systemTelemetryProxyServer,omitempty"`
		TaskManagerBlockEndTask                               *bool                                        `json:"taskManagerBlockEndTask,omitempty"`
		TenantLockdownRequireNetworkDuringOutOfBoxExperience  *bool                                        `json:"tenantLockdownRequireNetworkDuringOutOfBoxExperience,omitempty"`
		UninstallBuiltInApps                                  *bool                                        `json:"uninstallBuiltInApps,omitempty"`
		UsbBlocked                                            *bool                                        `json:"usbBlocked,omitempty"`
		VoiceRecordingBlocked                                 *bool                                        `json:"voiceRecordingBlocked,omitempty"`
		WebRtcBlockLocalhostIPAddress                         *bool                                        `json:"webRtcBlockLocalhostIpAddress,omitempty"`
		WiFiBlockAutomaticConnectHotspots                     *bool                                        `json:"wiFiBlockAutomaticConnectHotspots,omitempty"`
		WiFiBlockManualConfiguration                          *bool                                        `json:"wiFiBlockManualConfiguration,omitempty"`
		WiFiBlocked                                           *bool                                        `json:"wiFiBlocked,omitempty"`
		WiFiScanInterval                                      nullable.Type[int64]                         `json:"wiFiScanInterval,omitempty"`
		Windows10AppsForceUpdateSchedule                      *Windows10AppsForceUpdateSchedule            `json:"windows10AppsForceUpdateSchedule,omitempty"`
		WindowsSpotlightBlockConsumerSpecificFeatures         *bool                                        `json:"windowsSpotlightBlockConsumerSpecificFeatures,omitempty"`
		WindowsSpotlightBlockOnActionCenter                   *bool                                        `json:"windowsSpotlightBlockOnActionCenter,omitempty"`
		WindowsSpotlightBlockTailoredExperiences              *bool                                        `json:"windowsSpotlightBlockTailoredExperiences,omitempty"`
		WindowsSpotlightBlockThirdPartyNotifications          *bool                                        `json:"windowsSpotlightBlockThirdPartyNotifications,omitempty"`
		WindowsSpotlightBlockWelcomeExperience                *bool                                        `json:"windowsSpotlightBlockWelcomeExperience,omitempty"`
		WindowsSpotlightBlockWindowsTips                      *bool                                        `json:"windowsSpotlightBlockWindowsTips,omitempty"`
		WindowsSpotlightBlocked                               *bool                                        `json:"windowsSpotlightBlocked,omitempty"`
		WindowsSpotlightConfigureOnLockScreen                 *WindowsSpotlightEnablementSettings          `json:"windowsSpotlightConfigureOnLockScreen,omitempty"`
		WindowsStoreBlockAutoUpdate                           *bool                                        `json:"windowsStoreBlockAutoUpdate,omitempty"`
		WindowsStoreBlocked                                   *bool                                        `json:"windowsStoreBlocked,omitempty"`
		WindowsStoreEnablePrivateStoreOnly                    *bool                                        `json:"windowsStoreEnablePrivateStoreOnly,omitempty"`
		WirelessDisplayBlockProjectionToThisDevice            *bool                                        `json:"wirelessDisplayBlockProjectionToThisDevice,omitempty"`
		WirelessDisplayBlockUserInputFromReceiver             *bool                                        `json:"wirelessDisplayBlockUserInputFromReceiver,omitempty"`
		WirelessDisplayRequirePinForPairing                   *bool                                        `json:"wirelessDisplayRequirePinForPairing,omitempty"`
		Assignments                                           *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
		CreatedDateTime                                       *string                                      `json:"createdDateTime,omitempty"`
		Description                                           nullable.Type[string]                        `json:"description,omitempty"`
		DeviceManagementApplicabilityRuleDeviceMode           *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
		DeviceManagementApplicabilityRuleOsEdition            *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
		DeviceManagementApplicabilityRuleOsVersion            *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
		DeviceSettingStateSummaries                           *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
		DeviceStatusOverview                                  *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
		DeviceStatuses                                        *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
		DisplayName                                           *string                                      `json:"displayName,omitempty"`
		GroupAssignments                                      *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
		LastModifiedDateTime                                  *string                                      `json:"lastModifiedDateTime,omitempty"`
		RoleScopeTagIds                                       *[]string                                    `json:"roleScopeTagIds,omitempty"`
		SupportsScopeTags                                     *bool                                        `json:"supportsScopeTags,omitempty"`
		UserStatusOverview                                    *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
		UserStatuses                                          *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
		Version                                               *int64                                       `json:"version,omitempty"`
		Id                                                    *string                                      `json:"id,omitempty"`
		ODataId                                               *string                                      `json:"@odata.id,omitempty"`
		ODataType                                             *string                                      `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountsBlockAddingNonMicrosoftAccountEmail = decoded.AccountsBlockAddingNonMicrosoftAccountEmail
	s.ActivateAppsWithVoice = decoded.ActivateAppsWithVoice
	s.AntiTheftModeBlocked = decoded.AntiTheftModeBlocked
	s.AppManagementMSIAllowUserControlOverInstall = decoded.AppManagementMSIAllowUserControlOverInstall
	s.AppManagementMSIAlwaysInstallWithElevatedPrivileges = decoded.AppManagementMSIAlwaysInstallWithElevatedPrivileges
	s.AppManagementPackageFamilyNamesToLaunchAfterLogOn = decoded.AppManagementPackageFamilyNamesToLaunchAfterLogOn
	s.AppsAllowTrustedAppsSideloading = decoded.AppsAllowTrustedAppsSideloading
	s.AppsBlockWindowsStoreOriginatedApps = decoded.AppsBlockWindowsStoreOriginatedApps
	s.AuthenticationAllowSecondaryDevice = decoded.AuthenticationAllowSecondaryDevice
	s.AuthenticationPreferredAzureADTenantDomainName = decoded.AuthenticationPreferredAzureADTenantDomainName
	s.AuthenticationWebSignIn = decoded.AuthenticationWebSignIn
	s.BluetoothAllowedServices = decoded.BluetoothAllowedServices
	s.BluetoothBlockAdvertising = decoded.BluetoothBlockAdvertising
	s.BluetoothBlockDiscoverableMode = decoded.BluetoothBlockDiscoverableMode
	s.BluetoothBlockPrePairing = decoded.BluetoothBlockPrePairing
	s.BluetoothBlockPromptedProximalConnections = decoded.BluetoothBlockPromptedProximalConnections
	s.BluetoothBlocked = decoded.BluetoothBlocked
	s.CameraBlocked = decoded.CameraBlocked
	s.CellularBlockDataWhenRoaming = decoded.CellularBlockDataWhenRoaming
	s.CellularBlockVpn = decoded.CellularBlockVpn
	s.CellularBlockVpnWhenRoaming = decoded.CellularBlockVpnWhenRoaming
	s.CellularData = decoded.CellularData
	s.CertificatesBlockManualRootCertificateInstallation = decoded.CertificatesBlockManualRootCertificateInstallation
	s.ConfigureTimeZone = decoded.ConfigureTimeZone
	s.ConnectedDevicesServiceBlocked = decoded.ConnectedDevicesServiceBlocked
	s.CopyPasteBlocked = decoded.CopyPasteBlocked
	s.CortanaBlocked = decoded.CortanaBlocked
	s.CryptographyAllowFipsAlgorithmPolicy = decoded.CryptographyAllowFipsAlgorithmPolicy
	s.DataProtectionBlockDirectMemoryAccess = decoded.DataProtectionBlockDirectMemoryAccess
	s.DefenderBlockEndUserAccess = decoded.DefenderBlockEndUserAccess
	s.DefenderBlockOnAccessProtection = decoded.DefenderBlockOnAccessProtection
	s.DefenderCloudBlockLevel = decoded.DefenderCloudBlockLevel
	s.DefenderCloudExtendedTimeout = decoded.DefenderCloudExtendedTimeout
	s.DefenderCloudExtendedTimeoutInSeconds = decoded.DefenderCloudExtendedTimeoutInSeconds
	s.DefenderDaysBeforeDeletingQuarantinedMalware = decoded.DefenderDaysBeforeDeletingQuarantinedMalware
	s.DefenderDetectedMalwareActions = decoded.DefenderDetectedMalwareActions
	s.DefenderDisableCatchupFullScan = decoded.DefenderDisableCatchupFullScan
	s.DefenderDisableCatchupQuickScan = decoded.DefenderDisableCatchupQuickScan
	s.DefenderFileExtensionsToExclude = decoded.DefenderFileExtensionsToExclude
	s.DefenderFilesAndFoldersToExclude = decoded.DefenderFilesAndFoldersToExclude
	s.DefenderMonitorFileActivity = decoded.DefenderMonitorFileActivity
	s.DefenderPotentiallyUnwantedAppAction = decoded.DefenderPotentiallyUnwantedAppAction
	s.DefenderPotentiallyUnwantedAppActionSetting = decoded.DefenderPotentiallyUnwantedAppActionSetting
	s.DefenderProcessesToExclude = decoded.DefenderProcessesToExclude
	s.DefenderPromptForSampleSubmission = decoded.DefenderPromptForSampleSubmission
	s.DefenderRequireBehaviorMonitoring = decoded.DefenderRequireBehaviorMonitoring
	s.DefenderRequireCloudProtection = decoded.DefenderRequireCloudProtection
	s.DefenderRequireNetworkInspectionSystem = decoded.DefenderRequireNetworkInspectionSystem
	s.DefenderRequireRealTimeMonitoring = decoded.DefenderRequireRealTimeMonitoring
	s.DefenderScanArchiveFiles = decoded.DefenderScanArchiveFiles
	s.DefenderScanDownloads = decoded.DefenderScanDownloads
	s.DefenderScanIncomingMail = decoded.DefenderScanIncomingMail
	s.DefenderScanMappedNetworkDrivesDuringFullScan = decoded.DefenderScanMappedNetworkDrivesDuringFullScan
	s.DefenderScanMaxCpu = decoded.DefenderScanMaxCpu
	s.DefenderScanNetworkFiles = decoded.DefenderScanNetworkFiles
	s.DefenderScanRemovableDrivesDuringFullScan = decoded.DefenderScanRemovableDrivesDuringFullScan
	s.DefenderScanScriptsLoadedInInternetExplorer = decoded.DefenderScanScriptsLoadedInInternetExplorer
	s.DefenderScanType = decoded.DefenderScanType
	s.DefenderScheduleScanEnableLowCpuPriority = decoded.DefenderScheduleScanEnableLowCpuPriority
	s.DefenderScheduledQuickScanTime = decoded.DefenderScheduledQuickScanTime
	s.DefenderScheduledScanTime = decoded.DefenderScheduledScanTime
	s.DefenderSignatureUpdateIntervalInHours = decoded.DefenderSignatureUpdateIntervalInHours
	s.DefenderSubmitSamplesConsentType = decoded.DefenderSubmitSamplesConsentType
	s.DefenderSystemScanSchedule = decoded.DefenderSystemScanSchedule
	s.DeveloperUnlockSetting = decoded.DeveloperUnlockSetting
	s.DeviceManagementBlockFactoryResetOnMobile = decoded.DeviceManagementBlockFactoryResetOnMobile
	s.DeviceManagementBlockManualUnenroll = decoded.DeviceManagementBlockManualUnenroll
	s.DiagnosticsDataSubmissionMode = decoded.DiagnosticsDataSubmissionMode
	s.DisplayAppListWithGdiDPIScalingTurnedOff = decoded.DisplayAppListWithGdiDPIScalingTurnedOff
	s.DisplayAppListWithGdiDPIScalingTurnedOn = decoded.DisplayAppListWithGdiDPIScalingTurnedOn
	s.EdgeAllowStartPagesModification = decoded.EdgeAllowStartPagesModification
	s.EdgeBlockAccessToAboutFlags = decoded.EdgeBlockAccessToAboutFlags
	s.EdgeBlockAddressBarDropdown = decoded.EdgeBlockAddressBarDropdown
	s.EdgeBlockAutofill = decoded.EdgeBlockAutofill
	s.EdgeBlockCompatibilityList = decoded.EdgeBlockCompatibilityList
	s.EdgeBlockDeveloperTools = decoded.EdgeBlockDeveloperTools
	s.EdgeBlockEditFavorites = decoded.EdgeBlockEditFavorites
	s.EdgeBlockExtensions = decoded.EdgeBlockExtensions
	s.EdgeBlockFullScreenMode = decoded.EdgeBlockFullScreenMode
	s.EdgeBlockInPrivateBrowsing = decoded.EdgeBlockInPrivateBrowsing
	s.EdgeBlockJavaScript = decoded.EdgeBlockJavaScript
	s.EdgeBlockLiveTileDataCollection = decoded.EdgeBlockLiveTileDataCollection
	s.EdgeBlockPasswordManager = decoded.EdgeBlockPasswordManager
	s.EdgeBlockPopups = decoded.EdgeBlockPopups
	s.EdgeBlockPrelaunch = decoded.EdgeBlockPrelaunch
	s.EdgeBlockPrinting = decoded.EdgeBlockPrinting
	s.EdgeBlockSavingHistory = decoded.EdgeBlockSavingHistory
	s.EdgeBlockSearchEngineCustomization = decoded.EdgeBlockSearchEngineCustomization
	s.EdgeBlockSearchSuggestions = decoded.EdgeBlockSearchSuggestions
	s.EdgeBlockSendingDoNotTrackHeader = decoded.EdgeBlockSendingDoNotTrackHeader
	s.EdgeBlockSendingIntranetTrafficToInternetExplorer = decoded.EdgeBlockSendingIntranetTrafficToInternetExplorer
	s.EdgeBlockSideloadingExtensions = decoded.EdgeBlockSideloadingExtensions
	s.EdgeBlockTabPreloading = decoded.EdgeBlockTabPreloading
	s.EdgeBlockWebContentOnNewTabPage = decoded.EdgeBlockWebContentOnNewTabPage
	s.EdgeBlocked = decoded.EdgeBlocked
	s.EdgeClearBrowsingDataOnExit = decoded.EdgeClearBrowsingDataOnExit
	s.EdgeCookiePolicy = decoded.EdgeCookiePolicy
	s.EdgeDisableFirstRunPage = decoded.EdgeDisableFirstRunPage
	s.EdgeEnterpriseModeSiteListLocation = decoded.EdgeEnterpriseModeSiteListLocation
	s.EdgeFavoritesBarVisibility = decoded.EdgeFavoritesBarVisibility
	s.EdgeFavoritesListLocation = decoded.EdgeFavoritesListLocation
	s.EdgeFirstRunUrl = decoded.EdgeFirstRunUrl
	s.EdgeHomeButtonConfigurationEnabled = decoded.EdgeHomeButtonConfigurationEnabled
	s.EdgeHomepageUrls = decoded.EdgeHomepageUrls
	s.EdgeKioskModeRestriction = decoded.EdgeKioskModeRestriction
	s.EdgeKioskResetAfterIdleTimeInMinutes = decoded.EdgeKioskResetAfterIdleTimeInMinutes
	s.EdgeNewTabPageURL = decoded.EdgeNewTabPageURL
	s.EdgeOpensWith = decoded.EdgeOpensWith
	s.EdgePreventCertificateErrorOverride = decoded.EdgePreventCertificateErrorOverride
	s.EdgeRequireSmartScreen = decoded.EdgeRequireSmartScreen
	s.EdgeRequiredExtensionPackageFamilyNames = decoded.EdgeRequiredExtensionPackageFamilyNames
	s.EdgeSendIntranetTrafficToInternetExplorer = decoded.EdgeSendIntranetTrafficToInternetExplorer
	s.EdgeShowMessageWhenOpeningInternetExplorerSites = decoded.EdgeShowMessageWhenOpeningInternetExplorerSites
	s.EdgeSyncFavoritesWithInternetExplorer = decoded.EdgeSyncFavoritesWithInternetExplorer
	s.EdgeTelemetryForMicrosoft365Analytics = decoded.EdgeTelemetryForMicrosoft365Analytics
	s.EnableAutomaticRedeployment = decoded.EnableAutomaticRedeployment
	s.EnergySaverOnBatteryThresholdPercentage = decoded.EnergySaverOnBatteryThresholdPercentage
	s.EnergySaverPluggedInThresholdPercentage = decoded.EnergySaverPluggedInThresholdPercentage
	s.EnterpriseCloudPrintDiscoveryEndPoint = decoded.EnterpriseCloudPrintDiscoveryEndPoint
	s.EnterpriseCloudPrintDiscoveryMaxLimit = decoded.EnterpriseCloudPrintDiscoveryMaxLimit
	s.EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier = decoded.EnterpriseCloudPrintMopriaDiscoveryResourceIdentifier
	s.EnterpriseCloudPrintOAuthAuthority = decoded.EnterpriseCloudPrintOAuthAuthority
	s.EnterpriseCloudPrintOAuthClientIdentifier = decoded.EnterpriseCloudPrintOAuthClientIdentifier
	s.EnterpriseCloudPrintResourceIdentifier = decoded.EnterpriseCloudPrintResourceIdentifier
	s.ExperienceBlockDeviceDiscovery = decoded.ExperienceBlockDeviceDiscovery
	s.ExperienceBlockErrorDialogWhenNoSIM = decoded.ExperienceBlockErrorDialogWhenNoSIM
	s.ExperienceBlockTaskSwitcher = decoded.ExperienceBlockTaskSwitcher
	s.ExperienceDoNotSyncBrowserSettings = decoded.ExperienceDoNotSyncBrowserSettings
	s.FindMyFiles = decoded.FindMyFiles
	s.GameDvrBlocked = decoded.GameDvrBlocked
	s.InkWorkspaceAccess = decoded.InkWorkspaceAccess
	s.InkWorkspaceAccessState = decoded.InkWorkspaceAccessState
	s.InkWorkspaceBlockSuggestedApps = decoded.InkWorkspaceBlockSuggestedApps
	s.InternetSharingBlocked = decoded.InternetSharingBlocked
	s.LocationServicesBlocked = decoded.LocationServicesBlocked
	s.LockScreenActivateAppsWithVoice = decoded.LockScreenActivateAppsWithVoice
	s.LockScreenAllowTimeoutConfiguration = decoded.LockScreenAllowTimeoutConfiguration
	s.LockScreenBlockActionCenterNotifications = decoded.LockScreenBlockActionCenterNotifications
	s.LockScreenBlockCortana = decoded.LockScreenBlockCortana
	s.LockScreenBlockToastNotifications = decoded.LockScreenBlockToastNotifications
	s.LockScreenTimeoutInSeconds = decoded.LockScreenTimeoutInSeconds
	s.LogonBlockFastUserSwitching = decoded.LogonBlockFastUserSwitching
	s.MessagingBlockMMS = decoded.MessagingBlockMMS
	s.MessagingBlockRichCommunicationServices = decoded.MessagingBlockRichCommunicationServices
	s.MessagingBlockSync = decoded.MessagingBlockSync
	s.MicrosoftAccountBlockSettingsSync = decoded.MicrosoftAccountBlockSettingsSync
	s.MicrosoftAccountBlocked = decoded.MicrosoftAccountBlocked
	s.MicrosoftAccountSignInAssistantSettings = decoded.MicrosoftAccountSignInAssistantSettings
	s.NetworkProxyApplySettingsDeviceWide = decoded.NetworkProxyApplySettingsDeviceWide
	s.NetworkProxyAutomaticConfigurationUrl = decoded.NetworkProxyAutomaticConfigurationUrl
	s.NetworkProxyDisableAutoDetect = decoded.NetworkProxyDisableAutoDetect
	s.NetworkProxyServer = decoded.NetworkProxyServer
	s.NfcBlocked = decoded.NfcBlocked
	s.OneDriveDisableFileSync = decoded.OneDriveDisableFileSync
	s.PasswordBlockSimple = decoded.PasswordBlockSimple
	s.PasswordExpirationDays = decoded.PasswordExpirationDays
	s.PasswordMinimumAgeInDays = decoded.PasswordMinimumAgeInDays
	s.PasswordMinimumCharacterSetCount = decoded.PasswordMinimumCharacterSetCount
	s.PasswordMinimumLength = decoded.PasswordMinimumLength
	s.PasswordMinutesOfInactivityBeforeScreenTimeout = decoded.PasswordMinutesOfInactivityBeforeScreenTimeout
	s.PasswordPreviousPasswordBlockCount = decoded.PasswordPreviousPasswordBlockCount
	s.PasswordRequireWhenResumeFromIdleState = decoded.PasswordRequireWhenResumeFromIdleState
	s.PasswordRequired = decoded.PasswordRequired
	s.PasswordRequiredType = decoded.PasswordRequiredType
	s.PasswordSignInFailureCountBeforeFactoryReset = decoded.PasswordSignInFailureCountBeforeFactoryReset
	s.PersonalizationDesktopImageUrl = decoded.PersonalizationDesktopImageUrl
	s.PersonalizationLockScreenImageUrl = decoded.PersonalizationLockScreenImageUrl
	s.PowerButtonActionOnBattery = decoded.PowerButtonActionOnBattery
	s.PowerButtonActionPluggedIn = decoded.PowerButtonActionPluggedIn
	s.PowerHybridSleepOnBattery = decoded.PowerHybridSleepOnBattery
	s.PowerHybridSleepPluggedIn = decoded.PowerHybridSleepPluggedIn
	s.PowerLidCloseActionOnBattery = decoded.PowerLidCloseActionOnBattery
	s.PowerLidCloseActionPluggedIn = decoded.PowerLidCloseActionPluggedIn
	s.PowerSleepButtonActionOnBattery = decoded.PowerSleepButtonActionOnBattery
	s.PowerSleepButtonActionPluggedIn = decoded.PowerSleepButtonActionPluggedIn
	s.PrinterBlockAddition = decoded.PrinterBlockAddition
	s.PrinterDefaultName = decoded.PrinterDefaultName
	s.PrinterNames = decoded.PrinterNames
	s.PrivacyAccessControls = decoded.PrivacyAccessControls
	s.PrivacyAdvertisingId = decoded.PrivacyAdvertisingId
	s.PrivacyAutoAcceptPairingAndConsentPrompts = decoded.PrivacyAutoAcceptPairingAndConsentPrompts
	s.PrivacyBlockActivityFeed = decoded.PrivacyBlockActivityFeed
	s.PrivacyBlockInputPersonalization = decoded.PrivacyBlockInputPersonalization
	s.PrivacyBlockPublishUserActivities = decoded.PrivacyBlockPublishUserActivities
	s.PrivacyDisableLaunchExperience = decoded.PrivacyDisableLaunchExperience
	s.ResetProtectionModeBlocked = decoded.ResetProtectionModeBlocked
	s.SafeSearchFilter = decoded.SafeSearchFilter
	s.ScreenCaptureBlocked = decoded.ScreenCaptureBlocked
	s.SearchBlockDiacritics = decoded.SearchBlockDiacritics
	s.SearchBlockWebResults = decoded.SearchBlockWebResults
	s.SearchDisableAutoLanguageDetection = decoded.SearchDisableAutoLanguageDetection
	s.SearchDisableIndexerBackoff = decoded.SearchDisableIndexerBackoff
	s.SearchDisableIndexingEncryptedItems = decoded.SearchDisableIndexingEncryptedItems
	s.SearchDisableIndexingRemovableDrive = decoded.SearchDisableIndexingRemovableDrive
	s.SearchDisableLocation = decoded.SearchDisableLocation
	s.SearchDisableUseLocation = decoded.SearchDisableUseLocation
	s.SearchEnableAutomaticIndexSizeManangement = decoded.SearchEnableAutomaticIndexSizeManangement
	s.SearchEnableRemoteQueries = decoded.SearchEnableRemoteQueries
	s.SecurityBlockAzureADJoinedDevicesAutoEncryption = decoded.SecurityBlockAzureADJoinedDevicesAutoEncryption
	s.SettingsBlockAccountsPage = decoded.SettingsBlockAccountsPage
	s.SettingsBlockAddProvisioningPackage = decoded.SettingsBlockAddProvisioningPackage
	s.SettingsBlockAppsPage = decoded.SettingsBlockAppsPage
	s.SettingsBlockChangeLanguage = decoded.SettingsBlockChangeLanguage
	s.SettingsBlockChangePowerSleep = decoded.SettingsBlockChangePowerSleep
	s.SettingsBlockChangeRegion = decoded.SettingsBlockChangeRegion
	s.SettingsBlockChangeSystemTime = decoded.SettingsBlockChangeSystemTime
	s.SettingsBlockDevicesPage = decoded.SettingsBlockDevicesPage
	s.SettingsBlockEaseOfAccessPage = decoded.SettingsBlockEaseOfAccessPage
	s.SettingsBlockEditDeviceName = decoded.SettingsBlockEditDeviceName
	s.SettingsBlockGamingPage = decoded.SettingsBlockGamingPage
	s.SettingsBlockNetworkInternetPage = decoded.SettingsBlockNetworkInternetPage
	s.SettingsBlockPersonalizationPage = decoded.SettingsBlockPersonalizationPage
	s.SettingsBlockPrivacyPage = decoded.SettingsBlockPrivacyPage
	s.SettingsBlockRemoveProvisioningPackage = decoded.SettingsBlockRemoveProvisioningPackage
	s.SettingsBlockSettingsApp = decoded.SettingsBlockSettingsApp
	s.SettingsBlockSystemPage = decoded.SettingsBlockSystemPage
	s.SettingsBlockTimeLanguagePage = decoded.SettingsBlockTimeLanguagePage
	s.SettingsBlockUpdateSecurityPage = decoded.SettingsBlockUpdateSecurityPage
	s.SharedUserAppDataAllowed = decoded.SharedUserAppDataAllowed
	s.SmartScreenAppInstallControl = decoded.SmartScreenAppInstallControl
	s.SmartScreenBlockPromptOverride = decoded.SmartScreenBlockPromptOverride
	s.SmartScreenBlockPromptOverrideForFiles = decoded.SmartScreenBlockPromptOverrideForFiles
	s.SmartScreenEnableAppInstallControl = decoded.SmartScreenEnableAppInstallControl
	s.StartBlockUnpinningAppsFromTaskbar = decoded.StartBlockUnpinningAppsFromTaskbar
	s.StartMenuAppListVisibility = decoded.StartMenuAppListVisibility
	s.StartMenuHideChangeAccountSettings = decoded.StartMenuHideChangeAccountSettings
	s.StartMenuHideFrequentlyUsedApps = decoded.StartMenuHideFrequentlyUsedApps
	s.StartMenuHideHibernate = decoded.StartMenuHideHibernate
	s.StartMenuHideLock = decoded.StartMenuHideLock
	s.StartMenuHidePowerButton = decoded.StartMenuHidePowerButton
	s.StartMenuHideRecentJumpLists = decoded.StartMenuHideRecentJumpLists
	s.StartMenuHideRecentlyAddedApps = decoded.StartMenuHideRecentlyAddedApps
	s.StartMenuHideRestartOptions = decoded.StartMenuHideRestartOptions
	s.StartMenuHideShutDown = decoded.StartMenuHideShutDown
	s.StartMenuHideSignOut = decoded.StartMenuHideSignOut
	s.StartMenuHideSleep = decoded.StartMenuHideSleep
	s.StartMenuHideSwitchAccount = decoded.StartMenuHideSwitchAccount
	s.StartMenuHideUserTile = decoded.StartMenuHideUserTile
	s.StartMenuLayoutEdgeAssetsXml = decoded.StartMenuLayoutEdgeAssetsXml
	s.StartMenuLayoutXml = decoded.StartMenuLayoutXml
	s.StartMenuMode = decoded.StartMenuMode
	s.StartMenuPinnedFolderDocuments = decoded.StartMenuPinnedFolderDocuments
	s.StartMenuPinnedFolderDownloads = decoded.StartMenuPinnedFolderDownloads
	s.StartMenuPinnedFolderFileExplorer = decoded.StartMenuPinnedFolderFileExplorer
	s.StartMenuPinnedFolderHomeGroup = decoded.StartMenuPinnedFolderHomeGroup
	s.StartMenuPinnedFolderMusic = decoded.StartMenuPinnedFolderMusic
	s.StartMenuPinnedFolderNetwork = decoded.StartMenuPinnedFolderNetwork
	s.StartMenuPinnedFolderPersonalFolder = decoded.StartMenuPinnedFolderPersonalFolder
	s.StartMenuPinnedFolderPictures = decoded.StartMenuPinnedFolderPictures
	s.StartMenuPinnedFolderSettings = decoded.StartMenuPinnedFolderSettings
	s.StartMenuPinnedFolderVideos = decoded.StartMenuPinnedFolderVideos
	s.StorageBlockRemovableStorage = decoded.StorageBlockRemovableStorage
	s.StorageRequireMobileDeviceEncryption = decoded.StorageRequireMobileDeviceEncryption
	s.StorageRestrictAppDataToSystemVolume = decoded.StorageRestrictAppDataToSystemVolume
	s.StorageRestrictAppInstallToSystemVolume = decoded.StorageRestrictAppInstallToSystemVolume
	s.SystemTelemetryProxyServer = decoded.SystemTelemetryProxyServer
	s.TaskManagerBlockEndTask = decoded.TaskManagerBlockEndTask
	s.TenantLockdownRequireNetworkDuringOutOfBoxExperience = decoded.TenantLockdownRequireNetworkDuringOutOfBoxExperience
	s.UninstallBuiltInApps = decoded.UninstallBuiltInApps
	s.UsbBlocked = decoded.UsbBlocked
	s.VoiceRecordingBlocked = decoded.VoiceRecordingBlocked
	s.WebRtcBlockLocalhostIPAddress = decoded.WebRtcBlockLocalhostIPAddress
	s.WiFiBlockAutomaticConnectHotspots = decoded.WiFiBlockAutomaticConnectHotspots
	s.WiFiBlockManualConfiguration = decoded.WiFiBlockManualConfiguration
	s.WiFiBlocked = decoded.WiFiBlocked
	s.WiFiScanInterval = decoded.WiFiScanInterval
	s.Windows10AppsForceUpdateSchedule = decoded.Windows10AppsForceUpdateSchedule
	s.WindowsSpotlightBlockConsumerSpecificFeatures = decoded.WindowsSpotlightBlockConsumerSpecificFeatures
	s.WindowsSpotlightBlockOnActionCenter = decoded.WindowsSpotlightBlockOnActionCenter
	s.WindowsSpotlightBlockTailoredExperiences = decoded.WindowsSpotlightBlockTailoredExperiences
	s.WindowsSpotlightBlockThirdPartyNotifications = decoded.WindowsSpotlightBlockThirdPartyNotifications
	s.WindowsSpotlightBlockWelcomeExperience = decoded.WindowsSpotlightBlockWelcomeExperience
	s.WindowsSpotlightBlockWindowsTips = decoded.WindowsSpotlightBlockWindowsTips
	s.WindowsSpotlightBlocked = decoded.WindowsSpotlightBlocked
	s.WindowsSpotlightConfigureOnLockScreen = decoded.WindowsSpotlightConfigureOnLockScreen
	s.WindowsStoreBlockAutoUpdate = decoded.WindowsStoreBlockAutoUpdate
	s.WindowsStoreBlocked = decoded.WindowsStoreBlocked
	s.WindowsStoreEnablePrivateStoreOnly = decoded.WindowsStoreEnablePrivateStoreOnly
	s.WirelessDisplayBlockProjectionToThisDevice = decoded.WirelessDisplayBlockProjectionToThisDevice
	s.WirelessDisplayBlockUserInputFromReceiver = decoded.WirelessDisplayBlockUserInputFromReceiver
	s.WirelessDisplayRequirePinForPairing = decoded.WirelessDisplayRequirePinForPairing
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
		return fmt.Errorf("unmarshaling Windows10GeneralConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["edgeHomeButtonConfiguration"]; ok {
		impl, err := UnmarshalEdgeHomeButtonConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EdgeHomeButtonConfiguration' for 'Windows10GeneralConfiguration': %+v", err)
		}
		s.EdgeHomeButtonConfiguration = impl
	}

	if v, ok := temp["edgeSearchEngine"]; ok {
		impl, err := UnmarshalEdgeSearchEngineBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EdgeSearchEngine' for 'Windows10GeneralConfiguration': %+v", err)
		}
		s.EdgeSearchEngine = impl
	}

	return nil
}
