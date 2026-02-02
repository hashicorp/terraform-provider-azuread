package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityStatus string

const (
	CloudPCConnectivityStatus_Available            CloudPCConnectivityStatus = "available"
	CloudPCConnectivityStatus_AvailableWithWarning CloudPCConnectivityStatus = "availableWithWarning"
	CloudPCConnectivityStatus_Unavailable          CloudPCConnectivityStatus = "unavailable"
	CloudPCConnectivityStatus_Unknown              CloudPCConnectivityStatus = "unknown"
)

func PossibleValuesForCloudPCConnectivityStatus() []string {
	return []string{
		string(CloudPCConnectivityStatus_Available),
		string(CloudPCConnectivityStatus_AvailableWithWarning),
		string(CloudPCConnectivityStatus_Unavailable),
		string(CloudPCConnectivityStatus_Unknown),
	}
}

func (s *CloudPCConnectivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityStatus(input string) (*CloudPCConnectivityStatus, error) {
	vals := map[string]CloudPCConnectivityStatus{
		"available":            CloudPCConnectivityStatus_Available,
		"availablewithwarning": CloudPCConnectivityStatus_AvailableWithWarning,
		"unavailable":          CloudPCConnectivityStatus_Unavailable,
		"unknown":              CloudPCConnectivityStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityStatus(input)
	return &out, nil
}
