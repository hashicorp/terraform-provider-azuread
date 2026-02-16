package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewStageFilterByCurrentUserOptions string

const (
	AccessReviewStageFilterByCurrentUserOptions_Reviewer AccessReviewStageFilterByCurrentUserOptions = "reviewer"
)

func PossibleValuesForAccessReviewStageFilterByCurrentUserOptions() []string {
	return []string{
		string(AccessReviewStageFilterByCurrentUserOptions_Reviewer),
	}
}

func (s *AccessReviewStageFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewStageFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewStageFilterByCurrentUserOptions(input string) (*AccessReviewStageFilterByCurrentUserOptions, error) {
	vals := map[string]AccessReviewStageFilterByCurrentUserOptions{
		"reviewer": AccessReviewStageFilterByCurrentUserOptions_Reviewer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewStageFilterByCurrentUserOptions(input)
	return &out, nil
}
