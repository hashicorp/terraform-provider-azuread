package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessExternalTenantsMembershipKind string

const (
	ConditionalAccessExternalTenantsMembershipKind_All        ConditionalAccessExternalTenantsMembershipKind = "all"
	ConditionalAccessExternalTenantsMembershipKind_Enumerated ConditionalAccessExternalTenantsMembershipKind = "enumerated"
)

func PossibleValuesForConditionalAccessExternalTenantsMembershipKind() []string {
	return []string{
		string(ConditionalAccessExternalTenantsMembershipKind_All),
		string(ConditionalAccessExternalTenantsMembershipKind_Enumerated),
	}
}

func (s *ConditionalAccessExternalTenantsMembershipKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessExternalTenantsMembershipKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessExternalTenantsMembershipKind(input string) (*ConditionalAccessExternalTenantsMembershipKind, error) {
	vals := map[string]ConditionalAccessExternalTenantsMembershipKind{
		"all":        ConditionalAccessExternalTenantsMembershipKind_All,
		"enumerated": ConditionalAccessExternalTenantsMembershipKind_Enumerated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessExternalTenantsMembershipKind(input)
	return &out, nil
}
