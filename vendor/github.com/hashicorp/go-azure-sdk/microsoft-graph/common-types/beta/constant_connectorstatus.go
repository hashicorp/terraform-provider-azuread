package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatus string

const (
	ConnectorStatus_Active   ConnectorStatus = "active"
	ConnectorStatus_Inactive ConnectorStatus = "inactive"
)

func PossibleValuesForConnectorStatus() []string {
	return []string{
		string(ConnectorStatus_Active),
		string(ConnectorStatus_Inactive),
	}
}

func (s *ConnectorStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorStatus(input string) (*ConnectorStatus, error) {
	vals := map[string]ConnectorStatus{
		"active":   ConnectorStatus_Active,
		"inactive": ConnectorStatus_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorStatus(input)
	return &out, nil
}
