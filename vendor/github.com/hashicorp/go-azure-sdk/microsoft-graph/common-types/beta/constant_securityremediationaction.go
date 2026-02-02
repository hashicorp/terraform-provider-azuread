package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRemediationAction string

const (
	SecurityRemediationAction_HardDelete         SecurityRemediationAction = "hardDelete"
	SecurityRemediationAction_MoveToDeletedItems SecurityRemediationAction = "moveToDeletedItems"
	SecurityRemediationAction_MoveToInbox        SecurityRemediationAction = "moveToInbox"
	SecurityRemediationAction_MoveToJunk         SecurityRemediationAction = "moveToJunk"
	SecurityRemediationAction_SoftDelete         SecurityRemediationAction = "softDelete"
)

func PossibleValuesForSecurityRemediationAction() []string {
	return []string{
		string(SecurityRemediationAction_HardDelete),
		string(SecurityRemediationAction_MoveToDeletedItems),
		string(SecurityRemediationAction_MoveToInbox),
		string(SecurityRemediationAction_MoveToJunk),
		string(SecurityRemediationAction_SoftDelete),
	}
}

func (s *SecurityRemediationAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRemediationAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRemediationAction(input string) (*SecurityRemediationAction, error) {
	vals := map[string]SecurityRemediationAction{
		"harddelete":         SecurityRemediationAction_HardDelete,
		"movetodeleteditems": SecurityRemediationAction_MoveToDeletedItems,
		"movetoinbox":        SecurityRemediationAction_MoveToInbox,
		"movetojunk":         SecurityRemediationAction_MoveToJunk,
		"softdelete":         SecurityRemediationAction_SoftDelete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRemediationAction(input)
	return &out, nil
}
