package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxType string

const (
	MailboxType_Shared  MailboxType = "shared"
	MailboxType_Unknown MailboxType = "unknown"
	MailboxType_User    MailboxType = "user"
)

func PossibleValuesForMailboxType() []string {
	return []string{
		string(MailboxType_Shared),
		string(MailboxType_Unknown),
		string(MailboxType_User),
	}
}

func (s *MailboxType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxType(input string) (*MailboxType, error) {
	vals := map[string]MailboxType{
		"shared":  MailboxType_Shared,
		"unknown": MailboxType_Unknown,
		"user":    MailboxType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxType(input)
	return &out, nil
}
