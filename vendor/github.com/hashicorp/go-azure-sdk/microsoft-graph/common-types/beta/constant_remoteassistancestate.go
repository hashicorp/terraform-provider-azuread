package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistanceState string

const (
	RemoteAssistanceState_Disabled RemoteAssistanceState = "disabled"
	RemoteAssistanceState_Enabled  RemoteAssistanceState = "enabled"
)

func PossibleValuesForRemoteAssistanceState() []string {
	return []string{
		string(RemoteAssistanceState_Disabled),
		string(RemoteAssistanceState_Enabled),
	}
}

func (s *RemoteAssistanceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteAssistanceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteAssistanceState(input string) (*RemoteAssistanceState, error) {
	vals := map[string]RemoteAssistanceState{
		"disabled": RemoteAssistanceState_Disabled,
		"enabled":  RemoteAssistanceState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteAssistanceState(input)
	return &out, nil
}
