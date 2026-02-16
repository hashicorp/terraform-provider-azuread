package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsNetworkTransportProtocol string

const (
	CallRecordsNetworkTransportProtocol_Tcp     CallRecordsNetworkTransportProtocol = "tcp"
	CallRecordsNetworkTransportProtocol_Udp     CallRecordsNetworkTransportProtocol = "udp"
	CallRecordsNetworkTransportProtocol_Unknown CallRecordsNetworkTransportProtocol = "unknown"
)

func PossibleValuesForCallRecordsNetworkTransportProtocol() []string {
	return []string{
		string(CallRecordsNetworkTransportProtocol_Tcp),
		string(CallRecordsNetworkTransportProtocol_Udp),
		string(CallRecordsNetworkTransportProtocol_Unknown),
	}
}

func (s *CallRecordsNetworkTransportProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsNetworkTransportProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsNetworkTransportProtocol(input string) (*CallRecordsNetworkTransportProtocol, error) {
	vals := map[string]CallRecordsNetworkTransportProtocol{
		"tcp":     CallRecordsNetworkTransportProtocol_Tcp,
		"udp":     CallRecordsNetworkTransportProtocol_Udp,
		"unknown": CallRecordsNetworkTransportProtocol_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsNetworkTransportProtocol(input)
	return &out, nil
}
