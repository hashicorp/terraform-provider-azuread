package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventStatusType string

const (
	SecurityEventStatusType_Error        SecurityEventStatusType = "error"
	SecurityEventStatusType_NotAvaliable SecurityEventStatusType = "notAvaliable"
	SecurityEventStatusType_Pending      SecurityEventStatusType = "pending"
	SecurityEventStatusType_Success      SecurityEventStatusType = "success"
)

func PossibleValuesForSecurityEventStatusType() []string {
	return []string{
		string(SecurityEventStatusType_Error),
		string(SecurityEventStatusType_NotAvaliable),
		string(SecurityEventStatusType_Pending),
		string(SecurityEventStatusType_Success),
	}
}

func (s *SecurityEventStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEventStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEventStatusType(input string) (*SecurityEventStatusType, error) {
	vals := map[string]SecurityEventStatusType{
		"error":        SecurityEventStatusType_Error,
		"notavaliable": SecurityEventStatusType_NotAvaliable,
		"pending":      SecurityEventStatusType_Pending,
		"success":      SecurityEventStatusType_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventStatusType(input)
	return &out, nil
}
