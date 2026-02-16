package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEventType string

const (
	CloudPCConnectivityEventType_DeviceHealthCheck   CloudPCConnectivityEventType = "deviceHealthCheck"
	CloudPCConnectivityEventType_Unknown             CloudPCConnectivityEventType = "unknown"
	CloudPCConnectivityEventType_UserConnection      CloudPCConnectivityEventType = "userConnection"
	CloudPCConnectivityEventType_UserTroubleshooting CloudPCConnectivityEventType = "userTroubleshooting"
)

func PossibleValuesForCloudPCConnectivityEventType() []string {
	return []string{
		string(CloudPCConnectivityEventType_DeviceHealthCheck),
		string(CloudPCConnectivityEventType_Unknown),
		string(CloudPCConnectivityEventType_UserConnection),
		string(CloudPCConnectivityEventType_UserTroubleshooting),
	}
}

func (s *CloudPCConnectivityEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityEventType(input string) (*CloudPCConnectivityEventType, error) {
	vals := map[string]CloudPCConnectivityEventType{
		"devicehealthcheck":   CloudPCConnectivityEventType_DeviceHealthCheck,
		"unknown":             CloudPCConnectivityEventType_Unknown,
		"userconnection":      CloudPCConnectivityEventType_UserConnection,
		"usertroubleshooting": CloudPCConnectivityEventType_UserTroubleshooting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityEventType(input)
	return &out, nil
}
