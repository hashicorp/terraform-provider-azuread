package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowExecutionType string

const (
	IdentityGovernanceWorkflowExecutionType_OnDemand  IdentityGovernanceWorkflowExecutionType = "onDemand"
	IdentityGovernanceWorkflowExecutionType_Scheduled IdentityGovernanceWorkflowExecutionType = "scheduled"
)

func PossibleValuesForIdentityGovernanceWorkflowExecutionType() []string {
	return []string{
		string(IdentityGovernanceWorkflowExecutionType_OnDemand),
		string(IdentityGovernanceWorkflowExecutionType_Scheduled),
	}
}

func (s *IdentityGovernanceWorkflowExecutionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowExecutionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowExecutionType(input string) (*IdentityGovernanceWorkflowExecutionType, error) {
	vals := map[string]IdentityGovernanceWorkflowExecutionType{
		"ondemand":  IdentityGovernanceWorkflowExecutionType_OnDemand,
		"scheduled": IdentityGovernanceWorkflowExecutionType_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowExecutionType(input)
	return &out, nil
}
