package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoverySource string

const (
	DiscoverySource_AdminImport             DiscoverySource = "adminImport"
	DiscoverySource_DeviceEnrollmentProgram DiscoverySource = "deviceEnrollmentProgram"
	DiscoverySource_Unknown                 DiscoverySource = "unknown"
)

func PossibleValuesForDiscoverySource() []string {
	return []string{
		string(DiscoverySource_AdminImport),
		string(DiscoverySource_DeviceEnrollmentProgram),
		string(DiscoverySource_Unknown),
	}
}

func (s *DiscoverySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDiscoverySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDiscoverySource(input string) (*DiscoverySource, error) {
	vals := map[string]DiscoverySource{
		"adminimport":             DiscoverySource_AdminImport,
		"deviceenrollmentprogram": DiscoverySource_DeviceEnrollmentProgram,
		"unknown":                 DiscoverySource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiscoverySource(input)
	return &out, nil
}
