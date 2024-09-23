package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountTargetContentType string

const (
	AccountTargetContentType_AddressBook AccountTargetContentType = "addressBook"
	AccountTargetContentType_IncludeAll  AccountTargetContentType = "includeAll"
	AccountTargetContentType_Unknown     AccountTargetContentType = "unknown"
)

func PossibleValuesForAccountTargetContentType() []string {
	return []string{
		string(AccountTargetContentType_AddressBook),
		string(AccountTargetContentType_IncludeAll),
		string(AccountTargetContentType_Unknown),
	}
}

func (s *AccountTargetContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccountTargetContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccountTargetContentType(input string) (*AccountTargetContentType, error) {
	vals := map[string]AccountTargetContentType{
		"addressbook": AccountTargetContentType_AddressBook,
		"includeall":  AccountTargetContentType_IncludeAll,
		"unknown":     AccountTargetContentType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountTargetContentType(input)
	return &out, nil
}
