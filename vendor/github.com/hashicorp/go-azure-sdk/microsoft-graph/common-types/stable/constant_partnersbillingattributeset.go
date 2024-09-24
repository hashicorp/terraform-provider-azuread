package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnersBillingAttributeSet string

const (
	PartnersBillingAttributeSet_Basic PartnersBillingAttributeSet = "basic"
	PartnersBillingAttributeSet_Full  PartnersBillingAttributeSet = "full"
)

func PossibleValuesForPartnersBillingAttributeSet() []string {
	return []string{
		string(PartnersBillingAttributeSet_Basic),
		string(PartnersBillingAttributeSet_Full),
	}
}

func (s *PartnersBillingAttributeSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnersBillingAttributeSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnersBillingAttributeSet(input string) (*PartnersBillingAttributeSet, error) {
	vals := map[string]PartnersBillingAttributeSet{
		"basic": PartnersBillingAttributeSet_Basic,
		"full":  PartnersBillingAttributeSet_Full,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnersBillingAttributeSet(input)
	return &out, nil
}
