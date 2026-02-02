package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssociatedAssignmentPayloadType string

const (
	AssociatedAssignmentPayloadType_AndroidEnterpriseApp                            AssociatedAssignmentPayloadType = "androidEnterpriseApp"
	AssociatedAssignmentPayloadType_AndroidEnterpriseConfiguration                  AssociatedAssignmentPayloadType = "androidEnterpriseConfiguration"
	AssociatedAssignmentPayloadType_Application                                     AssociatedAssignmentPayloadType = "application"
	AssociatedAssignmentPayloadType_DeviceConfigurationAndCompliance                AssociatedAssignmentPayloadType = "deviceConfigurationAndCompliance"
	AssociatedAssignmentPayloadType_DeviceFirmwareConfigurationInterfacePolicy      AssociatedAssignmentPayloadType = "deviceFirmwareConfigurationInterfacePolicy"
	AssociatedAssignmentPayloadType_DeviceManagmentConfigurationAndCompliancePolicy AssociatedAssignmentPayloadType = "deviceManagmentConfigurationAndCompliancePolicy"
	AssociatedAssignmentPayloadType_EnrollmentConfiguration                         AssociatedAssignmentPayloadType = "enrollmentConfiguration"
	AssociatedAssignmentPayloadType_GroupPolicyConfiguration                        AssociatedAssignmentPayloadType = "groupPolicyConfiguration"
	AssociatedAssignmentPayloadType_ResourceAccessPolicy                            AssociatedAssignmentPayloadType = "resourceAccessPolicy"
	AssociatedAssignmentPayloadType_Unknown                                         AssociatedAssignmentPayloadType = "unknown"
	AssociatedAssignmentPayloadType_Win32app                                        AssociatedAssignmentPayloadType = "win32app"
	AssociatedAssignmentPayloadType_ZeroTouchDeploymentDeviceConfigProfile          AssociatedAssignmentPayloadType = "zeroTouchDeploymentDeviceConfigProfile"
)

func PossibleValuesForAssociatedAssignmentPayloadType() []string {
	return []string{
		string(AssociatedAssignmentPayloadType_AndroidEnterpriseApp),
		string(AssociatedAssignmentPayloadType_AndroidEnterpriseConfiguration),
		string(AssociatedAssignmentPayloadType_Application),
		string(AssociatedAssignmentPayloadType_DeviceConfigurationAndCompliance),
		string(AssociatedAssignmentPayloadType_DeviceFirmwareConfigurationInterfacePolicy),
		string(AssociatedAssignmentPayloadType_DeviceManagmentConfigurationAndCompliancePolicy),
		string(AssociatedAssignmentPayloadType_EnrollmentConfiguration),
		string(AssociatedAssignmentPayloadType_GroupPolicyConfiguration),
		string(AssociatedAssignmentPayloadType_ResourceAccessPolicy),
		string(AssociatedAssignmentPayloadType_Unknown),
		string(AssociatedAssignmentPayloadType_Win32app),
		string(AssociatedAssignmentPayloadType_ZeroTouchDeploymentDeviceConfigProfile),
	}
}

func (s *AssociatedAssignmentPayloadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssociatedAssignmentPayloadType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssociatedAssignmentPayloadType(input string) (*AssociatedAssignmentPayloadType, error) {
	vals := map[string]AssociatedAssignmentPayloadType{
		"androidenterpriseapp":                            AssociatedAssignmentPayloadType_AndroidEnterpriseApp,
		"androidenterpriseconfiguration":                  AssociatedAssignmentPayloadType_AndroidEnterpriseConfiguration,
		"application":                                     AssociatedAssignmentPayloadType_Application,
		"deviceconfigurationandcompliance":                AssociatedAssignmentPayloadType_DeviceConfigurationAndCompliance,
		"devicefirmwareconfigurationinterfacepolicy":      AssociatedAssignmentPayloadType_DeviceFirmwareConfigurationInterfacePolicy,
		"devicemanagmentconfigurationandcompliancepolicy": AssociatedAssignmentPayloadType_DeviceManagmentConfigurationAndCompliancePolicy,
		"enrollmentconfiguration":                         AssociatedAssignmentPayloadType_EnrollmentConfiguration,
		"grouppolicyconfiguration":                        AssociatedAssignmentPayloadType_GroupPolicyConfiguration,
		"resourceaccesspolicy":                            AssociatedAssignmentPayloadType_ResourceAccessPolicy,
		"unknown":                                         AssociatedAssignmentPayloadType_Unknown,
		"win32app":                                        AssociatedAssignmentPayloadType_Win32app,
		"zerotouchdeploymentdeviceconfigprofile":          AssociatedAssignmentPayloadType_ZeroTouchDeploymentDeviceConfigProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssociatedAssignmentPayloadType(input)
	return &out, nil
}
