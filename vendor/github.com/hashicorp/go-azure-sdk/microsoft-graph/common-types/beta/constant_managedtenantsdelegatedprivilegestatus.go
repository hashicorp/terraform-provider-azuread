package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDelegatedPrivilegeStatus string

const (
	ManagedTenantsDelegatedPrivilegeStatus_DelegatedAdminPrivileges                     ManagedTenantsDelegatedPrivilegeStatus = "delegatedAdminPrivileges"
	ManagedTenantsDelegatedPrivilegeStatus_DelegatedAndGranularDelegetedAdminPrivileges ManagedTenantsDelegatedPrivilegeStatus = "delegatedAndGranularDelegetedAdminPrivileges"
	ManagedTenantsDelegatedPrivilegeStatus_GranularDelegatedAdminPrivileges             ManagedTenantsDelegatedPrivilegeStatus = "granularDelegatedAdminPrivileges"
	ManagedTenantsDelegatedPrivilegeStatus_None                                         ManagedTenantsDelegatedPrivilegeStatus = "none"
)

func PossibleValuesForManagedTenantsDelegatedPrivilegeStatus() []string {
	return []string{
		string(ManagedTenantsDelegatedPrivilegeStatus_DelegatedAdminPrivileges),
		string(ManagedTenantsDelegatedPrivilegeStatus_DelegatedAndGranularDelegetedAdminPrivileges),
		string(ManagedTenantsDelegatedPrivilegeStatus_GranularDelegatedAdminPrivileges),
		string(ManagedTenantsDelegatedPrivilegeStatus_None),
	}
}

func (s *ManagedTenantsDelegatedPrivilegeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsDelegatedPrivilegeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsDelegatedPrivilegeStatus(input string) (*ManagedTenantsDelegatedPrivilegeStatus, error) {
	vals := map[string]ManagedTenantsDelegatedPrivilegeStatus{
		"delegatedadminprivileges":                     ManagedTenantsDelegatedPrivilegeStatus_DelegatedAdminPrivileges,
		"delegatedandgranulardelegetedadminprivileges": ManagedTenantsDelegatedPrivilegeStatus_DelegatedAndGranularDelegetedAdminPrivileges,
		"granulardelegatedadminprivileges":             ManagedTenantsDelegatedPrivilegeStatus_GranularDelegatedAdminPrivileges,
		"none":                                         ManagedTenantsDelegatedPrivilegeStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsDelegatedPrivilegeStatus(input)
	return &out, nil
}
