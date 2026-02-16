package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileType string

const (
	AndroidDeviceOwnerEnrollmentProfileType_DedicatedDevice AndroidDeviceOwnerEnrollmentProfileType = "dedicatedDevice"
	AndroidDeviceOwnerEnrollmentProfileType_FullyManaged    AndroidDeviceOwnerEnrollmentProfileType = "fullyManaged"
	AndroidDeviceOwnerEnrollmentProfileType_NotConfigured   AndroidDeviceOwnerEnrollmentProfileType = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileType() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileType_DedicatedDevice),
		string(AndroidDeviceOwnerEnrollmentProfileType_FullyManaged),
		string(AndroidDeviceOwnerEnrollmentProfileType_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerEnrollmentProfileType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnrollmentProfileType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnrollmentProfileType(input string) (*AndroidDeviceOwnerEnrollmentProfileType, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileType{
		"dedicateddevice": AndroidDeviceOwnerEnrollmentProfileType_DedicatedDevice,
		"fullymanaged":    AndroidDeviceOwnerEnrollmentProfileType_FullyManaged,
		"notconfigured":   AndroidDeviceOwnerEnrollmentProfileType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileType(input)
	return &out, nil
}
