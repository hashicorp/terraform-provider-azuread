package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageOsStatus string

const (
	CloudPCDeviceImageOsStatus_Supported            CloudPCDeviceImageOsStatus = "supported"
	CloudPCDeviceImageOsStatus_SupportedWithWarning CloudPCDeviceImageOsStatus = "supportedWithWarning"
	CloudPCDeviceImageOsStatus_Unknown              CloudPCDeviceImageOsStatus = "unknown"
)

func PossibleValuesForCloudPCDeviceImageOsStatus() []string {
	return []string{
		string(CloudPCDeviceImageOsStatus_Supported),
		string(CloudPCDeviceImageOsStatus_SupportedWithWarning),
		string(CloudPCDeviceImageOsStatus_Unknown),
	}
}

func (s *CloudPCDeviceImageOsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDeviceImageOsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDeviceImageOsStatus(input string) (*CloudPCDeviceImageOsStatus, error) {
	vals := map[string]CloudPCDeviceImageOsStatus{
		"supported":            CloudPCDeviceImageOsStatus_Supported,
		"supportedwithwarning": CloudPCDeviceImageOsStatus_SupportedWithWarning,
		"unknown":              CloudPCDeviceImageOsStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageOsStatus(input)
	return &out, nil
}
