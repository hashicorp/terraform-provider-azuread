package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DecisionItemPrincipalResourceMembershipType string

const (
	DecisionItemPrincipalResourceMembershipType_Direct   DecisionItemPrincipalResourceMembershipType = "direct"
	DecisionItemPrincipalResourceMembershipType_Indirect DecisionItemPrincipalResourceMembershipType = "indirect"
)

func PossibleValuesForDecisionItemPrincipalResourceMembershipType() []string {
	return []string{
		string(DecisionItemPrincipalResourceMembershipType_Direct),
		string(DecisionItemPrincipalResourceMembershipType_Indirect),
	}
}

func (s *DecisionItemPrincipalResourceMembershipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDecisionItemPrincipalResourceMembershipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDecisionItemPrincipalResourceMembershipType(input string) (*DecisionItemPrincipalResourceMembershipType, error) {
	vals := map[string]DecisionItemPrincipalResourceMembershipType{
		"direct":   DecisionItemPrincipalResourceMembershipType_Direct,
		"indirect": DecisionItemPrincipalResourceMembershipType_Indirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DecisionItemPrincipalResourceMembershipType(input)
	return &out, nil
}
