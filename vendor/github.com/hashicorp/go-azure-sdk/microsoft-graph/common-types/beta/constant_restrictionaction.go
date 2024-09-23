package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictionAction string

const (
	RestrictionAction_Audit RestrictionAction = "audit"
	RestrictionAction_Block RestrictionAction = "block"
	RestrictionAction_Warn  RestrictionAction = "warn"
)

func PossibleValuesForRestrictionAction() []string {
	return []string{
		string(RestrictionAction_Audit),
		string(RestrictionAction_Block),
		string(RestrictionAction_Warn),
	}
}

func (s *RestrictionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestrictionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestrictionAction(input string) (*RestrictionAction, error) {
	vals := map[string]RestrictionAction{
		"audit": RestrictionAction_Audit,
		"block": RestrictionAction_Block,
		"warn":  RestrictionAction_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictionAction(input)
	return &out, nil
}
