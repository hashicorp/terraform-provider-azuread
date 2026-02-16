package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorSyncType string

const (
	DeviceManagementExchangeConnectorSyncType_DeltaSync DeviceManagementExchangeConnectorSyncType = "deltaSync"
	DeviceManagementExchangeConnectorSyncType_FullSync  DeviceManagementExchangeConnectorSyncType = "fullSync"
)

func PossibleValuesForDeviceManagementExchangeConnectorSyncType() []string {
	return []string{
		string(DeviceManagementExchangeConnectorSyncType_DeltaSync),
		string(DeviceManagementExchangeConnectorSyncType_FullSync),
	}
}

func (s *DeviceManagementExchangeConnectorSyncType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeConnectorSyncType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeConnectorSyncType(input string) (*DeviceManagementExchangeConnectorSyncType, error) {
	vals := map[string]DeviceManagementExchangeConnectorSyncType{
		"deltasync": DeviceManagementExchangeConnectorSyncType_DeltaSync,
		"fullsync":  DeviceManagementExchangeConnectorSyncType_FullSync,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeConnectorSyncType(input)
	return &out, nil
}
