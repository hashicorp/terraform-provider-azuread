package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnualEventType string

const (
	PersonAnnualEventType_Birthday PersonAnnualEventType = "birthday"
	PersonAnnualEventType_Other    PersonAnnualEventType = "other"
	PersonAnnualEventType_Wedding  PersonAnnualEventType = "wedding"
	PersonAnnualEventType_Work     PersonAnnualEventType = "work"
)

func PossibleValuesForPersonAnnualEventType() []string {
	return []string{
		string(PersonAnnualEventType_Birthday),
		string(PersonAnnualEventType_Other),
		string(PersonAnnualEventType_Wedding),
		string(PersonAnnualEventType_Work),
	}
}

func (s *PersonAnnualEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePersonAnnualEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePersonAnnualEventType(input string) (*PersonAnnualEventType, error) {
	vals := map[string]PersonAnnualEventType{
		"birthday": PersonAnnualEventType_Birthday,
		"other":    PersonAnnualEventType_Other,
		"wedding":  PersonAnnualEventType_Wedding,
		"work":     PersonAnnualEventType_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnualEventType(input)
	return &out, nil
}
