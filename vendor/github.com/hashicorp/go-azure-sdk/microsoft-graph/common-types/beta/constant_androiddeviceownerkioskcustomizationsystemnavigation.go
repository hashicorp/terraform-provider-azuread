package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskCustomizationSystemNavigation string

const (
	AndroidDeviceOwnerKioskCustomizationSystemNavigation_HomeButtonOnly    AndroidDeviceOwnerKioskCustomizationSystemNavigation = "homeButtonOnly"
	AndroidDeviceOwnerKioskCustomizationSystemNavigation_NavigationEnabled AndroidDeviceOwnerKioskCustomizationSystemNavigation = "navigationEnabled"
	AndroidDeviceOwnerKioskCustomizationSystemNavigation_NotConfigured     AndroidDeviceOwnerKioskCustomizationSystemNavigation = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerKioskCustomizationSystemNavigation() []string {
	return []string{
		string(AndroidDeviceOwnerKioskCustomizationSystemNavigation_HomeButtonOnly),
		string(AndroidDeviceOwnerKioskCustomizationSystemNavigation_NavigationEnabled),
		string(AndroidDeviceOwnerKioskCustomizationSystemNavigation_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerKioskCustomizationSystemNavigation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerKioskCustomizationSystemNavigation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerKioskCustomizationSystemNavigation(input string) (*AndroidDeviceOwnerKioskCustomizationSystemNavigation, error) {
	vals := map[string]AndroidDeviceOwnerKioskCustomizationSystemNavigation{
		"homebuttononly":    AndroidDeviceOwnerKioskCustomizationSystemNavigation_HomeButtonOnly,
		"navigationenabled": AndroidDeviceOwnerKioskCustomizationSystemNavigation_NavigationEnabled,
		"notconfigured":     AndroidDeviceOwnerKioskCustomizationSystemNavigation_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerKioskCustomizationSystemNavigation(input)
	return &out, nil
}
