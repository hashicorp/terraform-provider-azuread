package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnboardingStatus string

const (
	OnboardingStatus_Failed      OnboardingStatus = "failed"
	OnboardingStatus_Inprogress  OnboardingStatus = "inprogress"
	OnboardingStatus_Offboarding OnboardingStatus = "offboarding"
	OnboardingStatus_Onboarded   OnboardingStatus = "onboarded"
	OnboardingStatus_Unknown     OnboardingStatus = "unknown"
)

func PossibleValuesForOnboardingStatus() []string {
	return []string{
		string(OnboardingStatus_Failed),
		string(OnboardingStatus_Inprogress),
		string(OnboardingStatus_Offboarding),
		string(OnboardingStatus_Onboarded),
		string(OnboardingStatus_Unknown),
	}
}

func (s *OnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnboardingStatus(input string) (*OnboardingStatus, error) {
	vals := map[string]OnboardingStatus{
		"failed":      OnboardingStatus_Failed,
		"inprogress":  OnboardingStatus_Inprogress,
		"offboarding": OnboardingStatus_Offboarding,
		"onboarded":   OnboardingStatus_Onboarded,
		"unknown":     OnboardingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnboardingStatus(input)
	return &out, nil
}
