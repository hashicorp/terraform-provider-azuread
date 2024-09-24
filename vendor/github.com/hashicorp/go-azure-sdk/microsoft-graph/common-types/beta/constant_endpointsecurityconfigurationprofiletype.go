package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointSecurityConfigurationProfileType string

const (
	EndpointSecurityConfigurationProfileType_AccountProtection            EndpointSecurityConfigurationProfileType = "accountProtection"
	EndpointSecurityConfigurationProfileType_Antivirus                    EndpointSecurityConfigurationProfileType = "antivirus"
	EndpointSecurityConfigurationProfileType_AppAndBrowserIsolation       EndpointSecurityConfigurationProfileType = "appAndBrowserIsolation"
	EndpointSecurityConfigurationProfileType_ApplicationControl           EndpointSecurityConfigurationProfileType = "applicationControl"
	EndpointSecurityConfigurationProfileType_AttackSurfaceReductionRules  EndpointSecurityConfigurationProfileType = "attackSurfaceReductionRules"
	EndpointSecurityConfigurationProfileType_BitLocker                    EndpointSecurityConfigurationProfileType = "bitLocker"
	EndpointSecurityConfigurationProfileType_DeviceControl                EndpointSecurityConfigurationProfileType = "deviceControl"
	EndpointSecurityConfigurationProfileType_EndpointDetectionAndResponse EndpointSecurityConfigurationProfileType = "endpointDetectionAndResponse"
	EndpointSecurityConfigurationProfileType_ExploitProtection            EndpointSecurityConfigurationProfileType = "exploitProtection"
	EndpointSecurityConfigurationProfileType_FileVault                    EndpointSecurityConfigurationProfileType = "fileVault"
	EndpointSecurityConfigurationProfileType_Firewall                     EndpointSecurityConfigurationProfileType = "firewall"
	EndpointSecurityConfigurationProfileType_FirewallRules                EndpointSecurityConfigurationProfileType = "firewallRules"
	EndpointSecurityConfigurationProfileType_Unknown                      EndpointSecurityConfigurationProfileType = "unknown"
	EndpointSecurityConfigurationProfileType_WebProtection                EndpointSecurityConfigurationProfileType = "webProtection"
	EndpointSecurityConfigurationProfileType_WindowsSecurity              EndpointSecurityConfigurationProfileType = "windowsSecurity"
)

func PossibleValuesForEndpointSecurityConfigurationProfileType() []string {
	return []string{
		string(EndpointSecurityConfigurationProfileType_AccountProtection),
		string(EndpointSecurityConfigurationProfileType_Antivirus),
		string(EndpointSecurityConfigurationProfileType_AppAndBrowserIsolation),
		string(EndpointSecurityConfigurationProfileType_ApplicationControl),
		string(EndpointSecurityConfigurationProfileType_AttackSurfaceReductionRules),
		string(EndpointSecurityConfigurationProfileType_BitLocker),
		string(EndpointSecurityConfigurationProfileType_DeviceControl),
		string(EndpointSecurityConfigurationProfileType_EndpointDetectionAndResponse),
		string(EndpointSecurityConfigurationProfileType_ExploitProtection),
		string(EndpointSecurityConfigurationProfileType_FileVault),
		string(EndpointSecurityConfigurationProfileType_Firewall),
		string(EndpointSecurityConfigurationProfileType_FirewallRules),
		string(EndpointSecurityConfigurationProfileType_Unknown),
		string(EndpointSecurityConfigurationProfileType_WebProtection),
		string(EndpointSecurityConfigurationProfileType_WindowsSecurity),
	}
}

func (s *EndpointSecurityConfigurationProfileType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointSecurityConfigurationProfileType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointSecurityConfigurationProfileType(input string) (*EndpointSecurityConfigurationProfileType, error) {
	vals := map[string]EndpointSecurityConfigurationProfileType{
		"accountprotection":            EndpointSecurityConfigurationProfileType_AccountProtection,
		"antivirus":                    EndpointSecurityConfigurationProfileType_Antivirus,
		"appandbrowserisolation":       EndpointSecurityConfigurationProfileType_AppAndBrowserIsolation,
		"applicationcontrol":           EndpointSecurityConfigurationProfileType_ApplicationControl,
		"attacksurfacereductionrules":  EndpointSecurityConfigurationProfileType_AttackSurfaceReductionRules,
		"bitlocker":                    EndpointSecurityConfigurationProfileType_BitLocker,
		"devicecontrol":                EndpointSecurityConfigurationProfileType_DeviceControl,
		"endpointdetectionandresponse": EndpointSecurityConfigurationProfileType_EndpointDetectionAndResponse,
		"exploitprotection":            EndpointSecurityConfigurationProfileType_ExploitProtection,
		"filevault":                    EndpointSecurityConfigurationProfileType_FileVault,
		"firewall":                     EndpointSecurityConfigurationProfileType_Firewall,
		"firewallrules":                EndpointSecurityConfigurationProfileType_FirewallRules,
		"unknown":                      EndpointSecurityConfigurationProfileType_Unknown,
		"webprotection":                EndpointSecurityConfigurationProfileType_WebProtection,
		"windowssecurity":              EndpointSecurityConfigurationProfileType_WindowsSecurity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointSecurityConfigurationProfileType(input)
	return &out, nil
}
