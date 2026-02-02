package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = ServicePrincipal{}

type ServicePrincipal struct {
	// true if the service principal account is enabled; otherwise, false. If set to false, then no users are able to sign
	// in to this app, even if they're assigned to it. Supports $filter (eq, ne, not, in).
	AccountEnabled nullable.Type[bool] `json:"accountEnabled,omitempty"`

	// Defines custom behavior that a consuming service can use to call an app in specific contexts. For example,
	// applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This lets
	// services like Microsoft 365 call the application in the context of a document the user is working on.
	AddIns *[]AddIn `json:"addIns,omitempty"`

	// Used to retrieve service principals by subscription, identify resource group and full resource IDs for managed
	// identities. Supports $filter (eq, not, ge, le, startsWith).
	AlternativeNames *[]string `json:"alternativeNames,omitempty"`

	// The description exposed by the associated application.
	AppDescription nullable.Type[string] `json:"appDescription,omitempty"`

	// The display name exposed by the associated application. Maximum length is 256 characters.
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The unique identifier for the associated application (its appId property). Alternate key. Supports $filter (eq, ne,
	// not, in, startsWith).
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The appManagementPolicy applied to this service principal.
	AppManagementPolicies *[]AppManagementPolicy `json:"appManagementPolicies,omitempty"`

	// Contains the tenant ID where the application is registered. This is applicable only to service principals backed by
	// applications. Supports $filter (eq, ne, NOT, ge, le).
	AppOwnerOrganizationId nullable.Type[string] `json:"appOwnerOrganizationId,omitempty"`

	// App role assignments for this app or service, granted to users, groups, and other service principals.Supports
	// $expand.
	AppRoleAssignedTo *[]AppRoleAssignment `json:"appRoleAssignedTo,omitempty"`

	// Specifies whether users or other service principals need to be granted an app role assignment for this service
	// principal before users can sign in or apps can get tokens. The default value is false. Not nullable. Supports $filter
	// (eq, ne, NOT).
	AppRoleAssignmentRequired *bool `json:"appRoleAssignmentRequired,omitempty"`

	// App role assignment for another app or service, granted to this service principal. Supports $expand.
	AppRoleAssignments *[]AppRoleAssignment `json:"appRoleAssignments,omitempty"`

	// The roles exposed by the application, which this service principal represents. For more information, see the appRoles
	// property definition on the application entity. Not nullable.
	AppRoles *[]AppRole `json:"appRoles,omitempty"`

	// Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne). Read-only. null if the app wasn't
	// created from an application template.
	ApplicationTemplateId nullable.Type[string] `json:"applicationTemplateId,omitempty"`

	// The claimsMappingPolicies assigned to this service principal. Supports $expand.
	ClaimsMappingPolicies *[]ClaimsMappingPolicy `json:"claimsMappingPolicies,omitempty"`

	// A claims policy that allows application admins to customize the claims that will be emitted in tokens affected by
	// this policy.
	ClaimsPolicy *CustomClaimsPolicy `json:"claimsPolicy,omitempty"`

	// Directory objects created by this service principal. Read-only. Nullable.
	CreatedObjects *[]DirectoryObject `json:"createdObjects,omitempty"`

	// List of OData IDs for `CreatedObjects` to bind to this entity
	CreatedObjects_ODataBind *[]string `json:"createdObjects@odata.bind,omitempty"`

	// An open complex type that holds the value of a custom security attribute that is assigned to a directory object.
	// Nullable. Returned only on $select. Supports $filter (eq, ne, not, startsWith). Filter value is case sensitive.To
	// read this property, the calling app must be assigned the CustomSecAttributeAssignment.Read.All permission. To write
	// this property, the calling app must be assigned the CustomSecAttributeAssignment.ReadWrite.All permissions. To read
	// or write this property in delegated scenarios, the admin must be assigned the Attribute Assignment Administrator
	// role.
	CustomSecurityAttributes *CustomSecurityAttributeValue `json:"customSecurityAttributes,omitempty"`

	// The permission classifications for delegated permissions exposed by the app that this service principal represents.
	// Supports $expand.
	DelegatedPermissionClassifications *[]DelegatedPermissionClassification `json:"delegatedPermissionClassifications,omitempty"`

	// Free text field to provide an internal end-user facing description of the service principal. End-user portals such
	// MyApps displays the application description in this field. The maximum allowed size is 1,024 characters. Supports
	// $filter (eq, ne, not, ge, le, startsWith) and $search.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value),
	// NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious
	// activity, or a violation of the Microsoft Services Agreement). Supports $filter (eq, ne, not).
	DisabledByMicrosoftStatus nullable.Type[string] `json:"disabledByMicrosoftStatus,omitempty"`

	// The display name for the service principal. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null
	// values), $search, and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Endpoints available for discovery. Services like Sharepoint populate this property with a tenant specific SharePoint
	// endpoints that other applications can discover and use in their experiences.
	Endpoints *[]Endpoint `json:"endpoints,omitempty"`

	// Deprecated. Don't use.
	ErrorUrl nullable.Type[string] `json:"errorUrl,omitempty"`

	FederatedIdentityCredentials *[]FederatedIdentityCredential `json:"federatedIdentityCredentials,omitempty"`

	// The homeRealmDiscoveryPolicies assigned to this service principal. Supports $expand.
	HomeRealmDiscoveryPolicies *[]HomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`

	// Home page or landing page of the application.
	Homepage nullable.Type[string] `json:"homepage,omitempty"`

	// Basic profile information of the acquired application such as app's marketing, support, terms of service and privacy
	// statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience.
	// For more info, see How to: Add Terms of service and privacy statement for registered Microsoft Entra apps. Supports
	// $filter (eq, ne, not, ge, le, and eq on null values).
	Info *InformationalUrl `json:"info,omitempty"`

	// The collection of key credentials associated with the service principal. Not nullable. Supports $filter (eq, not, ge,
	// le).
	KeyCredentials *[]KeyCredential `json:"keyCredentials,omitempty"`

	LicenseDetails *[]LicenseDetails `json:"licenseDetails,omitempty"`

	// Specifies the URL where the service provider redirects the user to Microsoft Entra ID to authenticate. Microsoft
	// Entra ID uses the URL to launch the application from Microsoft 365 or the Microsoft Entra My Apps. When blank,
	// Microsoft Entra ID performs IdP-initiated sign-on for applications configured with SAML-based single sign-on. The
	// user launches the application from Microsoft 365, the Microsoft Entra My Apps, or the Microsoft Entra SSO URL.
	LoginUrl nullable.Type[string] `json:"loginUrl,omitempty"`

	// Specifies the URL that the Microsoft's authorization service uses to sign out a user using OpenId Connect
	// front-channel, back-channel, or SAML sign out protocols.
	LogoutUrl nullable.Type[string] `json:"logoutUrl,omitempty"`

	// Roles that this service principal is a member of. HTTP Methods: GET Read-only. Nullable. Supports $expand.
	MemberOf *[]DirectoryObject `json:"memberOf,omitempty"`

	// List of OData IDs for `MemberOf` to bind to this entity
	MemberOf_ODataBind *[]string `json:"memberOf@odata.bind,omitempty"`

	// Free text field to capture information about the service principal, typically used for operational purposes. Maximum
	// allowed size is 1,024 characters.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// Specifies the list of email addresses where Microsoft Entra ID sends a notification when the active certificate is
	// near the expiration date. This is only for the certificates used to sign the SAML token issued for Microsoft Entra
	// Gallery applications.
	NotificationEmailAddresses *[]string `json:"notificationEmailAddresses,omitempty"`

	// Delegated permission grants authorizing this service principal to access an API on behalf of a signed-in user.
	// Read-only. Nullable.
	OAuth2PermissionGrants *[]OAuth2PermissionGrant `json:"oauth2PermissionGrants,omitempty"`

	// Directory objects that are owned by this service principal. Read-only. Nullable. Supports $expand and $filter
	// (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
	OwnedObjects *[]DirectoryObject `json:"ownedObjects,omitempty"`

	// List of OData IDs for `OwnedObjects` to bind to this entity
	OwnedObjects_ODataBind *[]string `json:"ownedObjects@odata.bind,omitempty"`

	// Directory objects that are owners of this servicePrincipal. The owners are a set of nonadmin users or
	// servicePrincipals who are allowed to modify this object. Supports $expand and $filter (/$count eq 0, /$count ne 0,
	// /$count eq 1, /$count ne 1).
	Owners *[]DirectoryObject `json:"owners,omitempty"`

	// List of OData IDs for `Owners` to bind to this entity
	Owners_ODataBind *[]string `json:"owners@odata.bind,omitempty"`

	// The collection of password credentials associated with the service principal. Not nullable.
	PasswordCredentials *[]PasswordCredential `json:"passwordCredentials,omitempty"`

	// The collection for settings related to password single sign-on. Use $select=passwordSingleSignOnSettings to read the
	// property. Read-only for applicationTemplates except for custom applicationTemplates.
	PasswordSingleSignOnSettings *PasswordSingleSignOnSettings `json:"passwordSingleSignOnSettings,omitempty"`

	PermissionGrantPreApprovalPolicies *[]PermissionGrantPreApprovalPolicy `json:"permissionGrantPreApprovalPolicies,omitempty"`

	// Specifies the single sign-on mode configured for this application. Microsoft Entra ID uses the preferred single
	// sign-on mode to launch the application from Microsoft 365 or the Microsoft Entra My Apps. The supported values are
	// password, saml, notSupported, and oidc. Note: This field might be null for older SAML apps and for OIDC applications
	// where it isn't set automatically.
	PreferredSingleSignOnMode nullable.Type[string] `json:"preferredSingleSignOnMode,omitempty"`

	// Specifies the expiration date of the keyCredential used for token signing, marked by
	// preferredTokenSigningKeyThumbprint. Updating this attribute isn't currently supported. For details, see
	// ServicePrincipal property differences.
	PreferredTokenSigningKeyEndDateTime nullable.Type[string] `json:"preferredTokenSigningKeyEndDateTime,omitempty"`

	// This property can be used on SAML applications (apps that have preferredSingleSignOnMode set to saml) to control
	// which certificate is used to sign the SAML responses. For applications that aren't SAML, don't write or otherwise
	// rely on this property.
	PreferredTokenSigningKeyThumbprint nullable.Type[string] `json:"preferredTokenSigningKeyThumbprint,omitempty"`

	// The delegated permissions exposed by the application. For more information, see the oauth2PermissionScopes property
	// on the application entity's api property. Not nullable. Note: This property is named oauth2PermissionScopes in v1.0.
	PublishedPermissionScopes *[]PermissionScope `json:"publishedPermissionScopes,omitempty"`

	// The name of the Microsoft Entra tenant that published the application.
	PublisherName nullable.Type[string] `json:"publisherName,omitempty"`

	// The remoteDesktopSecurityConfiguration object applied to this service principal. Supports $filter (eq) for
	// isRemoteDesktopProtocolEnabled property.
	RemoteDesktopSecurityConfiguration *RemoteDesktopSecurityConfiguration `json:"remoteDesktopSecurityConfiguration,omitempty"`

	// The URLs that user tokens are sent to for sign in with the associated application, or the redirect URIs that OAuth
	// 2.0 authorization codes and access tokens are sent to for the associated application. Not nullable.
	ReplyUrls *[]string `json:"replyUrls,omitempty"`

	// The url where the service exposes SAML metadata for federation.
	SamlMetadataUrl nullable.Type[string] `json:"samlMetadataUrl,omitempty"`

	// The collection for settings related to saml single sign-on.
	SamlSingleSignOnSettings *SamlSingleSignOnSettings `json:"samlSingleSignOnSettings,omitempty"`

	// Contains the list of identifiersUris, copied over from the associated application. More values can be added to hybrid
	// applications. These values can be used to identify the permissions exposed by this app within Microsoft Entra ID. For
	// example,Client apps can specify a resource URI that is based on the values of this property to acquire an access
	// token, which is the URI returned in the 'aud' claim.The any operator is required for filter expressions on
	// multi-valued properties. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
	ServicePrincipalNames *[]string `json:"servicePrincipalNames,omitempty"`

	// Identifies if the service principal represents an application or a managed identity. This is set by Microsoft Entra
	// ID internally. For a service principal that represents an application this is set as Application. For a service
	// principal that represents a managed identity this is set as ManagedIdentity. The SocialIdp type is for internal use.
	ServicePrincipalType nullable.Type[string] `json:"servicePrincipalType,omitempty"`

	// Specifies the Microsoft accounts that are supported for the current application. Read-only. Supported values
	// are:AzureADMyOrg: Users with a Microsoft work or school account in my organization's Microsoft Entra tenant
	// (single-tenant).AzureADMultipleOrgs: Users with a Microsoft work or school account in any organization's Microsoft
	// Entra tenant (multitenant).AzureADandPersonalMicrosoftAccount: Users with a personal Microsoft account, or a work or
	// school account in any organization's Microsoft Entra tenant.PersonalMicrosoftAccount: Users with a personal Microsoft
	// account only.
	SignInAudience nullable.Type[string] `json:"signInAudience,omitempty"`

	// Represents the capability for Microsoft Entra identity synchronization through the Microsoft Graph API.
	Synchronization *Synchronization `json:"synchronization,omitempty"`

	// Custom strings that can be used to categorize and identify the service principal. Not nullable. The value is the
	// union of strings set here and on the associated application entity's tags property.Supports $filter (eq, not, ge, le,
	// startsWith).
	Tags *[]string `json:"tags,omitempty"`

	// Specifies the keyId of a public key from the keyCredentials collection. When configured, Microsoft Entra ID issues
	// tokens for this application encrypted using the key specified by this property. The application code that receives
	// the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in
	// user.
	TokenEncryptionKeyId nullable.Type[string] `json:"tokenEncryptionKeyId,omitempty"`

	// The tokenIssuancePolicies assigned to this service principal. Supports $expand.
	TokenIssuancePolicies *[]TokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`

	// The tokenLifetimePolicies assigned to this service principal. Supports $expand.
	TokenLifetimePolicies *[]TokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`

	TransitiveMemberOf *[]DirectoryObject `json:"transitiveMemberOf,omitempty"`

	// List of OData IDs for `TransitiveMemberOf` to bind to this entity
	TransitiveMemberOf_ODataBind *[]string `json:"transitiveMemberOf@odata.bind,omitempty"`

	// Specifies the verified publisher of the application that's linked to this service principal.
	VerifiedPublisher *VerifiedPublisher `json:"verifiedPublisher,omitempty"`

	// Fields inherited from DirectoryObject

	// Date and time when this object was deleted. Always null when the object hasn't been deleted.
	DeletedDateTime nullable.Type[string] `json:"deletedDateTime,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ServicePrincipal) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s ServicePrincipal) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServicePrincipal{}

func (s ServicePrincipal) MarshalJSON() ([]byte, error) {
	type wrapper ServicePrincipal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServicePrincipal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServicePrincipal: %+v", err)
	}

	delete(decoded, "applicationTemplateId")
	delete(decoded, "createdObjects")
	delete(decoded, "memberOf")
	delete(decoded, "oauth2PermissionGrants")
	delete(decoded, "ownedObjects")
	delete(decoded, "signInAudience")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.servicePrincipal"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServicePrincipal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ServicePrincipal{}

func (s *ServicePrincipal) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccountEnabled                      nullable.Type[bool]                  `json:"accountEnabled,omitempty"`
		AddIns                              *[]AddIn                             `json:"addIns,omitempty"`
		AlternativeNames                    *[]string                            `json:"alternativeNames,omitempty"`
		AppDescription                      nullable.Type[string]                `json:"appDescription,omitempty"`
		AppDisplayName                      nullable.Type[string]                `json:"appDisplayName,omitempty"`
		AppId                               nullable.Type[string]                `json:"appId,omitempty"`
		AppManagementPolicies               *[]AppManagementPolicy               `json:"appManagementPolicies,omitempty"`
		AppOwnerOrganizationId              nullable.Type[string]                `json:"appOwnerOrganizationId,omitempty"`
		AppRoleAssignedTo                   *[]AppRoleAssignment                 `json:"appRoleAssignedTo,omitempty"`
		AppRoleAssignmentRequired           *bool                                `json:"appRoleAssignmentRequired,omitempty"`
		AppRoleAssignments                  *[]AppRoleAssignment                 `json:"appRoleAssignments,omitempty"`
		AppRoles                            *[]AppRole                           `json:"appRoles,omitempty"`
		ApplicationTemplateId               nullable.Type[string]                `json:"applicationTemplateId,omitempty"`
		ClaimsMappingPolicies               *[]ClaimsMappingPolicy               `json:"claimsMappingPolicies,omitempty"`
		ClaimsPolicy                        *CustomClaimsPolicy                  `json:"claimsPolicy,omitempty"`
		CreatedObjects_ODataBind            *[]string                            `json:"createdObjects@odata.bind,omitempty"`
		CustomSecurityAttributes            *CustomSecurityAttributeValue        `json:"customSecurityAttributes,omitempty"`
		DelegatedPermissionClassifications  *[]DelegatedPermissionClassification `json:"delegatedPermissionClassifications,omitempty"`
		Description                         nullable.Type[string]                `json:"description,omitempty"`
		DisabledByMicrosoftStatus           nullable.Type[string]                `json:"disabledByMicrosoftStatus,omitempty"`
		DisplayName                         nullable.Type[string]                `json:"displayName,omitempty"`
		Endpoints                           *[]Endpoint                          `json:"endpoints,omitempty"`
		ErrorUrl                            nullable.Type[string]                `json:"errorUrl,omitempty"`
		FederatedIdentityCredentials        *[]FederatedIdentityCredential       `json:"federatedIdentityCredentials,omitempty"`
		HomeRealmDiscoveryPolicies          *[]HomeRealmDiscoveryPolicy          `json:"homeRealmDiscoveryPolicies,omitempty"`
		Homepage                            nullable.Type[string]                `json:"homepage,omitempty"`
		Info                                *InformationalUrl                    `json:"info,omitempty"`
		KeyCredentials                      *[]KeyCredential                     `json:"keyCredentials,omitempty"`
		LicenseDetails                      *[]LicenseDetails                    `json:"licenseDetails,omitempty"`
		LoginUrl                            nullable.Type[string]                `json:"loginUrl,omitempty"`
		LogoutUrl                           nullable.Type[string]                `json:"logoutUrl,omitempty"`
		MemberOf_ODataBind                  *[]string                            `json:"memberOf@odata.bind,omitempty"`
		Notes                               nullable.Type[string]                `json:"notes,omitempty"`
		NotificationEmailAddresses          *[]string                            `json:"notificationEmailAddresses,omitempty"`
		OAuth2PermissionGrants              *[]OAuth2PermissionGrant             `json:"oauth2PermissionGrants,omitempty"`
		OwnedObjects_ODataBind              *[]string                            `json:"ownedObjects@odata.bind,omitempty"`
		Owners_ODataBind                    *[]string                            `json:"owners@odata.bind,omitempty"`
		PasswordCredentials                 *[]PasswordCredential                `json:"passwordCredentials,omitempty"`
		PasswordSingleSignOnSettings        *PasswordSingleSignOnSettings        `json:"passwordSingleSignOnSettings,omitempty"`
		PermissionGrantPreApprovalPolicies  *[]PermissionGrantPreApprovalPolicy  `json:"permissionGrantPreApprovalPolicies,omitempty"`
		PreferredSingleSignOnMode           nullable.Type[string]                `json:"preferredSingleSignOnMode,omitempty"`
		PreferredTokenSigningKeyEndDateTime nullable.Type[string]                `json:"preferredTokenSigningKeyEndDateTime,omitempty"`
		PreferredTokenSigningKeyThumbprint  nullable.Type[string]                `json:"preferredTokenSigningKeyThumbprint,omitempty"`
		PublishedPermissionScopes           *[]PermissionScope                   `json:"publishedPermissionScopes,omitempty"`
		PublisherName                       nullable.Type[string]                `json:"publisherName,omitempty"`
		RemoteDesktopSecurityConfiguration  *RemoteDesktopSecurityConfiguration  `json:"remoteDesktopSecurityConfiguration,omitempty"`
		ReplyUrls                           *[]string                            `json:"replyUrls,omitempty"`
		SamlMetadataUrl                     nullable.Type[string]                `json:"samlMetadataUrl,omitempty"`
		SamlSingleSignOnSettings            *SamlSingleSignOnSettings            `json:"samlSingleSignOnSettings,omitempty"`
		ServicePrincipalNames               *[]string                            `json:"servicePrincipalNames,omitempty"`
		ServicePrincipalType                nullable.Type[string]                `json:"servicePrincipalType,omitempty"`
		SignInAudience                      nullable.Type[string]                `json:"signInAudience,omitempty"`
		Synchronization                     *Synchronization                     `json:"synchronization,omitempty"`
		Tags                                *[]string                            `json:"tags,omitempty"`
		TokenEncryptionKeyId                nullable.Type[string]                `json:"tokenEncryptionKeyId,omitempty"`
		TokenIssuancePolicies               *[]TokenIssuancePolicy               `json:"tokenIssuancePolicies,omitempty"`
		TokenLifetimePolicies               *[]TokenLifetimePolicy               `json:"tokenLifetimePolicies,omitempty"`
		TransitiveMemberOf_ODataBind        *[]string                            `json:"transitiveMemberOf@odata.bind,omitempty"`
		VerifiedPublisher                   *VerifiedPublisher                   `json:"verifiedPublisher,omitempty"`
		DeletedDateTime                     nullable.Type[string]                `json:"deletedDateTime,omitempty"`
		Id                                  *string                              `json:"id,omitempty"`
		ODataId                             *string                              `json:"@odata.id,omitempty"`
		ODataType                           *string                              `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccountEnabled = decoded.AccountEnabled
	s.AddIns = decoded.AddIns
	s.AlternativeNames = decoded.AlternativeNames
	s.AppDescription = decoded.AppDescription
	s.AppDisplayName = decoded.AppDisplayName
	s.AppId = decoded.AppId
	s.AppManagementPolicies = decoded.AppManagementPolicies
	s.AppOwnerOrganizationId = decoded.AppOwnerOrganizationId
	s.AppRoleAssignedTo = decoded.AppRoleAssignedTo
	s.AppRoleAssignmentRequired = decoded.AppRoleAssignmentRequired
	s.AppRoleAssignments = decoded.AppRoleAssignments
	s.AppRoles = decoded.AppRoles
	s.ApplicationTemplateId = decoded.ApplicationTemplateId
	s.ClaimsMappingPolicies = decoded.ClaimsMappingPolicies
	s.ClaimsPolicy = decoded.ClaimsPolicy
	s.CreatedObjects_ODataBind = decoded.CreatedObjects_ODataBind
	s.CustomSecurityAttributes = decoded.CustomSecurityAttributes
	s.DelegatedPermissionClassifications = decoded.DelegatedPermissionClassifications
	s.Description = decoded.Description
	s.DisabledByMicrosoftStatus = decoded.DisabledByMicrosoftStatus
	s.DisplayName = decoded.DisplayName
	s.Endpoints = decoded.Endpoints
	s.ErrorUrl = decoded.ErrorUrl
	s.FederatedIdentityCredentials = decoded.FederatedIdentityCredentials
	s.HomeRealmDiscoveryPolicies = decoded.HomeRealmDiscoveryPolicies
	s.Homepage = decoded.Homepage
	s.Info = decoded.Info
	s.KeyCredentials = decoded.KeyCredentials
	s.LicenseDetails = decoded.LicenseDetails
	s.LoginUrl = decoded.LoginUrl
	s.LogoutUrl = decoded.LogoutUrl
	s.MemberOf_ODataBind = decoded.MemberOf_ODataBind
	s.Notes = decoded.Notes
	s.NotificationEmailAddresses = decoded.NotificationEmailAddresses
	s.OAuth2PermissionGrants = decoded.OAuth2PermissionGrants
	s.OwnedObjects_ODataBind = decoded.OwnedObjects_ODataBind
	s.Owners_ODataBind = decoded.Owners_ODataBind
	s.PasswordCredentials = decoded.PasswordCredentials
	s.PasswordSingleSignOnSettings = decoded.PasswordSingleSignOnSettings
	s.PermissionGrantPreApprovalPolicies = decoded.PermissionGrantPreApprovalPolicies
	s.PreferredSingleSignOnMode = decoded.PreferredSingleSignOnMode
	s.PreferredTokenSigningKeyEndDateTime = decoded.PreferredTokenSigningKeyEndDateTime
	s.PreferredTokenSigningKeyThumbprint = decoded.PreferredTokenSigningKeyThumbprint
	s.PublishedPermissionScopes = decoded.PublishedPermissionScopes
	s.PublisherName = decoded.PublisherName
	s.RemoteDesktopSecurityConfiguration = decoded.RemoteDesktopSecurityConfiguration
	s.ReplyUrls = decoded.ReplyUrls
	s.SamlMetadataUrl = decoded.SamlMetadataUrl
	s.SamlSingleSignOnSettings = decoded.SamlSingleSignOnSettings
	s.ServicePrincipalNames = decoded.ServicePrincipalNames
	s.ServicePrincipalType = decoded.ServicePrincipalType
	s.SignInAudience = decoded.SignInAudience
	s.Synchronization = decoded.Synchronization
	s.Tags = decoded.Tags
	s.TokenEncryptionKeyId = decoded.TokenEncryptionKeyId
	s.TokenIssuancePolicies = decoded.TokenIssuancePolicies
	s.TokenLifetimePolicies = decoded.TokenLifetimePolicies
	s.TransitiveMemberOf_ODataBind = decoded.TransitiveMemberOf_ODataBind
	s.VerifiedPublisher = decoded.VerifiedPublisher
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServicePrincipal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdObjects"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling CreatedObjects into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'CreatedObjects' for 'ServicePrincipal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.CreatedObjects = &output
	}

	if v, ok := temp["memberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MemberOf' for 'ServicePrincipal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MemberOf = &output
	}

	if v, ok := temp["ownedObjects"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling OwnedObjects into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'OwnedObjects' for 'ServicePrincipal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.OwnedObjects = &output
	}

	if v, ok := temp["owners"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Owners into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Owners' for 'ServicePrincipal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Owners = &output
	}

	if v, ok := temp["transitiveMemberOf"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling TransitiveMemberOf into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'TransitiveMemberOf' for 'ServicePrincipal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.TransitiveMemberOf = &output
	}

	return nil
}
