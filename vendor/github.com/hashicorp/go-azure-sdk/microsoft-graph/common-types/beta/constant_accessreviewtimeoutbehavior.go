package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewTimeoutBehavior string

const (
	AccessReviewTimeoutBehavior_AcceptAccessRecommendation AccessReviewTimeoutBehavior = "acceptAccessRecommendation"
	AccessReviewTimeoutBehavior_KeepAccess                 AccessReviewTimeoutBehavior = "keepAccess"
	AccessReviewTimeoutBehavior_RemoveAccess               AccessReviewTimeoutBehavior = "removeAccess"
)

func PossibleValuesForAccessReviewTimeoutBehavior() []string {
	return []string{
		string(AccessReviewTimeoutBehavior_AcceptAccessRecommendation),
		string(AccessReviewTimeoutBehavior_KeepAccess),
		string(AccessReviewTimeoutBehavior_RemoveAccess),
	}
}

func (s *AccessReviewTimeoutBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewTimeoutBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewTimeoutBehavior(input string) (*AccessReviewTimeoutBehavior, error) {
	vals := map[string]AccessReviewTimeoutBehavior{
		"acceptaccessrecommendation": AccessReviewTimeoutBehavior_AcceptAccessRecommendation,
		"keepaccess":                 AccessReviewTimeoutBehavior_KeepAccess,
		"removeaccess":               AccessReviewTimeoutBehavior_RemoveAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewTimeoutBehavior(input)
	return &out, nil
}
