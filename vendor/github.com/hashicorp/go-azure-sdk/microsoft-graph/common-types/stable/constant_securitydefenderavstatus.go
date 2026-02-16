package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDefenderAvStatus string

const (
	SecurityDefenderAvStatus_Disabled     SecurityDefenderAvStatus = "disabled"
	SecurityDefenderAvStatus_NotReporting SecurityDefenderAvStatus = "notReporting"
	SecurityDefenderAvStatus_NotSupported SecurityDefenderAvStatus = "notSupported"
	SecurityDefenderAvStatus_NotUpdated   SecurityDefenderAvStatus = "notUpdated"
	SecurityDefenderAvStatus_Unknown      SecurityDefenderAvStatus = "unknown"
	SecurityDefenderAvStatus_Updated      SecurityDefenderAvStatus = "updated"
)

func PossibleValuesForSecurityDefenderAvStatus() []string {
	return []string{
		string(SecurityDefenderAvStatus_Disabled),
		string(SecurityDefenderAvStatus_NotReporting),
		string(SecurityDefenderAvStatus_NotSupported),
		string(SecurityDefenderAvStatus_NotUpdated),
		string(SecurityDefenderAvStatus_Unknown),
		string(SecurityDefenderAvStatus_Updated),
	}
}

func (s *SecurityDefenderAvStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDefenderAvStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDefenderAvStatus(input string) (*SecurityDefenderAvStatus, error) {
	vals := map[string]SecurityDefenderAvStatus{
		"disabled":     SecurityDefenderAvStatus_Disabled,
		"notreporting": SecurityDefenderAvStatus_NotReporting,
		"notsupported": SecurityDefenderAvStatus_NotSupported,
		"notupdated":   SecurityDefenderAvStatus_NotUpdated,
		"unknown":      SecurityDefenderAvStatus_Unknown,
		"updated":      SecurityDefenderAvStatus_Updated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDefenderAvStatus(input)
	return &out, nil
}
