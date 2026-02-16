package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoPciDssVersion string

const (
	SecurityAppInfoPciDssVersion_NotSupported SecurityAppInfoPciDssVersion = "notSupported"
	SecurityAppInfoPciDssVersion_Unknown      SecurityAppInfoPciDssVersion = "unknown"
	SecurityAppInfoPciDssVersion_V1           SecurityAppInfoPciDssVersion = "v1"
	SecurityAppInfoPciDssVersion_V2           SecurityAppInfoPciDssVersion = "v2"
	SecurityAppInfoPciDssVersion_V3           SecurityAppInfoPciDssVersion = "v3"
	SecurityAppInfoPciDssVersion_V31          SecurityAppInfoPciDssVersion = "v3_1"
	SecurityAppInfoPciDssVersion_V32          SecurityAppInfoPciDssVersion = "v3_2"
	SecurityAppInfoPciDssVersion_V321         SecurityAppInfoPciDssVersion = "v3_2_1"
	SecurityAppInfoPciDssVersion_V4           SecurityAppInfoPciDssVersion = "v4"
)

func PossibleValuesForSecurityAppInfoPciDssVersion() []string {
	return []string{
		string(SecurityAppInfoPciDssVersion_NotSupported),
		string(SecurityAppInfoPciDssVersion_Unknown),
		string(SecurityAppInfoPciDssVersion_V1),
		string(SecurityAppInfoPciDssVersion_V2),
		string(SecurityAppInfoPciDssVersion_V3),
		string(SecurityAppInfoPciDssVersion_V31),
		string(SecurityAppInfoPciDssVersion_V32),
		string(SecurityAppInfoPciDssVersion_V321),
		string(SecurityAppInfoPciDssVersion_V4),
	}
}

func (s *SecurityAppInfoPciDssVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoPciDssVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoPciDssVersion(input string) (*SecurityAppInfoPciDssVersion, error) {
	vals := map[string]SecurityAppInfoPciDssVersion{
		"notsupported": SecurityAppInfoPciDssVersion_NotSupported,
		"unknown":      SecurityAppInfoPciDssVersion_Unknown,
		"v1":           SecurityAppInfoPciDssVersion_V1,
		"v2":           SecurityAppInfoPciDssVersion_V2,
		"v3":           SecurityAppInfoPciDssVersion_V3,
		"v3_1":         SecurityAppInfoPciDssVersion_V31,
		"v3_2":         SecurityAppInfoPciDssVersion_V32,
		"v3_2_1":       SecurityAppInfoPciDssVersion_V321,
		"v4":           SecurityAppInfoPciDssVersion_V4,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoPciDssVersion(input)
	return &out, nil
}
