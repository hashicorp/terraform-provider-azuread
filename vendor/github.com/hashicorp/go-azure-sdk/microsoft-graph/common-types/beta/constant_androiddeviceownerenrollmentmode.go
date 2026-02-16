package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentMode string

const (
	AndroidDeviceOwnerEnrollmentMode_CorporateOwnedAOSPUserAssociatedDevice AndroidDeviceOwnerEnrollmentMode = "corporateOwnedAOSPUserAssociatedDevice"
	AndroidDeviceOwnerEnrollmentMode_CorporateOwnedAOSPUserlessDevice       AndroidDeviceOwnerEnrollmentMode = "corporateOwnedAOSPUserlessDevice"
	AndroidDeviceOwnerEnrollmentMode_CorporateOwnedDedicatedDevice          AndroidDeviceOwnerEnrollmentMode = "corporateOwnedDedicatedDevice"
	AndroidDeviceOwnerEnrollmentMode_CorporateOwnedFullyManaged             AndroidDeviceOwnerEnrollmentMode = "corporateOwnedFullyManaged"
	AndroidDeviceOwnerEnrollmentMode_CorporateOwnedWorkProfile              AndroidDeviceOwnerEnrollmentMode = "corporateOwnedWorkProfile"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentMode() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentMode_CorporateOwnedAOSPUserAssociatedDevice),
		string(AndroidDeviceOwnerEnrollmentMode_CorporateOwnedAOSPUserlessDevice),
		string(AndroidDeviceOwnerEnrollmentMode_CorporateOwnedDedicatedDevice),
		string(AndroidDeviceOwnerEnrollmentMode_CorporateOwnedFullyManaged),
		string(AndroidDeviceOwnerEnrollmentMode_CorporateOwnedWorkProfile),
	}
}

func (s *AndroidDeviceOwnerEnrollmentMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnrollmentMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnrollmentMode(input string) (*AndroidDeviceOwnerEnrollmentMode, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentMode{
		"corporateownedaospuserassociateddevice": AndroidDeviceOwnerEnrollmentMode_CorporateOwnedAOSPUserAssociatedDevice,
		"corporateownedaospuserlessdevice":       AndroidDeviceOwnerEnrollmentMode_CorporateOwnedAOSPUserlessDevice,
		"corporateowneddedicateddevice":          AndroidDeviceOwnerEnrollmentMode_CorporateOwnedDedicatedDevice,
		"corporateownedfullymanaged":             AndroidDeviceOwnerEnrollmentMode_CorporateOwnedFullyManaged,
		"corporateownedworkprofile":              AndroidDeviceOwnerEnrollmentMode_CorporateOwnedWorkProfile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentMode(input)
	return &out, nil
}
