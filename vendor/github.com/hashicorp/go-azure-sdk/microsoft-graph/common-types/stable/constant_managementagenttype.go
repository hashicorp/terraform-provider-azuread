package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementAgentType string

const (
	ManagementAgentType_ConfigurationManagerClient        ManagementAgentType = "configurationManagerClient"
	ManagementAgentType_ConfigurationManagerClientMdm     ManagementAgentType = "configurationManagerClientMdm"
	ManagementAgentType_ConfigurationManagerClientMdmEas  ManagementAgentType = "configurationManagerClientMdmEas"
	ManagementAgentType_Eas                               ManagementAgentType = "eas"
	ManagementAgentType_EasIntuneClient                   ManagementAgentType = "easIntuneClient"
	ManagementAgentType_EasMdm                            ManagementAgentType = "easMdm"
	ManagementAgentType_GoogleCloudDevicePolicyController ManagementAgentType = "googleCloudDevicePolicyController"
	ManagementAgentType_IntuneClient                      ManagementAgentType = "intuneClient"
	ManagementAgentType_Jamf                              ManagementAgentType = "jamf"
	ManagementAgentType_Mdm                               ManagementAgentType = "mdm"
	ManagementAgentType_Microsoft365ManagedMdm            ManagementAgentType = "microsoft365ManagedMdm"
	ManagementAgentType_MsSense                           ManagementAgentType = "msSense"
	ManagementAgentType_Unknown                           ManagementAgentType = "unknown"
)

func PossibleValuesForManagementAgentType() []string {
	return []string{
		string(ManagementAgentType_ConfigurationManagerClient),
		string(ManagementAgentType_ConfigurationManagerClientMdm),
		string(ManagementAgentType_ConfigurationManagerClientMdmEas),
		string(ManagementAgentType_Eas),
		string(ManagementAgentType_EasIntuneClient),
		string(ManagementAgentType_EasMdm),
		string(ManagementAgentType_GoogleCloudDevicePolicyController),
		string(ManagementAgentType_IntuneClient),
		string(ManagementAgentType_Jamf),
		string(ManagementAgentType_Mdm),
		string(ManagementAgentType_Microsoft365ManagedMdm),
		string(ManagementAgentType_MsSense),
		string(ManagementAgentType_Unknown),
	}
}

func (s *ManagementAgentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagementAgentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagementAgentType(input string) (*ManagementAgentType, error) {
	vals := map[string]ManagementAgentType{
		"configurationmanagerclient":        ManagementAgentType_ConfigurationManagerClient,
		"configurationmanagerclientmdm":     ManagementAgentType_ConfigurationManagerClientMdm,
		"configurationmanagerclientmdmeas":  ManagementAgentType_ConfigurationManagerClientMdmEas,
		"eas":                               ManagementAgentType_Eas,
		"easintuneclient":                   ManagementAgentType_EasIntuneClient,
		"easmdm":                            ManagementAgentType_EasMdm,
		"googleclouddevicepolicycontroller": ManagementAgentType_GoogleCloudDevicePolicyController,
		"intuneclient":                      ManagementAgentType_IntuneClient,
		"jamf":                              ManagementAgentType_Jamf,
		"mdm":                               ManagementAgentType_Mdm,
		"microsoft365managedmdm":            ManagementAgentType_Microsoft365ManagedMdm,
		"mssense":                           ManagementAgentType_MsSense,
		"unknown":                           ManagementAgentType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementAgentType(input)
	return &out, nil
}
