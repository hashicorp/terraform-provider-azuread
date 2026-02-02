package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerKioskModeFolderIcon string

const (
	AndroidDeviceOwnerKioskModeFolderIcon_DarkCircle    AndroidDeviceOwnerKioskModeFolderIcon = "darkCircle"
	AndroidDeviceOwnerKioskModeFolderIcon_DarkSquare    AndroidDeviceOwnerKioskModeFolderIcon = "darkSquare"
	AndroidDeviceOwnerKioskModeFolderIcon_LightCircle   AndroidDeviceOwnerKioskModeFolderIcon = "lightCircle"
	AndroidDeviceOwnerKioskModeFolderIcon_LightSquare   AndroidDeviceOwnerKioskModeFolderIcon = "lightSquare"
	AndroidDeviceOwnerKioskModeFolderIcon_NotConfigured AndroidDeviceOwnerKioskModeFolderIcon = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerKioskModeFolderIcon() []string {
	return []string{
		string(AndroidDeviceOwnerKioskModeFolderIcon_DarkCircle),
		string(AndroidDeviceOwnerKioskModeFolderIcon_DarkSquare),
		string(AndroidDeviceOwnerKioskModeFolderIcon_LightCircle),
		string(AndroidDeviceOwnerKioskModeFolderIcon_LightSquare),
		string(AndroidDeviceOwnerKioskModeFolderIcon_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerKioskModeFolderIcon) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerKioskModeFolderIcon(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerKioskModeFolderIcon(input string) (*AndroidDeviceOwnerKioskModeFolderIcon, error) {
	vals := map[string]AndroidDeviceOwnerKioskModeFolderIcon{
		"darkcircle":    AndroidDeviceOwnerKioskModeFolderIcon_DarkCircle,
		"darksquare":    AndroidDeviceOwnerKioskModeFolderIcon_DarkSquare,
		"lightcircle":   AndroidDeviceOwnerKioskModeFolderIcon_LightCircle,
		"lightsquare":   AndroidDeviceOwnerKioskModeFolderIcon_LightSquare,
		"notconfigured": AndroidDeviceOwnerKioskModeFolderIcon_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerKioskModeFolderIcon(input)
	return &out, nil
}
