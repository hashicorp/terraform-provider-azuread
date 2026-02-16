package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceHealthStatus string

const (
	SecurityDeviceHealthStatus_Active                            SecurityDeviceHealthStatus = "active"
	SecurityDeviceHealthStatus_ImpairedCommunication             SecurityDeviceHealthStatus = "impairedCommunication"
	SecurityDeviceHealthStatus_Inactive                          SecurityDeviceHealthStatus = "inactive"
	SecurityDeviceHealthStatus_NoSensorData                      SecurityDeviceHealthStatus = "noSensorData"
	SecurityDeviceHealthStatus_NoSensorDataImpairedCommunication SecurityDeviceHealthStatus = "noSensorDataImpairedCommunication"
	SecurityDeviceHealthStatus_Unknown                           SecurityDeviceHealthStatus = "unknown"
)

func PossibleValuesForSecurityDeviceHealthStatus() []string {
	return []string{
		string(SecurityDeviceHealthStatus_Active),
		string(SecurityDeviceHealthStatus_ImpairedCommunication),
		string(SecurityDeviceHealthStatus_Inactive),
		string(SecurityDeviceHealthStatus_NoSensorData),
		string(SecurityDeviceHealthStatus_NoSensorDataImpairedCommunication),
		string(SecurityDeviceHealthStatus_Unknown),
	}
}

func (s *SecurityDeviceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceHealthStatus(input string) (*SecurityDeviceHealthStatus, error) {
	vals := map[string]SecurityDeviceHealthStatus{
		"active":                            SecurityDeviceHealthStatus_Active,
		"impairedcommunication":             SecurityDeviceHealthStatus_ImpairedCommunication,
		"inactive":                          SecurityDeviceHealthStatus_Inactive,
		"nosensordata":                      SecurityDeviceHealthStatus_NoSensorData,
		"nosensordataimpairedcommunication": SecurityDeviceHealthStatus_NoSensorDataImpairedCommunication,
		"unknown":                           SecurityDeviceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceHealthStatus(input)
	return &out, nil
}
