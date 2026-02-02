package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportLocation string

const (
	SecurityExportLocation_NonresponsiveLocations SecurityExportLocation = "nonresponsiveLocations"
	SecurityExportLocation_ResponsiveLocations    SecurityExportLocation = "responsiveLocations"
)

func PossibleValuesForSecurityExportLocation() []string {
	return []string{
		string(SecurityExportLocation_NonresponsiveLocations),
		string(SecurityExportLocation_ResponsiveLocations),
	}
}

func (s *SecurityExportLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityExportLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityExportLocation(input string) (*SecurityExportLocation, error) {
	vals := map[string]SecurityExportLocation{
		"nonresponsivelocations": SecurityExportLocation_NonresponsiveLocations,
		"responsivelocations":    SecurityExportLocation_ResponsiveLocations,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityExportLocation(input)
	return &out, nil
}
