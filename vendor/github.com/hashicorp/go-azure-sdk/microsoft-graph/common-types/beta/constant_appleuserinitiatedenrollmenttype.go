package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentType string

const (
	AppleUserInitiatedEnrollmentType_AccountDrivenUserEnrollment AppleUserInitiatedEnrollmentType = "accountDrivenUserEnrollment"
	AppleUserInitiatedEnrollmentType_Device                      AppleUserInitiatedEnrollmentType = "device"
	AppleUserInitiatedEnrollmentType_Unknown                     AppleUserInitiatedEnrollmentType = "unknown"
	AppleUserInitiatedEnrollmentType_User                        AppleUserInitiatedEnrollmentType = "user"
	AppleUserInitiatedEnrollmentType_WebDeviceEnrollment         AppleUserInitiatedEnrollmentType = "webDeviceEnrollment"
)

func PossibleValuesForAppleUserInitiatedEnrollmentType() []string {
	return []string{
		string(AppleUserInitiatedEnrollmentType_AccountDrivenUserEnrollment),
		string(AppleUserInitiatedEnrollmentType_Device),
		string(AppleUserInitiatedEnrollmentType_Unknown),
		string(AppleUserInitiatedEnrollmentType_User),
		string(AppleUserInitiatedEnrollmentType_WebDeviceEnrollment),
	}
}

func (s *AppleUserInitiatedEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleUserInitiatedEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleUserInitiatedEnrollmentType(input string) (*AppleUserInitiatedEnrollmentType, error) {
	vals := map[string]AppleUserInitiatedEnrollmentType{
		"accountdrivenuserenrollment": AppleUserInitiatedEnrollmentType_AccountDrivenUserEnrollment,
		"device":                      AppleUserInitiatedEnrollmentType_Device,
		"unknown":                     AppleUserInitiatedEnrollmentType_Unknown,
		"user":                        AppleUserInitiatedEnrollmentType_User,
		"webdeviceenrollment":         AppleUserInitiatedEnrollmentType_WebDeviceEnrollment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleUserInitiatedEnrollmentType(input)
	return &out, nil
}
