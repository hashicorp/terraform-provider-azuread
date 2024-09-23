package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantOnboardingEligibilityReason string

const (
	ManagedTenantsTenantOnboardingEligibilityReason_ContractType             ManagedTenantsTenantOnboardingEligibilityReason = "contractType"
	ManagedTenantsTenantOnboardingEligibilityReason_DelegatedAdminPrivileges ManagedTenantsTenantOnboardingEligibilityReason = "delegatedAdminPrivileges"
	ManagedTenantsTenantOnboardingEligibilityReason_License                  ManagedTenantsTenantOnboardingEligibilityReason = "license"
	ManagedTenantsTenantOnboardingEligibilityReason_None                     ManagedTenantsTenantOnboardingEligibilityReason = "none"
	ManagedTenantsTenantOnboardingEligibilityReason_UsersCount               ManagedTenantsTenantOnboardingEligibilityReason = "usersCount"
)

func PossibleValuesForManagedTenantsTenantOnboardingEligibilityReason() []string {
	return []string{
		string(ManagedTenantsTenantOnboardingEligibilityReason_ContractType),
		string(ManagedTenantsTenantOnboardingEligibilityReason_DelegatedAdminPrivileges),
		string(ManagedTenantsTenantOnboardingEligibilityReason_License),
		string(ManagedTenantsTenantOnboardingEligibilityReason_None),
		string(ManagedTenantsTenantOnboardingEligibilityReason_UsersCount),
	}
}

func (s *ManagedTenantsTenantOnboardingEligibilityReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsTenantOnboardingEligibilityReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsTenantOnboardingEligibilityReason(input string) (*ManagedTenantsTenantOnboardingEligibilityReason, error) {
	vals := map[string]ManagedTenantsTenantOnboardingEligibilityReason{
		"contracttype":             ManagedTenantsTenantOnboardingEligibilityReason_ContractType,
		"delegatedadminprivileges": ManagedTenantsTenantOnboardingEligibilityReason_DelegatedAdminPrivileges,
		"license":                  ManagedTenantsTenantOnboardingEligibilityReason_License,
		"none":                     ManagedTenantsTenantOnboardingEligibilityReason_None,
		"userscount":               ManagedTenantsTenantOnboardingEligibilityReason_UsersCount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantOnboardingEligibilityReason(input)
	return &out, nil
}
