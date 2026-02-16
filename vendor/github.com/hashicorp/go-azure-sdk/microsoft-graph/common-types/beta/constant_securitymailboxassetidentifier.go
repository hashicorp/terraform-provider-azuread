package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxAssetIdentifier string

const (
	SecurityMailboxAssetIdentifier_AccountUpn                  SecurityMailboxAssetIdentifier = "accountUpn"
	SecurityMailboxAssetIdentifier_FileOwnerUpn                SecurityMailboxAssetIdentifier = "fileOwnerUpn"
	SecurityMailboxAssetIdentifier_InitiatingProcessAccountUpn SecurityMailboxAssetIdentifier = "initiatingProcessAccountUpn"
	SecurityMailboxAssetIdentifier_LastModifyingAccountUpn     SecurityMailboxAssetIdentifier = "lastModifyingAccountUpn"
	SecurityMailboxAssetIdentifier_RecipientEmailAddress       SecurityMailboxAssetIdentifier = "recipientEmailAddress"
	SecurityMailboxAssetIdentifier_SenderDisplayName           SecurityMailboxAssetIdentifier = "senderDisplayName"
	SecurityMailboxAssetIdentifier_SenderFromAddress           SecurityMailboxAssetIdentifier = "senderFromAddress"
	SecurityMailboxAssetIdentifier_SenderMailFromAddress       SecurityMailboxAssetIdentifier = "senderMailFromAddress"
	SecurityMailboxAssetIdentifier_TargetAccountUpn            SecurityMailboxAssetIdentifier = "targetAccountUpn"
)

func PossibleValuesForSecurityMailboxAssetIdentifier() []string {
	return []string{
		string(SecurityMailboxAssetIdentifier_AccountUpn),
		string(SecurityMailboxAssetIdentifier_FileOwnerUpn),
		string(SecurityMailboxAssetIdentifier_InitiatingProcessAccountUpn),
		string(SecurityMailboxAssetIdentifier_LastModifyingAccountUpn),
		string(SecurityMailboxAssetIdentifier_RecipientEmailAddress),
		string(SecurityMailboxAssetIdentifier_SenderDisplayName),
		string(SecurityMailboxAssetIdentifier_SenderFromAddress),
		string(SecurityMailboxAssetIdentifier_SenderMailFromAddress),
		string(SecurityMailboxAssetIdentifier_TargetAccountUpn),
	}
}

func (s *SecurityMailboxAssetIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailboxAssetIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailboxAssetIdentifier(input string) (*SecurityMailboxAssetIdentifier, error) {
	vals := map[string]SecurityMailboxAssetIdentifier{
		"accountupn":                  SecurityMailboxAssetIdentifier_AccountUpn,
		"fileownerupn":                SecurityMailboxAssetIdentifier_FileOwnerUpn,
		"initiatingprocessaccountupn": SecurityMailboxAssetIdentifier_InitiatingProcessAccountUpn,
		"lastmodifyingaccountupn":     SecurityMailboxAssetIdentifier_LastModifyingAccountUpn,
		"recipientemailaddress":       SecurityMailboxAssetIdentifier_RecipientEmailAddress,
		"senderdisplayname":           SecurityMailboxAssetIdentifier_SenderDisplayName,
		"senderfromaddress":           SecurityMailboxAssetIdentifier_SenderFromAddress,
		"sendermailfromaddress":       SecurityMailboxAssetIdentifier_SenderMailFromAddress,
		"targetaccountupn":            SecurityMailboxAssetIdentifier_TargetAccountUpn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxAssetIdentifier(input)
	return &out, nil
}
