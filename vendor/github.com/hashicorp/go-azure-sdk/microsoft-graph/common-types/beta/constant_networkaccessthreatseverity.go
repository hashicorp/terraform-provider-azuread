package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessThreatSeverity string

const (
	NetworkaccessThreatSeverity_Critical NetworkaccessThreatSeverity = "critical"
	NetworkaccessThreatSeverity_High     NetworkaccessThreatSeverity = "high"
	NetworkaccessThreatSeverity_Low      NetworkaccessThreatSeverity = "low"
	NetworkaccessThreatSeverity_Medium   NetworkaccessThreatSeverity = "medium"
)

func PossibleValuesForNetworkaccessThreatSeverity() []string {
	return []string{
		string(NetworkaccessThreatSeverity_Critical),
		string(NetworkaccessThreatSeverity_High),
		string(NetworkaccessThreatSeverity_Low),
		string(NetworkaccessThreatSeverity_Medium),
	}
}

func (s *NetworkaccessThreatSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessThreatSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessThreatSeverity(input string) (*NetworkaccessThreatSeverity, error) {
	vals := map[string]NetworkaccessThreatSeverity{
		"critical": NetworkaccessThreatSeverity_Critical,
		"high":     NetworkaccessThreatSeverity_High,
		"low":      NetworkaccessThreatSeverity_Low,
		"medium":   NetworkaccessThreatSeverity_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessThreatSeverity(input)
	return &out, nil
}
