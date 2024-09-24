package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOnboardingStatus string

const (
	SecurityOnboardingStatus_CanBeOnboarded   SecurityOnboardingStatus = "canBeOnboarded"
	SecurityOnboardingStatus_InsufficientInfo SecurityOnboardingStatus = "insufficientInfo"
	SecurityOnboardingStatus_Onboarded        SecurityOnboardingStatus = "onboarded"
	SecurityOnboardingStatus_Unsupported      SecurityOnboardingStatus = "unsupported"
)

func PossibleValuesForSecurityOnboardingStatus() []string {
	return []string{
		string(SecurityOnboardingStatus_CanBeOnboarded),
		string(SecurityOnboardingStatus_InsufficientInfo),
		string(SecurityOnboardingStatus_Onboarded),
		string(SecurityOnboardingStatus_Unsupported),
	}
}

func (s *SecurityOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityOnboardingStatus(input string) (*SecurityOnboardingStatus, error) {
	vals := map[string]SecurityOnboardingStatus{
		"canbeonboarded":   SecurityOnboardingStatus_CanBeOnboarded,
		"insufficientinfo": SecurityOnboardingStatus_InsufficientInfo,
		"onboarded":        SecurityOnboardingStatus_Onboarded,
		"unsupported":      SecurityOnboardingStatus_Unsupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOnboardingStatus(input)
	return &out, nil
}
