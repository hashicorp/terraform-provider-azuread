package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySourceType string

const (
	SecuritySourceType_Mailbox SecuritySourceType = "mailbox"
	SecuritySourceType_Site    SecuritySourceType = "site"
)

func PossibleValuesForSecuritySourceType() []string {
	return []string{
		string(SecuritySourceType_Mailbox),
		string(SecuritySourceType_Site),
	}
}

func (s *SecuritySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySourceType(input string) (*SecuritySourceType, error) {
	vals := map[string]SecuritySourceType{
		"mailbox": SecuritySourceType_Mailbox,
		"site":    SecuritySourceType_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySourceType(input)
	return &out, nil
}
