package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerSystemUpdateInstallType string

const (
	AndroidDeviceOwnerSystemUpdateInstallType_Automatic     AndroidDeviceOwnerSystemUpdateInstallType = "automatic"
	AndroidDeviceOwnerSystemUpdateInstallType_DeviceDefault AndroidDeviceOwnerSystemUpdateInstallType = "deviceDefault"
	AndroidDeviceOwnerSystemUpdateInstallType_Postpone      AndroidDeviceOwnerSystemUpdateInstallType = "postpone"
	AndroidDeviceOwnerSystemUpdateInstallType_Windowed      AndroidDeviceOwnerSystemUpdateInstallType = "windowed"
)

func PossibleValuesForAndroidDeviceOwnerSystemUpdateInstallType() []string {
	return []string{
		string(AndroidDeviceOwnerSystemUpdateInstallType_Automatic),
		string(AndroidDeviceOwnerSystemUpdateInstallType_DeviceDefault),
		string(AndroidDeviceOwnerSystemUpdateInstallType_Postpone),
		string(AndroidDeviceOwnerSystemUpdateInstallType_Windowed),
	}
}

func (s *AndroidDeviceOwnerSystemUpdateInstallType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerSystemUpdateInstallType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerSystemUpdateInstallType(input string) (*AndroidDeviceOwnerSystemUpdateInstallType, error) {
	vals := map[string]AndroidDeviceOwnerSystemUpdateInstallType{
		"automatic":     AndroidDeviceOwnerSystemUpdateInstallType_Automatic,
		"devicedefault": AndroidDeviceOwnerSystemUpdateInstallType_DeviceDefault,
		"postpone":      AndroidDeviceOwnerSystemUpdateInstallType_Postpone,
		"windowed":      AndroidDeviceOwnerSystemUpdateInstallType_Windowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerSystemUpdateInstallType(input)
	return &out, nil
}
