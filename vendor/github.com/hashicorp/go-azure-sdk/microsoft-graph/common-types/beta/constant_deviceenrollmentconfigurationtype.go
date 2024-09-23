package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentConfigurationType string

const (
	DeviceEnrollmentConfigurationType_DefaultLimit                                          DeviceEnrollmentConfigurationType = "defaultLimit"
	DeviceEnrollmentConfigurationType_DefaultPlatformRestrictions                           DeviceEnrollmentConfigurationType = "defaultPlatformRestrictions"
	DeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration DeviceEnrollmentConfigurationType = "defaultWindows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness                        DeviceEnrollmentConfigurationType = "defaultWindowsHelloForBusiness"
	DeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration              DeviceEnrollmentConfigurationType = "deviceComanagementAuthorityConfiguration"
	DeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration                  DeviceEnrollmentConfigurationType = "enrollmentNotificationsConfiguration"
	DeviceEnrollmentConfigurationType_Limit                                                 DeviceEnrollmentConfigurationType = "limit"
	DeviceEnrollmentConfigurationType_PlatformRestrictions                                  DeviceEnrollmentConfigurationType = "platformRestrictions"
	DeviceEnrollmentConfigurationType_SinglePlatformRestriction                             DeviceEnrollmentConfigurationType = "singlePlatformRestriction"
	DeviceEnrollmentConfigurationType_Unknown                                               DeviceEnrollmentConfigurationType = "unknown"
	DeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration        DeviceEnrollmentConfigurationType = "windows10EnrollmentCompletionPageConfiguration"
	DeviceEnrollmentConfigurationType_WindowsHelloForBusiness                               DeviceEnrollmentConfigurationType = "windowsHelloForBusiness"
)

func PossibleValuesForDeviceEnrollmentConfigurationType() []string {
	return []string{
		string(DeviceEnrollmentConfigurationType_DefaultLimit),
		string(DeviceEnrollmentConfigurationType_DefaultPlatformRestrictions),
		string(DeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness),
		string(DeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration),
		string(DeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration),
		string(DeviceEnrollmentConfigurationType_Limit),
		string(DeviceEnrollmentConfigurationType_PlatformRestrictions),
		string(DeviceEnrollmentConfigurationType_SinglePlatformRestriction),
		string(DeviceEnrollmentConfigurationType_Unknown),
		string(DeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration),
		string(DeviceEnrollmentConfigurationType_WindowsHelloForBusiness),
	}
}

func (s *DeviceEnrollmentConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentConfigurationType(input string) (*DeviceEnrollmentConfigurationType, error) {
	vals := map[string]DeviceEnrollmentConfigurationType{
		"defaultlimit":                DeviceEnrollmentConfigurationType_DefaultLimit,
		"defaultplatformrestrictions": DeviceEnrollmentConfigurationType_DefaultPlatformRestrictions,
		"defaultwindows10enrollmentcompletionpageconfiguration": DeviceEnrollmentConfigurationType_DefaultWindows10EnrollmentCompletionPageConfiguration,
		"defaultwindowshelloforbusiness":                        DeviceEnrollmentConfigurationType_DefaultWindowsHelloForBusiness,
		"devicecomanagementauthorityconfiguration":              DeviceEnrollmentConfigurationType_DeviceComanagementAuthorityConfiguration,
		"enrollmentnotificationsconfiguration":                  DeviceEnrollmentConfigurationType_EnrollmentNotificationsConfiguration,
		"limit":                                                 DeviceEnrollmentConfigurationType_Limit,
		"platformrestrictions":                                  DeviceEnrollmentConfigurationType_PlatformRestrictions,
		"singleplatformrestriction":                             DeviceEnrollmentConfigurationType_SinglePlatformRestriction,
		"unknown":                                               DeviceEnrollmentConfigurationType_Unknown,
		"windows10enrollmentcompletionpageconfiguration":        DeviceEnrollmentConfigurationType_Windows10EnrollmentCompletionPageConfiguration,
		"windowshelloforbusiness":                               DeviceEnrollmentConfigurationType_WindowsHelloForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentConfigurationType(input)
	return &out, nil
}
