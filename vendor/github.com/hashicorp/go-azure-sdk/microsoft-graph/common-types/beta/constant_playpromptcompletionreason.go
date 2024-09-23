package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlayPromptCompletionReason string

const (
	PlayPromptCompletionReason_CompletedSuccessfully  PlayPromptCompletionReason = "completedSuccessfully"
	PlayPromptCompletionReason_MediaOperationCanceled PlayPromptCompletionReason = "mediaOperationCanceled"
	PlayPromptCompletionReason_Unknown                PlayPromptCompletionReason = "unknown"
)

func PossibleValuesForPlayPromptCompletionReason() []string {
	return []string{
		string(PlayPromptCompletionReason_CompletedSuccessfully),
		string(PlayPromptCompletionReason_MediaOperationCanceled),
		string(PlayPromptCompletionReason_Unknown),
	}
}

func (s *PlayPromptCompletionReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlayPromptCompletionReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlayPromptCompletionReason(input string) (*PlayPromptCompletionReason, error) {
	vals := map[string]PlayPromptCompletionReason{
		"completedsuccessfully":  PlayPromptCompletionReason_CompletedSuccessfully,
		"mediaoperationcanceled": PlayPromptCompletionReason_MediaOperationCanceled,
		"unknown":                PlayPromptCompletionReason_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlayPromptCompletionReason(input)
	return &out, nil
}
