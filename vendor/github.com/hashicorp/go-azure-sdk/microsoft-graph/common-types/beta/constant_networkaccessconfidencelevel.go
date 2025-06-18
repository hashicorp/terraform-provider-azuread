package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConfidenceLevel string

const (
	NetworkaccessConfidenceLevel_High    NetworkaccessConfidenceLevel = "high"
	NetworkaccessConfidenceLevel_Low     NetworkaccessConfidenceLevel = "low"
	NetworkaccessConfidenceLevel_Unknown NetworkaccessConfidenceLevel = "unknown"
)

func PossibleValuesForNetworkaccessConfidenceLevel() []string {
	return []string{
		string(NetworkaccessConfidenceLevel_High),
		string(NetworkaccessConfidenceLevel_Low),
		string(NetworkaccessConfidenceLevel_Unknown),
	}
}

func (s *NetworkaccessConfidenceLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessConfidenceLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessConfidenceLevel(input string) (*NetworkaccessConfidenceLevel, error) {
	vals := map[string]NetworkaccessConfidenceLevel{
		"high":    NetworkaccessConfidenceLevel_High,
		"low":     NetworkaccessConfidenceLevel_Low,
		"unknown": NetworkaccessConfidenceLevel_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessConfidenceLevel(input)
	return &out, nil
}
