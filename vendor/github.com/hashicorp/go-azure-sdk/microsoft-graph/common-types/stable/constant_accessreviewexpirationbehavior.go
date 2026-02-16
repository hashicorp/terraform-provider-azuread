package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewExpirationBehavior string

const (
	AccessReviewExpirationBehavior_AcceptAccessRecommendation AccessReviewExpirationBehavior = "acceptAccessRecommendation"
	AccessReviewExpirationBehavior_KeepAccess                 AccessReviewExpirationBehavior = "keepAccess"
	AccessReviewExpirationBehavior_RemoveAccess               AccessReviewExpirationBehavior = "removeAccess"
)

func PossibleValuesForAccessReviewExpirationBehavior() []string {
	return []string{
		string(AccessReviewExpirationBehavior_AcceptAccessRecommendation),
		string(AccessReviewExpirationBehavior_KeepAccess),
		string(AccessReviewExpirationBehavior_RemoveAccess),
	}
}

func (s *AccessReviewExpirationBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewExpirationBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewExpirationBehavior(input string) (*AccessReviewExpirationBehavior, error) {
	vals := map[string]AccessReviewExpirationBehavior{
		"acceptaccessrecommendation": AccessReviewExpirationBehavior_AcceptAccessRecommendation,
		"keepaccess":                 AccessReviewExpirationBehavior_KeepAccess,
		"removeaccess":               AccessReviewExpirationBehavior_RemoveAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewExpirationBehavior(input)
	return &out, nil
}
