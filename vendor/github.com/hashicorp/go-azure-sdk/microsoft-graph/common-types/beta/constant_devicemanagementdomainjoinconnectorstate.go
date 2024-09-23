package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDomainJoinConnectorState string

const (
	DeviceManagementDomainJoinConnectorState_Active   DeviceManagementDomainJoinConnectorState = "active"
	DeviceManagementDomainJoinConnectorState_Error    DeviceManagementDomainJoinConnectorState = "error"
	DeviceManagementDomainJoinConnectorState_Inactive DeviceManagementDomainJoinConnectorState = "inactive"
)

func PossibleValuesForDeviceManagementDomainJoinConnectorState() []string {
	return []string{
		string(DeviceManagementDomainJoinConnectorState_Active),
		string(DeviceManagementDomainJoinConnectorState_Error),
		string(DeviceManagementDomainJoinConnectorState_Inactive),
	}
}

func (s *DeviceManagementDomainJoinConnectorState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementDomainJoinConnectorState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementDomainJoinConnectorState(input string) (*DeviceManagementDomainJoinConnectorState, error) {
	vals := map[string]DeviceManagementDomainJoinConnectorState{
		"active":   DeviceManagementDomainJoinConnectorState_Active,
		"error":    DeviceManagementDomainJoinConnectorState_Error,
		"inactive": DeviceManagementDomainJoinConnectorState_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDomainJoinConnectorState(input)
	return &out, nil
}
