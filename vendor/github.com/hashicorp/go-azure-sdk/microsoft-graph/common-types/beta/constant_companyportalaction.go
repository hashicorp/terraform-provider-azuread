package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalAction string

const (
	CompanyPortalAction_Remove  CompanyPortalAction = "remove"
	CompanyPortalAction_Reset   CompanyPortalAction = "reset"
	CompanyPortalAction_Unknown CompanyPortalAction = "unknown"
)

func PossibleValuesForCompanyPortalAction() []string {
	return []string{
		string(CompanyPortalAction_Remove),
		string(CompanyPortalAction_Reset),
		string(CompanyPortalAction_Unknown),
	}
}

func (s *CompanyPortalAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompanyPortalAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompanyPortalAction(input string) (*CompanyPortalAction, error) {
	vals := map[string]CompanyPortalAction{
		"remove":  CompanyPortalAction_Remove,
		"reset":   CompanyPortalAction_Reset,
		"unknown": CompanyPortalAction_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompanyPortalAction(input)
	return &out, nil
}
