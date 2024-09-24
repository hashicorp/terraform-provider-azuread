package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailDestinationRoutingReason string

const (
	MailDestinationRoutingReason_AdvancedSpamFiltering MailDestinationRoutingReason = "advancedSpamFiltering"
	MailDestinationRoutingReason_AutoPurgeToDeleted    MailDestinationRoutingReason = "autoPurgeToDeleted"
	MailDestinationRoutingReason_AutoPurgeToInbox      MailDestinationRoutingReason = "autoPurgeToInbox"
	MailDestinationRoutingReason_AutoPurgeToJunk       MailDestinationRoutingReason = "autoPurgeToJunk"
	MailDestinationRoutingReason_BlockedSender         MailDestinationRoutingReason = "blockedSender"
	MailDestinationRoutingReason_DomainAllowList       MailDestinationRoutingReason = "domainAllowList"
	MailDestinationRoutingReason_DomainBlockList       MailDestinationRoutingReason = "domainBlockList"
	MailDestinationRoutingReason_FirstTimeSender       MailDestinationRoutingReason = "firstTimeSender"
	MailDestinationRoutingReason_Junk                  MailDestinationRoutingReason = "junk"
	MailDestinationRoutingReason_MailFlowRule          MailDestinationRoutingReason = "mailFlowRule"
	MailDestinationRoutingReason_None                  MailDestinationRoutingReason = "none"
	MailDestinationRoutingReason_NotInAddressBook      MailDestinationRoutingReason = "notInAddressBook"
	MailDestinationRoutingReason_NotJunk               MailDestinationRoutingReason = "notJunk"
	MailDestinationRoutingReason_Outbound              MailDestinationRoutingReason = "outbound"
	MailDestinationRoutingReason_SafeSender            MailDestinationRoutingReason = "safeSender"
)

func PossibleValuesForMailDestinationRoutingReason() []string {
	return []string{
		string(MailDestinationRoutingReason_AdvancedSpamFiltering),
		string(MailDestinationRoutingReason_AutoPurgeToDeleted),
		string(MailDestinationRoutingReason_AutoPurgeToInbox),
		string(MailDestinationRoutingReason_AutoPurgeToJunk),
		string(MailDestinationRoutingReason_BlockedSender),
		string(MailDestinationRoutingReason_DomainAllowList),
		string(MailDestinationRoutingReason_DomainBlockList),
		string(MailDestinationRoutingReason_FirstTimeSender),
		string(MailDestinationRoutingReason_Junk),
		string(MailDestinationRoutingReason_MailFlowRule),
		string(MailDestinationRoutingReason_None),
		string(MailDestinationRoutingReason_NotInAddressBook),
		string(MailDestinationRoutingReason_NotJunk),
		string(MailDestinationRoutingReason_Outbound),
		string(MailDestinationRoutingReason_SafeSender),
	}
}

func (s *MailDestinationRoutingReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailDestinationRoutingReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailDestinationRoutingReason(input string) (*MailDestinationRoutingReason, error) {
	vals := map[string]MailDestinationRoutingReason{
		"advancedspamfiltering": MailDestinationRoutingReason_AdvancedSpamFiltering,
		"autopurgetodeleted":    MailDestinationRoutingReason_AutoPurgeToDeleted,
		"autopurgetoinbox":      MailDestinationRoutingReason_AutoPurgeToInbox,
		"autopurgetojunk":       MailDestinationRoutingReason_AutoPurgeToJunk,
		"blockedsender":         MailDestinationRoutingReason_BlockedSender,
		"domainallowlist":       MailDestinationRoutingReason_DomainAllowList,
		"domainblocklist":       MailDestinationRoutingReason_DomainBlockList,
		"firsttimesender":       MailDestinationRoutingReason_FirstTimeSender,
		"junk":                  MailDestinationRoutingReason_Junk,
		"mailflowrule":          MailDestinationRoutingReason_MailFlowRule,
		"none":                  MailDestinationRoutingReason_None,
		"notinaddressbook":      MailDestinationRoutingReason_NotInAddressBook,
		"notjunk":               MailDestinationRoutingReason_NotJunk,
		"outbound":              MailDestinationRoutingReason_Outbound,
		"safesender":            MailDestinationRoutingReason_SafeSender,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailDestinationRoutingReason(input)
	return &out, nil
}
