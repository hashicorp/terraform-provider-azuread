package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderThreatAction string

const (
	DefenderThreatAction_Allow         DefenderThreatAction = "allow"
	DefenderThreatAction_Block         DefenderThreatAction = "block"
	DefenderThreatAction_Clean         DefenderThreatAction = "clean"
	DefenderThreatAction_DeviceDefault DefenderThreatAction = "deviceDefault"
	DefenderThreatAction_Quarantine    DefenderThreatAction = "quarantine"
	DefenderThreatAction_Remove        DefenderThreatAction = "remove"
	DefenderThreatAction_UserDefined   DefenderThreatAction = "userDefined"
)

func PossibleValuesForDefenderThreatAction() []string {
	return []string{
		string(DefenderThreatAction_Allow),
		string(DefenderThreatAction_Block),
		string(DefenderThreatAction_Clean),
		string(DefenderThreatAction_DeviceDefault),
		string(DefenderThreatAction_Quarantine),
		string(DefenderThreatAction_Remove),
		string(DefenderThreatAction_UserDefined),
	}
}

func (s *DefenderThreatAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderThreatAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderThreatAction(input string) (*DefenderThreatAction, error) {
	vals := map[string]DefenderThreatAction{
		"allow":         DefenderThreatAction_Allow,
		"block":         DefenderThreatAction_Block,
		"clean":         DefenderThreatAction_Clean,
		"devicedefault": DefenderThreatAction_DeviceDefault,
		"quarantine":    DefenderThreatAction_Quarantine,
		"remove":        DefenderThreatAction_Remove,
		"userdefined":   DefenderThreatAction_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderThreatAction(input)
	return &out, nil
}
