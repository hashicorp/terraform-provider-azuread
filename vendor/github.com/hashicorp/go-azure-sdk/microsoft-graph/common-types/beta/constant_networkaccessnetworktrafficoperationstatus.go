package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkTrafficOperationStatus string

const (
	NetworkaccessNetworkTrafficOperationStatus_Failure NetworkaccessNetworkTrafficOperationStatus = "failure"
	NetworkaccessNetworkTrafficOperationStatus_Success NetworkaccessNetworkTrafficOperationStatus = "success"
)

func PossibleValuesForNetworkaccessNetworkTrafficOperationStatus() []string {
	return []string{
		string(NetworkaccessNetworkTrafficOperationStatus_Failure),
		string(NetworkaccessNetworkTrafficOperationStatus_Success),
	}
}

func (s *NetworkaccessNetworkTrafficOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkTrafficOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkTrafficOperationStatus(input string) (*NetworkaccessNetworkTrafficOperationStatus, error) {
	vals := map[string]NetworkaccessNetworkTrafficOperationStatus{
		"failure": NetworkaccessNetworkTrafficOperationStatus_Failure,
		"success": NetworkaccessNetworkTrafficOperationStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkTrafficOperationStatus(input)
	return &out, nil
}
