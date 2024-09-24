package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionState string

const (
	ActionState_Active       ActionState = "active"
	ActionState_Canceled     ActionState = "canceled"
	ActionState_Done         ActionState = "done"
	ActionState_Failed       ActionState = "failed"
	ActionState_None         ActionState = "none"
	ActionState_NotSupported ActionState = "notSupported"
	ActionState_Pending      ActionState = "pending"
)

func PossibleValuesForActionState() []string {
	return []string{
		string(ActionState_Active),
		string(ActionState_Canceled),
		string(ActionState_Done),
		string(ActionState_Failed),
		string(ActionState_None),
		string(ActionState_NotSupported),
		string(ActionState_Pending),
	}
}

func (s *ActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActionState(input string) (*ActionState, error) {
	vals := map[string]ActionState{
		"active":       ActionState_Active,
		"canceled":     ActionState_Canceled,
		"done":         ActionState_Done,
		"failed":       ActionState_Failed,
		"none":         ActionState_None,
		"notsupported": ActionState_NotSupported,
		"pending":      ActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionState(input)
	return &out, nil
}
