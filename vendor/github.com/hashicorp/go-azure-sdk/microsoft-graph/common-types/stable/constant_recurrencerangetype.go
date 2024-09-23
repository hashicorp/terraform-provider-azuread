package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrenceRangeType string

const (
	RecurrenceRangeType_EndDate  RecurrenceRangeType = "endDate"
	RecurrenceRangeType_NoEnd    RecurrenceRangeType = "noEnd"
	RecurrenceRangeType_Numbered RecurrenceRangeType = "numbered"
)

func PossibleValuesForRecurrenceRangeType() []string {
	return []string{
		string(RecurrenceRangeType_EndDate),
		string(RecurrenceRangeType_NoEnd),
		string(RecurrenceRangeType_Numbered),
	}
}

func (s *RecurrenceRangeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecurrenceRangeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecurrenceRangeType(input string) (*RecurrenceRangeType, error) {
	vals := map[string]RecurrenceRangeType{
		"enddate":  RecurrenceRangeType_EndDate,
		"noend":    RecurrenceRangeType_NoEnd,
		"numbered": RecurrenceRangeType_Numbered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceRangeType(input)
	return &out, nil
}
