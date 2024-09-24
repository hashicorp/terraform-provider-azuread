package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientState string

const (
	ConfigurationManagerClientState_CommunicationError ConfigurationManagerClientState = "communicationError"
	ConfigurationManagerClientState_Healthy            ConfigurationManagerClientState = "healthy"
	ConfigurationManagerClientState_InstallFailed      ConfigurationManagerClientState = "installFailed"
	ConfigurationManagerClientState_Installed          ConfigurationManagerClientState = "installed"
	ConfigurationManagerClientState_Unknown            ConfigurationManagerClientState = "unknown"
	ConfigurationManagerClientState_UpdateFailed       ConfigurationManagerClientState = "updateFailed"
)

func PossibleValuesForConfigurationManagerClientState() []string {
	return []string{
		string(ConfigurationManagerClientState_CommunicationError),
		string(ConfigurationManagerClientState_Healthy),
		string(ConfigurationManagerClientState_InstallFailed),
		string(ConfigurationManagerClientState_Installed),
		string(ConfigurationManagerClientState_Unknown),
		string(ConfigurationManagerClientState_UpdateFailed),
	}
}

func (s *ConfigurationManagerClientState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerClientState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerClientState(input string) (*ConfigurationManagerClientState, error) {
	vals := map[string]ConfigurationManagerClientState{
		"communicationerror": ConfigurationManagerClientState_CommunicationError,
		"healthy":            ConfigurationManagerClientState_Healthy,
		"installfailed":      ConfigurationManagerClientState_InstallFailed,
		"installed":          ConfigurationManagerClientState_Installed,
		"unknown":            ConfigurationManagerClientState_Unknown,
		"updatefailed":       ConfigurationManagerClientState_UpdateFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerClientState(input)
	return &out, nil
}
