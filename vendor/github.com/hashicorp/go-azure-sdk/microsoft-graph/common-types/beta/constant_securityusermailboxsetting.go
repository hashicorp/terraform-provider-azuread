package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserMailboxSetting string

const (
	SecurityUserMailboxSetting_CustomRule                             SecurityUserMailboxSetting = "customRule"
	SecurityUserMailboxSetting_Exclusive                              SecurityUserMailboxSetting = "exclusive"
	SecurityUserMailboxSetting_FromFirstTimeSender                    SecurityUserMailboxSetting = "fromFirstTimeSender"
	SecurityUserMailboxSetting_IsFromAddressInAddressBlockList        SecurityUserMailboxSetting = "isFromAddressInAddressBlockList"
	SecurityUserMailboxSetting_IsFromAddressInAddressBook             SecurityUserMailboxSetting = "isFromAddressInAddressBook"
	SecurityUserMailboxSetting_IsFromAddressInAddressImplicitJunkList SecurityUserMailboxSetting = "isFromAddressInAddressImplicitJunkList"
	SecurityUserMailboxSetting_IsFromAddressInAddressImplicitSafeList SecurityUserMailboxSetting = "isFromAddressInAddressImplicitSafeList"
	SecurityUserMailboxSetting_IsFromAddressInAddressSafeList         SecurityUserMailboxSetting = "isFromAddressInAddressSafeList"
	SecurityUserMailboxSetting_IsFromDomainInDomainBlockList          SecurityUserMailboxSetting = "isFromDomainInDomainBlockList"
	SecurityUserMailboxSetting_IsFromDomainInDomainSafeList           SecurityUserMailboxSetting = "isFromDomainInDomainSafeList"
	SecurityUserMailboxSetting_IsJunkMailRuleEnabled                  SecurityUserMailboxSetting = "isJunkMailRuleEnabled"
	SecurityUserMailboxSetting_IsRecipientInRecipientSafeList         SecurityUserMailboxSetting = "isRecipientInRecipientSafeList"
	SecurityUserMailboxSetting_JunkMailDeletion                       SecurityUserMailboxSetting = "junkMailDeletion"
	SecurityUserMailboxSetting_JunkMailRule                           SecurityUserMailboxSetting = "junkMailRule"
	SecurityUserMailboxSetting_None                                   SecurityUserMailboxSetting = "none"
	SecurityUserMailboxSetting_PriorSeenPass                          SecurityUserMailboxSetting = "priorSeenPass"
	SecurityUserMailboxSetting_SenderAuthenticationSucceeded          SecurityUserMailboxSetting = "senderAuthenticationSucceeded"
	SecurityUserMailboxSetting_SenderPraPresent                       SecurityUserMailboxSetting = "senderPraPresent"
)

func PossibleValuesForSecurityUserMailboxSetting() []string {
	return []string{
		string(SecurityUserMailboxSetting_CustomRule),
		string(SecurityUserMailboxSetting_Exclusive),
		string(SecurityUserMailboxSetting_FromFirstTimeSender),
		string(SecurityUserMailboxSetting_IsFromAddressInAddressBlockList),
		string(SecurityUserMailboxSetting_IsFromAddressInAddressBook),
		string(SecurityUserMailboxSetting_IsFromAddressInAddressImplicitJunkList),
		string(SecurityUserMailboxSetting_IsFromAddressInAddressImplicitSafeList),
		string(SecurityUserMailboxSetting_IsFromAddressInAddressSafeList),
		string(SecurityUserMailboxSetting_IsFromDomainInDomainBlockList),
		string(SecurityUserMailboxSetting_IsFromDomainInDomainSafeList),
		string(SecurityUserMailboxSetting_IsJunkMailRuleEnabled),
		string(SecurityUserMailboxSetting_IsRecipientInRecipientSafeList),
		string(SecurityUserMailboxSetting_JunkMailDeletion),
		string(SecurityUserMailboxSetting_JunkMailRule),
		string(SecurityUserMailboxSetting_None),
		string(SecurityUserMailboxSetting_PriorSeenPass),
		string(SecurityUserMailboxSetting_SenderAuthenticationSucceeded),
		string(SecurityUserMailboxSetting_SenderPraPresent),
	}
}

func (s *SecurityUserMailboxSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserMailboxSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserMailboxSetting(input string) (*SecurityUserMailboxSetting, error) {
	vals := map[string]SecurityUserMailboxSetting{
		"customrule":                             SecurityUserMailboxSetting_CustomRule,
		"exclusive":                              SecurityUserMailboxSetting_Exclusive,
		"fromfirsttimesender":                    SecurityUserMailboxSetting_FromFirstTimeSender,
		"isfromaddressinaddressblocklist":        SecurityUserMailboxSetting_IsFromAddressInAddressBlockList,
		"isfromaddressinaddressbook":             SecurityUserMailboxSetting_IsFromAddressInAddressBook,
		"isfromaddressinaddressimplicitjunklist": SecurityUserMailboxSetting_IsFromAddressInAddressImplicitJunkList,
		"isfromaddressinaddressimplicitsafelist": SecurityUserMailboxSetting_IsFromAddressInAddressImplicitSafeList,
		"isfromaddressinaddresssafelist":         SecurityUserMailboxSetting_IsFromAddressInAddressSafeList,
		"isfromdomainindomainblocklist":          SecurityUserMailboxSetting_IsFromDomainInDomainBlockList,
		"isfromdomainindomainsafelist":           SecurityUserMailboxSetting_IsFromDomainInDomainSafeList,
		"isjunkmailruleenabled":                  SecurityUserMailboxSetting_IsJunkMailRuleEnabled,
		"isrecipientinrecipientsafelist":         SecurityUserMailboxSetting_IsRecipientInRecipientSafeList,
		"junkmaildeletion":                       SecurityUserMailboxSetting_JunkMailDeletion,
		"junkmailrule":                           SecurityUserMailboxSetting_JunkMailRule,
		"none":                                   SecurityUserMailboxSetting_None,
		"priorseenpass":                          SecurityUserMailboxSetting_PriorSeenPass,
		"senderauthenticationsucceeded":          SecurityUserMailboxSetting_SenderAuthenticationSucceeded,
		"senderprapresent":                       SecurityUserMailboxSetting_SenderPraPresent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserMailboxSetting(input)
	return &out, nil
}
