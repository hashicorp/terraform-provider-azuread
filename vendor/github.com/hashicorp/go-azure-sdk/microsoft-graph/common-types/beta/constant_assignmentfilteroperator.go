package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterOperator string

const (
	AssignmentFilterOperator_Contains            AssignmentFilterOperator = "contains"
	AssignmentFilterOperator_EndsWith            AssignmentFilterOperator = "endsWith"
	AssignmentFilterOperator_Equals              AssignmentFilterOperator = "equals"
	AssignmentFilterOperator_GreaterThan         AssignmentFilterOperator = "greaterThan"
	AssignmentFilterOperator_GreaterThanOrEquals AssignmentFilterOperator = "greaterThanOrEquals"
	AssignmentFilterOperator_In                  AssignmentFilterOperator = "in"
	AssignmentFilterOperator_LessThan            AssignmentFilterOperator = "lessThan"
	AssignmentFilterOperator_LessThanOrEquals    AssignmentFilterOperator = "lessThanOrEquals"
	AssignmentFilterOperator_NotContains         AssignmentFilterOperator = "notContains"
	AssignmentFilterOperator_NotEndsWith         AssignmentFilterOperator = "notEndsWith"
	AssignmentFilterOperator_NotEquals           AssignmentFilterOperator = "notEquals"
	AssignmentFilterOperator_NotIn               AssignmentFilterOperator = "notIn"
	AssignmentFilterOperator_NotSet              AssignmentFilterOperator = "notSet"
	AssignmentFilterOperator_NotStartsWith       AssignmentFilterOperator = "notStartsWith"
	AssignmentFilterOperator_StartsWith          AssignmentFilterOperator = "startsWith"
)

func PossibleValuesForAssignmentFilterOperator() []string {
	return []string{
		string(AssignmentFilterOperator_Contains),
		string(AssignmentFilterOperator_EndsWith),
		string(AssignmentFilterOperator_Equals),
		string(AssignmentFilterOperator_GreaterThan),
		string(AssignmentFilterOperator_GreaterThanOrEquals),
		string(AssignmentFilterOperator_In),
		string(AssignmentFilterOperator_LessThan),
		string(AssignmentFilterOperator_LessThanOrEquals),
		string(AssignmentFilterOperator_NotContains),
		string(AssignmentFilterOperator_NotEndsWith),
		string(AssignmentFilterOperator_NotEquals),
		string(AssignmentFilterOperator_NotIn),
		string(AssignmentFilterOperator_NotSet),
		string(AssignmentFilterOperator_NotStartsWith),
		string(AssignmentFilterOperator_StartsWith),
	}
}

func (s *AssignmentFilterOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterOperator(input string) (*AssignmentFilterOperator, error) {
	vals := map[string]AssignmentFilterOperator{
		"contains":            AssignmentFilterOperator_Contains,
		"endswith":            AssignmentFilterOperator_EndsWith,
		"equals":              AssignmentFilterOperator_Equals,
		"greaterthan":         AssignmentFilterOperator_GreaterThan,
		"greaterthanorequals": AssignmentFilterOperator_GreaterThanOrEquals,
		"in":                  AssignmentFilterOperator_In,
		"lessthan":            AssignmentFilterOperator_LessThan,
		"lessthanorequals":    AssignmentFilterOperator_LessThanOrEquals,
		"notcontains":         AssignmentFilterOperator_NotContains,
		"notendswith":         AssignmentFilterOperator_NotEndsWith,
		"notequals":           AssignmentFilterOperator_NotEquals,
		"notin":               AssignmentFilterOperator_NotIn,
		"notset":              AssignmentFilterOperator_NotSet,
		"notstartswith":       AssignmentFilterOperator_NotStartsWith,
		"startswith":          AssignmentFilterOperator_StartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterOperator(input)
	return &out, nil
}
