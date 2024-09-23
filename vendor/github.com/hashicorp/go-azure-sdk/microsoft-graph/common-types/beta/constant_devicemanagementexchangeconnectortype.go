package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeConnectorType string

const (
	DeviceManagementExchangeConnectorType_Dedicated        DeviceManagementExchangeConnectorType = "dedicated"
	DeviceManagementExchangeConnectorType_Hosted           DeviceManagementExchangeConnectorType = "hosted"
	DeviceManagementExchangeConnectorType_OnPremises       DeviceManagementExchangeConnectorType = "onPremises"
	DeviceManagementExchangeConnectorType_ServiceToService DeviceManagementExchangeConnectorType = "serviceToService"
)

func PossibleValuesForDeviceManagementExchangeConnectorType() []string {
	return []string{
		string(DeviceManagementExchangeConnectorType_Dedicated),
		string(DeviceManagementExchangeConnectorType_Hosted),
		string(DeviceManagementExchangeConnectorType_OnPremises),
		string(DeviceManagementExchangeConnectorType_ServiceToService),
	}
}

func (s *DeviceManagementExchangeConnectorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeConnectorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeConnectorType(input string) (*DeviceManagementExchangeConnectorType, error) {
	vals := map[string]DeviceManagementExchangeConnectorType{
		"dedicated":        DeviceManagementExchangeConnectorType_Dedicated,
		"hosted":           DeviceManagementExchangeConnectorType_Hosted,
		"onpremises":       DeviceManagementExchangeConnectorType_OnPremises,
		"servicetoservice": DeviceManagementExchangeConnectorType_ServiceToService,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeConnectorType(input)
	return &out, nil
}
