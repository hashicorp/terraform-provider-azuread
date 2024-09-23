package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointSecurityConfigurationType string

const (
	EndpointSecurityConfigurationType_AccountProtection            EndpointSecurityConfigurationType = "accountProtection"
	EndpointSecurityConfigurationType_Antivirus                    EndpointSecurityConfigurationType = "antivirus"
	EndpointSecurityConfigurationType_AttackSurfaceReduction       EndpointSecurityConfigurationType = "attackSurfaceReduction"
	EndpointSecurityConfigurationType_DiskEncryption               EndpointSecurityConfigurationType = "diskEncryption"
	EndpointSecurityConfigurationType_EndpointDetectionAndResponse EndpointSecurityConfigurationType = "endpointDetectionAndResponse"
	EndpointSecurityConfigurationType_Firewall                     EndpointSecurityConfigurationType = "firewall"
	EndpointSecurityConfigurationType_Unknown                      EndpointSecurityConfigurationType = "unknown"
)

func PossibleValuesForEndpointSecurityConfigurationType() []string {
	return []string{
		string(EndpointSecurityConfigurationType_AccountProtection),
		string(EndpointSecurityConfigurationType_Antivirus),
		string(EndpointSecurityConfigurationType_AttackSurfaceReduction),
		string(EndpointSecurityConfigurationType_DiskEncryption),
		string(EndpointSecurityConfigurationType_EndpointDetectionAndResponse),
		string(EndpointSecurityConfigurationType_Firewall),
		string(EndpointSecurityConfigurationType_Unknown),
	}
}

func (s *EndpointSecurityConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointSecurityConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointSecurityConfigurationType(input string) (*EndpointSecurityConfigurationType, error) {
	vals := map[string]EndpointSecurityConfigurationType{
		"accountprotection":            EndpointSecurityConfigurationType_AccountProtection,
		"antivirus":                    EndpointSecurityConfigurationType_Antivirus,
		"attacksurfacereduction":       EndpointSecurityConfigurationType_AttackSurfaceReduction,
		"diskencryption":               EndpointSecurityConfigurationType_DiskEncryption,
		"endpointdetectionandresponse": EndpointSecurityConfigurationType_EndpointDetectionAndResponse,
		"firewall":                     EndpointSecurityConfigurationType_Firewall,
		"unknown":                      EndpointSecurityConfigurationType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointSecurityConfigurationType(input)
	return &out, nil
}
