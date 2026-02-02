package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRetentionTrigger string

const (
	SecurityRetentionTrigger_DateCreated  SecurityRetentionTrigger = "dateCreated"
	SecurityRetentionTrigger_DateLabeled  SecurityRetentionTrigger = "dateLabeled"
	SecurityRetentionTrigger_DateModified SecurityRetentionTrigger = "dateModified"
	SecurityRetentionTrigger_DateOfEvent  SecurityRetentionTrigger = "dateOfEvent"
)

func PossibleValuesForSecurityRetentionTrigger() []string {
	return []string{
		string(SecurityRetentionTrigger_DateCreated),
		string(SecurityRetentionTrigger_DateLabeled),
		string(SecurityRetentionTrigger_DateModified),
		string(SecurityRetentionTrigger_DateOfEvent),
	}
}

func (s *SecurityRetentionTrigger) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRetentionTrigger(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRetentionTrigger(input string) (*SecurityRetentionTrigger, error) {
	vals := map[string]SecurityRetentionTrigger{
		"datecreated":  SecurityRetentionTrigger_DateCreated,
		"datelabeled":  SecurityRetentionTrigger_DateLabeled,
		"datemodified": SecurityRetentionTrigger_DateModified,
		"dateofevent":  SecurityRetentionTrigger_DateOfEvent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRetentionTrigger(input)
	return &out, nil
}
