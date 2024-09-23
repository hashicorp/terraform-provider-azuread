package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationTemplateFamily string

const (
	DeviceManagementConfigurationTemplateFamily_AppQuietTime                                 DeviceManagementConfigurationTemplateFamily = "appQuietTime"
	DeviceManagementConfigurationTemplateFamily_Baseline                                     DeviceManagementConfigurationTemplateFamily = "baseline"
	DeviceManagementConfigurationTemplateFamily_CompanyPortal                                DeviceManagementConfigurationTemplateFamily = "companyPortal"
	DeviceManagementConfigurationTemplateFamily_DeviceConfigurationPolicies                  DeviceManagementConfigurationTemplateFamily = "deviceConfigurationPolicies"
	DeviceManagementConfigurationTemplateFamily_DeviceConfigurationScripts                   DeviceManagementConfigurationTemplateFamily = "deviceConfigurationScripts"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityAccountProtection            DeviceManagementConfigurationTemplateFamily = "endpointSecurityAccountProtection"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityAntivirus                    DeviceManagementConfigurationTemplateFamily = "endpointSecurityAntivirus"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityApplicationControl           DeviceManagementConfigurationTemplateFamily = "endpointSecurityApplicationControl"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityAttackSurfaceReduction       DeviceManagementConfigurationTemplateFamily = "endpointSecurityAttackSurfaceReduction"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityDiskEncryption               DeviceManagementConfigurationTemplateFamily = "endpointSecurityDiskEncryption"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityEndpointDetectionAndResponse DeviceManagementConfigurationTemplateFamily = "endpointSecurityEndpointDetectionAndResponse"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityEndpointPrivilegeManagement  DeviceManagementConfigurationTemplateFamily = "endpointSecurityEndpointPrivilegeManagement"
	DeviceManagementConfigurationTemplateFamily_EndpointSecurityFirewall                     DeviceManagementConfigurationTemplateFamily = "endpointSecurityFirewall"
	DeviceManagementConfigurationTemplateFamily_EnrollmentConfiguration                      DeviceManagementConfigurationTemplateFamily = "enrollmentConfiguration"
	DeviceManagementConfigurationTemplateFamily_None                                         DeviceManagementConfigurationTemplateFamily = "none"
	DeviceManagementConfigurationTemplateFamily_WindowsOsRecoveryPolicies                    DeviceManagementConfigurationTemplateFamily = "windowsOsRecoveryPolicies"
)

func PossibleValuesForDeviceManagementConfigurationTemplateFamily() []string {
	return []string{
		string(DeviceManagementConfigurationTemplateFamily_AppQuietTime),
		string(DeviceManagementConfigurationTemplateFamily_Baseline),
		string(DeviceManagementConfigurationTemplateFamily_CompanyPortal),
		string(DeviceManagementConfigurationTemplateFamily_DeviceConfigurationPolicies),
		string(DeviceManagementConfigurationTemplateFamily_DeviceConfigurationScripts),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityAccountProtection),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityAntivirus),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityApplicationControl),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityAttackSurfaceReduction),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityDiskEncryption),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityEndpointDetectionAndResponse),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityEndpointPrivilegeManagement),
		string(DeviceManagementConfigurationTemplateFamily_EndpointSecurityFirewall),
		string(DeviceManagementConfigurationTemplateFamily_EnrollmentConfiguration),
		string(DeviceManagementConfigurationTemplateFamily_None),
		string(DeviceManagementConfigurationTemplateFamily_WindowsOsRecoveryPolicies),
	}
}

func (s *DeviceManagementConfigurationTemplateFamily) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationTemplateFamily(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationTemplateFamily(input string) (*DeviceManagementConfigurationTemplateFamily, error) {
	vals := map[string]DeviceManagementConfigurationTemplateFamily{
		"appquiettime":                                 DeviceManagementConfigurationTemplateFamily_AppQuietTime,
		"baseline":                                     DeviceManagementConfigurationTemplateFamily_Baseline,
		"companyportal":                                DeviceManagementConfigurationTemplateFamily_CompanyPortal,
		"deviceconfigurationpolicies":                  DeviceManagementConfigurationTemplateFamily_DeviceConfigurationPolicies,
		"deviceconfigurationscripts":                   DeviceManagementConfigurationTemplateFamily_DeviceConfigurationScripts,
		"endpointsecurityaccountprotection":            DeviceManagementConfigurationTemplateFamily_EndpointSecurityAccountProtection,
		"endpointsecurityantivirus":                    DeviceManagementConfigurationTemplateFamily_EndpointSecurityAntivirus,
		"endpointsecurityapplicationcontrol":           DeviceManagementConfigurationTemplateFamily_EndpointSecurityApplicationControl,
		"endpointsecurityattacksurfacereduction":       DeviceManagementConfigurationTemplateFamily_EndpointSecurityAttackSurfaceReduction,
		"endpointsecuritydiskencryption":               DeviceManagementConfigurationTemplateFamily_EndpointSecurityDiskEncryption,
		"endpointsecurityendpointdetectionandresponse": DeviceManagementConfigurationTemplateFamily_EndpointSecurityEndpointDetectionAndResponse,
		"endpointsecurityendpointprivilegemanagement":  DeviceManagementConfigurationTemplateFamily_EndpointSecurityEndpointPrivilegeManagement,
		"endpointsecurityfirewall":                     DeviceManagementConfigurationTemplateFamily_EndpointSecurityFirewall,
		"enrollmentconfiguration":                      DeviceManagementConfigurationTemplateFamily_EnrollmentConfiguration,
		"none":                                         DeviceManagementConfigurationTemplateFamily_None,
		"windowsosrecoverypolicies":                    DeviceManagementConfigurationTemplateFamily_WindowsOsRecoveryPolicies,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationTemplateFamily(input)
	return &out, nil
}
