package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerActionType string

const (
	PowerActionType_Hibernate     PowerActionType = "hibernate"
	PowerActionType_NoAction      PowerActionType = "noAction"
	PowerActionType_NotConfigured PowerActionType = "notConfigured"
	PowerActionType_Shutdown      PowerActionType = "shutdown"
	PowerActionType_Sleep         PowerActionType = "sleep"
)

func PossibleValuesForPowerActionType() []string {
	return []string{
		string(PowerActionType_Hibernate),
		string(PowerActionType_NoAction),
		string(PowerActionType_NotConfigured),
		string(PowerActionType_Shutdown),
		string(PowerActionType_Sleep),
	}
}

func (s *PowerActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePowerActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePowerActionType(input string) (*PowerActionType, error) {
	vals := map[string]PowerActionType{
		"hibernate":     PowerActionType_Hibernate,
		"noaction":      PowerActionType_NoAction,
		"notconfigured": PowerActionType_NotConfigured,
		"shutdown":      PowerActionType_Shutdown,
		"sleep":         PowerActionType_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PowerActionType(input)
	return &out, nil
}
