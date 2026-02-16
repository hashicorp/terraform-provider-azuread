package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityEventResult string

const (
	CloudPCConnectivityEventResult_Failure CloudPCConnectivityEventResult = "failure"
	CloudPCConnectivityEventResult_Success CloudPCConnectivityEventResult = "success"
	CloudPCConnectivityEventResult_Unknown CloudPCConnectivityEventResult = "unknown"
)

func PossibleValuesForCloudPCConnectivityEventResult() []string {
	return []string{
		string(CloudPCConnectivityEventResult_Failure),
		string(CloudPCConnectivityEventResult_Success),
		string(CloudPCConnectivityEventResult_Unknown),
	}
}

func (s *CloudPCConnectivityEventResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCConnectivityEventResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCConnectivityEventResult(input string) (*CloudPCConnectivityEventResult, error) {
	vals := map[string]CloudPCConnectivityEventResult{
		"failure": CloudPCConnectivityEventResult_Failure,
		"success": CloudPCConnectivityEventResult_Success,
		"unknown": CloudPCConnectivityEventResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCConnectivityEventResult(input)
	return &out, nil
}
