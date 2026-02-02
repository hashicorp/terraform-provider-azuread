package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnersBillingBillingPeriod string

const (
	PartnersBillingBillingPeriod_Current PartnersBillingBillingPeriod = "current"
	PartnersBillingBillingPeriod_Last    PartnersBillingBillingPeriod = "last"
)

func PossibleValuesForPartnersBillingBillingPeriod() []string {
	return []string{
		string(PartnersBillingBillingPeriod_Current),
		string(PartnersBillingBillingPeriod_Last),
	}
}

func (s *PartnersBillingBillingPeriod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnersBillingBillingPeriod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnersBillingBillingPeriod(input string) (*PartnersBillingBillingPeriod, error) {
	vals := map[string]PartnersBillingBillingPeriod{
		"current": PartnersBillingBillingPeriod_Current,
		"last":    PartnersBillingBillingPeriod_Last,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnersBillingBillingPeriod(input)
	return &out, nil
}
