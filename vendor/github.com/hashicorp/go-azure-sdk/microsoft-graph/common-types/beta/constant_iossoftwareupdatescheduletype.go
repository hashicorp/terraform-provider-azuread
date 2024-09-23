package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosSoftwareUpdateScheduleType string

const (
	IosSoftwareUpdateScheduleType_AlwaysUpdate               IosSoftwareUpdateScheduleType = "alwaysUpdate"
	IosSoftwareUpdateScheduleType_UpdateDuringTimeWindows    IosSoftwareUpdateScheduleType = "updateDuringTimeWindows"
	IosSoftwareUpdateScheduleType_UpdateOutsideOfActiveHours IosSoftwareUpdateScheduleType = "updateOutsideOfActiveHours"
	IosSoftwareUpdateScheduleType_UpdateOutsideOfTimeWindows IosSoftwareUpdateScheduleType = "updateOutsideOfTimeWindows"
)

func PossibleValuesForIosSoftwareUpdateScheduleType() []string {
	return []string{
		string(IosSoftwareUpdateScheduleType_AlwaysUpdate),
		string(IosSoftwareUpdateScheduleType_UpdateDuringTimeWindows),
		string(IosSoftwareUpdateScheduleType_UpdateOutsideOfActiveHours),
		string(IosSoftwareUpdateScheduleType_UpdateOutsideOfTimeWindows),
	}
}

func (s *IosSoftwareUpdateScheduleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosSoftwareUpdateScheduleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosSoftwareUpdateScheduleType(input string) (*IosSoftwareUpdateScheduleType, error) {
	vals := map[string]IosSoftwareUpdateScheduleType{
		"alwaysupdate":               IosSoftwareUpdateScheduleType_AlwaysUpdate,
		"updateduringtimewindows":    IosSoftwareUpdateScheduleType_UpdateDuringTimeWindows,
		"updateoutsideofactivehours": IosSoftwareUpdateScheduleType_UpdateOutsideOfActiveHours,
		"updateoutsideoftimewindows": IosSoftwareUpdateScheduleType_UpdateOutsideOfTimeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosSoftwareUpdateScheduleType(input)
	return &out, nil
}
