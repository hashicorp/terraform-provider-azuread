package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyPlatformType string

const (
	PolicyPlatformType_All                PolicyPlatformType = "all"
	PolicyPlatformType_Android            PolicyPlatformType = "android"
	PolicyPlatformType_AndroidAOSP        PolicyPlatformType = "androidAOSP"
	PolicyPlatformType_AndroidForWork     PolicyPlatformType = "androidForWork"
	PolicyPlatformType_AndroidWorkProfile PolicyPlatformType = "androidWorkProfile"
	PolicyPlatformType_IOS                PolicyPlatformType = "iOS"
	PolicyPlatformType_MacOS              PolicyPlatformType = "macOS"
	PolicyPlatformType_Windows10AndLater  PolicyPlatformType = "windows10AndLater"
	PolicyPlatformType_Windows10XProfile  PolicyPlatformType = "windows10XProfile"
	PolicyPlatformType_Windows81AndLater  PolicyPlatformType = "windows81AndLater"
	PolicyPlatformType_WindowsPhone81     PolicyPlatformType = "windowsPhone81"
)

func PossibleValuesForPolicyPlatformType() []string {
	return []string{
		string(PolicyPlatformType_All),
		string(PolicyPlatformType_Android),
		string(PolicyPlatformType_AndroidAOSP),
		string(PolicyPlatformType_AndroidForWork),
		string(PolicyPlatformType_AndroidWorkProfile),
		string(PolicyPlatformType_IOS),
		string(PolicyPlatformType_MacOS),
		string(PolicyPlatformType_Windows10AndLater),
		string(PolicyPlatformType_Windows10XProfile),
		string(PolicyPlatformType_Windows81AndLater),
		string(PolicyPlatformType_WindowsPhone81),
	}
}

func (s *PolicyPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyPlatformType(input string) (*PolicyPlatformType, error) {
	vals := map[string]PolicyPlatformType{
		"all":                PolicyPlatformType_All,
		"android":            PolicyPlatformType_Android,
		"androidaosp":        PolicyPlatformType_AndroidAOSP,
		"androidforwork":     PolicyPlatformType_AndroidForWork,
		"androidworkprofile": PolicyPlatformType_AndroidWorkProfile,
		"ios":                PolicyPlatformType_IOS,
		"macos":              PolicyPlatformType_MacOS,
		"windows10andlater":  PolicyPlatformType_Windows10AndLater,
		"windows10xprofile":  PolicyPlatformType_Windows10XProfile,
		"windows81andlater":  PolicyPlatformType_Windows81AndLater,
		"windowsphone81":     PolicyPlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyPlatformType(input)
	return &out, nil
}
