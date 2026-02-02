package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorHealthState string

const (
	ConnectorHealthState_Healthy   ConnectorHealthState = "healthy"
	ConnectorHealthState_Unhealthy ConnectorHealthState = "unhealthy"
	ConnectorHealthState_Unknown   ConnectorHealthState = "unknown"
	ConnectorHealthState_Warning   ConnectorHealthState = "warning"
)

func PossibleValuesForConnectorHealthState() []string {
	return []string{
		string(ConnectorHealthState_Healthy),
		string(ConnectorHealthState_Unhealthy),
		string(ConnectorHealthState_Unknown),
		string(ConnectorHealthState_Warning),
	}
}

func (s *ConnectorHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorHealthState(input string) (*ConnectorHealthState, error) {
	vals := map[string]ConnectorHealthState{
		"healthy":   ConnectorHealthState_Healthy,
		"unhealthy": ConnectorHealthState_Unhealthy,
		"unknown":   ConnectorHealthState_Unknown,
		"warning":   ConnectorHealthState_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorHealthState(input)
	return &out, nil
}
