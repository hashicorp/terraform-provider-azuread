package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringAlertType string

const (
	HealthMonitoringAlertType_CompliantDeviceSignInFailure   HealthMonitoringAlertType = "compliantDeviceSignInFailure"
	HealthMonitoringAlertType_ConditionalAccessBlockedSignIn HealthMonitoringAlertType = "conditionalAccessBlockedSignIn"
	HealthMonitoringAlertType_ManagedDeviceSignInFailure     HealthMonitoringAlertType = "managedDeviceSignInFailure"
	HealthMonitoringAlertType_MfaSignInFailure               HealthMonitoringAlertType = "mfaSignInFailure"
	HealthMonitoringAlertType_Unknown                        HealthMonitoringAlertType = "unknown"
)

func PossibleValuesForHealthMonitoringAlertType() []string {
	return []string{
		string(HealthMonitoringAlertType_CompliantDeviceSignInFailure),
		string(HealthMonitoringAlertType_ConditionalAccessBlockedSignIn),
		string(HealthMonitoringAlertType_ManagedDeviceSignInFailure),
		string(HealthMonitoringAlertType_MfaSignInFailure),
		string(HealthMonitoringAlertType_Unknown),
	}
}

func (s *HealthMonitoringAlertType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthMonitoringAlertType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthMonitoringAlertType(input string) (*HealthMonitoringAlertType, error) {
	vals := map[string]HealthMonitoringAlertType{
		"compliantdevicesigninfailure":   HealthMonitoringAlertType_CompliantDeviceSignInFailure,
		"conditionalaccessblockedsignin": HealthMonitoringAlertType_ConditionalAccessBlockedSignIn,
		"manageddevicesigninfailure":     HealthMonitoringAlertType_ManagedDeviceSignInFailure,
		"mfasigninfailure":               HealthMonitoringAlertType_MfaSignInFailure,
		"unknown":                        HealthMonitoringAlertType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthMonitoringAlertType(input)
	return &out, nil
}
