package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppSafetyNetDeviceAttestationType string

const (
	AndroidManagedAppSafetyNetDeviceAttestationType_BasicIntegrity                       AndroidManagedAppSafetyNetDeviceAttestationType = "basicIntegrity"
	AndroidManagedAppSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification AndroidManagedAppSafetyNetDeviceAttestationType = "basicIntegrityAndDeviceCertification"
	AndroidManagedAppSafetyNetDeviceAttestationType_None                                 AndroidManagedAppSafetyNetDeviceAttestationType = "none"
)

func PossibleValuesForAndroidManagedAppSafetyNetDeviceAttestationType() []string {
	return []string{
		string(AndroidManagedAppSafetyNetDeviceAttestationType_BasicIntegrity),
		string(AndroidManagedAppSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification),
		string(AndroidManagedAppSafetyNetDeviceAttestationType_None),
	}
}

func (s *AndroidManagedAppSafetyNetDeviceAttestationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppSafetyNetDeviceAttestationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppSafetyNetDeviceAttestationType(input string) (*AndroidManagedAppSafetyNetDeviceAttestationType, error) {
	vals := map[string]AndroidManagedAppSafetyNetDeviceAttestationType{
		"basicintegrity":                       AndroidManagedAppSafetyNetDeviceAttestationType_BasicIntegrity,
		"basicintegrityanddevicecertification": AndroidManagedAppSafetyNetDeviceAttestationType_BasicIntegrityAndDeviceCertification,
		"none":                                 AndroidManagedAppSafetyNetDeviceAttestationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppSafetyNetDeviceAttestationType(input)
	return &out, nil
}
