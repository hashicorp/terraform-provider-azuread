package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetail string

const (
	RiskDetail_AdminConfirmedAccountSafe                 RiskDetail = "adminConfirmedAccountSafe"
	RiskDetail_AdminConfirmedServicePrincipalCompromised RiskDetail = "adminConfirmedServicePrincipalCompromised"
	RiskDetail_AdminConfirmedSigninCompromised           RiskDetail = "adminConfirmedSigninCompromised"
	RiskDetail_AdminConfirmedSigninSafe                  RiskDetail = "adminConfirmedSigninSafe"
	RiskDetail_AdminConfirmedUserCompromised             RiskDetail = "adminConfirmedUserCompromised"
	RiskDetail_AdminDismissedAllRiskForServicePrincipal  RiskDetail = "adminDismissedAllRiskForServicePrincipal"
	RiskDetail_AdminDismissedAllRiskForUser              RiskDetail = "adminDismissedAllRiskForUser"
	RiskDetail_AdminDismissedRiskForSignIn               RiskDetail = "adminDismissedRiskForSignIn"
	RiskDetail_AdminGeneratedTemporaryPassword           RiskDetail = "adminGeneratedTemporaryPassword"
	RiskDetail_AiConfirmedSigninSafe                     RiskDetail = "aiConfirmedSigninSafe"
	RiskDetail_Hidden                                    RiskDetail = "hidden"
	RiskDetail_M365DAdminDismissedDetection              RiskDetail = "m365DAdminDismissedDetection"
	RiskDetail_None                                      RiskDetail = "none"
	RiskDetail_UserChangedPasswordOnPremises             RiskDetail = "userChangedPasswordOnPremises"
	RiskDetail_UserPassedMFADrivenByRiskBasedPolicy      RiskDetail = "userPassedMFADrivenByRiskBasedPolicy"
	RiskDetail_UserPerformedSecuredPasswordChange        RiskDetail = "userPerformedSecuredPasswordChange"
	RiskDetail_UserPerformedSecuredPasswordReset         RiskDetail = "userPerformedSecuredPasswordReset"
)

func PossibleValuesForRiskDetail() []string {
	return []string{
		string(RiskDetail_AdminConfirmedAccountSafe),
		string(RiskDetail_AdminConfirmedServicePrincipalCompromised),
		string(RiskDetail_AdminConfirmedSigninCompromised),
		string(RiskDetail_AdminConfirmedSigninSafe),
		string(RiskDetail_AdminConfirmedUserCompromised),
		string(RiskDetail_AdminDismissedAllRiskForServicePrincipal),
		string(RiskDetail_AdminDismissedAllRiskForUser),
		string(RiskDetail_AdminDismissedRiskForSignIn),
		string(RiskDetail_AdminGeneratedTemporaryPassword),
		string(RiskDetail_AiConfirmedSigninSafe),
		string(RiskDetail_Hidden),
		string(RiskDetail_M365DAdminDismissedDetection),
		string(RiskDetail_None),
		string(RiskDetail_UserChangedPasswordOnPremises),
		string(RiskDetail_UserPassedMFADrivenByRiskBasedPolicy),
		string(RiskDetail_UserPerformedSecuredPasswordChange),
		string(RiskDetail_UserPerformedSecuredPasswordReset),
	}
}

func (s *RiskDetail) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetail(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetail(input string) (*RiskDetail, error) {
	vals := map[string]RiskDetail{
		"adminconfirmedaccountsafe":                 RiskDetail_AdminConfirmedAccountSafe,
		"adminconfirmedserviceprincipalcompromised": RiskDetail_AdminConfirmedServicePrincipalCompromised,
		"adminconfirmedsignincompromised":           RiskDetail_AdminConfirmedSigninCompromised,
		"adminconfirmedsigninsafe":                  RiskDetail_AdminConfirmedSigninSafe,
		"adminconfirmedusercompromised":             RiskDetail_AdminConfirmedUserCompromised,
		"admindismissedallriskforserviceprincipal":  RiskDetail_AdminDismissedAllRiskForServicePrincipal,
		"admindismissedallriskforuser":              RiskDetail_AdminDismissedAllRiskForUser,
		"admindismissedriskforsignin":               RiskDetail_AdminDismissedRiskForSignIn,
		"admingeneratedtemporarypassword":           RiskDetail_AdminGeneratedTemporaryPassword,
		"aiconfirmedsigninsafe":                     RiskDetail_AiConfirmedSigninSafe,
		"hidden":                                    RiskDetail_Hidden,
		"m365dadmindismisseddetection":              RiskDetail_M365DAdminDismissedDetection,
		"none":                                      RiskDetail_None,
		"userchangedpasswordonpremises":             RiskDetail_UserChangedPasswordOnPremises,
		"userpassedmfadrivenbyriskbasedpolicy":      RiskDetail_UserPassedMFADrivenByRiskBasedPolicy,
		"userperformedsecuredpasswordchange":        RiskDetail_UserPerformedSecuredPasswordChange,
		"userperformedsecuredpasswordreset":         RiskDetail_UserPerformedSecuredPasswordReset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetail(input)
	return &out, nil
}
