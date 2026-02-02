package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceFilterByCurrentUserOptions string

const (
	AccessReviewInstanceFilterByCurrentUserOptions_Reviewer AccessReviewInstanceFilterByCurrentUserOptions = "reviewer"
)

func PossibleValuesForAccessReviewInstanceFilterByCurrentUserOptions() []string {
	return []string{
		string(AccessReviewInstanceFilterByCurrentUserOptions_Reviewer),
	}
}

func (s *AccessReviewInstanceFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewInstanceFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewInstanceFilterByCurrentUserOptions(input string) (*AccessReviewInstanceFilterByCurrentUserOptions, error) {
	vals := map[string]AccessReviewInstanceFilterByCurrentUserOptions{
		"reviewer": AccessReviewInstanceFilterByCurrentUserOptions_Reviewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewInstanceFilterByCurrentUserOptions(input)
	return &out, nil
}
