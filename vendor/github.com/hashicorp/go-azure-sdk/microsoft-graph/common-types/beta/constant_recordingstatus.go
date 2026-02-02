package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordingStatus string

const (
	RecordingStatus_Failed       RecordingStatus = "failed"
	RecordingStatus_NotRecording RecordingStatus = "notRecording"
	RecordingStatus_Recording    RecordingStatus = "recording"
	RecordingStatus_Unknown      RecordingStatus = "unknown"
)

func PossibleValuesForRecordingStatus() []string {
	return []string{
		string(RecordingStatus_Failed),
		string(RecordingStatus_NotRecording),
		string(RecordingStatus_Recording),
		string(RecordingStatus_Unknown),
	}
}

func (s *RecordingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecordingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecordingStatus(input string) (*RecordingStatus, error) {
	vals := map[string]RecordingStatus{
		"failed":       RecordingStatus_Failed,
		"notrecording": RecordingStatus_NotRecording,
		"recording":    RecordingStatus_Recording,
		"unknown":      RecordingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordingStatus(input)
	return &out, nil
}
