package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentType string

const (
	DeviceEnrollmentType_AndroidAOSPUserOwnedDeviceEnrollment  DeviceEnrollmentType = "androidAOSPUserOwnedDeviceEnrollment"
	DeviceEnrollmentType_AndroidAOSPUserlessDeviceEnrollment   DeviceEnrollmentType = "androidAOSPUserlessDeviceEnrollment"
	DeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile DeviceEnrollmentType = "androidEnterpriseCorporateWorkProfile"
	DeviceEnrollmentType_AndroidEnterpriseDedicatedDevice      DeviceEnrollmentType = "androidEnterpriseDedicatedDevice"
	DeviceEnrollmentType_AndroidEnterpriseFullyManaged         DeviceEnrollmentType = "androidEnterpriseFullyManaged"
	DeviceEnrollmentType_AppleBulkWithUser                     DeviceEnrollmentType = "appleBulkWithUser"
	DeviceEnrollmentType_AppleBulkWithoutUser                  DeviceEnrollmentType = "appleBulkWithoutUser"
	DeviceEnrollmentType_AppleUserEnrollment                   DeviceEnrollmentType = "appleUserEnrollment"
	DeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount DeviceEnrollmentType = "appleUserEnrollmentWithServiceAccount"
	DeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension      DeviceEnrollmentType = "azureAdJoinUsingAzureVmExtension"
	DeviceEnrollmentType_DeviceEnrollmentManager               DeviceEnrollmentType = "deviceEnrollmentManager"
	DeviceEnrollmentType_Unknown                               DeviceEnrollmentType = "unknown"
	DeviceEnrollmentType_UserEnrollment                        DeviceEnrollmentType = "userEnrollment"
	DeviceEnrollmentType_WindowsAutoEnrollment                 DeviceEnrollmentType = "windowsAutoEnrollment"
	DeviceEnrollmentType_WindowsAzureADJoin                    DeviceEnrollmentType = "windowsAzureADJoin"
	DeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth     DeviceEnrollmentType = "windowsAzureADJoinUsingDeviceAuth"
	DeviceEnrollmentType_WindowsBulkAzureDomainJoin            DeviceEnrollmentType = "windowsBulkAzureDomainJoin"
	DeviceEnrollmentType_WindowsBulkUserless                   DeviceEnrollmentType = "windowsBulkUserless"
	DeviceEnrollmentType_WindowsCoManagement                   DeviceEnrollmentType = "windowsCoManagement"
)

func PossibleValuesForDeviceEnrollmentType() []string {
	return []string{
		string(DeviceEnrollmentType_AndroidAOSPUserOwnedDeviceEnrollment),
		string(DeviceEnrollmentType_AndroidAOSPUserlessDeviceEnrollment),
		string(DeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile),
		string(DeviceEnrollmentType_AndroidEnterpriseDedicatedDevice),
		string(DeviceEnrollmentType_AndroidEnterpriseFullyManaged),
		string(DeviceEnrollmentType_AppleBulkWithUser),
		string(DeviceEnrollmentType_AppleBulkWithoutUser),
		string(DeviceEnrollmentType_AppleUserEnrollment),
		string(DeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount),
		string(DeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension),
		string(DeviceEnrollmentType_DeviceEnrollmentManager),
		string(DeviceEnrollmentType_Unknown),
		string(DeviceEnrollmentType_UserEnrollment),
		string(DeviceEnrollmentType_WindowsAutoEnrollment),
		string(DeviceEnrollmentType_WindowsAzureADJoin),
		string(DeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth),
		string(DeviceEnrollmentType_WindowsBulkAzureDomainJoin),
		string(DeviceEnrollmentType_WindowsBulkUserless),
		string(DeviceEnrollmentType_WindowsCoManagement),
	}
}

func (s *DeviceEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentType(input string) (*DeviceEnrollmentType, error) {
	vals := map[string]DeviceEnrollmentType{
		"androidaospuserowneddeviceenrollment":  DeviceEnrollmentType_AndroidAOSPUserOwnedDeviceEnrollment,
		"androidaospuserlessdeviceenrollment":   DeviceEnrollmentType_AndroidAOSPUserlessDeviceEnrollment,
		"androidenterprisecorporateworkprofile": DeviceEnrollmentType_AndroidEnterpriseCorporateWorkProfile,
		"androidenterprisededicateddevice":      DeviceEnrollmentType_AndroidEnterpriseDedicatedDevice,
		"androidenterprisefullymanaged":         DeviceEnrollmentType_AndroidEnterpriseFullyManaged,
		"applebulkwithuser":                     DeviceEnrollmentType_AppleBulkWithUser,
		"applebulkwithoutuser":                  DeviceEnrollmentType_AppleBulkWithoutUser,
		"appleuserenrollment":                   DeviceEnrollmentType_AppleUserEnrollment,
		"appleuserenrollmentwithserviceaccount": DeviceEnrollmentType_AppleUserEnrollmentWithServiceAccount,
		"azureadjoinusingazurevmextension":      DeviceEnrollmentType_AzureAdJoinUsingAzureVmExtension,
		"deviceenrollmentmanager":               DeviceEnrollmentType_DeviceEnrollmentManager,
		"unknown":                               DeviceEnrollmentType_Unknown,
		"userenrollment":                        DeviceEnrollmentType_UserEnrollment,
		"windowsautoenrollment":                 DeviceEnrollmentType_WindowsAutoEnrollment,
		"windowsazureadjoin":                    DeviceEnrollmentType_WindowsAzureADJoin,
		"windowsazureadjoinusingdeviceauth":     DeviceEnrollmentType_WindowsAzureADJoinUsingDeviceAuth,
		"windowsbulkazuredomainjoin":            DeviceEnrollmentType_WindowsBulkAzureDomainJoin,
		"windowsbulkuserless":                   DeviceEnrollmentType_WindowsBulkUserless,
		"windowscomanagement":                   DeviceEnrollmentType_WindowsCoManagement,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentType(input)
	return &out, nil
}
