package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentTokenType string

const (
	AndroidDeviceOwnerEnrollmentTokenType_CorporateOwnedDedicatedDeviceWithAzureADSharedMode AndroidDeviceOwnerEnrollmentTokenType = "corporateOwnedDedicatedDeviceWithAzureADSharedMode"
	AndroidDeviceOwnerEnrollmentTokenType_Default                                            AndroidDeviceOwnerEnrollmentTokenType = "default"
	AndroidDeviceOwnerEnrollmentTokenType_DeviceStaging                                      AndroidDeviceOwnerEnrollmentTokenType = "deviceStaging"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentTokenType() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentTokenType_CorporateOwnedDedicatedDeviceWithAzureADSharedMode),
		string(AndroidDeviceOwnerEnrollmentTokenType_Default),
		string(AndroidDeviceOwnerEnrollmentTokenType_DeviceStaging),
	}
}

func (s *AndroidDeviceOwnerEnrollmentTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnrollmentTokenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnrollmentTokenType(input string) (*AndroidDeviceOwnerEnrollmentTokenType, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentTokenType{
		"corporateowneddedicateddevicewithazureadsharedmode": AndroidDeviceOwnerEnrollmentTokenType_CorporateOwnedDedicatedDeviceWithAzureADSharedMode,
		"default":       AndroidDeviceOwnerEnrollmentTokenType_Default,
		"devicestaging": AndroidDeviceOwnerEnrollmentTokenType_DeviceStaging,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentTokenType(input)
	return &out, nil
}
