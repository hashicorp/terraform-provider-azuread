package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConnectionStatus string

const (
	NetworkaccessConnectionStatus_Active NetworkaccessConnectionStatus = "active"
	NetworkaccessConnectionStatus_Closed NetworkaccessConnectionStatus = "closed"
	NetworkaccessConnectionStatus_Open   NetworkaccessConnectionStatus = "open"
)

func PossibleValuesForNetworkaccessConnectionStatus() []string {
	return []string{
		string(NetworkaccessConnectionStatus_Active),
		string(NetworkaccessConnectionStatus_Closed),
		string(NetworkaccessConnectionStatus_Open),
	}
}

func (s *NetworkaccessConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessConnectionStatus(input string) (*NetworkaccessConnectionStatus, error) {
	vals := map[string]NetworkaccessConnectionStatus{
		"active": NetworkaccessConnectionStatus_Active,
		"closed": NetworkaccessConnectionStatus_Closed,
		"open":   NetworkaccessConnectionStatus_Open,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessConnectionStatus(input)
	return &out, nil
}
