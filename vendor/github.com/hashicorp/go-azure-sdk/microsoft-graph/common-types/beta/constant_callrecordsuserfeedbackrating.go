package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsUserFeedbackRating string

const (
	CallRecordsUserFeedbackRating_Bad       CallRecordsUserFeedbackRating = "bad"
	CallRecordsUserFeedbackRating_Excellent CallRecordsUserFeedbackRating = "excellent"
	CallRecordsUserFeedbackRating_Fair      CallRecordsUserFeedbackRating = "fair"
	CallRecordsUserFeedbackRating_Good      CallRecordsUserFeedbackRating = "good"
	CallRecordsUserFeedbackRating_NotRated  CallRecordsUserFeedbackRating = "notRated"
	CallRecordsUserFeedbackRating_Poor      CallRecordsUserFeedbackRating = "poor"
)

func PossibleValuesForCallRecordsUserFeedbackRating() []string {
	return []string{
		string(CallRecordsUserFeedbackRating_Bad),
		string(CallRecordsUserFeedbackRating_Excellent),
		string(CallRecordsUserFeedbackRating_Fair),
		string(CallRecordsUserFeedbackRating_Good),
		string(CallRecordsUserFeedbackRating_NotRated),
		string(CallRecordsUserFeedbackRating_Poor),
	}
}

func (s *CallRecordsUserFeedbackRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsUserFeedbackRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsUserFeedbackRating(input string) (*CallRecordsUserFeedbackRating, error) {
	vals := map[string]CallRecordsUserFeedbackRating{
		"bad":       CallRecordsUserFeedbackRating_Bad,
		"excellent": CallRecordsUserFeedbackRating_Excellent,
		"fair":      CallRecordsUserFeedbackRating_Fair,
		"good":      CallRecordsUserFeedbackRating_Good,
		"notrated":  CallRecordsUserFeedbackRating_NotRated,
		"poor":      CallRecordsUserFeedbackRating_Poor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsUserFeedbackRating(input)
	return &out, nil
}
