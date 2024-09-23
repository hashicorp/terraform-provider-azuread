package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDisasterRecoveryCapabilityType string

const (
	CloudPCDisasterRecoveryCapabilityType_Failback CloudPCDisasterRecoveryCapabilityType = "failback"
	CloudPCDisasterRecoveryCapabilityType_Failover CloudPCDisasterRecoveryCapabilityType = "failover"
	CloudPCDisasterRecoveryCapabilityType_None     CloudPCDisasterRecoveryCapabilityType = "none"
)

func PossibleValuesForCloudPCDisasterRecoveryCapabilityType() []string {
	return []string{
		string(CloudPCDisasterRecoveryCapabilityType_Failback),
		string(CloudPCDisasterRecoveryCapabilityType_Failover),
		string(CloudPCDisasterRecoveryCapabilityType_None),
	}
}

func (s *CloudPCDisasterRecoveryCapabilityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDisasterRecoveryCapabilityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDisasterRecoveryCapabilityType(input string) (*CloudPCDisasterRecoveryCapabilityType, error) {
	vals := map[string]CloudPCDisasterRecoveryCapabilityType{
		"failback": CloudPCDisasterRecoveryCapabilityType_Failback,
		"failover": CloudPCDisasterRecoveryCapabilityType_Failover,
		"none":     CloudPCDisasterRecoveryCapabilityType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDisasterRecoveryCapabilityType(input)
	return &out, nil
}
