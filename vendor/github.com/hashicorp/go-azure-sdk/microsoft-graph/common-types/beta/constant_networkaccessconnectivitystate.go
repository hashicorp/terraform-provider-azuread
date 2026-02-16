package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConnectivityState string

const (
	NetworkaccessConnectivityState_Connected NetworkaccessConnectivityState = "connected"
	NetworkaccessConnectivityState_Error     NetworkaccessConnectivityState = "error"
	NetworkaccessConnectivityState_Inactive  NetworkaccessConnectivityState = "inactive"
	NetworkaccessConnectivityState_Pending   NetworkaccessConnectivityState = "pending"
)

func PossibleValuesForNetworkaccessConnectivityState() []string {
	return []string{
		string(NetworkaccessConnectivityState_Connected),
		string(NetworkaccessConnectivityState_Error),
		string(NetworkaccessConnectivityState_Inactive),
		string(NetworkaccessConnectivityState_Pending),
	}
}

func (s *NetworkaccessConnectivityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessConnectivityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessConnectivityState(input string) (*NetworkaccessConnectivityState, error) {
	vals := map[string]NetworkaccessConnectivityState{
		"connected": NetworkaccessConnectivityState_Connected,
		"error":     NetworkaccessConnectivityState_Error,
		"inactive":  NetworkaccessConnectivityState_Inactive,
		"pending":   NetworkaccessConnectivityState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessConnectivityState(input)
	return &out, nil
}
