package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPurgeType string

const (
	SecurityPurgeType_PermanentlyDeleted SecurityPurgeType = "permanentlyDeleted"
	SecurityPurgeType_Recoverable        SecurityPurgeType = "recoverable"
)

func PossibleValuesForSecurityPurgeType() []string {
	return []string{
		string(SecurityPurgeType_PermanentlyDeleted),
		string(SecurityPurgeType_Recoverable),
	}
}

func (s *SecurityPurgeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityPurgeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityPurgeType(input string) (*SecurityPurgeType, error) {
	vals := map[string]SecurityPurgeType{
		"permanentlydeleted": SecurityPurgeType_PermanentlyDeleted,
		"recoverable":        SecurityPurgeType_Recoverable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityPurgeType(input)
	return &out, nil
}
