package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeScreenOrientation string

const (
	AndroidDeviceOwnerKioskModeScreenOrientation_AutoRotate    AndroidDeviceOwnerKioskModeScreenOrientation = "autoRotate"
	AndroidDeviceOwnerKioskModeScreenOrientation_Landscape     AndroidDeviceOwnerKioskModeScreenOrientation = "landscape"
	AndroidDeviceOwnerKioskModeScreenOrientation_NotConfigured AndroidDeviceOwnerKioskModeScreenOrientation = "notConfigured"
	AndroidDeviceOwnerKioskModeScreenOrientation_Portrait      AndroidDeviceOwnerKioskModeScreenOrientation = "portrait"
)

func PossibleValuesForAndroidDeviceOwnerKioskModeScreenOrientation() []string {
	return []string{
		string(AndroidDeviceOwnerKioskModeScreenOrientation_AutoRotate),
		string(AndroidDeviceOwnerKioskModeScreenOrientation_Landscape),
		string(AndroidDeviceOwnerKioskModeScreenOrientation_NotConfigured),
		string(AndroidDeviceOwnerKioskModeScreenOrientation_Portrait),
	}
}

func (s *AndroidDeviceOwnerKioskModeScreenOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerKioskModeScreenOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerKioskModeScreenOrientation(input string) (*AndroidDeviceOwnerKioskModeScreenOrientation, error) {
	vals := map[string]AndroidDeviceOwnerKioskModeScreenOrientation{
		"autorotate":    AndroidDeviceOwnerKioskModeScreenOrientation_AutoRotate,
		"landscape":     AndroidDeviceOwnerKioskModeScreenOrientation_Landscape,
		"notconfigured": AndroidDeviceOwnerKioskModeScreenOrientation_NotConfigured,
		"portrait":      AndroidDeviceOwnerKioskModeScreenOrientation_Portrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerKioskModeScreenOrientation(input)
	return &out, nil
}
