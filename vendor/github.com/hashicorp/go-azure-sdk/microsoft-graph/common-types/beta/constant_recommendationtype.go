package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationType string

const (
	RecommendationType_AadConnectDeprecated                RecommendationType = "aadConnectDeprecated"
	RecommendationType_AadGraphDeprecationApplication      RecommendationType = "aadGraphDeprecationApplication"
	RecommendationType_AadGraphDeprecationServicePrincipal RecommendationType = "aadGraphDeprecationServicePrincipal"
	RecommendationType_AdalToMsalMigration                 RecommendationType = "adalToMsalMigration"
	RecommendationType_AdfsAppsMigration                   RecommendationType = "adfsAppsMigration"
	RecommendationType_AdminMFAV2                          RecommendationType = "adminMFAV2"
	RecommendationType_AppRoleAssignmentsGroups            RecommendationType = "appRoleAssignmentsGroups"
	RecommendationType_AppRoleAssignmentsUsers             RecommendationType = "appRoleAssignmentsUsers"
	RecommendationType_ApplicationCredentialExpiry         RecommendationType = "applicationCredentialExpiry"
	RecommendationType_BlockLegacyAuthentication           RecommendationType = "blockLegacyAuthentication"
	RecommendationType_EnableDesktopSSO                    RecommendationType = "enableDesktopSSO"
	RecommendationType_EnablePHS                           RecommendationType = "enablePHS"
	RecommendationType_EnableProvisioning                  RecommendationType = "enableProvisioning"
	RecommendationType_InactiveGuests                      RecommendationType = "inactiveGuests"
	RecommendationType_IntegratedApps                      RecommendationType = "integratedApps"
	RecommendationType_LongLivedCredentials                RecommendationType = "longLivedCredentials"
	RecommendationType_ManagedIdentity                     RecommendationType = "managedIdentity"
	RecommendationType_MfaRegistrationV2                   RecommendationType = "mfaRegistrationV2"
	RecommendationType_MfaServerDeprecation                RecommendationType = "mfaServerDeprecation"
	RecommendationType_OneAdmin                            RecommendationType = "oneAdmin"
	RecommendationType_OverprivilegedApps                  RecommendationType = "overprivilegedApps"
	RecommendationType_OwnerlessApps                       RecommendationType = "ownerlessApps"
	RecommendationType_PasswordHashSync                    RecommendationType = "passwordHashSync"
	RecommendationType_PrivateLinkForAAD                   RecommendationType = "privateLinkForAAD"
	RecommendationType_PwagePolicyNew                      RecommendationType = "pwagePolicyNew"
	RecommendationType_RoleOverlap                         RecommendationType = "roleOverlap"
	RecommendationType_SelfServicePasswordReset            RecommendationType = "selfServicePasswordReset"
	RecommendationType_ServicePrincipalKeyExpiry           RecommendationType = "servicePrincipalKeyExpiry"
	RecommendationType_SigninRiskPolicy                    RecommendationType = "signinRiskPolicy"
	RecommendationType_StaleAppCreds                       RecommendationType = "staleAppCreds"
	RecommendationType_StaleApps                           RecommendationType = "staleApps"
	RecommendationType_SwitchFromPerUserMFA                RecommendationType = "switchFromPerUserMFA"
	RecommendationType_TenantMFA                           RecommendationType = "tenantMFA"
	RecommendationType_ThirdPartyApps                      RecommendationType = "thirdPartyApps"
	RecommendationType_TurnOffPerUserMFA                   RecommendationType = "turnOffPerUserMFA"
	RecommendationType_UseAuthenticatorApp                 RecommendationType = "useAuthenticatorApp"
	RecommendationType_UseMyApps                           RecommendationType = "useMyApps"
	RecommendationType_UserRiskPolicy                      RecommendationType = "userRiskPolicy"
	RecommendationType_VerifyAppPublisher                  RecommendationType = "verifyAppPublisher"
)

func PossibleValuesForRecommendationType() []string {
	return []string{
		string(RecommendationType_AadConnectDeprecated),
		string(RecommendationType_AadGraphDeprecationApplication),
		string(RecommendationType_AadGraphDeprecationServicePrincipal),
		string(RecommendationType_AdalToMsalMigration),
		string(RecommendationType_AdfsAppsMigration),
		string(RecommendationType_AdminMFAV2),
		string(RecommendationType_AppRoleAssignmentsGroups),
		string(RecommendationType_AppRoleAssignmentsUsers),
		string(RecommendationType_ApplicationCredentialExpiry),
		string(RecommendationType_BlockLegacyAuthentication),
		string(RecommendationType_EnableDesktopSSO),
		string(RecommendationType_EnablePHS),
		string(RecommendationType_EnableProvisioning),
		string(RecommendationType_InactiveGuests),
		string(RecommendationType_IntegratedApps),
		string(RecommendationType_LongLivedCredentials),
		string(RecommendationType_ManagedIdentity),
		string(RecommendationType_MfaRegistrationV2),
		string(RecommendationType_MfaServerDeprecation),
		string(RecommendationType_OneAdmin),
		string(RecommendationType_OverprivilegedApps),
		string(RecommendationType_OwnerlessApps),
		string(RecommendationType_PasswordHashSync),
		string(RecommendationType_PrivateLinkForAAD),
		string(RecommendationType_PwagePolicyNew),
		string(RecommendationType_RoleOverlap),
		string(RecommendationType_SelfServicePasswordReset),
		string(RecommendationType_ServicePrincipalKeyExpiry),
		string(RecommendationType_SigninRiskPolicy),
		string(RecommendationType_StaleAppCreds),
		string(RecommendationType_StaleApps),
		string(RecommendationType_SwitchFromPerUserMFA),
		string(RecommendationType_TenantMFA),
		string(RecommendationType_ThirdPartyApps),
		string(RecommendationType_TurnOffPerUserMFA),
		string(RecommendationType_UseAuthenticatorApp),
		string(RecommendationType_UseMyApps),
		string(RecommendationType_UserRiskPolicy),
		string(RecommendationType_VerifyAppPublisher),
	}
}

func (s *RecommendationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationType(input string) (*RecommendationType, error) {
	vals := map[string]RecommendationType{
		"aadconnectdeprecated":                RecommendationType_AadConnectDeprecated,
		"aadgraphdeprecationapplication":      RecommendationType_AadGraphDeprecationApplication,
		"aadgraphdeprecationserviceprincipal": RecommendationType_AadGraphDeprecationServicePrincipal,
		"adaltomsalmigration":                 RecommendationType_AdalToMsalMigration,
		"adfsappsmigration":                   RecommendationType_AdfsAppsMigration,
		"adminmfav2":                          RecommendationType_AdminMFAV2,
		"approleassignmentsgroups":            RecommendationType_AppRoleAssignmentsGroups,
		"approleassignmentsusers":             RecommendationType_AppRoleAssignmentsUsers,
		"applicationcredentialexpiry":         RecommendationType_ApplicationCredentialExpiry,
		"blocklegacyauthentication":           RecommendationType_BlockLegacyAuthentication,
		"enabledesktopsso":                    RecommendationType_EnableDesktopSSO,
		"enablephs":                           RecommendationType_EnablePHS,
		"enableprovisioning":                  RecommendationType_EnableProvisioning,
		"inactiveguests":                      RecommendationType_InactiveGuests,
		"integratedapps":                      RecommendationType_IntegratedApps,
		"longlivedcredentials":                RecommendationType_LongLivedCredentials,
		"managedidentity":                     RecommendationType_ManagedIdentity,
		"mfaregistrationv2":                   RecommendationType_MfaRegistrationV2,
		"mfaserverdeprecation":                RecommendationType_MfaServerDeprecation,
		"oneadmin":                            RecommendationType_OneAdmin,
		"overprivilegedapps":                  RecommendationType_OverprivilegedApps,
		"ownerlessapps":                       RecommendationType_OwnerlessApps,
		"passwordhashsync":                    RecommendationType_PasswordHashSync,
		"privatelinkforaad":                   RecommendationType_PrivateLinkForAAD,
		"pwagepolicynew":                      RecommendationType_PwagePolicyNew,
		"roleoverlap":                         RecommendationType_RoleOverlap,
		"selfservicepasswordreset":            RecommendationType_SelfServicePasswordReset,
		"serviceprincipalkeyexpiry":           RecommendationType_ServicePrincipalKeyExpiry,
		"signinriskpolicy":                    RecommendationType_SigninRiskPolicy,
		"staleappcreds":                       RecommendationType_StaleAppCreds,
		"staleapps":                           RecommendationType_StaleApps,
		"switchfromperusermfa":                RecommendationType_SwitchFromPerUserMFA,
		"tenantmfa":                           RecommendationType_TenantMFA,
		"thirdpartyapps":                      RecommendationType_ThirdPartyApps,
		"turnoffperusermfa":                   RecommendationType_TurnOffPerUserMFA,
		"useauthenticatorapp":                 RecommendationType_UseAuthenticatorApp,
		"usemyapps":                           RecommendationType_UseMyApps,
		"userriskpolicy":                      RecommendationType_UserRiskPolicy,
		"verifyapppublisher":                  RecommendationType_VerifyAppPublisher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationType(input)
	return &out, nil
}
