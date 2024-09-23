package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KioskModeType string

const (
	KioskModeType_MultiAppMode  KioskModeType = "multiAppMode"
	KioskModeType_NotConfigured KioskModeType = "notConfigured"
	KioskModeType_SingleAppMode KioskModeType = "singleAppMode"
)

func PossibleValuesForKioskModeType() []string {
	return []string{
		string(KioskModeType_MultiAppMode),
		string(KioskModeType_NotConfigured),
		string(KioskModeType_SingleAppMode),
	}
}

func (s *KioskModeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKioskModeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKioskModeType(input string) (*KioskModeType, error) {
	vals := map[string]KioskModeType{
		"multiappmode":  KioskModeType_MultiAppMode,
		"notconfigured": KioskModeType_NotConfigured,
		"singleappmode": KioskModeType_SingleAppMode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KioskModeType(input)
	return &out, nil
}
