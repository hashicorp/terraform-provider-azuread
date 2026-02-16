package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionDirection string

const (
	ConnectionDirection_Inbound  ConnectionDirection = "inbound"
	ConnectionDirection_Outbound ConnectionDirection = "outbound"
	ConnectionDirection_Unknown  ConnectionDirection = "unknown"
)

func PossibleValuesForConnectionDirection() []string {
	return []string{
		string(ConnectionDirection_Inbound),
		string(ConnectionDirection_Outbound),
		string(ConnectionDirection_Unknown),
	}
}

func (s *ConnectionDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionDirection(input string) (*ConnectionDirection, error) {
	vals := map[string]ConnectionDirection{
		"inbound":  ConnectionDirection_Inbound,
		"outbound": ConnectionDirection_Outbound,
		"unknown":  ConnectionDirection_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionDirection(input)
	return &out, nil
}
