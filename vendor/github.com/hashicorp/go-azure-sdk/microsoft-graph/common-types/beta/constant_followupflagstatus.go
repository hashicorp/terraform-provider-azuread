package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowupFlagStatus string

const (
	FollowupFlagStatus_Complete   FollowupFlagStatus = "complete"
	FollowupFlagStatus_Flagged    FollowupFlagStatus = "flagged"
	FollowupFlagStatus_NotFlagged FollowupFlagStatus = "notFlagged"
)

func PossibleValuesForFollowupFlagStatus() []string {
	return []string{
		string(FollowupFlagStatus_Complete),
		string(FollowupFlagStatus_Flagged),
		string(FollowupFlagStatus_NotFlagged),
	}
}

func (s *FollowupFlagStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFollowupFlagStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFollowupFlagStatus(input string) (*FollowupFlagStatus, error) {
	vals := map[string]FollowupFlagStatus{
		"complete":   FollowupFlagStatus_Complete,
		"flagged":    FollowupFlagStatus_Flagged,
		"notflagged": FollowupFlagStatus_NotFlagged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FollowupFlagStatus(input)
	return &out, nil
}
