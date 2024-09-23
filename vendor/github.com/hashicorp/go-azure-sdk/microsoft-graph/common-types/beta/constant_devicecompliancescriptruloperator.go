package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptRulOperator string

const (
	DeviceComplianceScriptRulOperator_AllOf            DeviceComplianceScriptRulOperator = "allOf"
	DeviceComplianceScriptRulOperator_And              DeviceComplianceScriptRulOperator = "and"
	DeviceComplianceScriptRulOperator_BeginsWith       DeviceComplianceScriptRulOperator = "beginsWith"
	DeviceComplianceScriptRulOperator_Between          DeviceComplianceScriptRulOperator = "between"
	DeviceComplianceScriptRulOperator_Contains         DeviceComplianceScriptRulOperator = "contains"
	DeviceComplianceScriptRulOperator_DayTimeBetween   DeviceComplianceScriptRulOperator = "dayTimeBetween"
	DeviceComplianceScriptRulOperator_EndsWith         DeviceComplianceScriptRulOperator = "endsWith"
	DeviceComplianceScriptRulOperator_ExcludesAll      DeviceComplianceScriptRulOperator = "excludesAll"
	DeviceComplianceScriptRulOperator_GreaterEquals    DeviceComplianceScriptRulOperator = "greaterEquals"
	DeviceComplianceScriptRulOperator_GreaterThan      DeviceComplianceScriptRulOperator = "greaterThan"
	DeviceComplianceScriptRulOperator_IsEquals         DeviceComplianceScriptRulOperator = "isEquals"
	DeviceComplianceScriptRulOperator_LessEquals       DeviceComplianceScriptRulOperator = "lessEquals"
	DeviceComplianceScriptRulOperator_LessThan         DeviceComplianceScriptRulOperator = "lessThan"
	DeviceComplianceScriptRulOperator_None             DeviceComplianceScriptRulOperator = "none"
	DeviceComplianceScriptRulOperator_NoneOf           DeviceComplianceScriptRulOperator = "noneOf"
	DeviceComplianceScriptRulOperator_NotBeginsWith    DeviceComplianceScriptRulOperator = "notBeginsWith"
	DeviceComplianceScriptRulOperator_NotBetween       DeviceComplianceScriptRulOperator = "notBetween"
	DeviceComplianceScriptRulOperator_NotContains      DeviceComplianceScriptRulOperator = "notContains"
	DeviceComplianceScriptRulOperator_NotEndsWith      DeviceComplianceScriptRulOperator = "notEndsWith"
	DeviceComplianceScriptRulOperator_NotEquals        DeviceComplianceScriptRulOperator = "notEquals"
	DeviceComplianceScriptRulOperator_OneOf            DeviceComplianceScriptRulOperator = "oneOf"
	DeviceComplianceScriptRulOperator_Or               DeviceComplianceScriptRulOperator = "or"
	DeviceComplianceScriptRulOperator_OrderedSetEquals DeviceComplianceScriptRulOperator = "orderedSetEquals"
	DeviceComplianceScriptRulOperator_SetEquals        DeviceComplianceScriptRulOperator = "setEquals"
	DeviceComplianceScriptRulOperator_SubsetOf         DeviceComplianceScriptRulOperator = "subsetOf"
)

func PossibleValuesForDeviceComplianceScriptRulOperator() []string {
	return []string{
		string(DeviceComplianceScriptRulOperator_AllOf),
		string(DeviceComplianceScriptRulOperator_And),
		string(DeviceComplianceScriptRulOperator_BeginsWith),
		string(DeviceComplianceScriptRulOperator_Between),
		string(DeviceComplianceScriptRulOperator_Contains),
		string(DeviceComplianceScriptRulOperator_DayTimeBetween),
		string(DeviceComplianceScriptRulOperator_EndsWith),
		string(DeviceComplianceScriptRulOperator_ExcludesAll),
		string(DeviceComplianceScriptRulOperator_GreaterEquals),
		string(DeviceComplianceScriptRulOperator_GreaterThan),
		string(DeviceComplianceScriptRulOperator_IsEquals),
		string(DeviceComplianceScriptRulOperator_LessEquals),
		string(DeviceComplianceScriptRulOperator_LessThan),
		string(DeviceComplianceScriptRulOperator_None),
		string(DeviceComplianceScriptRulOperator_NoneOf),
		string(DeviceComplianceScriptRulOperator_NotBeginsWith),
		string(DeviceComplianceScriptRulOperator_NotBetween),
		string(DeviceComplianceScriptRulOperator_NotContains),
		string(DeviceComplianceScriptRulOperator_NotEndsWith),
		string(DeviceComplianceScriptRulOperator_NotEquals),
		string(DeviceComplianceScriptRulOperator_OneOf),
		string(DeviceComplianceScriptRulOperator_Or),
		string(DeviceComplianceScriptRulOperator_OrderedSetEquals),
		string(DeviceComplianceScriptRulOperator_SetEquals),
		string(DeviceComplianceScriptRulOperator_SubsetOf),
	}
}

func (s *DeviceComplianceScriptRulOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceScriptRulOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceScriptRulOperator(input string) (*DeviceComplianceScriptRulOperator, error) {
	vals := map[string]DeviceComplianceScriptRulOperator{
		"allof":            DeviceComplianceScriptRulOperator_AllOf,
		"and":              DeviceComplianceScriptRulOperator_And,
		"beginswith":       DeviceComplianceScriptRulOperator_BeginsWith,
		"between":          DeviceComplianceScriptRulOperator_Between,
		"contains":         DeviceComplianceScriptRulOperator_Contains,
		"daytimebetween":   DeviceComplianceScriptRulOperator_DayTimeBetween,
		"endswith":         DeviceComplianceScriptRulOperator_EndsWith,
		"excludesall":      DeviceComplianceScriptRulOperator_ExcludesAll,
		"greaterequals":    DeviceComplianceScriptRulOperator_GreaterEquals,
		"greaterthan":      DeviceComplianceScriptRulOperator_GreaterThan,
		"isequals":         DeviceComplianceScriptRulOperator_IsEquals,
		"lessequals":       DeviceComplianceScriptRulOperator_LessEquals,
		"lessthan":         DeviceComplianceScriptRulOperator_LessThan,
		"none":             DeviceComplianceScriptRulOperator_None,
		"noneof":           DeviceComplianceScriptRulOperator_NoneOf,
		"notbeginswith":    DeviceComplianceScriptRulOperator_NotBeginsWith,
		"notbetween":       DeviceComplianceScriptRulOperator_NotBetween,
		"notcontains":      DeviceComplianceScriptRulOperator_NotContains,
		"notendswith":      DeviceComplianceScriptRulOperator_NotEndsWith,
		"notequals":        DeviceComplianceScriptRulOperator_NotEquals,
		"oneof":            DeviceComplianceScriptRulOperator_OneOf,
		"or":               DeviceComplianceScriptRulOperator_Or,
		"orderedsetequals": DeviceComplianceScriptRulOperator_OrderedSetEquals,
		"setequals":        DeviceComplianceScriptRulOperator_SetEquals,
		"subsetof":         DeviceComplianceScriptRulOperator_SubsetOf,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptRulOperator(input)
	return &out, nil
}
