package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderRealtimeScanDirection string

const (
	DefenderRealtimeScanDirection_MonitorAllFiles          DefenderRealtimeScanDirection = "monitorAllFiles"
	DefenderRealtimeScanDirection_MonitorIncomingFilesOnly DefenderRealtimeScanDirection = "monitorIncomingFilesOnly"
	DefenderRealtimeScanDirection_MonitorOutgoingFilesOnly DefenderRealtimeScanDirection = "monitorOutgoingFilesOnly"
)

func PossibleValuesForDefenderRealtimeScanDirection() []string {
	return []string{
		string(DefenderRealtimeScanDirection_MonitorAllFiles),
		string(DefenderRealtimeScanDirection_MonitorIncomingFilesOnly),
		string(DefenderRealtimeScanDirection_MonitorOutgoingFilesOnly),
	}
}

func (s *DefenderRealtimeScanDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderRealtimeScanDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderRealtimeScanDirection(input string) (*DefenderRealtimeScanDirection, error) {
	vals := map[string]DefenderRealtimeScanDirection{
		"monitorallfiles":          DefenderRealtimeScanDirection_MonitorAllFiles,
		"monitorincomingfilesonly": DefenderRealtimeScanDirection_MonitorIncomingFilesOnly,
		"monitoroutgoingfilesonly": DefenderRealtimeScanDirection_MonitorOutgoingFilesOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderRealtimeScanDirection(input)
	return &out, nil
}
