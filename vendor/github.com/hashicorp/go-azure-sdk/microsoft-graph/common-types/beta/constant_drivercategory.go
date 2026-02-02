package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriverCategory string

const (
	DriverCategory_Other              DriverCategory = "other"
	DriverCategory_PreviouslyApproved DriverCategory = "previouslyApproved"
	DriverCategory_Recommended        DriverCategory = "recommended"
)

func PossibleValuesForDriverCategory() []string {
	return []string{
		string(DriverCategory_Other),
		string(DriverCategory_PreviouslyApproved),
		string(DriverCategory_Recommended),
	}
}

func (s *DriverCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDriverCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDriverCategory(input string) (*DriverCategory, error) {
	vals := map[string]DriverCategory{
		"other":              DriverCategory_Other,
		"previouslyapproved": DriverCategory_PreviouslyApproved,
		"recommended":        DriverCategory_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriverCategory(input)
	return &out, nil
}
