package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSubscriptionState string

const (
	DeviceManagementSubscriptionState_Active    DeviceManagementSubscriptionState = "active"
	DeviceManagementSubscriptionState_Blocked   DeviceManagementSubscriptionState = "blocked"
	DeviceManagementSubscriptionState_Deleted   DeviceManagementSubscriptionState = "deleted"
	DeviceManagementSubscriptionState_Disabled  DeviceManagementSubscriptionState = "disabled"
	DeviceManagementSubscriptionState_LockedOut DeviceManagementSubscriptionState = "lockedOut"
	DeviceManagementSubscriptionState_Pending   DeviceManagementSubscriptionState = "pending"
	DeviceManagementSubscriptionState_Warning   DeviceManagementSubscriptionState = "warning"
)

func PossibleValuesForDeviceManagementSubscriptionState() []string {
	return []string{
		string(DeviceManagementSubscriptionState_Active),
		string(DeviceManagementSubscriptionState_Blocked),
		string(DeviceManagementSubscriptionState_Deleted),
		string(DeviceManagementSubscriptionState_Disabled),
		string(DeviceManagementSubscriptionState_LockedOut),
		string(DeviceManagementSubscriptionState_Pending),
		string(DeviceManagementSubscriptionState_Warning),
	}
}

func (s *DeviceManagementSubscriptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementSubscriptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementSubscriptionState(input string) (*DeviceManagementSubscriptionState, error) {
	vals := map[string]DeviceManagementSubscriptionState{
		"active":    DeviceManagementSubscriptionState_Active,
		"blocked":   DeviceManagementSubscriptionState_Blocked,
		"deleted":   DeviceManagementSubscriptionState_Deleted,
		"disabled":  DeviceManagementSubscriptionState_Disabled,
		"lockedout": DeviceManagementSubscriptionState_LockedOut,
		"pending":   DeviceManagementSubscriptionState_Pending,
		"warning":   DeviceManagementSubscriptionState_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementSubscriptionState(input)
	return &out, nil
}
