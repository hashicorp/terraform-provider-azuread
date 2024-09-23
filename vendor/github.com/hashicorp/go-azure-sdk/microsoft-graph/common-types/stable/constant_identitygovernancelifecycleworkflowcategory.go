package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceLifecycleWorkflowCategory string

const (
	IdentityGovernanceLifecycleWorkflowCategory_Joiner IdentityGovernanceLifecycleWorkflowCategory = "joiner"
	IdentityGovernanceLifecycleWorkflowCategory_Leaver IdentityGovernanceLifecycleWorkflowCategory = "leaver"
	IdentityGovernanceLifecycleWorkflowCategory_Mover  IdentityGovernanceLifecycleWorkflowCategory = "mover"
)

func PossibleValuesForIdentityGovernanceLifecycleWorkflowCategory() []string {
	return []string{
		string(IdentityGovernanceLifecycleWorkflowCategory_Joiner),
		string(IdentityGovernanceLifecycleWorkflowCategory_Leaver),
		string(IdentityGovernanceLifecycleWorkflowCategory_Mover),
	}
}

func (s *IdentityGovernanceLifecycleWorkflowCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceLifecycleWorkflowCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceLifecycleWorkflowCategory(input string) (*IdentityGovernanceLifecycleWorkflowCategory, error) {
	vals := map[string]IdentityGovernanceLifecycleWorkflowCategory{
		"joiner": IdentityGovernanceLifecycleWorkflowCategory_Joiner,
		"leaver": IdentityGovernanceLifecycleWorkflowCategory_Leaver,
		"mover":  IdentityGovernanceLifecycleWorkflowCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceLifecycleWorkflowCategory(input)
	return &out, nil
}
