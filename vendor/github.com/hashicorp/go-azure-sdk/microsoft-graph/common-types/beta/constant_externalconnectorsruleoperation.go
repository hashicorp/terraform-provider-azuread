package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsRuleOperation string

const (
	ExternalConnectorsRuleOperation_Contains    ExternalConnectorsRuleOperation = "contains"
	ExternalConnectorsRuleOperation_Equals      ExternalConnectorsRuleOperation = "equals"
	ExternalConnectorsRuleOperation_GreaterThan ExternalConnectorsRuleOperation = "greaterThan"
	ExternalConnectorsRuleOperation_LessThan    ExternalConnectorsRuleOperation = "lessThan"
	ExternalConnectorsRuleOperation_NotContains ExternalConnectorsRuleOperation = "notContains"
	ExternalConnectorsRuleOperation_NotEquals   ExternalConnectorsRuleOperation = "notEquals"
	ExternalConnectorsRuleOperation_Null        ExternalConnectorsRuleOperation = "null"
	ExternalConnectorsRuleOperation_StartsWith  ExternalConnectorsRuleOperation = "startsWith"
)

func PossibleValuesForExternalConnectorsRuleOperation() []string {
	return []string{
		string(ExternalConnectorsRuleOperation_Contains),
		string(ExternalConnectorsRuleOperation_Equals),
		string(ExternalConnectorsRuleOperation_GreaterThan),
		string(ExternalConnectorsRuleOperation_LessThan),
		string(ExternalConnectorsRuleOperation_NotContains),
		string(ExternalConnectorsRuleOperation_NotEquals),
		string(ExternalConnectorsRuleOperation_Null),
		string(ExternalConnectorsRuleOperation_StartsWith),
	}
}

func (s *ExternalConnectorsRuleOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsRuleOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsRuleOperation(input string) (*ExternalConnectorsRuleOperation, error) {
	vals := map[string]ExternalConnectorsRuleOperation{
		"contains":    ExternalConnectorsRuleOperation_Contains,
		"equals":      ExternalConnectorsRuleOperation_Equals,
		"greaterthan": ExternalConnectorsRuleOperation_GreaterThan,
		"lessthan":    ExternalConnectorsRuleOperation_LessThan,
		"notcontains": ExternalConnectorsRuleOperation_NotContains,
		"notequals":   ExternalConnectorsRuleOperation_NotEquals,
		"null":        ExternalConnectorsRuleOperation_Null,
		"startswith":  ExternalConnectorsRuleOperation_StartsWith,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsRuleOperation(input)
	return &out, nil
}
