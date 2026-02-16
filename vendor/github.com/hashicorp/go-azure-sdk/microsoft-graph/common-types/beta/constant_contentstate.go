package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentState string

const (
	ContentState_Motion ContentState = "motion"
	ContentState_Rest   ContentState = "rest"
	ContentState_Use    ContentState = "use"
)

func PossibleValuesForContentState() []string {
	return []string{
		string(ContentState_Motion),
		string(ContentState_Rest),
		string(ContentState_Use),
	}
}

func (s *ContentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseContentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseContentState(input string) (*ContentState, error) {
	vals := map[string]ContentState{
		"motion": ContentState_Motion,
		"rest":   ContentState_Rest,
		"use":    ContentState_Use,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentState(input)
	return &out, nil
}
