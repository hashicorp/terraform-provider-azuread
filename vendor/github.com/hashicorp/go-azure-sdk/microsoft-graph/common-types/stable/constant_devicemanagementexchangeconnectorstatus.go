package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorStatus string

const (
	DeviceManagementExchangeConnectorStatus_Connected         DeviceManagementExchangeConnectorStatus = "connected"
	DeviceManagementExchangeConnectorStatus_ConnectionPending DeviceManagementExchangeConnectorStatus = "connectionPending"
	DeviceManagementExchangeConnectorStatus_Disconnected      DeviceManagementExchangeConnectorStatus = "disconnected"
	DeviceManagementExchangeConnectorStatus_None              DeviceManagementExchangeConnectorStatus = "none"
)

func PossibleValuesForDeviceManagementExchangeConnectorStatus() []string {
	return []string{
		string(DeviceManagementExchangeConnectorStatus_Connected),
		string(DeviceManagementExchangeConnectorStatus_ConnectionPending),
		string(DeviceManagementExchangeConnectorStatus_Disconnected),
		string(DeviceManagementExchangeConnectorStatus_None),
	}
}

func (s *DeviceManagementExchangeConnectorStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeConnectorStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeConnectorStatus(input string) (*DeviceManagementExchangeConnectorStatus, error) {
	vals := map[string]DeviceManagementExchangeConnectorStatus{
		"connected":         DeviceManagementExchangeConnectorStatus_Connected,
		"connectionpending": DeviceManagementExchangeConnectorStatus_ConnectionPending,
		"disconnected":      DeviceManagementExchangeConnectorStatus_Disconnected,
		"none":              DeviceManagementExchangeConnectorStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeConnectorStatus(input)
	return &out, nil
}
