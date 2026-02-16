package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewHistoryDecisionFilter string

const (
	AccessReviewHistoryDecisionFilter_Approve     AccessReviewHistoryDecisionFilter = "approve"
	AccessReviewHistoryDecisionFilter_Deny        AccessReviewHistoryDecisionFilter = "deny"
	AccessReviewHistoryDecisionFilter_DontKnow    AccessReviewHistoryDecisionFilter = "dontKnow"
	AccessReviewHistoryDecisionFilter_NotNotified AccessReviewHistoryDecisionFilter = "notNotified"
	AccessReviewHistoryDecisionFilter_NotReviewed AccessReviewHistoryDecisionFilter = "notReviewed"
)

func PossibleValuesForAccessReviewHistoryDecisionFilter() []string {
	return []string{
		string(AccessReviewHistoryDecisionFilter_Approve),
		string(AccessReviewHistoryDecisionFilter_Deny),
		string(AccessReviewHistoryDecisionFilter_DontKnow),
		string(AccessReviewHistoryDecisionFilter_NotNotified),
		string(AccessReviewHistoryDecisionFilter_NotReviewed),
	}
}

func (s *AccessReviewHistoryDecisionFilter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessReviewHistoryDecisionFilter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessReviewHistoryDecisionFilter(input string) (*AccessReviewHistoryDecisionFilter, error) {
	vals := map[string]AccessReviewHistoryDecisionFilter{
		"approve":     AccessReviewHistoryDecisionFilter_Approve,
		"deny":        AccessReviewHistoryDecisionFilter_Deny,
		"dontknow":    AccessReviewHistoryDecisionFilter_DontKnow,
		"notnotified": AccessReviewHistoryDecisionFilter_NotNotified,
		"notreviewed": AccessReviewHistoryDecisionFilter_NotReviewed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessReviewHistoryDecisionFilter(input)
	return &out, nil
}
