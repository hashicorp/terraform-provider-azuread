package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAgentHealthCheckResultType string

const (
	CloudPCAgentHealthCheckResultType_Failed    CloudPCAgentHealthCheckResultType = "failed"
	CloudPCAgentHealthCheckResultType_Succeeded CloudPCAgentHealthCheckResultType = "succeeded"
	CloudPCAgentHealthCheckResultType_Warning   CloudPCAgentHealthCheckResultType = "warning"
)

func PossibleValuesForCloudPCAgentHealthCheckResultType() []string {
	return []string{
		string(CloudPCAgentHealthCheckResultType_Failed),
		string(CloudPCAgentHealthCheckResultType_Succeeded),
		string(CloudPCAgentHealthCheckResultType_Warning),
	}
}

func (s *CloudPCAgentHealthCheckResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAgentHealthCheckResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAgentHealthCheckResultType(input string) (*CloudPCAgentHealthCheckResultType, error) {
	vals := map[string]CloudPCAgentHealthCheckResultType{
		"failed":    CloudPCAgentHealthCheckResultType_Failed,
		"succeeded": CloudPCAgentHealthCheckResultType_Succeeded,
		"warning":   CloudPCAgentHealthCheckResultType_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAgentHealthCheckResultType(input)
	return &out, nil
}
