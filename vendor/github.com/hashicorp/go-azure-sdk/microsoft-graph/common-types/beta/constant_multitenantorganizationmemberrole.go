package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberRole string

const (
	MultiTenantOrganizationMemberRole_Member MultiTenantOrganizationMemberRole = "member"
	MultiTenantOrganizationMemberRole_Owner  MultiTenantOrganizationMemberRole = "owner"
)

func PossibleValuesForMultiTenantOrganizationMemberRole() []string {
	return []string{
		string(MultiTenantOrganizationMemberRole_Member),
		string(MultiTenantOrganizationMemberRole_Owner),
	}
}

func (s *MultiTenantOrganizationMemberRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationMemberRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationMemberRole(input string) (*MultiTenantOrganizationMemberRole, error) {
	vals := map[string]MultiTenantOrganizationMemberRole{
		"member": MultiTenantOrganizationMemberRole_Member,
		"owner":  MultiTenantOrganizationMemberRole_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberRole(input)
	return &out, nil
}
