package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberState string

const (
	MultiTenantOrganizationMemberState_Active  MultiTenantOrganizationMemberState = "active"
	MultiTenantOrganizationMemberState_Pending MultiTenantOrganizationMemberState = "pending"
	MultiTenantOrganizationMemberState_Removed MultiTenantOrganizationMemberState = "removed"
)

func PossibleValuesForMultiTenantOrganizationMemberState() []string {
	return []string{
		string(MultiTenantOrganizationMemberState_Active),
		string(MultiTenantOrganizationMemberState_Pending),
		string(MultiTenantOrganizationMemberState_Removed),
	}
}

func (s *MultiTenantOrganizationMemberState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationMemberState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationMemberState(input string) (*MultiTenantOrganizationMemberState, error) {
	vals := map[string]MultiTenantOrganizationMemberState{
		"active":  MultiTenantOrganizationMemberState_Active,
		"pending": MultiTenantOrganizationMemberState_Pending,
		"removed": MultiTenantOrganizationMemberState_Removed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberState(input)
	return &out, nil
}
