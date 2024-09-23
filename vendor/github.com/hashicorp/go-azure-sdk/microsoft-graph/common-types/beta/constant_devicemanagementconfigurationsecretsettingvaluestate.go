package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSecretSettingValueState string

const (
	DeviceManagementConfigurationSecretSettingValueState_EncryptedValueToken DeviceManagementConfigurationSecretSettingValueState = "encryptedValueToken"
	DeviceManagementConfigurationSecretSettingValueState_Invalid             DeviceManagementConfigurationSecretSettingValueState = "invalid"
	DeviceManagementConfigurationSecretSettingValueState_NotEncrypted        DeviceManagementConfigurationSecretSettingValueState = "notEncrypted"
)

func PossibleValuesForDeviceManagementConfigurationSecretSettingValueState() []string {
	return []string{
		string(DeviceManagementConfigurationSecretSettingValueState_EncryptedValueToken),
		string(DeviceManagementConfigurationSecretSettingValueState_Invalid),
		string(DeviceManagementConfigurationSecretSettingValueState_NotEncrypted),
	}
}

func (s *DeviceManagementConfigurationSecretSettingValueState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSecretSettingValueState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSecretSettingValueState(input string) (*DeviceManagementConfigurationSecretSettingValueState, error) {
	vals := map[string]DeviceManagementConfigurationSecretSettingValueState{
		"encryptedvaluetoken": DeviceManagementConfigurationSecretSettingValueState_EncryptedValueToken,
		"invalid":             DeviceManagementConfigurationSecretSettingValueState_Invalid,
		"notencrypted":        DeviceManagementConfigurationSecretSettingValueState_NotEncrypted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSecretSettingValueState(input)
	return &out, nil
}
