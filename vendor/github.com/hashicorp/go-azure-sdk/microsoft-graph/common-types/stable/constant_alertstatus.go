package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertStatus string

const (
	AlertStatus_Dismissed  AlertStatus = "dismissed"
	AlertStatus_InProgress AlertStatus = "inProgress"
	AlertStatus_NewAlert   AlertStatus = "newAlert"
	AlertStatus_Resolved   AlertStatus = "resolved"
	AlertStatus_Unknown    AlertStatus = "unknown"
)

func PossibleValuesForAlertStatus() []string {
	return []string{
		string(AlertStatus_Dismissed),
		string(AlertStatus_InProgress),
		string(AlertStatus_NewAlert),
		string(AlertStatus_Resolved),
		string(AlertStatus_Unknown),
	}
}

func (s *AlertStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertStatus(input string) (*AlertStatus, error) {
	vals := map[string]AlertStatus{
		"dismissed":  AlertStatus_Dismissed,
		"inprogress": AlertStatus_InProgress,
		"newalert":   AlertStatus_NewAlert,
		"resolved":   AlertStatus_Resolved,
		"unknown":    AlertStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertStatus(input)
	return &out, nil
}
