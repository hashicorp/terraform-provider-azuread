package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditActivityResult string

const (
	CloudPCAuditActivityResult_ClientError CloudPCAuditActivityResult = "clientError"
	CloudPCAuditActivityResult_Failure     CloudPCAuditActivityResult = "failure"
	CloudPCAuditActivityResult_Success     CloudPCAuditActivityResult = "success"
	CloudPCAuditActivityResult_Timeout     CloudPCAuditActivityResult = "timeout"
)

func PossibleValuesForCloudPCAuditActivityResult() []string {
	return []string{
		string(CloudPCAuditActivityResult_ClientError),
		string(CloudPCAuditActivityResult_Failure),
		string(CloudPCAuditActivityResult_Success),
		string(CloudPCAuditActivityResult_Timeout),
	}
}

func (s *CloudPCAuditActivityResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCAuditActivityResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCAuditActivityResult(input string) (*CloudPCAuditActivityResult, error) {
	vals := map[string]CloudPCAuditActivityResult{
		"clienterror": CloudPCAuditActivityResult_ClientError,
		"failure":     CloudPCAuditActivityResult_Failure,
		"success":     CloudPCAuditActivityResult_Success,
		"timeout":     CloudPCAuditActivityResult_Timeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCAuditActivityResult(input)
	return &out, nil
}
