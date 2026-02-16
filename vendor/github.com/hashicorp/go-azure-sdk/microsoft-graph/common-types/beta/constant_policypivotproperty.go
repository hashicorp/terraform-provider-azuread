package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyPivotProperty string

const (
	PolicyPivotProperty_Activity PolicyPivotProperty = "activity"
	PolicyPivotProperty_Location PolicyPivotProperty = "location"
	PolicyPivotProperty_None     PolicyPivotProperty = "none"
)

func PossibleValuesForPolicyPivotProperty() []string {
	return []string{
		string(PolicyPivotProperty_Activity),
		string(PolicyPivotProperty_Location),
		string(PolicyPivotProperty_None),
	}
}

func (s *PolicyPivotProperty) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyPivotProperty(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyPivotProperty(input string) (*PolicyPivotProperty, error) {
	vals := map[string]PolicyPivotProperty{
		"activity": PolicyPivotProperty_Activity,
		"location": PolicyPivotProperty_Location,
		"none":     PolicyPivotProperty_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyPivotProperty(input)
	return &out, nil
}
