package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceWorkflowTriggerTimeBasedAttribute string

const (
	IdentityGovernanceWorkflowTriggerTimeBasedAttribute_CreatedDateTime       IdentityGovernanceWorkflowTriggerTimeBasedAttribute = "createdDateTime"
	IdentityGovernanceWorkflowTriggerTimeBasedAttribute_EmployeeHireDate      IdentityGovernanceWorkflowTriggerTimeBasedAttribute = "employeeHireDate"
	IdentityGovernanceWorkflowTriggerTimeBasedAttribute_EmployeeLeaveDateTime IdentityGovernanceWorkflowTriggerTimeBasedAttribute = "employeeLeaveDateTime"
)

func PossibleValuesForIdentityGovernanceWorkflowTriggerTimeBasedAttribute() []string {
	return []string{
		string(IdentityGovernanceWorkflowTriggerTimeBasedAttribute_CreatedDateTime),
		string(IdentityGovernanceWorkflowTriggerTimeBasedAttribute_EmployeeHireDate),
		string(IdentityGovernanceWorkflowTriggerTimeBasedAttribute_EmployeeLeaveDateTime),
	}
}

func (s *IdentityGovernanceWorkflowTriggerTimeBasedAttribute) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceWorkflowTriggerTimeBasedAttribute(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceWorkflowTriggerTimeBasedAttribute(input string) (*IdentityGovernanceWorkflowTriggerTimeBasedAttribute, error) {
	vals := map[string]IdentityGovernanceWorkflowTriggerTimeBasedAttribute{
		"createddatetime":       IdentityGovernanceWorkflowTriggerTimeBasedAttribute_CreatedDateTime,
		"employeehiredate":      IdentityGovernanceWorkflowTriggerTimeBasedAttribute_EmployeeHireDate,
		"employeeleavedatetime": IdentityGovernanceWorkflowTriggerTimeBasedAttribute_EmployeeLeaveDateTime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceWorkflowTriggerTimeBasedAttribute(input)
	return &out, nil
}
