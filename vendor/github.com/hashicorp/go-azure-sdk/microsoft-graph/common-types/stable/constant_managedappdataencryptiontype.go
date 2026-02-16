package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDataEncryptionType string

const (
	ManagedAppDataEncryptionType_AfterDeviceRestart              ManagedAppDataEncryptionType = "afterDeviceRestart"
	ManagedAppDataEncryptionType_UseDeviceSettings               ManagedAppDataEncryptionType = "useDeviceSettings"
	ManagedAppDataEncryptionType_WhenDeviceLocked                ManagedAppDataEncryptionType = "whenDeviceLocked"
	ManagedAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles ManagedAppDataEncryptionType = "whenDeviceLockedExceptOpenFiles"
)

func PossibleValuesForManagedAppDataEncryptionType() []string {
	return []string{
		string(ManagedAppDataEncryptionType_AfterDeviceRestart),
		string(ManagedAppDataEncryptionType_UseDeviceSettings),
		string(ManagedAppDataEncryptionType_WhenDeviceLocked),
		string(ManagedAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles),
	}
}

func (s *ManagedAppDataEncryptionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppDataEncryptionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppDataEncryptionType(input string) (*ManagedAppDataEncryptionType, error) {
	vals := map[string]ManagedAppDataEncryptionType{
		"afterdevicerestart":              ManagedAppDataEncryptionType_AfterDeviceRestart,
		"usedevicesettings":               ManagedAppDataEncryptionType_UseDeviceSettings,
		"whendevicelocked":                ManagedAppDataEncryptionType_WhenDeviceLocked,
		"whendevicelockedexceptopenfiles": ManagedAppDataEncryptionType_WhenDeviceLockedExceptOpenFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppDataEncryptionType(input)
	return &out, nil
}
