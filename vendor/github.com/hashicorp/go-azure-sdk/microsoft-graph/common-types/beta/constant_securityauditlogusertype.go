package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAuditLogUserType string

const (
	SecurityAuditLogUserType_Admin             SecurityAuditLogUserType = "Admin"
	SecurityAuditLogUserType_Application       SecurityAuditLogUserType = "Application"
	SecurityAuditLogUserType_CustomPolicy      SecurityAuditLogUserType = "CustomPolicy"
	SecurityAuditLogUserType_DcAdmin           SecurityAuditLogUserType = "DcAdmin"
	SecurityAuditLogUserType_Guest             SecurityAuditLogUserType = "Guest"
	SecurityAuditLogUserType_PartnerTechnician SecurityAuditLogUserType = "PartnerTechnician"
	SecurityAuditLogUserType_Regular           SecurityAuditLogUserType = "Regular"
	SecurityAuditLogUserType_Reserved          SecurityAuditLogUserType = "Reserved"
	SecurityAuditLogUserType_ServicePrincipal  SecurityAuditLogUserType = "ServicePrincipal"
	SecurityAuditLogUserType_System            SecurityAuditLogUserType = "System"
	SecurityAuditLogUserType_SystemPolicy      SecurityAuditLogUserType = "SystemPolicy"
)

func PossibleValuesForSecurityAuditLogUserType() []string {
	return []string{
		string(SecurityAuditLogUserType_Admin),
		string(SecurityAuditLogUserType_Application),
		string(SecurityAuditLogUserType_CustomPolicy),
		string(SecurityAuditLogUserType_DcAdmin),
		string(SecurityAuditLogUserType_Guest),
		string(SecurityAuditLogUserType_PartnerTechnician),
		string(SecurityAuditLogUserType_Regular),
		string(SecurityAuditLogUserType_Reserved),
		string(SecurityAuditLogUserType_ServicePrincipal),
		string(SecurityAuditLogUserType_System),
		string(SecurityAuditLogUserType_SystemPolicy),
	}
}

func (s *SecurityAuditLogUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAuditLogUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAuditLogUserType(input string) (*SecurityAuditLogUserType, error) {
	vals := map[string]SecurityAuditLogUserType{
		"admin":             SecurityAuditLogUserType_Admin,
		"application":       SecurityAuditLogUserType_Application,
		"custompolicy":      SecurityAuditLogUserType_CustomPolicy,
		"dcadmin":           SecurityAuditLogUserType_DcAdmin,
		"guest":             SecurityAuditLogUserType_Guest,
		"partnertechnician": SecurityAuditLogUserType_PartnerTechnician,
		"regular":           SecurityAuditLogUserType_Regular,
		"reserved":          SecurityAuditLogUserType_Reserved,
		"serviceprincipal":  SecurityAuditLogUserType_ServicePrincipal,
		"system":            SecurityAuditLogUserType_System,
		"systempolicy":      SecurityAuditLogUserType_SystemPolicy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAuditLogUserType(input)
	return &out, nil
}
