package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantOnboardingStatus string

const (
	ManagedTenantsTenantOnboardingStatus_Active     ManagedTenantsTenantOnboardingStatus = "active"
	ManagedTenantsTenantOnboardingStatus_Disabled   ManagedTenantsTenantOnboardingStatus = "disabled"
	ManagedTenantsTenantOnboardingStatus_InProcess  ManagedTenantsTenantOnboardingStatus = "inProcess"
	ManagedTenantsTenantOnboardingStatus_Inactive   ManagedTenantsTenantOnboardingStatus = "inactive"
	ManagedTenantsTenantOnboardingStatus_Ineligible ManagedTenantsTenantOnboardingStatus = "ineligible"
)

func PossibleValuesForManagedTenantsTenantOnboardingStatus() []string {
	return []string{
		string(ManagedTenantsTenantOnboardingStatus_Active),
		string(ManagedTenantsTenantOnboardingStatus_Disabled),
		string(ManagedTenantsTenantOnboardingStatus_InProcess),
		string(ManagedTenantsTenantOnboardingStatus_Inactive),
		string(ManagedTenantsTenantOnboardingStatus_Ineligible),
	}
}

func (s *ManagedTenantsTenantOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsTenantOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsTenantOnboardingStatus(input string) (*ManagedTenantsTenantOnboardingStatus, error) {
	vals := map[string]ManagedTenantsTenantOnboardingStatus{
		"active":     ManagedTenantsTenantOnboardingStatus_Active,
		"disabled":   ManagedTenantsTenantOnboardingStatus_Disabled,
		"inprocess":  ManagedTenantsTenantOnboardingStatus_InProcess,
		"inactive":   ManagedTenantsTenantOnboardingStatus_Inactive,
		"ineligible": ManagedTenantsTenantOnboardingStatus_Ineligible,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantOnboardingStatus(input)
	return &out, nil
}
