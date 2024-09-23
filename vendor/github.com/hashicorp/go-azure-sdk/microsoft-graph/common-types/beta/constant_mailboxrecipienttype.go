package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxRecipientType string

const (
	MailboxRecipientType_Equipment MailboxRecipientType = "equipment"
	MailboxRecipientType_Linked    MailboxRecipientType = "linked"
	MailboxRecipientType_Others    MailboxRecipientType = "others"
	MailboxRecipientType_Room      MailboxRecipientType = "room"
	MailboxRecipientType_Shared    MailboxRecipientType = "shared"
	MailboxRecipientType_Unknown   MailboxRecipientType = "unknown"
	MailboxRecipientType_User      MailboxRecipientType = "user"
)

func PossibleValuesForMailboxRecipientType() []string {
	return []string{
		string(MailboxRecipientType_Equipment),
		string(MailboxRecipientType_Linked),
		string(MailboxRecipientType_Others),
		string(MailboxRecipientType_Room),
		string(MailboxRecipientType_Shared),
		string(MailboxRecipientType_Unknown),
		string(MailboxRecipientType_User),
	}
}

func (s *MailboxRecipientType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxRecipientType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxRecipientType(input string) (*MailboxRecipientType, error) {
	vals := map[string]MailboxRecipientType{
		"equipment": MailboxRecipientType_Equipment,
		"linked":    MailboxRecipientType_Linked,
		"others":    MailboxRecipientType_Others,
		"room":      MailboxRecipientType_Room,
		"shared":    MailboxRecipientType_Shared,
		"unknown":   MailboxRecipientType_Unknown,
		"user":      MailboxRecipientType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxRecipientType(input)
	return &out, nil
}
