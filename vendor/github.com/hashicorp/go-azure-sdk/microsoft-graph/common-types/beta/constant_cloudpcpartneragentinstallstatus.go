package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPartnerAgentInstallStatus string

const (
	CloudPCPartnerAgentInstallStatus_InstallFailed   CloudPCPartnerAgentInstallStatus = "installFailed"
	CloudPCPartnerAgentInstallStatus_Installed       CloudPCPartnerAgentInstallStatus = "installed"
	CloudPCPartnerAgentInstallStatus_Installing      CloudPCPartnerAgentInstallStatus = "installing"
	CloudPCPartnerAgentInstallStatus_Licensed        CloudPCPartnerAgentInstallStatus = "licensed"
	CloudPCPartnerAgentInstallStatus_UninstallFailed CloudPCPartnerAgentInstallStatus = "uninstallFailed"
	CloudPCPartnerAgentInstallStatus_Uninstalling    CloudPCPartnerAgentInstallStatus = "uninstalling"
)

func PossibleValuesForCloudPCPartnerAgentInstallStatus() []string {
	return []string{
		string(CloudPCPartnerAgentInstallStatus_InstallFailed),
		string(CloudPCPartnerAgentInstallStatus_Installed),
		string(CloudPCPartnerAgentInstallStatus_Installing),
		string(CloudPCPartnerAgentInstallStatus_Licensed),
		string(CloudPCPartnerAgentInstallStatus_UninstallFailed),
		string(CloudPCPartnerAgentInstallStatus_Uninstalling),
	}
}

func (s *CloudPCPartnerAgentInstallStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPartnerAgentInstallStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPartnerAgentInstallStatus(input string) (*CloudPCPartnerAgentInstallStatus, error) {
	vals := map[string]CloudPCPartnerAgentInstallStatus{
		"installfailed":   CloudPCPartnerAgentInstallStatus_InstallFailed,
		"installed":       CloudPCPartnerAgentInstallStatus_Installed,
		"installing":      CloudPCPartnerAgentInstallStatus_Installing,
		"licensed":        CloudPCPartnerAgentInstallStatus_Licensed,
		"uninstallfailed": CloudPCPartnerAgentInstallStatus_UninstallFailed,
		"uninstalling":    CloudPCPartnerAgentInstallStatus_Uninstalling,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPartnerAgentInstallStatus(input)
	return &out, nil
}
