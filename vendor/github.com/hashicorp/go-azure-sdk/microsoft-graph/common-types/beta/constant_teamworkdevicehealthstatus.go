package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceHealthStatus string

const (
	TeamworkDeviceHealthStatus_Critical  TeamworkDeviceHealthStatus = "critical"
	TeamworkDeviceHealthStatus_Healthy   TeamworkDeviceHealthStatus = "healthy"
	TeamworkDeviceHealthStatus_NonUrgent TeamworkDeviceHealthStatus = "nonUrgent"
	TeamworkDeviceHealthStatus_Offline   TeamworkDeviceHealthStatus = "offline"
	TeamworkDeviceHealthStatus_Unknown   TeamworkDeviceHealthStatus = "unknown"
)

func PossibleValuesForTeamworkDeviceHealthStatus() []string {
	return []string{
		string(TeamworkDeviceHealthStatus_Critical),
		string(TeamworkDeviceHealthStatus_Healthy),
		string(TeamworkDeviceHealthStatus_NonUrgent),
		string(TeamworkDeviceHealthStatus_Offline),
		string(TeamworkDeviceHealthStatus_Unknown),
	}
}

func (s *TeamworkDeviceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkDeviceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkDeviceHealthStatus(input string) (*TeamworkDeviceHealthStatus, error) {
	vals := map[string]TeamworkDeviceHealthStatus{
		"critical":  TeamworkDeviceHealthStatus_Critical,
		"healthy":   TeamworkDeviceHealthStatus_Healthy,
		"nonurgent": TeamworkDeviceHealthStatus_NonUrgent,
		"offline":   TeamworkDeviceHealthStatus_Offline,
		"unknown":   TeamworkDeviceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceHealthStatus(input)
	return &out, nil
}
