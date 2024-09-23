package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsModality string

const (
	CallRecordsModality_Audio                   CallRecordsModality = "audio"
	CallRecordsModality_Data                    CallRecordsModality = "data"
	CallRecordsModality_ScreenSharing           CallRecordsModality = "screenSharing"
	CallRecordsModality_Video                   CallRecordsModality = "video"
	CallRecordsModality_VideoBasedScreenSharing CallRecordsModality = "videoBasedScreenSharing"
)

func PossibleValuesForCallRecordsModality() []string {
	return []string{
		string(CallRecordsModality_Audio),
		string(CallRecordsModality_Data),
		string(CallRecordsModality_ScreenSharing),
		string(CallRecordsModality_Video),
		string(CallRecordsModality_VideoBasedScreenSharing),
	}
}

func (s *CallRecordsModality) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsModality(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsModality(input string) (*CallRecordsModality, error) {
	vals := map[string]CallRecordsModality{
		"audio":                   CallRecordsModality_Audio,
		"data":                    CallRecordsModality_Data,
		"screensharing":           CallRecordsModality_ScreenSharing,
		"video":                   CallRecordsModality_Video,
		"videobasedscreensharing": CallRecordsModality_VideoBasedScreenSharing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsModality(input)
	return &out, nil
}
