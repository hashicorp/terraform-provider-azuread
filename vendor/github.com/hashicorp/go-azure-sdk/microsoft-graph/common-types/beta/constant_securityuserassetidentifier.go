package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserAssetIdentifier string

const (
	SecurityUserAssetIdentifier_AccountDomain               SecurityUserAssetIdentifier = "accountDomain"
	SecurityUserAssetIdentifier_AccountId                   SecurityUserAssetIdentifier = "accountId"
	SecurityUserAssetIdentifier_AccountName                 SecurityUserAssetIdentifier = "accountName"
	SecurityUserAssetIdentifier_AccountObjectId             SecurityUserAssetIdentifier = "accountObjectId"
	SecurityUserAssetIdentifier_AccountSid                  SecurityUserAssetIdentifier = "accountSid"
	SecurityUserAssetIdentifier_AccountUpn                  SecurityUserAssetIdentifier = "accountUpn"
	SecurityUserAssetIdentifier_InitiatingAccountDomain     SecurityUserAssetIdentifier = "initiatingAccountDomain"
	SecurityUserAssetIdentifier_InitiatingAccountName       SecurityUserAssetIdentifier = "initiatingAccountName"
	SecurityUserAssetIdentifier_InitiatingAccountSid        SecurityUserAssetIdentifier = "initiatingAccountSid"
	SecurityUserAssetIdentifier_InitiatingProcessAccountUpn SecurityUserAssetIdentifier = "initiatingProcessAccountUpn"
	SecurityUserAssetIdentifier_ProcessAccountObjectId      SecurityUserAssetIdentifier = "processAccountObjectId"
	SecurityUserAssetIdentifier_RecipientObjectId           SecurityUserAssetIdentifier = "recipientObjectId"
	SecurityUserAssetIdentifier_RequestAccountDomain        SecurityUserAssetIdentifier = "requestAccountDomain"
	SecurityUserAssetIdentifier_RequestAccountName          SecurityUserAssetIdentifier = "requestAccountName"
	SecurityUserAssetIdentifier_RequestAccountSid           SecurityUserAssetIdentifier = "requestAccountSid"
	SecurityUserAssetIdentifier_ServicePrincipalId          SecurityUserAssetIdentifier = "servicePrincipalId"
	SecurityUserAssetIdentifier_ServicePrincipalName        SecurityUserAssetIdentifier = "servicePrincipalName"
	SecurityUserAssetIdentifier_TargetAccountUpn            SecurityUserAssetIdentifier = "targetAccountUpn"
)

func PossibleValuesForSecurityUserAssetIdentifier() []string {
	return []string{
		string(SecurityUserAssetIdentifier_AccountDomain),
		string(SecurityUserAssetIdentifier_AccountId),
		string(SecurityUserAssetIdentifier_AccountName),
		string(SecurityUserAssetIdentifier_AccountObjectId),
		string(SecurityUserAssetIdentifier_AccountSid),
		string(SecurityUserAssetIdentifier_AccountUpn),
		string(SecurityUserAssetIdentifier_InitiatingAccountDomain),
		string(SecurityUserAssetIdentifier_InitiatingAccountName),
		string(SecurityUserAssetIdentifier_InitiatingAccountSid),
		string(SecurityUserAssetIdentifier_InitiatingProcessAccountUpn),
		string(SecurityUserAssetIdentifier_ProcessAccountObjectId),
		string(SecurityUserAssetIdentifier_RecipientObjectId),
		string(SecurityUserAssetIdentifier_RequestAccountDomain),
		string(SecurityUserAssetIdentifier_RequestAccountName),
		string(SecurityUserAssetIdentifier_RequestAccountSid),
		string(SecurityUserAssetIdentifier_ServicePrincipalId),
		string(SecurityUserAssetIdentifier_ServicePrincipalName),
		string(SecurityUserAssetIdentifier_TargetAccountUpn),
	}
}

func (s *SecurityUserAssetIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserAssetIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserAssetIdentifier(input string) (*SecurityUserAssetIdentifier, error) {
	vals := map[string]SecurityUserAssetIdentifier{
		"accountdomain":               SecurityUserAssetIdentifier_AccountDomain,
		"accountid":                   SecurityUserAssetIdentifier_AccountId,
		"accountname":                 SecurityUserAssetIdentifier_AccountName,
		"accountobjectid":             SecurityUserAssetIdentifier_AccountObjectId,
		"accountsid":                  SecurityUserAssetIdentifier_AccountSid,
		"accountupn":                  SecurityUserAssetIdentifier_AccountUpn,
		"initiatingaccountdomain":     SecurityUserAssetIdentifier_InitiatingAccountDomain,
		"initiatingaccountname":       SecurityUserAssetIdentifier_InitiatingAccountName,
		"initiatingaccountsid":        SecurityUserAssetIdentifier_InitiatingAccountSid,
		"initiatingprocessaccountupn": SecurityUserAssetIdentifier_InitiatingProcessAccountUpn,
		"processaccountobjectid":      SecurityUserAssetIdentifier_ProcessAccountObjectId,
		"recipientobjectid":           SecurityUserAssetIdentifier_RecipientObjectId,
		"requestaccountdomain":        SecurityUserAssetIdentifier_RequestAccountDomain,
		"requestaccountname":          SecurityUserAssetIdentifier_RequestAccountName,
		"requestaccountsid":           SecurityUserAssetIdentifier_RequestAccountSid,
		"serviceprincipalid":          SecurityUserAssetIdentifier_ServicePrincipalId,
		"serviceprincipalname":        SecurityUserAssetIdentifier_ServicePrincipalName,
		"targetaccountupn":            SecurityUserAssetIdentifier_TargetAccountUpn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserAssetIdentifier(input)
	return &out, nil
}
