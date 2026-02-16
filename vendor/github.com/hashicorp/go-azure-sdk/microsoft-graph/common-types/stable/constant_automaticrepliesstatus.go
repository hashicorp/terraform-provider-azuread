package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticRepliesStatus string

const (
	AutomaticRepliesStatus_AlwaysEnabled AutomaticRepliesStatus = "alwaysEnabled"
	AutomaticRepliesStatus_Disabled      AutomaticRepliesStatus = "disabled"
	AutomaticRepliesStatus_Scheduled     AutomaticRepliesStatus = "scheduled"
)

func PossibleValuesForAutomaticRepliesStatus() []string {
	return []string{
		string(AutomaticRepliesStatus_AlwaysEnabled),
		string(AutomaticRepliesStatus_Disabled),
		string(AutomaticRepliesStatus_Scheduled),
	}
}

func (s *AutomaticRepliesStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticRepliesStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticRepliesStatus(input string) (*AutomaticRepliesStatus, error) {
	vals := map[string]AutomaticRepliesStatus{
		"alwaysenabled": AutomaticRepliesStatus_AlwaysEnabled,
		"disabled":      AutomaticRepliesStatus_Disabled,
		"scheduled":     AutomaticRepliesStatus_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticRepliesStatus(input)
	return &out, nil
}
