package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateScheduleType string

const (
	MacOSSoftwareUpdateScheduleType_AlwaysUpdate               MacOSSoftwareUpdateScheduleType = "alwaysUpdate"
	MacOSSoftwareUpdateScheduleType_UpdateDuringTimeWindows    MacOSSoftwareUpdateScheduleType = "updateDuringTimeWindows"
	MacOSSoftwareUpdateScheduleType_UpdateOutsideOfTimeWindows MacOSSoftwareUpdateScheduleType = "updateOutsideOfTimeWindows"
)

func PossibleValuesForMacOSSoftwareUpdateScheduleType() []string {
	return []string{
		string(MacOSSoftwareUpdateScheduleType_AlwaysUpdate),
		string(MacOSSoftwareUpdateScheduleType_UpdateDuringTimeWindows),
		string(MacOSSoftwareUpdateScheduleType_UpdateOutsideOfTimeWindows),
	}
}

func (s *MacOSSoftwareUpdateScheduleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateScheduleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateScheduleType(input string) (*MacOSSoftwareUpdateScheduleType, error) {
	vals := map[string]MacOSSoftwareUpdateScheduleType{
		"alwaysupdate":               MacOSSoftwareUpdateScheduleType_AlwaysUpdate,
		"updateduringtimewindows":    MacOSSoftwareUpdateScheduleType_UpdateDuringTimeWindows,
		"updateoutsideoftimewindows": MacOSSoftwareUpdateScheduleType_UpdateOutsideOfTimeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateScheduleType(input)
	return &out, nil
}
