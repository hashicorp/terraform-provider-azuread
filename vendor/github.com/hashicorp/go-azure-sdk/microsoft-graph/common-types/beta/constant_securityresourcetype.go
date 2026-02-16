package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResourceType string

const (
	SecurityResourceType_Attacked SecurityResourceType = "attacked"
	SecurityResourceType_Related  SecurityResourceType = "related"
	SecurityResourceType_Unknown  SecurityResourceType = "unknown"
)

func PossibleValuesForSecurityResourceType() []string {
	return []string{
		string(SecurityResourceType_Attacked),
		string(SecurityResourceType_Related),
		string(SecurityResourceType_Unknown),
	}
}

func (s *SecurityResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityResourceType(input string) (*SecurityResourceType, error) {
	vals := map[string]SecurityResourceType{
		"attacked": SecurityResourceType_Attacked,
		"related":  SecurityResourceType_Related,
		"unknown":  SecurityResourceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityResourceType(input)
	return &out, nil
}
