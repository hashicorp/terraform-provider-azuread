// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package applications

const (
	applicationResourceName = "azuread_application"
)

const (
	AppRoleAllowedMemberTypeApplication = "Application"
	AppRoleAllowedMemberTypeUser        = "User"
)

var possibleValuesForAppRoleAllowedMemberType = []string{AppRoleAllowedMemberTypeApplication, AppRoleAllowedMemberTypeUser}

const (
	GroupMembershipClaimAll              = "All"
	GroupMembershipClaimNone             = "None"
	GroupMembershipClaimApplicationGroup = "ApplicationGroup"
	GroupMembershipClaimDirectoryRole    = "DirectoryRole"
	GroupMembershipClaimSecurityGroup    = "SecurityGroup"
)

var possibleValuesForGroupMembershipClaim = []string{GroupMembershipClaimAll, GroupMembershipClaimNone, GroupMembershipClaimApplicationGroup, GroupMembershipClaimDirectoryRole, GroupMembershipClaimSecurityGroup}

const (
	RedirectUriTypePublicClient = "PublicClient"
	RedirectUriTypeSPA          = "SPA"
	RedirectUriTypeWeb          = "Web"
)

var possibleValuesForRedirectUriType = []string{RedirectUriTypePublicClient, RedirectUriTypeSPA, RedirectUriTypeWeb}

const (
	PermissionScopeTypeAdmin = "Admin"
	PermissionScopeTypeUser  = "User"
)

var possibleValuesForPermissionScopeType = []string{PermissionScopeTypeAdmin, PermissionScopeTypeUser}

const (
	ResourceAccessTypeRole  = "Role"
	ResourceAccessTypeScope = "Scope"
)

var possibleValuesForResourceAccessType = []string{ResourceAccessTypeRole, ResourceAccessTypeScope}

const (
	SignInAudienceAzureADMyOrg                       = "AzureADMyOrg"
	SignInAudienceAzureADMultipleOrgs                = "AzureADMultipleOrgs"
	SignInAudienceAzureADandPersonalMicrosoftAccount = "AzureADandPersonalMicrosoftAccount"
	SignInAudiencePersonalMicrosoftAccount           = "PersonalMicrosoftAccount"
)

var possibleValuesForSignInAudience = []string{SignInAudienceAzureADMyOrg, SignInAudienceAzureADMultipleOrgs, SignInAudienceAzureADandPersonalMicrosoftAccount, SignInAudiencePersonalMicrosoftAccount}
