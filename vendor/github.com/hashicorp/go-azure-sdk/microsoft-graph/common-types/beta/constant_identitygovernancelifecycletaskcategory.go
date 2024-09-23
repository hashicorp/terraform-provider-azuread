package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceLifecycleTaskCategory string

const (
	IdentityGovernanceLifecycleTaskCategory_Joiner IdentityGovernanceLifecycleTaskCategory = "joiner"
	IdentityGovernanceLifecycleTaskCategory_Leaver IdentityGovernanceLifecycleTaskCategory = "leaver"
	IdentityGovernanceLifecycleTaskCategory_Mover  IdentityGovernanceLifecycleTaskCategory = "mover"
)

func PossibleValuesForIdentityGovernanceLifecycleTaskCategory() []string {
	return []string{
		string(IdentityGovernanceLifecycleTaskCategory_Joiner),
		string(IdentityGovernanceLifecycleTaskCategory_Leaver),
		string(IdentityGovernanceLifecycleTaskCategory_Mover),
	}
}

func (s *IdentityGovernanceLifecycleTaskCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceLifecycleTaskCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceLifecycleTaskCategory(input string) (*IdentityGovernanceLifecycleTaskCategory, error) {
	vals := map[string]IdentityGovernanceLifecycleTaskCategory{
		"joiner": IdentityGovernanceLifecycleTaskCategory_Joiner,
		"leaver": IdentityGovernanceLifecycleTaskCategory_Leaver,
		"mover":  IdentityGovernanceLifecycleTaskCategory_Mover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceLifecycleTaskCategory(input)
	return &out, nil
}
