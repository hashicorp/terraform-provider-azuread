package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskCustomizationStatusBar string

const (
	AndroidDeviceOwnerKioskCustomizationStatusBar_NotConfigured                     AndroidDeviceOwnerKioskCustomizationStatusBar = "notConfigured"
	AndroidDeviceOwnerKioskCustomizationStatusBar_NotificationsAndSystemInfoEnabled AndroidDeviceOwnerKioskCustomizationStatusBar = "notificationsAndSystemInfoEnabled"
	AndroidDeviceOwnerKioskCustomizationStatusBar_SystemInfoOnly                    AndroidDeviceOwnerKioskCustomizationStatusBar = "systemInfoOnly"
)

func PossibleValuesForAndroidDeviceOwnerKioskCustomizationStatusBar() []string {
	return []string{
		string(AndroidDeviceOwnerKioskCustomizationStatusBar_NotConfigured),
		string(AndroidDeviceOwnerKioskCustomizationStatusBar_NotificationsAndSystemInfoEnabled),
		string(AndroidDeviceOwnerKioskCustomizationStatusBar_SystemInfoOnly),
	}
}

func (s *AndroidDeviceOwnerKioskCustomizationStatusBar) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerKioskCustomizationStatusBar(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerKioskCustomizationStatusBar(input string) (*AndroidDeviceOwnerKioskCustomizationStatusBar, error) {
	vals := map[string]AndroidDeviceOwnerKioskCustomizationStatusBar{
		"notconfigured":                     AndroidDeviceOwnerKioskCustomizationStatusBar_NotConfigured,
		"notificationsandsysteminfoenabled": AndroidDeviceOwnerKioskCustomizationStatusBar_NotificationsAndSystemInfoEnabled,
		"systeminfoonly":                    AndroidDeviceOwnerKioskCustomizationStatusBar_SystemInfoOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerKioskCustomizationStatusBar(input)
	return &out, nil
}
