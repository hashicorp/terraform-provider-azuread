package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentHealthStatus string

const (
	CloudPCAgentHealthStatus_Healthy     CloudPCAgentHealthStatus = "healthy"
	CloudPCAgentHealthStatus_Unavailable CloudPCAgentHealthStatus = "unavailable"
	CloudPCAgentHealthStatus_Warning     CloudPCAgentHealthStatus = "warning"
)

func PossibleValuesForCloudPCAgentHealthStatus() []string {
	return []string{
		string(CloudPCAgentHealthStatus_Healthy),
		string(CloudPCAgentHealthStatus_Unavailable),
		string(CloudPCAgentHealthStatus_Warning),
	}
}

func (s *CloudPCAgentHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAgentHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAgentHealthStatus(input string) (*CloudPCAgentHealthStatus, error) {
	vals := map[string]CloudPCAgentHealthStatus{
		"healthy":     CloudPCAgentHealthStatus_Healthy,
		"unavailable": CloudPCAgentHealthStatus_Unavailable,
		"warning":     CloudPCAgentHealthStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAgentHealthStatus(input)
	return &out, nil
}
