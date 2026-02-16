package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistanceOnboardingStatus string

const (
	RemoteAssistanceOnboardingStatus_NotOnboarded RemoteAssistanceOnboardingStatus = "notOnboarded"
	RemoteAssistanceOnboardingStatus_Onboarded    RemoteAssistanceOnboardingStatus = "onboarded"
	RemoteAssistanceOnboardingStatus_Onboarding   RemoteAssistanceOnboardingStatus = "onboarding"
)

func PossibleValuesForRemoteAssistanceOnboardingStatus() []string {
	return []string{
		string(RemoteAssistanceOnboardingStatus_NotOnboarded),
		string(RemoteAssistanceOnboardingStatus_Onboarded),
		string(RemoteAssistanceOnboardingStatus_Onboarding),
	}
}

func (s *RemoteAssistanceOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemoteAssistanceOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemoteAssistanceOnboardingStatus(input string) (*RemoteAssistanceOnboardingStatus, error) {
	vals := map[string]RemoteAssistanceOnboardingStatus{
		"notonboarded": RemoteAssistanceOnboardingStatus_NotOnboarded,
		"onboarded":    RemoteAssistanceOnboardingStatus_Onboarded,
		"onboarding":   RemoteAssistanceOnboardingStatus_Onboarding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteAssistanceOnboardingStatus(input)
	return &out, nil
}
