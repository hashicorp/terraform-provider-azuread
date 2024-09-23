package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseOperationStatus string

const (
	EdiscoveryCaseOperationStatus_Failed             EdiscoveryCaseOperationStatus = "failed"
	EdiscoveryCaseOperationStatus_NotStarted         EdiscoveryCaseOperationStatus = "notStarted"
	EdiscoveryCaseOperationStatus_PartiallySucceeded EdiscoveryCaseOperationStatus = "partiallySucceeded"
	EdiscoveryCaseOperationStatus_Running            EdiscoveryCaseOperationStatus = "running"
	EdiscoveryCaseOperationStatus_SubmissionFailed   EdiscoveryCaseOperationStatus = "submissionFailed"
	EdiscoveryCaseOperationStatus_Succeeded          EdiscoveryCaseOperationStatus = "succeeded"
)

func PossibleValuesForEdiscoveryCaseOperationStatus() []string {
	return []string{
		string(EdiscoveryCaseOperationStatus_Failed),
		string(EdiscoveryCaseOperationStatus_NotStarted),
		string(EdiscoveryCaseOperationStatus_PartiallySucceeded),
		string(EdiscoveryCaseOperationStatus_Running),
		string(EdiscoveryCaseOperationStatus_SubmissionFailed),
		string(EdiscoveryCaseOperationStatus_Succeeded),
	}
}

func (s *EdiscoveryCaseOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseOperationStatus(input string) (*EdiscoveryCaseOperationStatus, error) {
	vals := map[string]EdiscoveryCaseOperationStatus{
		"failed":             EdiscoveryCaseOperationStatus_Failed,
		"notstarted":         EdiscoveryCaseOperationStatus_NotStarted,
		"partiallysucceeded": EdiscoveryCaseOperationStatus_PartiallySucceeded,
		"running":            EdiscoveryCaseOperationStatus_Running,
		"submissionfailed":   EdiscoveryCaseOperationStatus_SubmissionFailed,
		"succeeded":          EdiscoveryCaseOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseOperationStatus(input)
	return &out, nil
}
