package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimConditionUserType string

const (
	ClaimConditionUserType_AadGuests      ClaimConditionUserType = "aadGuests"
	ClaimConditionUserType_AllGuests      ClaimConditionUserType = "allGuests"
	ClaimConditionUserType_Any            ClaimConditionUserType = "any"
	ClaimConditionUserType_ExternalGuests ClaimConditionUserType = "externalGuests"
	ClaimConditionUserType_Members        ClaimConditionUserType = "members"
)

func PossibleValuesForClaimConditionUserType() []string {
	return []string{
		string(ClaimConditionUserType_AadGuests),
		string(ClaimConditionUserType_AllGuests),
		string(ClaimConditionUserType_Any),
		string(ClaimConditionUserType_ExternalGuests),
		string(ClaimConditionUserType_Members),
	}
}

func (s *ClaimConditionUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClaimConditionUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClaimConditionUserType(input string) (*ClaimConditionUserType, error) {
	vals := map[string]ClaimConditionUserType{
		"aadguests":      ClaimConditionUserType_AadGuests,
		"allguests":      ClaimConditionUserType_AllGuests,
		"any":            ClaimConditionUserType_Any,
		"externalguests": ClaimConditionUserType_ExternalGuests,
		"members":        ClaimConditionUserType_Members,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClaimConditionUserType(input)
	return &out, nil
}
