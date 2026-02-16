package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemFilterByCurrentUserOptions string

const (
	AccessReviewInstanceDecisionItemFilterByCurrentUserOptions_Reviewer AccessReviewInstanceDecisionItemFilterByCurrentUserOptions = "reviewer"
)

func PossibleValuesForAccessReviewInstanceDecisionItemFilterByCurrentUserOptions() []string {
	return []string{
		string(AccessReviewInstanceDecisionItemFilterByCurrentUserOptions_Reviewer),
	}
}

func (s *AccessReviewInstanceDecisionItemFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewInstanceDecisionItemFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewInstanceDecisionItemFilterByCurrentUserOptions(input string) (*AccessReviewInstanceDecisionItemFilterByCurrentUserOptions, error) {
	vals := map[string]AccessReviewInstanceDecisionItemFilterByCurrentUserOptions{
		"reviewer": AccessReviewInstanceDecisionItemFilterByCurrentUserOptions_Reviewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewInstanceDecisionItemFilterByCurrentUserOptions(input)
	return &out, nil
}
