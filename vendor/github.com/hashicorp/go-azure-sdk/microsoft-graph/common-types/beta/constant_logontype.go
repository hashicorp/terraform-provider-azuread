package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonType string

const (
	LogonType_Batch             LogonType = "batch"
	LogonType_Interactive       LogonType = "interactive"
	LogonType_Network           LogonType = "network"
	LogonType_RemoteInteractive LogonType = "remoteInteractive"
	LogonType_Service           LogonType = "service"
	LogonType_Unknown           LogonType = "unknown"
)

func PossibleValuesForLogonType() []string {
	return []string{
		string(LogonType_Batch),
		string(LogonType_Interactive),
		string(LogonType_Network),
		string(LogonType_RemoteInteractive),
		string(LogonType_Service),
		string(LogonType_Unknown),
	}
}

func (s *LogonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogonType(input string) (*LogonType, error) {
	vals := map[string]LogonType{
		"batch":             LogonType_Batch,
		"interactive":       LogonType_Interactive,
		"network":           LogonType_Network,
		"remoteinteractive": LogonType_RemoteInteractive,
		"service":           LogonType_Service,
		"unknown":           LogonType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogonType(input)
	return &out, nil
}
