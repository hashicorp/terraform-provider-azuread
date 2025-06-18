package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = PolicyRoot{}

type PolicyRoot struct {
	// The policy that contains directory-level access review settings.
	AccessReviewPolicy *AccessReviewPolicy `json:"accessReviewPolicy,omitempty"`

	// The policy that controls the idle time-out for web sessions for applications.
	ActivityBasedTimeoutPolicies *[]ActivityBasedTimeoutPolicy `json:"activityBasedTimeoutPolicies,omitempty"`

	// The policy by which consent requests are created and managed for the entire tenant.
	AdminConsentRequestPolicy *AdminConsentRequestPolicy `json:"adminConsentRequestPolicy,omitempty"`

	// The policies that enforce app management restrictions for specific applications and service principals, overriding
	// the defaultAppManagementPolicy.
	AppManagementPolicies *[]AppManagementPolicy `json:"appManagementPolicies,omitempty"`

	// The policy configuration of the self-service sign-up experience of guests.
	AuthenticationFlowsPolicy *AuthenticationFlowsPolicy `json:"authenticationFlowsPolicy,omitempty"`

	// The authentication methods and the users that are allowed to use them to sign in and perform multifactor
	// authentication (MFA) in Microsoft Entra ID.
	AuthenticationMethodsPolicy *AuthenticationMethodsPolicy `json:"authenticationMethodsPolicy,omitempty"`

	// The authentication method combinations that are to be used in scenarios defined by Microsoft Entra Conditional
	// Access.
	AuthenticationStrengthPolicies *[]AuthenticationStrengthPolicy `json:"authenticationStrengthPolicies,omitempty"`

	// The policy that controls Microsoft Entra authorization settings.
	AuthorizationPolicy *[]AuthorizationPolicy `json:"authorizationPolicy,omitempty"`

	// The Azure AD B2C policies that define how end users register via local accounts.
	B2cAuthenticationMethodsPolicy *B2cAuthenticationMethodsPolicy `json:"b2cAuthenticationMethodsPolicy,omitempty"`

	// The claim-mapping policies for WS-Fed, SAML, OAuth 2.0, and OpenID Connect protocols, for tokens issued to a specific
	// application.
	ClaimsMappingPolicies *[]ClaimsMappingPolicy `json:"claimsMappingPolicies,omitempty"`

	// The custom rules that define an access scenario when interacting with external Microsoft Entra tenants.
	CrossTenantAccessPolicy *CrossTenantAccessPolicy `json:"crossTenantAccessPolicy,omitempty"`

	// The tenant-wide policy that enforces app management restrictions for all applications and service principals.
	DefaultAppManagementPolicy *TenantAppManagementPolicy `json:"defaultAppManagementPolicy,omitempty"`

	// Represents the policy scope that controls quota restrictions, additional authentication, and authorization policies
	// to register device identities to your organization.
	DeviceRegistrationPolicy *DeviceRegistrationPolicy `json:"deviceRegistrationPolicy,omitempty"`

	DirectoryRoleAccessReviewPolicy *DirectoryRoleAccessReviewPolicy `json:"directoryRoleAccessReviewPolicy,omitempty"`

	// Represents the tenant-wide policy that controls whether guests can leave a Microsoft Entra tenant via self-service
	// controls.
	ExternalIdentitiesPolicy *ExternalIdentitiesPolicy `json:"externalIdentitiesPolicy,omitempty"`

	// The feature rollout policy associated with a directory object.
	FeatureRolloutPolicies *[]FeatureRolloutPolicy `json:"featureRolloutPolicies,omitempty"`

	// Represents a policy to control enabling or disabling validation of federation authentication tokens.
	FederatedTokenValidationPolicy *FederatedTokenValidationPolicy `json:"federatedTokenValidationPolicy,omitempty"`

	// The policy to control Microsoft Entra authentication behavior for federated users.
	HomeRealmDiscoveryPolicies *[]HomeRealmDiscoveryPolicy `json:"homeRealmDiscoveryPolicies,omitempty"`

	// The policy that represents the security defaults that protect against common attacks.
	IdentitySecurityDefaultsEnforcementPolicy *IdentitySecurityDefaultsEnforcementPolicy `json:"identitySecurityDefaultsEnforcementPolicy,omitempty"`

	// The policy that defines autoenrollment configuration for a mobility management (MDM or MAM) application.
	MobileAppManagementPolicies *[]MobilityManagementPolicy `json:"mobileAppManagementPolicies,omitempty"`

	MobileDeviceManagementPolicies *[]MobilityManagementPolicy `json:"mobileDeviceManagementPolicies,omitempty"`

	// The policy that specifies the conditions under which consent can be granted.
	PermissionGrantPolicies *[]PermissionGrantPolicy `json:"permissionGrantPolicies,omitempty"`

	// Policies that specify the conditions under which consent can be granted to a specific application.
	PermissionGrantPreApprovalPolicies *[]PermissionGrantPreApprovalPolicy `json:"permissionGrantPreApprovalPolicies,omitempty"`

	// Represents the role management policies.
	RoleManagementPolicies *[]UnifiedRoleManagementPolicy `json:"roleManagementPolicies,omitempty"`

	// Represents the role management policy assignments.
	RoleManagementPolicyAssignments *[]UnifiedRoleManagementPolicyAssignment `json:"roleManagementPolicyAssignments,omitempty"`

	ServicePrincipalCreationPolicies *[]ServicePrincipalCreationPolicy `json:"servicePrincipalCreationPolicies,omitempty"`

	// The policy that specifies the characteristics of SAML tokens issued by Microsoft Entra ID.
	TokenIssuancePolicies *[]TokenIssuancePolicy `json:"tokenIssuancePolicies,omitempty"`

	// The policy that controls the lifetime of a JWT access token, an ID token, or a SAML 1.1/2.0 token issued by Microsoft
	// Entra ID.
	TokenLifetimePolicies *[]TokenLifetimePolicy `json:"tokenLifetimePolicies,omitempty"`

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

func (s PolicyRoot) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PolicyRoot{}

func (s PolicyRoot) MarshalJSON() ([]byte, error) {
	type wrapper PolicyRoot
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PolicyRoot: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PolicyRoot: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.policyRoot"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PolicyRoot: %+v", err)
	}

	return encoded, nil
}
