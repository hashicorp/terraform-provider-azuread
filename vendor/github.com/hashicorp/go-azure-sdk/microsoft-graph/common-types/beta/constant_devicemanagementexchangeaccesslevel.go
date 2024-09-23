package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessLevel string

const (
	DeviceManagementExchangeAccessLevel_Allow      DeviceManagementExchangeAccessLevel = "allow"
	DeviceManagementExchangeAccessLevel_Block      DeviceManagementExchangeAccessLevel = "block"
	DeviceManagementExchangeAccessLevel_None       DeviceManagementExchangeAccessLevel = "none"
	DeviceManagementExchangeAccessLevel_Quarantine DeviceManagementExchangeAccessLevel = "quarantine"
)

func PossibleValuesForDeviceManagementExchangeAccessLevel() []string {
	return []string{
		string(DeviceManagementExchangeAccessLevel_Allow),
		string(DeviceManagementExchangeAccessLevel_Block),
		string(DeviceManagementExchangeAccessLevel_None),
		string(DeviceManagementExchangeAccessLevel_Quarantine),
	}
}

func (s *DeviceManagementExchangeAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeAccessLevel(input string) (*DeviceManagementExchangeAccessLevel, error) {
	vals := map[string]DeviceManagementExchangeAccessLevel{
		"allow":      DeviceManagementExchangeAccessLevel_Allow,
		"block":      DeviceManagementExchangeAccessLevel_Block,
		"none":       DeviceManagementExchangeAccessLevel_None,
		"quarantine": DeviceManagementExchangeAccessLevel_Quarantine,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeAccessLevel(input)
	return &out, nil
}
