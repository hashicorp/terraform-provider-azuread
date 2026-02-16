package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateDelayPolicy string

const (
	MacOSSoftwareUpdateDelayPolicy_DelayAppUpdateVisibility     MacOSSoftwareUpdateDelayPolicy = "delayAppUpdateVisibility"
	MacOSSoftwareUpdateDelayPolicy_DelayMajorOsUpdateVisibility MacOSSoftwareUpdateDelayPolicy = "delayMajorOsUpdateVisibility"
	MacOSSoftwareUpdateDelayPolicy_DelayOSUpdateVisibility      MacOSSoftwareUpdateDelayPolicy = "delayOSUpdateVisibility"
	MacOSSoftwareUpdateDelayPolicy_None                         MacOSSoftwareUpdateDelayPolicy = "none"
)

func PossibleValuesForMacOSSoftwareUpdateDelayPolicy() []string {
	return []string{
		string(MacOSSoftwareUpdateDelayPolicy_DelayAppUpdateVisibility),
		string(MacOSSoftwareUpdateDelayPolicy_DelayMajorOsUpdateVisibility),
		string(MacOSSoftwareUpdateDelayPolicy_DelayOSUpdateVisibility),
		string(MacOSSoftwareUpdateDelayPolicy_None),
	}
}

func (s *MacOSSoftwareUpdateDelayPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateDelayPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateDelayPolicy(input string) (*MacOSSoftwareUpdateDelayPolicy, error) {
	vals := map[string]MacOSSoftwareUpdateDelayPolicy{
		"delayappupdatevisibility":     MacOSSoftwareUpdateDelayPolicy_DelayAppUpdateVisibility,
		"delaymajorosupdatevisibility": MacOSSoftwareUpdateDelayPolicy_DelayMajorOsUpdateVisibility,
		"delayosupdatevisibility":      MacOSSoftwareUpdateDelayPolicy_DelayOSUpdateVisibility,
		"none":                         MacOSSoftwareUpdateDelayPolicy_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateDelayPolicy(input)
	return &out, nil
}
