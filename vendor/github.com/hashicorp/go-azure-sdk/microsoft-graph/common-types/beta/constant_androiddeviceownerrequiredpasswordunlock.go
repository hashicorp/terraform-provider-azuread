package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerRequiredPasswordUnlock string

const (
	AndroidDeviceOwnerRequiredPasswordUnlock_Daily             AndroidDeviceOwnerRequiredPasswordUnlock = "daily"
	AndroidDeviceOwnerRequiredPasswordUnlock_DeviceDefault     AndroidDeviceOwnerRequiredPasswordUnlock = "deviceDefault"
	AndroidDeviceOwnerRequiredPasswordUnlock_UnkownFutureValue AndroidDeviceOwnerRequiredPasswordUnlock = "unkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerRequiredPasswordUnlock() []string {
	return []string{
		string(AndroidDeviceOwnerRequiredPasswordUnlock_Daily),
		string(AndroidDeviceOwnerRequiredPasswordUnlock_DeviceDefault),
		string(AndroidDeviceOwnerRequiredPasswordUnlock_UnkownFutureValue),
	}
}

func (s *AndroidDeviceOwnerRequiredPasswordUnlock) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerRequiredPasswordUnlock(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerRequiredPasswordUnlock(input string) (*AndroidDeviceOwnerRequiredPasswordUnlock, error) {
	vals := map[string]AndroidDeviceOwnerRequiredPasswordUnlock{
		"daily":             AndroidDeviceOwnerRequiredPasswordUnlock_Daily,
		"devicedefault":     AndroidDeviceOwnerRequiredPasswordUnlock_DeviceDefault,
		"unkownfuturevalue": AndroidDeviceOwnerRequiredPasswordUnlock_UnkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerRequiredPasswordUnlock(input)
	return &out, nil
}
