package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffReasonIconType string

const (
	TimeOffReasonIconType_Cake        TimeOffReasonIconType = "cake"
	TimeOffReasonIconType_Calendar    TimeOffReasonIconType = "calendar"
	TimeOffReasonIconType_Car         TimeOffReasonIconType = "car"
	TimeOffReasonIconType_Clock       TimeOffReasonIconType = "clock"
	TimeOffReasonIconType_Cup         TimeOffReasonIconType = "cup"
	TimeOffReasonIconType_Doctor      TimeOffReasonIconType = "doctor"
	TimeOffReasonIconType_Dog         TimeOffReasonIconType = "dog"
	TimeOffReasonIconType_FirstAid    TimeOffReasonIconType = "firstAid"
	TimeOffReasonIconType_Globe       TimeOffReasonIconType = "globe"
	TimeOffReasonIconType_JuryDuty    TimeOffReasonIconType = "juryDuty"
	TimeOffReasonIconType_None        TimeOffReasonIconType = "none"
	TimeOffReasonIconType_NotWorking  TimeOffReasonIconType = "notWorking"
	TimeOffReasonIconType_Phone       TimeOffReasonIconType = "phone"
	TimeOffReasonIconType_PiggyBank   TimeOffReasonIconType = "piggyBank"
	TimeOffReasonIconType_Pin         TimeOffReasonIconType = "pin"
	TimeOffReasonIconType_Plane       TimeOffReasonIconType = "plane"
	TimeOffReasonIconType_Running     TimeOffReasonIconType = "running"
	TimeOffReasonIconType_Sunny       TimeOffReasonIconType = "sunny"
	TimeOffReasonIconType_TrafficCone TimeOffReasonIconType = "trafficCone"
	TimeOffReasonIconType_Umbrella    TimeOffReasonIconType = "umbrella"
	TimeOffReasonIconType_Weather     TimeOffReasonIconType = "weather"
)

func PossibleValuesForTimeOffReasonIconType() []string {
	return []string{
		string(TimeOffReasonIconType_Cake),
		string(TimeOffReasonIconType_Calendar),
		string(TimeOffReasonIconType_Car),
		string(TimeOffReasonIconType_Clock),
		string(TimeOffReasonIconType_Cup),
		string(TimeOffReasonIconType_Doctor),
		string(TimeOffReasonIconType_Dog),
		string(TimeOffReasonIconType_FirstAid),
		string(TimeOffReasonIconType_Globe),
		string(TimeOffReasonIconType_JuryDuty),
		string(TimeOffReasonIconType_None),
		string(TimeOffReasonIconType_NotWorking),
		string(TimeOffReasonIconType_Phone),
		string(TimeOffReasonIconType_PiggyBank),
		string(TimeOffReasonIconType_Pin),
		string(TimeOffReasonIconType_Plane),
		string(TimeOffReasonIconType_Running),
		string(TimeOffReasonIconType_Sunny),
		string(TimeOffReasonIconType_TrafficCone),
		string(TimeOffReasonIconType_Umbrella),
		string(TimeOffReasonIconType_Weather),
	}
}

func (s *TimeOffReasonIconType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeOffReasonIconType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeOffReasonIconType(input string) (*TimeOffReasonIconType, error) {
	vals := map[string]TimeOffReasonIconType{
		"cake":        TimeOffReasonIconType_Cake,
		"calendar":    TimeOffReasonIconType_Calendar,
		"car":         TimeOffReasonIconType_Car,
		"clock":       TimeOffReasonIconType_Clock,
		"cup":         TimeOffReasonIconType_Cup,
		"doctor":      TimeOffReasonIconType_Doctor,
		"dog":         TimeOffReasonIconType_Dog,
		"firstaid":    TimeOffReasonIconType_FirstAid,
		"globe":       TimeOffReasonIconType_Globe,
		"juryduty":    TimeOffReasonIconType_JuryDuty,
		"none":        TimeOffReasonIconType_None,
		"notworking":  TimeOffReasonIconType_NotWorking,
		"phone":       TimeOffReasonIconType_Phone,
		"piggybank":   TimeOffReasonIconType_PiggyBank,
		"pin":         TimeOffReasonIconType_Pin,
		"plane":       TimeOffReasonIconType_Plane,
		"running":     TimeOffReasonIconType_Running,
		"sunny":       TimeOffReasonIconType_Sunny,
		"trafficcone": TimeOffReasonIconType_TrafficCone,
		"umbrella":    TimeOffReasonIconType_Umbrella,
		"weather":     TimeOffReasonIconType_Weather,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeOffReasonIconType(input)
	return &out, nil
}
