package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionStatus string

const (
	CloudPCOnPremisesConnectionStatus_Failed        CloudPCOnPremisesConnectionStatus = "failed"
	CloudPCOnPremisesConnectionStatus_Informational CloudPCOnPremisesConnectionStatus = "informational"
	CloudPCOnPremisesConnectionStatus_Passed        CloudPCOnPremisesConnectionStatus = "passed"
	CloudPCOnPremisesConnectionStatus_Pending       CloudPCOnPremisesConnectionStatus = "pending"
	CloudPCOnPremisesConnectionStatus_Running       CloudPCOnPremisesConnectionStatus = "running"
	CloudPCOnPremisesConnectionStatus_Warning       CloudPCOnPremisesConnectionStatus = "warning"
)

func PossibleValuesForCloudPCOnPremisesConnectionStatus() []string {
	return []string{
		string(CloudPCOnPremisesConnectionStatus_Failed),
		string(CloudPCOnPremisesConnectionStatus_Informational),
		string(CloudPCOnPremisesConnectionStatus_Passed),
		string(CloudPCOnPremisesConnectionStatus_Pending),
		string(CloudPCOnPremisesConnectionStatus_Running),
		string(CloudPCOnPremisesConnectionStatus_Warning),
	}
}

func (s *CloudPCOnPremisesConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOnPremisesConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOnPremisesConnectionStatus(input string) (*CloudPCOnPremisesConnectionStatus, error) {
	vals := map[string]CloudPCOnPremisesConnectionStatus{
		"failed":        CloudPCOnPremisesConnectionStatus_Failed,
		"informational": CloudPCOnPremisesConnectionStatus_Informational,
		"passed":        CloudPCOnPremisesConnectionStatus_Passed,
		"pending":       CloudPCOnPremisesConnectionStatus_Pending,
		"running":       CloudPCOnPremisesConnectionStatus_Running,
		"warning":       CloudPCOnPremisesConnectionStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionStatus(input)
	return &out, nil
}
