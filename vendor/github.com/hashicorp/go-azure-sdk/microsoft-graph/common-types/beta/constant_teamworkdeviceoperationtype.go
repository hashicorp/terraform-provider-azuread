package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceOperationType string

const (
	TeamworkDeviceOperationType_ConfigUpdate                      TeamworkDeviceOperationType = "configUpdate"
	TeamworkDeviceOperationType_DeviceDiagnostics                 TeamworkDeviceOperationType = "deviceDiagnostics"
	TeamworkDeviceOperationType_DeviceManagementAgentConfigUpdate TeamworkDeviceOperationType = "deviceManagementAgentConfigUpdate"
	TeamworkDeviceOperationType_DeviceRestart                     TeamworkDeviceOperationType = "deviceRestart"
	TeamworkDeviceOperationType_RemoteLogin                       TeamworkDeviceOperationType = "remoteLogin"
	TeamworkDeviceOperationType_RemoteLogout                      TeamworkDeviceOperationType = "remoteLogout"
	TeamworkDeviceOperationType_SoftwareUpdate                    TeamworkDeviceOperationType = "softwareUpdate"
)

func PossibleValuesForTeamworkDeviceOperationType() []string {
	return []string{
		string(TeamworkDeviceOperationType_ConfigUpdate),
		string(TeamworkDeviceOperationType_DeviceDiagnostics),
		string(TeamworkDeviceOperationType_DeviceManagementAgentConfigUpdate),
		string(TeamworkDeviceOperationType_DeviceRestart),
		string(TeamworkDeviceOperationType_RemoteLogin),
		string(TeamworkDeviceOperationType_RemoteLogout),
		string(TeamworkDeviceOperationType_SoftwareUpdate),
	}
}

func (s *TeamworkDeviceOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkDeviceOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkDeviceOperationType(input string) (*TeamworkDeviceOperationType, error) {
	vals := map[string]TeamworkDeviceOperationType{
		"configupdate":                      TeamworkDeviceOperationType_ConfigUpdate,
		"devicediagnostics":                 TeamworkDeviceOperationType_DeviceDiagnostics,
		"devicemanagementagentconfigupdate": TeamworkDeviceOperationType_DeviceManagementAgentConfigUpdate,
		"devicerestart":                     TeamworkDeviceOperationType_DeviceRestart,
		"remotelogin":                       TeamworkDeviceOperationType_RemoteLogin,
		"remotelogout":                      TeamworkDeviceOperationType_RemoteLogout,
		"softwareupdate":                    TeamworkDeviceOperationType_SoftwareUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceOperationType(input)
	return &out, nil
}
