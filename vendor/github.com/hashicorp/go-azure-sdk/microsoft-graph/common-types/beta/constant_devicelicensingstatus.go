package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLicensingStatus string

const (
	DeviceLicensingStatus_AcquiringDeviceLicense                DeviceLicensingStatus = "acquiringDeviceLicense"
	DeviceLicensingStatus_DeviceIdentityVerificationFailed      DeviceLicensingStatus = "deviceIdentityVerificationFailed"
	DeviceLicensingStatus_DeviceIsNotAzureActiveDirectoryJoined DeviceLicensingStatus = "deviceIsNotAzureActiveDirectoryJoined"
	DeviceLicensingStatus_DeviceLicenseRefreshFailed            DeviceLicensingStatus = "deviceLicenseRefreshFailed"
	DeviceLicensingStatus_DeviceLicenseRefreshSucceed           DeviceLicensingStatus = "deviceLicenseRefreshSucceed"
	DeviceLicensingStatus_DeviceLicenseRemoveFailed             DeviceLicensingStatus = "deviceLicenseRemoveFailed"
	DeviceLicensingStatus_DeviceLicenseRemoveSucceed            DeviceLicensingStatus = "deviceLicenseRemoveSucceed"
	DeviceLicensingStatus_LicenseRefreshPending                 DeviceLicensingStatus = "licenseRefreshPending"
	DeviceLicensingStatus_LicenseRefreshStarted                 DeviceLicensingStatus = "licenseRefreshStarted"
	DeviceLicensingStatus_MicrosoftAccountVerificationFailed    DeviceLicensingStatus = "microsoftAccountVerificationFailed"
	DeviceLicensingStatus_RefreshingDeviceLicense               DeviceLicensingStatus = "refreshingDeviceLicense"
	DeviceLicensingStatus_RemovingDeviceLicense                 DeviceLicensingStatus = "removingDeviceLicense"
	DeviceLicensingStatus_Unknown                               DeviceLicensingStatus = "unknown"
	DeviceLicensingStatus_VerifyingMicrosoftAccountIdentity     DeviceLicensingStatus = "verifyingMicrosoftAccountIdentity"
	DeviceLicensingStatus_VerifyingMicrosoftDeviceIdentity      DeviceLicensingStatus = "verifyingMicrosoftDeviceIdentity"
)

func PossibleValuesForDeviceLicensingStatus() []string {
	return []string{
		string(DeviceLicensingStatus_AcquiringDeviceLicense),
		string(DeviceLicensingStatus_DeviceIdentityVerificationFailed),
		string(DeviceLicensingStatus_DeviceIsNotAzureActiveDirectoryJoined),
		string(DeviceLicensingStatus_DeviceLicenseRefreshFailed),
		string(DeviceLicensingStatus_DeviceLicenseRefreshSucceed),
		string(DeviceLicensingStatus_DeviceLicenseRemoveFailed),
		string(DeviceLicensingStatus_DeviceLicenseRemoveSucceed),
		string(DeviceLicensingStatus_LicenseRefreshPending),
		string(DeviceLicensingStatus_LicenseRefreshStarted),
		string(DeviceLicensingStatus_MicrosoftAccountVerificationFailed),
		string(DeviceLicensingStatus_RefreshingDeviceLicense),
		string(DeviceLicensingStatus_RemovingDeviceLicense),
		string(DeviceLicensingStatus_Unknown),
		string(DeviceLicensingStatus_VerifyingMicrosoftAccountIdentity),
		string(DeviceLicensingStatus_VerifyingMicrosoftDeviceIdentity),
	}
}

func (s *DeviceLicensingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceLicensingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceLicensingStatus(input string) (*DeviceLicensingStatus, error) {
	vals := map[string]DeviceLicensingStatus{
		"acquiringdevicelicense":                DeviceLicensingStatus_AcquiringDeviceLicense,
		"deviceidentityverificationfailed":      DeviceLicensingStatus_DeviceIdentityVerificationFailed,
		"deviceisnotazureactivedirectoryjoined": DeviceLicensingStatus_DeviceIsNotAzureActiveDirectoryJoined,
		"devicelicenserefreshfailed":            DeviceLicensingStatus_DeviceLicenseRefreshFailed,
		"devicelicenserefreshsucceed":           DeviceLicensingStatus_DeviceLicenseRefreshSucceed,
		"devicelicenseremovefailed":             DeviceLicensingStatus_DeviceLicenseRemoveFailed,
		"devicelicenseremovesucceed":            DeviceLicensingStatus_DeviceLicenseRemoveSucceed,
		"licenserefreshpending":                 DeviceLicensingStatus_LicenseRefreshPending,
		"licenserefreshstarted":                 DeviceLicensingStatus_LicenseRefreshStarted,
		"microsoftaccountverificationfailed":    DeviceLicensingStatus_MicrosoftAccountVerificationFailed,
		"refreshingdevicelicense":               DeviceLicensingStatus_RefreshingDeviceLicense,
		"removingdevicelicense":                 DeviceLicensingStatus_RemovingDeviceLicense,
		"unknown":                               DeviceLicensingStatus_Unknown,
		"verifyingmicrosoftaccountidentity":     DeviceLicensingStatus_VerifyingMicrosoftAccountIdentity,
		"verifyingmicrosoftdeviceidentity":      DeviceLicensingStatus_VerifyingMicrosoftDeviceIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLicensingStatus(input)
	return &out, nil
}
