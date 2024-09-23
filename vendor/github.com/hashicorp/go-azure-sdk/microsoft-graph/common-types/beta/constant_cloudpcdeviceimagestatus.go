package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageStatus string

const (
	CloudPCDeviceImageStatus_Failed  CloudPCDeviceImageStatus = "failed"
	CloudPCDeviceImageStatus_Pending CloudPCDeviceImageStatus = "pending"
	CloudPCDeviceImageStatus_Ready   CloudPCDeviceImageStatus = "ready"
)

func PossibleValuesForCloudPCDeviceImageStatus() []string {
	return []string{
		string(CloudPCDeviceImageStatus_Failed),
		string(CloudPCDeviceImageStatus_Pending),
		string(CloudPCDeviceImageStatus_Ready),
	}
}

func (s *CloudPCDeviceImageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDeviceImageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDeviceImageStatus(input string) (*CloudPCDeviceImageStatus, error) {
	vals := map[string]CloudPCDeviceImageStatus{
		"failed":  CloudPCDeviceImageStatus_Failed,
		"pending": CloudPCDeviceImageStatus_Pending,
		"ready":   CloudPCDeviceImageStatus_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageStatus(input)
	return &out, nil
}
