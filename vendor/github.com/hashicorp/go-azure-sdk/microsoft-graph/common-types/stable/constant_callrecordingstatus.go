package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordingStatus string

const (
	CallRecordingStatus_ChunkFinished CallRecordingStatus = "chunkFinished"
	CallRecordingStatus_Failure       CallRecordingStatus = "failure"
	CallRecordingStatus_Initial       CallRecordingStatus = "initial"
	CallRecordingStatus_Success       CallRecordingStatus = "success"
)

func PossibleValuesForCallRecordingStatus() []string {
	return []string{
		string(CallRecordingStatus_ChunkFinished),
		string(CallRecordingStatus_Failure),
		string(CallRecordingStatus_Initial),
		string(CallRecordingStatus_Success),
	}
}

func (s *CallRecordingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordingStatus(input string) (*CallRecordingStatus, error) {
	vals := map[string]CallRecordingStatus{
		"chunkfinished": CallRecordingStatus_ChunkFinished,
		"failure":       CallRecordingStatus_Failure,
		"initial":       CallRecordingStatus_Initial,
		"success":       CallRecordingStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordingStatus(input)
	return &out, nil
}
