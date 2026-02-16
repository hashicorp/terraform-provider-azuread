package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringAction string

const (
	WindowsUpdatesMonitoringAction_AlertError      WindowsUpdatesMonitoringAction = "alertError"
	WindowsUpdatesMonitoringAction_OfferFallback   WindowsUpdatesMonitoringAction = "offerFallback"
	WindowsUpdatesMonitoringAction_PauseDeployment WindowsUpdatesMonitoringAction = "pauseDeployment"
)

func PossibleValuesForWindowsUpdatesMonitoringAction() []string {
	return []string{
		string(WindowsUpdatesMonitoringAction_AlertError),
		string(WindowsUpdatesMonitoringAction_OfferFallback),
		string(WindowsUpdatesMonitoringAction_PauseDeployment),
	}
}

func (s *WindowsUpdatesMonitoringAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesMonitoringAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesMonitoringAction(input string) (*WindowsUpdatesMonitoringAction, error) {
	vals := map[string]WindowsUpdatesMonitoringAction{
		"alerterror":      WindowsUpdatesMonitoringAction_AlertError,
		"offerfallback":   WindowsUpdatesMonitoringAction_OfferFallback,
		"pausedeployment": WindowsUpdatesMonitoringAction_PauseDeployment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesMonitoringAction(input)
	return &out, nil
}
