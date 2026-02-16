package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceIdEntityIdentifier string

const (
	SecurityDeviceIdEntityIdentifier_DeviceId SecurityDeviceIdEntityIdentifier = "deviceId"
)

func PossibleValuesForSecurityDeviceIdEntityIdentifier() []string {
	return []string{
		string(SecurityDeviceIdEntityIdentifier_DeviceId),
	}
}

func (s *SecurityDeviceIdEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceIdEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceIdEntityIdentifier(input string) (*SecurityDeviceIdEntityIdentifier, error) {
	vals := map[string]SecurityDeviceIdEntityIdentifier{
		"deviceid": SecurityDeviceIdEntityIdentifier_DeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceIdEntityIdentifier(input)
	return &out, nil
}
