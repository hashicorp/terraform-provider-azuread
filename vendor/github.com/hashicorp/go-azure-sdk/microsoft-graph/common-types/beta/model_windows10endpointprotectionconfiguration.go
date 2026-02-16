package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeviceConfiguration = Windows10EndpointProtectionConfiguration{}

type Windows10EndpointProtectionConfiguration struct {
	// Possible values of AppLocker Application Control Types
	AppLockerApplicationControl *AppLockerApplicationControlType `json:"appLockerApplicationControl,omitempty"`

	// Gets or sets whether applications inside Microsoft Defender Application Guard can access the device’s camera and
	// microphone.
	ApplicationGuardAllowCameraMicrophoneRedirection nullable.Type[bool] `json:"applicationGuardAllowCameraMicrophoneRedirection,omitempty"`

	// Allow users to download files from Edge in the application guard container and save them on the host file system
	ApplicationGuardAllowFileSaveOnHost *bool `json:"applicationGuardAllowFileSaveOnHost,omitempty"`

	// Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
	ApplicationGuardAllowPersistence *bool `json:"applicationGuardAllowPersistence,omitempty"`

	// Allow printing to Local Printers from Container
	ApplicationGuardAllowPrintToLocalPrinters *bool `json:"applicationGuardAllowPrintToLocalPrinters,omitempty"`

	// Allow printing to Network Printers from Container
	ApplicationGuardAllowPrintToNetworkPrinters *bool `json:"applicationGuardAllowPrintToNetworkPrinters,omitempty"`

	// Allow printing to PDF from Container
	ApplicationGuardAllowPrintToPDF *bool `json:"applicationGuardAllowPrintToPDF,omitempty"`

	// Allow printing to XPS from Container
	ApplicationGuardAllowPrintToXPS *bool `json:"applicationGuardAllowPrintToXPS,omitempty"`

	// Allow application guard to use virtual GPU
	ApplicationGuardAllowVirtualGPU *bool `json:"applicationGuardAllowVirtualGPU,omitempty"`

	// Possible values for applicationGuardBlockClipboardSharingType
	ApplicationGuardBlockClipboardSharing *ApplicationGuardBlockClipboardSharingType `json:"applicationGuardBlockClipboardSharing,omitempty"`

	// Possible values for applicationGuardBlockFileTransfer
	ApplicationGuardBlockFileTransfer *ApplicationGuardBlockFileTransferType `json:"applicationGuardBlockFileTransfer,omitempty"`

	// Block enterprise sites to load non-enterprise content, such as third party plug-ins
	ApplicationGuardBlockNonEnterpriseContent *bool `json:"applicationGuardBlockNonEnterpriseContent,omitempty"`

	// Allows certain device level Root Certificates to be shared with the Microsoft Defender Application Guard container.
	ApplicationGuardCertificateThumbprints *[]string `json:"applicationGuardCertificateThumbprints,omitempty"`

	// Enable Windows Defender Application Guard
	ApplicationGuardEnabled *bool `json:"applicationGuardEnabled,omitempty"`

	// Possible values for ApplicationGuardEnabledOptions
	ApplicationGuardEnabledOptions *ApplicationGuardEnabledOptions `json:"applicationGuardEnabledOptions,omitempty"`

	// Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user
	// login-logoff, use of privilege rights, software installation, system changes, etc.)
	ApplicationGuardForceAuditing *bool `json:"applicationGuardForceAuditing,omitempty"`

	// Allows the admin to allow standard users to enable encrpytion during Azure AD Join.
	BitLockerAllowStandardUserEncryption *bool `json:"bitLockerAllowStandardUserEncryption,omitempty"`

	// Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
	BitLockerDisableWarningForOtherDiskEncryption *bool `json:"bitLockerDisableWarningForOtherDiskEncryption,omitempty"`

	// Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
	BitLockerEnableStorageCardEncryptionOnMobile *bool `json:"bitLockerEnableStorageCardEncryptionOnMobile,omitempty"`

	// Allows the admin to require encryption to be turned on using BitLocker.
	BitLockerEncryptDevice *bool `json:"bitLockerEncryptDevice,omitempty"`

	// BitLocker Fixed Drive Policy.
	BitLockerFixedDrivePolicy *BitLockerFixedDrivePolicy `json:"bitLockerFixedDrivePolicy,omitempty"`

	// BitLocker recovery password rotation type
	BitLockerRecoveryPasswordRotation *BitLockerRecoveryPasswordRotationType `json:"bitLockerRecoveryPasswordRotation,omitempty"`

	// BitLocker Removable Drive Policy.
	BitLockerRemovableDrivePolicy *BitLockerRemovableDrivePolicy `json:"bitLockerRemovableDrivePolicy,omitempty"`

	// BitLocker System Drive Policy.
	BitLockerSystemDrivePolicy *BitLockerSystemDrivePolicy `json:"bitLockerSystemDrivePolicy,omitempty"`

	// List of folder paths to be added to the list of protected folders
	DefenderAdditionalGuardedFolders *[]string `json:"defenderAdditionalGuardedFolders,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderAdobeReaderLaunchChildProcess *DefenderProtectionType `json:"defenderAdobeReaderLaunchChildProcess,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderAdvancedRansomewareProtectionType *DefenderProtectionType `json:"defenderAdvancedRansomewareProtectionType,omitempty"`

	// Allows or disallows Windows Defender Behavior Monitoring functionality.
	DefenderAllowBehaviorMonitoring nullable.Type[bool] `json:"defenderAllowBehaviorMonitoring,omitempty"`

	// To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft
	// will analyze that information, learn more about problems affecting you and other customers, and offer improved
	// solutions.
	DefenderAllowCloudProtection nullable.Type[bool] `json:"defenderAllowCloudProtection,omitempty"`

	// Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will
	// also be suppressed.
	DefenderAllowEndUserAccess nullable.Type[bool] `json:"defenderAllowEndUserAccess,omitempty"`

	// Allows or disallows Windows Defender Intrusion Prevention functionality.
	DefenderAllowIntrusionPreventionSystem nullable.Type[bool] `json:"defenderAllowIntrusionPreventionSystem,omitempty"`

	// Allows or disallows Windows Defender On Access Protection functionality.
	DefenderAllowOnAccessProtection nullable.Type[bool] `json:"defenderAllowOnAccessProtection,omitempty"`

	// Allows or disallows Windows Defender Realtime Monitoring functionality.
	DefenderAllowRealTimeMonitoring nullable.Type[bool] `json:"defenderAllowRealTimeMonitoring,omitempty"`

	// Allows or disallows scanning of archives.
	DefenderAllowScanArchiveFiles nullable.Type[bool] `json:"defenderAllowScanArchiveFiles,omitempty"`

	// Allows or disallows Windows Defender IOAVP Protection functionality.
	DefenderAllowScanDownloads nullable.Type[bool] `json:"defenderAllowScanDownloads,omitempty"`

	// Allows or disallows a scanning of network files.
	DefenderAllowScanNetworkFiles nullable.Type[bool] `json:"defenderAllowScanNetworkFiles,omitempty"`

	// Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
	DefenderAllowScanRemovableDrivesDuringFullScan nullable.Type[bool] `json:"defenderAllowScanRemovableDrivesDuringFullScan,omitempty"`

	// Allows or disallows Windows Defender Script Scanning functionality.
	DefenderAllowScanScriptsLoadedInInternetExplorer nullable.Type[bool] `json:"defenderAllowScanScriptsLoadedInInternetExplorer,omitempty"`

	// List of exe files and folders to be excluded from attack surface reduction rules
	DefenderAttackSurfaceReductionExcludedPaths *[]string `json:"defenderAttackSurfaceReductionExcludedPaths,omitempty"`

	// Allows or disallows user access to the Windows Defender UI. If disallowed, all Windows Defender notifications will
	// also be suppressed.
	DefenderBlockEndUserAccess nullable.Type[bool] `json:"defenderBlockEndUserAccess,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderBlockPersistenceThroughWmiType *DefenderAttackSurfaceType `json:"defenderBlockPersistenceThroughWmiType,omitempty"`

	// This policy setting allows you to manage whether a check for new virus and spyware definitions will occur before
	// running a scan.
	DefenderCheckForSignaturesBeforeRunningScan nullable.Type[bool] `json:"defenderCheckForSignaturesBeforeRunningScan,omitempty"`

	// Added in Windows 10, version 1709. This policy setting determines how aggressive Windows Defender Antivirus will be
	// in blocking and scanning suspicious files. Value type is integer. This feature requires the 'Join Microsoft MAPS'
	// setting enabled in order to function. Possible values are: notConfigured, high, highPlus, zeroTolerance.
	DefenderCloudBlockLevel *DefenderCloudBlockLevelType `json:"defenderCloudBlockLevel,omitempty"`

	// Added in Windows 10, version 1709. This feature allows Windows Defender Antivirus to block a suspicious file for up
	// to 60 seconds, and scan it in the cloud to make sure it's safe. Value type is integer, range is 0 - 50. This feature
	// depends on three other MAPS settings the must all be enabled- 'Configure the 'Block at First Sight' feature; 'Join
	// Microsoft MAPS'; 'Send file samples when further analysis is required'. Valid values 0 to 50
	DefenderCloudExtendedTimeoutInSeconds nullable.Type[int64] `json:"defenderCloudExtendedTimeoutInSeconds,omitempty"`

	// Time period (in days) that quarantine items will be stored on the system. Valid values 0 to 90
	DefenderDaysBeforeDeletingQuarantinedMalware nullable.Type[int64] `json:"defenderDaysBeforeDeletingQuarantinedMalware,omitempty"`

	// Allows an administrator to specify any valid threat severity levels and the corresponding default action ID to take.
	DefenderDetectedMalwareActions *DefenderDetectedMalwareActions `json:"defenderDetectedMalwareActions,omitempty"`

	// Allows or disallows Windows Defender Behavior Monitoring functionality.
	DefenderDisableBehaviorMonitoring nullable.Type[bool] `json:"defenderDisableBehaviorMonitoring,omitempty"`

	// This policy setting allows you to configure catch-up scans for scheduled full scans. A catch-up scan is a scan that
	// is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the
	// computer was turned off at the scheduled time.
	DefenderDisableCatchupFullScan nullable.Type[bool] `json:"defenderDisableCatchupFullScan,omitempty"`

	// This policy setting allows you to configure catch-up scans for scheduled quick scans. A catch-up scan is a scan that
	// is initiated because a regularly scheduled scan was missed. Usually these scheduled scans are missed because the
	// computer was turned off at the scheduled time.
	DefenderDisableCatchupQuickScan nullable.Type[bool] `json:"defenderDisableCatchupQuickScan,omitempty"`

	// To best protect your PC, Windows Defender will send information to Microsoft about any problems it finds. Microsoft
	// will analyze that information, learn more about problems affecting you and other customers, and offer improved
	// solutions.
	DefenderDisableCloudProtection nullable.Type[bool] `json:"defenderDisableCloudProtection,omitempty"`

	// Allows or disallows Windows Defender Intrusion Prevention functionality.
	DefenderDisableIntrusionPreventionSystem nullable.Type[bool] `json:"defenderDisableIntrusionPreventionSystem,omitempty"`

	// Allows or disallows Windows Defender On Access Protection functionality.
	DefenderDisableOnAccessProtection nullable.Type[bool] `json:"defenderDisableOnAccessProtection,omitempty"`

	// Allows or disallows Windows Defender Realtime Monitoring functionality.
	DefenderDisableRealTimeMonitoring nullable.Type[bool] `json:"defenderDisableRealTimeMonitoring,omitempty"`

	// Allows or disallows scanning of archives.
	DefenderDisableScanArchiveFiles nullable.Type[bool] `json:"defenderDisableScanArchiveFiles,omitempty"`

	// Allows or disallows Windows Defender IOAVP Protection functionality.
	DefenderDisableScanDownloads nullable.Type[bool] `json:"defenderDisableScanDownloads,omitempty"`

	// Allows or disallows a scanning of network files.
	DefenderDisableScanNetworkFiles nullable.Type[bool] `json:"defenderDisableScanNetworkFiles,omitempty"`

	// Allows or disallows a full scan of removable drives. During a quick scan, removable drives may still be scanned.
	DefenderDisableScanRemovableDrivesDuringFullScan nullable.Type[bool] `json:"defenderDisableScanRemovableDrivesDuringFullScan,omitempty"`

	// Allows or disallows Windows Defender Script Scanning functionality.
	DefenderDisableScanScriptsLoadedInInternetExplorer nullable.Type[bool] `json:"defenderDisableScanScriptsLoadedInInternetExplorer,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderEmailContentExecution *DefenderProtectionType `json:"defenderEmailContentExecution,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderEmailContentExecutionType *DefenderAttackSurfaceType `json:"defenderEmailContentExecutionType,omitempty"`

	// This policy setting allows you to enable or disable low CPU priority for scheduled scans.
	DefenderEnableLowCpuPriority nullable.Type[bool] `json:"defenderEnableLowCpuPriority,omitempty"`

	// Allows or disallows scanning of email.
	DefenderEnableScanIncomingMail nullable.Type[bool] `json:"defenderEnableScanIncomingMail,omitempty"`

	// Allows or disallows a full scan of mapped network drives.
	DefenderEnableScanMappedNetworkDrivesDuringFullScan nullable.Type[bool] `json:"defenderEnableScanMappedNetworkDrivesDuringFullScan,omitempty"`

	// Xml content containing information regarding exploit protection details.
	DefenderExploitProtectionXml nullable.Type[string] `json:"defenderExploitProtectionXml,omitempty"`

	// Name of the file from which DefenderExploitProtectionXml was obtained.
	DefenderExploitProtectionXmlFileName nullable.Type[string] `json:"defenderExploitProtectionXmlFileName,omitempty"`

	// File extensions to exclude from scans and real time protection.
	DefenderFileExtensionsToExclude *[]string `json:"defenderFileExtensionsToExclude,omitempty"`

	// Files and folder to exclude from scans and real time protection.
	DefenderFilesAndFoldersToExclude *[]string `json:"defenderFilesAndFoldersToExclude,omitempty"`

	// Possible values of Folder Protection
	DefenderGuardMyFoldersType *FolderProtectionType `json:"defenderGuardMyFoldersType,omitempty"`

	// List of paths to exe that are allowed to access protected folders
	DefenderGuardedFoldersAllowedAppPaths *[]string `json:"defenderGuardedFoldersAllowedAppPaths,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderNetworkProtectionType *DefenderProtectionType `json:"defenderNetworkProtectionType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderOfficeAppsExecutableContentCreationOrLaunch *DefenderProtectionType `json:"defenderOfficeAppsExecutableContentCreationOrLaunch,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderOfficeAppsExecutableContentCreationOrLaunchType *DefenderAttackSurfaceType `json:"defenderOfficeAppsExecutableContentCreationOrLaunchType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderOfficeAppsLaunchChildProcess *DefenderProtectionType `json:"defenderOfficeAppsLaunchChildProcess,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderOfficeAppsLaunchChildProcessType *DefenderAttackSurfaceType `json:"defenderOfficeAppsLaunchChildProcessType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderOfficeAppsOtherProcessInjection *DefenderProtectionType `json:"defenderOfficeAppsOtherProcessInjection,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderOfficeAppsOtherProcessInjectionType *DefenderAttackSurfaceType `json:"defenderOfficeAppsOtherProcessInjectionType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderOfficeCommunicationAppsLaunchChildProcess *DefenderProtectionType `json:"defenderOfficeCommunicationAppsLaunchChildProcess,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderOfficeMacroCodeAllowWin32Imports *DefenderProtectionType `json:"defenderOfficeMacroCodeAllowWin32Imports,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderOfficeMacroCodeAllowWin32ImportsType *DefenderAttackSurfaceType `json:"defenderOfficeMacroCodeAllowWin32ImportsType,omitempty"`

	// Added in Windows 10, version 1607. Specifies the level of detection for potentially unwanted applications (PUAs).
	// Windows Defender alerts you when potentially unwanted software is being downloaded or attempts to install itself on
	// your computer. Possible values are: userDefined, enable, auditMode, warn, notConfigured.
	DefenderPotentiallyUnwantedAppAction *DefenderProtectionType `json:"defenderPotentiallyUnwantedAppAction,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderPreventCredentialStealingType *DefenderProtectionType `json:"defenderPreventCredentialStealingType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderProcessCreation *DefenderProtectionType `json:"defenderProcessCreation,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderProcessCreationType *DefenderAttackSurfaceType `json:"defenderProcessCreationType,omitempty"`

	// Processes to exclude from scans and real time protection.
	DefenderProcessesToExclude *[]string `json:"defenderProcessesToExclude,omitempty"`

	// Controls which sets of files should be monitored. Possible values are: monitorAllFiles, monitorIncomingFilesOnly,
	// monitorOutgoingFilesOnly.
	DefenderScanDirection *DefenderRealtimeScanDirection `json:"defenderScanDirection,omitempty"`

	// Represents the average CPU load factor for the Windows Defender scan (in percent). The default value is 50. Valid
	// values 0 to 100
	DefenderScanMaxCpuPercentage nullable.Type[int64] `json:"defenderScanMaxCpuPercentage,omitempty"`

	// Selects whether to perform a quick scan or full scan. Possible values are: userDefined, disabled, quick, full.
	DefenderScanType *DefenderScanType `json:"defenderScanType,omitempty"`

	// Selects the time of day that the Windows Defender quick scan should run. For example, a value of 0=12:00AM, a value
	// of 60=1:00AM, a value of 120=2:00, and so on, up to a value of 1380=11:00PM. The default value is 120
	DefenderScheduledQuickScanTime nullable.Type[string] `json:"defenderScheduledQuickScanTime,omitempty"`

	// Selects the day that the Windows Defender scan should run. Possible values are: userDefined, everyday, sunday,
	// monday, tuesday, wednesday, thursday, friday, saturday, noScheduledScan.
	DefenderScheduledScanDay *WeeklySchedule `json:"defenderScheduledScanDay,omitempty"`

	// Selects the time of day that the Windows Defender scan should run.
	DefenderScheduledScanTime nullable.Type[string] `json:"defenderScheduledScanTime,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderScriptDownloadedPayloadExecution *DefenderProtectionType `json:"defenderScriptDownloadedPayloadExecution,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderScriptDownloadedPayloadExecutionType *DefenderAttackSurfaceType `json:"defenderScriptDownloadedPayloadExecutionType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderScriptObfuscatedMacroCode *DefenderProtectionType `json:"defenderScriptObfuscatedMacroCode,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderScriptObfuscatedMacroCodeType *DefenderAttackSurfaceType `json:"defenderScriptObfuscatedMacroCodeType,omitempty"`

	// Indicates whether or not to block user from overriding Exploit Protection settings.
	DefenderSecurityCenterBlockExploitProtectionOverride *bool `json:"defenderSecurityCenterBlockExploitProtectionOverride,omitempty"`

	// Used to disable the display of the account protection area.
	DefenderSecurityCenterDisableAccountUI nullable.Type[bool] `json:"defenderSecurityCenterDisableAccountUI,omitempty"`

	// Used to disable the display of the app and browser protection area.
	DefenderSecurityCenterDisableAppBrowserUI nullable.Type[bool] `json:"defenderSecurityCenterDisableAppBrowserUI,omitempty"`

	// Used to disable the display of the Clear TPM button.
	DefenderSecurityCenterDisableClearTpmUI nullable.Type[bool] `json:"defenderSecurityCenterDisableClearTpmUI,omitempty"`

	// Used to disable the display of the family options area.
	DefenderSecurityCenterDisableFamilyUI nullable.Type[bool] `json:"defenderSecurityCenterDisableFamilyUI,omitempty"`

	// Used to disable the display of the hardware protection area.
	DefenderSecurityCenterDisableHardwareUI nullable.Type[bool] `json:"defenderSecurityCenterDisableHardwareUI,omitempty"`

	// Used to disable the display of the device performance and health area.
	DefenderSecurityCenterDisableHealthUI nullable.Type[bool] `json:"defenderSecurityCenterDisableHealthUI,omitempty"`

	// Used to disable the display of the firewall and network protection area.
	DefenderSecurityCenterDisableNetworkUI nullable.Type[bool] `json:"defenderSecurityCenterDisableNetworkUI,omitempty"`

	// Used to disable the display of the notification area control. The user needs to either sign out and sign in or reboot
	// the computer for this setting to take effect.
	DefenderSecurityCenterDisableNotificationAreaUI nullable.Type[bool] `json:"defenderSecurityCenterDisableNotificationAreaUI,omitempty"`

	// Used to disable the display of the ransomware protection area.
	DefenderSecurityCenterDisableRansomwareUI nullable.Type[bool] `json:"defenderSecurityCenterDisableRansomwareUI,omitempty"`

	// Used to disable the display of the secure boot area under Device security.
	DefenderSecurityCenterDisableSecureBootUI nullable.Type[bool] `json:"defenderSecurityCenterDisableSecureBootUI,omitempty"`

	// Used to disable the display of the security process troubleshooting under Device security.
	DefenderSecurityCenterDisableTroubleshootingUI nullable.Type[bool] `json:"defenderSecurityCenterDisableTroubleshootingUI,omitempty"`

	// Used to disable the display of the virus and threat protection area.
	DefenderSecurityCenterDisableVirusUI nullable.Type[bool] `json:"defenderSecurityCenterDisableVirusUI,omitempty"`

	// Used to disable the display of update TPM Firmware when a vulnerable firmware is detected.
	DefenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI nullable.Type[bool] `json:"defenderSecurityCenterDisableVulnerableTpmFirmwareUpdateUI,omitempty"`

	// The email address that is displayed to users.
	DefenderSecurityCenterHelpEmail nullable.Type[string] `json:"defenderSecurityCenterHelpEmail,omitempty"`

	// The phone number or Skype ID that is displayed to users.
	DefenderSecurityCenterHelpPhone nullable.Type[string] `json:"defenderSecurityCenterHelpPhone,omitempty"`

	// The help portal URL this is displayed to users.
	DefenderSecurityCenterHelpURL nullable.Type[string] `json:"defenderSecurityCenterHelpURL,omitempty"`

	// Possible values for defenderSecurityCenterITContactDisplay
	DefenderSecurityCenterITContactDisplay *DefenderSecurityCenterITContactDisplayType `json:"defenderSecurityCenterITContactDisplay,omitempty"`

	// Possible values for defenderSecurityCenterNotificationsFromApp
	DefenderSecurityCenterNotificationsFromApp *DefenderSecurityCenterNotificationsFromAppType `json:"defenderSecurityCenterNotificationsFromApp,omitempty"`

	// The company name that is displayed to the users.
	DefenderSecurityCenterOrganizationDisplayName nullable.Type[string] `json:"defenderSecurityCenterOrganizationDisplayName,omitempty"`

	// Specifies the interval (in hours) that will be used to check for signatures, so instead of using the ScheduleDay and
	// ScheduleTime the check for new signatures will be set according to the interval. Valid values 0 to 24
	DefenderSignatureUpdateIntervalInHours nullable.Type[int64] `json:"defenderSignatureUpdateIntervalInHours,omitempty"`

	// Checks for the user consent level in Windows Defender to send data. Possible values are:
	// sendSafeSamplesAutomatically, alwaysPrompt, neverSend, sendAllSamplesAutomatically.
	DefenderSubmitSamplesConsentType *DefenderSubmitSamplesConsentType `json:"defenderSubmitSamplesConsentType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderUntrustedExecutable *DefenderProtectionType `json:"defenderUntrustedExecutable,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderUntrustedExecutableType *DefenderAttackSurfaceType `json:"defenderUntrustedExecutableType,omitempty"`

	// Possible values of Defender PUA Protection
	DefenderUntrustedUSBProcess *DefenderProtectionType `json:"defenderUntrustedUSBProcess,omitempty"`

	// Possible values of Defender Attack Surface Reduction Rules
	DefenderUntrustedUSBProcessType *DefenderAttackSurfaceType `json:"defenderUntrustedUSBProcessType,omitempty"`

	// This property will be deprecated in May 2019 and will be replaced with property DeviceGuardSecureBootWithDMA.
	// Specifies whether Platform Security Level is enabled at next reboot.
	DeviceGuardEnableSecureBootWithDMA *bool `json:"deviceGuardEnableSecureBootWithDMA,omitempty"`

	// Turns On Virtualization Based Security(VBS).
	DeviceGuardEnableVirtualizationBasedSecurity *bool `json:"deviceGuardEnableVirtualizationBasedSecurity,omitempty"`

	// Possible values of a property
	DeviceGuardLaunchSystemGuard *Enablement `json:"deviceGuardLaunchSystemGuard,omitempty"`

	// Possible values of Credential Guard settings.
	DeviceGuardLocalSystemAuthorityCredentialGuardSettings *DeviceGuardLocalSystemAuthorityCredentialGuardType `json:"deviceGuardLocalSystemAuthorityCredentialGuardSettings,omitempty"`

	// Possible values of Secure Boot with DMA
	DeviceGuardSecureBootWithDMA *SecureBootWithDMAType `json:"deviceGuardSecureBootWithDMA,omitempty"`

	// Possible values of the DmaGuardDeviceEnumerationPolicy.
	DmaGuardDeviceEnumerationPolicy *DmaGuardDeviceEnumerationPolicyType `json:"dmaGuardDeviceEnumerationPolicy,omitempty"`

	// Blocks stateful FTP connections to the device
	FirewallBlockStatefulFTP nullable.Type[bool] `json:"firewallBlockStatefulFTP,omitempty"`

	// Possible values for firewallCertificateRevocationListCheckMethod
	FirewallCertificateRevocationListCheckMethod *FirewallCertificateRevocationListCheckMethodType `json:"firewallCertificateRevocationListCheckMethod,omitempty"`

	// Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
	FirewallIPSecExemptionsAllowDHCP *bool `json:"firewallIPSecExemptionsAllowDHCP,omitempty"`

	// Configures IPSec exemptions to allow ICMP
	FirewallIPSecExemptionsAllowICMP *bool `json:"firewallIPSecExemptionsAllowICMP,omitempty"`

	// Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
	FirewallIPSecExemptionsAllowNeighborDiscovery *bool `json:"firewallIPSecExemptionsAllowNeighborDiscovery,omitempty"`

	// Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
	FirewallIPSecExemptionsAllowRouterDiscovery *bool `json:"firewallIPSecExemptionsAllowRouterDiscovery,omitempty"`

	// Configures IPSec exemptions to no exemptions
	FirewallIPSecExemptionsNone *bool `json:"firewallIPSecExemptionsNone,omitempty"`

	// Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period
	// after which security associations will expire and be deleted. Valid values 300 to 3600
	FirewallIdleTimeoutForSecurityAssociationInSeconds nullable.Type[int64] `json:"firewallIdleTimeoutForSecurityAssociationInSeconds,omitempty"`

	// If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported
	// authentication suites rather than the entire set
	FirewallMergeKeyingModuleSettings nullable.Type[bool] `json:"firewallMergeKeyingModuleSettings,omitempty"`

	// Possible values for firewallPacketQueueingMethod
	FirewallPacketQueueingMethod *FirewallPacketQueueingMethodType `json:"firewallPacketQueueingMethod,omitempty"`

	// Possible values for firewallPreSharedKeyEncodingMethod
	FirewallPreSharedKeyEncodingMethod *FirewallPreSharedKeyEncodingMethodType `json:"firewallPreSharedKeyEncodingMethod,omitempty"`

	// Configures the firewall profile settings for domain networks
	FirewallProfileDomain *WindowsFirewallNetworkProfile `json:"firewallProfileDomain,omitempty"`

	// Configures the firewall profile settings for private networks
	FirewallProfilePrivate *WindowsFirewallNetworkProfile `json:"firewallProfilePrivate,omitempty"`

	// Configures the firewall profile settings for public networks
	FirewallProfilePublic *WindowsFirewallNetworkProfile `json:"firewallProfilePublic,omitempty"`

	// Configures the firewall rule settings. This collection can contain a maximum of 150 elements.
	FirewallRules *[]WindowsFirewallRule `json:"firewallRules,omitempty"`

	// Possible values for LanManagerAuthenticationLevel
	LanManagerAuthenticationLevel *LanManagerAuthenticationLevel `json:"lanManagerAuthenticationLevel,omitempty"`

	// If enabled,the SMB client will allow insecure guest logons. If not configured, the SMB client will reject insecure
	// guest logons.
	LanManagerWorkstationDisableInsecureGuestLogons *bool `json:"lanManagerWorkstationDisableInsecureGuestLogons,omitempty"`

	// Define a different account name to be associated with the security identifier (SID) for the account 'Administrator'.
	LocalSecurityOptionsAdministratorAccountName nullable.Type[string] `json:"localSecurityOptionsAdministratorAccountName,omitempty"`

	// Possible values for LocalSecurityOptionsAdministratorElevationPromptBehavior
	LocalSecurityOptionsAdministratorElevationPromptBehavior *LocalSecurityOptionsAdministratorElevationPromptBehaviorType `json:"localSecurityOptionsAdministratorElevationPromptBehavior,omitempty"`

	// This security setting determines whether to allows anonymous users to perform certain activities, such as enumerating
	// the names of domain accounts and network shares.
	LocalSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares *bool `json:"localSecurityOptionsAllowAnonymousEnumerationOfSAMAccountsAndShares,omitempty"`

	// Block PKU2U authentication requests to this device to use online identities.
	LocalSecurityOptionsAllowPKU2UAuthenticationRequests *bool `json:"localSecurityOptionsAllowPKU2UAuthenticationRequests,omitempty"`

	// Edit the default Security Descriptor Definition Language string to allow or deny users and groups to make remote
	// calls to the SAM.
	LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager nullable.Type[string] `json:"localSecurityOptionsAllowRemoteCallsToSecurityAccountsManager,omitempty"`

	// UI helper boolean for LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManager entity
	LocalSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool *bool `json:"localSecurityOptionsAllowRemoteCallsToSecurityAccountsManagerHelperBool,omitempty"`

	// This security setting determines whether a computer can be shut down without having to log on to Windows.
	LocalSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn *bool `json:"localSecurityOptionsAllowSystemToBeShutDownWithoutHavingToLogOn,omitempty"`

	// Allow UIAccess apps to prompt for elevation without using the secure desktop.
	LocalSecurityOptionsAllowUIAccessApplicationElevation *bool `json:"localSecurityOptionsAllowUIAccessApplicationElevation,omitempty"`

	// Allow UIAccess apps to prompt for elevation without using the secure desktop.Default is enabled
	LocalSecurityOptionsAllowUIAccessApplicationsForSecureLocations *bool `json:"localSecurityOptionsAllowUIAccessApplicationsForSecureLocations,omitempty"`

	// Prevent a portable computer from being undocked without having to log in.
	LocalSecurityOptionsAllowUndockWithoutHavingToLogon *bool `json:"localSecurityOptionsAllowUndockWithoutHavingToLogon,omitempty"`

	// Prevent users from adding new Microsoft accounts to this computer.
	LocalSecurityOptionsBlockMicrosoftAccounts *bool `json:"localSecurityOptionsBlockMicrosoftAccounts,omitempty"`

	// Enable Local accounts that are not password protected to log on from locations other than the physical device.Default
	// is enabled
	LocalSecurityOptionsBlockRemoteLogonWithBlankPassword *bool `json:"localSecurityOptionsBlockRemoteLogonWithBlankPassword,omitempty"`

	// Enabling this settings allows only interactively logged on user to access CD-ROM media.
	LocalSecurityOptionsBlockRemoteOpticalDriveAccess *bool `json:"localSecurityOptionsBlockRemoteOpticalDriveAccess,omitempty"`

	// Restrict installing printer drivers as part of connecting to a shared printer to admins only.
	LocalSecurityOptionsBlockUsersInstallingPrinterDrivers *bool `json:"localSecurityOptionsBlockUsersInstallingPrinterDrivers,omitempty"`

	// This security setting determines whether the virtual memory pagefile is cleared when the system is shut down.
	LocalSecurityOptionsClearVirtualMemoryPageFile *bool `json:"localSecurityOptionsClearVirtualMemoryPageFile,omitempty"`

	// This security setting determines whether packet signing is required by the SMB client component.
	LocalSecurityOptionsClientDigitallySignCommunicationsAlways *bool `json:"localSecurityOptionsClientDigitallySignCommunicationsAlways,omitempty"`

	// If this security setting is enabled, the Server Message Block (SMB) redirector is allowed to send plaintext passwords
	// to non-Microsoft SMB servers that do not support password encryption during authentication.
	LocalSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers *bool `json:"localSecurityOptionsClientSendUnencryptedPasswordToThirdPartySMBServers,omitempty"`

	// App installations requiring elevated privileges will prompt for admin credentials.Default is enabled
	LocalSecurityOptionsDetectApplicationInstallationsAndPromptForElevation *bool `json:"localSecurityOptionsDetectApplicationInstallationsAndPromptForElevation,omitempty"`

	// Determines whether the Local Administrator account is enabled or disabled.
	LocalSecurityOptionsDisableAdministratorAccount *bool `json:"localSecurityOptionsDisableAdministratorAccount,omitempty"`

	// This security setting determines whether the SMB client attempts to negotiate SMB packet signing.
	LocalSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees *bool `json:"localSecurityOptionsDisableClientDigitallySignCommunicationsIfServerAgrees,omitempty"`

	// Determines if the Guest account is enabled or disabled.
	LocalSecurityOptionsDisableGuestAccount *bool `json:"localSecurityOptionsDisableGuestAccount,omitempty"`

	// This security setting determines whether packet signing is required by the SMB server component.
	LocalSecurityOptionsDisableServerDigitallySignCommunicationsAlways *bool `json:"localSecurityOptionsDisableServerDigitallySignCommunicationsAlways,omitempty"`

	// This security setting determines whether the SMB server will negotiate SMB packet signing with clients that request
	// it.
	LocalSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees *bool `json:"localSecurityOptionsDisableServerDigitallySignCommunicationsIfClientAgrees,omitempty"`

	// This security setting determines what additional permissions will be granted for anonymous connections to the
	// computer.
	LocalSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts *bool `json:"localSecurityOptionsDoNotAllowAnonymousEnumerationOfSAMAccounts,omitempty"`

	// Require CTRL+ALT+DEL to be pressed before a user can log on.
	LocalSecurityOptionsDoNotRequireCtrlAltDel *bool `json:"localSecurityOptionsDoNotRequireCtrlAltDel,omitempty"`

	// This security setting determines if, at the next password change, the LAN Manager (LM) hash value for the new
	// password is stored. It’s not stored by default.
	LocalSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange *bool `json:"localSecurityOptionsDoNotStoreLANManagerHashValueOnNextPasswordChange,omitempty"`

	// Possible values for LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType `json:"localSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser,omitempty"`

	// Define a different account name to be associated with the security identifier (SID) for the account 'Guest'.
	LocalSecurityOptionsGuestAccountName nullable.Type[string] `json:"localSecurityOptionsGuestAccountName,omitempty"`

	// Do not display the username of the last person who signed in on this device.
	LocalSecurityOptionsHideLastSignedInUser *bool `json:"localSecurityOptionsHideLastSignedInUser,omitempty"`

	// Do not display the username of the person signing in to this device after credentials are entered and before the
	// device’s desktop is shown.
	LocalSecurityOptionsHideUsernameAtSignIn *bool `json:"localSecurityOptionsHideUsernameAtSignIn,omitempty"`

	// Possible values for LocalSecurityOptionsInformationDisplayedOnLockScreen
	LocalSecurityOptionsInformationDisplayedOnLockScreen *LocalSecurityOptionsInformationDisplayedOnLockScreenType `json:"localSecurityOptionsInformationDisplayedOnLockScreen,omitempty"`

	// Possible values for LocalSecurityOptionsInformationShownOnLockScreenType
	LocalSecurityOptionsInformationShownOnLockScreen *LocalSecurityOptionsInformationShownOnLockScreenType `json:"localSecurityOptionsInformationShownOnLockScreen,omitempty"`

	// Set message text for users attempting to log in.
	LocalSecurityOptionsLogOnMessageText nullable.Type[string] `json:"localSecurityOptionsLogOnMessageText,omitempty"`

	// Set message title for users attempting to log in.
	LocalSecurityOptionsLogOnMessageTitle nullable.Type[string] `json:"localSecurityOptionsLogOnMessageTitle,omitempty"`

	// Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid
	// values 0 to 9999
	LocalSecurityOptionsMachineInactivityLimit nullable.Type[int64] `json:"localSecurityOptionsMachineInactivityLimit,omitempty"`

	// Define maximum minutes of inactivity on the interactive desktop’s login screen until the screen saver runs. Valid
	// values 0 to 9999
	LocalSecurityOptionsMachineInactivityLimitInMinutes nullable.Type[int64] `json:"localSecurityOptionsMachineInactivityLimitInMinutes,omitempty"`

	// Possible values for LocalSecurityOptionsMinimumSessionSecurity
	LocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients *LocalSecurityOptionsMinimumSessionSecurity `json:"localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedClients,omitempty"`

	// Possible values for LocalSecurityOptionsMinimumSessionSecurity
	LocalSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers *LocalSecurityOptionsMinimumSessionSecurity `json:"localSecurityOptionsMinimumSessionSecurityForNtlmSspBasedServers,omitempty"`

	// Enforce PKI certification path validation for a given executable file before it is permitted to run.
	LocalSecurityOptionsOnlyElevateSignedExecutables *bool `json:"localSecurityOptionsOnlyElevateSignedExecutables,omitempty"`

	// By default, this security setting restricts anonymous access to shares and pipes to the settings for named pipes that
	// can be accessed anonymously and Shares that can be accessed anonymously
	LocalSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares *bool `json:"localSecurityOptionsRestrictAnonymousAccessToNamedPipesAndShares,omitempty"`

	// Possible values for LocalSecurityOptionsSmartCardRemovalBehaviorType
	LocalSecurityOptionsSmartCardRemovalBehavior *LocalSecurityOptionsSmartCardRemovalBehaviorType `json:"localSecurityOptionsSmartCardRemovalBehavior,omitempty"`

	// Possible values for LocalSecurityOptionsStandardUserElevationPromptBehavior
	LocalSecurityOptionsStandardUserElevationPromptBehavior *LocalSecurityOptionsStandardUserElevationPromptBehaviorType `json:"localSecurityOptionsStandardUserElevationPromptBehavior,omitempty"`

	// Enable all elevation requests to go to the interactive user's desktop rather than the secure desktop. Prompt behavior
	// policy settings for admins and standard users are used.
	LocalSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation *bool `json:"localSecurityOptionsSwitchToSecureDesktopWhenPromptingForElevation,omitempty"`

	// Defines whether the built-in admin account uses Admin Approval Mode or runs all apps with full admin
	// privileges.Default is enabled
	LocalSecurityOptionsUseAdminApprovalMode *bool `json:"localSecurityOptionsUseAdminApprovalMode,omitempty"`

	// Define whether Admin Approval Mode and all UAC policy settings are enabled, default is enabled
	LocalSecurityOptionsUseAdminApprovalModeForAdministrators *bool `json:"localSecurityOptionsUseAdminApprovalModeForAdministrators,omitempty"`

	// Virtualize file and registry write failures to per user locations
	LocalSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations *bool `json:"localSecurityOptionsVirtualizeFileAndRegistryWriteFailuresToPerUserLocations,omitempty"`

	// Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
	SmartScreenBlockOverrideForFiles *bool `json:"smartScreenBlockOverrideForFiles,omitempty"`

	// Allows IT Admins to configure SmartScreen for Windows.
	SmartScreenEnableInShell *bool `json:"smartScreenEnableInShell,omitempty"`

	// This user right is used by Credential Manager during Backup/Restore. Users' saved credentials might be compromised if
	// this privilege is given to other entities. Only states NotConfigured and Allowed are supported
	UserRightsAccessCredentialManagerAsTrustedCaller *DeviceManagementUserRightsSetting `json:"userRightsAccessCredentialManagerAsTrustedCaller,omitempty"`

	// This user right allows a process to impersonate any user without authentication. The process can therefore gain
	// access to the same local resources as that user. Only states NotConfigured and Allowed are supported
	UserRightsActAsPartOfTheOperatingSystem *DeviceManagementUserRightsSetting `json:"userRightsActAsPartOfTheOperatingSystem,omitempty"`

	// This user right determines which users and groups are allowed to connect to the computer over the network. State
	// Allowed is supported.
	UserRightsAllowAccessFromNetwork *DeviceManagementUserRightsSetting `json:"userRightsAllowAccessFromNetwork,omitempty"`

	// This user right determines which users can bypass file, directory, registry, and other persistent objects permissions
	// when backing up files and directories. Only states NotConfigured and Allowed are supported
	UserRightsBackupData *DeviceManagementUserRightsSetting `json:"userRightsBackupData,omitempty"`

	// This user right determines which users and groups are block from connecting to the computer over the network. State
	// Block is supported.
	UserRightsBlockAccessFromNetwork *DeviceManagementUserRightsSetting `json:"userRightsBlockAccessFromNetwork,omitempty"`

	// This user right determines which users and groups can change the time and date on the internal clock of the computer.
	// Only states NotConfigured and Allowed are supported
	UserRightsChangeSystemTime *DeviceManagementUserRightsSetting `json:"userRightsChangeSystemTime,omitempty"`

	// This security setting determines whether users can create global objects that are available to all sessions. Users
	// who can create global objects could affect processes that run under other users' sessions, which could lead to
	// application failure or data corruption. Only states NotConfigured and Allowed are supported
	UserRightsCreateGlobalObjects *DeviceManagementUserRightsSetting `json:"userRightsCreateGlobalObjects,omitempty"`

	// This user right determines which users and groups can call an internal API to create and change the size of a page
	// file. Only states NotConfigured and Allowed are supported
	UserRightsCreatePageFile *DeviceManagementUserRightsSetting `json:"userRightsCreatePageFile,omitempty"`

	// This user right determines which accounts can be used by processes to create a directory object using the object
	// manager. Only states NotConfigured and Allowed are supported
	UserRightsCreatePermanentSharedObjects *DeviceManagementUserRightsSetting `json:"userRightsCreatePermanentSharedObjects,omitempty"`

	// This user right determines if the user can create a symbolic link from the computer to which they are logged on. Only
	// states NotConfigured and Allowed are supported
	UserRightsCreateSymbolicLinks *DeviceManagementUserRightsSetting `json:"userRightsCreateSymbolicLinks,omitempty"`

	// This user right determines which users/groups can be used by processes to create a token that can then be used to get
	// access to any local resources when the process uses an internal API to create an access token. Only states
	// NotConfigured and Allowed are supported
	UserRightsCreateToken *DeviceManagementUserRightsSetting `json:"userRightsCreateToken,omitempty"`

	// This user right determines which users can attach a debugger to any process or to the kernel. Only states
	// NotConfigured and Allowed are supported
	UserRightsDebugPrograms *DeviceManagementUserRightsSetting `json:"userRightsDebugPrograms,omitempty"`

	// This user right determines which users can set the Trusted for Delegation setting on a user or computer object. Only
	// states NotConfigured and Allowed are supported.
	UserRightsDelegation *DeviceManagementUserRightsSetting `json:"userRightsDelegation,omitempty"`

	// This user right determines which users cannot log on to the computer. States NotConfigured, Blocked are supported
	UserRightsDenyLocalLogOn *DeviceManagementUserRightsSetting `json:"userRightsDenyLocalLogOn,omitempty"`

	// This user right determines which accounts can be used by a process to add entries to the security log. The security
	// log is used to trace unauthorized system access. Only states NotConfigured and Allowed are supported.
	UserRightsGenerateSecurityAudits *DeviceManagementUserRightsSetting `json:"userRightsGenerateSecurityAudits,omitempty"`

	// Assigning this user right to a user allows programs running on behalf of that user to impersonate a client. Requiring
	// this user right for this kind of impersonation prevents an unauthorized user from convincing a client to connect to a
	// service that they have created and then impersonating that client, which can elevate the unauthorized user's
	// permissions to administrative or system levels. Only states NotConfigured and Allowed are supported.
	UserRightsImpersonateClient *DeviceManagementUserRightsSetting `json:"userRightsImpersonateClient,omitempty"`

	// This user right determines which accounts can use a process with Write Property access to another process to increase
	// the execution priority assigned to the other process. Only states NotConfigured and Allowed are supported.
	UserRightsIncreaseSchedulingPriority *DeviceManagementUserRightsSetting `json:"userRightsIncreaseSchedulingPriority,omitempty"`

	// This user right determines which users can dynamically load and unload device drivers or other code in to kernel
	// mode. Only states NotConfigured and Allowed are supported.
	UserRightsLoadUnloadDrivers *DeviceManagementUserRightsSetting `json:"userRightsLoadUnloadDrivers,omitempty"`

	// This user right determines which users can log on to the computer. States NotConfigured, Allowed are supported
	UserRightsLocalLogOn *DeviceManagementUserRightsSetting `json:"userRightsLocalLogOn,omitempty"`

	// This user right determines which accounts can use a process to keep data in physical memory, which prevents the
	// system from paging the data to virtual memory on disk. Only states NotConfigured and Allowed are supported.
	UserRightsLockMemory *DeviceManagementUserRightsSetting `json:"userRightsLockMemory,omitempty"`

	// This user right determines which users can specify object access auditing options for individual resources, such as
	// files, Active Directory objects, and registry keys. Only states NotConfigured and Allowed are supported.
	UserRightsManageAuditingAndSecurityLogs *DeviceManagementUserRightsSetting `json:"userRightsManageAuditingAndSecurityLogs,omitempty"`

	// This user right determines which users and groups can run maintenance tasks on a volume, such as remote
	// defragmentation. Only states NotConfigured and Allowed are supported.
	UserRightsManageVolumes *DeviceManagementUserRightsSetting `json:"userRightsManageVolumes,omitempty"`

	// This user right determines who can modify firmware environment values. Only states NotConfigured and Allowed are
	// supported.
	UserRightsModifyFirmwareEnvironment *DeviceManagementUserRightsSetting `json:"userRightsModifyFirmwareEnvironment,omitempty"`

	// This user right determines which user accounts can modify the integrity label of objects, such as files, registry
	// keys, or processes owned by other users. Only states NotConfigured and Allowed are supported.
	UserRightsModifyObjectLabels *DeviceManagementUserRightsSetting `json:"userRightsModifyObjectLabels,omitempty"`

	// This user right determines which users can use performance monitoring tools to monitor the performance of system
	// processes. Only states NotConfigured and Allowed are supported.
	UserRightsProfileSingleProcess *DeviceManagementUserRightsSetting `json:"userRightsProfileSingleProcess,omitempty"`

	// This user right determines which users and groups are prohibited from logging on as a Remote Desktop Services client.
	// Only states NotConfigured and Blocked are supported
	UserRightsRemoteDesktopServicesLogOn *DeviceManagementUserRightsSetting `json:"userRightsRemoteDesktopServicesLogOn,omitempty"`

	// This user right determines which users are allowed to shut down a computer from a remote location on the network.
	// Misuse of this user right can result in a denial of service. Only states NotConfigured and Allowed are supported.
	UserRightsRemoteShutdown *DeviceManagementUserRightsSetting `json:"userRightsRemoteShutdown,omitempty"`

	// This user right determines which users can bypass file, directory, registry, and other persistent objects permissions
	// when restoring backed up files and directories, and determines which users can set any valid security principal as
	// the owner of an object. Only states NotConfigured and Allowed are supported.
	UserRightsRestoreData *DeviceManagementUserRightsSetting `json:"userRightsRestoreData,omitempty"`

	// This user right determines which users can take ownership of any securable object in the system, including Active
	// Directory objects, files and folders, printers, registry keys, processes, and threads. Only states NotConfigured and
	// Allowed are supported.
	UserRightsTakeOwnership *DeviceManagementUserRightsSetting `json:"userRightsTakeOwnership,omitempty"`

	// Defender TamperProtection setting options
	WindowsDefenderTamperProtection *WindowsDefenderTamperProtectionOptions `json:"windowsDefenderTamperProtection,omitempty"`

	// Possible values of xbox service start type
	XboxServicesAccessoryManagementServiceStartupMode *ServiceStartType `json:"xboxServicesAccessoryManagementServiceStartupMode,omitempty"`

	// This setting determines whether xbox game save is enabled (1) or disabled (0).
	XboxServicesEnableXboxGameSaveTask *bool `json:"xboxServicesEnableXboxGameSaveTask,omitempty"`

	// Possible values of xbox service start type
	XboxServicesLiveAuthManagerServiceStartupMode *ServiceStartType `json:"xboxServicesLiveAuthManagerServiceStartupMode,omitempty"`

	// Possible values of xbox service start type
	XboxServicesLiveGameSaveServiceStartupMode *ServiceStartType `json:"xboxServicesLiveGameSaveServiceStartupMode,omitempty"`

	// Possible values of xbox service start type
	XboxServicesLiveNetworkingServiceStartupMode *ServiceStartType `json:"xboxServicesLiveNetworkingServiceStartupMode,omitempty"`

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

func (s Windows10EndpointProtectionConfiguration) DeviceConfiguration() BaseDeviceConfigurationImpl {
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

func (s Windows10EndpointProtectionConfiguration) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Windows10EndpointProtectionConfiguration{}

func (s Windows10EndpointProtectionConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper Windows10EndpointProtectionConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Windows10EndpointProtectionConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Windows10EndpointProtectionConfiguration: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windows10EndpointProtectionConfiguration"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Windows10EndpointProtectionConfiguration: %+v", err)
	}

	return encoded, nil
}
