package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkConnectionType string

const (
	CallRecordsNetworkConnectionType_Mobile  CallRecordsNetworkConnectionType = "mobile"
	CallRecordsNetworkConnectionType_Tunnel  CallRecordsNetworkConnectionType = "tunnel"
	CallRecordsNetworkConnectionType_Unknown CallRecordsNetworkConnectionType = "unknown"
	CallRecordsNetworkConnectionType_Wifi    CallRecordsNetworkConnectionType = "wifi"
	CallRecordsNetworkConnectionType_Wired   CallRecordsNetworkConnectionType = "wired"
)

func PossibleValuesForCallRecordsNetworkConnectionType() []string {
	return []string{
		string(CallRecordsNetworkConnectionType_Mobile),
		string(CallRecordsNetworkConnectionType_Tunnel),
		string(CallRecordsNetworkConnectionType_Unknown),
		string(CallRecordsNetworkConnectionType_Wifi),
		string(CallRecordsNetworkConnectionType_Wired),
	}
}

func (s *CallRecordsNetworkConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsNetworkConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsNetworkConnectionType(input string) (*CallRecordsNetworkConnectionType, error) {
	vals := map[string]CallRecordsNetworkConnectionType{
		"mobile":  CallRecordsNetworkConnectionType_Mobile,
		"tunnel":  CallRecordsNetworkConnectionType_Tunnel,
		"unknown": CallRecordsNetworkConnectionType_Unknown,
		"wifi":    CallRecordsNetworkConnectionType_Wifi,
		"wired":   CallRecordsNetworkConnectionType_Wired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkConnectionType(input)
	return &out, nil
}
