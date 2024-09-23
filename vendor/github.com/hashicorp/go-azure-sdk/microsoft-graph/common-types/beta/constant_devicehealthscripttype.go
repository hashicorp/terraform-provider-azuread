package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptType string

const (
	DeviceHealthScriptType_DeviceHealthScript     DeviceHealthScriptType = "deviceHealthScript"
	DeviceHealthScriptType_ManagedInstallerScript DeviceHealthScriptType = "managedInstallerScript"
)

func PossibleValuesForDeviceHealthScriptType() []string {
	return []string{
		string(DeviceHealthScriptType_DeviceHealthScript),
		string(DeviceHealthScriptType_ManagedInstallerScript),
	}
}

func (s *DeviceHealthScriptType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthScriptType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthScriptType(input string) (*DeviceHealthScriptType, error) {
	vals := map[string]DeviceHealthScriptType{
		"devicehealthscript":     DeviceHealthScriptType_DeviceHealthScript,
		"managedinstallerscript": DeviceHealthScriptType_ManagedInstallerScript,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptType(input)
	return &out, nil
}
