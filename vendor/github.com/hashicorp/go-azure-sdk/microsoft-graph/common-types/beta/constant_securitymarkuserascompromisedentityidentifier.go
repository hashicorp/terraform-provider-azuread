package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMarkUserAsCompromisedEntityIdentifier string

const (
	SecurityMarkUserAsCompromisedEntityIdentifier_AccountObjectId                  SecurityMarkUserAsCompromisedEntityIdentifier = "accountObjectId"
	SecurityMarkUserAsCompromisedEntityIdentifier_InitiatingProcessAccountObjectId SecurityMarkUserAsCompromisedEntityIdentifier = "initiatingProcessAccountObjectId"
	SecurityMarkUserAsCompromisedEntityIdentifier_RecipientObjectId                SecurityMarkUserAsCompromisedEntityIdentifier = "recipientObjectId"
	SecurityMarkUserAsCompromisedEntityIdentifier_ServicePrincipalId               SecurityMarkUserAsCompromisedEntityIdentifier = "servicePrincipalId"
)

func PossibleValuesForSecurityMarkUserAsCompromisedEntityIdentifier() []string {
	return []string{
		string(SecurityMarkUserAsCompromisedEntityIdentifier_AccountObjectId),
		string(SecurityMarkUserAsCompromisedEntityIdentifier_InitiatingProcessAccountObjectId),
		string(SecurityMarkUserAsCompromisedEntityIdentifier_RecipientObjectId),
		string(SecurityMarkUserAsCompromisedEntityIdentifier_ServicePrincipalId),
	}
}

func (s *SecurityMarkUserAsCompromisedEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMarkUserAsCompromisedEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMarkUserAsCompromisedEntityIdentifier(input string) (*SecurityMarkUserAsCompromisedEntityIdentifier, error) {
	vals := map[string]SecurityMarkUserAsCompromisedEntityIdentifier{
		"accountobjectid":                  SecurityMarkUserAsCompromisedEntityIdentifier_AccountObjectId,
		"initiatingprocessaccountobjectid": SecurityMarkUserAsCompromisedEntityIdentifier_InitiatingProcessAccountObjectId,
		"recipientobjectid":                SecurityMarkUserAsCompromisedEntityIdentifier_RecipientObjectId,
		"serviceprincipalid":               SecurityMarkUserAsCompromisedEntityIdentifier_ServicePrincipalId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMarkUserAsCompromisedEntityIdentifier(input)
	return &out, nil
}
