package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsHealthMonitoringScope string

const (
	WindowsHealthMonitoringScope_BootPerformance     WindowsHealthMonitoringScope = "bootPerformance"
	WindowsHealthMonitoringScope_HealthMonitoring    WindowsHealthMonitoringScope = "healthMonitoring"
	WindowsHealthMonitoringScope_PrivilegeManagement WindowsHealthMonitoringScope = "privilegeManagement"
	WindowsHealthMonitoringScope_Undefined           WindowsHealthMonitoringScope = "undefined"
	WindowsHealthMonitoringScope_WindowsUpdates      WindowsHealthMonitoringScope = "windowsUpdates"
)

func PossibleValuesForWindowsHealthMonitoringScope() []string {
	return []string{
		string(WindowsHealthMonitoringScope_BootPerformance),
		string(WindowsHealthMonitoringScope_HealthMonitoring),
		string(WindowsHealthMonitoringScope_PrivilegeManagement),
		string(WindowsHealthMonitoringScope_Undefined),
		string(WindowsHealthMonitoringScope_WindowsUpdates),
	}
}

func (s *WindowsHealthMonitoringScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsHealthMonitoringScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsHealthMonitoringScope(input string) (*WindowsHealthMonitoringScope, error) {
	vals := map[string]WindowsHealthMonitoringScope{
		"bootperformance":     WindowsHealthMonitoringScope_BootPerformance,
		"healthmonitoring":    WindowsHealthMonitoringScope_HealthMonitoring,
		"privilegemanagement": WindowsHealthMonitoringScope_PrivilegeManagement,
		"undefined":           WindowsHealthMonitoringScope_Undefined,
		"windowsupdates":      WindowsHealthMonitoringScope_WindowsUpdates,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsHealthMonitoringScope(input)
	return &out, nil
}
