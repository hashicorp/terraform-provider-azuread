package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderMonitorFileActivity string

const (
	DefenderMonitorFileActivity_Disable                  DefenderMonitorFileActivity = "disable"
	DefenderMonitorFileActivity_MonitorAllFiles          DefenderMonitorFileActivity = "monitorAllFiles"
	DefenderMonitorFileActivity_MonitorIncomingFilesOnly DefenderMonitorFileActivity = "monitorIncomingFilesOnly"
	DefenderMonitorFileActivity_MonitorOutgoingFilesOnly DefenderMonitorFileActivity = "monitorOutgoingFilesOnly"
	DefenderMonitorFileActivity_UserDefined              DefenderMonitorFileActivity = "userDefined"
)

func PossibleValuesForDefenderMonitorFileActivity() []string {
	return []string{
		string(DefenderMonitorFileActivity_Disable),
		string(DefenderMonitorFileActivity_MonitorAllFiles),
		string(DefenderMonitorFileActivity_MonitorIncomingFilesOnly),
		string(DefenderMonitorFileActivity_MonitorOutgoingFilesOnly),
		string(DefenderMonitorFileActivity_UserDefined),
	}
}

func (s *DefenderMonitorFileActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderMonitorFileActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderMonitorFileActivity(input string) (*DefenderMonitorFileActivity, error) {
	vals := map[string]DefenderMonitorFileActivity{
		"disable":                  DefenderMonitorFileActivity_Disable,
		"monitorallfiles":          DefenderMonitorFileActivity_MonitorAllFiles,
		"monitorincomingfilesonly": DefenderMonitorFileActivity_MonitorIncomingFilesOnly,
		"monitoroutgoingfilesonly": DefenderMonitorFileActivity_MonitorOutgoingFilesOnly,
		"userdefined":              DefenderMonitorFileActivity_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderMonitorFileActivity(input)
	return &out, nil
}
