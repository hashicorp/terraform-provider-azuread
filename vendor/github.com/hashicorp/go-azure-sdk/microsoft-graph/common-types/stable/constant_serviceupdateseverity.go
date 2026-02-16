package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateSeverity string

const (
	ServiceUpdateSeverity_Critical ServiceUpdateSeverity = "critical"
	ServiceUpdateSeverity_High     ServiceUpdateSeverity = "high"
	ServiceUpdateSeverity_Normal   ServiceUpdateSeverity = "normal"
)

func PossibleValuesForServiceUpdateSeverity() []string {
	return []string{
		string(ServiceUpdateSeverity_Critical),
		string(ServiceUpdateSeverity_High),
		string(ServiceUpdateSeverity_Normal),
	}
}

func (s *ServiceUpdateSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceUpdateSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceUpdateSeverity(input string) (*ServiceUpdateSeverity, error) {
	vals := map[string]ServiceUpdateSeverity{
		"critical": ServiceUpdateSeverity_Critical,
		"high":     ServiceUpdateSeverity_High,
		"normal":   ServiceUpdateSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceUpdateSeverity(input)
	return &out, nil
}
