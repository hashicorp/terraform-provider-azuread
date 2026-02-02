package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResultantAppStateDetail string

const (
	ResultantAppStateDetail_AppRemovedBySupersedence            ResultantAppStateDetail = "appRemovedBySupersedence"
	ResultantAppStateDetail_AutoInstallDisabled                 ResultantAppStateDetail = "autoInstallDisabled"
	ResultantAppStateDetail_ContentDownloaded                   ResultantAppStateDetail = "contentDownloaded"
	ResultantAppStateDetail_DependencyFailedToInstall           ResultantAppStateDetail = "dependencyFailedToInstall"
	ResultantAppStateDetail_DependencyPendingReboot             ResultantAppStateDetail = "dependencyPendingReboot"
	ResultantAppStateDetail_DependencyWithAutoInstallDisabled   ResultantAppStateDetail = "dependencyWithAutoInstallDisabled"
	ResultantAppStateDetail_DependencyWithRequirementsNotMet    ResultantAppStateDetail = "dependencyWithRequirementsNotMet"
	ResultantAppStateDetail_FileSystemRequirementNotMet         ResultantAppStateDetail = "fileSystemRequirementNotMet"
	ResultantAppStateDetail_InstallingDependencies              ResultantAppStateDetail = "installingDependencies"
	ResultantAppStateDetail_IosAppStoreUpdateFailedToInstall    ResultantAppStateDetail = "iosAppStoreUpdateFailedToInstall"
	ResultantAppStateDetail_ManagedAppNoLongerPresent           ResultantAppStateDetail = "managedAppNoLongerPresent"
	ResultantAppStateDetail_MinimumCpuSpeedNotMet               ResultantAppStateDetail = "minimumCpuSpeedNotMet"
	ResultantAppStateDetail_MinimumDiskSpaceNotMet              ResultantAppStateDetail = "minimumDiskSpaceNotMet"
	ResultantAppStateDetail_MinimumLogicalProcessorCountNotMet  ResultantAppStateDetail = "minimumLogicalProcessorCountNotMet"
	ResultantAppStateDetail_MinimumOsVersionNotMet              ResultantAppStateDetail = "minimumOsVersionNotMet"
	ResultantAppStateDetail_MinimumPhysicalMemoryNotMet         ResultantAppStateDetail = "minimumPhysicalMemoryNotMet"
	ResultantAppStateDetail_NoAdditionalDetails                 ResultantAppStateDetail = "noAdditionalDetails"
	ResultantAppStateDetail_PendingReboot                       ResultantAppStateDetail = "pendingReboot"
	ResultantAppStateDetail_PlatformNotApplicable               ResultantAppStateDetail = "platformNotApplicable"
	ResultantAppStateDetail_PowerShellScriptRequirementNotMet   ResultantAppStateDetail = "powerShellScriptRequirementNotMet"
	ResultantAppStateDetail_ProcessorArchitectureNotApplicable  ResultantAppStateDetail = "processorArchitectureNotApplicable"
	ResultantAppStateDetail_RegistryRequirementNotMet           ResultantAppStateDetail = "registryRequirementNotMet"
	ResultantAppStateDetail_RemovingSupersededApps              ResultantAppStateDetail = "removingSupersededApps"
	ResultantAppStateDetail_SeeInstallErrorCode                 ResultantAppStateDetail = "seeInstallErrorCode"
	ResultantAppStateDetail_SeeUninstallErrorCode               ResultantAppStateDetail = "seeUninstallErrorCode"
	ResultantAppStateDetail_SupersededAppUninstallFailed        ResultantAppStateDetail = "supersededAppUninstallFailed"
	ResultantAppStateDetail_SupersededAppUninstallPendingReboot ResultantAppStateDetail = "supersededAppUninstallPendingReboot"
	ResultantAppStateDetail_SupersededAppsDetected              ResultantAppStateDetail = "supersededAppsDetected"
	ResultantAppStateDetail_SupersedingAppsDetected             ResultantAppStateDetail = "supersedingAppsDetected"
	ResultantAppStateDetail_SupersedingAppsNotApplicable        ResultantAppStateDetail = "supersedingAppsNotApplicable"
	ResultantAppStateDetail_UninstallPendingReboot              ResultantAppStateDetail = "uninstallPendingReboot"
	ResultantAppStateDetail_UntargetedSupersedingAppsDetected   ResultantAppStateDetail = "untargetedSupersedingAppsDetected"
	ResultantAppStateDetail_UserIsNotLoggedIntoAppStore         ResultantAppStateDetail = "userIsNotLoggedIntoAppStore"
	ResultantAppStateDetail_UserRejectedInstall                 ResultantAppStateDetail = "userRejectedInstall"
	ResultantAppStateDetail_UserRejectedUpdate                  ResultantAppStateDetail = "userRejectedUpdate"
	ResultantAppStateDetail_VppAppHasUpdateAvailable            ResultantAppStateDetail = "vppAppHasUpdateAvailable"
)

func PossibleValuesForResultantAppStateDetail() []string {
	return []string{
		string(ResultantAppStateDetail_AppRemovedBySupersedence),
		string(ResultantAppStateDetail_AutoInstallDisabled),
		string(ResultantAppStateDetail_ContentDownloaded),
		string(ResultantAppStateDetail_DependencyFailedToInstall),
		string(ResultantAppStateDetail_DependencyPendingReboot),
		string(ResultantAppStateDetail_DependencyWithAutoInstallDisabled),
		string(ResultantAppStateDetail_DependencyWithRequirementsNotMet),
		string(ResultantAppStateDetail_FileSystemRequirementNotMet),
		string(ResultantAppStateDetail_InstallingDependencies),
		string(ResultantAppStateDetail_IosAppStoreUpdateFailedToInstall),
		string(ResultantAppStateDetail_ManagedAppNoLongerPresent),
		string(ResultantAppStateDetail_MinimumCpuSpeedNotMet),
		string(ResultantAppStateDetail_MinimumDiskSpaceNotMet),
		string(ResultantAppStateDetail_MinimumLogicalProcessorCountNotMet),
		string(ResultantAppStateDetail_MinimumOsVersionNotMet),
		string(ResultantAppStateDetail_MinimumPhysicalMemoryNotMet),
		string(ResultantAppStateDetail_NoAdditionalDetails),
		string(ResultantAppStateDetail_PendingReboot),
		string(ResultantAppStateDetail_PlatformNotApplicable),
		string(ResultantAppStateDetail_PowerShellScriptRequirementNotMet),
		string(ResultantAppStateDetail_ProcessorArchitectureNotApplicable),
		string(ResultantAppStateDetail_RegistryRequirementNotMet),
		string(ResultantAppStateDetail_RemovingSupersededApps),
		string(ResultantAppStateDetail_SeeInstallErrorCode),
		string(ResultantAppStateDetail_SeeUninstallErrorCode),
		string(ResultantAppStateDetail_SupersededAppUninstallFailed),
		string(ResultantAppStateDetail_SupersededAppUninstallPendingReboot),
		string(ResultantAppStateDetail_SupersededAppsDetected),
		string(ResultantAppStateDetail_SupersedingAppsDetected),
		string(ResultantAppStateDetail_SupersedingAppsNotApplicable),
		string(ResultantAppStateDetail_UninstallPendingReboot),
		string(ResultantAppStateDetail_UntargetedSupersedingAppsDetected),
		string(ResultantAppStateDetail_UserIsNotLoggedIntoAppStore),
		string(ResultantAppStateDetail_UserRejectedInstall),
		string(ResultantAppStateDetail_UserRejectedUpdate),
		string(ResultantAppStateDetail_VppAppHasUpdateAvailable),
	}
}

func (s *ResultantAppStateDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResultantAppStateDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResultantAppStateDetail(input string) (*ResultantAppStateDetail, error) {
	vals := map[string]ResultantAppStateDetail{
		"appremovedbysupersedence":            ResultantAppStateDetail_AppRemovedBySupersedence,
		"autoinstalldisabled":                 ResultantAppStateDetail_AutoInstallDisabled,
		"contentdownloaded":                   ResultantAppStateDetail_ContentDownloaded,
		"dependencyfailedtoinstall":           ResultantAppStateDetail_DependencyFailedToInstall,
		"dependencypendingreboot":             ResultantAppStateDetail_DependencyPendingReboot,
		"dependencywithautoinstalldisabled":   ResultantAppStateDetail_DependencyWithAutoInstallDisabled,
		"dependencywithrequirementsnotmet":    ResultantAppStateDetail_DependencyWithRequirementsNotMet,
		"filesystemrequirementnotmet":         ResultantAppStateDetail_FileSystemRequirementNotMet,
		"installingdependencies":              ResultantAppStateDetail_InstallingDependencies,
		"iosappstoreupdatefailedtoinstall":    ResultantAppStateDetail_IosAppStoreUpdateFailedToInstall,
		"managedappnolongerpresent":           ResultantAppStateDetail_ManagedAppNoLongerPresent,
		"minimumcpuspeednotmet":               ResultantAppStateDetail_MinimumCpuSpeedNotMet,
		"minimumdiskspacenotmet":              ResultantAppStateDetail_MinimumDiskSpaceNotMet,
		"minimumlogicalprocessorcountnotmet":  ResultantAppStateDetail_MinimumLogicalProcessorCountNotMet,
		"minimumosversionnotmet":              ResultantAppStateDetail_MinimumOsVersionNotMet,
		"minimumphysicalmemorynotmet":         ResultantAppStateDetail_MinimumPhysicalMemoryNotMet,
		"noadditionaldetails":                 ResultantAppStateDetail_NoAdditionalDetails,
		"pendingreboot":                       ResultantAppStateDetail_PendingReboot,
		"platformnotapplicable":               ResultantAppStateDetail_PlatformNotApplicable,
		"powershellscriptrequirementnotmet":   ResultantAppStateDetail_PowerShellScriptRequirementNotMet,
		"processorarchitecturenotapplicable":  ResultantAppStateDetail_ProcessorArchitectureNotApplicable,
		"registryrequirementnotmet":           ResultantAppStateDetail_RegistryRequirementNotMet,
		"removingsupersededapps":              ResultantAppStateDetail_RemovingSupersededApps,
		"seeinstallerrorcode":                 ResultantAppStateDetail_SeeInstallErrorCode,
		"seeuninstallerrorcode":               ResultantAppStateDetail_SeeUninstallErrorCode,
		"supersededappuninstallfailed":        ResultantAppStateDetail_SupersededAppUninstallFailed,
		"supersededappuninstallpendingreboot": ResultantAppStateDetail_SupersededAppUninstallPendingReboot,
		"supersededappsdetected":              ResultantAppStateDetail_SupersededAppsDetected,
		"supersedingappsdetected":             ResultantAppStateDetail_SupersedingAppsDetected,
		"supersedingappsnotapplicable":        ResultantAppStateDetail_SupersedingAppsNotApplicable,
		"uninstallpendingreboot":              ResultantAppStateDetail_UninstallPendingReboot,
		"untargetedsupersedingappsdetected":   ResultantAppStateDetail_UntargetedSupersedingAppsDetected,
		"userisnotloggedintoappstore":         ResultantAppStateDetail_UserIsNotLoggedIntoAppStore,
		"userrejectedinstall":                 ResultantAppStateDetail_UserRejectedInstall,
		"userrejectedupdate":                  ResultantAppStateDetail_UserRejectedUpdate,
		"vppapphasupdateavailable":            ResultantAppStateDetail_VppAppHasUpdateAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResultantAppStateDetail(input)
	return &out, nil
}
