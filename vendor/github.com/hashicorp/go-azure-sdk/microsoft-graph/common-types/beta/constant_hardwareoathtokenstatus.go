package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareOathTokenStatus string

const (
	HardwareOathTokenStatus_Activated        HardwareOathTokenStatus = "activated"
	HardwareOathTokenStatus_Assigned         HardwareOathTokenStatus = "assigned"
	HardwareOathTokenStatus_Available        HardwareOathTokenStatus = "available"
	HardwareOathTokenStatus_FailedActivation HardwareOathTokenStatus = "failedActivation"
)

func PossibleValuesForHardwareOathTokenStatus() []string {
	return []string{
		string(HardwareOathTokenStatus_Activated),
		string(HardwareOathTokenStatus_Assigned),
		string(HardwareOathTokenStatus_Available),
		string(HardwareOathTokenStatus_FailedActivation),
	}
}

func (s *HardwareOathTokenStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareOathTokenStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareOathTokenStatus(input string) (*HardwareOathTokenStatus, error) {
	vals := map[string]HardwareOathTokenStatus{
		"activated":        HardwareOathTokenStatus_Activated,
		"assigned":         HardwareOathTokenStatus_Assigned,
		"available":        HardwareOathTokenStatus_Available,
		"failedactivation": HardwareOathTokenStatus_FailedActivation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareOathTokenStatus(input)
	return &out, nil
}
