package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHuntingRuleErrorCode string

const (
	SecurityHuntingRuleErrorCode_AlertCreationFailed      SecurityHuntingRuleErrorCode = "alertCreationFailed"
	SecurityHuntingRuleErrorCode_AlertReportNotFound      SecurityHuntingRuleErrorCode = "alertReportNotFound"
	SecurityHuntingRuleErrorCode_NoImpactedEntity         SecurityHuntingRuleErrorCode = "noImpactedEntity"
	SecurityHuntingRuleErrorCode_PartialRowsFailed        SecurityHuntingRuleErrorCode = "partialRowsFailed"
	SecurityHuntingRuleErrorCode_QueryExceededResultSize  SecurityHuntingRuleErrorCode = "queryExceededResultSize"
	SecurityHuntingRuleErrorCode_QueryExecutionFailed     SecurityHuntingRuleErrorCode = "queryExecutionFailed"
	SecurityHuntingRuleErrorCode_QueryExecutionThrottling SecurityHuntingRuleErrorCode = "queryExecutionThrottling"
	SecurityHuntingRuleErrorCode_QueryLimitsExceeded      SecurityHuntingRuleErrorCode = "queryLimitsExceeded"
	SecurityHuntingRuleErrorCode_QueryTimeout             SecurityHuntingRuleErrorCode = "queryTimeout"
)

func PossibleValuesForSecurityHuntingRuleErrorCode() []string {
	return []string{
		string(SecurityHuntingRuleErrorCode_AlertCreationFailed),
		string(SecurityHuntingRuleErrorCode_AlertReportNotFound),
		string(SecurityHuntingRuleErrorCode_NoImpactedEntity),
		string(SecurityHuntingRuleErrorCode_PartialRowsFailed),
		string(SecurityHuntingRuleErrorCode_QueryExceededResultSize),
		string(SecurityHuntingRuleErrorCode_QueryExecutionFailed),
		string(SecurityHuntingRuleErrorCode_QueryExecutionThrottling),
		string(SecurityHuntingRuleErrorCode_QueryLimitsExceeded),
		string(SecurityHuntingRuleErrorCode_QueryTimeout),
	}
}

func (s *SecurityHuntingRuleErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHuntingRuleErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHuntingRuleErrorCode(input string) (*SecurityHuntingRuleErrorCode, error) {
	vals := map[string]SecurityHuntingRuleErrorCode{
		"alertcreationfailed":      SecurityHuntingRuleErrorCode_AlertCreationFailed,
		"alertreportnotfound":      SecurityHuntingRuleErrorCode_AlertReportNotFound,
		"noimpactedentity":         SecurityHuntingRuleErrorCode_NoImpactedEntity,
		"partialrowsfailed":        SecurityHuntingRuleErrorCode_PartialRowsFailed,
		"queryexceededresultsize":  SecurityHuntingRuleErrorCode_QueryExceededResultSize,
		"queryexecutionfailed":     SecurityHuntingRuleErrorCode_QueryExecutionFailed,
		"queryexecutionthrottling": SecurityHuntingRuleErrorCode_QueryExecutionThrottling,
		"querylimitsexceeded":      SecurityHuntingRuleErrorCode_QueryLimitsExceeded,
		"querytimeout":             SecurityHuntingRuleErrorCode_QueryTimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHuntingRuleErrorCode(input)
	return &out, nil
}
