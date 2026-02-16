package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpRoleType string

const (
	GcpRoleType_Custom GcpRoleType = "custom"
	GcpRoleType_System GcpRoleType = "system"
)

func PossibleValuesForGcpRoleType() []string {
	return []string{
		string(GcpRoleType_Custom),
		string(GcpRoleType_System),
	}
}

func (s *GcpRoleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGcpRoleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGcpRoleType(input string) (*GcpRoleType, error) {
	vals := map[string]GcpRoleType{
		"custom": GcpRoleType_Custom,
		"system": GcpRoleType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GcpRoleType(input)
	return &out, nil
}
