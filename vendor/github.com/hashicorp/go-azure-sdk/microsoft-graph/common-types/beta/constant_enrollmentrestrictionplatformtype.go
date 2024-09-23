package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentRestrictionPlatformType string

const (
	EnrollmentRestrictionPlatformType_AllPlatforms   EnrollmentRestrictionPlatformType = "allPlatforms"
	EnrollmentRestrictionPlatformType_Android        EnrollmentRestrictionPlatformType = "android"
	EnrollmentRestrictionPlatformType_AndroidForWork EnrollmentRestrictionPlatformType = "androidForWork"
	EnrollmentRestrictionPlatformType_Ios            EnrollmentRestrictionPlatformType = "ios"
	EnrollmentRestrictionPlatformType_Linux          EnrollmentRestrictionPlatformType = "linux"
	EnrollmentRestrictionPlatformType_Mac            EnrollmentRestrictionPlatformType = "mac"
	EnrollmentRestrictionPlatformType_Windows        EnrollmentRestrictionPlatformType = "windows"
	EnrollmentRestrictionPlatformType_WindowsPhone   EnrollmentRestrictionPlatformType = "windowsPhone"
)

func PossibleValuesForEnrollmentRestrictionPlatformType() []string {
	return []string{
		string(EnrollmentRestrictionPlatformType_AllPlatforms),
		string(EnrollmentRestrictionPlatformType_Android),
		string(EnrollmentRestrictionPlatformType_AndroidForWork),
		string(EnrollmentRestrictionPlatformType_Ios),
		string(EnrollmentRestrictionPlatformType_Linux),
		string(EnrollmentRestrictionPlatformType_Mac),
		string(EnrollmentRestrictionPlatformType_Windows),
		string(EnrollmentRestrictionPlatformType_WindowsPhone),
	}
}

func (s *EnrollmentRestrictionPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentRestrictionPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentRestrictionPlatformType(input string) (*EnrollmentRestrictionPlatformType, error) {
	vals := map[string]EnrollmentRestrictionPlatformType{
		"allplatforms":   EnrollmentRestrictionPlatformType_AllPlatforms,
		"android":        EnrollmentRestrictionPlatformType_Android,
		"androidforwork": EnrollmentRestrictionPlatformType_AndroidForWork,
		"ios":            EnrollmentRestrictionPlatformType_Ios,
		"linux":          EnrollmentRestrictionPlatformType_Linux,
		"mac":            EnrollmentRestrictionPlatformType_Mac,
		"windows":        EnrollmentRestrictionPlatformType_Windows,
		"windowsphone":   EnrollmentRestrictionPlatformType_WindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentRestrictionPlatformType(input)
	return &out, nil
}
