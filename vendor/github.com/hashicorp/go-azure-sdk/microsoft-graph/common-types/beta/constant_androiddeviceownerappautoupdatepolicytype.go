package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerAppAutoUpdatePolicyType string

const (
	AndroidDeviceOwnerAppAutoUpdatePolicyType_Always        AndroidDeviceOwnerAppAutoUpdatePolicyType = "always"
	AndroidDeviceOwnerAppAutoUpdatePolicyType_Never         AndroidDeviceOwnerAppAutoUpdatePolicyType = "never"
	AndroidDeviceOwnerAppAutoUpdatePolicyType_NotConfigured AndroidDeviceOwnerAppAutoUpdatePolicyType = "notConfigured"
	AndroidDeviceOwnerAppAutoUpdatePolicyType_UserChoice    AndroidDeviceOwnerAppAutoUpdatePolicyType = "userChoice"
	AndroidDeviceOwnerAppAutoUpdatePolicyType_WiFiOnly      AndroidDeviceOwnerAppAutoUpdatePolicyType = "wiFiOnly"
)

func PossibleValuesForAndroidDeviceOwnerAppAutoUpdatePolicyType() []string {
	return []string{
		string(AndroidDeviceOwnerAppAutoUpdatePolicyType_Always),
		string(AndroidDeviceOwnerAppAutoUpdatePolicyType_Never),
		string(AndroidDeviceOwnerAppAutoUpdatePolicyType_NotConfigured),
		string(AndroidDeviceOwnerAppAutoUpdatePolicyType_UserChoice),
		string(AndroidDeviceOwnerAppAutoUpdatePolicyType_WiFiOnly),
	}
}

func (s *AndroidDeviceOwnerAppAutoUpdatePolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerAppAutoUpdatePolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerAppAutoUpdatePolicyType(input string) (*AndroidDeviceOwnerAppAutoUpdatePolicyType, error) {
	vals := map[string]AndroidDeviceOwnerAppAutoUpdatePolicyType{
		"always":        AndroidDeviceOwnerAppAutoUpdatePolicyType_Always,
		"never":         AndroidDeviceOwnerAppAutoUpdatePolicyType_Never,
		"notconfigured": AndroidDeviceOwnerAppAutoUpdatePolicyType_NotConfigured,
		"userchoice":    AndroidDeviceOwnerAppAutoUpdatePolicyType_UserChoice,
		"wifionly":      AndroidDeviceOwnerAppAutoUpdatePolicyType_WiFiOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerAppAutoUpdatePolicyType(input)
	return &out, nil
}
