package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppRemediationAction string

const (
	ManagedAppRemediationAction_Block                       ManagedAppRemediationAction = "block"
	ManagedAppRemediationAction_BlockWhenSettingIsSupported ManagedAppRemediationAction = "blockWhenSettingIsSupported"
	ManagedAppRemediationAction_Warn                        ManagedAppRemediationAction = "warn"
	ManagedAppRemediationAction_Wipe                        ManagedAppRemediationAction = "wipe"
)

func PossibleValuesForManagedAppRemediationAction() []string {
	return []string{
		string(ManagedAppRemediationAction_Block),
		string(ManagedAppRemediationAction_BlockWhenSettingIsSupported),
		string(ManagedAppRemediationAction_Warn),
		string(ManagedAppRemediationAction_Wipe),
	}
}

func (s *ManagedAppRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppRemediationAction(input string) (*ManagedAppRemediationAction, error) {
	vals := map[string]ManagedAppRemediationAction{
		"block":                       ManagedAppRemediationAction_Block,
		"blockwhensettingissupported": ManagedAppRemediationAction_BlockWhenSettingIsSupported,
		"warn":                        ManagedAppRemediationAction_Warn,
		"wipe":                        ManagedAppRemediationAction_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppRemediationAction(input)
	return &out, nil
}
