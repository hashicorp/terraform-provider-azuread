package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingRiskLevel string

const (
	DeviceManagementConfigurationSettingRiskLevel_High   DeviceManagementConfigurationSettingRiskLevel = "high"
	DeviceManagementConfigurationSettingRiskLevel_Low    DeviceManagementConfigurationSettingRiskLevel = "low"
	DeviceManagementConfigurationSettingRiskLevel_Medium DeviceManagementConfigurationSettingRiskLevel = "medium"
)

func PossibleValuesForDeviceManagementConfigurationSettingRiskLevel() []string {
	return []string{
		string(DeviceManagementConfigurationSettingRiskLevel_High),
		string(DeviceManagementConfigurationSettingRiskLevel_Low),
		string(DeviceManagementConfigurationSettingRiskLevel_Medium),
	}
}

func (s *DeviceManagementConfigurationSettingRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingRiskLevel(input string) (*DeviceManagementConfigurationSettingRiskLevel, error) {
	vals := map[string]DeviceManagementConfigurationSettingRiskLevel{
		"high":   DeviceManagementConfigurationSettingRiskLevel_High,
		"low":    DeviceManagementConfigurationSettingRiskLevel_Low,
		"medium": DeviceManagementConfigurationSettingRiskLevel_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingRiskLevel(input)
	return &out, nil
}
