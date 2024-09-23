package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Operator string

const (
	Operator_AllOf            Operator = "allOf"
	Operator_And              Operator = "and"
	Operator_BeginsWith       Operator = "beginsWith"
	Operator_Between          Operator = "between"
	Operator_Contains         Operator = "contains"
	Operator_DayTimeBetween   Operator = "dayTimeBetween"
	Operator_EndsWith         Operator = "endsWith"
	Operator_ExcludesAll      Operator = "excludesAll"
	Operator_GreaterEquals    Operator = "greaterEquals"
	Operator_GreaterThan      Operator = "greaterThan"
	Operator_IsEquals         Operator = "isEquals"
	Operator_LessEquals       Operator = "lessEquals"
	Operator_LessThan         Operator = "lessThan"
	Operator_None             Operator = "none"
	Operator_NoneOf           Operator = "noneOf"
	Operator_NotBeginsWith    Operator = "notBeginsWith"
	Operator_NotBetween       Operator = "notBetween"
	Operator_NotContains      Operator = "notContains"
	Operator_NotEndsWith      Operator = "notEndsWith"
	Operator_NotEquals        Operator = "notEquals"
	Operator_OneOf            Operator = "oneOf"
	Operator_Or               Operator = "or"
	Operator_OrderedSetEquals Operator = "orderedSetEquals"
	Operator_SetEquals        Operator = "setEquals"
	Operator_SubsetOf         Operator = "subsetOf"
)

func PossibleValuesForOperator() []string {
	return []string{
		string(Operator_AllOf),
		string(Operator_And),
		string(Operator_BeginsWith),
		string(Operator_Between),
		string(Operator_Contains),
		string(Operator_DayTimeBetween),
		string(Operator_EndsWith),
		string(Operator_ExcludesAll),
		string(Operator_GreaterEquals),
		string(Operator_GreaterThan),
		string(Operator_IsEquals),
		string(Operator_LessEquals),
		string(Operator_LessThan),
		string(Operator_None),
		string(Operator_NoneOf),
		string(Operator_NotBeginsWith),
		string(Operator_NotBetween),
		string(Operator_NotContains),
		string(Operator_NotEndsWith),
		string(Operator_NotEquals),
		string(Operator_OneOf),
		string(Operator_Or),
		string(Operator_OrderedSetEquals),
		string(Operator_SetEquals),
		string(Operator_SubsetOf),
	}
}

func (s *Operator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperator(input string) (*Operator, error) {
	vals := map[string]Operator{
		"allof":            Operator_AllOf,
		"and":              Operator_And,
		"beginswith":       Operator_BeginsWith,
		"between":          Operator_Between,
		"contains":         Operator_Contains,
		"daytimebetween":   Operator_DayTimeBetween,
		"endswith":         Operator_EndsWith,
		"excludesall":      Operator_ExcludesAll,
		"greaterequals":    Operator_GreaterEquals,
		"greaterthan":      Operator_GreaterThan,
		"isequals":         Operator_IsEquals,
		"lessequals":       Operator_LessEquals,
		"lessthan":         Operator_LessThan,
		"none":             Operator_None,
		"noneof":           Operator_NoneOf,
		"notbeginswith":    Operator_NotBeginsWith,
		"notbetween":       Operator_NotBetween,
		"notcontains":      Operator_NotContains,
		"notendswith":      Operator_NotEndsWith,
		"notequals":        Operator_NotEquals,
		"oneof":            Operator_OneOf,
		"or":               Operator_Or,
		"orderedsetequals": Operator_OrderedSetEquals,
		"setequals":        Operator_SetEquals,
		"subsetof":         Operator_SubsetOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Operator(input)
	return &out, nil
}
