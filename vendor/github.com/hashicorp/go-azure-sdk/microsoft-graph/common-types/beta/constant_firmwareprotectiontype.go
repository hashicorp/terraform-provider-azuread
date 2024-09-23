package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirmwareProtectionType string

const (
	FirmwareProtectionType_Disabled                       FirmwareProtectionType = "disabled"
	FirmwareProtectionType_FirmwareAttackSurfaceReduction FirmwareProtectionType = "firmwareAttackSurfaceReduction"
	FirmwareProtectionType_NotApplicable                  FirmwareProtectionType = "notApplicable"
	FirmwareProtectionType_SystemGuardSecureLaunch        FirmwareProtectionType = "systemGuardSecureLaunch"
)

func PossibleValuesForFirmwareProtectionType() []string {
	return []string{
		string(FirmwareProtectionType_Disabled),
		string(FirmwareProtectionType_FirmwareAttackSurfaceReduction),
		string(FirmwareProtectionType_NotApplicable),
		string(FirmwareProtectionType_SystemGuardSecureLaunch),
	}
}

func (s *FirmwareProtectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirmwareProtectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirmwareProtectionType(input string) (*FirmwareProtectionType, error) {
	vals := map[string]FirmwareProtectionType{
		"disabled":                       FirmwareProtectionType_Disabled,
		"firmwareattacksurfacereduction": FirmwareProtectionType_FirmwareAttackSurfaceReduction,
		"notapplicable":                  FirmwareProtectionType_NotApplicable,
		"systemguardsecurelaunch":        FirmwareProtectionType_SystemGuardSecureLaunch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirmwareProtectionType(input)
	return &out, nil
}
