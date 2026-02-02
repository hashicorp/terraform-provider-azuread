package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FreeBusyStatus string

const (
	FreeBusyStatus_Busy             FreeBusyStatus = "busy"
	FreeBusyStatus_Free             FreeBusyStatus = "free"
	FreeBusyStatus_Oof              FreeBusyStatus = "oof"
	FreeBusyStatus_Tentative        FreeBusyStatus = "tentative"
	FreeBusyStatus_Unknown          FreeBusyStatus = "unknown"
	FreeBusyStatus_WorkingElsewhere FreeBusyStatus = "workingElsewhere"
)

func PossibleValuesForFreeBusyStatus() []string {
	return []string{
		string(FreeBusyStatus_Busy),
		string(FreeBusyStatus_Free),
		string(FreeBusyStatus_Oof),
		string(FreeBusyStatus_Tentative),
		string(FreeBusyStatus_Unknown),
		string(FreeBusyStatus_WorkingElsewhere),
	}
}

func (s *FreeBusyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFreeBusyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFreeBusyStatus(input string) (*FreeBusyStatus, error) {
	vals := map[string]FreeBusyStatus{
		"busy":             FreeBusyStatus_Busy,
		"free":             FreeBusyStatus_Free,
		"oof":              FreeBusyStatus_Oof,
		"tentative":        FreeBusyStatus_Tentative,
		"unknown":          FreeBusyStatus_Unknown,
		"workingelsewhere": FreeBusyStatus_WorkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FreeBusyStatus(input)
	return &out, nil
}
