package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRecipientType string

const (
	SecurityRecipientType_RoleGroup SecurityRecipientType = "roleGroup"
	SecurityRecipientType_User      SecurityRecipientType = "user"
)

func PossibleValuesForSecurityRecipientType() []string {
	return []string{
		string(SecurityRecipientType_RoleGroup),
		string(SecurityRecipientType_User),
	}
}

func (s *SecurityRecipientType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRecipientType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRecipientType(input string) (*SecurityRecipientType, error) {
	vals := map[string]SecurityRecipientType{
		"rolegroup": SecurityRecipientType_RoleGroup,
		"user":      SecurityRecipientType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRecipientType(input)
	return &out, nil
}
