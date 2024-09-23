package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPurgeAreas string

const (
	SecurityPurgeAreas_Mailboxes     SecurityPurgeAreas = "mailboxes"
	SecurityPurgeAreas_TeamsMessages SecurityPurgeAreas = "teamsMessages"
)

func PossibleValuesForSecurityPurgeAreas() []string {
	return []string{
		string(SecurityPurgeAreas_Mailboxes),
		string(SecurityPurgeAreas_TeamsMessages),
	}
}

func (s *SecurityPurgeAreas) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityPurgeAreas(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityPurgeAreas(input string) (*SecurityPurgeAreas, error) {
	vals := map[string]SecurityPurgeAreas{
		"mailboxes":     SecurityPurgeAreas_Mailboxes,
		"teamsmessages": SecurityPurgeAreas_TeamsMessages,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityPurgeAreas(input)
	return &out, nil
}
