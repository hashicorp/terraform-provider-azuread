package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplateSubtype string

const (
	DeviceManagementTemplateSubtype_AccountProtection        DeviceManagementTemplateSubtype = "accountProtection"
	DeviceManagementTemplateSubtype_Antivirus                DeviceManagementTemplateSubtype = "antivirus"
	DeviceManagementTemplateSubtype_AttackSurfaceReduction   DeviceManagementTemplateSubtype = "attackSurfaceReduction"
	DeviceManagementTemplateSubtype_DiskEncryption           DeviceManagementTemplateSubtype = "diskEncryption"
	DeviceManagementTemplateSubtype_EndpointDetectionReponse DeviceManagementTemplateSubtype = "endpointDetectionReponse"
	DeviceManagementTemplateSubtype_Firewall                 DeviceManagementTemplateSubtype = "firewall"
	DeviceManagementTemplateSubtype_FirewallSharedAppList    DeviceManagementTemplateSubtype = "firewallSharedAppList"
	DeviceManagementTemplateSubtype_FirewallSharedIPList     DeviceManagementTemplateSubtype = "firewallSharedIpList"
	DeviceManagementTemplateSubtype_FirewallSharedPortlist   DeviceManagementTemplateSubtype = "firewallSharedPortlist"
	DeviceManagementTemplateSubtype_None                     DeviceManagementTemplateSubtype = "none"
)

func PossibleValuesForDeviceManagementTemplateSubtype() []string {
	return []string{
		string(DeviceManagementTemplateSubtype_AccountProtection),
		string(DeviceManagementTemplateSubtype_Antivirus),
		string(DeviceManagementTemplateSubtype_AttackSurfaceReduction),
		string(DeviceManagementTemplateSubtype_DiskEncryption),
		string(DeviceManagementTemplateSubtype_EndpointDetectionReponse),
		string(DeviceManagementTemplateSubtype_Firewall),
		string(DeviceManagementTemplateSubtype_FirewallSharedAppList),
		string(DeviceManagementTemplateSubtype_FirewallSharedIPList),
		string(DeviceManagementTemplateSubtype_FirewallSharedPortlist),
		string(DeviceManagementTemplateSubtype_None),
	}
}

func (s *DeviceManagementTemplateSubtype) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementTemplateSubtype(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementTemplateSubtype(input string) (*DeviceManagementTemplateSubtype, error) {
	vals := map[string]DeviceManagementTemplateSubtype{
		"accountprotection":        DeviceManagementTemplateSubtype_AccountProtection,
		"antivirus":                DeviceManagementTemplateSubtype_Antivirus,
		"attacksurfacereduction":   DeviceManagementTemplateSubtype_AttackSurfaceReduction,
		"diskencryption":           DeviceManagementTemplateSubtype_DiskEncryption,
		"endpointdetectionreponse": DeviceManagementTemplateSubtype_EndpointDetectionReponse,
		"firewall":                 DeviceManagementTemplateSubtype_Firewall,
		"firewallsharedapplist":    DeviceManagementTemplateSubtype_FirewallSharedAppList,
		"firewallsharediplist":     DeviceManagementTemplateSubtype_FirewallSharedIPList,
		"firewallsharedportlist":   DeviceManagementTemplateSubtype_FirewallSharedPortlist,
		"none":                     DeviceManagementTemplateSubtype_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplateSubtype(input)
	return &out, nil
}
