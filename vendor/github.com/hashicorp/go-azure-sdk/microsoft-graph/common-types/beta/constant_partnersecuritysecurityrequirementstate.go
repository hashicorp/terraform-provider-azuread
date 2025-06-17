package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityRequirementState string

const (
	PartnerSecuritySecurityRequirementState_Active  PartnerSecuritySecurityRequirementState = "active"
	PartnerSecuritySecurityRequirementState_Preview PartnerSecuritySecurityRequirementState = "preview"
)

func PossibleValuesForPartnerSecuritySecurityRequirementState() []string {
	return []string{
		string(PartnerSecuritySecurityRequirementState_Active),
		string(PartnerSecuritySecurityRequirementState_Preview),
	}
}

func (s *PartnerSecuritySecurityRequirementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecuritySecurityRequirementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecuritySecurityRequirementState(input string) (*PartnerSecuritySecurityRequirementState, error) {
	vals := map[string]PartnerSecuritySecurityRequirementState{
		"active":  PartnerSecuritySecurityRequirementState_Active,
		"preview": PartnerSecuritySecurityRequirementState_Preview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecuritySecurityRequirementState(input)
	return &out, nil
}
