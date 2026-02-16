package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountStatus string

const (
	AccountStatus_Active    AccountStatus = "active"
	AccountStatus_Deleted   AccountStatus = "deleted"
	AccountStatus_Staged    AccountStatus = "staged"
	AccountStatus_Suspended AccountStatus = "suspended"
	AccountStatus_Unknown   AccountStatus = "unknown"
)

func PossibleValuesForAccountStatus() []string {
	return []string{
		string(AccountStatus_Active),
		string(AccountStatus_Deleted),
		string(AccountStatus_Staged),
		string(AccountStatus_Suspended),
		string(AccountStatus_Unknown),
	}
}

func (s *AccountStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccountStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccountStatus(input string) (*AccountStatus, error) {
	vals := map[string]AccountStatus{
		"active":    AccountStatus_Active,
		"deleted":   AccountStatus_Deleted,
		"staged":    AccountStatus_Staged,
		"suspended": AccountStatus_Suspended,
		"unknown":   AccountStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountStatus(input)
	return &out, nil
}
