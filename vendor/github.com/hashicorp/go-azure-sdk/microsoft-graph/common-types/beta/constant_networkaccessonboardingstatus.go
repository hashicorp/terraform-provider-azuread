package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessOnboardingStatus string

const (
	NetworkaccessOnboardingStatus_Offboarded               NetworkaccessOnboardingStatus = "offboarded"
	NetworkaccessOnboardingStatus_OffboardingErrorOccurred NetworkaccessOnboardingStatus = "offboardingErrorOccurred"
	NetworkaccessOnboardingStatus_OffboardingInProgress    NetworkaccessOnboardingStatus = "offboardingInProgress"
	NetworkaccessOnboardingStatus_Onboarded                NetworkaccessOnboardingStatus = "onboarded"
	NetworkaccessOnboardingStatus_OnboardingErrorOccurred  NetworkaccessOnboardingStatus = "onboardingErrorOccurred"
	NetworkaccessOnboardingStatus_OnboardingInProgress     NetworkaccessOnboardingStatus = "onboardingInProgress"
)

func PossibleValuesForNetworkaccessOnboardingStatus() []string {
	return []string{
		string(NetworkaccessOnboardingStatus_Offboarded),
		string(NetworkaccessOnboardingStatus_OffboardingErrorOccurred),
		string(NetworkaccessOnboardingStatus_OffboardingInProgress),
		string(NetworkaccessOnboardingStatus_Onboarded),
		string(NetworkaccessOnboardingStatus_OnboardingErrorOccurred),
		string(NetworkaccessOnboardingStatus_OnboardingInProgress),
	}
}

func (s *NetworkaccessOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessOnboardingStatus(input string) (*NetworkaccessOnboardingStatus, error) {
	vals := map[string]NetworkaccessOnboardingStatus{
		"offboarded":               NetworkaccessOnboardingStatus_Offboarded,
		"offboardingerroroccurred": NetworkaccessOnboardingStatus_OffboardingErrorOccurred,
		"offboardinginprogress":    NetworkaccessOnboardingStatus_OffboardingInProgress,
		"onboarded":                NetworkaccessOnboardingStatus_Onboarded,
		"onboardingerroroccurred":  NetworkaccessOnboardingStatus_OnboardingErrorOccurred,
		"onboardinginprogress":     NetworkaccessOnboardingStatus_OnboardingInProgress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessOnboardingStatus(input)
	return &out, nil
}
