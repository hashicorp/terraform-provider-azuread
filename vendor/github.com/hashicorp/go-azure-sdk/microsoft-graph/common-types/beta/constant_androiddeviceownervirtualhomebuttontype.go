package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerVirtualHomeButtonType string

const (
	AndroidDeviceOwnerVirtualHomeButtonType_Floating      AndroidDeviceOwnerVirtualHomeButtonType = "floating"
	AndroidDeviceOwnerVirtualHomeButtonType_NotConfigured AndroidDeviceOwnerVirtualHomeButtonType = "notConfigured"
	AndroidDeviceOwnerVirtualHomeButtonType_SwipeUp       AndroidDeviceOwnerVirtualHomeButtonType = "swipeUp"
)

func PossibleValuesForAndroidDeviceOwnerVirtualHomeButtonType() []string {
	return []string{
		string(AndroidDeviceOwnerVirtualHomeButtonType_Floating),
		string(AndroidDeviceOwnerVirtualHomeButtonType_NotConfigured),
		string(AndroidDeviceOwnerVirtualHomeButtonType_SwipeUp),
	}
}

func (s *AndroidDeviceOwnerVirtualHomeButtonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerVirtualHomeButtonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerVirtualHomeButtonType(input string) (*AndroidDeviceOwnerVirtualHomeButtonType, error) {
	vals := map[string]AndroidDeviceOwnerVirtualHomeButtonType{
		"floating":      AndroidDeviceOwnerVirtualHomeButtonType_Floating,
		"notconfigured": AndroidDeviceOwnerVirtualHomeButtonType_NotConfigured,
		"swipeup":       AndroidDeviceOwnerVirtualHomeButtonType_SwipeUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerVirtualHomeButtonType(input)
	return &out, nil
}
