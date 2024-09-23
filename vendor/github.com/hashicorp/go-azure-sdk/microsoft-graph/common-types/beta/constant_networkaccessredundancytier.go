package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRedundancyTier string

const (
	NetworkaccessRedundancyTier_NoRedundancy   NetworkaccessRedundancyTier = "noRedundancy"
	NetworkaccessRedundancyTier_ZoneRedundancy NetworkaccessRedundancyTier = "zoneRedundancy"
)

func PossibleValuesForNetworkaccessRedundancyTier() []string {
	return []string{
		string(NetworkaccessRedundancyTier_NoRedundancy),
		string(NetworkaccessRedundancyTier_ZoneRedundancy),
	}
}

func (s *NetworkaccessRedundancyTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessRedundancyTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessRedundancyTier(input string) (*NetworkaccessRedundancyTier, error) {
	vals := map[string]NetworkaccessRedundancyTier{
		"noredundancy":   NetworkaccessRedundancyTier_NoRedundancy,
		"zoneredundancy": NetworkaccessRedundancyTier_ZoneRedundancy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessRedundancyTier(input)
	return &out, nil
}
