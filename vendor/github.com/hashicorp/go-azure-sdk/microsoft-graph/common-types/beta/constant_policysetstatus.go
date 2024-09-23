package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicySetStatus string

const (
	PolicySetStatus_Error          PolicySetStatus = "error"
	PolicySetStatus_NotAssigned    PolicySetStatus = "notAssigned"
	PolicySetStatus_PartialSuccess PolicySetStatus = "partialSuccess"
	PolicySetStatus_Success        PolicySetStatus = "success"
	PolicySetStatus_Unknown        PolicySetStatus = "unknown"
	PolicySetStatus_Validating     PolicySetStatus = "validating"
)

func PossibleValuesForPolicySetStatus() []string {
	return []string{
		string(PolicySetStatus_Error),
		string(PolicySetStatus_NotAssigned),
		string(PolicySetStatus_PartialSuccess),
		string(PolicySetStatus_Success),
		string(PolicySetStatus_Unknown),
		string(PolicySetStatus_Validating),
	}
}

func (s *PolicySetStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicySetStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicySetStatus(input string) (*PolicySetStatus, error) {
	vals := map[string]PolicySetStatus{
		"error":          PolicySetStatus_Error,
		"notassigned":    PolicySetStatus_NotAssigned,
		"partialsuccess": PolicySetStatus_PartialSuccess,
		"success":        PolicySetStatus_Success,
		"unknown":        PolicySetStatus_Unknown,
		"validating":     PolicySetStatus_Validating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicySetStatus(input)
	return &out, nil
}
