package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSContentCachingParentSelectionPolicy string

const (
	MacOSContentCachingParentSelectionPolicy_FirstAvailable  MacOSContentCachingParentSelectionPolicy = "firstAvailable"
	MacOSContentCachingParentSelectionPolicy_NotConfigured   MacOSContentCachingParentSelectionPolicy = "notConfigured"
	MacOSContentCachingParentSelectionPolicy_Random          MacOSContentCachingParentSelectionPolicy = "random"
	MacOSContentCachingParentSelectionPolicy_RoundRobin      MacOSContentCachingParentSelectionPolicy = "roundRobin"
	MacOSContentCachingParentSelectionPolicy_StickyAvailable MacOSContentCachingParentSelectionPolicy = "stickyAvailable"
	MacOSContentCachingParentSelectionPolicy_UrlPathHash     MacOSContentCachingParentSelectionPolicy = "urlPathHash"
)

func PossibleValuesForMacOSContentCachingParentSelectionPolicy() []string {
	return []string{
		string(MacOSContentCachingParentSelectionPolicy_FirstAvailable),
		string(MacOSContentCachingParentSelectionPolicy_NotConfigured),
		string(MacOSContentCachingParentSelectionPolicy_Random),
		string(MacOSContentCachingParentSelectionPolicy_RoundRobin),
		string(MacOSContentCachingParentSelectionPolicy_StickyAvailable),
		string(MacOSContentCachingParentSelectionPolicy_UrlPathHash),
	}
}

func (s *MacOSContentCachingParentSelectionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSContentCachingParentSelectionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSContentCachingParentSelectionPolicy(input string) (*MacOSContentCachingParentSelectionPolicy, error) {
	vals := map[string]MacOSContentCachingParentSelectionPolicy{
		"firstavailable":  MacOSContentCachingParentSelectionPolicy_FirstAvailable,
		"notconfigured":   MacOSContentCachingParentSelectionPolicy_NotConfigured,
		"random":          MacOSContentCachingParentSelectionPolicy_Random,
		"roundrobin":      MacOSContentCachingParentSelectionPolicy_RoundRobin,
		"stickyavailable": MacOSContentCachingParentSelectionPolicy_StickyAvailable,
		"urlpathhash":     MacOSContentCachingParentSelectionPolicy_UrlPathHash,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSContentCachingParentSelectionPolicy(input)
	return &out, nil
}
