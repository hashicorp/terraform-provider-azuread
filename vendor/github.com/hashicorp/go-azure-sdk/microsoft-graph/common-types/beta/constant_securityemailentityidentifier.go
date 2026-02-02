package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailEntityIdentifier string

const (
	SecurityEmailEntityIdentifier_NetworkMessageId      SecurityEmailEntityIdentifier = "networkMessageId"
	SecurityEmailEntityIdentifier_RecipientEmailAddress SecurityEmailEntityIdentifier = "recipientEmailAddress"
)

func PossibleValuesForSecurityEmailEntityIdentifier() []string {
	return []string{
		string(SecurityEmailEntityIdentifier_NetworkMessageId),
		string(SecurityEmailEntityIdentifier_RecipientEmailAddress),
	}
}

func (s *SecurityEmailEntityIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailEntityIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailEntityIdentifier(input string) (*SecurityEmailEntityIdentifier, error) {
	vals := map[string]SecurityEmailEntityIdentifier{
		"networkmessageid":      SecurityEmailEntityIdentifier_NetworkMessageId,
		"recipientemailaddress": SecurityEmailEntityIdentifier_RecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailEntityIdentifier(input)
	return &out, nil
}
