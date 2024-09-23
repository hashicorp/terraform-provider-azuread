package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBulkActionStatus string

const (
	CloudPCBulkActionStatus_Failed    CloudPCBulkActionStatus = "failed"
	CloudPCBulkActionStatus_Pending   CloudPCBulkActionStatus = "pending"
	CloudPCBulkActionStatus_Succeeded CloudPCBulkActionStatus = "succeeded"
)

func PossibleValuesForCloudPCBulkActionStatus() []string {
	return []string{
		string(CloudPCBulkActionStatus_Failed),
		string(CloudPCBulkActionStatus_Pending),
		string(CloudPCBulkActionStatus_Succeeded),
	}
}

func (s *CloudPCBulkActionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCBulkActionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCBulkActionStatus(input string) (*CloudPCBulkActionStatus, error) {
	vals := map[string]CloudPCBulkActionStatus{
		"failed":    CloudPCBulkActionStatus_Failed,
		"pending":   CloudPCBulkActionStatus_Pending,
		"succeeded": CloudPCBulkActionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCBulkActionStatus(input)
	return &out, nil
}
