package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaConnectorState string

const (
	ZebraFotaConnectorState_Connected    ZebraFotaConnectorState = "connected"
	ZebraFotaConnectorState_Disconnected ZebraFotaConnectorState = "disconnected"
	ZebraFotaConnectorState_None         ZebraFotaConnectorState = "none"
)

func PossibleValuesForZebraFotaConnectorState() []string {
	return []string{
		string(ZebraFotaConnectorState_Connected),
		string(ZebraFotaConnectorState_Disconnected),
		string(ZebraFotaConnectorState_None),
	}
}

func (s *ZebraFotaConnectorState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaConnectorState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaConnectorState(input string) (*ZebraFotaConnectorState, error) {
	vals := map[string]ZebraFotaConnectorState{
		"connected":    ZebraFotaConnectorState_Connected,
		"disconnected": ZebraFotaConnectorState_Disconnected,
		"none":         ZebraFotaConnectorState_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaConnectorState(input)
	return &out, nil
}
