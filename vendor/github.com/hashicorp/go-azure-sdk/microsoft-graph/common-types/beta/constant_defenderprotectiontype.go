package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderProtectionType string

const (
	DefenderProtectionType_AuditMode     DefenderProtectionType = "auditMode"
	DefenderProtectionType_Enable        DefenderProtectionType = "enable"
	DefenderProtectionType_NotConfigured DefenderProtectionType = "notConfigured"
	DefenderProtectionType_UserDefined   DefenderProtectionType = "userDefined"
	DefenderProtectionType_Warn          DefenderProtectionType = "warn"
)

func PossibleValuesForDefenderProtectionType() []string {
	return []string{
		string(DefenderProtectionType_AuditMode),
		string(DefenderProtectionType_Enable),
		string(DefenderProtectionType_NotConfigured),
		string(DefenderProtectionType_UserDefined),
		string(DefenderProtectionType_Warn),
	}
}

func (s *DefenderProtectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderProtectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderProtectionType(input string) (*DefenderProtectionType, error) {
	vals := map[string]DefenderProtectionType{
		"auditmode":     DefenderProtectionType_AuditMode,
		"enable":        DefenderProtectionType_Enable,
		"notconfigured": DefenderProtectionType_NotConfigured,
		"userdefined":   DefenderProtectionType_UserDefined,
		"warn":          DefenderProtectionType_Warn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderProtectionType(input)
	return &out, nil
}
