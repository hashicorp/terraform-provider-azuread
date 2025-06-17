package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentHealthCheckState string

const (
	CloudPCAgentHealthCheckState_Canceled   CloudPCAgentHealthCheckState = "canceled"
	CloudPCAgentHealthCheckState_Conflict   CloudPCAgentHealthCheckState = "conflict"
	CloudPCAgentHealthCheckState_Failed     CloudPCAgentHealthCheckState = "failed"
	CloudPCAgentHealthCheckState_Pending    CloudPCAgentHealthCheckState = "pending"
	CloudPCAgentHealthCheckState_Processing CloudPCAgentHealthCheckState = "processing"
	CloudPCAgentHealthCheckState_Succeeded  CloudPCAgentHealthCheckState = "succeeded"
)

func PossibleValuesForCloudPCAgentHealthCheckState() []string {
	return []string{
		string(CloudPCAgentHealthCheckState_Canceled),
		string(CloudPCAgentHealthCheckState_Conflict),
		string(CloudPCAgentHealthCheckState_Failed),
		string(CloudPCAgentHealthCheckState_Pending),
		string(CloudPCAgentHealthCheckState_Processing),
		string(CloudPCAgentHealthCheckState_Succeeded),
	}
}

func (s *CloudPCAgentHealthCheckState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAgentHealthCheckState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAgentHealthCheckState(input string) (*CloudPCAgentHealthCheckState, error) {
	vals := map[string]CloudPCAgentHealthCheckState{
		"canceled":   CloudPCAgentHealthCheckState_Canceled,
		"conflict":   CloudPCAgentHealthCheckState_Conflict,
		"failed":     CloudPCAgentHealthCheckState_Failed,
		"pending":    CloudPCAgentHealthCheckState_Pending,
		"processing": CloudPCAgentHealthCheckState_Processing,
		"succeeded":  CloudPCAgentHealthCheckState_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAgentHealthCheckState(input)
	return &out, nil
}
