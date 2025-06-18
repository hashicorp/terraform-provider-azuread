package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCPolicyApplyActionStatus string

const (
	CloudPCPolicyApplyActionStatus_Failed     CloudPCPolicyApplyActionStatus = "failed"
	CloudPCPolicyApplyActionStatus_Processing CloudPCPolicyApplyActionStatus = "processing"
	CloudPCPolicyApplyActionStatus_Succeeded  CloudPCPolicyApplyActionStatus = "succeeded"
)

func PossibleValuesForCloudPCPolicyApplyActionStatus() []string {
	return []string{
		string(CloudPCPolicyApplyActionStatus_Failed),
		string(CloudPCPolicyApplyActionStatus_Processing),
		string(CloudPCPolicyApplyActionStatus_Succeeded),
	}
}

func (s *CloudPCPolicyApplyActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCPolicyApplyActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCPolicyApplyActionStatus(input string) (*CloudPCPolicyApplyActionStatus, error) {
	vals := map[string]CloudPCPolicyApplyActionStatus{
		"failed":     CloudPCPolicyApplyActionStatus_Failed,
		"processing": CloudPCPolicyApplyActionStatus_Processing,
		"succeeded":  CloudPCPolicyApplyActionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCPolicyApplyActionStatus(input)
	return &out, nil
}
