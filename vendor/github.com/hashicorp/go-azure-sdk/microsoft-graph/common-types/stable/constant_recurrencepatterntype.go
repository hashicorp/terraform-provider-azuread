package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternType string

const (
	RecurrencePatternType_AbsoluteMonthly RecurrencePatternType = "absoluteMonthly"
	RecurrencePatternType_AbsoluteYearly  RecurrencePatternType = "absoluteYearly"
	RecurrencePatternType_Daily           RecurrencePatternType = "daily"
	RecurrencePatternType_RelativeMonthly RecurrencePatternType = "relativeMonthly"
	RecurrencePatternType_RelativeYearly  RecurrencePatternType = "relativeYearly"
	RecurrencePatternType_Weekly          RecurrencePatternType = "weekly"
)

func PossibleValuesForRecurrencePatternType() []string {
	return []string{
		string(RecurrencePatternType_AbsoluteMonthly),
		string(RecurrencePatternType_AbsoluteYearly),
		string(RecurrencePatternType_Daily),
		string(RecurrencePatternType_RelativeMonthly),
		string(RecurrencePatternType_RelativeYearly),
		string(RecurrencePatternType_Weekly),
	}
}

func (s *RecurrencePatternType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrencePatternType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrencePatternType(input string) (*RecurrencePatternType, error) {
	vals := map[string]RecurrencePatternType{
		"absolutemonthly": RecurrencePatternType_AbsoluteMonthly,
		"absoluteyearly":  RecurrencePatternType_AbsoluteYearly,
		"daily":           RecurrencePatternType_Daily,
		"relativemonthly": RecurrencePatternType_RelativeMonthly,
		"relativeyearly":  RecurrencePatternType_RelativeYearly,
		"weekly":          RecurrencePatternType_Weekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrencePatternType(input)
	return &out, nil
}
