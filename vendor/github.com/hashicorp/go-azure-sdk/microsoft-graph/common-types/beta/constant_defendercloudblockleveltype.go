package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderCloudBlockLevelType string

const (
	DefenderCloudBlockLevelType_High          DefenderCloudBlockLevelType = "high"
	DefenderCloudBlockLevelType_HighPlus      DefenderCloudBlockLevelType = "highPlus"
	DefenderCloudBlockLevelType_NotConfigured DefenderCloudBlockLevelType = "notConfigured"
	DefenderCloudBlockLevelType_ZeroTolerance DefenderCloudBlockLevelType = "zeroTolerance"
)

func PossibleValuesForDefenderCloudBlockLevelType() []string {
	return []string{
		string(DefenderCloudBlockLevelType_High),
		string(DefenderCloudBlockLevelType_HighPlus),
		string(DefenderCloudBlockLevelType_NotConfigured),
		string(DefenderCloudBlockLevelType_ZeroTolerance),
	}
}

func (s *DefenderCloudBlockLevelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderCloudBlockLevelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderCloudBlockLevelType(input string) (*DefenderCloudBlockLevelType, error) {
	vals := map[string]DefenderCloudBlockLevelType{
		"high":          DefenderCloudBlockLevelType_High,
		"highplus":      DefenderCloudBlockLevelType_HighPlus,
		"notconfigured": DefenderCloudBlockLevelType_NotConfigured,
		"zerotolerance": DefenderCloudBlockLevelType_ZeroTolerance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderCloudBlockLevelType(input)
	return &out, nil
}
