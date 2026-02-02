package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareConfigurationFormat string

const (
	HardwareConfigurationFormat_Dell        HardwareConfigurationFormat = "dell"
	HardwareConfigurationFormat_Surface     HardwareConfigurationFormat = "surface"
	HardwareConfigurationFormat_SurfaceDock HardwareConfigurationFormat = "surfaceDock"
)

func PossibleValuesForHardwareConfigurationFormat() []string {
	return []string{
		string(HardwareConfigurationFormat_Dell),
		string(HardwareConfigurationFormat_Surface),
		string(HardwareConfigurationFormat_SurfaceDock),
	}
}

func (s *HardwareConfigurationFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareConfigurationFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareConfigurationFormat(input string) (*HardwareConfigurationFormat, error) {
	vals := map[string]HardwareConfigurationFormat{
		"dell":        HardwareConfigurationFormat_Dell,
		"surface":     HardwareConfigurationFormat_Surface,
		"surfacedock": HardwareConfigurationFormat_SurfaceDock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareConfigurationFormat(input)
	return &out, nil
}
