package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowBlockListAction string

const (
	SecurityTenantAllowBlockListAction_Allow SecurityTenantAllowBlockListAction = "allow"
	SecurityTenantAllowBlockListAction_Block SecurityTenantAllowBlockListAction = "block"
)

func PossibleValuesForSecurityTenantAllowBlockListAction() []string {
	return []string{
		string(SecurityTenantAllowBlockListAction_Allow),
		string(SecurityTenantAllowBlockListAction_Block),
	}
}

func (s *SecurityTenantAllowBlockListAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTenantAllowBlockListAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTenantAllowBlockListAction(input string) (*SecurityTenantAllowBlockListAction, error) {
	vals := map[string]SecurityTenantAllowBlockListAction{
		"allow": SecurityTenantAllowBlockListAction_Allow,
		"block": SecurityTenantAllowBlockListAction_Block,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowBlockListAction(input)
	return &out, nil
}
