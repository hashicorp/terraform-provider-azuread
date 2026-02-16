package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionStatus string

const (
	ConnectionStatus_Attempted ConnectionStatus = "attempted"
	ConnectionStatus_Blocked   ConnectionStatus = "blocked"
	ConnectionStatus_Failed    ConnectionStatus = "failed"
	ConnectionStatus_Succeeded ConnectionStatus = "succeeded"
	ConnectionStatus_Unknown   ConnectionStatus = "unknown"
)

func PossibleValuesForConnectionStatus() []string {
	return []string{
		string(ConnectionStatus_Attempted),
		string(ConnectionStatus_Blocked),
		string(ConnectionStatus_Failed),
		string(ConnectionStatus_Succeeded),
		string(ConnectionStatus_Unknown),
	}
}

func (s *ConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionStatus(input string) (*ConnectionStatus, error) {
	vals := map[string]ConnectionStatus{
		"attempted": ConnectionStatus_Attempted,
		"blocked":   ConnectionStatus_Blocked,
		"failed":    ConnectionStatus_Failed,
		"succeeded": ConnectionStatus_Succeeded,
		"unknown":   ConnectionStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionStatus(input)
	return &out, nil
}
