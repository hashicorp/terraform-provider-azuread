package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeIconSize string

const (
	AndroidDeviceOwnerKioskModeIconSize_Large         AndroidDeviceOwnerKioskModeIconSize = "large"
	AndroidDeviceOwnerKioskModeIconSize_Largest       AndroidDeviceOwnerKioskModeIconSize = "largest"
	AndroidDeviceOwnerKioskModeIconSize_NotConfigured AndroidDeviceOwnerKioskModeIconSize = "notConfigured"
	AndroidDeviceOwnerKioskModeIconSize_Regular       AndroidDeviceOwnerKioskModeIconSize = "regular"
	AndroidDeviceOwnerKioskModeIconSize_Small         AndroidDeviceOwnerKioskModeIconSize = "small"
	AndroidDeviceOwnerKioskModeIconSize_Smallest      AndroidDeviceOwnerKioskModeIconSize = "smallest"
)

func PossibleValuesForAndroidDeviceOwnerKioskModeIconSize() []string {
	return []string{
		string(AndroidDeviceOwnerKioskModeIconSize_Large),
		string(AndroidDeviceOwnerKioskModeIconSize_Largest),
		string(AndroidDeviceOwnerKioskModeIconSize_NotConfigured),
		string(AndroidDeviceOwnerKioskModeIconSize_Regular),
		string(AndroidDeviceOwnerKioskModeIconSize_Small),
		string(AndroidDeviceOwnerKioskModeIconSize_Smallest),
	}
}

func (s *AndroidDeviceOwnerKioskModeIconSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerKioskModeIconSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerKioskModeIconSize(input string) (*AndroidDeviceOwnerKioskModeIconSize, error) {
	vals := map[string]AndroidDeviceOwnerKioskModeIconSize{
		"large":         AndroidDeviceOwnerKioskModeIconSize_Large,
		"largest":       AndroidDeviceOwnerKioskModeIconSize_Largest,
		"notconfigured": AndroidDeviceOwnerKioskModeIconSize_NotConfigured,
		"regular":       AndroidDeviceOwnerKioskModeIconSize_Regular,
		"small":         AndroidDeviceOwnerKioskModeIconSize_Small,
		"smallest":      AndroidDeviceOwnerKioskModeIconSize_Smallest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerKioskModeIconSize(input)
	return &out, nil
}
