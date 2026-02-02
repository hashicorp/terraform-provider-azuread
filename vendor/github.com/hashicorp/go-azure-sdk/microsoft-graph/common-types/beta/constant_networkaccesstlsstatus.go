package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTlsStatus string

const (
	NetworkaccessTlsStatus_Failure NetworkaccessTlsStatus = "failure"
	NetworkaccessTlsStatus_Success NetworkaccessTlsStatus = "success"
)

func PossibleValuesForNetworkaccessTlsStatus() []string {
	return []string{
		string(NetworkaccessTlsStatus_Failure),
		string(NetworkaccessTlsStatus_Success),
	}
}

func (s *NetworkaccessTlsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTlsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTlsStatus(input string) (*NetworkaccessTlsStatus, error) {
	vals := map[string]NetworkaccessTlsStatus{
		"failure": NetworkaccessTlsStatus_Failure,
		"success": NetworkaccessTlsStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTlsStatus(input)
	return &out, nil
}
