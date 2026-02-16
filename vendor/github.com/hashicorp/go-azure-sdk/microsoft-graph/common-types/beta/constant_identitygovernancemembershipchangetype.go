package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceMembershipChangeType string

const (
	IdentityGovernanceMembershipChangeType_Add    IdentityGovernanceMembershipChangeType = "add"
	IdentityGovernanceMembershipChangeType_Remove IdentityGovernanceMembershipChangeType = "remove"
)

func PossibleValuesForIdentityGovernanceMembershipChangeType() []string {
	return []string{
		string(IdentityGovernanceMembershipChangeType_Add),
		string(IdentityGovernanceMembershipChangeType_Remove),
	}
}

func (s *IdentityGovernanceMembershipChangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceMembershipChangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceMembershipChangeType(input string) (*IdentityGovernanceMembershipChangeType, error) {
	vals := map[string]IdentityGovernanceMembershipChangeType{
		"add":    IdentityGovernanceMembershipChangeType_Add,
		"remove": IdentityGovernanceMembershipChangeType_Remove,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceMembershipChangeType(input)
	return &out, nil
}
