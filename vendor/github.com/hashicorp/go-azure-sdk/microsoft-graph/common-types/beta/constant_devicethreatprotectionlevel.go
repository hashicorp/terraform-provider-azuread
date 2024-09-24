package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceThreatProtectionLevel string

const (
	DeviceThreatProtectionLevel_High        DeviceThreatProtectionLevel = "high"
	DeviceThreatProtectionLevel_Low         DeviceThreatProtectionLevel = "low"
	DeviceThreatProtectionLevel_Medium      DeviceThreatProtectionLevel = "medium"
	DeviceThreatProtectionLevel_NotSet      DeviceThreatProtectionLevel = "notSet"
	DeviceThreatProtectionLevel_Secured     DeviceThreatProtectionLevel = "secured"
	DeviceThreatProtectionLevel_Unavailable DeviceThreatProtectionLevel = "unavailable"
)

func PossibleValuesForDeviceThreatProtectionLevel() []string {
	return []string{
		string(DeviceThreatProtectionLevel_High),
		string(DeviceThreatProtectionLevel_Low),
		string(DeviceThreatProtectionLevel_Medium),
		string(DeviceThreatProtectionLevel_NotSet),
		string(DeviceThreatProtectionLevel_Secured),
		string(DeviceThreatProtectionLevel_Unavailable),
	}
}

func (s *DeviceThreatProtectionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceThreatProtectionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceThreatProtectionLevel(input string) (*DeviceThreatProtectionLevel, error) {
	vals := map[string]DeviceThreatProtectionLevel{
		"high":        DeviceThreatProtectionLevel_High,
		"low":         DeviceThreatProtectionLevel_Low,
		"medium":      DeviceThreatProtectionLevel_Medium,
		"notset":      DeviceThreatProtectionLevel_NotSet,
		"secured":     DeviceThreatProtectionLevel_Secured,
		"unavailable": DeviceThreatProtectionLevel_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceThreatProtectionLevel(input)
	return &out, nil
}
