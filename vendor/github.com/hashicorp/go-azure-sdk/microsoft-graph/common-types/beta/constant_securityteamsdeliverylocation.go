package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTeamsDeliveryLocation string

const (
	SecurityTeamsDeliveryLocation_Failed     SecurityTeamsDeliveryLocation = "failed"
	SecurityTeamsDeliveryLocation_Quarantine SecurityTeamsDeliveryLocation = "quarantine"
	SecurityTeamsDeliveryLocation_Teams      SecurityTeamsDeliveryLocation = "teams"
	SecurityTeamsDeliveryLocation_Unknown    SecurityTeamsDeliveryLocation = "unknown"
)

func PossibleValuesForSecurityTeamsDeliveryLocation() []string {
	return []string{
		string(SecurityTeamsDeliveryLocation_Failed),
		string(SecurityTeamsDeliveryLocation_Quarantine),
		string(SecurityTeamsDeliveryLocation_Teams),
		string(SecurityTeamsDeliveryLocation_Unknown),
	}
}

func (s *SecurityTeamsDeliveryLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTeamsDeliveryLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTeamsDeliveryLocation(input string) (*SecurityTeamsDeliveryLocation, error) {
	vals := map[string]SecurityTeamsDeliveryLocation{
		"failed":     SecurityTeamsDeliveryLocation_Failed,
		"quarantine": SecurityTeamsDeliveryLocation_Quarantine,
		"teams":      SecurityTeamsDeliveryLocation_Teams,
		"unknown":    SecurityTeamsDeliveryLocation_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTeamsDeliveryLocation(input)
	return &out, nil
}
