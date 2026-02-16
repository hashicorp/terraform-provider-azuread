package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessState string

const (
	DeviceManagementExchangeAccessState_Allowed     DeviceManagementExchangeAccessState = "allowed"
	DeviceManagementExchangeAccessState_Blocked     DeviceManagementExchangeAccessState = "blocked"
	DeviceManagementExchangeAccessState_None        DeviceManagementExchangeAccessState = "none"
	DeviceManagementExchangeAccessState_Quarantined DeviceManagementExchangeAccessState = "quarantined"
	DeviceManagementExchangeAccessState_Unknown     DeviceManagementExchangeAccessState = "unknown"
)

func PossibleValuesForDeviceManagementExchangeAccessState() []string {
	return []string{
		string(DeviceManagementExchangeAccessState_Allowed),
		string(DeviceManagementExchangeAccessState_Blocked),
		string(DeviceManagementExchangeAccessState_None),
		string(DeviceManagementExchangeAccessState_Quarantined),
		string(DeviceManagementExchangeAccessState_Unknown),
	}
}

func (s *DeviceManagementExchangeAccessState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeAccessState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeAccessState(input string) (*DeviceManagementExchangeAccessState, error) {
	vals := map[string]DeviceManagementExchangeAccessState{
		"allowed":     DeviceManagementExchangeAccessState_Allowed,
		"blocked":     DeviceManagementExchangeAccessState_Blocked,
		"none":        DeviceManagementExchangeAccessState_None,
		"quarantined": DeviceManagementExchangeAccessState_Quarantined,
		"unknown":     DeviceManagementExchangeAccessState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeAccessState(input)
	return &out, nil
}
