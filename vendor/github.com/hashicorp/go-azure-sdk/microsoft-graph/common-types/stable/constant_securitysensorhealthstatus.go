package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySensorHealthStatus string

const (
	SecuritySensorHealthStatus_Healthy          SecuritySensorHealthStatus = "healthy"
	SecuritySensorHealthStatus_NotHealthyHigh   SecuritySensorHealthStatus = "notHealthyHigh"
	SecuritySensorHealthStatus_NotHealthyLow    SecuritySensorHealthStatus = "notHealthyLow"
	SecuritySensorHealthStatus_NotHealthyMedium SecuritySensorHealthStatus = "notHealthyMedium"
)

func PossibleValuesForSecuritySensorHealthStatus() []string {
	return []string{
		string(SecuritySensorHealthStatus_Healthy),
		string(SecuritySensorHealthStatus_NotHealthyHigh),
		string(SecuritySensorHealthStatus_NotHealthyLow),
		string(SecuritySensorHealthStatus_NotHealthyMedium),
	}
}

func (s *SecuritySensorHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySensorHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySensorHealthStatus(input string) (*SecuritySensorHealthStatus, error) {
	vals := map[string]SecuritySensorHealthStatus{
		"healthy":          SecuritySensorHealthStatus_Healthy,
		"nothealthyhigh":   SecuritySensorHealthStatus_NotHealthyHigh,
		"nothealthylow":    SecuritySensorHealthStatus_NotHealthyLow,
		"nothealthymedium": SecuritySensorHealthStatus_NotHealthyMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySensorHealthStatus(input)
	return &out, nil
}
