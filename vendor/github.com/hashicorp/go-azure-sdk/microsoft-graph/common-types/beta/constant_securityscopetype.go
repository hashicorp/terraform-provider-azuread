package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityScopeType string

const (
	SecurityScopeType_DeviceGroup SecurityScopeType = "deviceGroup"
)

func PossibleValuesForSecurityScopeType() []string {
	return []string{
		string(SecurityScopeType_DeviceGroup),
	}
}

func (s *SecurityScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityScopeType(input string) (*SecurityScopeType, error) {
	vals := map[string]SecurityScopeType{
		"devicegroup": SecurityScopeType_DeviceGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityScopeType(input)
	return &out, nil
}
