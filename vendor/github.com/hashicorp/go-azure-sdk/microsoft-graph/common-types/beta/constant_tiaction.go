package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TiAction string

const (
	TiAction_Alert   TiAction = "alert"
	TiAction_Allow   TiAction = "allow"
	TiAction_Block   TiAction = "block"
	TiAction_Unknown TiAction = "unknown"
)

func PossibleValuesForTiAction() []string {
	return []string{
		string(TiAction_Alert),
		string(TiAction_Allow),
		string(TiAction_Block),
		string(TiAction_Unknown),
	}
}

func (s *TiAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTiAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTiAction(input string) (*TiAction, error) {
	vals := map[string]TiAction{
		"alert":   TiAction_Alert,
		"allow":   TiAction_Allow,
		"block":   TiAction_Block,
		"unknown": TiAction_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TiAction(input)
	return &out, nil
}
