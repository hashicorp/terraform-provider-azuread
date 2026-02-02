package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplateType string

const (
	DeviceManagementTemplateType_AdvancedThreatProtectionSecurityBaseline  DeviceManagementTemplateType = "advancedThreatProtectionSecurityBaseline"
	DeviceManagementTemplateType_CloudPC                                   DeviceManagementTemplateType = "cloudPC"
	DeviceManagementTemplateType_Custom                                    DeviceManagementTemplateType = "custom"
	DeviceManagementTemplateType_DeviceCompliance                          DeviceManagementTemplateType = "deviceCompliance"
	DeviceManagementTemplateType_DeviceConfiguration                       DeviceManagementTemplateType = "deviceConfiguration"
	DeviceManagementTemplateType_DeviceConfigurationForOffice365           DeviceManagementTemplateType = "deviceConfigurationForOffice365"
	DeviceManagementTemplateType_FirewallSharedSettings                    DeviceManagementTemplateType = "firewallSharedSettings"
	DeviceManagementTemplateType_MicrosoftEdgeSecurityBaseline             DeviceManagementTemplateType = "microsoftEdgeSecurityBaseline"
	DeviceManagementTemplateType_MicrosoftOffice365ProPlusSecurityBaseline DeviceManagementTemplateType = "microsoftOffice365ProPlusSecurityBaseline"
	DeviceManagementTemplateType_SecurityBaseline                          DeviceManagementTemplateType = "securityBaseline"
	DeviceManagementTemplateType_SecurityTemplate                          DeviceManagementTemplateType = "securityTemplate"
	DeviceManagementTemplateType_SpecializedDevices                        DeviceManagementTemplateType = "specializedDevices"
)

func PossibleValuesForDeviceManagementTemplateType() []string {
	return []string{
		string(DeviceManagementTemplateType_AdvancedThreatProtectionSecurityBaseline),
		string(DeviceManagementTemplateType_CloudPC),
		string(DeviceManagementTemplateType_Custom),
		string(DeviceManagementTemplateType_DeviceCompliance),
		string(DeviceManagementTemplateType_DeviceConfiguration),
		string(DeviceManagementTemplateType_DeviceConfigurationForOffice365),
		string(DeviceManagementTemplateType_FirewallSharedSettings),
		string(DeviceManagementTemplateType_MicrosoftEdgeSecurityBaseline),
		string(DeviceManagementTemplateType_MicrosoftOffice365ProPlusSecurityBaseline),
		string(DeviceManagementTemplateType_SecurityBaseline),
		string(DeviceManagementTemplateType_SecurityTemplate),
		string(DeviceManagementTemplateType_SpecializedDevices),
	}
}

func (s *DeviceManagementTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementTemplateType(input string) (*DeviceManagementTemplateType, error) {
	vals := map[string]DeviceManagementTemplateType{
		"advancedthreatprotectionsecuritybaseline": DeviceManagementTemplateType_AdvancedThreatProtectionSecurityBaseline,
		"cloudpc":                                   DeviceManagementTemplateType_CloudPC,
		"custom":                                    DeviceManagementTemplateType_Custom,
		"devicecompliance":                          DeviceManagementTemplateType_DeviceCompliance,
		"deviceconfiguration":                       DeviceManagementTemplateType_DeviceConfiguration,
		"deviceconfigurationforoffice365":           DeviceManagementTemplateType_DeviceConfigurationForOffice365,
		"firewallsharedsettings":                    DeviceManagementTemplateType_FirewallSharedSettings,
		"microsoftedgesecuritybaseline":             DeviceManagementTemplateType_MicrosoftEdgeSecurityBaseline,
		"microsoftoffice365proplussecuritybaseline": DeviceManagementTemplateType_MicrosoftOffice365ProPlusSecurityBaseline,
		"securitybaseline":                          DeviceManagementTemplateType_SecurityBaseline,
		"securitytemplate":                          DeviceManagementTemplateType_SecurityTemplate,
		"specializeddevices":                        DeviceManagementTemplateType_SpecializedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplateType(input)
	return &out, nil
}
