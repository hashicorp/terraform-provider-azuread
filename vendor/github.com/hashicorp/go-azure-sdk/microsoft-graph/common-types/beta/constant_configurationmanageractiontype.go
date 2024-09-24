package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionType string

const (
	ConfigurationManagerActionType_AppEvaluation                   ConfigurationManagerActionType = "appEvaluation"
	ConfigurationManagerActionType_FullScan                        ConfigurationManagerActionType = "fullScan"
	ConfigurationManagerActionType_QuickScan                       ConfigurationManagerActionType = "quickScan"
	ConfigurationManagerActionType_RefreshMachinePolicy            ConfigurationManagerActionType = "refreshMachinePolicy"
	ConfigurationManagerActionType_RefreshUserPolicy               ConfigurationManagerActionType = "refreshUserPolicy"
	ConfigurationManagerActionType_WakeUpClient                    ConfigurationManagerActionType = "wakeUpClient"
	ConfigurationManagerActionType_WindowsDefenderUpdateSignatures ConfigurationManagerActionType = "windowsDefenderUpdateSignatures"
)

func PossibleValuesForConfigurationManagerActionType() []string {
	return []string{
		string(ConfigurationManagerActionType_AppEvaluation),
		string(ConfigurationManagerActionType_FullScan),
		string(ConfigurationManagerActionType_QuickScan),
		string(ConfigurationManagerActionType_RefreshMachinePolicy),
		string(ConfigurationManagerActionType_RefreshUserPolicy),
		string(ConfigurationManagerActionType_WakeUpClient),
		string(ConfigurationManagerActionType_WindowsDefenderUpdateSignatures),
	}
}

func (s *ConfigurationManagerActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerActionType(input string) (*ConfigurationManagerActionType, error) {
	vals := map[string]ConfigurationManagerActionType{
		"appevaluation":                   ConfigurationManagerActionType_AppEvaluation,
		"fullscan":                        ConfigurationManagerActionType_FullScan,
		"quickscan":                       ConfigurationManagerActionType_QuickScan,
		"refreshmachinepolicy":            ConfigurationManagerActionType_RefreshMachinePolicy,
		"refreshuserpolicy":               ConfigurationManagerActionType_RefreshUserPolicy,
		"wakeupclient":                    ConfigurationManagerActionType_WakeUpClient,
		"windowsdefenderupdatesignatures": ConfigurationManagerActionType_WindowsDefenderUpdateSignatures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionType(input)
	return &out, nil
}
