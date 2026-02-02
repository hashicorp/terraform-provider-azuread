package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowBlockListEntryType string

const (
	SecurityTenantAllowBlockListEntryType_FileHash  SecurityTenantAllowBlockListEntryType = "fileHash"
	SecurityTenantAllowBlockListEntryType_Recipient SecurityTenantAllowBlockListEntryType = "recipient"
	SecurityTenantAllowBlockListEntryType_Sender    SecurityTenantAllowBlockListEntryType = "sender"
	SecurityTenantAllowBlockListEntryType_Url       SecurityTenantAllowBlockListEntryType = "url"
)

func PossibleValuesForSecurityTenantAllowBlockListEntryType() []string {
	return []string{
		string(SecurityTenantAllowBlockListEntryType_FileHash),
		string(SecurityTenantAllowBlockListEntryType_Recipient),
		string(SecurityTenantAllowBlockListEntryType_Sender),
		string(SecurityTenantAllowBlockListEntryType_Url),
	}
}

func (s *SecurityTenantAllowBlockListEntryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTenantAllowBlockListEntryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTenantAllowBlockListEntryType(input string) (*SecurityTenantAllowBlockListEntryType, error) {
	vals := map[string]SecurityTenantAllowBlockListEntryType{
		"filehash":  SecurityTenantAllowBlockListEntryType_FileHash,
		"recipient": SecurityTenantAllowBlockListEntryType_Recipient,
		"sender":    SecurityTenantAllowBlockListEntryType_Sender,
		"url":       SecurityTenantAllowBlockListEntryType_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTenantAllowBlockListEntryType(input)
	return &out, nil
}
