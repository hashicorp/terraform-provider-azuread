package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FrontlineCloudPCAvailability string

const (
	FrontlineCloudPCAvailability_Available     FrontlineCloudPCAvailability = "available"
	FrontlineCloudPCAvailability_NotApplicable FrontlineCloudPCAvailability = "notApplicable"
	FrontlineCloudPCAvailability_NotAvailable  FrontlineCloudPCAvailability = "notAvailable"
)

func PossibleValuesForFrontlineCloudPCAvailability() []string {
	return []string{
		string(FrontlineCloudPCAvailability_Available),
		string(FrontlineCloudPCAvailability_NotApplicable),
		string(FrontlineCloudPCAvailability_NotAvailable),
	}
}

func (s *FrontlineCloudPCAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFrontlineCloudPCAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFrontlineCloudPCAvailability(input string) (*FrontlineCloudPCAvailability, error) {
	vals := map[string]FrontlineCloudPCAvailability{
		"available":     FrontlineCloudPCAvailability_Available,
		"notapplicable": FrontlineCloudPCAvailability_NotApplicable,
		"notavailable":  FrontlineCloudPCAvailability_NotAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FrontlineCloudPCAvailability(input)
	return &out, nil
}
