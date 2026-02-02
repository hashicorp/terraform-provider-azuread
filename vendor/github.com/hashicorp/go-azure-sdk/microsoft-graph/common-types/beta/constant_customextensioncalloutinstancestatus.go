package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionCalloutInstanceStatus string

const (
	CustomExtensionCalloutInstanceStatus_CallbackReceived   CustomExtensionCalloutInstanceStatus = "callbackReceived"
	CustomExtensionCalloutInstanceStatus_CallbackTimedOut   CustomExtensionCalloutInstanceStatus = "callbackTimedOut"
	CustomExtensionCalloutInstanceStatus_CalloutFailed      CustomExtensionCalloutInstanceStatus = "calloutFailed"
	CustomExtensionCalloutInstanceStatus_CalloutSent        CustomExtensionCalloutInstanceStatus = "calloutSent"
	CustomExtensionCalloutInstanceStatus_WaitingForCallback CustomExtensionCalloutInstanceStatus = "waitingForCallback"
)

func PossibleValuesForCustomExtensionCalloutInstanceStatus() []string {
	return []string{
		string(CustomExtensionCalloutInstanceStatus_CallbackReceived),
		string(CustomExtensionCalloutInstanceStatus_CallbackTimedOut),
		string(CustomExtensionCalloutInstanceStatus_CalloutFailed),
		string(CustomExtensionCalloutInstanceStatus_CalloutSent),
		string(CustomExtensionCalloutInstanceStatus_WaitingForCallback),
	}
}

func (s *CustomExtensionCalloutInstanceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomExtensionCalloutInstanceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomExtensionCalloutInstanceStatus(input string) (*CustomExtensionCalloutInstanceStatus, error) {
	vals := map[string]CustomExtensionCalloutInstanceStatus{
		"callbackreceived":   CustomExtensionCalloutInstanceStatus_CallbackReceived,
		"callbacktimedout":   CustomExtensionCalloutInstanceStatus_CallbackTimedOut,
		"calloutfailed":      CustomExtensionCalloutInstanceStatus_CalloutFailed,
		"calloutsent":        CustomExtensionCalloutInstanceStatus_CalloutSent,
		"waitingforcallback": CustomExtensionCalloutInstanceStatus_WaitingForCallback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionCalloutInstanceStatus(input)
	return &out, nil
}
