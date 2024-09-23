package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WeeklySchedule string

const (
	WeeklySchedule_Everyday    WeeklySchedule = "everyday"
	WeeklySchedule_Friday      WeeklySchedule = "friday"
	WeeklySchedule_Monday      WeeklySchedule = "monday"
	WeeklySchedule_Saturday    WeeklySchedule = "saturday"
	WeeklySchedule_Sunday      WeeklySchedule = "sunday"
	WeeklySchedule_Thursday    WeeklySchedule = "thursday"
	WeeklySchedule_Tuesday     WeeklySchedule = "tuesday"
	WeeklySchedule_UserDefined WeeklySchedule = "userDefined"
	WeeklySchedule_Wednesday   WeeklySchedule = "wednesday"
)

func PossibleValuesForWeeklySchedule() []string {
	return []string{
		string(WeeklySchedule_Everyday),
		string(WeeklySchedule_Friday),
		string(WeeklySchedule_Monday),
		string(WeeklySchedule_Saturday),
		string(WeeklySchedule_Sunday),
		string(WeeklySchedule_Thursday),
		string(WeeklySchedule_Tuesday),
		string(WeeklySchedule_UserDefined),
		string(WeeklySchedule_Wednesday),
	}
}

func (s *WeeklySchedule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWeeklySchedule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWeeklySchedule(input string) (*WeeklySchedule, error) {
	vals := map[string]WeeklySchedule{
		"everyday":    WeeklySchedule_Everyday,
		"friday":      WeeklySchedule_Friday,
		"monday":      WeeklySchedule_Monday,
		"saturday":    WeeklySchedule_Saturday,
		"sunday":      WeeklySchedule_Sunday,
		"thursday":    WeeklySchedule_Thursday,
		"tuesday":     WeeklySchedule_Tuesday,
		"userdefined": WeeklySchedule_UserDefined,
		"wednesday":   WeeklySchedule_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WeeklySchedule(input)
	return &out, nil
}
