package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderSecurityCenterITContactDisplayType string

const (
	DefenderSecurityCenterITContactDisplayType_DisplayInAppAndInNotifications DefenderSecurityCenterITContactDisplayType = "displayInAppAndInNotifications"
	DefenderSecurityCenterITContactDisplayType_DisplayOnlyInApp               DefenderSecurityCenterITContactDisplayType = "displayOnlyInApp"
	DefenderSecurityCenterITContactDisplayType_DisplayOnlyInNotifications     DefenderSecurityCenterITContactDisplayType = "displayOnlyInNotifications"
	DefenderSecurityCenterITContactDisplayType_NotConfigured                  DefenderSecurityCenterITContactDisplayType = "notConfigured"
)

func PossibleValuesForDefenderSecurityCenterITContactDisplayType() []string {
	return []string{
		string(DefenderSecurityCenterITContactDisplayType_DisplayInAppAndInNotifications),
		string(DefenderSecurityCenterITContactDisplayType_DisplayOnlyInApp),
		string(DefenderSecurityCenterITContactDisplayType_DisplayOnlyInNotifications),
		string(DefenderSecurityCenterITContactDisplayType_NotConfigured),
	}
}

func (s *DefenderSecurityCenterITContactDisplayType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderSecurityCenterITContactDisplayType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderSecurityCenterITContactDisplayType(input string) (*DefenderSecurityCenterITContactDisplayType, error) {
	vals := map[string]DefenderSecurityCenterITContactDisplayType{
		"displayinappandinnotifications": DefenderSecurityCenterITContactDisplayType_DisplayInAppAndInNotifications,
		"displayonlyinapp":               DefenderSecurityCenterITContactDisplayType_DisplayOnlyInApp,
		"displayonlyinnotifications":     DefenderSecurityCenterITContactDisplayType_DisplayOnlyInNotifications,
		"notconfigured":                  DefenderSecurityCenterITContactDisplayType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderSecurityCenterITContactDisplayType(input)
	return &out, nil
}
