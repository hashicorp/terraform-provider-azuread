package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUsageStatus string

const (
	NetworkaccessUsageStatus_FrequentlyUsed NetworkaccessUsageStatus = "frequentlyUsed"
	NetworkaccessUsageStatus_RarelyUsed     NetworkaccessUsageStatus = "rarelyUsed"
)

func PossibleValuesForNetworkaccessUsageStatus() []string {
	return []string{
		string(NetworkaccessUsageStatus_FrequentlyUsed),
		string(NetworkaccessUsageStatus_RarelyUsed),
	}
}

func (s *NetworkaccessUsageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessUsageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessUsageStatus(input string) (*NetworkaccessUsageStatus, error) {
	vals := map[string]NetworkaccessUsageStatus{
		"frequentlyused": NetworkaccessUsageStatus_FrequentlyUsed,
		"rarelyused":     NetworkaccessUsageStatus_RarelyUsed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessUsageStatus(input)
	return &out, nil
}
