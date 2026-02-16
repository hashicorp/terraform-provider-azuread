package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationSecret string

const (
	SynchronizationSecret_AppKey                          SynchronizationSecret = "AppKey"
	SynchronizationSecret_ApplicationTemplateIdentifier   SynchronizationSecret = "ApplicationTemplateIdentifier"
	SynchronizationSecret_AuthenticationType              SynchronizationSecret = "AuthenticationType"
	SynchronizationSecret_BaseAddress                     SynchronizationSecret = "BaseAddress"
	SynchronizationSecret_ClientIdentifier                SynchronizationSecret = "ClientIdentifier"
	SynchronizationSecret_ClientSecret                    SynchronizationSecret = "ClientSecret"
	SynchronizationSecret_CompanyId                       SynchronizationSecret = "CompanyId"
	SynchronizationSecret_ConnectionString                SynchronizationSecret = "ConnectionString"
	SynchronizationSecret_ConsumerKey                     SynchronizationSecret = "ConsumerKey"
	SynchronizationSecret_ConsumerSecret                  SynchronizationSecret = "ConsumerSecret"
	SynchronizationSecret_Domain                          SynchronizationSecret = "Domain"
	SynchronizationSecret_EnforceDomain                   SynchronizationSecret = "EnforceDomain"
	SynchronizationSecret_HardDeletesEnabled              SynchronizationSecret = "HardDeletesEnabled"
	SynchronizationSecret_InstanceName                    SynchronizationSecret = "InstanceName"
	SynchronizationSecret_None                            SynchronizationSecret = "None"
	SynchronizationSecret_OAuth2AccessToken               SynchronizationSecret = "Oauth2AccessToken"
	SynchronizationSecret_OAuth2AccessTokenCreationTime   SynchronizationSecret = "Oauth2AccessTokenCreationTime"
	SynchronizationSecret_OAuth2AuthorizationCode         SynchronizationSecret = "Oauth2AuthorizationCode"
	SynchronizationSecret_OAuth2AuthorizationUri          SynchronizationSecret = "Oauth2AuthorizationUri"
	SynchronizationSecret_OAuth2ClientId                  SynchronizationSecret = "Oauth2ClientId"
	SynchronizationSecret_OAuth2ClientSecret              SynchronizationSecret = "Oauth2ClientSecret"
	SynchronizationSecret_OAuth2RedirectUri               SynchronizationSecret = "Oauth2RedirectUri"
	SynchronizationSecret_OAuth2RefreshToken              SynchronizationSecret = "Oauth2RefreshToken"
	SynchronizationSecret_OAuth2TokenExchangeUri          SynchronizationSecret = "Oauth2TokenExchangeUri"
	SynchronizationSecret_Password                        SynchronizationSecret = "Password"
	SynchronizationSecret_PerformInboundEntitlementGrants SynchronizationSecret = "PerformInboundEntitlementGrants"
	SynchronizationSecret_Sandbox                         SynchronizationSecret = "Sandbox"
	SynchronizationSecret_SandboxName                     SynchronizationSecret = "SandboxName"
	SynchronizationSecret_SecretToken                     SynchronizationSecret = "SecretToken"
	SynchronizationSecret_Server                          SynchronizationSecret = "Server"
	SynchronizationSecret_SingleSignOnType                SynchronizationSecret = "SingleSignOnType"
	SynchronizationSecret_SkipOutOfScopeDeletions         SynchronizationSecret = "SkipOutOfScopeDeletions"
	SynchronizationSecret_SyncAgentADContainer            SynchronizationSecret = "SyncAgentADContainer"
	SynchronizationSecret_SyncAgentCompatibilityKey       SynchronizationSecret = "SyncAgentCompatibilityKey"
	SynchronizationSecret_SyncAll                         SynchronizationSecret = "SyncAll"
	SynchronizationSecret_SyncNotificationSettings        SynchronizationSecret = "SyncNotificationSettings"
	SynchronizationSecret_SynchronizationSchedule         SynchronizationSecret = "SynchronizationSchedule"
	SynchronizationSecret_SystemOfRecord                  SynchronizationSecret = "SystemOfRecord"
	SynchronizationSecret_TestReferences                  SynchronizationSecret = "TestReferences"
	SynchronizationSecret_TokenExpiration                 SynchronizationSecret = "TokenExpiration"
	SynchronizationSecret_TokenKey                        SynchronizationSecret = "TokenKey"
	SynchronizationSecret_UpdateKeyOnSoftDelete           SynchronizationSecret = "UpdateKeyOnSoftDelete"
	SynchronizationSecret_Url                             SynchronizationSecret = "Url"
	SynchronizationSecret_UserName                        SynchronizationSecret = "UserName"
	SynchronizationSecret_ValidateDomain                  SynchronizationSecret = "ValidateDomain"
)

func PossibleValuesForSynchronizationSecret() []string {
	return []string{
		string(SynchronizationSecret_AppKey),
		string(SynchronizationSecret_ApplicationTemplateIdentifier),
		string(SynchronizationSecret_AuthenticationType),
		string(SynchronizationSecret_BaseAddress),
		string(SynchronizationSecret_ClientIdentifier),
		string(SynchronizationSecret_ClientSecret),
		string(SynchronizationSecret_CompanyId),
		string(SynchronizationSecret_ConnectionString),
		string(SynchronizationSecret_ConsumerKey),
		string(SynchronizationSecret_ConsumerSecret),
		string(SynchronizationSecret_Domain),
		string(SynchronizationSecret_EnforceDomain),
		string(SynchronizationSecret_HardDeletesEnabled),
		string(SynchronizationSecret_InstanceName),
		string(SynchronizationSecret_None),
		string(SynchronizationSecret_OAuth2AccessToken),
		string(SynchronizationSecret_OAuth2AccessTokenCreationTime),
		string(SynchronizationSecret_OAuth2AuthorizationCode),
		string(SynchronizationSecret_OAuth2AuthorizationUri),
		string(SynchronizationSecret_OAuth2ClientId),
		string(SynchronizationSecret_OAuth2ClientSecret),
		string(SynchronizationSecret_OAuth2RedirectUri),
		string(SynchronizationSecret_OAuth2RefreshToken),
		string(SynchronizationSecret_OAuth2TokenExchangeUri),
		string(SynchronizationSecret_Password),
		string(SynchronizationSecret_PerformInboundEntitlementGrants),
		string(SynchronizationSecret_Sandbox),
		string(SynchronizationSecret_SandboxName),
		string(SynchronizationSecret_SecretToken),
		string(SynchronizationSecret_Server),
		string(SynchronizationSecret_SingleSignOnType),
		string(SynchronizationSecret_SkipOutOfScopeDeletions),
		string(SynchronizationSecret_SyncAgentADContainer),
		string(SynchronizationSecret_SyncAgentCompatibilityKey),
		string(SynchronizationSecret_SyncAll),
		string(SynchronizationSecret_SyncNotificationSettings),
		string(SynchronizationSecret_SynchronizationSchedule),
		string(SynchronizationSecret_SystemOfRecord),
		string(SynchronizationSecret_TestReferences),
		string(SynchronizationSecret_TokenExpiration),
		string(SynchronizationSecret_TokenKey),
		string(SynchronizationSecret_UpdateKeyOnSoftDelete),
		string(SynchronizationSecret_Url),
		string(SynchronizationSecret_UserName),
		string(SynchronizationSecret_ValidateDomain),
	}
}

func (s *SynchronizationSecret) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationSecret(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationSecret(input string) (*SynchronizationSecret, error) {
	vals := map[string]SynchronizationSecret{
		"appkey":                          SynchronizationSecret_AppKey,
		"applicationtemplateidentifier":   SynchronizationSecret_ApplicationTemplateIdentifier,
		"authenticationtype":              SynchronizationSecret_AuthenticationType,
		"baseaddress":                     SynchronizationSecret_BaseAddress,
		"clientidentifier":                SynchronizationSecret_ClientIdentifier,
		"clientsecret":                    SynchronizationSecret_ClientSecret,
		"companyid":                       SynchronizationSecret_CompanyId,
		"connectionstring":                SynchronizationSecret_ConnectionString,
		"consumerkey":                     SynchronizationSecret_ConsumerKey,
		"consumersecret":                  SynchronizationSecret_ConsumerSecret,
		"domain":                          SynchronizationSecret_Domain,
		"enforcedomain":                   SynchronizationSecret_EnforceDomain,
		"harddeletesenabled":              SynchronizationSecret_HardDeletesEnabled,
		"instancename":                    SynchronizationSecret_InstanceName,
		"none":                            SynchronizationSecret_None,
		"oauth2accesstoken":               SynchronizationSecret_OAuth2AccessToken,
		"oauth2accesstokencreationtime":   SynchronizationSecret_OAuth2AccessTokenCreationTime,
		"oauth2authorizationcode":         SynchronizationSecret_OAuth2AuthorizationCode,
		"oauth2authorizationuri":          SynchronizationSecret_OAuth2AuthorizationUri,
		"oauth2clientid":                  SynchronizationSecret_OAuth2ClientId,
		"oauth2clientsecret":              SynchronizationSecret_OAuth2ClientSecret,
		"oauth2redirecturi":               SynchronizationSecret_OAuth2RedirectUri,
		"oauth2refreshtoken":              SynchronizationSecret_OAuth2RefreshToken,
		"oauth2tokenexchangeuri":          SynchronizationSecret_OAuth2TokenExchangeUri,
		"password":                        SynchronizationSecret_Password,
		"performinboundentitlementgrants": SynchronizationSecret_PerformInboundEntitlementGrants,
		"sandbox":                         SynchronizationSecret_Sandbox,
		"sandboxname":                     SynchronizationSecret_SandboxName,
		"secrettoken":                     SynchronizationSecret_SecretToken,
		"server":                          SynchronizationSecret_Server,
		"singlesignontype":                SynchronizationSecret_SingleSignOnType,
		"skipoutofscopedeletions":         SynchronizationSecret_SkipOutOfScopeDeletions,
		"syncagentadcontainer":            SynchronizationSecret_SyncAgentADContainer,
		"syncagentcompatibilitykey":       SynchronizationSecret_SyncAgentCompatibilityKey,
		"syncall":                         SynchronizationSecret_SyncAll,
		"syncnotificationsettings":        SynchronizationSecret_SyncNotificationSettings,
		"synchronizationschedule":         SynchronizationSecret_SynchronizationSchedule,
		"systemofrecord":                  SynchronizationSecret_SystemOfRecord,
		"testreferences":                  SynchronizationSecret_TestReferences,
		"tokenexpiration":                 SynchronizationSecret_TokenExpiration,
		"tokenkey":                        SynchronizationSecret_TokenKey,
		"updatekeyonsoftdelete":           SynchronizationSecret_UpdateKeyOnSoftDelete,
		"url":                             SynchronizationSecret_Url,
		"username":                        SynchronizationSecret_UserName,
		"validatedomain":                  SynchronizationSecret_ValidateDomain,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationSecret(input)
	return &out, nil
}
