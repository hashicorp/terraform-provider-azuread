package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDisableUserEntityIdentifier string

const (
	SecurityDisableUserEntityIdentifier_AccountSid                  SecurityDisableUserEntityIdentifier = "accountSid"
	SecurityDisableUserEntityIdentifier_InitiatingProcessAccountSid SecurityDisableUserEntityIdentifier = "initiatingProcessAccountSid"
	SecurityDisableUserEntityIdentifier_OnPremSid                   SecurityDisableUserEntityIdentifier = "onPremSid"
	SecurityDisableUserEntityIdentifier_RequestAccountSid           SecurityDisableUserEntityIdentifier = "requestAccountSid"
)

func PossibleValuesForSecurityDisableUserEntityIdentifier() []string {
	return []string{
		string(SecurityDisableUserEntityIdentifier_AccountSid),
		string(SecurityDisableUserEntityIdentifier_InitiatingProcessAccountSid),
		string(SecurityDisableUserEntityIdentifier_OnPremSid),
		string(SecurityDisableUserEntityIdentifier_RequestAccountSid),
	}
}

func (s *SecurityDisableUserEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDisableUserEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDisableUserEntityIdentifier(input string) (*SecurityDisableUserEntityIdentifier, error) {
	vals := map[string]SecurityDisableUserEntityIdentifier{
		"accountsid":                  SecurityDisableUserEntityIdentifier_AccountSid,
		"initiatingprocessaccountsid": SecurityDisableUserEntityIdentifier_InitiatingProcessAccountSid,
		"onpremsid":                   SecurityDisableUserEntityIdentifier_OnPremSid,
		"requestaccountsid":           SecurityDisableUserEntityIdentifier_RequestAccountSid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDisableUserEntityIdentifier(input)
	return &out, nil
}
