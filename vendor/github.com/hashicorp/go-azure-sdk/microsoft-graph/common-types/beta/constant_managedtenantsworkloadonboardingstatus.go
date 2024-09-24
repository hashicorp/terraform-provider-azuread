package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadOnboardingStatus string

const (
	ManagedTenantsWorkloadOnboardingStatus_NotOnboarded ManagedTenantsWorkloadOnboardingStatus = "notOnboarded"
	ManagedTenantsWorkloadOnboardingStatus_Onboarded    ManagedTenantsWorkloadOnboardingStatus = "onboarded"
)

func PossibleValuesForManagedTenantsWorkloadOnboardingStatus() []string {
	return []string{
		string(ManagedTenantsWorkloadOnboardingStatus_NotOnboarded),
		string(ManagedTenantsWorkloadOnboardingStatus_Onboarded),
	}
}

func (s *ManagedTenantsWorkloadOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsWorkloadOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsWorkloadOnboardingStatus(input string) (*ManagedTenantsWorkloadOnboardingStatus, error) {
	vals := map[string]ManagedTenantsWorkloadOnboardingStatus{
		"notonboarded": ManagedTenantsWorkloadOnboardingStatus_NotOnboarded,
		"onboarded":    ManagedTenantsWorkloadOnboardingStatus_Onboarded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsWorkloadOnboardingStatus(input)
	return &out, nil
}
