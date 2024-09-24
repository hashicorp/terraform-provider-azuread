package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateReasonValue string

const (
	WindowsUpdatesDeploymentStateReasonValue_FaultedByContentOutdated WindowsUpdatesDeploymentStateReasonValue = "faultedByContentOutdated"
	WindowsUpdatesDeploymentStateReasonValue_OfferingByRequest        WindowsUpdatesDeploymentStateReasonValue = "offeringByRequest"
	WindowsUpdatesDeploymentStateReasonValue_PausedByMonitoring       WindowsUpdatesDeploymentStateReasonValue = "pausedByMonitoring"
	WindowsUpdatesDeploymentStateReasonValue_PausedByRequest          WindowsUpdatesDeploymentStateReasonValue = "pausedByRequest"
	WindowsUpdatesDeploymentStateReasonValue_ScheduledByOfferWindow   WindowsUpdatesDeploymentStateReasonValue = "scheduledByOfferWindow"
)

func PossibleValuesForWindowsUpdatesDeploymentStateReasonValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateReasonValue_FaultedByContentOutdated),
		string(WindowsUpdatesDeploymentStateReasonValue_OfferingByRequest),
		string(WindowsUpdatesDeploymentStateReasonValue_PausedByMonitoring),
		string(WindowsUpdatesDeploymentStateReasonValue_PausedByRequest),
		string(WindowsUpdatesDeploymentStateReasonValue_ScheduledByOfferWindow),
	}
}

func (s *WindowsUpdatesDeploymentStateReasonValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesDeploymentStateReasonValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesDeploymentStateReasonValue(input string) (*WindowsUpdatesDeploymentStateReasonValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateReasonValue{
		"faultedbycontentoutdated": WindowsUpdatesDeploymentStateReasonValue_FaultedByContentOutdated,
		"offeringbyrequest":        WindowsUpdatesDeploymentStateReasonValue_OfferingByRequest,
		"pausedbymonitoring":       WindowsUpdatesDeploymentStateReasonValue_PausedByMonitoring,
		"pausedbyrequest":          WindowsUpdatesDeploymentStateReasonValue_PausedByRequest,
		"scheduledbyofferwindow":   WindowsUpdatesDeploymentStateReasonValue_ScheduledByOfferWindow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateReasonValue(input)
	return &out, nil
}
