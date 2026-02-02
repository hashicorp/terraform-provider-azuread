package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessStatus string

const (
	NetworkaccessStatus_Disabled NetworkaccessStatus = "disabled"
	NetworkaccessStatus_Enabled  NetworkaccessStatus = "enabled"
)

func PossibleValuesForNetworkaccessStatus() []string {
	return []string{
		string(NetworkaccessStatus_Disabled),
		string(NetworkaccessStatus_Enabled),
	}
}

func (s *NetworkaccessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessStatus(input string) (*NetworkaccessStatus, error) {
	vals := map[string]NetworkaccessStatus{
		"disabled": NetworkaccessStatus_Disabled,
		"enabled":  NetworkaccessStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessStatus(input)
	return &out, nil
}
