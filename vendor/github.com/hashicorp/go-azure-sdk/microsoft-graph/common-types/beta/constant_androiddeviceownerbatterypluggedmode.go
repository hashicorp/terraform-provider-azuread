package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerBatteryPluggedMode string

const (
	AndroidDeviceOwnerBatteryPluggedMode_Ac            AndroidDeviceOwnerBatteryPluggedMode = "ac"
	AndroidDeviceOwnerBatteryPluggedMode_NotConfigured AndroidDeviceOwnerBatteryPluggedMode = "notConfigured"
	AndroidDeviceOwnerBatteryPluggedMode_Usb           AndroidDeviceOwnerBatteryPluggedMode = "usb"
	AndroidDeviceOwnerBatteryPluggedMode_Wireless      AndroidDeviceOwnerBatteryPluggedMode = "wireless"
)

func PossibleValuesForAndroidDeviceOwnerBatteryPluggedMode() []string {
	return []string{
		string(AndroidDeviceOwnerBatteryPluggedMode_Ac),
		string(AndroidDeviceOwnerBatteryPluggedMode_NotConfigured),
		string(AndroidDeviceOwnerBatteryPluggedMode_Usb),
		string(AndroidDeviceOwnerBatteryPluggedMode_Wireless),
	}
}

func (s *AndroidDeviceOwnerBatteryPluggedMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerBatteryPluggedMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerBatteryPluggedMode(input string) (*AndroidDeviceOwnerBatteryPluggedMode, error) {
	vals := map[string]AndroidDeviceOwnerBatteryPluggedMode{
		"ac":            AndroidDeviceOwnerBatteryPluggedMode_Ac,
		"notconfigured": AndroidDeviceOwnerBatteryPluggedMode_NotConfigured,
		"usb":           AndroidDeviceOwnerBatteryPluggedMode_Usb,
		"wireless":      AndroidDeviceOwnerBatteryPluggedMode_Wireless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerBatteryPluggedMode(input)
	return &out, nil
}
