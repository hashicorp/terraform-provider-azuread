package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseOperationStatus string

const (
	SecurityCaseOperationStatus_Failed             SecurityCaseOperationStatus = "failed"
	SecurityCaseOperationStatus_NotStarted         SecurityCaseOperationStatus = "notStarted"
	SecurityCaseOperationStatus_PartiallySucceeded SecurityCaseOperationStatus = "partiallySucceeded"
	SecurityCaseOperationStatus_Running            SecurityCaseOperationStatus = "running"
	SecurityCaseOperationStatus_SubmissionFailed   SecurityCaseOperationStatus = "submissionFailed"
	SecurityCaseOperationStatus_Succeeded          SecurityCaseOperationStatus = "succeeded"
)

func PossibleValuesForSecurityCaseOperationStatus() []string {
	return []string{
		string(SecurityCaseOperationStatus_Failed),
		string(SecurityCaseOperationStatus_NotStarted),
		string(SecurityCaseOperationStatus_PartiallySucceeded),
		string(SecurityCaseOperationStatus_Running),
		string(SecurityCaseOperationStatus_SubmissionFailed),
		string(SecurityCaseOperationStatus_Succeeded),
	}
}

func (s *SecurityCaseOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCaseOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCaseOperationStatus(input string) (*SecurityCaseOperationStatus, error) {
	vals := map[string]SecurityCaseOperationStatus{
		"failed":             SecurityCaseOperationStatus_Failed,
		"notstarted":         SecurityCaseOperationStatus_NotStarted,
		"partiallysucceeded": SecurityCaseOperationStatus_PartiallySucceeded,
		"running":            SecurityCaseOperationStatus_Running,
		"submissionfailed":   SecurityCaseOperationStatus_SubmissionFailed,
		"succeeded":          SecurityCaseOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseOperationStatus(input)
	return &out, nil
}
