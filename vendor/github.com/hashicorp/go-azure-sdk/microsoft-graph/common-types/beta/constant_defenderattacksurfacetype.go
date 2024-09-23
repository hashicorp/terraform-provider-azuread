package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderAttackSurfaceType string

const (
	DefenderAttackSurfaceType_AuditMode   DefenderAttackSurfaceType = "auditMode"
	DefenderAttackSurfaceType_Block       DefenderAttackSurfaceType = "block"
	DefenderAttackSurfaceType_Disable     DefenderAttackSurfaceType = "disable"
	DefenderAttackSurfaceType_UserDefined DefenderAttackSurfaceType = "userDefined"
	DefenderAttackSurfaceType_Warn        DefenderAttackSurfaceType = "warn"
)

func PossibleValuesForDefenderAttackSurfaceType() []string {
	return []string{
		string(DefenderAttackSurfaceType_AuditMode),
		string(DefenderAttackSurfaceType_Block),
		string(DefenderAttackSurfaceType_Disable),
		string(DefenderAttackSurfaceType_UserDefined),
		string(DefenderAttackSurfaceType_Warn),
	}
}

func (s *DefenderAttackSurfaceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderAttackSurfaceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderAttackSurfaceType(input string) (*DefenderAttackSurfaceType, error) {
	vals := map[string]DefenderAttackSurfaceType{
		"auditmode":   DefenderAttackSurfaceType_AuditMode,
		"block":       DefenderAttackSurfaceType_Block,
		"disable":     DefenderAttackSurfaceType_Disable,
		"userdefined": DefenderAttackSurfaceType_UserDefined,
		"warn":        DefenderAttackSurfaceType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderAttackSurfaceType(input)
	return &out, nil
}
