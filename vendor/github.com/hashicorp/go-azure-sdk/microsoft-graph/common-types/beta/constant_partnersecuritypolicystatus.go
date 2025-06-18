package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecurityPolicyStatus string

const (
	PartnerSecurityPolicyStatus_Disabled PartnerSecurityPolicyStatus = "disabled"
	PartnerSecurityPolicyStatus_Enabled  PartnerSecurityPolicyStatus = "enabled"
)

func PossibleValuesForPartnerSecurityPolicyStatus() []string {
	return []string{
		string(PartnerSecurityPolicyStatus_Disabled),
		string(PartnerSecurityPolicyStatus_Enabled),
	}
}

func (s *PartnerSecurityPolicyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecurityPolicyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecurityPolicyStatus(input string) (*PartnerSecurityPolicyStatus, error) {
	vals := map[string]PartnerSecurityPolicyStatus{
		"disabled": PartnerSecurityPolicyStatus_Disabled,
		"enabled":  PartnerSecurityPolicyStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecurityPolicyStatus(input)
	return &out, nil
}
