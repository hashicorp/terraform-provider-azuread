package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsCallType string

const (
	CallRecordsCallType_GroupCall  CallRecordsCallType = "groupCall"
	CallRecordsCallType_PeerToPeer CallRecordsCallType = "peerToPeer"
	CallRecordsCallType_Unknown    CallRecordsCallType = "unknown"
)

func PossibleValuesForCallRecordsCallType() []string {
	return []string{
		string(CallRecordsCallType_GroupCall),
		string(CallRecordsCallType_PeerToPeer),
		string(CallRecordsCallType_Unknown),
	}
}

func (s *CallRecordsCallType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsCallType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsCallType(input string) (*CallRecordsCallType, error) {
	vals := map[string]CallRecordsCallType{
		"groupcall":  CallRecordsCallType_GroupCall,
		"peertopeer": CallRecordsCallType_PeerToPeer,
		"unknown":    CallRecordsCallType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsCallType(input)
	return &out, nil
}
