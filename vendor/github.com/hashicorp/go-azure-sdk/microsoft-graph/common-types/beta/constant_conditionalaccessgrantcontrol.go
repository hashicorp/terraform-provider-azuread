package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessGrantControl string

const (
	ConditionalAccessGrantControl_ApprovedApplication  ConditionalAccessGrantControl = "approvedApplication"
	ConditionalAccessGrantControl_Block                ConditionalAccessGrantControl = "block"
	ConditionalAccessGrantControl_CompliantApplication ConditionalAccessGrantControl = "compliantApplication"
	ConditionalAccessGrantControl_CompliantDevice      ConditionalAccessGrantControl = "compliantDevice"
	ConditionalAccessGrantControl_DomainJoinedDevice   ConditionalAccessGrantControl = "domainJoinedDevice"
	ConditionalAccessGrantControl_Mfa                  ConditionalAccessGrantControl = "mfa"
	ConditionalAccessGrantControl_PasswordChange       ConditionalAccessGrantControl = "passwordChange"
)

func PossibleValuesForConditionalAccessGrantControl() []string {
	return []string{
		string(ConditionalAccessGrantControl_ApprovedApplication),
		string(ConditionalAccessGrantControl_Block),
		string(ConditionalAccessGrantControl_CompliantApplication),
		string(ConditionalAccessGrantControl_CompliantDevice),
		string(ConditionalAccessGrantControl_DomainJoinedDevice),
		string(ConditionalAccessGrantControl_Mfa),
		string(ConditionalAccessGrantControl_PasswordChange),
	}
}

func (s *ConditionalAccessGrantControl) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessGrantControl(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessGrantControl(input string) (*ConditionalAccessGrantControl, error) {
	vals := map[string]ConditionalAccessGrantControl{
		"approvedapplication":  ConditionalAccessGrantControl_ApprovedApplication,
		"block":                ConditionalAccessGrantControl_Block,
		"compliantapplication": ConditionalAccessGrantControl_CompliantApplication,
		"compliantdevice":      ConditionalAccessGrantControl_CompliantDevice,
		"domainjoineddevice":   ConditionalAccessGrantControl_DomainJoinedDevice,
		"mfa":                  ConditionalAccessGrantControl_Mfa,
		"passwordchange":       ConditionalAccessGrantControl_PasswordChange,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessGrantControl(input)
	return &out, nil
}
