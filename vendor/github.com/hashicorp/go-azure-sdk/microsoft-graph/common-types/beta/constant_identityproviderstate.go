package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProviderState string

const (
	IdentityProviderState_Disabled IdentityProviderState = "disabled"
	IdentityProviderState_Enabled  IdentityProviderState = "enabled"
)

func PossibleValuesForIdentityProviderState() []string {
	return []string{
		string(IdentityProviderState_Disabled),
		string(IdentityProviderState_Enabled),
	}
}

func (s *IdentityProviderState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityProviderState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityProviderState(input string) (*IdentityProviderState, error) {
	vals := map[string]IdentityProviderState{
		"disabled": IdentityProviderState_Disabled,
		"enabled":  IdentityProviderState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityProviderState(input)
	return &out, nil
}
