package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityForceUserPasswordResetEntityIdentifier string

const (
	SecurityForceUserPasswordResetEntityIdentifier_AccountSid                  SecurityForceUserPasswordResetEntityIdentifier = "accountSid"
	SecurityForceUserPasswordResetEntityIdentifier_InitiatingProcessAccountSid SecurityForceUserPasswordResetEntityIdentifier = "initiatingProcessAccountSid"
	SecurityForceUserPasswordResetEntityIdentifier_OnPremSid                   SecurityForceUserPasswordResetEntityIdentifier = "onPremSid"
	SecurityForceUserPasswordResetEntityIdentifier_RequestAccountSid           SecurityForceUserPasswordResetEntityIdentifier = "requestAccountSid"
)

func PossibleValuesForSecurityForceUserPasswordResetEntityIdentifier() []string {
	return []string{
		string(SecurityForceUserPasswordResetEntityIdentifier_AccountSid),
		string(SecurityForceUserPasswordResetEntityIdentifier_InitiatingProcessAccountSid),
		string(SecurityForceUserPasswordResetEntityIdentifier_OnPremSid),
		string(SecurityForceUserPasswordResetEntityIdentifier_RequestAccountSid),
	}
}

func (s *SecurityForceUserPasswordResetEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityForceUserPasswordResetEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityForceUserPasswordResetEntityIdentifier(input string) (*SecurityForceUserPasswordResetEntityIdentifier, error) {
	vals := map[string]SecurityForceUserPasswordResetEntityIdentifier{
		"accountsid":                  SecurityForceUserPasswordResetEntityIdentifier_AccountSid,
		"initiatingprocessaccountsid": SecurityForceUserPasswordResetEntityIdentifier_InitiatingProcessAccountSid,
		"onpremsid":                   SecurityForceUserPasswordResetEntityIdentifier_OnPremSid,
		"requestaccountsid":           SecurityForceUserPasswordResetEntityIdentifier_RequestAccountSid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityForceUserPasswordResetEntityIdentifier(input)
	return &out, nil
}
