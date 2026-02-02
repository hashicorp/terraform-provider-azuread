package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppPublishingState string

const (
	TeamsAppPublishingState_Published TeamsAppPublishingState = "published"
	TeamsAppPublishingState_Rejected  TeamsAppPublishingState = "rejected"
	TeamsAppPublishingState_Submitted TeamsAppPublishingState = "submitted"
)

func PossibleValuesForTeamsAppPublishingState() []string {
	return []string{
		string(TeamsAppPublishingState_Published),
		string(TeamsAppPublishingState_Rejected),
		string(TeamsAppPublishingState_Submitted),
	}
}

func (s *TeamsAppPublishingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppPublishingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppPublishingState(input string) (*TeamsAppPublishingState, error) {
	vals := map[string]TeamsAppPublishingState{
		"published": TeamsAppPublishingState_Published,
		"rejected":  TeamsAppPublishingState_Rejected,
		"submitted": TeamsAppPublishingState_Submitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppPublishingState(input)
	return &out, nil
}
