package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderPotentiallyUnwantedAppAction string

const (
	DefenderPotentiallyUnwantedAppAction_Audit         DefenderPotentiallyUnwantedAppAction = "audit"
	DefenderPotentiallyUnwantedAppAction_Block         DefenderPotentiallyUnwantedAppAction = "block"
	DefenderPotentiallyUnwantedAppAction_DeviceDefault DefenderPotentiallyUnwantedAppAction = "deviceDefault"
)

func PossibleValuesForDefenderPotentiallyUnwantedAppAction() []string {
	return []string{
		string(DefenderPotentiallyUnwantedAppAction_Audit),
		string(DefenderPotentiallyUnwantedAppAction_Block),
		string(DefenderPotentiallyUnwantedAppAction_DeviceDefault),
	}
}

func (s *DefenderPotentiallyUnwantedAppAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderPotentiallyUnwantedAppAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderPotentiallyUnwantedAppAction(input string) (*DefenderPotentiallyUnwantedAppAction, error) {
	vals := map[string]DefenderPotentiallyUnwantedAppAction{
		"audit":         DefenderPotentiallyUnwantedAppAction_Audit,
		"block":         DefenderPotentiallyUnwantedAppAction_Block,
		"devicedefault": DefenderPotentiallyUnwantedAppAction_DeviceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderPotentiallyUnwantedAppAction(input)
	return &out, nil
}
