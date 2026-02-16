package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAction string

const (
	RemoteAction_ActivateDeviceEsim                                   RemoteAction = "activateDeviceEsim"
	RemoteAction_AutomaticRedeployment                                RemoteAction = "automaticRedeployment"
	RemoteAction_ChangeAssignments                                    RemoteAction = "changeAssignments"
	RemoteAction_CleanWindowsDevice                                   RemoteAction = "cleanWindowsDevice"
	RemoteAction_Delete                                               RemoteAction = "delete"
	RemoteAction_Deprovision                                          RemoteAction = "deprovision"
	RemoteAction_Disable                                              RemoteAction = "disable"
	RemoteAction_DisableLostMode                                      RemoteAction = "disableLostMode"
	RemoteAction_EnableLostMode                                       RemoteAction = "enableLostMode"
	RemoteAction_FactoryReset                                         RemoteAction = "factoryReset"
	RemoteAction_FactoryResetKeepEnrollmentData                       RemoteAction = "factoryResetKeepEnrollmentData"
	RemoteAction_FullScan                                             RemoteAction = "fullScan"
	RemoteAction_GetFileVaultKey                                      RemoteAction = "getFileVaultKey"
	RemoteAction_InitiateDeviceAttestation                            RemoteAction = "initiateDeviceAttestation"
	RemoteAction_InitiateMobileDeviceManagementKeyRecovery            RemoteAction = "initiateMobileDeviceManagementKeyRecovery"
	RemoteAction_InitiateOnDemandProactiveRemediation                 RemoteAction = "initiateOnDemandProactiveRemediation"
	RemoteAction_LaunchRemoteHelp                                     RemoteAction = "launchRemoteHelp"
	RemoteAction_LocateDevice                                         RemoteAction = "locateDevice"
	RemoteAction_LogoutSharedAppleDeviceActiveUser                    RemoteAction = "logoutSharedAppleDeviceActiveUser"
	RemoteAction_MoveDeviceToOrganizationalUnit                       RemoteAction = "moveDeviceToOrganizationalUnit"
	RemoteAction_PauseConfigurationRefresh                            RemoteAction = "pauseConfigurationRefresh"
	RemoteAction_QuickScan                                            RemoteAction = "quickScan"
	RemoteAction_RebootNow                                            RemoteAction = "rebootNow"
	RemoteAction_RecoverPasscode                                      RemoteAction = "recoverPasscode"
	RemoteAction_Reenable                                             RemoteAction = "reenable"
	RemoteAction_RemoteLock                                           RemoteAction = "remoteLock"
	RemoteAction_RemoveCompanyData                                    RemoteAction = "removeCompanyData"
	RemoteAction_RemoveDeviceFirmwareConfigurationInterfaceManagement RemoteAction = "removeDeviceFirmwareConfigurationInterfaceManagement"
	RemoteAction_ResetPasscode                                        RemoteAction = "resetPasscode"
	RemoteAction_RevokeAppleVppLicenses                               RemoteAction = "revokeAppleVppLicenses"
	RemoteAction_RotateBitLockerKeys                                  RemoteAction = "rotateBitLockerKeys"
	RemoteAction_RotateFileVaultKey                                   RemoteAction = "rotateFileVaultKey"
	RemoteAction_RotateLocalAdminPassword                             RemoteAction = "rotateLocalAdminPassword"
	RemoteAction_SetDeviceName                                        RemoteAction = "setDeviceName"
	RemoteAction_ShutDown                                             RemoteAction = "shutDown"
	RemoteAction_Unknown                                              RemoteAction = "unknown"
	RemoteAction_UpdateDeviceAccount                                  RemoteAction = "updateDeviceAccount"
	RemoteAction_WindowsDefenderUpdateSignatures                      RemoteAction = "windowsDefenderUpdateSignatures"
)

func PossibleValuesForRemoteAction() []string {
	return []string{
		string(RemoteAction_ActivateDeviceEsim),
		string(RemoteAction_AutomaticRedeployment),
		string(RemoteAction_ChangeAssignments),
		string(RemoteAction_CleanWindowsDevice),
		string(RemoteAction_Delete),
		string(RemoteAction_Deprovision),
		string(RemoteAction_Disable),
		string(RemoteAction_DisableLostMode),
		string(RemoteAction_EnableLostMode),
		string(RemoteAction_FactoryReset),
		string(RemoteAction_FactoryResetKeepEnrollmentData),
		string(RemoteAction_FullScan),
		string(RemoteAction_GetFileVaultKey),
		string(RemoteAction_InitiateDeviceAttestation),
		string(RemoteAction_InitiateMobileDeviceManagementKeyRecovery),
		string(RemoteAction_InitiateOnDemandProactiveRemediation),
		string(RemoteAction_LaunchRemoteHelp),
		string(RemoteAction_LocateDevice),
		string(RemoteAction_LogoutSharedAppleDeviceActiveUser),
		string(RemoteAction_MoveDeviceToOrganizationalUnit),
		string(RemoteAction_PauseConfigurationRefresh),
		string(RemoteAction_QuickScan),
		string(RemoteAction_RebootNow),
		string(RemoteAction_RecoverPasscode),
		string(RemoteAction_Reenable),
		string(RemoteAction_RemoteLock),
		string(RemoteAction_RemoveCompanyData),
		string(RemoteAction_RemoveDeviceFirmwareConfigurationInterfaceManagement),
		string(RemoteAction_ResetPasscode),
		string(RemoteAction_RevokeAppleVppLicenses),
		string(RemoteAction_RotateBitLockerKeys),
		string(RemoteAction_RotateFileVaultKey),
		string(RemoteAction_RotateLocalAdminPassword),
		string(RemoteAction_SetDeviceName),
		string(RemoteAction_ShutDown),
		string(RemoteAction_Unknown),
		string(RemoteAction_UpdateDeviceAccount),
		string(RemoteAction_WindowsDefenderUpdateSignatures),
	}
}

func (s *RemoteAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteAction(input string) (*RemoteAction, error) {
	vals := map[string]RemoteAction{
		"activatedeviceesim":             RemoteAction_ActivateDeviceEsim,
		"automaticredeployment":          RemoteAction_AutomaticRedeployment,
		"changeassignments":              RemoteAction_ChangeAssignments,
		"cleanwindowsdevice":             RemoteAction_CleanWindowsDevice,
		"delete":                         RemoteAction_Delete,
		"deprovision":                    RemoteAction_Deprovision,
		"disable":                        RemoteAction_Disable,
		"disablelostmode":                RemoteAction_DisableLostMode,
		"enablelostmode":                 RemoteAction_EnableLostMode,
		"factoryreset":                   RemoteAction_FactoryReset,
		"factoryresetkeepenrollmentdata": RemoteAction_FactoryResetKeepEnrollmentData,
		"fullscan":                       RemoteAction_FullScan,
		"getfilevaultkey":                RemoteAction_GetFileVaultKey,
		"initiatedeviceattestation":      RemoteAction_InitiateDeviceAttestation,
		"initiatemobiledevicemanagementkeyrecovery": RemoteAction_InitiateMobileDeviceManagementKeyRecovery,
		"initiateondemandproactiveremediation":      RemoteAction_InitiateOnDemandProactiveRemediation,
		"launchremotehelp":                          RemoteAction_LaunchRemoteHelp,
		"locatedevice":                              RemoteAction_LocateDevice,
		"logoutsharedappledeviceactiveuser":         RemoteAction_LogoutSharedAppleDeviceActiveUser,
		"movedevicetoorganizationalunit":            RemoteAction_MoveDeviceToOrganizationalUnit,
		"pauseconfigurationrefresh":                 RemoteAction_PauseConfigurationRefresh,
		"quickscan":                                 RemoteAction_QuickScan,
		"rebootnow":                                 RemoteAction_RebootNow,
		"recoverpasscode":                           RemoteAction_RecoverPasscode,
		"reenable":                                  RemoteAction_Reenable,
		"remotelock":                                RemoteAction_RemoteLock,
		"removecompanydata":                         RemoteAction_RemoveCompanyData,
		"removedevicefirmwareconfigurationinterfacemanagement": RemoteAction_RemoveDeviceFirmwareConfigurationInterfaceManagement,
		"resetpasscode":                   RemoteAction_ResetPasscode,
		"revokeapplevpplicenses":          RemoteAction_RevokeAppleVppLicenses,
		"rotatebitlockerkeys":             RemoteAction_RotateBitLockerKeys,
		"rotatefilevaultkey":              RemoteAction_RotateFileVaultKey,
		"rotatelocaladminpassword":        RemoteAction_RotateLocalAdminPassword,
		"setdevicename":                   RemoteAction_SetDeviceName,
		"shutdown":                        RemoteAction_ShutDown,
		"unknown":                         RemoteAction_Unknown,
		"updatedeviceaccount":             RemoteAction_UpdateDeviceAccount,
		"windowsdefenderupdatesignatures": RemoteAction_WindowsDefenderUpdateSignatures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteAction(input)
	return &out, nil
}
