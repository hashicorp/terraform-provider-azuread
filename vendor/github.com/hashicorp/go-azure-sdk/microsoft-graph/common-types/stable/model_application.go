package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DirectoryObject = Application{}

type Application struct {
	// Defines custom behavior that a consuming service can use to call an app in specific contexts. For example,
	// applications that can render file streams can set the addIns property for its 'FileHandler' functionality. This lets
	// services like Microsoft 365 call the application in the context of a document the user is working on.
	AddIns *[]AddIn `json:"addIns,omitempty"`

	// Specifies settings for an application that implements a web API.
	Api *ApiApplication `json:"api,omitempty"`

	// The unique identifier for the application that is assigned to an application by Microsoft Entra ID. Not nullable.
	// Read-only. Alternate key. Supports $filter (eq).
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The appManagementPolicy applied to this application.
	AppManagementPolicies *[]AppManagementPolicy `json:"appManagementPolicies,omitempty"`

	// The collection of roles defined for the application. With app role assignments, these roles can be assigned to users,
	// groups, or service principals associated with other applications. Not nullable.
	AppRoles *[]AppRole `json:"appRoles,omitempty"`

	// Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne). Read-only. null if the app wasn't
	// created from an application template.
	ApplicationTemplateId nullable.Type[string] `json:"applicationTemplateId,omitempty"`

	// Specifies the certification status of the application.
	Certification *Certification `json:"certification,omitempty"`

	// The date and time the application was registered. The DateTimeOffset type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Read-only. Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderby.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Supports $filter (/$count eq 0, /$count ne 0). Read-only.
	CreatedOnBehalfOf *DirectoryObject `json:"createdOnBehalfOf,omitempty"`

	// OData ID for `CreatedOnBehalfOf` to bind to this entity
	CreatedOnBehalfOf_ODataBind *string `json:"createdOnBehalfOf@odata.bind,omitempty"`

	DefaultRedirectUri nullable.Type[string] `json:"defaultRedirectUri,omitempty"`

	// Free text field to provide a description of the application object to end users. The maximum allowed size is 1,024
	// characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value),
	// NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons include suspicious, abusive, or malicious
	// activity, or a violation of the Microsoft Services Agreement). Supports $filter (eq, ne, not).
	DisabledByMicrosoftStatus nullable.Type[string] `json:"disabledByMicrosoftStatus,omitempty"`

	// The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values),
	// $search, and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0).
	ExtensionProperties *[]ExtensionProperty `json:"extensionProperties,omitempty"`

	// Federated identities for applications. Supports $expand and $filter (startsWith, /$count eq 0, /$count ne 0).
	FederatedIdentityCredentials *[]FederatedIdentityCredential `json:"federatedIdentityCredentials,omitempty"`

	// Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this
	// attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Microsoft Entra
	// roles), All (this gets all of the security groups, distribution groups, and Microsoft Entra directory roles that the
	// signed-in user is a member of).
	GroupMembershipClaims nullable.Type[string] `json:"groupMembershipClaims,omitempty"`

	HomeRealmDiscoveryPolicies *[]HomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`

	// Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as
	// the prefix for the scopes you reference in your API's code, and it must be globally unique. You can use the default
	// value provided, which is in the form api://<appId>, or specify a more readable URI like https://contoso.com/api. For
	// more information on valid identifierUris patterns and best practices, see Microsoft Entra application registration
	// security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
	IdentifierUris *[]string `json:"identifierUris,omitempty"`

	// Basic profile information of the application such as app's marketing, support, terms of service and privacy statement
	// URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more
	// info, see How to: Add Terms of service and privacy statement for registered Microsoft Entra apps. Supports $filter
	// (eq, ne, not, ge, le, and eq on null values).
	Info *InformationalUrl `json:"info,omitempty"`

	// Specifies whether this application supports device authentication without a user. The default is false.
	IsDeviceOnlyAuthSupported nullable.Type[bool] `json:"isDeviceOnlyAuthSupported,omitempty"`

	// Specifies the fallback application type as public client, such as an installed application running on a mobile
	// device. The default value is false, which means the fallback application type is confidential client such as a web
	// app. There are certain scenarios where Microsoft Entra ID can't determine the client application type. For example,
	// the ROPC flow where it's configured without specifying a redirect URI. In those cases, Microsoft Entra ID interprets
	// the application type based on the value of this property.
	IsFallbackPublicClient nullable.Type[bool] `json:"isFallbackPublicClient,omitempty"`

	// The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
	KeyCredentials *[]KeyCredential `json:"keyCredentials,omitempty"`

	// The main logo for the application. Not nullable.
	Logo *string `json:"logo,omitempty"`

	// Specifies whether the Native Authentication APIs are enabled for the application. The possible values are: none and
	// all. Default is none. For more information, see Native Authentication.
	NativeAuthenticationApisEnabled *NativeAuthenticationApisEnabled `json:"nativeAuthenticationApisEnabled,omitempty"`

	// Notes relevant for the management of the application.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	OAuth2RequirePostResponse *bool `json:"oauth2RequirePostResponse,omitempty"`

	// Application developers can configure optional claims in their Microsoft Entra applications to specify the claims that
	// are sent to their application by the Microsoft security token service. For more information, see How to: Provide
	// optional claims to your app.
	OptionalClaims *OptionalClaims `json:"optionalClaims,omitempty"`

	// Directory objects that are owners of the application. Read-only. Nullable. Supports $expand, $filter (/$count eq 0,
	// /$count ne 0, /$count eq 1, /$count ne 1), and $select nested in $expand.
	Owners *[]DirectoryObject `json:"owners,omitempty"`

	// List of OData IDs for `Owners` to bind to this entity
	Owners_ODataBind *[]string `json:"owners@odata.bind,omitempty"`

	// Specifies parental control settings for an application.
	ParentalControlSettings *ParentalControlSettings `json:"parentalControlSettings,omitempty"`

	// The collection of password credentials associated with the application. Not nullable.
	PasswordCredentials *[]PasswordCredential `json:"passwordCredentials,omitempty"`

	// Specifies settings for installed clients such as desktop or mobile devices.
	PublicClient *PublicClientApplication `json:"publicClient,omitempty"`

	// The verified publisher domain for the application. Read-only. For more information, see How to: Configure an
	// application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
	PublisherDomain nullable.Type[string] `json:"publisherDomain,omitempty"`

	// Specifies whether this application requires Microsoft Entra ID to verify the signed authentication requests.
	RequestSignatureVerification *RequestSignatureVerification `json:"requestSignatureVerification,omitempty"`

	// Specifies the resources that the application needs to access. This property also specifies the set of delegated
	// permissions and application roles that it needs for each of those resources. This configuration of access to the
	// required resources drives the consent experience. No more than 50 resource services (APIs) can be configured.
	// Beginning mid-October 2021, the total number of required permissions must not exceed 400. For more information, see
	// Limits on requested permissions per app. Not nullable. Supports $filter (eq, not, ge, le).
	RequiredResourceAccess *[]RequiredResourceAccess `json:"requiredResourceAccess,omitempty"`

	// The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant
	// applications. Nullable.
	SamlMetadataUrl nullable.Type[string] `json:"samlMetadataUrl,omitempty"`

	// References application or service contact information from a Service or Asset Management database. Nullable.
	ServiceManagementReference nullable.Type[string] `json:"serviceManagementReference,omitempty"`

	// Specifies whether sensitive properties of a multitenant application should be locked for editing after the
	// application is provisioned in a tenant. Nullable. null by default.
	ServicePrincipalLockConfiguration *ServicePrincipalLockConfiguration `json:"servicePrincipalLockConfiguration,omitempty"`

	// Specifies the Microsoft accounts that are supported for the current application. The possible values are:
	// AzureADMyOrg (default), AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, and PersonalMicrosoftAccount. See
	// more in the table. The value of this object also limits the number of permissions an app can request. For more
	// information, see Limits on requested permissions per app. The value for this property has implications on other app
	// object properties. As a result, if you change this property, you might need to change other properties first. For
	// more information, see Validation differences for signInAudience.Supports $filter (eq, ne, not).
	SignInAudience nullable.Type[string] `json:"signInAudience,omitempty"`

	// Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes
	// and access tokens.
	Spa *SpaApplication `json:"spa,omitempty"`

	// Represents the capability for Microsoft Entra identity synchronization through the Microsoft Graph API.
	Synchronization *Synchronization `json:"synchronization,omitempty"`

	// Custom strings that can be used to categorize and identify the application. Not nullable. Strings added here will
	// also appear in the tags property of any associated service principals.Supports $filter (eq, not, ge, le, startsWith)
	// and $search.
	Tags *[]string `json:"tags,omitempty"`

	// Specifies the keyId of a public key from the keyCredentials collection. When configured, Microsoft Entra ID encrypts
	// all the tokens it emits by using the key this property points to. The application code that receives the encrypted
	// token must use the matching private key to decrypt the token before it can be used for the signed-in user.
	TokenEncryptionKeyId nullable.Type[string] `json:"tokenEncryptionKeyId,omitempty"`

	TokenIssuancePolicies *[]TokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies *[]TokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`

	// The unique identifier that can be assigned to an application and used as an alternate key. Immutable. Read-only.
	UniqueName nullable.Type[string] `json:"uniqueName,omitempty"`

	// Specifies the verified publisher of the application. For more information about how publisher verification helps
	// support application security, trustworthiness, and compliance, see Publisher verification.
	VerifiedPublisher *VerifiedPublisher `json:"verifiedPublisher,omitempty"`

	// Specifies settings for a web application.
	Web *WebApplication `json:"web,omitempty"`

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

func (s Application) DirectoryObject() BaseDirectoryObjectImpl {
	return BaseDirectoryObjectImpl{
		DeletedDateTime: s.DeletedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s Application) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Application{}

func (s Application) MarshalJSON() ([]byte, error) {
	type wrapper Application
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Application: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Application: %+v", err)
	}

	delete(decoded, "appId")
	delete(decoded, "applicationTemplateId")
	delete(decoded, "createdDateTime")
	delete(decoded, "createdOnBehalfOf")
	delete(decoded, "extensionProperties")
	delete(decoded, "owners")
	delete(decoded, "publisherDomain")
	delete(decoded, "uniqueName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.application"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Application: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Application{}

func (s *Application) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AddIns                            *[]AddIn                           `json:"addIns,omitempty"`
		Api                               *ApiApplication                    `json:"api,omitempty"`
		AppId                             nullable.Type[string]              `json:"appId,omitempty"`
		AppManagementPolicies             *[]AppManagementPolicy             `json:"appManagementPolicies,omitempty"`
		AppRoles                          *[]AppRole                         `json:"appRoles,omitempty"`
		ApplicationTemplateId             nullable.Type[string]              `json:"applicationTemplateId,omitempty"`
		Certification                     *Certification                     `json:"certification,omitempty"`
		CreatedDateTime                   nullable.Type[string]              `json:"createdDateTime,omitempty"`
		CreatedOnBehalfOf_ODataBind       *string                            `json:"createdOnBehalfOf@odata.bind,omitempty"`
		DefaultRedirectUri                nullable.Type[string]              `json:"defaultRedirectUri,omitempty"`
		Description                       nullable.Type[string]              `json:"description,omitempty"`
		DisabledByMicrosoftStatus         nullable.Type[string]              `json:"disabledByMicrosoftStatus,omitempty"`
		DisplayName                       nullable.Type[string]              `json:"displayName,omitempty"`
		ExtensionProperties               *[]ExtensionProperty               `json:"extensionProperties,omitempty"`
		FederatedIdentityCredentials      *[]FederatedIdentityCredential     `json:"federatedIdentityCredentials,omitempty"`
		GroupMembershipClaims             nullable.Type[string]              `json:"groupMembershipClaims,omitempty"`
		HomeRealmDiscoveryPolicies        *[]HomeRealmDiscoveryPolicy        `json:"homeRealmDiscoveryPolicies,omitempty"`
		IdentifierUris                    *[]string                          `json:"identifierUris,omitempty"`
		Info                              *InformationalUrl                  `json:"info,omitempty"`
		IsDeviceOnlyAuthSupported         nullable.Type[bool]                `json:"isDeviceOnlyAuthSupported,omitempty"`
		IsFallbackPublicClient            nullable.Type[bool]                `json:"isFallbackPublicClient,omitempty"`
		KeyCredentials                    *[]KeyCredential                   `json:"keyCredentials,omitempty"`
		Logo                              *string                            `json:"logo,omitempty"`
		NativeAuthenticationApisEnabled   *NativeAuthenticationApisEnabled   `json:"nativeAuthenticationApisEnabled,omitempty"`
		Notes                             nullable.Type[string]              `json:"notes,omitempty"`
		OAuth2RequirePostResponse         *bool                              `json:"oauth2RequirePostResponse,omitempty"`
		OptionalClaims                    *OptionalClaims                    `json:"optionalClaims,omitempty"`
		Owners_ODataBind                  *[]string                          `json:"owners@odata.bind,omitempty"`
		ParentalControlSettings           *ParentalControlSettings           `json:"parentalControlSettings,omitempty"`
		PasswordCredentials               *[]PasswordCredential              `json:"passwordCredentials,omitempty"`
		PublicClient                      *PublicClientApplication           `json:"publicClient,omitempty"`
		PublisherDomain                   nullable.Type[string]              `json:"publisherDomain,omitempty"`
		RequestSignatureVerification      *RequestSignatureVerification      `json:"requestSignatureVerification,omitempty"`
		RequiredResourceAccess            *[]RequiredResourceAccess          `json:"requiredResourceAccess,omitempty"`
		SamlMetadataUrl                   nullable.Type[string]              `json:"samlMetadataUrl,omitempty"`
		ServiceManagementReference        nullable.Type[string]              `json:"serviceManagementReference,omitempty"`
		ServicePrincipalLockConfiguration *ServicePrincipalLockConfiguration `json:"servicePrincipalLockConfiguration,omitempty"`
		SignInAudience                    nullable.Type[string]              `json:"signInAudience,omitempty"`
		Spa                               *SpaApplication                    `json:"spa,omitempty"`
		Synchronization                   *Synchronization                   `json:"synchronization,omitempty"`
		Tags                              *[]string                          `json:"tags,omitempty"`
		TokenEncryptionKeyId              nullable.Type[string]              `json:"tokenEncryptionKeyId,omitempty"`
		TokenIssuancePolicies             *[]TokenIssuancePolicy             `json:"tokenIssuancePolicies,omitempty"`
		TokenLifetimePolicies             *[]TokenLifetimePolicy             `json:"tokenLifetimePolicies,omitempty"`
		UniqueName                        nullable.Type[string]              `json:"uniqueName,omitempty"`
		VerifiedPublisher                 *VerifiedPublisher                 `json:"verifiedPublisher,omitempty"`
		Web                               *WebApplication                    `json:"web,omitempty"`
		DeletedDateTime                   nullable.Type[string]              `json:"deletedDateTime,omitempty"`
		Id                                *string                            `json:"id,omitempty"`
		ODataId                           *string                            `json:"@odata.id,omitempty"`
		ODataType                         *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AddIns = decoded.AddIns
	s.Api = decoded.Api
	s.AppId = decoded.AppId
	s.AppManagementPolicies = decoded.AppManagementPolicies
	s.AppRoles = decoded.AppRoles
	s.ApplicationTemplateId = decoded.ApplicationTemplateId
	s.Certification = decoded.Certification
	s.CreatedDateTime = decoded.CreatedDateTime
	s.CreatedOnBehalfOf_ODataBind = decoded.CreatedOnBehalfOf_ODataBind
	s.DefaultRedirectUri = decoded.DefaultRedirectUri
	s.Description = decoded.Description
	s.DisabledByMicrosoftStatus = decoded.DisabledByMicrosoftStatus
	s.DisplayName = decoded.DisplayName
	s.ExtensionProperties = decoded.ExtensionProperties
	s.FederatedIdentityCredentials = decoded.FederatedIdentityCredentials
	s.GroupMembershipClaims = decoded.GroupMembershipClaims
	s.HomeRealmDiscoveryPolicies = decoded.HomeRealmDiscoveryPolicies
	s.IdentifierUris = decoded.IdentifierUris
	s.Info = decoded.Info
	s.IsDeviceOnlyAuthSupported = decoded.IsDeviceOnlyAuthSupported
	s.IsFallbackPublicClient = decoded.IsFallbackPublicClient
	s.KeyCredentials = decoded.KeyCredentials
	s.Logo = decoded.Logo
	s.NativeAuthenticationApisEnabled = decoded.NativeAuthenticationApisEnabled
	s.Notes = decoded.Notes
	s.OAuth2RequirePostResponse = decoded.OAuth2RequirePostResponse
	s.OptionalClaims = decoded.OptionalClaims
	s.Owners_ODataBind = decoded.Owners_ODataBind
	s.ParentalControlSettings = decoded.ParentalControlSettings
	s.PasswordCredentials = decoded.PasswordCredentials
	s.PublicClient = decoded.PublicClient
	s.PublisherDomain = decoded.PublisherDomain
	s.RequestSignatureVerification = decoded.RequestSignatureVerification
	s.RequiredResourceAccess = decoded.RequiredResourceAccess
	s.SamlMetadataUrl = decoded.SamlMetadataUrl
	s.ServiceManagementReference = decoded.ServiceManagementReference
	s.ServicePrincipalLockConfiguration = decoded.ServicePrincipalLockConfiguration
	s.SignInAudience = decoded.SignInAudience
	s.Spa = decoded.Spa
	s.Synchronization = decoded.Synchronization
	s.Tags = decoded.Tags
	s.TokenEncryptionKeyId = decoded.TokenEncryptionKeyId
	s.TokenIssuancePolicies = decoded.TokenIssuancePolicies
	s.TokenLifetimePolicies = decoded.TokenLifetimePolicies
	s.UniqueName = decoded.UniqueName
	s.VerifiedPublisher = decoded.VerifiedPublisher
	s.Web = decoded.Web
	s.DeletedDateTime = decoded.DeletedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Application into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdOnBehalfOf"]; ok {
		impl, err := UnmarshalDirectoryObjectImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedOnBehalfOf' for 'Application': %+v", err)
		}
		s.CreatedOnBehalfOf = &impl
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
				return fmt.Errorf("unmarshaling index %d field 'Owners' for 'Application': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Owners = &output
	}

	return nil
}
