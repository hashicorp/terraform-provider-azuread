package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIncidentStatus string

const (
	SecurityIncidentStatus_Active         SecurityIncidentStatus = "active"
	SecurityIncidentStatus_AwaitingAction SecurityIncidentStatus = "awaitingAction"
	SecurityIncidentStatus_InProgress     SecurityIncidentStatus = "inProgress"
	SecurityIncidentStatus_Redirected     SecurityIncidentStatus = "redirected"
	SecurityIncidentStatus_Resolved       SecurityIncidentStatus = "resolved"
)

func PossibleValuesForSecurityIncidentStatus() []string {
	return []string{
		string(SecurityIncidentStatus_Active),
		string(SecurityIncidentStatus_AwaitingAction),
		string(SecurityIncidentStatus_InProgress),
		string(SecurityIncidentStatus_Redirected),
		string(SecurityIncidentStatus_Resolved),
	}
}

func (s *SecurityIncidentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIncidentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIncidentStatus(input string) (*SecurityIncidentStatus, error) {
	vals := map[string]SecurityIncidentStatus{
		"active":         SecurityIncidentStatus_Active,
		"awaitingaction": SecurityIncidentStatus_AwaitingAction,
		"inprogress":     SecurityIncidentStatus_InProgress,
		"redirected":     SecurityIncidentStatus_Redirected,
		"resolved":       SecurityIncidentStatus_Resolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIncidentStatus(input)
	return &out, nil
}
