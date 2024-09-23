package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAuditLogQueryStatus string

const (
	SecurityAuditLogQueryStatus_Cancelled  SecurityAuditLogQueryStatus = "cancelled"
	SecurityAuditLogQueryStatus_Failed     SecurityAuditLogQueryStatus = "failed"
	SecurityAuditLogQueryStatus_NotStarted SecurityAuditLogQueryStatus = "notStarted"
	SecurityAuditLogQueryStatus_Running    SecurityAuditLogQueryStatus = "running"
	SecurityAuditLogQueryStatus_Succeeded  SecurityAuditLogQueryStatus = "succeeded"
)

func PossibleValuesForSecurityAuditLogQueryStatus() []string {
	return []string{
		string(SecurityAuditLogQueryStatus_Cancelled),
		string(SecurityAuditLogQueryStatus_Failed),
		string(SecurityAuditLogQueryStatus_NotStarted),
		string(SecurityAuditLogQueryStatus_Running),
		string(SecurityAuditLogQueryStatus_Succeeded),
	}
}

func (s *SecurityAuditLogQueryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAuditLogQueryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAuditLogQueryStatus(input string) (*SecurityAuditLogQueryStatus, error) {
	vals := map[string]SecurityAuditLogQueryStatus{
		"cancelled":  SecurityAuditLogQueryStatus_Cancelled,
		"failed":     SecurityAuditLogQueryStatus_Failed,
		"notstarted": SecurityAuditLogQueryStatus_NotStarted,
		"running":    SecurityAuditLogQueryStatus_Running,
		"succeeded":  SecurityAuditLogQueryStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAuditLogQueryStatus(input)
	return &out, nil
}
