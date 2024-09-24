package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerLocationMode string

const (
	AndroidDeviceOwnerLocationMode_Disabled      AndroidDeviceOwnerLocationMode = "disabled"
	AndroidDeviceOwnerLocationMode_NotConfigured AndroidDeviceOwnerLocationMode = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerLocationMode() []string {
	return []string{
		string(AndroidDeviceOwnerLocationMode_Disabled),
		string(AndroidDeviceOwnerLocationMode_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerLocationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerLocationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerLocationMode(input string) (*AndroidDeviceOwnerLocationMode, error) {
	vals := map[string]AndroidDeviceOwnerLocationMode{
		"disabled":      AndroidDeviceOwnerLocationMode_Disabled,
		"notconfigured": AndroidDeviceOwnerLocationMode_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerLocationMode(input)
	return &out, nil
}
