package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdatesInstallStatus string

const (
	IosUpdatesInstallStatus_Available                          IosUpdatesInstallStatus = "available"
	IosUpdatesInstallStatus_DeviceOsHigherThanDesiredOsVersion IosUpdatesInstallStatus = "deviceOsHigherThanDesiredOsVersion"
	IosUpdatesInstallStatus_DownloadFailed                     IosUpdatesInstallStatus = "downloadFailed"
	IosUpdatesInstallStatus_DownloadInsufficientNetwork        IosUpdatesInstallStatus = "downloadInsufficientNetwork"
	IosUpdatesInstallStatus_DownloadInsufficientPower          IosUpdatesInstallStatus = "downloadInsufficientPower"
	IosUpdatesInstallStatus_DownloadInsufficientSpace          IosUpdatesInstallStatus = "downloadInsufficientSpace"
	IosUpdatesInstallStatus_DownloadRequiresComputer           IosUpdatesInstallStatus = "downloadRequiresComputer"
	IosUpdatesInstallStatus_Downloading                        IosUpdatesInstallStatus = "downloading"
	IosUpdatesInstallStatus_Idle                               IosUpdatesInstallStatus = "idle"
	IosUpdatesInstallStatus_InstallFailed                      IosUpdatesInstallStatus = "installFailed"
	IosUpdatesInstallStatus_InstallInsufficientPower           IosUpdatesInstallStatus = "installInsufficientPower"
	IosUpdatesInstallStatus_InstallInsufficientSpace           IosUpdatesInstallStatus = "installInsufficientSpace"
	IosUpdatesInstallStatus_InstallPhoneCallInProgress         IosUpdatesInstallStatus = "installPhoneCallInProgress"
	IosUpdatesInstallStatus_Installing                         IosUpdatesInstallStatus = "installing"
	IosUpdatesInstallStatus_MdmClientCrashed                   IosUpdatesInstallStatus = "mdmClientCrashed"
	IosUpdatesInstallStatus_NotSupportedOperation              IosUpdatesInstallStatus = "notSupportedOperation"
	IosUpdatesInstallStatus_SharedDeviceUserLoggedInError      IosUpdatesInstallStatus = "sharedDeviceUserLoggedInError"
	IosUpdatesInstallStatus_Success                            IosUpdatesInstallStatus = "success"
	IosUpdatesInstallStatus_Timeout                            IosUpdatesInstallStatus = "timeout"
	IosUpdatesInstallStatus_Unknown                            IosUpdatesInstallStatus = "unknown"
	IosUpdatesInstallStatus_UpdateError                        IosUpdatesInstallStatus = "updateError"
	IosUpdatesInstallStatus_UpdateScanFailed                   IosUpdatesInstallStatus = "updateScanFailed"
)

func PossibleValuesForIosUpdatesInstallStatus() []string {
	return []string{
		string(IosUpdatesInstallStatus_Available),
		string(IosUpdatesInstallStatus_DeviceOsHigherThanDesiredOsVersion),
		string(IosUpdatesInstallStatus_DownloadFailed),
		string(IosUpdatesInstallStatus_DownloadInsufficientNetwork),
		string(IosUpdatesInstallStatus_DownloadInsufficientPower),
		string(IosUpdatesInstallStatus_DownloadInsufficientSpace),
		string(IosUpdatesInstallStatus_DownloadRequiresComputer),
		string(IosUpdatesInstallStatus_Downloading),
		string(IosUpdatesInstallStatus_Idle),
		string(IosUpdatesInstallStatus_InstallFailed),
		string(IosUpdatesInstallStatus_InstallInsufficientPower),
		string(IosUpdatesInstallStatus_InstallInsufficientSpace),
		string(IosUpdatesInstallStatus_InstallPhoneCallInProgress),
		string(IosUpdatesInstallStatus_Installing),
		string(IosUpdatesInstallStatus_MdmClientCrashed),
		string(IosUpdatesInstallStatus_NotSupportedOperation),
		string(IosUpdatesInstallStatus_SharedDeviceUserLoggedInError),
		string(IosUpdatesInstallStatus_Success),
		string(IosUpdatesInstallStatus_Timeout),
		string(IosUpdatesInstallStatus_Unknown),
		string(IosUpdatesInstallStatus_UpdateError),
		string(IosUpdatesInstallStatus_UpdateScanFailed),
	}
}

func (s *IosUpdatesInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosUpdatesInstallStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosUpdatesInstallStatus(input string) (*IosUpdatesInstallStatus, error) {
	vals := map[string]IosUpdatesInstallStatus{
		"available":                          IosUpdatesInstallStatus_Available,
		"deviceoshigherthandesiredosversion": IosUpdatesInstallStatus_DeviceOsHigherThanDesiredOsVersion,
		"downloadfailed":                     IosUpdatesInstallStatus_DownloadFailed,
		"downloadinsufficientnetwork":        IosUpdatesInstallStatus_DownloadInsufficientNetwork,
		"downloadinsufficientpower":          IosUpdatesInstallStatus_DownloadInsufficientPower,
		"downloadinsufficientspace":          IosUpdatesInstallStatus_DownloadInsufficientSpace,
		"downloadrequirescomputer":           IosUpdatesInstallStatus_DownloadRequiresComputer,
		"downloading":                        IosUpdatesInstallStatus_Downloading,
		"idle":                               IosUpdatesInstallStatus_Idle,
		"installfailed":                      IosUpdatesInstallStatus_InstallFailed,
		"installinsufficientpower":           IosUpdatesInstallStatus_InstallInsufficientPower,
		"installinsufficientspace":           IosUpdatesInstallStatus_InstallInsufficientSpace,
		"installphonecallinprogress":         IosUpdatesInstallStatus_InstallPhoneCallInProgress,
		"installing":                         IosUpdatesInstallStatus_Installing,
		"mdmclientcrashed":                   IosUpdatesInstallStatus_MdmClientCrashed,
		"notsupportedoperation":              IosUpdatesInstallStatus_NotSupportedOperation,
		"shareddeviceuserloggedinerror":      IosUpdatesInstallStatus_SharedDeviceUserLoggedInError,
		"success":                            IosUpdatesInstallStatus_Success,
		"timeout":                            IosUpdatesInstallStatus_Timeout,
		"unknown":                            IosUpdatesInstallStatus_Unknown,
		"updateerror":                        IosUpdatesInstallStatus_UpdateError,
		"updatescanfailed":                   IosUpdatesInstallStatus_UpdateScanFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdatesInstallStatus(input)
	return &out, nil
}
