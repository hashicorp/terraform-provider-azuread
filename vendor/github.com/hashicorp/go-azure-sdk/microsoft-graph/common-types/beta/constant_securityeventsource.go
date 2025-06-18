package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventSource string

const (
	SecurityEventSource_Admin  SecurityEventSource = "admin"
	SecurityEventSource_System SecurityEventSource = "system"
	SecurityEventSource_User   SecurityEventSource = "user"
)

func PossibleValuesForSecurityEventSource() []string {
	return []string{
		string(SecurityEventSource_Admin),
		string(SecurityEventSource_System),
		string(SecurityEventSource_User),
	}
}

func (s *SecurityEventSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEventSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEventSource(input string) (*SecurityEventSource, error) {
	vals := map[string]SecurityEventSource{
		"admin":  SecurityEventSource_Admin,
		"system": SecurityEventSource_System,
		"user":   SecurityEventSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventSource(input)
	return &out, nil
}
