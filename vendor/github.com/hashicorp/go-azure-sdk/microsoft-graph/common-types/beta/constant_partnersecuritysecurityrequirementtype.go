package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerSecuritySecurityRequirementType string

const (
	PartnerSecuritySecurityRequirementType_MfaEnforcedForAdmins                           PartnerSecuritySecurityRequirementType = "mfaEnforcedForAdmins"
	PartnerSecuritySecurityRequirementType_MfaEnforcedForAdminsOfCustomers                PartnerSecuritySecurityRequirementType = "mfaEnforcedForAdminsOfCustomers"
	PartnerSecuritySecurityRequirementType_SecurityAlertsPromptlyResolved                 PartnerSecuritySecurityRequirementType = "securityAlertsPromptlyResolved"
	PartnerSecuritySecurityRequirementType_SecurityContactProvided                        PartnerSecuritySecurityRequirementType = "securityContactProvided"
	PartnerSecuritySecurityRequirementType_SpendingBudgetSetForCustomerAzureSubscriptions PartnerSecuritySecurityRequirementType = "spendingBudgetSetForCustomerAzureSubscriptions"
)

func PossibleValuesForPartnerSecuritySecurityRequirementType() []string {
	return []string{
		string(PartnerSecuritySecurityRequirementType_MfaEnforcedForAdmins),
		string(PartnerSecuritySecurityRequirementType_MfaEnforcedForAdminsOfCustomers),
		string(PartnerSecuritySecurityRequirementType_SecurityAlertsPromptlyResolved),
		string(PartnerSecuritySecurityRequirementType_SecurityContactProvided),
		string(PartnerSecuritySecurityRequirementType_SpendingBudgetSetForCustomerAzureSubscriptions),
	}
}

func (s *PartnerSecuritySecurityRequirementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerSecuritySecurityRequirementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerSecuritySecurityRequirementType(input string) (*PartnerSecuritySecurityRequirementType, error) {
	vals := map[string]PartnerSecuritySecurityRequirementType{
		"mfaenforcedforadmins":                           PartnerSecuritySecurityRequirementType_MfaEnforcedForAdmins,
		"mfaenforcedforadminsofcustomers":                PartnerSecuritySecurityRequirementType_MfaEnforcedForAdminsOfCustomers,
		"securityalertspromptlyresolved":                 PartnerSecuritySecurityRequirementType_SecurityAlertsPromptlyResolved,
		"securitycontactprovided":                        PartnerSecuritySecurityRequirementType_SecurityContactProvided,
		"spendingbudgetsetforcustomerazuresubscriptions": PartnerSecuritySecurityRequirementType_SpendingBudgetSetForCustomerAzureSubscriptions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerSecuritySecurityRequirementType(input)
	return &out, nil
}
