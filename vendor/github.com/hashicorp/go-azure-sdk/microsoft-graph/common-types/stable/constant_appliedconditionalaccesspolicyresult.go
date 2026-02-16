package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicyResult string

const (
	AppliedConditionalAccessPolicyResult_Failure               AppliedConditionalAccessPolicyResult = "failure"
	AppliedConditionalAccessPolicyResult_NotApplied            AppliedConditionalAccessPolicyResult = "notApplied"
	AppliedConditionalAccessPolicyResult_NotEnabled            AppliedConditionalAccessPolicyResult = "notEnabled"
	AppliedConditionalAccessPolicyResult_ReportOnlyFailure     AppliedConditionalAccessPolicyResult = "reportOnlyFailure"
	AppliedConditionalAccessPolicyResult_ReportOnlyInterrupted AppliedConditionalAccessPolicyResult = "reportOnlyInterrupted"
	AppliedConditionalAccessPolicyResult_ReportOnlyNotApplied  AppliedConditionalAccessPolicyResult = "reportOnlyNotApplied"
	AppliedConditionalAccessPolicyResult_ReportOnlySuccess     AppliedConditionalAccessPolicyResult = "reportOnlySuccess"
	AppliedConditionalAccessPolicyResult_Success               AppliedConditionalAccessPolicyResult = "success"
	AppliedConditionalAccessPolicyResult_Unknown               AppliedConditionalAccessPolicyResult = "unknown"
)

func PossibleValuesForAppliedConditionalAccessPolicyResult() []string {
	return []string{
		string(AppliedConditionalAccessPolicyResult_Failure),
		string(AppliedConditionalAccessPolicyResult_NotApplied),
		string(AppliedConditionalAccessPolicyResult_NotEnabled),
		string(AppliedConditionalAccessPolicyResult_ReportOnlyFailure),
		string(AppliedConditionalAccessPolicyResult_ReportOnlyInterrupted),
		string(AppliedConditionalAccessPolicyResult_ReportOnlyNotApplied),
		string(AppliedConditionalAccessPolicyResult_ReportOnlySuccess),
		string(AppliedConditionalAccessPolicyResult_Success),
		string(AppliedConditionalAccessPolicyResult_Unknown),
	}
}

func (s *AppliedConditionalAccessPolicyResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppliedConditionalAccessPolicyResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppliedConditionalAccessPolicyResult(input string) (*AppliedConditionalAccessPolicyResult, error) {
	vals := map[string]AppliedConditionalAccessPolicyResult{
		"failure":               AppliedConditionalAccessPolicyResult_Failure,
		"notapplied":            AppliedConditionalAccessPolicyResult_NotApplied,
		"notenabled":            AppliedConditionalAccessPolicyResult_NotEnabled,
		"reportonlyfailure":     AppliedConditionalAccessPolicyResult_ReportOnlyFailure,
		"reportonlyinterrupted": AppliedConditionalAccessPolicyResult_ReportOnlyInterrupted,
		"reportonlynotapplied":  AppliedConditionalAccessPolicyResult_ReportOnlyNotApplied,
		"reportonlysuccess":     AppliedConditionalAccessPolicyResult_ReportOnlySuccess,
		"success":               AppliedConditionalAccessPolicyResult_Success,
		"unknown":               AppliedConditionalAccessPolicyResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedConditionalAccessPolicyResult(input)
	return &out, nil
}
