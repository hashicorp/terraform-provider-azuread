package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperatingSystemUpgradeEligibility string

const (
	OperatingSystemUpgradeEligibility_Capable    OperatingSystemUpgradeEligibility = "capable"
	OperatingSystemUpgradeEligibility_NotCapable OperatingSystemUpgradeEligibility = "notCapable"
	OperatingSystemUpgradeEligibility_Unknown    OperatingSystemUpgradeEligibility = "unknown"
	OperatingSystemUpgradeEligibility_Upgraded   OperatingSystemUpgradeEligibility = "upgraded"
)

func PossibleValuesForOperatingSystemUpgradeEligibility() []string {
	return []string{
		string(OperatingSystemUpgradeEligibility_Capable),
		string(OperatingSystemUpgradeEligibility_NotCapable),
		string(OperatingSystemUpgradeEligibility_Unknown),
		string(OperatingSystemUpgradeEligibility_Upgraded),
	}
}

func (s *OperatingSystemUpgradeEligibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperatingSystemUpgradeEligibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperatingSystemUpgradeEligibility(input string) (*OperatingSystemUpgradeEligibility, error) {
	vals := map[string]OperatingSystemUpgradeEligibility{
		"capable":    OperatingSystemUpgradeEligibility_Capable,
		"notcapable": OperatingSystemUpgradeEligibility_NotCapable,
		"unknown":    OperatingSystemUpgradeEligibility_Unknown,
		"upgraded":   OperatingSystemUpgradeEligibility_Upgraded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatingSystemUpgradeEligibility(input)
	return &out, nil
}
