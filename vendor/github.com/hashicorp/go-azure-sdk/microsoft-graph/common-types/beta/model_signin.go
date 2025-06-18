package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SignIn{}

type SignIn struct {
	// Represents details about the agentic sign-in. Includes the type of agent as well as parentAppID in some cases
	Agent *AgenticAgentSignIn `json:"agent,omitempty"`

	// The application name displayed in the Microsoft Entra admin center. Supports $filter (eq, startsWith).
	AppDisplayName nullable.Type[string] `json:"appDisplayName,omitempty"`

	// The application identifier in Microsoft Entra ID. Supports $filter (eq).
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The identifier of the tenant that owns the client application. Supports $filter (eq).
	AppOwnerTenantId nullable.Type[string] `json:"appOwnerTenantId,omitempty"`

	// Token protection creates a cryptographically secure tie between the token and the device it's issued to. This field
	// indicates whether the app token was bound to the device.
	AppTokenProtectionStatus *TokenProtectionStatus `json:"appTokenProtectionStatus,omitempty"`

	// A list of conditional access policies that the corresponding sign-in activity triggers. Apps need more Conditional
	// Access-related privileges to read the details of this property. For more information, see Permissions for viewing
	// applied conditional access (CA) policies in sign-ins.
	AppliedConditionalAccessPolicies *[]AppliedConditionalAccessPolicy `json:"appliedConditionalAccessPolicies,omitempty"`

	// Detailed information about the listeners, such as Azure Logic Apps and Azure Functions, which the corresponding
	// events in the sign-in event triggered.
	AppliedEventListeners *[]AppliedAuthenticationEventListener `json:"appliedEventListeners,omitempty"`

	// Provides details about the app and device used during a Microsoft Entra authentication step.
	AuthenticationAppDeviceDetails *AuthenticationAppDeviceDetails `json:"authenticationAppDeviceDetails,omitempty"`

	// Provides details of the Microsoft Entra policies applied to a user and client authentication app during an
	// authentication step.
	AuthenticationAppPolicyEvaluationDetails *[]AuthenticationAppPolicyDetails `json:"authenticationAppPolicyEvaluationDetails,omitempty"`

	// Contains a collection of values that represent the conditional access authentication contexts applied to the sign-in.
	AuthenticationContextClassReferences *[]AuthenticationContext `json:"authenticationContextClassReferences,omitempty"`

	// The result of the authentication attempt and more details on the authentication method.
	AuthenticationDetails *[]AuthenticationDetail `json:"authenticationDetails,omitempty"`

	// The authentication methods used. Possible values: SMS, Authenticator App, App Verification code, Password, FIDO, PTA,
	// or PHS.
	AuthenticationMethodsUsed *[]string `json:"authenticationMethodsUsed,omitempty"`

	// More authentication processing details, such as the agent name for PTA and PHS, or a server or farm name for
	// federated authentication.
	AuthenticationProcessingDetails *[]KeyValue `json:"authenticationProcessingDetails,omitempty"`

	// Lists the protocol type or grant type used in the authentication. The possible values are: none, oAuth2, ropc,
	// wsFederation, saml20, deviceCode, unknownFutureValue, authenticationTransfer, nativeAuth,
	// implicitAccessTokenAndGetResponseMode, implicitIdTokenAndGetResponseMode, implicitAccessTokenAndPostResponseMode,
	// implicitIdTokenAndPostResponseMode, authorizationCodeWithoutPkce, authorizationCodeWithPkce, clientCredentials,
	// refreshTokenGrant, encryptedAuthorizeResponse, directUserGrant, kerberos, prtGrant, seamlessSso, prtBrokerBased,
	// prtNonBrokerBased, onBehalfOf, samlOnBehalfOf. Use the Prefer: include-unknown-enum-members request header to get the
	// following values from this {evolvable
	// enum}(/graph/best-practices-concept#handling-future-members-in-evolvable-enumerations): authenticationTransfer ,
	// nativeAuth , implicitAccessTokenAndGetResponseMode , implicitIdTokenAndGetResponseMode ,
	// implicitAccessTokenAndPostResponseMode , implicitIdTokenAndPostResponseMode , authorizationCodeWithoutPkce ,
	// authorizationCodeWithPkce , clientCredentials , refreshTokenGrant , encryptedAuthorizeResponse , directUserGrant ,
	// kerberos , prtGrant , seamlessSso , prtBrokerBased , prtNonBrokerBased , onBehalfOf , samlOnBehalfOf.
	AuthenticationProtocol *ProtocolType `json:"authenticationProtocol,omitempty"`

	// This holds the highest level of authentication needed through all the sign-in steps, for sign-in to succeed. Supports
	// $filter (eq, startsWith).
	AuthenticationRequirement nullable.Type[string] `json:"authenticationRequirement,omitempty"`

	// Sources of authentication requirement, such as conditional access, per-user MFA, identity protection, and security
	// defaults.
	AuthenticationRequirementPolicies *[]AuthenticationRequirementPolicy `json:"authenticationRequirementPolicies,omitempty"`

	// The Autonomous System Number (ASN) of the network used by the actor.
	AutonomousSystemNumber nullable.Type[int64] `json:"autonomousSystemNumber,omitempty"`

	// Contains a fully qualified Azure Resource Manager ID of an Azure resource accessed during the sign-in.
	AzureResourceId nullable.Type[string] `json:"azureResourceId,omitempty"`

	// The legacy client used for sign-in activity. For example: Browser, Exchange ActiveSync, Modern clients, IMAP, MAPI,
	// SMTP, or POP. Supports $filter (eq).
	ClientAppUsed nullable.Type[string] `json:"clientAppUsed,omitempty"`

	// Describes the credential type that a user client or service principal provided to Microsoft Entra ID to authenticate
	// itself. You can review this property to track and eliminate less secure credential types or to watch for clients and
	// service principals using anomalous credential types. The possible values are: none, clientSecret, clientAssertion,
	// federatedIdentityCredential, managedIdentity, certificate, unknownFutureValue.
	ClientCredentialType *ClientCredentialType `json:"clientCredentialType,omitempty"`

	// A list that indicates the audience that Conditional Access evaluated during a sign-in event. Supports $filter (eq).
	ConditionalAccessAudiences *[]string `json:"conditionalAccessAudiences,omitempty"`

	// The status of the conditional access policy triggered. Possible values: success, failure, notApplied, or
	// unknownFutureValue. Supports $filter (eq).
	ConditionalAccessStatus *ConditionalAccessStatus `json:"conditionalAccessStatus,omitempty"`

	// The identifier the client sends when sign-in is initiated. This property is used for troubleshooting the
	// corresponding sign-in activity when calling for support. Supports $filter (eq).
	CorrelationId nullable.Type[string] `json:"correlationId,omitempty"`

	// The date and time the sign-in was initiated. The Timestamp type is always in UTC time. For example, midnight UTC on
	// Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $orderby, $filter (eq, le, and ge).
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// Describes the type of cross-tenant access used by the actor to access the resource. Possible values are: none,
	// b2bCollaboration, b2bDirectConnect, microsoftSupport, serviceProvider, unknownFutureValue, passthrough. Use the
	// Prefer: include-unknown-enum-members request header to get the following value or values in this evolvable enum:
	// passthrough. If the sign in didn't cross tenant boundaries, the value is none.
	CrossTenantAccessType *SignInAccessType `json:"crossTenantAccessType,omitempty"`

	// The device information from where the sign-in occurred. Includes information such as deviceId, OS, and browser.
	// Supports $filter (eq, startsWith) on browser and operatingSystem properties.
	DeviceDetail *DeviceDetail `json:"deviceDetail,omitempty"`

	// Contains the identifier of an application's federated identity credential, if a federated identity credential was
	// used to sign in.
	FederatedCredentialId nullable.Type[string] `json:"federatedCredentialId,omitempty"`

	// During a failed sign-in, a user can select a button in the Azure portal to mark the failed event for tenant admins.
	// If a user selects the button to flag the failed sign-in, this value is true.
	FlaggedForReview nullable.Type[bool] `json:"flaggedForReview,omitempty"`

	// The Global Secure Access IP address that the sign-in was initiated from.
	GlobalSecureAccessIPAddress nullable.Type[string] `json:"globalSecureAccessIpAddress,omitempty"`

	// The tenant identifier of the user initiating the sign-in. Not applicable in Managed Identity or service principal
	// sign ins.
	HomeTenantId nullable.Type[string] `json:"homeTenantId,omitempty"`

	// For user sign ins, the identifier of the tenant that the user is a member of. Only populated in cases where the home
	// tenant provides affirmative consent to Microsoft Entra ID to show the tenant content.
	HomeTenantName nullable.Type[string] `json:"homeTenantName,omitempty"`

	// The IP address of the client from where the sign-in occurred. Supports $filter (eq, startsWith).
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The IP address a user used to reach a resource provider, used to determine Conditional Access compliance for some
	// policies. For example, when a user interacts with Exchange Online, the IP address that Microsoft Exchange receives
	// from the user can be recorded here. This value is often null.
	IPAddressFromResourceProvider nullable.Type[string] `json:"ipAddressFromResourceProvider,omitempty"`

	// Indicates the token types that were presented to Microsoft Entra ID to authenticate the actor in the sign in. The
	// possible values are: none, primaryRefreshToken, saml11, saml20, unknownFutureValue, remoteDesktopToken, refreshToken.
	// NOTE Microsoft Entra ID might have also used token types not listed in this enum type to authenticate the actor.
	// Don't infer the lack of a token if it isn't one of the types listed. Use the Prefer: include-unknown-enum-members
	// request header to get the following value or values in this evolvable enum: remoteDesktopToken, refreshToken.
	IncomingTokenType *IncomingTokenType `json:"incomingTokenType,omitempty"`

	// Indicates whether a user sign in is interactive. In interactive sign in, the user provides an authentication factor
	// to Microsoft Entra ID. These factors include passwords, responses to MFA challenges, biometric factors, or QR codes
	// that a user provides to Microsoft Entra ID or an associated app. In non-interactive sign in, the user doesn't provide
	// an authentication factor. Instead, the client app uses a token or code to authenticate or access a resource on behalf
	// of a user. Non-interactive sign ins are commonly used for a client to sign in on a user's behalf in a process
	// transparent to the user.
	IsInteractive nullable.Type[bool] `json:"isInteractive,omitempty"`

	// Shows whether the sign in event was subject to a Microsoft Entra tenant restriction policy.
	IsTenantRestricted nullable.Type[bool] `json:"isTenantRestricted,omitempty"`

	// Indicates whether a user came through Global Secure Access service.
	IsThroughGlobalSecureAccess nullable.Type[bool] `json:"isThroughGlobalSecureAccess,omitempty"`

	// The city, state, and two letter country code from where the sign-in occurred. Supports $filter (eq, startsWith) on
	// city, state, and countryOrRegion properties.
	Location *SignInLocation `json:"location,omitempty"`

	// Contains information about the managed identity used for the sign in, including its type, associated Azure Resource
	// Manager (ARM) resource ID, and federated token information.
	ManagedServiceIdentity *ManagedIdentity `json:"managedServiceIdentity,omitempty"`

	// This property is deprecated.
	MfaDetail *MfaDetail `json:"mfaDetail,omitempty"`

	// The network location details including the type of network used and its names.
	NetworkLocationDetails *[]NetworkLocationDetail `json:"networkLocationDetails,omitempty"`

	// The request identifier of the first request in the authentication sequence. Supports $filter (eq).
	OriginalRequestId nullable.Type[string] `json:"originalRequestId,omitempty"`

	// Transfer method used to initiate a session throughout all subsequent request. The possible values are: none,
	// deviceCodeFlow, authenticationTransfer, unknownFutureValue.
	OriginalTransferMethod *OriginalTransferMethods `json:"originalTransferMethod,omitempty"`

	// Contains information about the Microsoft Entra Private Link policy that is associated with the sign in event.
	PrivateLinkDetails *PrivateLinkDetails `json:"privateLinkDetails,omitempty"`

	// The request processing time in milliseconds in AD STS.
	ProcessingTimeInMilliseconds nullable.Type[int64] `json:"processingTimeInMilliseconds,omitempty"`

	// The name of the resource that the user signed in to. Supports $filter (eq).
	ResourceDisplayName nullable.Type[string] `json:"resourceDisplayName,omitempty"`

	// The identifier of the resource that the user signed in to. Supports $filter (eq).
	ResourceId nullable.Type[string] `json:"resourceId,omitempty"`

	// The identifier of the owner of the resource. Supports $filter (eq).
	ResourceOwnerTenantId nullable.Type[string] `json:"resourceOwnerTenantId,omitempty"`

	// The identifier of the service principal representing the target resource in the sign-in event.
	ResourceServicePrincipalId nullable.Type[string] `json:"resourceServicePrincipalId,omitempty"`

	// The tenant identifier of the resource referenced in the sign in.
	ResourceTenantId nullable.Type[string] `json:"resourceTenantId,omitempty"`

	// The reason behind a specific state of a risky user, sign-in, or a risk event. The possible values are none,
	// adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset,
	// adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser,
	// adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue,
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection,
	// userChangedPasswordOnPremises, adminDismissedRiskForSignIn, adminConfirmedAccountSafe. Use the Prefer:
	// include-unknown-enum-members request header to get the following value or values in this evolvable enum:
	// adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal, m365DAdminDismissedDetection,
	// userChangedPasswordOnPremises, adminDismissedRiskForSignIn, adminConfirmedAccountSafe.The value none means that
	// Microsoft Entra risk detection hasn't flagged the user or the sign-in as a risky event so far. Supports $filter (eq).
	// Note: Details for this property are only available for Microsoft Entra ID P2 customers. All other customers are
	// returned hidden.
	RiskDetail *RiskDetail `json:"riskDetail,omitempty"`

	// The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress,
	// maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials,
	// investigationsThreatIntelligence, generic, or unknownFutureValue. Supports $filter (eq, startsWith).
	RiskEventTypesv2 *[]string `json:"riskEventTypes_v2,omitempty"`

	// The aggregated risk level. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value hidden
	// means the user or sign-in wasn't enabled for Microsoft Entra ID Protection. Supports $filter (eq). Note: Details for
	// this property are only available for Microsoft Entra ID P2 customers. All other customers are returned hidden.
	RiskLevelAggregated *RiskLevel `json:"riskLevelAggregated,omitempty"`

	// The risk level during sign-in. Possible values: none, low, medium, high, hidden, or unknownFutureValue. The value
	// hidden means the user or sign-in wasn't enabled for Microsoft Entra ID Protection. Supports $filter (eq). Note:
	// Details for this property are only available for Microsoft Entra ID P2 customers. All other customers are returned
	// hidden.
	RiskLevelDuringSignIn *RiskLevel `json:"riskLevelDuringSignIn,omitempty"`

	// The risk state of a risky user, sign-in, or a risk event. Possible values: none, confirmedSafe, remediated,
	// dismissed, atRisk, confirmedCompromised, or unknownFutureValue. Supports $filter (eq).
	RiskState *RiskState `json:"riskState,omitempty"`

	// The unique identifier of the key credential used by the service principal to authenticate.
	ServicePrincipalCredentialKeyId nullable.Type[string] `json:"servicePrincipalCredentialKeyId,omitempty"`

	// The certificate thumbprint of the certificate used by the service principal to authenticate.
	ServicePrincipalCredentialThumbprint nullable.Type[string] `json:"servicePrincipalCredentialThumbprint,omitempty"`

	// The application identifier used for sign-in. This field is populated when you're signing in using an application.
	// Supports $filter (eq, startsWith).
	ServicePrincipalId *string `json:"servicePrincipalId,omitempty"`

	// The application name used for sign-in. This field is populated when you're signing in using an application. Supports
	// $filter (eq, startsWith).
	ServicePrincipalName nullable.Type[string] `json:"servicePrincipalName,omitempty"`

	// Identifier of the session that was generated during the sign-in.
	SessionId nullable.Type[string] `json:"sessionId,omitempty"`

	// Any conditional access session management policies that were applied during the sign-in event.
	SessionLifetimePolicies *[]SessionLifetimePolicy `json:"sessionLifetimePolicies,omitempty"`

	// Indicates the category of sign in that the event represents. For user sign ins, the category can be interactiveUser
	// or nonInteractiveUser and corresponds to the value for the isInteractive property on the signin resource. For managed
	// identity sign ins, the category is managedIdentity. For service principal sign-ins, the category is servicePrincipal.
	// Possible values are: interactiveUser, nonInteractiveUser, servicePrincipal, managedIdentity, unknownFutureValue.
	// Supports $filter (eq, ne). NOTE: Only interactive sign-ins are returned unless you set an explicit filter. For
	// example, the filter for getting non-interactive sign-ins is
	// https://graph.microsoft.com/beta/auditLogs/signIns?&$filter=signInEventTypes/any(t: t eq 'nonInteractiveUser').
	SignInEventTypes *[]string `json:"signInEventTypes,omitempty"`

	// The identification that the user provided to sign in. It can be the userPrincipalName, but is also populated when a
	// user signs in using other identifiers.
	SignInIdentifier nullable.Type[string] `json:"signInIdentifier,omitempty"`

	// The type of sign in identifier. Possible values are: userPrincipalName, phoneNumber, proxyAddress, qrCode,
	// onPremisesUserPrincipalName, unknownFutureValue.
	SignInIdentifierType *SignInIdentifierType `json:"signInIdentifierType,omitempty"`

	// Token protection creates a cryptographically secure tie between the token and the device it's issued to. This field
	// indicates whether the signin token was bound to the device or not. The possible values are: none, bound, unbound,
	// unknownFutureValue.
	SignInTokenProtectionStatus *TokenProtectionStatus `json:"signInTokenProtectionStatus,omitempty"`

	// The sign-in status. Includes the error code and description of the error (for a sign-in failure). Supports $filter
	// (eq) on errorCode property.
	Status *SignInStatus `json:"status,omitempty"`

	// The name of the identity provider. For example, sts.microsoft.com. Supports $filter (eq).
	TokenIssuerName nullable.Type[string] `json:"tokenIssuerName,omitempty"`

	// The type of identity provider. The possible values are: AzureAD, ADFederationServices, UnknownFutureValue,
	// AzureADBackupAuth, ADFederationServicesMFAAdapter, NPSExtension. Use the Prefer: include-unknown-enum-members request
	// header to get the following values in this evolvable enum: AzureADBackupAuth , ADFederationServicesMFAAdapter ,
	// NPSExtension.
	TokenIssuerType *TokenIssuerType `json:"tokenIssuerType,omitempty"`

	TokenProtectionStatusDetails *TokenProtectionStatusDetails `json:"tokenProtectionStatusDetails,omitempty"`

	// A unique base64-encoded request identifier used to track tokens issued by Microsoft Entra ID as they're redeemed at
	// resource providers.
	UniqueTokenIdentifier nullable.Type[string] `json:"uniqueTokenIdentifier,omitempty"`

	// The user agent information related to sign-in. Supports $filter (eq, startsWith).
	UserAgent nullable.Type[string] `json:"userAgent,omitempty"`

	// The display name of the user. Supports $filter (eq, startsWith).
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// The identifier of the user. Supports $filter (eq).
	UserId *string `json:"userId,omitempty"`

	// User principal name of the user that initiated the sign-in. This value is always in lowercase. For guest users whose
	// values in the user object typically contain #EXT# before the domain part, this property stores the value in both
	// lowercase and the 'true' format. For example, while the user object stores AdeleVance_fabrikam.com#EXT#@contoso.com,
	// the sign-in logs store adelevance@fabrikam.com. Supports $filter (eq, startsWith).
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Identifies whether the user is a member or guest in the tenant. Possible values are: member, guest,
	// unknownFutureValue.
	UserType *SignInUserType `json:"userType,omitempty"`

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

func (s SignIn) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SignIn{}

func (s SignIn) MarshalJSON() ([]byte, error) {
	type wrapper SignIn
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SignIn: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SignIn: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.signIn"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SignIn: %+v", err)
	}

	return encoded, nil
}
