package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChromeOSOnboardingStatus string

const (
	ChromeOSOnboardingStatus_Failed      ChromeOSOnboardingStatus = "failed"
	ChromeOSOnboardingStatus_Inprogress  ChromeOSOnboardingStatus = "inprogress"
	ChromeOSOnboardingStatus_Offboarding ChromeOSOnboardingStatus = "offboarding"
	ChromeOSOnboardingStatus_Onboarded   ChromeOSOnboardingStatus = "onboarded"
	ChromeOSOnboardingStatus_Unknown     ChromeOSOnboardingStatus = "unknown"
)

func PossibleValuesForChromeOSOnboardingStatus() []string {
	return []string{
		string(ChromeOSOnboardingStatus_Failed),
		string(ChromeOSOnboardingStatus_Inprogress),
		string(ChromeOSOnboardingStatus_Offboarding),
		string(ChromeOSOnboardingStatus_Onboarded),
		string(ChromeOSOnboardingStatus_Unknown),
	}
}

func (s *ChromeOSOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChromeOSOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChromeOSOnboardingStatus(input string) (*ChromeOSOnboardingStatus, error) {
	vals := map[string]ChromeOSOnboardingStatus{
		"failed":      ChromeOSOnboardingStatus_Failed,
		"inprogress":  ChromeOSOnboardingStatus_Inprogress,
		"offboarding": ChromeOSOnboardingStatus_Offboarding,
		"onboarded":   ChromeOSOnboardingStatus_Onboarded,
		"unknown":     ChromeOSOnboardingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChromeOSOnboardingStatus(input)
	return &out, nil
}
